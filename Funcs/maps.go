package Groupie_tracker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetCoordinates(address string) (float64, float64, error) {
	apiKey := "cVmzhl4sjxU6-5wMt_YwcYpbyOdw3cyfSfEA3s2-E_E"
	baseURL := "https://geocode.search.hereapi.com/v1/geocode"
	reqURL := fmt.Sprintf("%s?q=%s&apiKey=%s", baseURL, url.QueryEscape(address), apiKey)

	resp, err := http.Get(reqURL)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	var geoResp GeoResponse
	err = json.Unmarshal(body, &geoResp)
	if err != nil {
		return 0, 0, err
	}
	if len(geoResp.Items) > 0 {
		position := geoResp.Items[0].Position
		return position.Lat, position.Lng, nil
	}
	return 0, 0, fmt.Errorf("no coordinates found")
}
