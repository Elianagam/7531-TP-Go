package repository

import (
	"encoding/csv"
	"fmt"
	"github.com/Nicobugliot/7531-TP-Go/twitter/domain"
	"os"
	"strconv"
)

// Interface

type TwitterRepository interface {
	GetTweetsFromUser(user string) ([]*domain.Tweet, error)
}


// Implementations

type FileTwitterRepository struct {}

func NewFileTwitterRepository() *FileTwitterRepository {
	return &FileTwitterRepository{}
}

const DataPath = "/Users/akler/go/src/github.com/Nicobugliot/7531-TP-Go/twitterScraper/tweets/tweets_%s.csv"

func (f *FileTwitterRepository) GetTweetsFromUser(user string) ([]*domain.Tweet, error) {
	path := fmt.Sprintf(DataPath, user)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var tweets []*domain.Tweet

	for _, line := range lines {
		likes, _ := strconv.Atoi(line[3])
		retweets, _ := strconv.Atoi(line[4])

		tweets = append(tweets, &domain.Tweet{
			CreatedAt: line[0],
			User:      user,
			Text:      line[2],
			Likes:     likes,
			Retweets:  retweets,
		})
	}

	return tweets, nil
}
