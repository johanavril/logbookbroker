package constant

type message struct {
	LogbookFailed       string
	LogbookFormat       string
	LogbookSuccess      string
	LoginFailed         string
	LoginFormat         string
	LoginSuccess        string
	NotLoggedIn         string
	RequestEditReminder string
	SubmitReminder      string
}

var Message = registerMessage()

func registerMessage() *message {
	return &message{
		LogbookFailed:       "Submit logbook failed, please try again.",
		LogbookFormat:       "Please submit logbook with the following format:\nlogbook\nCLOCKIN\nCLOCKOUT\nACTIVITY\nDESCRIPTION",
		LogbookSuccess:      "Today logbook has been successfully inserted.",
		LoginFailed:         "Login failed, please try again.",
		LoginFormat:         "Please login with the following format:\nlogin\nUSERNAME\nPASSWORD",
		LoginSuccess:        "You're successfully logged in as",
		NotLoggedIn:         "You're not logged in, please login first.",
		RequestEditReminder: "Hi, don't forget to request edit logbook if you got any unfilled logbook, you can only request edit logbook every Friday 9AM - 12PM.",
		SubmitReminder:      "Hi, don't forget to submit your logbook today.",
	}
}
