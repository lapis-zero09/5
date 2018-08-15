package xkcd

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

const endPoint = "https://xkcd.com"

type XKCD struct {
	Num        int    `gorm:"primary_key" json:"num"`
	Year       string `json:"year"`
	Month      string `json:"month"`
	Day        string `json:"day"`
	Link       string `json:"link"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
}

func randomString() string {
	var n uint64
	binary.Read(cryptoRand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

func SampleXKCDs() []XKCD {
	random := rand.New(rand.NewSource(1))
	random.Seed(5)
	xkcds := make([]XKCD, 0, 100)
	for i := 0; i < 100; i++ {
		xkcds = append(xkcds,
			XKCD{
				Num:        i,
				Transcript: randomString(),
			})
	}
	return xkcds
}

func InitXKCDTable() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./xkcd.db")
	// db.LogMode(true)

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&XKCD{}) {
		db.CreateTable(&XKCD{})
		db.Set("gorm.table_options", "ENGINE=InnoDB").CreateTable(&XKCD{})
		// insert seeds
		// for _, xkcd := range SampleXKCDs() {
		// 	db.Create(&xkcd)
		// }
	}

	return db
}
