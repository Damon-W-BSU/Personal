package playlist

type Playlist struct {
	name string
	url  string
	href string
	id   string
	// may add a field for user later if that info is needed to be stored here
	tracklist []Track
}
