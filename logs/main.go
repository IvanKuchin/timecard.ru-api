package logs

import (
	"os"
	"os/signal"

	"go.uber.org/zap"
)

var Sugar *zap.SugaredLogger

func init() {
	l, _ := zap.NewDevelopment()
	Sugar = l.Sugar()

	go func() {
		defer l.Sync()

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		Sugar.Debug("Sync log system before exit")
	}()
}
