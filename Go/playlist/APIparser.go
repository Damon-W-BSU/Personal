package playlist

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Creates a Playlist struct from a Spotify playlist JSON file
// by moving data from unmarshalled structs into simpler format
func NewPlaylistFromJSON(JSON []byte, token string) (*Playlist, error) {

	// Unmarshall JSON data
	var JSONPL playlistJSON
	err := json.Unmarshal(JSON, &JSONPL)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	P := &Playlist{name: JSONPL.Name, total: JSONPL.Tracks.Total, url: JSONPL.Link.Url} // playlist to return

	// API call to get track details
	JSONtracks, err := P.getPlaylistTracks(token)
	if err != nil {
		return nil, err
	}

	var tracks *tracksJSON

	// NOTE: Nested range loops work because data isn't being read from playlist

	for _, jchunk := range JSONtracks {
		json.Unmarshal(jchunk, &tracks)
		var toAdd *Track // current track being added
		// traverse JSON formatted tracks
		for _, item := range tracks.Items {
			// create and append new track to output
			toAdd = newTrack(item.Track.Name)
			P.tracks = append(P.tracks, toAdd)
			// append name of each artist to new track
			for i, name := range item.Track.Artists {
				toAdd.artists = append(toAdd.artists, Artist{name: name.Name, href: item.Track.Artists[i].Href})
			}
		}
	}

	// API call to get artist details
	JSONartists, err := P.getArtistGenres(token)
	if err != nil {
		return nil, err
	}

	var artists trackJSON

	// NOTE: nested range loops were NOT working because they are
	// only making copies of the data seen in the playlist

	for _, jchunk := range JSONartists {
		json.Unmarshal(jchunk, &artists)

		for i := range P.tracks {
			for k := range P.tracks[i].artists {
				artist := &P.tracks[i].artists[k]
				for _, match := range artists.Artists {
					if artist.name == match.Name {
						artist.genres = append(artist.genres, match.Genres...)
						break
					}
				}
			}
		}

	}

	return P, nil
}

func ParseLink(link string) string {

	id := link[strings.LastIndex(link, "/")+1:]
	if strings.Contains(id, "?") {
		id = id[:strings.Index(id, "?")]
	}
	return id
}

// basic playlist info
type playlistJSON struct {
	Tracks      tracksJSON       `json:"tracks"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Link        external_urlJSON `json:"external_urls"`
}

type external_urlJSON struct {
	Url string `json:"spotify"`
}

// list of tracks
type tracksJSON struct {
	Items []itemJSON `json:"items"`
	Total int        `json:"total"`
}

// item containing track info
type itemJSON struct {
	Track trackJSON `json:"track"`
}

// actual track information
type trackJSON struct {
	Artists []artistJSON `json:"artists"`
	Name    string       `json:"name"`
}

// artist in artist list
type artistJSON struct {
	Name   string   `json:"name"`
	Href   string   `json:"href"`
	Genres []string `json:"genres"`
}

/**
	UNUSED METHODS
**/

// Creates a new Playlist with tracks from formatted TXT file
func NewPlaylistFromTXT(name string, filename string) *Playlist {
	P := &Playlist{name: name}

	// read track data from file
	data, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var next string   // next char
	var start int = 0 // starting index of each field
	var raw []byte    // raw byte data
	var track *Track  // track being added

	// Collect tracks and artists
	for i := range data {
		next = string(data[i])
		if next == "|" {
			raw = data[start : i-1]
			track = newTrack(string(raw))
			P.tracks = append(P.tracks, track)
		} else if next == "," || next == "\n" {
			raw = data[start+1 : i]
			track.artists = append(track.artists, Artist{name: string(raw)})
		} else {
			continue
		}

		i++
		start = i
	}

	return P
}
