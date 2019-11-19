package main

import (
	"fmt"
	"github.com/Nicobugliot/7531-TP-Go/twitter/domain"
	"github.com/Nicobugliot/7531-TP-Go/twitter/search"
	"regexp"
	"strings"
)

func main()  {
	users := []string {"alferdez", "mauriciomacri"}

	resultChannel := make(chan string)

	go search.Search(resultChannel, users, hasMoreThanNLikes(70000))

	var count int

	for message := range resultChannel {
		print(message + "\n\n")
		count++
	}

	print(fmt.Sprintf("\n\nTermino --------------- %d \n\n", count))
}

func containsQuery(query string) func(*domain.Tweet) bool {
	return func(tweet *domain.Tweet) bool {
		return strings.Contains(strings.ToLower(tweet.Text), strings.ToLower(query))
	}
}

func hasMoreThanNLikes(likes int) func(*domain.Tweet) bool {
	return func(tweet *domain.Tweet) bool {
		return tweet.Likes >= likes
	}
}

func containsAnEmoji(tweet *domain.Tweet) bool {
	var emojiRx = regexp.MustCompile("[\u1000-\uFFFF]+") // TODO encontrar una regex que funcione!
	return emojiRx.MatchString(tweet.Text)
}