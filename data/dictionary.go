package dictionary

var EnglishToSpanish = map[string]string{

	"Hello":  "Hola",
	"Car":    "Auto",
	"Day":    "Dia",
	"Server": "Servidor",
}

func Translation(word string) string {
	return EnglishToSpanish[word]
}
