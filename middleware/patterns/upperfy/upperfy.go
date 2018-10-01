package upperfy

import "strings"

// Textfy :
type Textfy string

// Arg :
type Args struct {
	Text string
}

// UpperText :
func (u *Textfy) UpperText(arg *Args, reply *string) error {
	*reply = strings.ToUpper(arg.Text)
	return nil
}
