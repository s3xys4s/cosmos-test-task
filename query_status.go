package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Status struct {
	latestBlockHeight float64
	latestBlockTime   float64
	blockDesyncTime   float64
}

func queryStatus(apiUrl string) Status {
	reqUrl := fmt.Sprintf("%s/%s", apiUrl, "status")

	res, err := http.Get(reqUrl)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := ApiStatusResponse{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}

	latestBlockHeight, err := strconv.ParseFloat(result.Result.SyncInfo.LatestBlockHeight, 64)
	if err != nil {
		log.Fatal(err)
	}

	latestBlockTime := result.Result.SyncInfo.LatestBlockTime.Unix()
	BlockDesyncTime := time.Now().Unix() - latestBlockTime

	return Status{
		latestBlockHeight,
		float64(latestBlockTime),
		float64(BlockDesyncTime),
	}
}
