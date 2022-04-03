package lib

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func get_plan(plan_location string) map[interface{}]interface{} {
	yfile, err := ioutil.ReadFile(plan_location)

	if err != nil {
		log.Fatal(err)
	}

	data := make(map[interface{}]interface{})
	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	return data
}

func validate_plan_contents(plan_data map[interface{}]interface{}) bool {
	for k, v := range plan_data {
		fmt.Printf("%s -> %d\n", k, v)
	}
	return true
}
