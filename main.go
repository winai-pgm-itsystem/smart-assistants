package main

import (
	"fmt"
	"log"
	"strings"

	tf "github.com/galeone/tensorflow/tensorflow/go"
)

func main() {
	// โหลดโมเดลจาก SavedModel directory
	model, err := tf.LoadSavedModel("BiLSTM_sentiment_model", []string{"serve"}, nil)
	if err != nil {
		log.Fatalf("Error loading model: %v", err)
	}
	defer model.Session.Close()

	// ตัวอย่างข้อความ input ที่จะทำการทำนาย
	input := "This is a great product!"

	// Preprocess ข้อมูล (tokenization, padding)
	xTest := manualTokenization(input)
	xTest = padSequence(xTest, 100) // max_length = 100

	// สร้าง Tensor จากข้อมูล input
	tensor, err := tf.NewTensor(xTest)
	if err != nil {
		log.Fatalf("Error creating tensor: %v", err)
	}

	// รันโมเดลเพื่อทำการทำนาย
	result, err := model.Session.Run(
		map[tf.Output]*tf.Tensor{
			model.Graph.Operation("serving_default_input_1").Output(0): tensor, // ปรับชื่อ input node ให้ตรง
		},
		[]tf.Output{
			model.Graph.Operation("StatefulPartitionedCall").Output(0), // ปรับชื่อ output node ให้ตรง
		},
		nil,
	)
	if err != nil {
		log.Fatalf("Error running the model: %v", err)
	}

	// ตรวจสอบผลลัพธ์
	prediction := result[0].Value().([][]float32)[0][0]

	// ตัดสินใจว่าผลลัพธ์เป็น positive หรือ negative
	var response string
	if prediction > 0.5 {
		response = "positive"
	} else {
		response = "negative"
	}

	// แสดงผลลัพธ์
	fmt.Printf("Prediction: %s\n", response)
}

// ฟังก์ชั่นเพื่อทำการ tokenization (แปลงข้อความเป็น token) ใน Go
func manualTokenization(input string) []int {
	// ตัวอย่างการแปลงข้อความเป็นตัวเลขโดยใช้ dictionary (ง่ายๆ)
	dictionary := map[string]int{
		"great":   1,
		"product": 2,
		"this":    3,
		"is":      4,
		"a":       5,
	}

	words := strings.Fields(input) // แยกคำจาก input
	tokens := make([]int, len(words))

	for i, word := range words {
		// หากคำมีอยู่ใน dictionary ให้แปลงเป็น token
		if token, exists := dictionary[word]; exists {
			tokens[i] = token
		} else {
			tokens[i] = 0 // หากไม่มีใน dictionary, ให้เป็น 0
		}
	}
	return tokens
}

// ฟังก์ชั่นเพื่อทำการ padding sequence ให้ยาวตาม max_length
func padSequence(sequence []int, maxLength int) []int {
	if len(sequence) >= maxLength {
		return sequence[:maxLength]
	}

	// ถ้าความยาว sequence น้อยกว่า maxLength, ทำการ padding
	paddedSequence := make([]int, maxLength)
	copy(paddedSequence, sequence)
	return paddedSequence
}
