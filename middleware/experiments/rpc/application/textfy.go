package application

import (
	"github.com/jhgv/gocodes/middleware/experiments/rpc/core"
	"strings"
)

// Textfy :
type Textfy string

// UpperText :
func (u *Textfy) UpperText(args *core.Args, reply *string) error {
	*reply = strings.ToUpper(args.Text)
	return nil
}
