package constant

type url struct {
	Broadcast          string
	Logbook            string
	LogbookInsert      string
	LogbookRequestEdit string
	Login              string
	Profile            string
}

var URL = registerUrl()

func registerUrl() *url {
	return &url{
		Broadcast:          "https://api.line.me/v2/bot/message/broadcast",
		Logbook:            "https://industry.socs.binus.ac.id/learning-plan/student/log-book",
		LogbookInsert:      "https://industry.socs.binus.ac.id/learning-plan/student/log-book/insert",
		LogbookRequestEdit: "https://industry.socs.binus.ac.id/learning-plan/student/log-book/edit-request",
		Login:              "https://industry.socs.binus.ac.id/learning-plan/auth/login",
		Profile:            "https://industry.socs.binus.ac.id/learning-plan/profile",
	}
}
