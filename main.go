package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

func setupAlias() {
	home := os.Getenv("HOME")
	if home == "" {
		return
	}

	shell := os.Getenv("SHELL")
	rc := filepath.Join(home, ".bashrc")
	base := filepath.Base(shell)
	if base == "zsh" {
		rc = filepath.Join(home, ".zshrc")
	} else if _, err := os.Stat(filepath.Join(home, ".zshrc")); err == nil && base != "bash" {
		rc = filepath.Join(home, ".zshrc")
	}

	dir, err := os.Getwd()
	if err != nil {
		return
	}
	aliasLine := fmt.Sprintf("alias news='cd %s && go run main.go'\n", dir)

	content, err := os.ReadFile(rc)
	if err != nil {
		_ = os.WriteFile(rc, []byte(aliasLine), 0644)
		return
	}

	if !strings.Contains(string(content), "alias news=") {
		f, err := os.OpenFile(rc, os.O_APPEND|os.O_WRONLY, 0644)
		if err == nil {
			f.WriteString(aliasLine)
			f.Close()
		}
	}
}

func main() {
	setupAlias()

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
