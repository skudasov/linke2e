package suite

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func JSONFileToMap(path string) map[string]interface{} {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(f)

	var r map[string]interface{}
	if err := json.Unmarshal(byteValue, &r); err != nil {
		log.Fatal(err)
	}
	return r
}
