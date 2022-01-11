package groupie

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(404)
		Templates.ExecuteTemplate(w, "errors.html", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		w.WriteHeader(405)
		Templates.ExecuteTemplate(w, "errors.html", http.StatusMethodNotAllowed)
		return
	}

	Templates.ExecuteTemplate(w, "mainPage.html", SearchArtists)
}

// WorkPage func
func ArtPage(w http.ResponseWriter, r *http.Request) {

	ind, err := strconv.Atoi(r.RequestURI[9:])

	if err != nil {
		w.WriteHeader(400)
		Templates.ExecuteTemplate(w, "errors.html", http.StatusBadRequest)
		return
	}

	if !isValidId(ind) {
		w.WriteHeader(404)
		Templates.ExecuteTemplate(w, "errors.html", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		w.WriteHeader(405)
		Templates.ExecuteTemplate(w, "errors.html", http.StatusMethodNotAllowed)
		return
	}

	res := &ArtOutput{}
	res.A = SearchArtists.Artists[ind-1]
	res.R = SearchArtists.Relations[ind-1]
	Templates.ExecuteTemplate(w, "artPage.html", res)
}

func SearchPage(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(405)
		Templates.ExecuteTemplate(w, "errors.html", http.StatusMethodNotAllowed)
	}

	out := &SearchOutput{}

	search := strings.Split(r.FormValue("sBar"), " - ")
	params := strings.Split(r.FormValue("Parametrs"), "/")

	if len(search) == 2 {
		params = strings.Split(search[1], "/")
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
		// fmt.Println("Name")
		for i := 0; i < len(SearchArtists.Artists); i++ {
			if searchWord == SearchArtists.Artists[i].Name || strings.HasPrefix(SearchArtists.Artists[i].Name, searchWord) == true {
				out.Art = append(out.Art, SearchArtists.Artists[i])
			}
		}
	case "Members":
		// fmt.Println("Members")
		for i := 0; i < len(SearchArtists.Artists); i++ {
			if FindMembers(searchWord, SearchArtists.Artists[i].Members) {
				out.Art = append(out.Art, SearchArtists.Artists[i])
			}
		}
	case "CreationDate":
		// fmt.Println("CreationDate")
		for i := 0; i < len(SearchArtists.Artists); i++ {
			if temp == SearchArtists.Artists[i].CreationDate {
				out.Art = append(out.Art, SearchArtists.Artists[i])
			}
		}
	case "FirstAlbum":
		// fmt.Println("FirstAlbum")
		for i := 0; i < len(SearchArtists.Artists); i++ {
			if searchWord == SearchArtists.Artists[i].FirstAlbum || strings.HasPrefix(SearchArtists.Artists[i].FirstAlbum, search[0]) == true {
				out.Art = append(out.Art, SearchArtists.Artists[i])

			}
		}
	case "DatesLocations":
		// fmt.Println("DatesLocation")
		for i := 0; i < len(SearchArtists.Artists); i++ {
			if FindDatLocations(searchWord, SearchArtists.Relations[i].DatesLocations) {
				out.Art = append(out.Art, SearchArtists.Artists[i])
			}
		}

	default:
		w.WriteHeader(400)
		Templates.ExecuteTemplate(w, "errors.html", http.StatusBadRequest)
		return
	}

	if out.Art == nil {
		Templates.ExecuteTemplate(w, "errors.html", 0)
		return

	}
	Templates.ExecuteTemplate(w, "searchPage.html", out)
}

func FilterPage(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(405)
		Templates.ExecuteTemplate(w, "errors.html", http.StatusBadRequest)
	}

	// r.ParseForm()
	// fmt.Println(r.Form)

	temp := &Filter{}
	out := &SearchOutput{}

	data := map[string][]string{
		"CD": []string{r.FormValue("CD"), r.FormValue("CDfrom"), r.FormValue("CDto")},
		"FA": []string{r.FormValue("FA"), r.FormValue("FAfrom"), r.FormValue("FAto")},
		"NM": []string{r.FormValue("NM"), r.FormValue("NMfrom"), r.FormValue("NMto")},
		"LC": []string{r.FormValue("LC"), r.FormValue("LCfrom")},
	}

	if CheckValue(data, temp) != nil {
		fmt.Println("err")
		w.WriteHeader(400)
		Templates.ExecuteTemplate(w, "errors.html", http.StatusBadRequest)
		return
	}
	// fmt.Println(temp.LocationsOfConcerts, "here232323")
	for i := 0; i < len(SearchArtists.Artists); i++ {
		// fmt.Println("here7788")
		if len(temp.CreationDate) != 0 && !(temp.CreationDate[0] <= SearchArtists.Artists[i].CreationDate && temp.CreationDate[1] >= SearchArtists.Artists[i].CreationDate) {
			// out.Art = append(out.Art, SearchArtists.Artists[i])
			continue
		}
		fa1 := strings.Split(SearchArtists.Artists[i].FirstAlbum, "-")
		// fmt.Println(fa1[2])
		fa, _ := strconv.Atoi(fa1[2])
		if len(temp.FirstAlbum) != 0 && !(temp.FirstAlbum[0] <= fa && temp.FirstAlbum[1] >= fa) {
			continue
		}
		if len(temp.NumberOfMembers) != 0 && !(temp.NumberOfMembers[0] <= len(SearchArtists.Artists[i].Members) && temp.NumberOfMembers[1] >= len(SearchArtists.Artists[i].Members)) {
			continue
		}
		if len(temp.LocationsOfConcerts) != 0 && !(FindDatLocations(temp.LocationsOfConcerts, SearchArtists.Relations[i].DatesLocations)) {
			continue
		}

		out.Art = append(out.Art, SearchArtists.Artists[i])
	}

	// fmt.Println(temp.CreationDate[0], temp.CreationDate[1], "hereAlmas", out.Art)

	Templates.ExecuteTemplate(w, "filterPage.html", out)
}
