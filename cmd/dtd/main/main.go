package main

import (
	"bufio"
	"dt/csv"
	"errors"
	"flag"
	"log"
	"os"
	"strings"
)

var usage = `tools commonds:
			 csv(default): transform log to csv format`

func main() {
	cmd := flag.String("cmd", "csv", usage)
	input := flag.String("input", "", "input log file")
	output := flag.String("output", "", "output csv file")
	fields := flag.String("fields", "", "csv title fields")
	flag.Parse()
	switch *cmd {
	case "csv":
		err := Csv(*input, *output, *fields)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	return
}

func Csv(input, output, fields string) error {
	f, err := os.Open(input)
	if err != nil {
		return err
	}
	defer f.Close()

	if output == "" || fields == "" {
		return errors.New("argv output and fields not set")
	}
	titles := strings.Split(fields, ",")
	csv, err := csv.NewCsv(output, titles)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		columns, err := FormatLogCsv(line, titles)
		if err != nil {
			return err
		}
		err = csv.WriteLines([][]string{columns})
		if err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func FormatLogCsv(line string, fields []string) ([]string, error) {
	strs := strings.Split(line, " ")
	var colums []string
	if len(strs) >= 13 {
		colums = strs[5:14]
	} else {
		return nil, errors.New("length is too low ")
	}
	return colums, nil
}
