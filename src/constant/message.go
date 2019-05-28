package constant

type message struct {
	EditFailed          string
	EditFormat          string
	EditSuccess         string
	Help                string
	HelpList            string
	HelpLogbook         string
	HelpLogin           string
	HelpStatus          string
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
		EditFailed:          "Edit logbook failed, please try again.",
		EditFormat:          "Edit Token can be found by acessing menu list and press the edit button on the logbook you want to edit, to unlock the edit button, you need to request edit first.\nPlease edit logbook with the following format:\nedit\nEDITTOKEN\nCLOCKIN\nCLOCKOUT\nACTIVITY\nDESCRIPTION",
		EditSuccess:         "Logbook has been successfully updated.",
		Help:                "Hi, welcome to Logbook Broker.\nWe are a portal for Binusian to easily submit their daily 3+1 logbook.\nFirst if you want to use our service you need to login first, you can login by using your industry.socs.binus.ac.id/learning-plan account\n\nFor more detailed information about our feature you can type  the following message:\nhelp login\nhelp logbook\nhelp status\nhelp list",
		HelpList:            "You can check your last 7 days logbook by sending us a message with text: list",
		HelpLogbook:         "To submit your today logbook you can send message with the following format:\nlogbook\nCLOCKIN\nCLOCKOUT\nACTIVITY\nDESCRIPTION",
		HelpLogin:           "You only need to login once and then you can enjoy our service.\nIf you want to change user then just login again, we will automatically replace your login credential.\n\nTo login into our service you can send message with the following format:\nlogin\nUSERNAME\nPASSWORD",
		HelpStatus:          "You can check your today logbook by sending us a message with text: status",
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
