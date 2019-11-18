package domain

import "fmt"

type Tweet struct {
	CreatedAt string
	User string
	Text string
	Likes int
	Retweets int
}

func (t *Tweet) ToString() string {
	return fmt.Sprintf(
		"@%s:\n%s\n%d likes | %d retweets | %s",
		t.User, t.Text, t.Likes, t.Retweets, t.CreatedAt)
}