package system

import (
	"fmt"
	"net/http"
	"net/http/httptrace"
	"time"
)

var MAX_LAT int64 = 3000

type LatStats struct {
	NA int64 `json:"na"`
	EU int64 `json:"eu"`
	AS int64 `json:"as"`
	CN int64 `json:"cn"`
}

func Lat() LatStats {
	stats := LatStats{
		NA: DoHTTPing("http://ping-na.pokenode.com"),
		EU: DoHTTPing("http://ping-eu.pokenode.com"),
		AS: DoHTTPing("http://ping-as.pokenode.com"),
		CN: DoHTTPing("http://ping-cn.pokenode.com"),
	}
	//PPrint(stats)
	return stats
}

func DoHTTPing(address string) int64 {
	var tList []int64
	// Do httping for 3 times
	for i := 0; i < 3; i++ {
		t := HTTPing(address)
		tList = append(tList, t)
	}
	//fmt.Println(tList)
	return GetMin(tList)
}

func GetMin(tList []int64) int64 {
	lat := tList[0]
	for _, t := range tList {
		if t < lat {
			lat = t
		}
	}
	return lat
}

func HTTPing(address string) int64 {
	var t0, t1, t2 int64

	req, _ := http.NewRequest("GET", address, nil)
	trace := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			t0 = time.Now().UnixNano()
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			t1 = time.Now().UnixNano()
			if info.Err != nil {
				fmt.Println(info.Err)
			}
		},
		ConnectStart: func(net, addr string) {
		},
		ConnectDone: func(net, addr string, err error) {
			if err != nil {
				fmt.Println(err)
			}
			t2 = time.Now().UnixNano()
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	c := &http.Client{
		Timeout: 3 * time.Second,
	}
	_, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
		return MAX_LAT
	}

	if t0 == 0 {
		t0 = t2
		t1 = t2
	}

	return (t2 - t1) / 1e6
}
