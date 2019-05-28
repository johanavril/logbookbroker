package constant

type command struct {
	Edit        string
	Help        string
	HelpList    string
	HelpLogin   string
	HelpLogbook string
	HelpStatus  string
	List        string
	Logbook     string
	Login       string
	Status      string
}

var Command = registerCommand()

func registerCommand() *command {
	return &command{
		Edit:        "edit",
		Help:        "help",
		HelpList:    "help list",
		HelpLogin:   "help login",
		HelpLogbook: "help logbook",
		HelpStatus:  "help status",
		List:        "list",
		Logbook:     "logbook",
		Login:       "login",
		Status:      "status",
	}
}
