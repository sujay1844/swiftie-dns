package swiftiedns

import "strings"

func getLyrics(songs Songs, name string) string {
	for _, song := range songs {
		if song.Name == name {
			return song.Lyrics
		}
	}
	return "Song not found"
}

func getResponses(songs Songs, name string) ([]string, error) {
	lyrics := getLyrics(songs, name)
	return strings.Split(lyrics, "\n"), nil
}
