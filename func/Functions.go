package groupie

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const UrlArt = "https://groupietrackers.herokuapp.com/api/artists"
const UrlRel = "https://groupietrackers.herokuapp.com/api/relation"

var templates, tempErr = template.ParseFiles("templates/mainPage.html", "templates/artPage.html", "templates/errors.html", "templates/searchPage.html")

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

// var Artists []DataArt

var SearchArtists struct {
	Artists   []DataArt
	Relations []Index `json:"index"`
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	// err1 := Unmarshal(urlArt, &SearchArtists.Artists)
	// err2 := Unmarshal(urlRel, &SearchArtists)

	// if err1 != nil && err2 != nil {
	// 	fmt.Println("ERROR")
	// 	return
	// }

	if r.URL.Path != "/" {
		w.WriteHeader(404)
		templates.ExecuteTemplate(w, "errors.html", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		w.WriteHeader(405)
		templates.ExecuteTemplate(w, "errors.html", http.StatusMethodNotAllowed)
		return
	}

	templates.ExecuteTemplate(w, "mainPage.html", SearchArtists)
}

// WorkPage func
func ArtPage(w http.ResponseWriter, r *http.Request) {

	ind, err := strconv.Atoi(r.RequestURI[9:])

	if err != nil {
		w.WriteHeader(400)
		templates.ExecuteTemplate(w, "errors.html", http.StatusBadRequest)
		return
	}

	if !isValidId(ind) {
		w.WriteHeader(404)
		templates.ExecuteTemplate(w, "errors.html", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		w.WriteHeader(405)
		templates.ExecuteTemplate(w, "errors.html", http.StatusMethodNotAllowed)
		return
	}

	res := &ArtOutput{}
	res.A = SearchArtists.Artists[ind-1]
	res.R = SearchArtists.Relations[ind-1]
	templates.ExecuteTemplate(w, "artPage.html", res)
}

func SearchPage(w http.ResponseWriter, r *http.Request) {
	// err1 := Unmarshal(urlArt, &SearchArtists.Artists)
	// err2 := Unmarshal(urlRel, &SearchArtists)

	// if err1 != nil && err2 != nil {
	// 	fmt.Println("ERROR")
	// 	return
	// }

	if r.Method != "GET" {
		w.WriteHeader(405)
		templates.ExecuteTemplate(w, "errors.html", http.StatusMethodNotAllowed)
	}

	out := &SearchOutput{}

	search := strings.Split(r.FormValue("sBar"), " - ")
	params := strings.Split(r.FormValue("Parametrs"), "/")

	if len(search) == 2 {
		params = strings.Split(search[1], "/")
		fmt.Println("Here777", params)
	}

	if params[0] == "Artist" {
		params[0] = "Name"
	} else if params[0] == "Location" {
		params[0] = "DatesLocations"
	}

	searchWord := search[0]

	temp, _ := strconv.Atoi(search[0])

	switch params[0] {
	case "Name":
		fmt.Println("Name")
		for i := 0; i < len(SearchArtists.Artists); i++ {
			if searchWord == SearchArtists.Artists[i].Name || strings.HasPrefix(SearchArtists.Artists[i].Name, searchWord) == true {
				out.Art = append(out.Art, SearchArtists.Artists[i])
			}
		}
	case "Members":
		fmt.Println("Members")
		for i := 0; i < len(SearchArtists.Artists); i++ {
			if FindMembers(searchWord, SearchArtists.Artists[i].Members) {
				out.Art = append(out.Art, SearchArtists.Artists[i])
			}
		}
	case "CreationDate":
		fmt.Println("CreationDate")
		for i := 0; i < len(SearchArtists.Artists); i++ {
			if temp == SearchArtists.Artists[i].CreationDate {
				out.Art = append(out.Art, SearchArtists.Artists[i])
			}
		}
	case "FirstAlbum":
		fmt.Println("FirstAlbum")
		for i := 0; i < len(SearchArtists.Artists); i++ {
			if searchWord == SearchArtists.Artists[i].FirstAlbum || strings.HasPrefix(SearchArtists.Artists[i].FirstAlbum, search[0]) == true {
				fmt.Println("HereFA")
				out.Art = append(out.Art, SearchArtists.Artists[i])

			}
		}
	case "DatesLocations":
		fmt.Println("DatesLocation")
		for i := 0; i < len(SearchArtists.Artists); i++ {
			if FindDatLocations(searchWord, SearchArtists.Relations[i].DatesLocations) {
				out.Art = append(out.Art, SearchArtists.Artists[i])
			}
		}

	default:
		w.WriteHeader(400)
		templates.ExecuteTemplate(w, "errors.html", http.StatusBadRequest)
		return
	}

	if out.Art == nil {
		templates.ExecuteTemplate(w, "errors.html", 0)
		return

	}
	templates.ExecuteTemplate(w, "searchPage.html", out)
}

func FilterPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		templates.ExecuteTemplate(w, "errors.html", http.StatusBadRequest)
	}

	// r.ParseForm()

	// forms := r.Form

	fmt.Println(r.Form)

	fmt.Println(r.FormValue("sub1"))
	fmt.Println(r.FormValue("sub2"))
	fmt.Println(r.FormValue("sub3"))
	fmt.Println(r.FormValue("sub4"))
	fmt.Println(r.FormValue("CDfrom"))
	fmt.Println(r.FormValue("CDto"))
	fmt.Println(r.FormValue("FAfrom"))
	fmt.Println(r.FormValue("FAto"))
	fmt.Println(r.FormValue("NMfrom"))
	fmt.Println(r.FormValue("NMto"))
	fmt.Println(r.FormValue("LCfrom"))
	fmt.Println(r.FormValue("LCto"))

}

// Unmarshal ...
func Unmarshal(s string, a interface{}) error {
	dataArt, err := http.Get(s)
	if err != nil {
		fmt.Println("No  response from request")
	}
	defer dataArt.Body.Close()
	body, err := ioutil.ReadAll(dataArt.Body) // response body is []byte
	error := json.Unmarshal(body, a)
	if error != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return err
}

func FindMembers(s string, sArr []string) bool {
	for i := 0; i < len(sArr); i++ {

		if sArr[i] == s || strings.HasPrefix(sArr[i], s) == true {
			return true
		}
	}
	return false
}

func FindDatLocations(s string, m map[string][]string) bool {
	for city, _ := range m {
		if city == s {
			return true
		}
	}
	return false
}
func isValidId(n int) bool {
	if n < 1 || n > 52 {
		return false
	}
	return true
}

func Fun(s string) {
	data, err := ioutil.ReadFile(s + ".txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
