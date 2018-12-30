package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tokuhirom/go-hsperfdata/hsperfdata"
)

func dumpStat(result *hsperfdata.Result) {
	proc_name := result.GetProcName()
	for key, value := range result.GetMap() {
		_, is_string := value.(string)
		if !is_string {
			fmt.Printf("%s %v proc_name=\"%v\"\n", key, value, proc_name)
		}
	}
}

func main() {
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		fmt.Printf("%v\n", hsperfdata.GetVersion())
		return
	}

	for _, dir := range os.Args {
		repo, err := hsperfdata.NewDir(dir)
		if err != nil {
			log.Fatal("new", err)
		}
		files, err := repo.GetFiles()
		if err == nil {
			for _, f := range files {
				result, err := f.Read()
				if err == nil {
					dumpStat(result)
				}
			}
		}
	}
}
