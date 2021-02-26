package consts

const (
	//response code
	RspStatusOkCode         int    = 10000
	RspStatusOkMsg          string = "Success"

	//common fail
	CommonServerFailCode int    = 11000
	CommonServerFailMsg  string = "Server internal fail "
	CommonParamsCheckFailCode int    = 11001
	CommonParamsCheckFailMsg  string = "Invalid parameter"
	CommonAuthFailCode int    = 11002
	CommonAuthFailMsg  string = "Authentication fail"


	//DB fail. [12000~13000)
	DBCommonFailCode int    = 12000
	DBCommonFailMsg  string = "Database operation common fail"
	DBInvalidIdCode int    = 12001
	DBInvalidIdMsg  string = "Database invalid id"
	//business fail. [13000~14000)

)
