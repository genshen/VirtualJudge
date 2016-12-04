package user

const (
	Login = 2 << iota    //if Login included,indicate that we must visit this page after login
	LoginJSON    //if you are not in login state,will return a json data which indicate you are not login
)

//const for user status
const (
	FREEZING int = iota
	UNACTIVATED
	STATUS_ACTIVE
)

