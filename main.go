package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yudppp/json2struct"
)

var version string

func main() {
	var opt json2struct.Options
	var parsed, file, outFile string
	var debug, help bool
	var err error
	var flagset pflag.FlagSet
	flagset.BoolVarP(&debug, "debug", "d", false, "Set debug mode")
	flagset.BoolVarP(&opt.UseOmitempty, "omitempty", "O", false, "Set omitempty mode")
	flagset.BoolVarP(&opt.UseShortStruct, "short", "S", false, "Set short struct name mode")
	flagset.BoolVarP(&opt.UseLocal, "local", "l", false, "Use local struct mode")
	flagset.BoolVarP(&opt.UseExample, "example", "e", false, "Use example tag mode")
	flagset.BoolVarP(&help, "help", "h", false, "Help")
	flagset.StringVarP(&opt.Prefix, "prefix", "p", "", "Set struct name prefix")
	flagset.StringVarP(&opt.Suffix, "suffix", "s", "", "Set struct name suffix")
	flagset.StringVarP(&opt.Name, "name", "n", json2struct.DefaultStructName, "Set struct name")
	flagset.StringVarP(&file, "file", "f", "", "JSON File to parse")
	flagset.StringVarP(&outFile, "file-out", "o", "", "File to be written. This will be the entire struct(s) in a complete .go file.")
	flagset.Parse(os.Args[1:])
	if help {
		fmt.Println(os.Args[0], "Version:", version, "\nUsage:")
		flagset.PrintDefaults()
		fmt.Println(os.Args[0], "Can be used to generate a go struct. Ex:", os.Args[0], "-f filein.json -o fileout.go -n \"struct name\"")
		os.Exit(1)
	}
	json2struct.SetDebug(debug)
	opt.Name = strings.ToLower(opt.Name)
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		//jsonBod, _ = ioutil.ReadAll(os.Stdin)
		parsed, err = json2struct.Parse(os.Stdin, opt)
	} else {
		//jsonBod := ReadInFile(file)
		jsonBod, err := os.Open(file)
		errorHandle(err, "os.Open", true)
		parsed, err = json2struct.Parse(jsonBod, opt)
		errorHandle(err, "json2struct.Parse", true)
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

// func ReadInFile(inFile string) (bod []byte) {
// 	var f os.File
// 	var err error
// 	f, err = os.OpenFile(inFile, os.O_RDONLY, 0644)
// 	errorHandle(err, "Can't Open the File "+inFile, true)
// 	defer f.Close()
// 	// Detect type of file
// 	buff := make([]byte, 512)
// 	_, err = f.Read(buff)
// 	errorHandle(err, "Read File buff", true)
// 	filetype := http.DetectContentType(buff)
// 	switch filetype {
// 	case "application/x-gzip":
// 		f.Seek(0, 0)
// 		fgz, err := gzip.NewReader(f)
// 		errorHandle(err, "Gzip in", true)
// 		bod, err = ioutil.ReadAll(fgz)
// 		errorHandle(err, "Read GZip File", true)
// 		return bod

// 	case "application/zip":
// 		z, err := zip.OpenReader(inFile)
// 		errorHandle(err, "Reading zip file...", true)
// 		defer z.Close()
// 		for _, f := range z.File {
// 			//c is contents
// 			c, err := f.Open()
// 			errorHandle(err, "Reading file in zip", true)
// 			_, err = c.Read(buff)
// 			filetype = http.DetectContentType(buff)
// 			if strings.Contains(filetype, "text/plain") {
// 				//f.Close()
// 				c, _ := f.Open()
// 				bod, err := ioutil.ReadAll(c)
// 				errorHandle(err, "Read in File", true)
// 				return bod
// 			}
// 			fmt.Printf("File: %v, filetype: %v\n ", f.Name, filetype)
// 		}
// 		fmt.Printf("No JSON or XML file in zip %v, %v\n", inFile, filetype)
// 		os.Exit(1)
// 	default:
// 		f.Seek(0, 0)
// 		bod, err := ioutil.ReadAll(f)
// 		errorHandle(err, "Read in File", true)
// 		head := make([]byte, 512)
// 		if len(bod) >= 512 {
// 			head = bod[512:]
// 		} else {
// 			head = bod
// 		}
// 		filetype = http.DetectContentType(head)
// 		return bod
// 	}
// 	// if we're this far return and empty []byte
// 	var ret []byte
// 	return ret
// }

// func fileExists(filename string) bool {
// 	info, err := os.Stat(filename)
// 	if os.IsNotExist(err) {
// 		return false
// 	}
// 	return !info.IsDir()
// }

func errorHandle(err error, str string, ex bool) {
	if err != nil {
		fmt.Println(str, "error:", err)
		if ex {
			os.Exit(1)
		}
	}
}
