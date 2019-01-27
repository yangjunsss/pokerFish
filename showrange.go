package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/yangjunsss/pokerfish/base"
	"regexp"
	"strings"
)

var showRangeCommand = cli.Command{
	Name:  "showrange",
	Usage: "Show echo position openrange",
	ArgsUsage: `<position>
    sb,bb,utg,mc,co,btn
  `,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "position,p",
			Value: "",
			Usage: "position",
		},
		cli.StringFlag{
			Name:  "config,c",
			Value: base.ConfPath,
			Usage: "Re-wirte the config file",
		},
	},
	Action: func(c *cli.Context) error {
		var err error
		if err = base.CheckArgs(c, 0, 10); err != nil {
			return err
		}
		pos := c.String("p")
		path := c.String("c")
		logrus.Info(pos, path)
		conf, err := base.ReadConfig(path)

		if err != nil {
			return err
		}
		showRange(pos, conf.OpenRange)
		return err
	},
}

func showRange(pos string, openRanges []base.Range) {
	for _, openRange := range openRanges {
		if pos != "" && pos != openRange.Name {
			continue
		}
		logrus.Info(openRange.Name)
		var exp *regexp.Regexp
		var err error
		if exp, err = regexp.Compile(openRange.Regexp); err != nil {
			logrus.Error(err)
		}
		base.PrintCards(func(v string) bool {
			if err != nil {
				return false
			}
			src := strings.TrimSpace(v)
			_, ok := openRange.Range[src]
			return exp.MatchString(src) || ok
		})
	}
}
