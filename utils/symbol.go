package utils

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func GetSymbolCurrentcyFormat(data float64) string {
	p := message.NewPrinter(language.English)
	currencySymbol := "฿"
	result := p.Sprintf("%s%.2f", currencySymbol, data)
	return result

}
