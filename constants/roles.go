package constants

// Enum for roles
type UserRole int

const (
	RoleAdmin UserRole = iota + 1
	
	RoleCount
)