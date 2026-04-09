package api

import (
	"encoding/json"
	"groupie-tracker/models"
	"net/http"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

func FetchArtists() ([]models.Artist, error) {
	resp, err := http.Get(baseURL + "/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []models.Artist
	return artists, json.NewDecoder(resp.Body).Decode(&artists)
}

func FetchLocations() ([]models.Location, error) {
	resp, err := http.Get(baseURL + "/locations")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Index []models.Location `json:"index"`
	}
	return data.Index, json.NewDecoder(resp.Body).Decode(&data)
}

func FetchDates() ([]models.Date, error) {
	resp, err := http.Get(baseURL + "/dates")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Index []models.Date `json:"index"`
	}
	return data.Index, json.NewDecoder(resp.Body).Decode(&data)
}

func FetchRelations() ([]models.Relation, error) {
	resp, err := http.Get(baseURL + "/relation")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Index []models.Relation `json:"index"`
	}
	return data.Index, json.NewDecoder(resp.Body).Decode(&data)
}
