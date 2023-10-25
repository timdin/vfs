package validation

import (
	"fmt"
	"regexp"

	"github.com/timdin/vfs/constants"
)

func InvalidCharacterValidation(args []string) error {
	// Compile the regular expression
	re := regexp.MustCompile(constants.ValidStringPattern)
	fmt.Println(args)
	for _, arg := range args {
		if !re.MatchString(arg) {
			return fmt.Errorf("Argument [%s] contains invalid character", arg)
		}
	}
	return nil
}
