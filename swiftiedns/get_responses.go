package swiftiedns

import (
	"fmt"
	"strings"

	"github.com/sahilm/fuzzy"
)

var ErrNoSongFound = fmt.Errorf("Song not found")

func getSong(songs Songs, name string) (Song, error) {
	results := fuzzy.FindFrom(name, songs)
	if len(results) == 0 {
		return Song{}, ErrNoSongFound
	}
	song := songs[results[0].Index]
	return song, nil
}

func getResponses(songs Songs, name string) ([]string, error) {
	song, err := getSong(songs, name)
	if err != nil {
		return nil, err
	}
	res := []string{"Name: " + song.Name, "Album: " + song.AlbumName, ""}
	res = append(res, strings.Split(song.Lyrics, "\n")...)
	return res, nil
}
