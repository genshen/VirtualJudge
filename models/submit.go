package models

import "time"

const (
	ContextIdDefault = -1  //means no contest was involved in
)

type Submission struct {
	Id               int          `json:"id"`          //
	Contest          int          `json:"contest_id" orm:"default(-1)"`
	StatusCode       int8         `json:"status"`
	ErrorDetail      string       `json:"error_detail"`
	ProblemId        int          `json:"problem_id"`  //
	OjType           int8         `json:"oj_type"`     // join
	OriginId         string       `json:"origin_id"`   // join
	OriginRunId      string       `json:"origin_run_id"`
	Language         int8         `json:"language"`    //
	ExecuteTime      string       `json:"execute_time"`
	Memory           string       `json:"memory"`
	SourceCode       string       `json:"source_code"` //
	CodeLength       int          `json:"code_length"` //
	Public           bool         `json:"public"`      //
	UserId           int          `json:"user_id"`     //
	UserName         string       `json:"username"`    // join
	OriginAccountId  string       `json:"_"`
	QueryCount       int          `json:"query_time"`
	OriginSubmitTime time.Time    `json:"origin_submit_time" orm:"null"`
	CreatedAt        time.Time    `json:"created_at"`  //
	UpdatedAt        time.Time    `json:"updated_at"`  //
}