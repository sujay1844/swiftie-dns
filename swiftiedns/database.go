package swiftiedns

import (
	"encoding/csv"
	"os"
)

type Song struct {
	ID        string
	Name      string
	AlbumID   string
	AlbumName string
	Lyrics    string
}

var Songs []Song

func InitDB(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	Songs = make([]Song, 0, len(records))
	for i, record := range records {
		if i == 0 {
			continue
		}
		song := Song{
			ID:        record[0],
			Name:      record[1],
			AlbumID:   record[2],
			AlbumName: record[3],
			Lyrics:    record[4],
		}
		Songs = append(Songs, song)
	}
}
