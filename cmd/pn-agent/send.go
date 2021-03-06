package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	sys "pokenode.com/pn-agent/pkg/system"
)

type MyDialer net.Dialer

func (d *MyDialer) Dial(network, address string) (net.Conn, error) {
	dd := &net.Dialer{}
	return dd.Dial("tcp4", address)
}

func SendNodeStats() {
	stats := sys.GetStats(MODE)
	b, err := json.Marshal(stats)
	if err != nil {
		fmt.Println(err)
		log.Error("Error occurred.")
		return
	}
	req, err := http.NewRequest("POST", API, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
		log.Error("Error occurred.")
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Pokenode-Agent-Nodeid", NODEID)

	// Do request
	d := &MyDialer{}
	client := &http.Client{
		Transport: &http.Transport{
			Dial: d.Dial,
		},
	}
	var resp *http.Response
	var retryCount int = 3
	for retryCount > 0 {
		resp, err = client.Do(req)
		if err != nil {
			fmt.Println(err)
			log.Error("Error occurred.")
			retryCount -= 1
			//fmt.Println("retryCount: ", retryCount)
		} else {
			break
		}
	}
	if resp != nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		if string(body) == "ok" {
			log.Info("Post ok.")
		} else {
			log.Info("Error occurred.")
		}
	} else {
		log.Error("Resp is nil.")
	}
}
