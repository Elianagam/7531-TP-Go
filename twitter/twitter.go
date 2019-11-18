package main

import (
	"github.com/Nicobugliot/7531-TP-Go/twitter/domain"
	"github.com/Nicobugliot/7531-TP-Go/twitter/repository"
	"regexp"
	"strings"
	"time"
)

func main()  {
	channelToReceive := make(chan string)
	
	users := []string {"alferdez", "mauriciomacri"}

	go search(channelToReceive, users, containsQuery("muert"))

	go func() {
		for {
			print(<- channelToReceive + "\n\n")
		}
	}()

	time.Sleep(5 * time.Minute)
}

func search(channelToReceive chan string, users []string, apply func(*domain.Tweet) bool) {

	channelToAggFunction := make(chan *domain.Tweet)

	for _,user := range users {

		go postTweetsFromUser(channelToAggFunction, user)
		go aggFunction(channelToReceive, channelToAggFunction, apply)
	}
	
}


func postTweetsFromUser(channel chan *domain.Tweet, user string) {

	var repo repository.TwitterRepository = repository.NewFileTwitterRepository()

	tweets,err := repo.GetTweetsFromUser(user)
	if err != nil {
		panic("Can't retrieve tweets for user " + user)
	}

	for _,tweet := range tweets {
		//time.Sleep(30 * time.Millisecond)
		channel <- tweet
	}
}

func aggFunction(channelToReceive chan string, channelToAggFunction chan *domain.Tweet, apply func(*domain.Tweet) bool)  {
	for {
		tweet := <- channelToAggFunction

		if apply(tweet) {
			channelToReceive <- tweet.ToString()
		}

	}
}

func containsQuery(query string) func(*domain.Tweet) bool {
	return func(tweet *domain.Tweet) bool {
		return strings.Contains(strings.ToLower(tweet.Text), strings.ToLower(query))
	}
}

func containsAnEmoji(tweet *domain.Tweet) bool {
	var emojiRx = regexp.MustCompile("[\u1000-\uFFFF]+") // TODO encontrar una regex que funcione!
	return emojiRx.MatchString(tweet.Text)
}