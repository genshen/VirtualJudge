package utils

const (
	LOCAL = iota
	POJ
	HUST
	HDU
	LANG_COUNT
)

const (
	LANG_C = iota
	LANG_CPP
	LANG_JAVA
	LANG_GCC
	LANG_GPP  //G++
)

const (
	STATUS_UNKNOWN_ERROR = iota + 1  //unknown error
	STATUS_KNOWN_ERROR      //known error
	STATUS_DIV_ERROR

	STATUS_SUBMITTING       //will submit to remote oj
	STATUS_SUBMITTED        //have submitted to remote oj
	STATUS_DIV_LOCAL

	STATUS_PENDING
	STATUS_QUEUEING
	STATUS_COMPILING
	STATUS_JUDGING   //including Running & Judging
	STATUS_DIV_REMOTE_PENDING

	STATUS_AC
	STATUS_PE   //Presentation Error
	STATUS_WA   //Wrong Answer
	STATUS_TLE  //Time Limit Exceed
	STATUS_MLE  //Memory Limit Exceed
	STATUS_OLE  //Output Limit Exceed
	STATUS_RE   //Runtime Error
	STATUS_CE   //Compile Error
	STATUS_DIV_END
)
