package handler

import (
	"fmt"
	"os"
	"smart-assistants/databases"
	"smart-assistants/entities"
	replymessages "smart-assistants/replyMessages"
	"smart-assistants/utils"
	"strconv"
)

func handleMessageEvent(webHookEvent *entities.WebHookEvent) {
	message := webHookEvent.Events[0].Message.Text
	replyToken := webHookEvent.Events[0].ReplyToken
	usserId := webHookEvent.Events[0].Source.UserID

	isMacherExpenseMode, expnenses := utils.MacherExpenseMode(message)

	if usserId != os.Getenv("USER_ID") {
		replymessages.TextMessage(replyToken, "user id is invalid.")
		fmt.Println(" user id is invalid.")

	} else if message == "test" {
		fmt.Println("ok message to new feature")

	} else if isMacherExpenseMode {
		category := entities.ExpenseCategories[expnenses[3]]
		remark := expnenses[4]
		amount, err := strconv.ParseFloat(expnenses[1], 32)
		if err != nil {
			replymessages.TextMessage(replyToken, fmt.Sprintf("amount is a  %s", err))
			return
		}
		_, recordErr := databases.AddExpenseHandler(category, float32(amount), remark)
		if recordErr != nil {
			replymessages.TextMessage(replyToken, recordErr.Error())
		}

		fmt.Printf("macher expense amount: %v , category: %s , remerk: %s", amount, category, remark)

	} else {
		replymessages.TextMessage(replyToken, fmt.Sprintf("message %s not macher.", message))
		fmt.Println(" message not macher")
	}

}
