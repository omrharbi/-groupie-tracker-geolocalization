package Groupie_tracker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// func to Get All data From json
func GetArtistsDataStruct() ([]JsonData, error) {
	var artistData []JsonData

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, fmt.Errorf("error fetching data from artist data: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching data from artist data: %d", response.StatusCode)
	}

	err = json.NewDecoder(response.Body).Decode(&artistData)
	if err != nil {
		return nil, fmt.Errorf("khata2 f t7wil JSON: %v", err)
	}
	newartist := []JsonData{}
	for _, val := range artistData {
		if val.Id == 21 || val.Id == 22 || val.Id == 43 {
			continue
		} else {
			newartist = append(newartist, val)
		}
	}
	return newartist, nil
}

var location Location

// / func to fetching data from any struct and return Struct Artist with Id user
func FetchDataRelationFromId(id, cityes string) (Artist, error) {
	fmt.Println(cityes)
	url := "https://groupietrackers.herokuapp.com/api"
	urlartist := url + "/artists/" + id
	var artist Artist
	err := GetanyStruct(urlartist, &artist)
	if err != nil {
		return Artist{}, fmt.Errorf("error fetching data from artist data: %w", err)
	}

	if artist.Id == 0 {
		return artist, nil
	}

	var date Date
	urldate := url + "/dates/" + id
	err = GetanyStruct(urldate, &date)
	if err != nil {
		return Artist{}, fmt.Errorf("error fetching data from artist data: %w", err)
	}

	urlLocation := url + "/locations/" + id
	err = GetanyStruct(urlLocation, &location)
	if err != nil {
		return Artist{}, fmt.Errorf("error fetching data from locations data: %w", err)
	}
	var relation Relation
	urlrelation := url + "/relation/" + id
	err = GetanyStruct(urlrelation, &relation)
	if err != nil {
		return Artist{}, fmt.Errorf("error fetching data from locations data: %w", err)
	}

	artist.Location = location.Location
	artist.Date = date.Date
	artist.DatesLocations = formatLocations(relation.DatesLocations)
	return artist, nil
}

// func to UnmarshalJSON from any struct with send url and any
// return error for has any error
func GetanyStruct(url string, result interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data from URL: %w", err)
	}
	defer response.Body.Close()
	// Decode the JSON response into the provided struct
	if err := json.NewDecoder(response.Body).Decode(result); err != nil {
		return fmt.Errorf("error decoding JSON data: %w", err)
	}
	return nil
}

func getCities() []string {
	cities := []string{}
	for _, item := range location.Location {
		city := strings.NewReplacer("-", " ", "_", " ").Replace(item)
		cities = append(cities, city)

	}
	return cities
}

func SendData(lat, lon float64, city string) map[string]interface{} {
	cities := getCities()
	return map[string]interface{}{
		"Token":        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJkODRkOWJmMy1lZWIwLTRjNTktOTFlNC1mMzMxYWJjYTAxYzAiLCJpZCI6MjM2NzE4LCJpYXQiOjE3MjQ0OTU5Mjd9.xUdL_xpp9kiw0sS_VQ5IOFdlFfN6ORhU4PPsyGVP_kc",
		"Cities":       cities,
		"Lat":          lat,
		"Lon":          lon,
		"SelectedCity": city,
	}
}

// func To Format String To remove '_' or '-' and Capitaliz text
func formatLocations(locations map[string][]string) map[string][]string {
	formatted := make(map[string][]string, len(locations))
	for location, dates := range locations {
		formattedLoc := strings.Title(strings.NewReplacer("-", " ", "_", " ").Replace(location))
		formatted[formattedLoc] = dates
	}
	return formatted
}
