package utils

import (
	"github.com/Nicobugliot/7531-TP-Go/twitter/domain"
	"regexp"
	"strings"
)

func ContainsQuery(query string) func(*domain.Tweet) bool {
	return func(tweet *domain.Tweet) bool {
		return strings.Contains(strings.ToLower(tweet.Text), strings.ToLower(query))
	}
}

func HasMoreThanNLikes(likes int) func(*domain.Tweet) bool {
	return func(tweet *domain.Tweet) bool {
		return tweet.Likes >= likes
	}
}

func ContainsAnEmoji(tweet *domain.Tweet) bool {
	var emojiRx = regexp.MustCompile("[\u1000-\uFFFF]+") // TODO encontrar una regex que funcione!
	return emojiRx.MatchString(tweet.Text)
}