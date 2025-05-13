package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"
)

func main() {
	filePath := "P:\\Projects\\go_github\\releases2.json"
	fl, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var jData []map[string]any
	jsonErr := json.Unmarshal(fl, &jData)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		os.Exit(1)
	}
	sort.Slice(jData, func(i, j int) bool {
		iTime, err := time.Parse(time.RFC3339, jData[i]["published_at"].(string))
		if err != nil {
			os.Exit(1)
		}
		jTime, err := time.Parse(time.RFC3339, jData[j]["published_at"].(string))
		if err != nil {
			os.Exit(1)
		}
		return iTime.Unix() > jTime.Unix()
	})
	for _, t := range jData {
		fmt.Println(t["published_at"])
	}
}
