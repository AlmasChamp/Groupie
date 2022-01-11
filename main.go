package main

import (
	"fmt"
	groupie "groupie/src"
	"log"
	"net/http"
)

func main() {
	// if tmplError != nil {
	// 	fmt.Println("Templates Error")
	// 	return
	// }

	err1 := groupie.Unmarshal(groupie.UrlArt, &groupie.SearchArtists.Artists)
	err2 := groupie.Unmarshal(groupie.UrlRel, &groupie.SearchArtists)
	if err1 != nil || err2 != nil {
		fmt.Println("ERROR")
		return
	}

	groupie.Fun("picture")
	fmt.Println("Server 8080 connected ...")
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js/"))))
	http.HandleFunc("/", groupie.HomePage)
	http.HandleFunc("/artists/", groupie.ArtPage)
	http.HandleFunc("/search/", groupie.SearchPage)
	http.HandleFunc("/filter", groupie.FilterPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func Exampl(w http.Response, r *http.Request)
