package main

import (
	"github.com/Nicobugliot/7531-TP-Go/twitter/domain"
)

func main()  {

	//users := []string{"alferdez", "mauriciomacri"}



	// TODO Realizar la búsqueda


	// TODO Imprimir resultados



}



func getTweetsFromUser(channel chan *domain.Tweet, user string) {

	// TODO instanciar repositorio

	// TODO Obtener los tweets del usuario

	// TODO Enviar tweets
}

func processTweets(resultChannel chan string, tweetsToProcess chan *domain.Tweet, query string)  {
	// TODO procesar
}