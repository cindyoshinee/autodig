package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/cindyoshinee/autodig/dep"
)

var (
	scanDir    string
	outputFile string
	tag        string
)

func init() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	flag.StringVar(&scanDir, "scans", fmt.Sprintf("%s/app", dir), "source code scan dirs, split with ','")
	flag.StringVar(&outputFile, "output", fmt.Sprintf("%s/app/entrypoint/autodig.go", dir), "output file path")
	flag.StringVar(&tag, "tag", "", "tag, only support one")
}

func main() {
	fmt.Println("=========autodig start==========")
	flag.Parse()
	scanDirFlag := flag.Lookup("scan")
	outputFileFlag := flag.Lookup("output")
	tagFlag := flag.Lookup("tag")
	if scanDirFlag != nil {
		scanDir = scanDirFlag.Value.String()
	}
	if outputFileFlag != nil {
		outputFile = outputFileFlag.Value.String()
	}
	if tagFlag != nil {
		tag = tagFlag.Value.String()
	}
	fmt.Println("dir", os.Args[0])
	fmt.Println("scanDir", scanDir)
	fmt.Println("outputFile", outputFile)
	err := dep.NewAutodig(strings.Split(scanDir, ","), outputFile, tag).GenDigFile()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("=========autodig success!!==========")
}
