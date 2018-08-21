package ln2latlng

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"log"
	"strings"

	"fmt"
	"os"

	"github.com/willxm/ln2latlng/utils"
)

type LocationPoint struct {
	Code     string
	Location string
	Lng      float64
	Lat      float64
}

var LocationPointData []byte
var AreaData map[string]string

func init() {
	log.Println("init location data...")
	var err error
	gopath := os.Getenv("GOPATH")
	dataPath := gopath + "/src/github.com/willxm/ln2latlng/data/"
	if LocationPointData == nil {
		LocationPointData, err = utils.ReadFile(dataPath + "location.json")
		if err != nil {
			panic(err)
		}
	}

	if AreaData == nil {
		AreaData = make(map[string]string)

		f, err := os.Open(dataPath + "area.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		rd := bufio.NewReader(f)
		for {
			line, err := rd.ReadString('\n')

			if err != nil || io.EOF == err {
				break
			}
			line = strings.Trim(line, "\n")
			res := strings.Split(line, "\t")
			if len(res) > 1 {
				AreaData[res[0]] = res[1]
			}
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

func IdCard2Area(idCard string) (string, string, string, error) {
	areaCode2 := idCard[:2]
	areaCode4 := idCard[:4]
	areaCode6 := idCard[:6]

	var area2, area4, area6 string
	var err error

	if v, ok := AreaData[areaCode6]; ok == true {
		area6 = v
	} else {
		err = fmt.Errorf("area6 not found")
	}

	if v, ok := AreaData[areaCode4]; ok == true {
		area4 = v
	} else {
		err = fmt.Errorf("area4 not found")
	}

	if v, ok := AreaData[areaCode2]; ok == true {
		area2 = v
	} else {
		err = fmt.Errorf("area2 not found")
	}
	return area2, area4, area6, err
}
