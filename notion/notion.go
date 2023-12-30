package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ExpensePayload struct {
	Parent     Parent     `json:"parent"`
	Properties Properties `json:"properties"`
}

type Parent struct {
	DatabaseID string `json:"database_id"`
}

type Properties struct {
	Expense  Expense  `json:"Expense"`
	Amount   Amount   `json:"Amount"`
	Date     Date     `json:"Date"`
	Category Category `json:"Category"`
}

type Expense struct {
	Title []Title `json:"title"`
}

type Title struct {
	Text Text `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}

type Amount struct {
	Number float64 `json:"number"`
}

type Date struct {
	Date DateValue `json:"date"`
}

type DateValue struct {
	Start string `json:"start"`
}

type Category struct {
	Select Select `json:"select"`
}

type Select struct {
	Name string `json:"name"`
}

func AddExpense(name string, amount float64, date, category string) error {
	DATABASE_ID := os.Getenv("DATABASE_ID")
	NOTION_SECRET := os.Getenv("NOTION_SECRET")

	expenseData := ExpensePayload{

		Parent: Parent{DatabaseID: DATABASE_ID},
		Properties: Properties{
			Expense: Expense{
				Title: []Title{
					{Text: Text{Content: name}},
				},
			},
			Amount: Amount{Number: amount},
			Date: Date{
				Date: DateValue{Start: date},
			},
			Category: Category{
				Select: Select{Name: category},
			},
		},
	}

	jsonData, err := json.Marshal(expenseData)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	jsonReader := bytes.NewReader(jsonData)

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api.notion.com/v1/pages", jsonReader)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+NOTION_SECRET)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", "2021-05-13")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return nil
}
