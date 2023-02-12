package loggerlib

import (
	"errors"
	"github.com/aliworkshop/configlib"
	"github.com/aliworkshop/loggerlib/customzap"
	"github.com/aliworkshop/loggerlib/logger"
	"github.com/aliworkshop/loggerlib/writers"
)

func GetLogger(registry configlib.Registry) (logger.Logger, error) {
	config := new(logger.Config)
	err := registry.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	if config.Type == "" {
		config.Type = "zap"
	}
	var wss []writers.Writer
	for k, _ := range config.Writers {
		ws, err := writers.GetWriter(k, registry.ValueOf("writers."+k))
		if err != nil {
			return nil, err
		}
		if ws != nil {
			wss = append(wss, ws)
		}
	}
	switch config.Type {
	case "zap":
		zl, err := customzap.NewLogger(registry, wss)
		if err != nil {
			return nil, err
		}
		return zl, nil
	}
	return nil, errors.New("no logger type matched")
}
