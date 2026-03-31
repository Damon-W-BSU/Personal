package playlist

import (
	"fmt"
	"strings"
)

// Represents a music track
type Track struct {
	name    string
	artists []Artist
}

// Represents a musical artist
type Artist struct {
	name   string
	genres []string
	href   string
}

// List containing music tracks
type Playlist struct {
	name   string
	url    string
	total  int // number of tracks
	tracks []*Track
}

func New(url string) *Playlist {

	token := GetAccessToken()
	id := ParseLink(url)
	JSON, err := GetPlaylist(id, token)
	if err != nil {
		return nil
	}
	P, err := NewPlaylistFromJSON(JSON, token)
	if err != nil {
		return nil
	}

	return P
}

// Creates a new track
func newTrack(name string) *Track {
	track := &Track{name: name}
	return track
}

// Sorts tracks alphabetically
func (P *Playlist) SortByName() {
	var lowest *Track
	var lowIndex int
	for i := range P.tracks {
		lowest = P.tracks[i]
		for j := i + 1; j < len(P.tracks); j++ {
			if strings.ToLower(P.tracks[j].name) < strings.ToLower(lowest.name) {
				lowest = P.tracks[j]
				lowIndex = j
			}
		}
		temp := P.tracks[i]
		P.tracks[i] = P.tracks[lowIndex]
		P.tracks[lowIndex] = temp
	}
}

// Displays
func (P *Playlist) ShowPlaylistStats() {

	// various counters to record stats
	var artistCount, genreCount, soloCount, collabCount, remixCount int

	// tallies how many times an artist or genre appears
	var artistAppearances = make(map[string]int)
	var genreAppearances = make(map[string]int)

	// genres currently associated with track
	var trackGenreList = make(map[string]struct{})

	// iterate through each track
	for _, curr := range P.tracks {

		// reset tracked genres
		for k := range trackGenreList {
			delete(trackGenreList, k)
		}

		// iterate though each artist
		for _, artist := range curr.artists {
			if artistAppearances[artist.name] == 0 {
				artistCount++
			}
			artistAppearances[artist.name]++

			// track genre appearances
			for _, genre := range artist.genres {

				// check if track has added genre already
				_, ok := trackGenreList[genre]

				// add genre and update count
				if !ok {
					trackGenreList[genre] = struct{}{}
					if genreAppearances[genre] == 0 {
						genreCount++
					}
					genreAppearances[genre]++
				}
			}
		}

		// update solo/collab/remix count
		if len(curr.artists) == 1 {
			soloCount++
		} else {
			collabCount++
		}
		if strings.Contains(strings.ToLower(curr.name), "remix") {
			remixCount++
		}
	}

	// Display
	fmt.Println("******************")
	fmt.Printf("Stats for %s\n", P.name)
	fmt.Println("******************")
	fmt.Printf("\nTracks: %d\nArtists: %d\nSolo tracks: %d\nCollabs: %d\nRemixes: %d\n\n",
		P.total,
		artistCount,
		soloCount,
		collabCount,
		remixCount)
	fmt.Println("\n# of artist appearances:")

	var highArtist string
	var highCount int
	var rank int

	for range artistCount {
		rank++
		highArtist = "" // set to high value string
		highCount = -1

		for artist, count := range artistAppearances {
			if count > highCount ||
				(count == highCount && strings.ToLower(artist) < strings.ToLower(highArtist)) {

				highArtist = artist
				highCount = count
			}
		}
		fmt.Printf("%3d. %-25s %d\n", rank, highArtist, artistAppearances[highArtist])
		delete(artistAppearances, highArtist)
	}

	fmt.Println("\n# of genre appearances:")

	var highGenre string
	highCount = 0
	rank = 0

	for range genreCount {
		rank++
		highGenre = ""
		highCount = -1

		for genre, count := range genreAppearances {
			if count > highCount ||
				(count == highCount && strings.ToLower(genre) < strings.ToLower(highGenre)) {

				highGenre = genre
				highCount = count
			}
		}
		fmt.Printf("%3d. %-25s %d\n", rank, highGenre, genreAppearances[highGenre])
		delete(genreAppearances, highGenre)
	}

}

// String representation of Playlist
func (P *Playlist) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s\n%s\n", P.name, P.url))
	for i := range P.tracks {
		sb.WriteString(fmt.Sprintf("%d. %s", i+1, P.tracks[i]))
	}

	return sb.String()
}

// String representation of a track
func (T *Track) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s - ", T.name))
	for i := 0; i < len(T.artists)-1; i++ {
		sb.WriteString(fmt.Sprintf("%s, ", T.artists[i].name))
	}
	sb.WriteString(fmt.Sprintf("%s\n", T.artists[len(T.artists)-1].name))
	return sb.String()
}
