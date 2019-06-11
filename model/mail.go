package model

type OutgoingMailConfig struct {
	SMTPHost string
	Port     int
	UseTLS   bool
	Username string
	Password string
}

type Mail struct {
	ToEmail   string
	ToName    string
	FromEmail string
	Subject   string
	BodyHTML  string
}
