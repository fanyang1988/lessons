package main

import (
	"github.com/cihub/seelog"
	"github.com/fanyang1988/lessons/logger"
)

func main() {
	logger.InitSeelogCfg("./log/", "example_log")
	defer seelog.Flush()

	func() {
		defer func() {
			err := recover()
			logger.LogIfPanic(err)
		}()

		testLogger := logger.SeeLogLogger()

		for i := 0; i < 100; i++ {
			testLogger.Debugf("debugf %v", i)
			testLogger.Infof("info %v", i)
			testLogger.Errorf("errf %v", i)
		}

		panic("panic !!!")
	}()
}
