package main

import (
	"github.com/Nicobugliot/7531-TP-Go/twitter/domain"
	"github.com/Nicobugliot/7531-TP-Go/twitter/search"
	"github.com/Nicobugliot/7531-TP-Go/twitter/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	// Define endpoints
	router.GET("/tweets/search", searchRouter)
	router.GET("/tweets/search/:userId", searchByUserRouter)

	router.Run()
}

func searchByUserRouter(c *gin.Context) {
	user := c.Param("userId")
	if !isUserAllowed(user) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or not allowed user id."})
		return
	}

	searchRouterCommon(c, []string{user})
}

func searchRouter(c *gin.Context) {
	searchRouterCommon(c, getAllowedUsers())
}

func searchRouterCommon(c *gin.Context, users []string) {
	query := c.Query("query")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Param 'query' is mandatory."})
		return
	}

	if response, err := searchTweets(query, users); err != nil {
		c.JSON(http.StatusInternalServerError, "Error trying to execute search.")
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func searchTweets(query string, users []string) (searchTweetsResponse, error) {

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
	Count   int             `json:"count"`
	Results []*domain.Tweet `json:"results"`
}

func getAllowedUsers() []string {
	return []string{"alferdez", "mauriciomacri", "jlespert", "NicolasdelCano", "juanjomalvinas"}
}

func isUserAllowed(userId string) bool {
	for _, user := range getAllowedUsers() {
		if userId == user {
			return true
		}
	}
	return false
}
