package writers

import (
	"fmt"
	"github.com/aliworkshop/configlib"
	"github.com/aliworkshop/errorslib"
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

func GetWriter(kind string, registry configlib.Registry) (Writer, errorslib.ErrorModel) {
	switch kind {
	case "stdout":
		act := new(isActive)
		err := registry.Unmarshal(act)
		if err != nil {
			return nil, errorslib.HandleError(err)
		}
		if act.Active {
			return newStdout(act.Level), nil
		}
		return nil, nil
	case "stderr":
		act := new(isActive)
		err := registry.Unmarshal(act)
		if err != nil {
			return nil, errorslib.HandleError(err)
		}
		if act.Active {
			return newStderr(act.Level), nil
		}
		return nil, nil
	default:
		return nil, errorslib.New(fmt.Errorf("logger writer not found"))
	}
}
