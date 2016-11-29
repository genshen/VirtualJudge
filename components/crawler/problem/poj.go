package problem

import (
	"regexp"
	"errors"
	"github.com/astaxie/beego/httplib"
	"gensh.me/VirtualJudge/components/crawler/utils"
)
//var r = regexp.MustCompile(`[\S|\s]+?ptt[\S|\s]+?>([\S|\s]*?)</div>[\S|\s]+?plm[\S|\s]+?</b>([\S|\s]*?)</td>[\S|\s]+?</b>([\S|\s]*?)</td>[\S|\s]+?ptx[\S|\s]+?>([\S|\s]*?)</div><p class="pst">[\S|\s]+?ptx[\S|\s]+?>([\S|\s]*?)</div><p class="pst">[\S|\s]+?ptx[\S|\s]+?>([\S|\s]*?)</div><p class="pst">[\S|\s]+?sio">([\S|\s]*?)</pre><p class="pst">[\S|\s]+?sio">([\S|\s]*?)</pre><p class="pst">[\S|\s]+?ptx[\S|\s]+?>([\S|\s]*?)</div>`)
//var r2 = regexp.MustCompile(`[\S|\s]+?ptt[\S|\s]+?>([\S|\s]*?)</div>[\S|\s]+?plm[\S|\s]+?</b>([\S|\s]*?)</td>[\S|\s]+?</b>([\S|\s]*?)</td>[\S|\s]+?ptx[\S|\s]+?>([\S|\s]+?)</div><p class="pst">[\S|\s]+?ptx[\S|\s]+?>([\S|\s]*?)</div><p class="pst">[\S|\s]+?ptx[\S|\s]+?>([\S|\s]*?)</div><p class="pst">[\S|\s]+?sio">([\S|\s]*?)</pre><p class="pst">[\S|\s]+?sio">([\S|\s]*?)</pre><p class="pst">[\S|\s]+?ptx[\S|\s]+?>([\S|\s]*?)</pre></div><p class="pst">[\S|\s]+?ptx[\S|\s]+?>([\S|\s]*?)</div>`)

const (
	pojDomain = "http://poj.org/problem?id="
)

var sourceRegexp = regexp.MustCompile(`<a href="([\S|\s]*?)">([\S|\s]*?)</a>`)

type PojProblemCrawler struct {

}

func (p PojProblemCrawler)RequestProblem(problemId string) (*ProblemMeta, error) {
	problemMeta := ProblemMeta{OriginUrl:pojDomain + problemId}
	s, err := httplib.Get(problemMeta.OriginUrl).Debug(true).String()
	if err != nil {
		return &problemMeta, err
	}

	title, po := utils.FindMatchString(s, `"ptt"`, ">", "</div>")
	if po != -1 && po < len(s) {
		//we think:if the title not exists,this is a not found page
		problemMeta.Title = title
		s = s[po:]
	} else {
		return &problemMeta, errors.New("problem not found")
	}

	time, po := utils.FindMatchString(s, `"plm"`, "</b>", "</td>")
	if po != -1 && po < len(s) {
		problemMeta.TimeLimit = time
		s = s[po:]
	}

	mem, po := utils.FindMatchString(s, "", "</b>", "</td>")
	if po != -1 && po < len(s) {
		problemMeta.MemLimit = mem
		s = s[po:]
	}

	describe, po := utils.FindMatchString(s, `"ptx"`, ">", "</div>")
	if po != -1 && po < len(s) {
		problemMeta.Describe = describe
		s = s[po:]
	}

	input, po := utils.FindMatchString(s, `"ptx"`, ">", "</div>")
	if po != -1 && po < len(s) {
		problemMeta.Input = input
		s = s[po:]
	}

	output, po := utils.FindMatchString(s, `"ptx"`, ">", "</div>")
	if po != -1 && po < len(s) {
		problemMeta.Output = output
		s = s[po:]
	}

	inputSample, po := utils.FindMatchString(s, "", "sio\">", "</pre><p class=\"pst\">")
	if po != -1 && po < len(s) {
		problemMeta.InputSample = inputSample
		s = s[po:]
	}

	outputSample, po := utils.FindMatchString(s, "", "sio\">", "</pre><p class=\"pst\">")
	if po != -1 && po < len(s) {
		problemMeta.OutputSample = outputSample
		s = s[po:]
	}

	hint, po := utils.FindMatchString(s, `Hint`, ">", "</div>")
	if po != -1 && po < len(s) {
		problemMeta.Hint = hint
		s = s[po:]
	}

	source, po := utils.FindMatchString(s, `"ptx"`, ">", "</div>")
	if po != -1 && po < len(s) {
		if sourceData := sourceRegexp.FindStringSubmatch(source); sourceData != nil {
			problemMeta.SourceUrl = sourceData[1]
			problemMeta.Source = sourceData[2]
		}
	}
	return &problemMeta, nil
}