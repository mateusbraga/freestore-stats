package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mateusbraga/gostat"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("You need to pass the file to summarize")
	}

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalln("Error to read file:", err)
		}

		scanner := bufio.NewScanner(file)

		var data []int64
		for scanner.Scan() {
			val, err := strconv.ParseInt(scanner.Text(), 10, 0)
			if err != nil {
				log.Fatalln("error parsing int:", err)
			}
			data = append(data, val)
		}
		if err := scanner.Err(); err != nil {
			log.Fatalln("error:", err)
		}

		filteredData := gostat.TakeExtremes(data)
		mean := gostat.Mean(filteredData)
		sd := gostat.StandardDeviation(filteredData)

		fmt.Printf("%v: %v %v\n", filename, mean, sd)
	}
}
