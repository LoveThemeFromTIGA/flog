package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"strings"

	"github.com/brianvoe/gofakeit"
)

// RandResourceURI generates a random resource URI
func RandResourceURI() string {
	var uri string
	num := gofakeit.Number(1, 4)
	for i := 0; i < num; i++ {
		uri += "/" + url.QueryEscape(gofakeit.BS())
	}
	uri = strings.ToLower(uri)
	return uri
}

// RandAuthUserID generates a random auth user id
func RandAuthUserID() string {
	candidates := []string{"-", strings.ToLower(gofakeit.Username())}
	return candidates[rand.Intn(2)]
}

// RandHTTPVersion returns a random http version
func RandHTTPVersion() string {
	versions := []string{"HTTP/1.0", "HTTP/1.1", "HTTP/2.0"}
	return versions[rand.Intn(3)]
}

func IPv4LanAddress() string {
	num := func() int { return 2 + rand.Intn(254) }
	return fmt.Sprintf("192.168.%d.%d", num(), num())
}

func RandIntRange(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn((max+1)-min) + min
}

func RandMillisecond(min, max int) float64 {
	num := RandIntRange(min, max)
	return float64(num) / float64(1000)
}
