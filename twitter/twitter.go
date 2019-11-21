package main

import (
	"github.com/Nicobugliot/7531-TP-Go/twitter/domain"
	"github.com/Nicobugliot/7531-TP-Go/twitter/search"
	"github.com/Nicobugliot/7531-TP-Go/twitter/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {

	router := gin.Default()

	// Define endpoints
	router.GET("/tweets/search", searchRouter)

	router.Run()
}

func searchRouter(c *gin.Context) {
	query := c.Query("query")

	if query == "" {
		c.JSON(http.StatusBadRequest, "Param 'query' is mandatory.")
		return
	}

	if response,err := searchTweets(query); err != nil {
		c.JSON(http.StatusInternalServerError, "Error trying to execute search.")
	}else {
		c.JSON(http.StatusOK, response)
	}
}

func searchTweets(query string) (searchTweetsResponse, error)  {
	users := []string {"alferdez", "mauriciomacri"}

	resultChannel := make(chan *domain.Tweet)

	go search.Search(resultChannel, users, utils.ContainsQuery(query))

	response := searchTweetsResponse{}

	for tweet := range resultChannel {
		response.Count++
		response.Results = append(response.Results, tweet)
	}

	return response, nil
}

type searchTweetsResponse struct {
	Count   int      `json:"count"`
	Results []*domain.Tweet `json:"results"`
} 