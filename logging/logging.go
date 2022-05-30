package logging

import (
	"github.com/withmandala/go-log"
	"os"
)

var Default *log.Logger

func init() {
	Default = log.New(os.Stderr)
}
