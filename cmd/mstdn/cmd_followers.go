package main

import (
	"context"
	"fmt"
	"time"

<<<<<<< HEAD
	"github.com/mattn/go-mastodon"
	"github.com/urfave/cli/v2"
=======
	"github.com/hanage999/go-mastodon"
	"github.com/urfave/cli"
>>>>>>> 6d9f14f (conform to go module)
)

func cmdFollowers(c *cli.Context) error {
	client := c.App.Metadata["client"].(*mastodon.Client)
	config := c.App.Metadata["config"].(*mastodon.Config)

	account, err := client.GetAccountCurrentUser(context.Background())
	if err != nil {
		return err
	}
	var followers []*mastodon.Account
	var pg mastodon.Pagination
	for {
		fs, err := client.GetAccountFollowers(context.Background(), account.ID, &pg)
		if err != nil {
			return err
		}
		followers = append(followers, fs...)
		if pg.MaxID == "" {
			break
		}
		pg.SinceID = ""
		pg.MinID = ""
		time.Sleep(10 * time.Second)
	}
	s := newScreen(config)
	for _, follower := range followers {
		fmt.Fprintf(c.App.Writer, "%v,%v\n", follower.ID, s.acct(follower.Acct))
	}
	return nil
}
