package main

// Reads as much of a user's last 3200 public Tweets as the Twitter API
// returns, and prints each Tweet to a file.  This example functions
// the same as user_timeline, but uses application-only auth.

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
	"encoding/csv"
)

var totalTweets []CsvTweet
var NAME = os.Args[1]

func LoadCredentials() (client *twittergo.Client, err error) {
	credentials, err := ioutil.ReadFile("CREDENTIALS")
	if err != nil {
		return
	}
	lines := strings.Split(string(credentials), "\n")
	config := &oauth1a.ClientConfig{
		ConsumerKey:    lines[0],
		ConsumerSecret: lines[1],
	}
	client = twittergo.NewClient(config, nil)
	return
}

type Args struct {
	ScreenName string
	OutputFile string
}

func parseArgs() *Args {
	a := &Args{}
	flag.StringVar(&a.ScreenName, "screen_name", NAME, "Screen name")
	flag.StringVar(&a.OutputFile, "out", fmt.Sprintf("./Json/%v.json", NAME), "Output file")
	flag.Parse()
	return a
}

type CsvTweet struct {
	Fecha 		string 	`json:"created_at"`
	Usuario 	string 	`json:"name"`
	Texto 		string 	`json:"full_text"`
	Likes 		int 	`json:"favorite_count"`
	Retweets 	int 	`json:"retweet_count"`
}

func writeCsv(args *Args) error{
	filePath := fmt.Sprintf("./Tweets/tweets_%v.csv", NAME)
	csvFile, err := os.Create(filePath)
	fmt.Println(args.ScreenName)

	if err != nil {
		fmt.Println(err)
		return err
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	for _, tweets := range totalTweets {
		var record []string
		record = append(record, tweets.Fecha)
		record = append(record, NAME)
		record = append(record, strings.Join(strings.Split(tweets.Texto, "\n"), " "))
		record = append(record, strconv.Itoa(tweets.Likes))
		record = append(record, strconv.Itoa(tweets.Retweets))
		writer.Write(record)
	}

	// remember to flush!
	writer.Flush()
	return nil
}

func parseJson(tweet []byte) error{
	var readJson CsvTweet
	err := json.Unmarshal([]byte(tweet), &readJson)
	if err != nil {
		fmt.Println(err)
		return err
	}
	totalTweets = append(totalTweets, readJson)
	return nil
}

func main() {
	var (
		err     error
		client  *twittergo.Client
		req     *http.Request
		resp    *twittergo.APIResponse
		args    *Args
		max_id  uint64
		out     *os.File
		query   url.Values
		results *twittergo.Timeline
		text    []byte
	)
	args = parseArgs()
	if client, err = LoadCredentials(); err != nil {
		fmt.Printf("Could not parse CREDENTIALS file: %v\n", err)
		os.Exit(1)
	}
	if out, err = os.Create(args.OutputFile); err != nil {
		fmt.Printf("Could not create output file: %v\n", args.OutputFile)
		os.Exit(1)
	}
	defer out.Close()
	const (
		count   int = 200
		urltmpl     = "/1.1/statuses/user_timeline.json?%v"
		minwait     = time.Duration(10) * time.Second
	)
	query = url.Values{}
	query.Set("exclude_replies", "true")
	query.Set("include_rts", "false")
	query.Set("tweet_mode", "extended")
	query.Set("count", fmt.Sprintf("%v", count))
	query.Set("screen_name", args.ScreenName)
	total := 0
	for {
		if max_id != 0 {
			query.Set("max_id", fmt.Sprintf("%v", max_id))
		}
		endpoint := fmt.Sprintf(urltmpl, query.Encode())
		if req, err = http.NewRequest("GET", endpoint, nil); err != nil {
			fmt.Printf("Could not parse request: %v\n", err)
			os.Exit(1)
		}
		if resp, err = client.SendRequest(req); err != nil {
			fmt.Printf("Could not send request: %v\n", err)
			os.Exit(1)
		}
		results = &twittergo.Timeline{}
		if err = resp.Parse(results); err != nil {
			if rle, ok := err.(twittergo.RateLimitError); ok {
				dur := rle.Reset.Sub(time.Now()) + time.Second
				if dur < minwait {
					// Don't wait less than minwait.
					dur = minwait
				}
				msg := "Rate limited. Reset at %v. Waiting for %v\n"
				fmt.Printf(msg, rle.Reset, dur)
				time.Sleep(dur)
				continue // Retry request.
			} else {
				fmt.Printf("Problem parsing response: %v\n", err)
			}
		}
		batch := len(*results)
		if batch == 0 {
			fmt.Printf("No more results, end of timeline.\n")
			break
		}
		for _, tweet := range *results {
			if text, err = json.Marshal(tweet); err != nil {
				fmt.Printf("Could not encode Tweet: %v\n", err)
				os.Exit(1)
			}
			err = parseJson(text)
			if err != nil {
				fmt.Println("Hubo un error al parsear el JSON")
			}
			out.Write(text)
			out.Write([]byte("\n"))
			max_id = tweet.Id() - 1
			total += 1
		}
		err = writeCsv(args)
		if err != nil {
			fmt.Println("Hubo un error en la escritura")
		}

		fmt.Printf("Got %v Tweets from %v", batch, NAME)
		if resp.HasRateLimit() {
			fmt.Printf(", %v calls available", resp.RateLimitRemaining())
		}
		fmt.Printf(".\n")

	}
	fmt.Printf("--------------------------------------------------------\n")
	fmt.Printf("Wrote %v Tweets to %v\n in %v", total, args.OutputFile, urltmpl)
}