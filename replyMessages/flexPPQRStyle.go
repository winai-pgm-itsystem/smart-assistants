package replymessages

import (
	"fmt"
)

func getFlexPPQRStyle(replyToken, imageUrl, amount, bankName string) string {

	bodyJson := fmt.Sprintf(`
	{
    "to": "%s",
    "messages": [
        {
            "type": "flex",
            "altText": "create ppqr successfuly",
            "contents": {
							"type": "bubble",
							"body": {
								"type": "box",
								"layout": "vertical",
								"contents": [
								{
									"type": "image",
									"url": "%s",
									"size": "3xl",
									"aspectMode": "cover"
								}
								]
							},
							"footer": {
								"type": "box",
								"layout": "vertical",
								"contents": [
								{
									"type": "text",
									"text": "ยอดชำระ: %s",
									"size": "xl",
									"align": "center"
								},
								{
									"type": "text",
									"text": "%s",
									"align": "center"
								},
								{
									"type": "text",
									"text": "*ส่งหลักฐานการชำระเงินด้วยนะครับ",
									"color": "#b80d21",
									"align": "center"
							}
						]
					}
				}
		}	
    ]
}
	`, replyToken, imageUrl, amount, bankName)

	return bodyJson
}
