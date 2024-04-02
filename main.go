package main

import (
	"fmt"
	"math/rand"

	"github.com/cdipaolo/goml/cluster"
)

func main() {
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
	guess, err := model.Predict([]float64{-3, 6})
	fmt.Println("Prediction")
	fmt.Println(guess)
	if err != nil {
		panic("prediction error")
	}

	// or if you want to get the clustering
	// results from the data
	fmt.Println("Cluster results")
	results := model.Guesses()
	fmt.Println(results)

	// you can also concat that with the
	// training set and save it to a file
	// (if you wanted to plot it or something)
	err = model.SaveClusteredData("./KMeansResults.csv")
	if err != nil {
		panic("file save error")
	}

	// you can also persist the model to a
	// file
	err = model.PersistToFile("./KMeans.json")
	if err != nil {
		panic("file save error")
	}

	// and also restore from file (at a
	// later time if you want)
	err = model.RestoreFromFile("./KMeans.json")
	if err != nil {
		panic("file save error")
	}

}
