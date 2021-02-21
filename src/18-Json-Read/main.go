package main

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Name        string `json:"name"`
	Capacity    int    `json:"capacity"`
	Time        string `json:"time"`
	Information Info   `json:"info"`
}

type Info struct {
	Description string `json: "description"`
}

func main() {
	fmt.Println("Json Reader Tutorial")

	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2020-11-15T08:12:00Z", "info": {"description": "a sensor reading"}}`

	// var reading SensorReading
	var reading map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &reading)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", reading)
}
