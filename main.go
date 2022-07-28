package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Params struct {
	APIKey      string `json:"apiKey"`
	N           int    `json:"n"`
	Min         int    `json:"min"`
	Max         int    `json:"max"`
	Replacement bool   `json:"replacement"`
	Base        int    `json:"base"`
}

type Payload struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
	ID      int    `json:"id"`
}

type RandomOrgResponse struct {
	Result struct {
		Random struct {
			Data []int `json:"data"`
		} `json:"random"`
	} `json:"result"`
}

type RandomStandard struct {
	Stddev float64 `json:"stddev"`
	Data   []int   `json:"data"`
}

var client = &http.Client{Timeout: 10 * time.Second}

func requestRandomNumbers(payload Payload) ([]int, error) {
	payloadBytes, _ := json.Marshal(payload)

	body := bytes.NewReader(payloadBytes)

	req, _ := http.NewRequest("POST", "https://api.random.org/json-rpc/4/invoke", body)

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var response RandomOrgResponse
	unmarshalErr := json.Unmarshal(respBody, &response)

	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
		return nil, unmarshalErr
	}

	randomNumbers := response.Result.Random.Data

	return randomNumbers, nil
}

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	r := gin.Default()
	r.GET("/random/mean", randomMeanHandler)
	r.Run(":" + PORT)
}

func randomMeanHandler(c *gin.Context) {
	concurrentRequestSize := c.DefaultQuery("requests", "1")
	n := c.DefaultQuery("length", "1")

	N, _ := strconv.Atoi(n)
	convertedConcurrentRequestSize, _ := strconv.Atoi(concurrentRequestSize)

	var standardResponse []RandomStandard
	var sum []int

	var wg sync.WaitGroup
	wg.Add(convertedConcurrentRequestSize)

	for i := 1; i <= convertedConcurrentRequestSize; i++ {
		data := Payload{
			Jsonrpc: "2.0",
			Method:  "generateIntegers",
			Params: Params{
				APIKey:      os.Getenv("API_KEY"),
				N:           N,
				Min:         1,
				Max:         10,
				Replacement: true,
				Base:        10,
			},
			ID: i,
		}

		go func() {
			randomNumbers, _ := requestRandomNumbers(data)

			if randomNumbers != nil {
				std := calculateStandardDeviation(randomNumbers)

				sum = append(sum, randomNumbers...)

				standardResponse = append(standardResponse, RandomStandard{Stddev: std, Data: randomNumbers})

			}

			wg.Done()
		}()
	}
	wg.Wait()

	std := calculateStandardDeviation(sum)
	standardResponse = append(standardResponse, RandomStandard{Stddev: std, Data: sum})

	c.JSON(http.StatusOK, standardResponse)
}
