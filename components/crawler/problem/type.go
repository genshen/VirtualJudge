package problem

import (
	"errors"
)

type ProblemMeta struct {
	Title        string
	OriginUrl    string
	TimeLimit    string
	MemLimit     string
	Describe     string
	Input        string
	Output       string
	InputSample  string
	OutputSample string
	Hint         string
	Source       string
	SourceUrl    string
}

type ProblemCrawler interface {
	RequestProblem(proId string) (*ProblemMeta, error)
}

//remember to follow the order in crawler/utils/values.go->const
var problemCrawlerInterfaces = []ProblemCrawler{new(PojProblemCrawler)}

func CrawlerProblem(id string, ojType int) (*ProblemMeta, error) {
	var pc ProblemCrawler
	if index := ojType - 1;index < len(problemCrawlerInterfaces) && index >= 0 {
		// 0 means Local problems
		pc = problemCrawlerInterfaces[index]
		return pc.RequestProblem(id)
	} else {
		return &ProblemMeta{}, errors.New("no OJ matched")
	}
}