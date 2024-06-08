package er

var messages = map[string]string{
	"1": "Oops! Something went wrong. Please try later",
	"2": "Unauthorized Request!",
	"3": "Record not found!",
	"4": "User not found!",
	"5": "Invalid Request Body!",
}

var codes = map[Code]string{
	UncaughtException:  "1",
	Unauthorized:       "2",
	RecordNotFound:     "3",
	UserNotFound:       "4",
	InvalidRequestBody: "5",
}
