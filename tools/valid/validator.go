package valid

import (
	"regexp"

	valid "github.com/asaskevich/govalidator"
)

var (
	rxPath = regexp.MustCompile(valid.UnixPath)
)

// IsPath checks if string is a valid path
func IsPath(str string) bool {
	return rxPath.MatchString(str)
}
