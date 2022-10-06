package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rmasci/json2struct"
	"github.com/spf13/pflag"
)

func main() {
	var opt json2struct.Options
	pflag.BoolVarP(&opt.Debug, "debug", "d", false, "Set debug mode")
	pflag.BoolVarP(&opt.Omitempty, "omitempty", "o", false, "Set omitempty mode")
	pflag.BoolVarP(&opt.Short, "short", "s", false, "Set short struct name mode")
	pflag.BoolVarP(&opt.Local, "local", "l", false, "Use local struct mode")
	pflag.BoolVarP(&opt.Example, "example", "e", false, "Use example tag mode")
	pflag.String(&opt.Prefix, "prefix", "p", "", "Set struct name prefix")
	pflag.String(&opt.Suffix, "suffix", "s", "", "Set struct name suffix")
	pflag.String(&opt.Name, "name", "n", json2struct.DefaultStructName, "Set struct name")
	flag.Parse()
	json2struct.SetDebug(*debug)
	opt.Name = strings.ToLower(opt.Name)
	parsed, err := json2struct.Parse(os.Stdin, opt)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsed)
}
