package input

import (
	"io/ioutil"
	"os"
)

func ReadString() (string, error) {
	b, err := ioutil.ReadAll(os.Stdin)
	return string(b), err
}
