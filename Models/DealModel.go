package Models

type Deal struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}

func (b *Deal) TableName() string {
	return "deal"
}
