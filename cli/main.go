package main

import (
	"github.com/jessevdk/go-flags"
	"os"
)

func main(){
	parser := flags.NewParser(nil, flags.Default)

	if _, err := parser.AddCommand("fetch", "",
		"", &fetchCmd{}); err != nil {
		panic(err)
	}

	_, err := parser.Parse()
	if err != nil {
		if _, ok := err.(*flags.Error); ok {
			parser.WriteHelp(os.Stdout)
		}

		os.Exit(1)
	}
}
