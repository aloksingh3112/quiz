package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("This is the test and you have given 30 seconds to ans all ques")
	fmt.Println("Type Y/y to start the test and N/n to exit the test")
	var input string
	fmt.Scanln(&input)

	correctData := make(map[string]string)
	incorrect := make(map[string]string)
	if input == "Y" || input == "y" {
		records := readCsv("./problem.csv")
		for _, record := range records {
			fmt.Println(record[0])
			var input string
			fmt.Scanln(&input)
			if input == record[1] {
				correctData[record[0]] = input
			} else {
				incorrect[record[0]] = input + "," + record[1]
			}
		}
		showData(records, correctData, incorrect)

	} else {
		os.Exit(1)
	}

}

func readCsv(file string) [][]string {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err, "file not found")
		os.Exit(1)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	data, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("unable to parser csv file", err)
		os.Exit(1)
	}
	return data
}

func showData(records [][]string, correct map[string]string, incorrect map[string]string) {
	fmt.Println("Total Question :", len(records))
	fmt.Println("correct Question :", len(correct))
	fmt.Println("incorrect Question :", len(incorrect))

	fmt.Println("wrong answers :-")

	for key, ans := range incorrect {
		a := strings.Split(ans, ",")
		fmt.Println("Question :", key, "Answer :", a[1], "Your Answer :", a[0])
	}

}
