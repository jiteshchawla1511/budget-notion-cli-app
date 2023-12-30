package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type ExpensePayload struct {
	Parent struct {
		DatabaseID string `json:"database_id"`
	} `json:"parent"`
	Properties struct {
		Expense struct {
			Title []struct {
				Text struct {
					Content string `json:"content"`
				} `json:"text"`
			} `json:"title"`
		} `json:"Expense"`
		Amount struct {
			Number float64 `json:"number"`
		} `json:"Amount"`
		Date struct {
			Date struct {
				Start string      `json:"start"`
				End   interface{} `json:"end"`
			} `json:"date"`
		} `json:"Date"`
		Category struct {
			Select struct {
				Name string `json:"name"`
			} `json:"select"`
		} `json:"Category"`
	} `json:"properties"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the dotenv file")
	}
	NOTION_SECRET := os.Getenv("NOTION_SECRET")
	DATABASE_ID := os.Getenv("DATABASE_ID")

	client := http.Client{}

	Url := "https://api.notion.com/v1/pages/"
	expenseData := ExpensePayload{
		Parent: struct {
			DatabaseID string `json:"database_id"`
		}{DatabaseID: DATABASE_ID},
		Properties: struct {
			Expense struct {
				Title []struct {
					Text struct {
						Content string `json:"content"`
					} `json:"text"`
				} `json:"title"`
			} `json:"Expense"`
			Amount struct {
				Number float64 `json:"number"`
			} `json:"Amount"`
			Date struct {
				Date struct {
					Start string      `json:"start"`
					End   interface{} `json:"end"`
				} `json:"date"`
			} `json:"Date"`
			Category struct {
				Select struct {
					Name string `json:"name"`
				} `json:"select"`
			} `json:"Category"`
		}{
			Expense: struct {
				Title []struct {
					Text struct {
						Content string `json:"content"`
					} `json:"text"`
				} `json:"title"`
			}{
				Title: []struct {
					Text struct {
						Content string `json:"content"`
					} `json:"text"`
				}{
					{
						Text: struct {
							Content string `json:"content"`
						}{Content: "Expense Name"},
					},
				},
			},
			Amount: struct {
				Number float64 `json:"number"`
			}{Number: 110},
			Date: struct {
				Date struct {
					Start string      `json:"start"`
					End   interface{} `json:"end"`
				} `json:"date"`
			}{Date: struct {
				Start string      `json:"start"`
				End   interface{} `json:"end"`
			}{Start: "2023-12-30", End: nil}},
			Category: struct {
				Select struct {
					Name string `json:"name"`
				} `json:"select"`
			}{Select: struct {
				Name string `json:"name"`
			}{Name: "Food"}},
		},
	}
	jsonData, err := json.Marshal(expenseData)

	jsonReader := strings.NewReader(string(jsonData))
	req, err := http.NewRequest("POST", Url, jsonReader)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+NOTION_SECRET)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}
