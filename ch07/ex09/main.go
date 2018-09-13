package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
	"time"
)

type multipleSort struct {
	prevSort, sort sort.Interface
}

func (s *multipleSort) Set(si sort.Interface) {
	s.prevSort = s.sort
	s.sort = si
}

func (s *multipleSort) Len() int {
	return s.sort.Len()
}

func (s *multipleSort) Less(i, j int) bool {
	if !equals(s.sort, i, j) {
		return s.sort.Less(i, j)
	}
	if s.prevSort != nil && !equals(s.prevSort, i, j) {
		return s.prevSort.Less(i, j)
	}
	return false
}

func (s *multipleSort) Swap(i, j int) {
	s.sort.Swap(i, j)
}

func equals(s sort.Interface, i, j int) bool {
	return !s.Less(i, j) && !s.Less(j, i)
}

//

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func byTitle(x, y *Track) bool  { return x.Title < y.Title }
func byArtist(x, y *Track) bool { return x.Artist < y.Artist }
func byYear(x, y *Track) bool   { return x.Year < y.Year }
func byLength(x, y *Track) bool { return x.Length < y.Length }
func byAlbum(x, y *Track) bool  { return x.Album < y.Album }

var lessMap = map[string]func(x, y *Track) bool{
	"title":  byTitle,
	"artist": byArtist,
	"year":   byYear,
	"length": byLength,
	"album":  byAlbum,
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	ms := &multipleSort{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := r.FormValue("sortedby")
		if len(c) > 0 {
			if f, ok := lessMap[c]; ok {
				ms.Set(customSort{tracks, f})
				sort.Sort(ms)
			}
		}
		printTracks(w, tracks)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

var tmpl = template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
	<title>Tracks</title>
	<link href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
</head>
<body>
<div class="jumbotron jumbotron-fluid">
  <div class="container">
    <h1>Tracks</h1>
	<table class="table">
		<thead>
			<tr>
				<th><a href="?sortedby=title">Title</a></th>
				<th><a href="?sortedby=artist">Artist</a></th>
				<th><a href="?sortedby=album">Album</a></th>
				<th><a href="?sortedby=year">Year</a></th>
				<th><a href="?sortedby=length">Length</a></th>
			</tr>
		</thead>
		<tbody>
			{{ range . }}
				<tr>
					<td>{{ .Title }}</td>
					<td>{{ .Artist }}</td>
					<td>{{ .Album }}</td>
					<td>{{ .Year }}</td>
					<td>{{ .Length }}</td>
				</tr>
			{{ end }}
		</tbody>
	</table>
  </div>
</div>
</body>
</html>
`))

func printTracks(w io.Writer, tracks []*Track) {
	tmpl.Execute(w, tracks)
}
