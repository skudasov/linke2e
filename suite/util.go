package suite

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/tidwall/gjson"
)

func GJSONFromFile(path string) gjson.Result {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(f)
	return gjson.ParseBytes(byteValue)
}

func GJSONToMap(g gjson.Result) map[string]interface{} {
	var r map[string]interface{}
	if err := json.Unmarshal([]byte(g.Raw), &r); err != nil {
		log.Fatal(err)
	}
	return r
}

func GJSONFromMap(m map[string]interface{}) gjson.Result {
	d, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return gjson.ParseBytes(d)
}
