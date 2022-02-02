package yaml

import (
	"io"
	"os"

	valid "github.com/asaskevich/govalidator"
	"gopkg.in/yaml.v2"

	"go.sancus.dev/core/errors"
)

func LoadReader(f io.Reader, mapping func(string) string, c interface{}) error {
	b, err := io.ReadAll(f)
	if err != nil {
		// read error
		return errors.Wrap(err, "ReadAll")
	}

	if mapping != nil {
		// expand $var and ${var}
		s := string(b)
		s = os.Expand(s, mapping)
		b = []byte(s)
	}

	if len(b) == 0 {
		// empty file
		return nil
	} else if err := yaml.Unmarshal(b, c); err != nil {
		// failed to decode
		return errors.Wrap(err, "Unmarshal")
	} else if ok, err := valid.ValidateStruct(c); ok {
		// ready
		return nil
	} else if err != nil {
		// failed to validate
		return errors.Wrap(err, "ValidateStruct")
	} else {
		// failed but no error given (can this happen?)
		return errors.New("ValidateStruct")

	}
}

func LoadFile(filename string, mapper func(string) string, c interface{}) error {
	var file io.ReadCloser

	if filename == "-" {
		file = os.Stdin
	} else if file, err := os.Open(filename); err != nil {
		return errors.Wrap(err, "Open")
	} else {
		defer file.Close()
	}

	return LoadReader(file, mapper, c)
}

func LoadFileEnv(filename string, c interface{}) error {
	return LoadFile(filename, os.Getenv, c)
}
