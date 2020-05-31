package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"time"
)

var urls = []string{
	"http://www.mocky.io/v2/5ecfd6473200009dc1e3d64e",
	"http://www.mocky.io/v2/5ecfd5dc3200006200e3d64b",
	"http://www.mocky.io/v2/5ecfd630320000f1aee3d64d",
}

type Charecter struct {
	Name     string  `json:"name"`
	MaxPower float64 `json:"max_power"`
	// TimesRequested int32  `json:"times_requested"`
}

var charList []Charecter

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

// func sort

type ByPowerRange []Charecter

func (a ByPowerRange) Len() int           { return len(a) }
func (a ByPowerRange) Less(i, j int) bool { return a[i].MaxPower < a[j].MaxPower }
func (a ByPowerRange) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func HandleAPICall(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s \n", body)

	var response map[string]interface{}
	json.Unmarshal(body, &response)

	fmt.Printf("After Marshal --> %v \n", response)

	charecterList := response["character"].([]interface{})

	for _, charecter := range charecterList {
		res := charecter.(map[string]interface{})
		s := getCaracters(res)
		if len(charList) >= 10 {
			//Define condition to remove the least used charecter from here
			// Satisfies second criteria. Removes charecter with least power.
			// Yet to implement first case.
			sort.Sort(ByPowerRange(charList))

			charList = append(charList[1:], *s)

		} else {
			charList = append(charList[:], *s)
		}

	}
}

func getCharecterDetail(charName string) float64 {
	for i := 0; i < len(charList); i++ {
		if charList[i].Name == charName {
			fmt.Printf(" Match Found %+v", charList[i])
			return charList[i].MaxPower
		}
	}
	return -1
}

func main() {

	HandleAPICall(urls[0])
	HandleAPICall(urls[1])
	time.Sleep(5 * time.Second)
	getCharecterDetail("Apocalype")
	// HandleAPICall(urls[2])

	/*
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

		charecterList := response["character"].([]interface{})

		for _, charecter := range charecterList {
			res := charecter.(map[string]interface{})
			s := getCaracters(res)
			if cap(charList) >= 5 {
				fmt.Printf(" Max Capacity")
			} else {
				charList = append(charList[:], *s)
			}
			// fmt.Printf("%+v \n", *s)

		}
		fmt.Printf("%v", charList)

		// if err != nil {
		// 	panic(err.Error())
		// }

		// fmt.Printf("%v ", s)

	*/
}
