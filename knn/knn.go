package main

import (
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	//"github.com/sjwhitworth/golearn/knn"
	"../../Documents/ML2/golearn/knn"
	"fmt"
	"time"
)

func main() {
	rawData, err := base.ParseCSVToInstances("loans_data2.csv", true)
	if err != nil {
		panic(err)
	}

	////////////////////////////////////////////////////////////////////
	// euclidean
	///////////////////////////////////////////////////////////////////
        fmt.Println("-> euclidean")

        stamp := time.Now()

	//Initialises a new KNN classifier
	cls := knn.NewKnnClassifier("euclidean", "kdtree", 40)
	cls.Weighted = true

	//Do a training-test split
	trainData, testData := base.InstancesTrainTestSplit(rawData, 11.0/12.0)
	cls.Fit(trainData)

	//Calculates the Euclidean distance and returns the most popular label
	predictions, err := cls.Predict(testData)
	if err != nil {
		panic(err)
	}
	//fmt.Println(predictions)

	// Prints precision/recall metrics
	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(confusionMat))
        fmt.Println("time: ", time.Since(stamp))
        
        ///////////////////////////////////////////////////////////////////////
        //manhattan
        //////////////////////////////////////////////////////////////////////
        fmt.Println("-> manhattan")


        stamp = time.Now()

	//Initialises a new KNN classifier
	cls = knn.NewKnnClassifier("manhattan", "kdtree", 40)
	cls.Weighted = true

	//Do a training-test split
	trainData, testData = base.InstancesTrainTestSplit(rawData, 11.0/12.0)
	cls.Fit(trainData)

	//Calculates the Euclidean distance and returns the most popular label
	predictions, err = cls.Predict(testData)
	if err != nil {
		panic(err)
	}
	//fmt.Println(predictions)

	// Prints precision/recall metrics
	confusionMat, err = evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(confusionMat))
        fmt.Println("time: ", time.Since(stamp))
        

        ////////////////////////////////////////////////////////////////////////
        // cosine
        ///////////////////////////////////////////////////////////////////////
        fmt.Println("-> cosine")
        
        stamp = time.Now()

	//Initialises a new KNN classifier
	cls = knn.NewKnnClassifier("cosine", "kdtree", 40)
	cls.Weighted = true

	//Do a training-test split
	trainData, testData = base.InstancesTrainTestSplit(rawData, 11.0/12.0)
	cls.Fit(trainData)

	//Calculates the Euclidean distance and returns the most popular label
	predictions, err = cls.Predict(testData)
	if err != nil {
		panic(err)
	}
	//fmt.Println(predictions)

	// Prints precision/recall metrics
	confusionMat, err = evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(confusionMat))
        fmt.Println("time: ", time.Since(stamp))
}
