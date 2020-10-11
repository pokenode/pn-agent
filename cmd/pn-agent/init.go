package main

import "os"

var (
	MODE   string
	API    string
	NODEID string
)

func init_env() {
	MODE = os.Getenv("MODE")
	if MODE == "" {
		MODE = "DEV"
	}
	API = os.Getenv("API")
	if API == "" {
		API = "https://api.pokenode.com/stats"
	}
	NODEID = os.Getenv("NODEID")
	if NODEID == "" {
		NODEID = "test_nodeid"
	}
}
