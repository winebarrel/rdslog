package main

import (
	"context"
	"log"
	"os"

	"github.com/alecthomas/kong"
	"github.com/winebarrel/rdslog"
)

var version string

func init() {
	log.SetFlags(0)
}

func parseArgs() *rdslog.Options {
	var cli struct {
		rdslog.Options
		Version kong.VersionFlag
	}

	parser := kong.Must(&cli, kong.Vars{"version": version})
	parser.Model.HelpFlag.Help = "Show help."
	_, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)

	return &cli.Options
}

func main() {
	ctx := context.Background()
	options := parseArgs()
	client, err := rdslog.NewClient(ctx, options)

	if err != nil {
		log.Fatal(err)
	}

	err = client.DownloadCompleteLogFile(ctx, os.Stdout)

	if err != nil {
		log.Fatal(err)
	}
}
