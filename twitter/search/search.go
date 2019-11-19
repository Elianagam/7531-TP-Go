package search

import (
	"github.com/Nicobugliot/7531-TP-Go/twitter/domain"
	"github.com/Nicobugliot/7531-TP-Go/twitter/repository"
)

func Search(resultsChannel chan string, users []string, apply func(*domain.Tweet) bool) {

	// Create channel to process tweets
	processChannel := make(chan *domain.Tweet, 10)

	// Create channels used to know when goroutines have finished
	quitUsersChannel := make(chan struct{})
	quitProcessChannel := make(chan struct{})

	// Start process goroutine
	go processTweets(resultsChannel, processChannel, apply, quitProcessChannel)

	// Start a goroutine to get tweets for each user
	for _,user := range users {
		go getTweetsFromUser(processChannel, user, quitUsersChannel)
	}

	// Wait until all user's goroutines finish
	// This channel will receive a message for each finished user goroutine
	for range users {
		<- quitUsersChannel
	}

	// We tell to process goroutine that user's goroutine have finish
	close(quitProcessChannel)

	// Wait until process goroutine finish (close his channel)
	for range quitProcessChannel {}

	close(resultsChannel)
}


func getTweetsFromUser(channel chan *domain.Tweet, user string, quitChannel chan struct{}) {
	defer func() {
		quitChannel <- struct{}{} // notify this function has finish
	}()

	var repo repository.TwitterRepository = repository.NewFileTwitterRepository()

	tweets,err := repo.GetTweetsFromUser(user)
	if err != nil {
		panic("Can't retrieve tweets for user " + user)
	}

	for _,tweet := range tweets {
		//time.Sleep(30 * time.Millisecond) // Para simular response time de la API
		channel <- tweet
	}
}

func processTweets(results chan string, tweetsToProcess chan *domain.Tweet, apply func(*domain.Tweet) bool, quitChannel chan struct{})  {

	defer close(quitChannel)

	for tweet := range tweetsToProcess {
		if apply(tweet) {
			results <- tweet.ToString()
		}
	}
}