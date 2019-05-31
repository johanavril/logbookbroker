package constant

type command struct {
	Edit            string
	Help            string
	HelpEdit        string
	HelpList        string
	HelpLogin       string
	HelpLogbook     string
	HelpRequestEdit string
	HelpStatus      string
	List            string
	Logbook         string
	Login           string
	RequestEdit     string
	Status          string
}

var Command = registerCommand()

func registerCommand() *command {
	return &command{
		Edit:            "edit",
		Help:            "help",
		HelpEdit:        "help edit",
		HelpList:        "help list",
		HelpLogin:       "help login",
		HelpLogbook:     "help logbook",
		HelpRequestEdit: "help request edit",
		HelpStatus:      "help status",
		List:            "list",
		Logbook:         "logbook",
		Login:           "login",
		RequestEdit:     "request",
		Status:          "status",
	}
}
