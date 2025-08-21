package library

import "time"

type RequestType string

const (
	Borrow RequestType = "BORROW"
	Return RequestType = "RETURN"
)

type UserRequest struct {
	ID          string
	UserID      string
	BookId      string
	Time        time.Time
	RequestType RequestType
}
