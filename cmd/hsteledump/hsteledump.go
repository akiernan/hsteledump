package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/tokuhirom/go-hsperfdata/hsperfdata"
)

func dumpStat(result *hsperfdata.Result) {
	proc_name := result.GetProcName()
	for key, value := range result.GetMap() {
		if str, ok := value.(string); ok {
			_, err := strconv.ParseFloat(str, 64)
			if err == nil {
				fmt.Printf("%s %v proc_name=\"%s\"\n", key, value, proc_name)
			}
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

	repo, err := hsperfdata.NewUser("root")
	if err != nil {
		log.Fatal("new", err)
	}
	files, err := repo.GetFiles()
	if err != nil {
		log.Fatal("repo", err)
	}

	for _, f := range files {
		result, err := f.Read()
		if err == nil {
			dumpStat(result)
		} else {
			fmt.Printf("%s unknown\n", f.GetPid())
		}
	}
}
