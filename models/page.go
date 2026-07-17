package models

type PageData struct {
	Title   string
	Artists []Artist
}

//ArtistPageData holds everything needed by artist.html
type ArtistPageData struct{
	Artist Artist
	Location []string
	Dates []string
	Relation map[string][]string
}