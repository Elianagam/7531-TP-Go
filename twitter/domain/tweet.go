package domain

import "fmt"

type Tweet struct {
	CreatedAt string `json:"created_at"`
	User string `json:"user"`
	Text string `json:"text"`
	Likes int `json:"likes"`
	Retweets int `json:"retweets"`
}

func (t *Tweet) ToString() string {
	return fmt.Sprintf(
		"@%s:\n%s\n%d likes | %d retweets | %s",
		t.User, t.Text, t.Likes, t.Retweets, t.CreatedAt)
}