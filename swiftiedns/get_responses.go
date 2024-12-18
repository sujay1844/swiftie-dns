package swiftiedns

import (
	"fmt"
	"strings"

	"github.com/sahilm/fuzzy"
)

var ErrNoSongFound = fmt.Errorf("Song not found")

func getLyrics(songs Songs, name string) (string, error) {
	results := fuzzy.FindFrom(name, songs)
	if len(results) == 0 {
		return "", ErrNoSongFound
	}
	song := songs[results[0].Index]
	return song.Lyrics, nil
}

func getResponses(songs Songs, name string) ([]string, error) {
	lyrics, err := getLyrics(songs, name)
	if err != nil {
		return nil, err
	}
	return strings.Split(lyrics, "\n"), nil
}
