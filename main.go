package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	NoIndex = -1
)

var (
	forceFrom      = flag.String("f", "", "from key for force specifing")
	forceTo        = flag.String("t", "", "to key for force specifing")
	path           = flag.String("p", "", "path to input file")
	forceIndexFrom = flag.Int("if", NoIndex, "from csv index for force specifing")
	forceIndexTo   = flag.Int("it", NoIndex, "to csv index for force specifing")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println("you can specify input from STDIN or -p input file path.")
	}
	flag.Parse()
	var input []byte
	if *path != "" {
		b, err := ioutil.ReadFile(*path)
		if err != nil {
			log.Fatalf("Read from file failed: %v", err)
		}
		input = b
	} else {
		b, err := readStdin()
		if err != nil {
			log.Fatalf("Read from stdin failed: %v", err)
		}
		input = b
	}
	sinput := string(input)
	if sinput == "" {
		flag.Usage()
		return
	}
	if (*forceIndexFrom == NoIndex && *forceIndexTo != NoIndex) || (*forceIndexFrom != NoIndex && *forceIndexTo == NoIndex) {
		flag.Usage()
		return
	}

	kf := NewKeyFinder(*forceFrom, *forceTo)
	// json is subset of yaml. so first try to parse as json
	var test interface{}
	var appended string
	var jerr, yerr, aerr error
	if jerr = json.Unmarshal(input, &test); jerr == nil {
		jr, err := Append(kf, &JSONMarshaler{}, sinput)
		if err != nil {
			aerr = fmt.Errorf("cant append json. %v", err)
		}
		appended = jr
	} else if yerr = yaml.Unmarshal(input, &test); yerr == nil {
		yr, err := Append(kf, &YAMLMarshaler{}, sinput)
		if err != nil {
			aerr = fmt.Errorf("cant append yaml. %v", err)
		}
		appended = yr
	}
	if aerr != nil {
		cr, err := AppendCSV(kf, sinput, *forceIndexFrom, *forceIndexTo)
		if err != nil {
			log.Printf("json unmarshal error: %v", jerr)
			log.Printf("yaml unmarshal error: %v", yerr)
			log.Printf("csv reader error: %v", err)
			log.Fatal("cant parse input by json, yaml, csv parser")
		}
		appended = cr
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
