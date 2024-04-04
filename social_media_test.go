package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/text"
	"github.com/stretchr/testify/assert"
)

const (
	class0 = "Class should be 0"
	file   = "/tmp/.goml/NaiveBayes.json"
)

func TestPersistPerceptronShouldPass1(t *testing.T) {
	Init()
	// create the channel of data and errors
	stream := make(chan base.TextDatapoint, 100)
	errors := make(chan error)

	model := text.NewNaiveBayes(stream, 3, base.OnlyWordsAndNumbers)

	go model.OnlineLearn(errors)

	stream <- base.TextDatapoint{
		X: "I love the city",
		Y: 0,
	}

	stream <- base.TextDatapoint{
		X: "I hate Los Angeles",
		Y: 1,
	}

	stream <- base.TextDatapoint{
		X: "My mother is not a nice lady",
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

	// now you can predict like normal
	class := model.Predict("My mother is normally in Los Angeles") // 0
	assert.EqualValues(t, 1, class, class0)

	// now persist to file
	err := model.PersistToFile(file)
	assert.Nil(t, err, "Persistance error should be nil")

	// reset model

	model = text.NewNaiveBayes(stream, 3, base.OnlyWordsAndNumbers)

	class = model.Predict("My mother is leaving Los Angeles") // 0
	assert.EqualValues(t, 0, class, class0)

	// restore from file
	err = model.RestoreFromFile(file)
	assert.Nil(t, err, "Persistance error should be nil")

	// now you can predict like normal
	class = model.Predict("My mother works in Los Angeles") // 0
	assert.EqualValues(t, 1, class, class0)

	// reset again

	model = text.NewNaiveBayes(stream, 3, base.OnlyWordsAndNumbers)

	class = model.Predict("My mother is in Los Angeles") // 0
	assert.EqualValues(t, 0, class, class0)

	// restore file straight from bytes now
	bytes, err := os.ReadFile(file)
	assert.Nil(t, err, "Read file error should be nil")

	assert.Nil(t, model.Restore(bytes), "Model restore error should be nil")

	class = model.Predict("My mother is in Los Angeles") // 0
	assert.EqualValues(t, 1, class, class0)
}

func Init() {
	// create the /tmp/.goml/ dir for persistance testing
	// if it doesn't already exist!
	err := os.MkdirAll("./tmp/.goml", os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("You should be able to create the directory for goml model persistance testing.\n\tError returned: %v\n", err.Error()))
	}
}
