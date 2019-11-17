package main

import (
	"fmt"
	"github.com/Nicobugliot/7531-TP-Go/twitter/domain"
	"github.com/Nicobugliot/7531-TP-Go/twitter/repository"
	"strings"
)

func main()  {
	channelToReceive := make(chan string)
	
	users := []string {"userA", "userB"}

	go search(channelToReceive, users, containsQuery("paste"))

	go func() {
		for {
			print(<- channelToReceive + "\n")
		}
	}()

	var input string
	fmt.Scanln(&input)
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
		channel <- tweet
	}
}

func aggFunction(channelToReceive chan string, channelToAggFunction chan *domain.Tweet, apply func(*domain.Tweet) bool)  {
	for {
		//time.Sleep(1 * time.Second)
		tweet := <- channelToAggFunction

		if apply(tweet) {
			channelToReceive <- tweet.Text
		}

	}
}

func containsQuery(query string) func(*domain.Tweet) bool {
	return func(tweet *domain.Tweet) bool {
		return strings.Contains(strings.ToLower(tweet.Text), strings.ToLower(query))
	}
}
