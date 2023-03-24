package main

import (
	"fmt"
	"github.com/kintar/filecleaner/model"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/v2"
	flag "github.com/spf13/pflag"
	"log"
	"os"
)

var k = koanf.New(".")

func main() {
	f := flag.NewFlagSet("config", flag.ContinueOnError)
	f.Usage = func() {
		fmt.Println("Scans a path for duplicate files, based on checksum")
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}
	f.StringP("path", "p", "", "path to scan")
	f.BoolP("recursive", "r", false, "if set, recurse into subdirectories")
	f.StringSliceP("exclude-dirs", "e", nil, "names of paths to exclude from scan")
	f.StringP("output", "o", "duplicates.json", "output file")
	err := f.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("failed to parse command line flags", err)
		os.Exit(0)
	}

	if err = k.Load(posflag.Provider(f, ".", nil), nil); err != nil {
		log.Fatal(err)
	}

	cfg := model.Config{}
	err = k.Unmarshal("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

}
