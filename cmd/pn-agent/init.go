package main

import "os"

var (
	API    string
	NODEID string
)

func init_env() {
	API = os.Getenv("API")
	if API == "" {
		API = "https://api.pokenode.com/stats"
	}
	NODEID = os.Getenv("NODEID")
	if NODEID == "" {
		NODEID = "test_nodeid"
	}
}
