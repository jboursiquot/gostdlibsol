package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	d := net.Dialer{
		Timeout: time.Millisecond * 5,
	}
	c := http.Client{
		Transport: &http.Transport{
			Dial: d.Dial,
		},
	}
	t := time.Now()
	_, err := c.Get("http://localhost:8080")
	if err != nil {
		log.Println(err)
	}
	log.Printf("%v", time.Now().Sub(t))
}
