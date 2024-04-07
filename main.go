package main

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

type HitsSource struct {
	Title string `json:"title"`
}

type HitsInHitsBody struct {
	Index  string     `json:"_index"`
	ID     string     `json:"_id"`
	Score  string     `json:"_score"`
	Source HitsSource `json:"_source"`
}

type HitsTotal struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type HitsInHits struct {
	Total    HitsTotal        `json:"total"`
	MaxScore float64          `json:"max_score"`
	Hits     []HitsInHitsBody `json:"hits"`
}

type EsRes struct {
	Took    int        `json:"took"`
	TimeOut bool       `json:"time_out"`
	Hits    HitsInHits `json:"hits"`
}

func main() {
	fmt.Println("elasticsearch - go implementation")
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	if err != nil {
		panic(err)
	}

	// fmt.Println(esClient.Info())

	// CREATE index
	// res, err := esClient.Index(
	// 	"go-index1",                             // Index name
	// 	strings.NewReader(`{"title" : "Test"}`), // Document body
	// 	esClient.Index.WithDocumentID("1"),      // Document ID
	// 	esClient.Index.WithRefresh("true"),      // Refresh
	// )
	// if err != nil {
	// 	log.Fatalf("ERROR: %s", err)
	// }
	// defer res.Body.Close()

	// GET
	// res, err := esClient.Get("go-index1", "1")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res.String())
	// defer res.Body.Close()

	// QUERY
	res, _ := esClient.Search(esClient.Search.WithIndex("go-index1"))
	fmt.Println(res)
	resBody := EsRes{}
	json.NewDecoder(res.Body).Decode(&resBody)
	fmt.Println(resBody.Took)
	fmt.Println(resBody.TimeOut)
	fmt.Println(resBody.Hits.Hits[0].Source.Title)
}
