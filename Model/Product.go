package Model

type Product struct {
	Code  string `json: "code"` //char(3)
	Name  string `json: "name"` //nvarchar(50)
	Type  string `json: "type"` //char(3)
	Price uint   `json: "price"`
}
