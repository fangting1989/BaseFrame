package utils

import "github.com/jbrodriguez/mlog"

func LogInfo(msg string, logtype string) {
	if logtype == "error" {
		mlog.Start(mlog.LevelInfo, "./logfile/error.log")
	} else {
		mlog.Start(mlog.LevelInfo, "./logfile/app.log")
	}
	mlog.Info(msg)
}
