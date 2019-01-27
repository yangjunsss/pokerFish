package main

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/yangjunsss/pokerfish/base"
	"os"
	"strings"
)

/*
utg:
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 26 27 28 29 39 42 43 52 56 57 70 71 84 85 98 99 112 126 140 154 168 30
mp:
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 26 27 28 29 39 42 43 52 56 57 70 71 84 85 98 99 112 126 140 154 168 30 18 31 44 40
co:
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 26 27 28 29 39 42 43 52 56 57 70 71 84 85 98 99 112 126 140 154 168 30 18 31 44 40 41 53 54 55 19 32 45 58 20 21 22 33 34 65 72 86 113
btn:
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 26 27 28 29 39 42 43 52 56 57 70 71 84 85 98 99 112 126 140 154 168 30 18 31 44 40 41 53 54 55 19 32 45 58 20 21 22 33 34 65 72 86 113 23 24 25 35 36 37 38 46 47 48 49 50 51 59 60 61 73 74 87 100 101 114 127 128 141 66 67
*/
var setRangeCommand = cli.Command{
	Name:  "setrange",
	Usage: "Set echo position openrange",
	ArgsUsage: `
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
		if err := base.CheckArgs(c, 0, 20); err != nil {
			return err
		}
		pos := c.String("p")
		path := c.String("c")
		logrus.Info(pos, path)
		conf, err := base.ReadConfig(path)

		if err != nil {
			return err
		}

		base.PrintCards(func(idx string) bool {
			return true
		})

		ranges := scanRange()
		setRange(pos, ranges, conf)
		logrus.Info(pos, ranges, conf)
		if base.WriteConf(conf, path); err != nil {
			return err
		}
		return nil
	},
}

func scanRange() map[string]int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	input0 := strings.Split(input, " ")
	ranges := make(map[string]int)
	for _, v := range input0 {
		ranges[v] = 0
	}
	return ranges
}

func setRange(pos string, ranges map[string]int, conf *base.Config) {
	if conf.OpenRange == nil {
		conf.OpenRange = []base.Range{
			base.Range{
				Name:  pos,
				Range: ranges,
			},
		}
	} else {
		for i, openrange := range conf.OpenRange {
			if openrange.Name == pos {
				conf.OpenRange[i].Range = ranges
				break
			}
			if i == len(conf.OpenRange)-1 {
				conf.OpenRange = append(conf.OpenRange, base.Range{
					Name:  pos,
					Range: ranges,
				})
			}
		}
	}
}
