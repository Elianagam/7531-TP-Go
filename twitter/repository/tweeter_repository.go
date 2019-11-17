package repository

import (
	"encoding/csv"
	"fmt"
	"github.com/Nicobugliot/7531-TP-Go/twitter/domain"
	"os"
)

type TwitterRepository interface {
	GetTweetsFromUser(user string) ([]*domain.Tweet, error)
}


type FileTwitterRepository struct {

}

func NewFileTwitterRepository() *FileTwitterRepository {
	return &FileTwitterRepository{}
}

const DataPath  = "/Users/akler/go/src/github.com/Nicobugliot/7531-TP-Go/twitter/repository/data/%s.csv"

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

	for _,line := range lines  {
		tweets = append(tweets, &domain.Tweet{
			User: user,
			Text: line[0],
		})
	}

	return tweets, nil
}