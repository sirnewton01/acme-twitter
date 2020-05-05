package main

import (
        "flag"
	"strconv"

        "github.com/coreos/pkg/flagutil"
        "github.com/dghubble/go-twitter/twitter"
)

func favorite (client *twitter.Client) {
        //flags := struct {
        //}{}
        flag.Parse()
        flagutil.SetFlagsFromEnv(flag.CommandLine, "TWITTER")

	// Retweet

	tweetId, err := strconv.ParseInt(flag.Arg(0), 10, 64)
	if err != nil {
		panic(err)
	}

	_, _, err = client.Favorites.Create(&twitter.FavoriteCreateParams{ID: tweetId})
	if err != nil {
		panic(err)
	}
}
