package constant

type command struct {
	List    string
	Logbook string
	Login   string
	Status  string
}

var Command = registerCommand()

func registerCommand() *command {
	return &command{
		List:    "list",
		Logbook: "logbook",
		Login:   "login",
		Status:  "status",
	}
}
