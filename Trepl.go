package main

import (
        "flag"
	"strconv"

        "github.com/coreos/pkg/flagutil"
        "github.com/dghubble/go-twitter/twitter"
)

func reply (client *twitter.Client) {
        //flags := struct {
        //}{}
        flag.Parse()
        flagutil.SetFlagsFromEnv(flag.CommandLine, "TWITTER")

	// Reply
	tweetId, err := strconv.ParseInt(flag.Arg(0), 10, 64)
	if err != nil {
		panic(err)
	}

	status := flag.Arg(1)

	_, _, err = client.Statuses.Update(status, &twitter.StatusUpdateParams{InReplyToStatusID: tweetId})
	if err != nil {
		panic(err)
	}
}
