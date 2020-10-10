package main

import "os"

var (
	API string
)

func init_env() {
	API = os.Getenv("API")
	if API == "" {
		API = "https://api.pokenode.com/stats"
	}
}
