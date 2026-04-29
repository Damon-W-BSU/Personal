package playlist

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetAccessToken() string {

	// load local client_id and client_secret
	err := godotenv.Load()

	// Encode POST values for API call
	var values url.Values = make(url.Values)
	values.Add("grant_type", "client_credentials")
	values.Add("client_id", os.Getenv("client_id"))
	values.Add("client_secret", os.Getenv("client_secret"))
	body := strings.NewReader(values.Encode())

	// Call Spotify API for access token
	response, err := http.Post("https://accounts.spotify.com/api/token", "application/x-www-form-urlencoded", body)
	if err != nil {
		return err.Error()
	}
	defer response.Body.Close()

	// Get response
	raw, err := io.ReadAll(response.Body)
	if err != nil {
		return err.Error()
	}

	// Parse only actual access token
	var start, end int
	for i := range raw {

		if string(raw[i]) == ":" && string(raw[i+1]) == "\"" {
			i += 2
			start = i
		}

		if string(raw[i]) == "\"" && string(raw[i+1]) == "," {
			end = i
			break
		}
	}

	// Return access token
	out := string(raw[start:end])
	return out
}

func GetPlaylist(link string, token string) ([]byte, error) {

	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s?fields=description%%2C+external_urls%%28spotify%%29%%2C+name%%2C+tracks%%28total%%2C+items%%28track%%28name%%2C+artists%%28name%%29%%29%%29%%29", link)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	var head http.Header = make(http.Header)
	head.Add("Authorization", "Bearer "+token)
	req.Header = head

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

func (P *Playlist) getPlaylistTracks(token string) ([][]byte, error) {

	id := ParseLink(P.url)

	var jchunks [][]byte

	for i := 0; i < P.total; i += 100 {

		url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks?fields=items%%28track%%28artists%%28href%%2C+name%%29%%2C+name%%29%%29&offset=%d", id, i)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		var head http.Header = make(http.Header)
		head.Add("Authorization", "Bearer "+token)
		req.Header = head

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		raw, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		jchunks = append(jchunks, raw)

	}

	return jchunks, nil
}

func (P *Playlist) getArtistGenres(token string) ([][]byte, error) {

	links := [][]string{}
	var jchunks [][]byte
	var found = make(map[string]int)

	for _, track := range P.tracks {
		for _, artist := range track.artists {
			if found[artist.href] == 0 {
				links = append(links, []string{ParseLink(artist.href)})
				found[artist.href] = 1
			}
		}
	}

	for i := 0; i < len(links); i += 50 {

		var sb strings.Builder
		sb.WriteString("https://api.spotify.com/v1/artists?ids=")

		for _, id := range links[i : i+50] {
			if len(id) != 0 {
				sb.WriteString(id[0] + "%2C")
			}
		}

		s, _ := strings.CutSuffix(sb.String(), "%2C")

		req, err := http.NewRequest("GET", s, nil)
		if err != nil {
			return nil, err
		}

		var head http.Header = make(http.Header)
		head.Add("Authorization", "Bearer "+token)
		req.Header = head

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		raw, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		jchunks = append(jchunks, raw)
		// sorting.DumpToTxt(raw)
	}

	return jchunks, nil
}

/**
	UNUSED METHODS
**/

func GAG(token string, href string) ([]string, error) {

	url := href
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var head http.Header = make(http.Header)
	head.Add("Authorization", "Bearer "+token)
	req.Header = head

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	_, data, _ := strings.Cut(string(raw), "\"genres\":[")
	data, _, _ = strings.Cut(data, "]")
	data = strings.ReplaceAll(data, "\"", "")
	data += ","

	raw = []byte(data)
	start := 0
	genres := []string{}

	for i, char := range raw {
		if string(char) == "," {
			genres = append(genres, string(raw[start:i]))
			start = i + 1
		}
	}

	fmt.Println(genres)

	return genres, nil
}
