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

func createWordIndex(texts []string) map[string]int {
	wordIndex := map[string]int{}
	index := 1 // เริ่มที่ 1 เพื่อให้ 0 ใช้สำหรับ padding
	for _, text := range texts {
		words := strings.Fields(text)
		for _, word := range words {
			if _, exists := wordIndex[word]; !exists {
				wordIndex[word] = index
				index++
			}
		}
	}
	return wordIndex
}

func textsToSequences(texts []string, wordIndex map[string]int) [][]int {
	sequences := [][]int{}
	for _, text := range texts {
		words := strings.Fields(text)
		sequence := []int{}
		for _, word := range words {
			if idx, exists := wordIndex[word]; exists {
				sequence = append(sequence, idx)
			}
		}
		sequences = append(sequences, sequence)
	}
	return sequences
}

func padSequences(sequences [][]int, maxLen int) [][]int {
	padded := [][]int{}
	for _, seq := range sequences {
		if len(seq) > maxLen {
			seq = seq[:maxLen] // ตัดความยาวให้สั้นลง
		} else {
			for len(seq) < maxLen {
				seq = append(seq, 0) // เติม 0 ด้านท้าย
			}
		}
		padded = append(padded, seq)
	}
	return padded
}
