package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	sys "pokenode.com/pn-agent/pkg/system"
)

func SendNodeStats() {
	stats := sys.GetStats()
	b, err := json.Marshal(stats)
	req, err := http.NewRequest("POST", API, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
