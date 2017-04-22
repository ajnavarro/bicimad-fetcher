package main

import (
	"github.com/ajnavarro/bicimad-fetcher"

	"time"
)

type fetchCmd struct {
	ID           string `short:"i" long:"id" optional:"no" description:"bicimad ID"`
	Key          string `short:"k" long:"key" optional:"no" description:"bicimad key"`
	DatabasePath string `short:"d" long:"database-path" optional:"no" description:"sqlite path"`
	Tts time.Duration `short:"t" long:"time" optional:"no" description:"Time to Sleep between calls"`
}

func (c *fetchCmd) Execute(args []string) error {
	client := bicimad_fetcher.NewClient(c.ID, c.Key)

	dao, err := bicimad_fetcher.NewStationsDAO(c.DatabasePath)
	if err != nil {
		return err
	}

	f := &bicimad_fetcher.Fetcher{
		Client: client,
		DAO:    dao,
		TTS:   c.Tts,
	}

	return f.Fetch()
}
