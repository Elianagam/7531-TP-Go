import twitter

"""
Downloads all tweets from a given user.
Uses twitter.Api.GetUserTimeline to retreive the last 3,200 tweets from a user.
Twitter doesn't allow retreiving more tweets than this through the API, so we get
as many as possible.
t.py should contain the imported variables.
"""

import json
import sys
import csv
import os

import twitter

CONSUMER_KEY = 'vlCeSQPk3vbxgHOLDRHW4HyEF'
CONSUMER_SECRET = 'sHeBkkGgD05hFNHK2IdcbE17BXXjah6CvAkI21WAq6ta6DDP96'
ACCESS_TOKEN_KEY = '1273577244-6wUVJKr8LEbSy7ZdeRbNF9GwVfYcrerAzvktJYw'
ACCESS_TOKEN_SECRET = 'XUX5UJpyyUZoQCpBulPjAgq7PAJbIWheFkjTlEN8zzGcL'


def get_tweets(api=None, screen_name=None, count=None):
    timeline = api.GetUserTimeline(screen_name=screen_name, count=count)
    for tweet in timeline:
        try:
            created_at = tweet.created_at
            username = screen_name
            text = " ".join(tweet.full_text.split("\n"))
            likes = tweet.favorite_count
            retweets = tweet.retweet_count
            hashtags = tweet.hashtags

            write_csv([created_at, username, text, likes, retweets, hashtags])
        except Exception as err:
            print(err)
            pass


def write_csv(list):
    if not os.path.isfile(f'tweets/tweets_{list[1]}.csv'):
        with open(f'tweets/tweets_{list[1]}.csv', mode='w') as csv_file:
            fieldnames = ['created_at', 'username', 'text', 'likes', 'retweets', 'hashtags']
            writer = csv.DictWriter(csv_file, fieldnames=fieldnames)
            writer.writeheader()

    with open(f'tweets/tweets_{list[1]}.csv', mode='a+') as csv_file:
        fieldnames = ['created_at', 'username', 'text', 'likes', 'retweets', 'hashtags']
        writer = csv.DictWriter(csv_file, fieldnames=fieldnames)

        writer.writerow({'created_at': list[0],
                         'username': list[1],
                         'text': list[2],
                         'likes': list[3],
                         'retweets': list[4],
                         'hashtags': list[5]})

if __name__ == "__main__":
    api = twitter.Api(
        CONSUMER_KEY, CONSUMER_SECRET, ACCESS_TOKEN_KEY, ACCESS_TOKEN_SECRET, tweet_mode='extended'
    )
    screen_name = sys.argv[1]
    count = sys.argv[2]
    timeline = get_tweets(api=api, screen_name=screen_name, count=count)

