package groupie

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func Unmarshal(s string, a interface{}) error {
	dataArt, err1 := http.Get(s)
	if err1 != nil {
		fmt.Println("No  response from request")
	}
	defer dataArt.Body.Close()
	body, err := ioutil.ReadAll(dataArt.Body) // response body is []byte
	err2 := json.Unmarshal(body, a)
	if err2 != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return err
}

func CheckValue(m map[string][]string, strc *Filter) error {
	var err error
	for key, value := range m {
		// fmt.Println("here0")
		if value[0] != "on" {
			continue
		}
		fmt.Println("here1")
		if key == "LC" {
			// fmt.Println(len(value), "here")
			if FillingStructureCity(key, value[1], strc) != nil {
				return errors.New("error here")
			}
			continue
		}
		if key == "NM" {
			if value[1] == "" {
				value[1] = "0"
			}
			if value[2] == "" {
				value[2] = "10"
			}
		}
		// fmt.Println(key, value, "here2")
		if value[1] == "" {
			value[1] = "1000"
		}
		if value[2] == "" {
			value[2] = "2022"
		}

		from, err1 := strconv.Atoi(value[1])
		fmt.Println(from, "here1")
		if err1 != nil || from < 0 {
			return errors.New("error -")
		}
		to, err2 := strconv.Atoi(value[2])
		fmt.Println(to, "here1")
		if err2 != nil || to < 0 {
			return errors.New("error -s")
		}

		FillingStructure(key, from, to, strc)
	}
	return err
}

func FillingStructure(key string, from int, to int, strc *Filter) error {
	// out := &Filter{}
	if key == "CD" {
		strc.CreationDate = append(strc.CreationDate, from)
		strc.CreationDate = append(strc.CreationDate, to)
	} else if key == "FA" {
		strc.FirstAlbum = append(strc.FirstAlbum, from)
		strc.FirstAlbum = append(strc.FirstAlbum, to)
	} else if key == "NM" {
		strc.NumberOfMembers = append(strc.NumberOfMembers, from)
		strc.NumberOfMembers = append(strc.NumberOfMembers, to)
	}
	// w.WriteHeader(400)
	fmt.Println(strc.CreationDate, "here2")
	// out.FirstAlbum, out.NumberOfMembers,
	return nil
}

func FillingStructureCity(key string, city string, strc *Filter) error {

	citySep := strings.Split(city, ", ")

	for i := 0; i < len(citySep[0]); i++ {
		if citySep[0][i] >= 65 && citySep[0][i] <= 90 || citySep[0][i] >= 97 && citySep[0][i] <= 122 || citySep[0][i] == 45 || citySep[0][i] == 95 || citySep[0][i] == 58 {
			continue
		}
		return errors.New("error here")
	}

	strc.LocationsOfConcerts = strings.ToLower(citySep[0])

	return nil
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
		if city == s || strings.HasPrefix(city, s) == true {
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
