# Notion-cli-tool
A CLI tool to add items in a simple budget template of Notion

## Prerequisites

- [Go](https://golang.org/dl/) installed on your system.
- A [Notion](https://www.notion.so/) Account

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/jiteshchawla1511/budget-notion-cli-app.git
   ```

2. Install the necessary dependencies:
   
   ```bash
   go mod download
   ```
3. You need to have a Notion Secret token and Database id
   - To get Notion secret, login in to [Integration](https://www.notion.so/integrations/all) and get the token
   - To get the Database id, follow this [Link](https://developers.notion.com/reference/retrieve-a-database)

4. Make .env file and paste the credentials
   ```
   NOTION_SECRET=
   DATABASE_ID=
   ```
5. Run the main file
   ```
   go run main.go
   ```
6. Add the expense according to you
7. Use command "exit" to exit the tool





