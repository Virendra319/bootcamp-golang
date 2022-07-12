package Models

type Scores struct {
	Id        uint   `json:"scoreId"`
	StudentId uint   `json:"studentId"`
	Subject   string `json:"subject"`
	Marks     int    `json:"marks"`
}

func (b *Scores) TableName() string {
	return "scores"
}
