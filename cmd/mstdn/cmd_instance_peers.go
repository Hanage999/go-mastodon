package main

import (
	"context"
	"fmt"

<<<<<<< HEAD
	"github.com/mattn/go-mastodon"
	"github.com/urfave/cli/v2"
=======
	"github.com/hanage999/go-mastodon"
	"github.com/urfave/cli"
>>>>>>> 6d9f14f (conform to go module)
)

func cmdInstancePeers(c *cli.Context) error {
	client := c.App.Metadata["client"].(*mastodon.Client)
	peers, err := client.GetInstancePeers(context.Background())
	if err != nil {
		return err
	}
	for _, peer := range peers {
		fmt.Fprintln(c.App.Writer, peer)
	}
	return nil
}
