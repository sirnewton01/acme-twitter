package main

import (
        "flag"

        "github.com/coreos/pkg/flagutil"
        "github.com/dghubble/go-twitter/twitter"
)

func tweet (client *twitter.Client) {
        //flags := struct {
        //}{}
        flag.Parse()
        flagutil.SetFlagsFromEnv(flag.CommandLine, "TWITTER")

	status := flag.Arg(0)

	_, _, err := client.Statuses.Update(status, &twitter.StatusUpdateParams{})
	if err != nil {
		panic(err)
	}
}
