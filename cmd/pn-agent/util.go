package main

import (
	"encoding/json"
	"fmt"
)

func PPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}
