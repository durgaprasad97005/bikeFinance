package constants

// Enum for status
type Status int

const (
	StatusPending  Status = iota + 1
	StatusApproved 
	StatusRejected 
	StatusInActive 

	StatusCount
)
