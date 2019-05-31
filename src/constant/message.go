package constant

type message struct {
	EditNotGranted      string
	EditFailed          string
	EditFormat          string
	EditSuccess         string
	Help                string
	HelpEdit            string
	HelpList            string
	HelpLogbook         string
	HelpLogin           string
	HelpRequestEdit     string
	HelpStatus          string
	LogbookFailed       string
	LogbookFormat       string
	LogbookSuccess      string
	LoginFailed         string
	LoginFormat         string
	LoginSuccess        string
	NotLoggedIn         string
	RequestEditFailed   string
	RequestEditFormat   string
	RequestEditReminder string
	RequestEditSuccess  string
	SubmitReminder      string
}

var Message = registerMessage()

func registerMessage() *message {
	return &message{
		EditNotGranted:      "You can't edit logbook on that day.",
		EditFailed:          "Edit logbook failed, please try again.",
		EditFormat:          "Please edit logbook with the following format:\nedit\nDD/MM/YYYY\nCLOCKIN\nCLOCKOUT\nACTIVITY\nDESCRIPTION",
		EditSuccess:         "Logbook has been successfully updated.",
		Help:                "Hi, welcome to Logbook Broker.\nWe are a portal for Binusian to easily submit their daily 3+1 logbook.\nFirst if you want to use our service you need to login first, you can login by using your industry.socs.binus.ac.id/learning-plan account\n\nFor more detailed information about our feature you can type  the following message:\nhelp login\nhelp logbook\nhelp status\nhelp list\nhelp edit\n help request edit",
		HelpEdit:            "To edit your logbook, you need to unlock edit button for that day, to unlock the edit button, you need to request edit for that day.\nTo edit logbook you can send message with the following format:\nedit\nDD/MM/YYYY\nCLOCKIN\nCLOCKOUT\nACTIVITY\nDESCRIPTION",
		HelpList:            "You can check your last 7 days logbook by sending us a message with text: list",
		HelpLogbook:         "To submit your today logbook, you can send message with the following format:\nlogbook\nCLOCKIN\nCLOCKOUT\nACTIVITY\nDESCRIPTION",
		HelpLogin:           "You only need to login once and then you can enjoy our service.\nIf you want to change user then just login again, we will automatically replace your login credential.\n\nTo login into our service you can send message with the following format:\nlogin\nUSERNAME\nPASSWORD",
		HelpRequestEdit:     "To request edit logbook, you can access menu list and press the request edit button. You can request edit logbook every friday 9AM - 12PM, you can only request 2 unfilled logbook each week.",
		HelpStatus:          "You can check your today logbook by sending us a message with text: status",
		LogbookFailed:       "Submit logbook failed, please try again.",
		LogbookFormat:       "Please submit logbook with the following format:\nlogbook\nCLOCKIN\nCLOCKOUT\nACTIVITY\nDESCRIPTION",
		LogbookSuccess:      "Today logbook has been successfully inserted.",
		LoginFailed:         "Login failed, please try again.",
		LoginFormat:         "Please login with the following format:\nlogin\nUSERNAME\nPASSWORD",
		LoginSuccess:        "You're successfully logged in as",
		NotLoggedIn:         "You're not logged in, please login first.",
		RequestEditFailed:   "Request edit failed, please try again.",
		RequestEditFormat:   "Please request edit logbook with the following format:\nrequest\nURL",
		RequestEditReminder: "Hi, don't forget to request edit logbook if you got any unfilled logbook, you can only request edit logbook every Friday 9AM - 12PM.",
		RequestEditSuccess:  "Request edit logbook success, please wait until 12PM for the edit logbook feature open.",
		SubmitReminder:      "Hi, don't forget to submit your logbook today.",
	}
}
