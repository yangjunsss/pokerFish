package base

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"os"
)

var cRed = color.New(color.FgRed)
var cGreen = color.New(color.FgGreen)
var cYellow = color.New(color.FgYellow)
var cBlack = color.New(color.FgBlack)

func CheckArgs(c *cli.Context, min, max int) error {
	var err error
	cmdName := c.Command.Name
	if c.NArg() < min || c.NArg() > max {
		err = fmt.Errorf("%s: %q require a minimum of %d and a maximum of %d arguments,%d", os.Args[0], cmdName, min, max, c.NArg())
	}

	if err != nil {
		cli.ShowCommandHelp(c, cmdName)
		return err
	}
	return nil
}

func PrintCards(bPrint func(idx string) bool) {
	for i, l := 0, len(Cards); i < l; i++ {
		for j := 0; j < l; j++ {
			var v string
			if i == j {
				v = fmt.Sprintf("%s%s  ", Cards[i], Cards[i])
				if bPrint(v) {
					cRed.Print(v)
				} else {
					cBlack.Print(v)
				}
			} else if i > j {
				v = fmt.Sprintf("%s%so ", Cards[j], Cards[i])
				if bPrint(v) {
					cYellow.Print(v)
				} else {
					cBlack.Print(v)
				}
			} else if i < j {
				v = fmt.Sprintf("%s%ss ", Cards[i], Cards[j])
				if bPrint(v) {
					cGreen.Print(v)
				} else {
					cBlack.Print(v)
				}
			}
		}
		fmt.Printf("\n")
	}
}
