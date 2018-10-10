package application

import "strings"

// Textfy :
type Textfy string

// Args :
type Args struct {
	Text string
}

// UpperText :
func (u *Textfy) UpperText(args *Args, reply *string) error {
	*reply = strings.ToUpper(args.Text)
	return nil
}
