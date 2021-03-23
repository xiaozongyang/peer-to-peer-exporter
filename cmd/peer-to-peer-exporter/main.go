package main

import (
	"time"

	_ "github.com/prometheus/client_golang/prometheus"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
}


func pick(peers []string, n int) (string, string) {
}

func probe(from string, to string) (bool, time.Duration){
}

func collect(from string, to string, success bool, duration time.Duration) {
}
