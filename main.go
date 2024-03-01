package main

import (
	"compress/gzip"
	"fmt"
	"os"
	"strings"

	clipboard "github.com/atotto/clipboard"
	"github.com/spf13/pflag"
	"github.com/yudppp/json2struct"
)

var version string

func main() {
	var opt json2struct.Options
	var parsed, file, outFile string
	var debug, help bool
	var err error
	pflag.BoolVarP(&debug, "debug", "d", false, "Set debug mode")
	pflag.BoolVarP(&opt.UseOmitempty, "omitempty", "O", false, "Set omitempty mode")
	pflag.BoolVarP(&opt.UseShortStruct, "short", "S", false, "Set short struct name mode")
	pflag.BoolVarP(&opt.UseLocal, "local", "l", false, "Use local struct mode")
	pflag.BoolVarP(&opt.UseExample, "example", "e", false, "Use example tag mode")
	pflag.BoolVarP(&help, "help", "h", false, "Help")
	pflag.StringVarP(&opt.Prefix, "prefix", "p", "", "Set struct name prefix")
	pflag.StringVarP(&opt.Suffix, "suffix", "s", "", "Set struct name suffix")
	pflag.StringVarP(&opt.Name, "name", "n", "", "Set struct name")
	pflag.StringVarP(&file, "file", "f", "", "JSON File to parse")
	pflag.StringVarP(&outFile, "file-out", "o", "", "File to be written. This will be the entire struct(s) in a complete .go file.")
	pflag.Parse()
	if help {
		fmt.Println(os.Args[0], "Version:", version, "\nUsage:")
		pflag.PrintDefaults()
		fmt.Println(os.Args[0], "Can be used to generate a go struct. Ex:", os.Args[0], "-f filein.json -o fileout.go -n \"struct name\"")
		os.Exit(1)
	}
	json2struct.SetDebug(debug)
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		//jsonBod, _ = ioutil.ReadAll(os.Stdin)
		parsed, err = json2struct.Parse(os.Stdin, opt)
	} else if file != "" {
		//jsonBod := ReadInFile(file)
		fmt.Println("Opening", file)
		jsonFile, err := os.Open(file)
		errorHandle(err, "os.Open", true)
		gzipReader, err := gzip.NewReader(jsonFile)
		if err != nil {
			parsed, err = json2struct.Parse(jsonFile, opt)
			errorHandle(err, "json2struct.Parse", true)
		} else {
			parsed, err = json2struct.Parse(gzipReader, opt)
			errorHandle(err, "json2struct.Parse gzip", true)
		}

	} else {
		//Read from clipboard
		text, err := clipboard.ReadAll()
		errorHandle(err, "Clipboard Read All", true)
		txtRdr := strings.NewReader(text)
		parsed, err = json2struct.Parse(txtRdr, opt)
		errorHandle(err, "error - is string valid JSON?", true)
	}
	if err != nil {
		panic(err)
	}
	if outFile != "" {
		parsed = fmt.Sprintf("package main\n\n%s", parsed)
		os.WriteFile(outFile, []byte(parsed), 0644)
	} else {
		fmt.Println(parsed)
	}
}

func errorHandle(err error, str string, ex bool) {
	if err != nil {
		fmt.Println(str, "error:", err)
		if ex {
			os.Exit(1)
		}
	}
}
