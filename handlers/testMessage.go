package handlers

import "strings"

func TestMessageNLP(message string) string {
	var replyText string

	if strings.Contains(message, "ขายอะไร") {
		replyText = "- กาแฟ และเครื่องดื่มค่ะ เช่น\n- คาปูชิโน่\n- ลาเต้\n- อเมริกาโน่\n- ชาไทย"
	} else if strings.Contains(message, "สวัสดี") {
		replyText = "สวัสดีค่ะ"
	} else {
		replyText = "ขออภัยค่ะ ฉันไม่เข้าใจคำถาม กรุณาถามคำถามใหม่ค่ะ"
	}

	return replyText
}
