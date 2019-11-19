package search

import (
	"github.com/Nicobugliot/7531-TP-Go/twitter/domain"
	"github.com/Nicobugliot/7531-TP-Go/twitter/repository"
)

func Search(channelToReceive chan string, users []string, apply func(*domain.Tweet) bool) {

	// Create channel to process tweets
	channelToAggFunction := make(chan *domain.Tweet) // TODO add buffer

	// Create channel used to know when goroutines have finished
	quitChannel := make(chan struct{})

	// Start process goroutine
	go processTweets(channelToReceive, channelToAggFunction, apply, quitChannel)

	// Initialize a goroutine for each user
	for _,user := range users {
		go postTweetsFromUser(channelToAggFunction, user, quitChannel)
	}

	// Wait until all  goroutines finish
	// This channel will receive a message for each finished
	// user goroutine and an extra message for process goroutine
	for range users  {
		<- quitChannel // wait users goroutines
	}

	// We tell to process goroutine that user's goroutine have finish
	quitChannel <- struct{}{}

	// Wait until process goroutine finish
	<- quitChannel

	close(channelToReceive)
}


func postTweetsFromUser(channel chan *domain.Tweet, user string, quitChannel chan struct{}) {

	var repo repository.TwitterRepository = repository.NewFileTwitterRepository()

	tweets,err := repo.GetTweetsFromUser(user)
	if err != nil {
		panic("Can't retrieve tweets for user " + user)
	}

	for _,tweet := range tweets {
		//time.Sleep(30 * time.Millisecond) Para simular response time de la API
		channel <- tweet
	}

	// notify this function has finish
	quitChannel <- struct{}{}
}

func processTweets(channelToReceive chan string, channelToAggFunction chan *domain.Tweet, apply func(*domain.Tweet) bool, quitChannel chan struct{})  {
	allGoRoutinesHaveFinish := false

	for {
		select {

		// Will entry when there is a tweet available for process
		case tweet := <- channelToAggFunction:
			if apply(tweet) {
				channelToReceive <- tweet.ToString()
			}

		// Will entry when all previous routines have finish
		case <- quitChannel:
			allGoRoutinesHaveFinish = true

		// Will entry if both cases are blocked
		default:
			if allGoRoutinesHaveFinish {
				quitChannel <- struct{}{}
			}
		}
	}
}