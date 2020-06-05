package models

type Msg struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     string `json:"age"`
	Address string `json:"address"`
}

func InsertExel(sql string) bool {
	db.Exec(sql)
	return true
}

func GetExelDate() (msg []Msg) {
	db.Select("name, age,address").Find(&msg)

	return
}

func GetExelDateTotal1() (count int) {
	db.Model(&Msg{}).Count(&count)

	return
}
