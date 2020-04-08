package ESign

import (
	"time"
)

type User struct {
	KtpNumber    int64  `json:"KtpNumber"`
	Name         string `json:"Name"`
	Address      string `json:"Address"`
	Gender       string `json:"Gender"`
	PlaceOfBirth string `json:"PlaceOfBirth"`
	Cityzenship  string `json:"Cityzenship"`
	DateOfBirth  string `json:"DateOfBirth"`
}

type SessionLink struct {
	UserId string    `json:"UserId"`
	Path   string    `json:"Path"`
	Valid  time.Time `json:"Valid"`
}

type UserSignature struct {
	UserId    string   `json:"UserId"`
	Signature []string `json:"Signature"`
}

type UserSignatureStatus struct {
	UserId     string `json:"UserId"`
	SignStatus int32  `json:"SignStatus"`
}
