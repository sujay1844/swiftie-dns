package swiftiedns

import "strings"

func getLyrics(name string) string {
	for _, song := range Songs {
		if song.Name == name {
			return song.Lyrics
		}
	}
	return "Song not found"
}

func getResponses(name string) ([]string, error) {
	lyrics := getLyrics(name)
	return strings.Split(lyrics, "\n"), nil
}
