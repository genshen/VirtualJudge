package status

import "gensh.me/VirtualJudge/components/crawler/utils"

//get submission status by name(such as:"Accept"),used by status package and submitter package
var ojStatuses = []map[string]int8{
	//match with accounts.OJ[0]
	map[string]int8{
		"Pending":utils.STATUS_PENDING,
		"Compiling":utils.STATUS_COMPILING,
		"Running & Judging":utils.STATUS_JUDGING,
		"Accepted":utils.STATUS_AC,
		"Presentation Error":utils.STATUS_PE,
		"Time Limit Exceeded":utils.STATUS_TLE,
		"Memory Limit Exceeded":utils.STATUS_MLE,
		"Wrong Answer":utils.STATUS_WA,
		"Runtime Error":utils.STATUS_RE,
		"Output Limit Exceeded":utils.STATUS_OLE,
		"Compile Error":utils.STATUS_CE,
	},
}

/**HU
Accepted
Presentation Error
Wrong Answer
Time Limit Exceed
Memory Limit Exceed
Output Limit Exceed
Runtime Error
Compile Error
Compile OK

Test Running Done
Pending
Pending Rejudging
Compiling
Running & Judging
*/

func GetStatusByOJType(ojType int8, status string) int8 {
	return GetStatusByOJTypeDefault(ojType, status, utils.STATUS_PENDING)
}

func GetStatusByOJTypeDefault(ojType int8, status string, defaultStatus int8) int8 {
	if value, ok := ojStatuses[ojType][status]; ok {
		return value
	} else {
		return defaultStatus
	}
}