package main

import (
	"time"

	_ "github.com/prometheus/client_golang/prometheus"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
}

/**
 * Pick 2 peers from first N peers, the picked 2 peers are different are guranteed
 *     1. if n == 0, pick 2 from all peers and set n = peers.size - 2
 *     2. if n == 1, pick 1 from all other peers
 * @author xiaozongyang
 */
func pick(peers []string, n int) (string, string) {
}

/**
 * Send a request to peer `from`, which sends a probe request to peer `to`. Return the status
 * whether the request sent from `from` to `to` succeeed, and the duration
 * @author xiaozongyang
 */
func probe(from string, to string) (bool, time.Duration) {
}

/**
 * Expose the success status and duration to prometheus
 * @author xiaozongyang
 */
func collect(from string, to string, success bool, duration time.Duration) {
}
