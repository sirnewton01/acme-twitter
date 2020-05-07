package main

import (
        "flag"
        "fmt"
        "text/template"
	"os"

        "github.com/coreos/pkg/flagutil"
        "github.com/dghubble/go-twitter/twitter"
)

const tweetTmpl = `__________________________
{{ .CreatedAtTime.Local.Format "Mon Jan 2 15:04:05" }} (Tuser {{.User.ScreenName}}) <https://twitter.com/{{.User.ScreenName}}/status/{{.IDStr}}>

{{ .FullText }}
{{ if .Entities }}{{ range .Entities.Urls }}* <{{ .ExpandedURL }}>
{{ end }}{{ end }}
(Trepl {{ .IDStr }} '@{{.User.ScreenName}} '):{{ .ReplyCount }} (Trt {{ .IDStr }}):{{ .RetweetCount }} (Tfav {{ .IDStr }}):{{ .FavoriteCount }}

`

func timeline (client *twitter.Client) {
        flags := struct {
                count          int
                since          int64
        }{}
        flag.IntVar(&flags.count, "count", 10, "Number of tweets to fetch")
        flag.Int64Var(&flags.since, "since", 0, "Retrieve tweets since this tweet ID")
        flag.Parse()
        flagutil.SetFlagsFromEnv(flag.CommandLine, "TWITTER")

        // home timeline
        timelineParams := &twitter.HomeTimelineParams{Count: flags.count, TweetMode: "extended", SinceID: flags.since}
        tweets, _, err := client.Timelines.HomeTimeline(timelineParams)
        if err != nil {
                panic(err)
        }
        tmpl, err := template.New("tweet").Parse(tweetTmpl)
        if err != nil {
                panic(err)
        }

        lastID := fmt.Sprintf("%d", flags.since)
        for idx := len(tweets)-1; idx >=0; idx-- {
                tweet := tweets[idx]
                err = tmpl.Execute(os.Stdout, tweet)
                if err != nil {
                        panic(err)
                }
                lastID = tweet.IDStr
        }

        fmt.Printf("__________________________\n")
        fmt.Printf("<Ttl -since %s", lastID)
}
