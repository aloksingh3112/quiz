package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("This is the test and you have given 30 seconds to ans all ques")
	fmt.Println("Type Y/y to start the test and N/n to exit the test")
	var input string
	fmt.Scanln(&input)

	correctData := make(map[string]string)
	incorrect := make(map[string]string)
	if input == "Y" || input == "y" {
		timer1 := time.NewTimer(10 * time.Second)
		records := readCsv("./problem.csv")
		c := make(chan string)
	loop:
		for _, record := range records {
			fmt.Println(record[0])
			go takeInput(c)
			select {
			case <-timer1.C:
				showData(records, correctData, incorrect)
				break loop
			case input := <-c:
				if input == record[1] {
					correctData[record[0]] = input
				} else {
					incorrect[record[0]] = input + "," + record[1]
				}
			}

		}

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

func takeInput(c chan string) {
	var input string
	fmt.Scanln(&input)
	c <- input

}

// func takeQuiz(records [][]string, correctData map[string]string, incorrect map[string]string, c chan string) {
// 	for _, record := range records {
// 		fmt.Println(record[0])
// 		go takeInput(c)
// 		input := <-c
// 		if input == record[1] {
// 			correctData[record[0]] = input
// 		} else {
// 			incorrect[record[0]] = input + "," + record[1]
// 		}
// 	}
// 	showData(records, correctData, incorrect)
// }

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
