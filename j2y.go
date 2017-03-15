package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var data interface{}

func main() {
	if len(os.Args) == 2 {
		bytes, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(bytes, &data)
		if err != nil {
			log.Fatal(err)
		}
		yaml, err := yaml.Marshal(&data)
		if err != nil {
			log.Fatal(err)

		}
		fmt.Printf("---\n%s", string(yaml))
	}
}
