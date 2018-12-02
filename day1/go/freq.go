package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var (
	inputFilePath = "../input"
)

func main() {
	ffV1, err := getFinalFreqV1(inputFilePath)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("The first part's final frequency is %v\n", ffV1)

	ffV2, err := getFinalFreqV2(inputFilePath)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("The second part's final frequency is %v\n", ffV2)
}

func getFinalFreqV1(filePath string) (int, error) {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return 0, err
	}

	freq := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		freq = updateFreq(freq, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return freq, nil
}

func getFinalFreqV2(filePath string) (int, error) {
	freqMap := make(map[int]bool)
	freqMap[0] = true
	freq := 0

	// Scary infinite while loop!
	for {
		f, err := os.Open(filePath)
		if err != nil {
			return 0, err
		}
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			freq = updateFreq(freq, scanner.Text())
			if freqMap[freq] {
				return freq, nil
			}
			freqMap[freq] = true
		}
		if err := scanner.Err(); err != nil {
			return 0, err
		}
	}
}

func updateFreq(freq int, input string) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		log.Panic(err)
	}
	return freq + val
}
