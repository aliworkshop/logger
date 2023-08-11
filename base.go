package logger

import (
	"errors"
	"github.com/aliworkshop/configer"
	"github.com/aliworkshop/logger/writers"
)

type config struct {
	Type    string
	Writers map[string]interface{}
}

func GetLogger(registry configer.Registry) (Logger, error) {
	cfg := new(config)
	err := registry.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	if cfg.Type == "" {
		cfg.Type = "zap"
	}
	var wss []writers.Writer
	for k, _ := range cfg.Writers {
		ws, err := writers.GetWriter(k, registry.ValueOf("writers."+k))
		if err != nil {
			return nil, err
		}
		if ws != nil {
			wss = append(wss, ws)
		}
	}
	switch cfg.Type {
	case "zap":
		zl, err := NewLogger(registry, wss)
		if err != nil {
			return nil, err
		}
		return zl, nil
	}
	return nil, errors.New("no logger type matched")
}
