package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type arrayFlag []string

func (a *arrayFlag) String() string {
	return fmt.Sprintf("%v", *a)
}

func (a *arrayFlag) Set(value string) error {
	*a = strings.Split(value, ",")
	return nil
}

var af arrayFlag

func main() {
	// flags
	urlPtr := flag.String("url", "", "Url to check.")
	flag.Var(&af, "urls", "comma separated string of urls")
	timePtr := flag.Duration("time", 60*time.Second, "maximum time to run check.")

	flag.Parse()

	var err error

	switch {
	case *urlPtr == "" && len(af) == 0:
		flag.PrintDefaults()
		os.Exit(1)
	case *urlPtr != "" && len(af) > 0:
		log.Fatal("Either specify -url or --urls exclusively to run this command.")
	case len(af) > 0:
		err = fetchURL(*timePtr, af)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	default:
		err = fetchURL(*timePtr, strings.Fields(*urlPtr))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}
}
