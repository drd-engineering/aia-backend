package session

type User struct {
	KtpNumber    int64  `json:"KtpNumber"`
	Name         string `json:"Name"`
	Address      string `json:"Address"`
	Gender       string `json:"Gender"`
	PlaceOfBirth string `json:"PlaceOfBirth"`
	Cityzenship  string `json:"Cityzenship"`
	DateOfBirth  string `json:"DateOfBirth"`
}

type LinkStatus struct {
	UserId     string `json:"UserId"`
	User       User
	Token      string `json:"Path"`
	LinkStatus int32  `json:"LinkStatus"`
}
