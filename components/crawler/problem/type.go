package problem

import (
	"errors"
	"gensh.me/VirtualJudge/components/crawler/utils"
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

func CrawlerProblem(id string, ojType int8) (*ProblemMeta, error) {
	var pc ProblemCrawler
	switch ojType {
	//case Self:
	case utils.POJ:
		pc = PojProblemCrawler{}
	default:
		return &ProblemMeta{}, errors.New("no OJ matched")
	}
	return pc.RequestProblem(id)

}