package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

}

func translation() {
	data := make(map[string]interface{})
	data["sourceLanguage"] = "en"
	data["targetLanguage"] = "de"
	data["texts"] = []string{"11e", "33en"}
	marshal, _ := json.Marshal(data)

	for i := 0; i < 1000; i++ {
		resp, err := http.Post("http://ai-translation.shoplineapp.com/api/translation", "application/json", bytes.NewReader(marshal))
		if err != nil {
			fmt.Print(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		respMap := make(map[string]interface{})
		json.Unmarshal(body, &respMap)
		if !respMap["success"].(bool) {
			fmt.Println(string(body))
		} else {
			fmt.Println("success")
		}
	}
}
