package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

type RSS struct {
	Channel struct {
		Items []struct {
			Title       string `xml:"title"`
			Description string `xml:"description"`
		} `xml:"item"`
	} `xml:"channel"`
}

func main() {
	client := &http.Client{Timeout: 15 * time.Second}

	resp, err := client.Get("https://techcrunch.com/feed/")
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic("respond faild and returned: " + resp.Status)
	}

	var feed RSS
	if err := xml.NewDecoder(resp.Body).Decode(&feed); err != nil {
		panic(err)
	}

	for _, item := range feed.Channel.Items {
		fmt.Printf("$ %s\n- %s\n\n", item.Title, item.Description)
	}
}
