package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Netinfo struct {
	peersCount float64
}

func queryNetinfo(apiUrl string) Netinfo {
	reqUrl := fmt.Sprintf("%s/%s", apiUrl, "net_info")

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

	result := ApiNetinfoResponse{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}

	peersCount, err := strconv.ParseFloat(result.Result.NPeers, 64)
	if err != nil {
		log.Fatal(err)
	}

	return Netinfo{
		peersCount: peersCount,
	}
}
