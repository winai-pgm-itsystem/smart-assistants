package handlers

import (
	"smart-assistants/entities"
	replymessages "smart-assistants/replyMessages"
	"strings"
)

func HandleMessageEvent(webHookEvent *entities.WebHookEvent) {
	message := webHookEvent.Events[0].Message.Text
	replyToken := webHookEvent.Events[0].ReplyToken
	// usserId := webHookEvent.Events[0].Source.UserID
	//
	// isMacherExpenseMode, expnenses := utils.MacherExpenseMode(message)
	// //
	// ppqrAmount, macherPPQR := utils.MacherPPQRMode(message)

	// if usserId != os.Getenv("USER_ID") {
	// 	replymessages.SendTextMessage(replyToken, "user id is invalid.")
	// 	fmt.Println(" user id is invalid.")

	// } else if message == "test" {
	// 	fmt.Println("ok message to new feature")

	// } else if isMacherExpenseMode {
	// 	category := entities.ExpenseCategories[expnenses[3]]
	// 	remark := expnenses[4]
	// 	amount, err := strconv.ParseFloat(expnenses[1], 32)
	// 	if err != nil {
	// 		replymessages.SendTextMessage(replyToken, fmt.Sprintf("amount is a  %s", err))
	// 		return
	// 	}
	// 	id, recordErr := databases.AddExpenseHandler(category, float32(amount), remark)
	// 	if recordErr != nil {
	// 		replymessages.SendTextMessage(replyToken, recordErr.Error())
	// 	}
	// 	var currentTotalPerMonth float64
	// 	var currentTotalPerDate float64
	// 	if id != "" {
	// 		month := utils.GetCurrentMonth()
	// 		today := utils.GetCurrentDate()
	// 		monthExpense, err := databases.GetExpenseHandler(month, "YYYY-MM")
	// 		if err != nil {
	// 			replymessages.SendTextMessage(replyToken, err.Error())
	// 		}
	// 		for i := 0; i < len(monthExpense.Records); i++ {
	// 			if today == monthExpense.Records[i].Fields.Date {
	// 				amountPerDate := monthExpense.Records[i].Fields.Amount
	// 				currentTotalPerDate += float64(amountPerDate)
	// 			}
	// 			amountPerMonth := monthExpense.Records[i].Fields.Amount
	// 			currentTotalPerMonth += float64(amountPerMonth)
	// 		}

	// 	}
	// 	limitPerMonth, err := strconv.ParseFloat(os.Getenv("LIMIT_EXPENSE_PER_MONTH"), 64)
	// 	if err != nil {
	// 		replymessages.SendTextMessage(replyToken, fmt.Sprintf("limit per month is a  %s", err))
	// 		return
	// 	}

	// 	replymessages.SendFlexExpenseMessage(utils.GetSymbolCurrentcyFormat(amount), category, utils.GetSymbolCurrentcyFormat(currentTotalPerDate), utils.GetSymbolCurrentcyFormat(limitPerMonth-currentTotalPerMonth), utils.GetToday())

	// } else if macherPPQR {

	// if ppqrAmount > 0 {
	// 	replymessages.SendFlexPPQRMessage(utils.GetSymbolCurrentcyFormat(ppqrAmount), ppqrAmount)

	// } else {
	// 		replymessages.SendTextMessage(replyToken, "ppqr not generated.")

	// 	}

	// }

	// else {
	// 	result, err := utils.AIGenearateText(message)
	// 	if err != nil {
	// 		replymessages.SendTextMessage(replyToken, fmt.Sprintf("message %s not macher.", message))
	// 	}

	// 	replymessages.SendTextMessage(replyToken, fmt.Sprintf(`%s`, result))

	// }
	var replyText string

	if strings.Contains(message, "ขายอะไร") {
		replyText = "- กาแฟ และเครื่องดื่มค่ะ เช่น คาปูชิโน่ ลาเต้ อเมริกาโน่ ชาไทย"
		replymessages.SendTextMessage(replyToken, replyText)
	} else if strings.Contains(message, "สวัสดี") {
		replyText = "สวัสดีค่ะ"
		replymessages.SendTextMessage(replyToken, replyText)
	} else {
		replyText = "ขออภัยค่ะ ฉันไม่เข้าใจคำถาม กรุณาถามคำถามใหม่ค่ะ"
		replymessages.SendTextMessage(replyToken, replyText)
	}

}
