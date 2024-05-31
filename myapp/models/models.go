package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Article struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Img      string `json:"img"`
	ImgAlt   string `json:"img_alt"`
	SubTitle string `json:"sub_title"`
	Text     string `json:"text"`
}

func LoadArticles(filename string) ([]Article, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	var articles []Article
	err = json.Unmarshal(bytes, &articles)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	return articles, nil
}
