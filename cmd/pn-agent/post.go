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
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", API, bytes.NewBuffer(b))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
