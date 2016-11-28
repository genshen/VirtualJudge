package models

import "time"

const(
	ProblemTableName = "problem"
	ProblemDetailTableName = "problem_detail"
)

type Problem struct {
	Id             int                `json:"id"`
	ProblemDetail  *ProblemDetail     `orm:"reverse(one)"`
	Title          string             `json:"title"`
	OriginUrl      string             `json:"origin_url"`
	Oj             int8               `json:"oj"`
	OriginId       string             `json:"origin_id"`
	MemLimit       string             `json:"men_limit"`
	TimeLimit      string             `json:"time_limit"`
	Source         string             `json:"source"`
	SourceUrl      string             `json:"source_url"`
	AcCount        int                `json:"ac_count"`
	SubmittedCount int                `json:"submitted_count"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}

type ProblemDetail struct {
	Id           int        `json:"id"`
	Problem      *Problem   `json:"problem_id" orm:"rel(one);index"`
	Describe     string     `json:"describe" orm:"type(text)"`
	Input        string     `json:"input" orm:"type(text)"`
	Output       string     `json:"output" orm:"type(text)"`
	InputSample  string     `json:"input_sample" orm:"type(text)"`
	OutputSample string     `json:"output_sample" orm:"type(text)"`
	Hint         string     `json:"hint" orm:"type(text)"`
	UpdatedAt    time.Time  `json:"updated_at"`
}


func (u *Problem) TableName() string {
	return ProblemTableName
}

func (u *ProblemDetail) TableName() string {
	return ProblemDetailTableName
}