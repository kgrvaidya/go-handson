package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var urls = []string{
	"http://www.mocky.io/v2/5ecfd6473200009dc1e3d64e",
	// "http://www.mocky.io/v2/5ecfd5dc3200006200e3d64b",
	// "http://www.mocky.io/v2/5ecfd630320000f1aee3d64d",
}

type Charecter struct {
	Name     string  `json:"name"`
	MaxPower float64 `json:"max_power"`
	// TimesRequested int32  `json:"times_requested"`
}

type CharectersList struct {
	Character []Charecter
}

// var jsonData = `{
// 	"name": "Mutants",
// 	"character": [
//     {
//       "name": "Iron man",
//       "max_power": 60
//     },
//     {
//       "name": "Captain America",
//       "max_power": 68
//     },
//     {
//       "name": "Spider man",
//       "max_power": 58
//     },
//     {
//       "name": "Black Panther",
//       "max_power": 68
//     },
//     {
//       "name": "Vision",
//       "max_power": 50
//     },
//     {
//       "name": "Hawk eye",
//       "max_power": 30
//     }
//   ]
// }`
func getCaracters(body map[string]interface{}) *Charecter {
	var s = new(Charecter)
	s.MaxPower, s.Name = body["max_power"].(float64), body["name"].(string)
	return s
}

func main() {

	res, err := http.Get(urls[0])
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var response map[string]interface{}
	json.Unmarshal(body, &response)

	// var charec []interface{}
	charecterList := response["character"].([]interface{})

	for _, charecter := range charecterList {
		res := charecter.(map[string]interface{})
		s := getCaracters(res)

		fmt.Printf("%+v \n", s)
	}

	// if err != nil {
	// 	panic(err.Error())
	// }

	// fmt.Printf("%v ", s)

}
