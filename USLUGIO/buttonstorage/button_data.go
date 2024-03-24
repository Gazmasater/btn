package buttonstorage

type Button struct {
	Selector string `json:"selector"`
	Name     string `json:"name"`
}

var LoginMap = map[string]string{
	"lew1968@list.ru":        "ГИПЕРМАРКЕТ УСЛ",
	"vash-agent48@yandex.ru": "ГАЗлайф",
	"vashagent48@gmail.com":  "Сервис Советски",
}
