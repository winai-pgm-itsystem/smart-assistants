package main

import (
	"fmt"
	"log"
	"smart-assistants/databases"
	"smart-assistants/utils"

	"github.com/joho/godotenv"
)

func main() {
	// comment this when deploy to vercel
	err := godotenv.Load("prod.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// app := gin.New()
	// handler.SetupRoutes(app)

	// app.Run(":8080")

	date := utils.GetCurrentDate("Asia/Bangkok")

	result, err := databases.GetExpenseHandler(date, "YYYY-MM-DD")

	if err != nil {
		fmt.Printf("get expense error : %s", err)
	}

	fmt.Println(result)

	for i := 0; i < len(result.Records); i++ {

		category := result.Records[i].Fields.Category
		amount := result.Records[i].Fields.Amount
		remark := result.Records[i].Fields.Remark
		fmt.Printf("c: %s , a: %v  , r: %s \n", category, amount, remark)

	}

	// rec, err := databases.AddExpenseHandler("food", 45.32, "test")
	// if err != nil {
	// 	fmt.Printf("add expense error : %s", err)
	// }

	// fmt.Println(rec)

}
