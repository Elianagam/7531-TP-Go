package main

import (
	"github.com/Nicobugliot/7531-TP-Go/twitter/domain"
	"github.com/Nicobugliot/7531-TP-Go/twitter/repository"
	"strings"
	"time"
)

func main()  {

	users := []string{"alferdez", "mauriciomacri"}
	query := "kir"

	resultChannel := make(chan string)

	go Search(resultChannel, users, query)

	go func() {
		for tweet := range resultChannel {
			print(tweet + "\n\n")
		}
	}()

	time.Sleep(5 * time.Second)
}

func Search(resultChannel chan string, users []string, query string) {

	processChannel := make(chan *domain.Tweet)

	for _,user := range users {
		go getTweetsFromUser(processChannel, user)
	}

	go processTweets(resultChannel, processChannel, query)

}



func getTweetsFromUser(processChannel chan *domain.Tweet, user string) {

	var repo repository.TwitterRepository = repository.NewFileTwitterRepository()

	tweets, err := repo.GetTweetsFromUser(user)
	if err != nil {
		panic("Error retrieving tweets")
	}

	for _,tweet := range tweets {
		processChannel <- tweet
	}
}

func processTweets(resultChannel chan string, tweetsToProcess chan *domain.Tweet, query string)  {

	for tweet := range tweetsToProcess {
		if strings.Contains(tweet.Text, query) {
			resultChannel <- tweet.ToString()
		}
	}

}