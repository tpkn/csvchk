package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"unicode/utf8"
)

var env = "development"
var version = "0.0.0"
var help = fmt.Sprintf(Help, version)

type Args struct {
	Cumulative bool
	Quiet      bool
	Delimiter  string
	Help       bool
	Version    bool
}

func main() {
	var args = Args{}
	flag.BoolVar(&args.Cumulative, "c", false, "Collect all csv errors and output the list at the end")
	flag.BoolVar(&args.Quiet, "q", false, "Silently terminate with exit(1) upon the first error encountered in the csv")
	flag.StringVar(&args.Delimiter, "d", ",", "Fields separator (default: comma)")
	flag.BoolVar(&args.Help, "h", false, "Help")
	flag.BoolVar(&args.Help, "help", false, "Alias for -h")
	flag.BoolVar(&args.Version, "v", false, "Version")
	flag.BoolVar(&args.Version, "version", false, "Alias for -v")
	flag.Parse()

	log.SetFlags(0)
	log.SetPrefix("csvchk: ")
	runtime.GOMAXPROCS(2)

	if args.Help {
		fmt.Println(help)
		os.Exit(0)
	}

	if args.Version {
		fmt.Println(version)
		os.Exit(0)
	}

	var errors_list strings.Builder

	var reader = csv.NewReader(os.Stdin)
	delimiter, _ := utf8.DecodeRuneInString(args.Delimiter)
	reader.Comma = delimiter
	reader.ReuseRecord = true

	for {
		_, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			if args.Quiet {
				os.Exit(1)
			} else {
				if args.Cumulative {
					errors_list.WriteString(err.Error() + "\n")
				} else {
					fmt.Print(err.Error())
					os.Exit(1)
				}
			}
		}
	}

	if !args.Quiet && args.Cumulative && errors_list.Len() > 0 {
		fmt.Print(errors_list.String())
		os.Exit(1)
	}
}
