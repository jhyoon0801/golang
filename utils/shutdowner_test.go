package utils

import (
	"strings"
	"testing"
	"time"
)

func TestShutdowner(t *testing.T) {
	shutdowner := new(ShutDowner)
	shutdowner.Init()
	shutdowner.SetCallback(callbackFunc)
	go func() {
		time.Sleep(5 * time.Second)
		shutdowner.Shutdown()
	}()
	shutdowner.WaitSignal(t, "test parameter")
}

func callbackFunc(param ...interface{}) error {
	convertedParam := param[0].([]interface{})
	a := convertedParam[0].(*testing.T)
	if !strings.EqualFold(convertedParam[1].(string), "test parameter") {
		a.Error("Parameters are not same")
	}
	return nil
}
