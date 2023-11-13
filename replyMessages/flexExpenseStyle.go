package replymessages

import "fmt"

func getFlexExpenseStyle(replyToken, amount, category, currentToday, limit, day string) string {

	bodyJson := fmt.Sprintf(`
	{
    "to": "%s",
    "messages": [
        {
            "type": "flex",
            "altText": "Update expense tracking successfuly",
            "contents": {
                "type": "bubble",
                "header": {
                    "type": "box",
                    "layout": "vertical",
                    "contents": [
                        {
                            "type": "text",
                            "text": "expense tracking",
                            "size": "lg",
                            "margin": "none",
                            "color": "#FFFFFF"
                        }
                    ],
                    "backgroundColor": "#2F4F4F"
                },
                "body": {
                    "type": "box",
                    "layout": "vertical",
                    "contents": [
                        {
                            "type": "text",
                            "text": "%s",
                            "weight": "bold",
                            "size": "xxl"
                        },
                        {
                            "type": "text",
                            "text": "%s",
                            "size": "md"
                        },
                        {
                            "type": "text",
                            "text": "recorded"
                        }
                    ],
                    "margin": "md"
                },
                "footer": {
                    "type": "box",
                    "layout": "horizontal",
                    "contents": [
                        {
                            "type": "box",
                            "layout": "vertical",
                            "contents": [
                                {
                                    "type": "text",
                                    "text": "today",
                                    "align": "start"
                                },
                                {
                                    "type": "text",
                                    "text": "%s",
                                    "align": "start"
                                }
                            ]
                        },
                        {
                            "type": "box",
                            "layout": "vertical",
                            "contents": [
                                {
                                    "type": "text",
                                    "text": "limit",
                                    "align": "start"
                                },
                                {
                                    "type": "text",
                                    "text": "%s",
                                    "align": "start"
                                }
                            ]
                        },
                        {
                            "type": "box",
                            "layout": "vertical",
                            "contents": [
                                {
                                    "type": "text",
                                    "text": "day",
                                    "align": "end"
                                },
                                {
                                    "type": "text",
                                    "text": "%s",
                                    "align": "end"
                                }
                            ]
                        }
                    ],
                    "backgroundColor": "#DCDCDC",
                    "justifyContent": "space-between",
                    "alignItems": "flex-end",
                    "paddingStart": "xxl"
                },
                "styles": {
                    "header": {
                        "separator": true
                    }
                }
            }
        }
    ]
}
	`, replyToken, amount, category, currentToday, limit, day)

	return bodyJson
}
