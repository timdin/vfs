package validation

import (
	"fmt"

	"github.com/timdin/vfs/constants"
)

func ValidLength(args []string) error {
	for _, arg := range args {
		if len(arg) > constants.ValidLength {
			return fmt.Errorf("[%s] too long, valid length: <=%d", arg, constants.ValidLength)
		}
	}
	return nil
}
