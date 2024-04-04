package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/text"
)

const (
	fileModel = "./models/NaiveBayes.json"
	plotFile   = "./social.csv"
)

func Social() {
	// create the channel of data and errors
	stream := make(chan base.TextDatapoint, 100)
	errors := make(chan error)

	model := text.NewNaiveBayes(stream, 3, base.OnlyWordsAndNumbers)

	go model.OnlineLearn(errors)

	stream <- base.TextDatapoint{
		X: "what a perfect combination for adding to my plates",
		Y: 2,
	}

	stream <- base.TextDatapoint{
		X: "sushi is my favourite food, this idea is great",
		Y: 2,
	}

	stream <- base.TextDatapoint{
		X: "I love japanese food. This a great product",
		Y: 2,
	}

	stream <- base.TextDatapoint{
		X: "idots like you are all around",
		Y: 0,
	}

	stream <- base.TextDatapoint{
		X: "this is stupid and your product sucks",
		Y: 0,
	}

	stream <- base.TextDatapoint{
		X: "why do you boder to show this",
		Y: 0,
	}

	stream <- base.TextDatapoint{
		X: "I will probably try it",
		Y: 1,
	}

	stream <- base.TextDatapoint{
		X: "I think it is the first time that I see this",
		Y: 1,
	}

	stream <- base.TextDatapoint{
		X: "I could imagine that this has interest to some people",
		Y: 1,
	}

	close(stream)

	for {
		err, more := <-errors
		if more {
			fmt.Printf("Error passed: %v", err)
		} else {
			// training is done!
			break
		}
	}

	// now persist to file
	err := model.PersistToFile(fileModel)
	if err != nil {
		fmt.Println("error guardando datos")
	}

	var comments = []string{"My mother is in Japan and she never saw this", "I probably like to test it", "what a bad idea you had", "I don't like this", "I eat sushi every week", "it looks interesting", "nice job, I could use this product for my restaurant", "tomorrow I will buy one", "where can I buy it?", "what a nice idea!", "is it possible to try it?"}

	// This a test for a prediction
	class := model.Predict("My mother is in Delhi a city in India, gets very hot")
	fmt.Println("the prediction is: ", class)


	file, err := os.Create(plotFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalClass0 := 0
	totalClass1 := 0
	totalClass2 := 0

	for _, comment := range comments {
		class = model.Predict(comment)

		switch class {
		case 0:
			totalClass0 = totalClass0 + 1
		case 1:
			totalClass1 = totalClass1 + 1
		case 2:
			totalClass2 = totalClass2 + 1
		default:
			fmt.Println("no class provided")
		}

	}

	// Include the headers
	file2, err := os.OpenFile(plotFile, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	writer := csv.NewWriter(file2)
	defer writer.Flush()
	// this defines the header value and data values for the new csv file
	headers := []string{"class", "value"}
	// Total values per class
	value1 := []string{"0", strconv.Itoa(totalClass0)}
	value2 := []string{"1", strconv.Itoa(totalClass1)}
	value3 := []string{"2", strconv.Itoa(totalClass2)}
	records := [][]string{headers, value1, value2, value3}
	writer.WriteAll(records)

}
