package main

import (
	"encoding/csv"
	"math/rand"
	"os"

	"github.com/cdipaolo/goml/cluster"
)


const (
	fileModelMean = "./models/KMeans.json"
	fileResultsMean = "./results/KMeansResults.csv"
	fileErr = "file save error"
)

func Values() {

	gaussian := [][]float64{}
	for i := 0; i < 40; i++ {
		x := rand.NormFloat64() + 4
		y := rand.NormFloat64()*0.25 + 5
		gaussian = append(gaussian, []float64{x, y})
	}
	for i := 0; i < 66; i++ {
		x := rand.NormFloat64()
		y := rand.NormFloat64() + 10
		gaussian = append(gaussian, []float64{x, y})
	}
	for i := 0; i < 100; i++ {
		x := rand.NormFloat64()*3 - 10
		y := rand.NormFloat64()*0.25 - 7
		gaussian = append(gaussian, []float64{x, y})
	}
	for i := 0; i < 23; i++ {
		x := rand.NormFloat64() * 2
		y := rand.NormFloat64() - 1.25
		gaussian = append(gaussian, []float64{x, y})
	}

	model := cluster.NewKMeans(4, 15, gaussian)

	if model.Learn() != nil {
		panic("Oh NO!!! There was an error learning!!")
	}

	// now you can predict like normal!
	// guess, err := model.Predict([]float64{-3, 6})
	// fmt.Println("Prediction")
	// fmt.Println(guess)
	// if err != nil {
	// 	panic("prediction error")
	// }

	// or if you want to get the clustering
	// results from the data
	// fmt.Println("Cluster results")
	// results := model.Guesses()
	// fmt.Println(results)

	// you can also concat that with the
	// training set and save it to a file
	// (if you wanted to plot it or something)
	err := model.SaveClusteredData(fileResultsMean)
	if err != nil {
		panic(fileErr)
	}

	// Write the CSV data
	file2, err := os.OpenFile(fileResultsMean, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	writer := csv.NewWriter(file2)
	defer writer.Flush()
	// this defines the header value and data values for the new csv file
	headers := []string{"x", "y", "class"}
	writer.Write(headers)

	// you can also persist the model to a
	// file
	err = model.PersistToFile(fileModelMean)
	if err != nil {
		panic(fileErr)
	}

	// and also restore from file (at a
	// later time if you want)
	err = model.RestoreFromFile(fileModelMean)
	if err != nil {
		panic(fileErr)
	}
}
