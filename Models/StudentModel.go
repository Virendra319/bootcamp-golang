package Models

type Student struct {
	Id       uint     `json:"id"`
	Name     string   `json:"first name"`
	LastName string   `json:"last name"`
	DOB      string   `json:"dob"`
	Address  string   `json:"address"`
	Scores   []Scores `json:"scores"`
}

func (b *Student) TableName() string {
	return "students"
}
