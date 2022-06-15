package postgres

type UserDTO struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Age  int16  `json:"age"`
}
