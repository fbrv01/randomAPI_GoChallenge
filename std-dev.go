package main

import (
	"math"
)

func calculateStandardDeviation(nums []int) float64 {
	var sum int
	var mean, sd float64

	for _, element := range nums {
		sum += element
	}
	mean = float64(sum) / float64(len(nums))

	for _, element := range nums {
		sd += math.Pow(float64(element)-mean, 2)
	}

	sd = math.Sqrt(sd / float64(len(nums)))

	return sd
}
