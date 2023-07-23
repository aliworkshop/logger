package writers

import (
	"fmt"
	"github.com/aliworkshop/configer"
	"github.com/aliworkshop/error"
	"github.com/aliworkshop/loggerlib/logger"
	"io"
)

type isActive struct {
	Active bool
	Level  string
}

type Writer interface {
	Level() (logger.Level, bool)
	io.Writer
}

func GetWriter(kind string, registry configer.Registry) (Writer, error.ErrorModel) {
	switch kind {
	case "stdout":
		act := new(isActive)
		err := registry.Unmarshal(act)
		if err != nil {
			return nil, error.HandleError(err)
		}
		if act.Active {
			return newStdout(act.Level), nil
		}
		return nil, nil
	case "stderr":
		act := new(isActive)
		err := registry.Unmarshal(act)
		if err != nil {
			return nil, error.HandleError(err)
		}
		if act.Active {
			return newStderr(act.Level), nil
		}
		return nil, nil
	default:
		return nil, error.New(fmt.Errorf("logger writer not found"))
	}
}
