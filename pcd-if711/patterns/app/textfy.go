package app

import (
	"strings"
)

// Textfy :
type Textfy struct {
}

// Args :
type Args struct {
	Text string
}

// UpperText :
func (u *Textfy) UpperText(args Args, reply *string) error {
	*reply = strings.ToUpper(args.Text)
	return nil
}
