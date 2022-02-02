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

// register cusom validator tags
func init() {
	valid.TagMap["path"] = valid.Validator(func(str string) bool {
		if str == "" {
			return true
		}

		return rxPath.MatchString(str)
	})
	valid.TagMap["pathdis"] = valid.Validator(func(str string) bool {
		if str == "" || str == "-" {
			return true
		}

		return rxPath.MatchString(str)
	})
}
