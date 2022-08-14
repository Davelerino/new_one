package main

import (
	"encoding/json"
	"flag"
	"os"
)

func main() {
	fileName := flag.String("file", "gopher.json", "ficher json à lire")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(f)
	_ = d
	var story packs.Story

	if err := d.Decode(&story); err != nil {
		fmt.Println("le ficher est illisible")
	}
	fmt.Printf("\n %+v", story)

}
