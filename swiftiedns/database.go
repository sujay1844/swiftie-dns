package swiftiedns

import (
	"encoding/csv"
	"io"
)

type Song struct {
	ID        string
	Name      string
	AlbumID   string
	AlbumName string
	AlbumPath string
	Lyrics    string
}

var Songs []Song

func InitDB(reader io.Reader) {
	r := csv.NewReader(reader)
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
			AlbumPath: record[4],
			Lyrics:    record[5],
		}
		Songs = append(Songs, song)
	}
}
