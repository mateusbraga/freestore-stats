package main

import (
    "os"
    "strconv"
    "bufio"
    "log"

    "github.com/mateusbraga/gostat"
)

func main() {
    file, err := os.Open("/home/mateus/read-latency-1.txt")
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
    log.Println("size:", len(filteredData), filteredData)
    mean := gostat.Mean(filteredData)
    sd := gostat.StandardDeviation(filteredData)

    log.Println(mean, sd)
}



