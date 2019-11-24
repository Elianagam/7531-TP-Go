package main

import (
	"fmt"
	"github.com/Nicobugliot/7531-TP-Go/twitter/domain"
	"github.com/Nicobugliot/7531-TP-Go/twitter/repository"
	"strings"
	"time"
)

func main()  {

	resultChannel := make(chan string)

	users := []string{"alferdez", "mauriciomacri"}
	query := "kir"


	go search(resultChannel, users, query)

	go func() {
		for {
			print(<- resultChannel + "\n\n")
		}
	}()

	// Para que no termine la ejecución
	var input string
	fmt.Scanln(&input)
}

func search(resultChannel chan string, users []string, query string) {

	tweetsToProcess := make(chan *domain.Tweet)

	for _,user := range users {
		go getTweetsFromUser(tweetsToProcess, user)
		go processTweets(resultChannel, tweetsToProcess, query)
	}

}

func getTweetsFromUser(channel chan *domain.Tweet, user string) {

	var repo repository.TwitterRepository = repository.NewFileTwitterRepository()

	tweets,err := repo.GetTweetsFromUser(user)
	if err != nil {
		panic("Can't retrieve tweets for user " + user)
	}

	for _,tweet := range tweets {
		time.Sleep(30 * time.Millisecond)

		channel <- tweet
	}
}

func processTweets(resultChannel chan string, tweetsToProcess chan *domain.Tweet, query string)  {
	for {
		tweet := <- tweetsToProcess
		if strings.Contains(tweet.Text, query) {
			resultChannel <- tweet.ToString()
		}
	}
}