package main

import (
	"InstagramComparer/internal/handler"
	"os"
)

func main() {
	followers, _ := os.Open("data/followers.json")
	following, _ := os.Open("data/following.json")

	handler.Compare(followers, following)
}
