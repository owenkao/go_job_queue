package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type Job struct {
	Action  string                 `json:"action"`
	Content map[string]interface{} `json:"content"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	noCareJSON()
	json2map()
	map2json()
	map2struct()

	person := Person{"test", 28}
	res := struct2map(person)
	fmt.Printf("%v\n", res)
}

func noCareJSON() {
	var job Job
	data := []byte(`{"action":"ota","content":{"a":"b", "c":"d", "e":{"f":"g"}}}`)
	err := json.Unmarshal(data, &job)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v %v\n", job, job.Content["e"])

	jsonStr, _ := json.Marshal(job)
	fmt.Printf("%s\n", jsonStr)

	jsondata, err := json.Marshal(job)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsondata))
}

func json2map() {
	jsonStr := `
    {
        "name":"liangyongxing",
        "age":12
    }
    `
	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &mapResult); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", mapResult)
	fmt.Printf("%s\n", mapResult["name"])
}

func map2json() {
	fmt.Println("map2json")
	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "liang637210"
	mapInstance["Age"] = 28
	mapInstance["Address"] = "ttttt"
	fmt.Printf("%+v\n", mapInstance)
	jsonStr, err := json.Marshal(mapInstance)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", jsonStr)
}

func map2struct() {
	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "liang637210"
	mapInstance["Age"] = 28

	var person Person
	if err := mapstructure.Decode(mapInstance, &person); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", person)
}

func struct2map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
