package models

type Problem struct {
	Id             int64   `json:"id"`
	Title          string  `json:"title"`
	Source         string  `json:"source"`
	OriginUrl      string  `json:"origin_url"`
	Oj             int     `json:"oj"`
	OriginId       string  `json:"origin_id"`
	MemLimit       int     `json:"men_limit"`
	TimeLimit      int     `json:"time_limit"`
	AcCount        int     `json:"ac_count"`
	SubmittedCount int     `json:"submitted_count"`
	CreatedAt      int     `json:"created_at"`
	UpdatedAt      int     `json:"updated_at"`
}
