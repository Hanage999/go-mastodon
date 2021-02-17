package main

import (
	"context"
	"errors"
	"fmt"

<<<<<<< HEAD
	"github.com/mattn/go-mastodon"
	"github.com/urfave/cli/v2"
=======
	"github.com/hanage999/go-mastodon"
	"github.com/urfave/cli"
>>>>>>> 6d9f14f (conform to go module)
)

func cmdToot(c *cli.Context) error {
	var toot string
	ff := c.String("ff")
	if ff != "" {
		text, err := readFile(ff)
		if err != nil {
			return err
		}
		toot = string(text)
	} else {
		if !c.Args().Present() {
			return errors.New("arguments required")
		}
		toot = argstr(c)
	}
	client := c.App.Metadata["client"].(*mastodon.Client)
	_, err := client.PostStatus(context.Background(), &mastodon.Toot{
		Status:      toot,
		InReplyToID: mastodon.ID(fmt.Sprint(c.String("i"))),
	})
	return err
}
