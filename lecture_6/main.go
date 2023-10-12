package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var jsonInput = `{
        "input_index": 0,
        "candidate_index": 0,
        "delivery_line_1": "1 N Rosedale St",
        "last_line": "Baltimore MD 21229-3737",
        "delivery_point_barcode": "212293737013",
        "components": {
            "primary_number": "1",
            "street_predirection": "N",
            "street_name": "Rosedale",
            "street_suffix": "St",
            "city_name": "Baltimore",
            "state_abbreviation": "MD",
            "zipcode": "21229",
            "plus4_code": "3737",
            "delivery_point": "01",
            "delivery_point_check_digit": "3"
        },
        "metadata": {
            "record_type": "S",
            "zip_type": "Standard",
            "county_fips": "24510",
            "county_name": "Baltimore City",
            "carrier_route": "C047",
            "congressional_district": "07",
            "rdi": "Residential",
            "elot_sequence": "0059",
            "elot_sort": "A",
            "latitude": 39.28602,
            "longitude": -76.6689,
            "precision": "Zip9",
            "time_zone": "Eastern",
            "utc_offset": -5,
            "dst": true
        },
        "analysis": {
            "dpv_match_code": "Y",
            "dpv_footnotes": "AABB",
            "dpv_cmra": "N",
            "dpv_vacant": "N",
            "active": "Y"
        }
    }
`

var str = ""

func main() {
	generateStructFromJSON(jsonInput)
}

func generateStructFromJSON(jsonInput string) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonInput), &data); err != nil {
		fmt.Println(err)
		return
	}
	str += "type T struct { \n"
	generate(data, 1)
	str += "}"
	fmt.Println(str)
}

func generate(data map[string]interface{}, lvl int) {
	for key, value := range data {
		tab := ""
		for i := 0; i < lvl; i++ {
			tab += "\t"
		}

		switch value.(type) {
		case map[string]interface{}:
			str += fmt.Sprintf("%s%s struct {\n", tab, cases.Title(language.English, cases.NoLower).String(key))
			generate(value.(map[string]interface{}), lvl+1)
			str += fmt.Sprintf("%s}\n", tab)
		default:
			str += fmt.Sprintf("%s%s %T `json:\"%s\"`\n", tab, cases.Title(language.English, cases.NoLower).String(key), value, key)
		}
	}
}
