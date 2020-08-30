package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	forceFrom = flag.String("f", "", "from key for force specifing")
	forceTo   = flag.String("t", "", "to keyfor force specifing")
	path      = flag.String("p", "", "path to input file")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println("you can specify input from STDIN or -p input file path.")
	}
	flag.Parse()
	var input string
	if *path != "" {
		b, err := ioutil.ReadFile(*path)
		if err != nil {
			log.Fatalf("Read from file failed: %v", err)
		}
		input = string(b)
	} else {
		b, err := readStdin()
		if err != nil {
			log.Fatalf("Read from stdin failed: %v", err)
		}
		input = string(b)
	}
	if input == "" {
		flag.Usage()
		return
	}

	kf := NewKeyFinder(*forceFrom, *forceTo)
	// json is subset of yaml. so first try to parse as json
	appended, jerr := Append(kf, &JSONMarshaler{}, input)
	if jerr != nil {
		ry, yerr := Append(kf, &YAMLMarshaler{}, input)
		if yerr != nil {
			rc, cerr := AppendCSV(kf, input)
			if cerr != nil {
				log.Printf("json unmarshal error: %v", jerr)
				log.Printf("yaml unmarshal error: %v", yerr)
				log.Printf("csv reader error: %v", cerr)
				log.Fatal("cant parse input by json, yaml, csv parser")
			}
			appended = rc
		} else {
			appended = ry
		}
	}
	fmt.Print(appended)

}

func readStdin() ([]byte, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return []byte{}, err
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return []byte{}, err
		}
		return b, nil
	} else {
		return []byte{}, nil
	}
}
