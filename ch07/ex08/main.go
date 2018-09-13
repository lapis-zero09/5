package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
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

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	ms := &multipleSort{}
	fmt.Println("---------------------org----------------------------------")
	printTracks(tracks)

	fmt.Println()
	fmt.Println("---------------------sort by title----------------------------------")
	ms.Set(customSort{tracks, byTitle})
	sort.Sort(ms)
	printTracks(tracks)

	fmt.Println()
	fmt.Println("---------------------sort by year----------------------------------")
	ms.Set(customSort{tracks, byYear})
	sort.Sort(ms)
	printTracks(tracks)

	fmt.Println()
	fmt.Println("---------------------sort by length----------------------------------")
	ms.Set(customSort{tracks, byLength})
	sort.Sort(ms)
	printTracks(tracks)

	fmt.Println()
	fmt.Println("---------------------sort by artist----------------------------------")
	ms.Set(customSort{tracks, byArtist})
	sort.Sort(ms)
	printTracks(tracks)
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}
