package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"io"
	"os"
)

var version = ""
var usage = ""

func main() {
	app := cli.NewApp()
	app.Name = "pokerFish"
	app.Usage = usage
	app.Version = version
	app.Commands = []cli.Command{
		showRangeCommand,
		setRangeCommand,
	}
	cli.ErrWriter = &FatalWriter{cli.ErrWriter}
	if err := app.Run(os.Args); err != nil {
		fatal(err)
	}
	os.Exit(0)
}

type FatalWriter struct {
	cliErrWriter io.Writer
}

func (f *FatalWriter) Write(p []byte) (n int, err error) {
	logrus.Error(string(p))
	return f.cliErrWriter.Write(p)
}

func fatal(err error) {
	logrus.Error(err)
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
