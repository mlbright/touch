package main

import (
	"flag"
	"log"
	"os"
	"time"
)

func touch(path string) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if _, err := os.Create(path); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	} else {
		now := time.Now()
		if err := os.Chtimes(path, now, now); err != nil {
            log.Fatal(err)
        }
	}
}

func main() {
	flag.Parse()
	for _, path := range flag.Args() {
		touch(path)
	}
}
