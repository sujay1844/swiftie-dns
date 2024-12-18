package swiftiedns

import (
	"encoding/csv"
	"fmt"
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

type Songs []Song

func (s Songs) Len() int {
	return len(s)
}

func (s Songs) String(i int) string {
	song := s[i]
	return fmt.Sprintf("%s %s", song.Name, song.AlbumName)
}

func InitDB(reader io.Reader) (Songs, error) {
	r := csv.NewReader(reader)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	songs := make([]Song, 0, len(records))
	for i, record := range records {
		if i == 0 {
			continue
		}
		if len(record) != 6 {
			return nil, fmt.Errorf("record on line %d: expected 6 fields, got %d", i+1, len(record))
		}
		song := Song{
			ID:        record[0],
			Name:      record[1],
			AlbumID:   record[2],
			AlbumName: record[3],
			AlbumPath: record[4],
			Lyrics:    record[5],
		}
		songs = append(songs, song)
	}
	return songs, nil
}
