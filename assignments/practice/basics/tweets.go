package main

import "os"

var apiKey = os.Getenv("TWITTER_BEARER_TOKEN")

type Tweet = struct{ Text string }
