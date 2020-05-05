package main

import (
        "flag"
	"strconv"

        "github.com/coreos/pkg/flagutil"
        "github.com/dghubble/go-twitter/twitter"
)

func retweet (client *twitter.Client) {
        //flags := struct {
        //}{}
        flag.Parse()
        flagutil.SetFlagsFromEnv(flag.CommandLine, "TWITTER")

	// Retweet

	tweetId, err := strconv.ParseInt(flag.Arg(0), 10, 64)
	if err != nil {
		panic(err)
	}

	_, _, err = client.Statuses.Retweet(tweetId, &twitter.StatusRetweetParams{})
	if err != nil {
		panic(err)
	}
}
