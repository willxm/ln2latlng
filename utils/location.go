package utils

import (
	"encoding/json"
	"errors"
)

type LocationPoint struct {
	Code     string
	Location string
	Lng      float64
	Lat      float64
}

var LocationPointData []byte

func init() {
	var err error
	if LocationPointData == nil {
		LocationPointData, err = ReadFile("../data/location.json")
		if err != nil {
			panic(err)
		}
	}

}

func IdCrad2LatLng(idCard string) (string, float64, float64, error) {
	var lps []LocationPoint
	json.Unmarshal(LocationPointData, &lps)
	for _, v := range lps {
		if idCard[:2] == v.Code {
			return v.Location, v.Lat, v.Lng, nil
		}
	}
	return "", 0.0, 0.0, errors.New("No found location")
}
