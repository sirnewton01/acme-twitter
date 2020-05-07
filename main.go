package main

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/go-twitter/twitter"

	"gopkg.in/yaml.v2"
)

type cfg struct {
	Profiles map[string]map[string] struct {
		ConsumerKey string `yaml:"consumer_key"`
		ConsumerSecret string `yaml:"consumer_secret"`
		Token string `yaml:"token"`
		Secret string `yaml:"secret"`
	} `yaml:"profiles"`
}

func main() {
	// FIXME find the home directory some other way
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	trc, err := os.Open(filepath.Join(usr.HomeDir, ".trc"))
	if err != nil {
		panic(err)
	}
	defer trc.Close()

	trcbytes, err := ioutil.ReadAll(trc)
	if err != nil {
		panic(err)
	}


	consumerKey := ""
	consumerSecret := ""
	token := ""
	secret := ""

	c := cfg{}
	err = yaml.Unmarshal(trcbytes, &c)
	for _, p := range c.Profiles {
		for _, q := range p {
			consumerKey = q.ConsumerKey
			consumerSecret = q.ConsumerSecret
			token = q.Token
			secret = q.Secret
		}
	}

	config := oauth1.NewConfig(
		consumerKey,
		consumerSecret,
	)
	accToken := oauth1.NewToken(token, secret)
	httpClient := config.Client(oauth1.NoContext, accToken)

	// Twitter client
	client := twitter.NewClient(httpClient)

	if os.Args[0] == "Trt" {
        	retweet(client)
	} else if os.Args[0] == "Tfav" {
        	favorite(client)
	} else if os.Args[0] == "Trepl" {
        	reply(client)
	} else if os.Args[0] == "Tweet" {
                tweet(client)
	} else {
		timeline(client)
	}
}
