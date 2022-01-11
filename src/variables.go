package groupie

import "html/template"

const UrlArt = "https://groupietrackers.herokuapp.com/api/artists"
const UrlRel = "https://groupietrackers.herokuapp.com/api/relation"

var Templates, TempErr = template.ParseFiles("templates/mainPage.html", "templates/artPage.html", "templates/errors.html", "templates/searchPage.html", "templates/filterPage.html")

// var templates, tempErr = template.ParseGlob("templates/*.html")

type DataArt struct {
	ID             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	Locations      string              `json:"locations"`
	DatesLocations map[string][]string `json:"datesLocations"`

	ConcertDates string `json:"concertDates"`
}

type Index struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type ArtOutput struct {
	A DataArt
	R Index
}

type SearchOutput struct {
	Art []DataArt
}

type Filter struct {
	CreationDate        []int
	FirstAlbum          []int
	NumberOfMembers     []int
	LocationsOfConcerts string
}

// var Artists []DataArt

var SearchArtists struct {
	Artists   []DataArt
	Relations []Index `json:"index"`
}
