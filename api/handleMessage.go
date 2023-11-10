package handler

import (
	"smart-assistants/entities"
	replymessages "smart-assistants/replyMessages"
)

func handleMessageEvent(webHookEvent *entities.WebHookEvent) {
	message := webHookEvent.Events[0].Message.Text
	replyToken := webHookEvent.Events[0].ReplyToken

	replymessages.TextMessage(replyToken, message)

}
