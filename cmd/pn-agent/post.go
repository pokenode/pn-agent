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
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest("POST", API, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Pokenode-Agent-Nodeid", NODEID)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
