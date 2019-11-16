package main

import (
	"strconv"
	"strings"
	"time"
)

func main()  {
	channelToReceive := make(chan string)
	
	users := []string {"userA", "userB"}
	query := "q"


	go search(channelToReceive, users, query)

	go func() {
		for {
			print(<- channelToReceive)
		}
	}()

	time.Sleep(5 * time.Second)
}

func search(channelToReceive chan string, users []string, query string) {

	channelToAggFunction := make(chan *Tweet)

	for i := 0; i < len(users); i++ {
		user := users[i]

		go postTweetsFromUser(channelToAggFunction, user)
		go aggFunction(channelToReceive, channelToAggFunction)
	}
	
}


func postTweetsFromUser(channel chan *Tweet, user string) {
	for i := 0 ; i < 5 ; i++ {
		tweet := &Tweet{
			user: user,
			text: user + strconv.Itoa(i) + "; ",
		}
		channel <- tweet
	}
}

func aggFunction(channelToReceive chan string, channelToAggFunction chan *Tweet)  {
	for {
		tweet := <- channelToAggFunction
		channelToReceive <- strings.ToUpper(tweet.text)
	}
}

type Tweet struct {
	user string
	text string
}