package main

import (
	"encoding/json"
	"fmt"
	"github.com/mh-cbon/ignore-file/ignored"
	"github.com/urfave/cli"
	"os"
	"path/filepath"
)

var VERSION = "0.0.0"

func main() {

	app := cli.NewApp()
	app.Name = "ignore-file"
	app.Version = VERSION
	app.Usage = "Test a ignore file"
	app.UsageText = "ignore-file <ignore file path> <options>"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "json,j",
			Usage: "return json response",
		},
		cli.StringSliceFlag{
			Name:  "also,a",
			Usage: "Add new rule",
		},
	}
	app.Action = run

	app.Run(os.Args)
}

func run(c *cli.Context) error {
	also := c.StringSlice("also")

	if c.NArg() == 0 {
		return cli.NewExitError("Missing <ignore file path> argument", 1)
	}

	file := c.Args().Get(0)
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return cli.NewExitError(err.Error(), 1)
	}
	dir := filepath.Dir(file)

	ignore := ignored.Ignored{}
	if err := ignore.Load(file); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	for _, r := range also {
		if err := ignore.Append(r); err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
	}

	computed := ignore.ComputeDirectory(dir)

	if c.Bool("json") {
		jsoned, err := json.MarshalIndent(computed, "", "  ")
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		fmt.Print(string(jsoned))
	} else {
		for _, l := range computed {
			fmt.Println(l)
		}
	}
	return nil
}
