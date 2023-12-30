package cli

import (
	"fmt"

	"github.com/jiteshchawla1511/budget-notion-cli-app/notion"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "notion-cli",
}
var addExpenseCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense to Notion",
	Run:   addExpense,
}

func init() {
	RootCmd.AddCommand(addExpenseCmd)
}

func addExpense(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	amount, _ := cmd.Flags().GetFloat64("amount")
	date, _ := cmd.Flags().GetString("date")
	category, _ := cmd.Flags().GetString("category")

	err := notion.AddExpense(name, amount, date, category)
	if err != nil {
		fmt.Println("Error adding expense:", err)
		return
	}

	fmt.Println("Expense added successfully to Notion!")
}

// AddInteractive is for interactive expense addition
func AddInteractive(cmd *cobra.Command, args []string) {
	fmt.Println("Interactive mode for adding expenses to Notion.")
	fmt.Println("Enter 'exit' to quit.")

	for {
		fmt.Print("Expense Name: ")
		var name string
		fmt.Scanln(&name)
		if name == "exit" {
			break
		}

		fmt.Print("Amount: ")
		var amount float64
		fmt.Scanln(&amount)

		fmt.Print("Date (YYYY-MM-DD): ")
		var date string
		fmt.Scanln(&date)

		fmt.Print("Category: ")
		var category string
		fmt.Scanln(&category)

		err := notion.AddExpense(name, amount, date, category)
		if err != nil {
			fmt.Println("Error adding expense:", err)
			return
		}

		fmt.Println("Expense added successfully to Notion!")
	}
}
