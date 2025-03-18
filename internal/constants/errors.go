package constants

import "errors"

var (
	ErrorEmptyStack        = errors.New("empty stack")
	ErrorIllegalMove       = errors.New("illegal move")
	ErrorInvalidDeckCounts = errors.New("invalid deck counts, must be five groups of five cards")
	ErrorInvalidCardCount  = errors.New("invalid card count")
	ErrorInvalidStack      = errors.New("invalid stack")
	ErrorInvalidUserID     = errors.New("invalid user id")
	ErrorInvalidUserSecret = errors.New("invalid user secret")
	ErrorUserInvalidName   = errors.New("user invalid name")
	ErrorUserNotFound      = errors.New("user not found")
	ErrorWrongUserTurn     = errors.New("wrong user turn")
)
