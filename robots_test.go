package main

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_robots_parse(t *testing.T) {
	//get robots content
	r, _ := http.Get("http://www.163.com/robots.txt")

	//parse robots
	rules, sitemaps, err := Parse(r.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(rules)
	fmt.Println(sitemaps)

	//check url dir
	robotsAllowed, matchData, delaySeconds := CheckPath("http://www.163.com/bbs", "Baiduspider", rules)

	if !robotsAllowed {
		fmt.Println(matchData)
		fmt.Println(delaySeconds)
		fmt.Println("It's not allowed.")
	}
}
