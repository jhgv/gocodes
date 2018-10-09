package app

import "strings"

// Textfy :
type Textfy string

// Args :
type Args struct {
	Text string
}

// UpperText :
func (u *Textfy) UpperText(message string, reply *string) error {
	*reply = strings.ToUpper(message)
	return nil
}
