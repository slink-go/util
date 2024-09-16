package util

import (
	"fmt"
	"github.com/slink-go/logging"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func HandleStopSignal(duration time.Duration) {
	go func() {
		qc := make(chan os.Signal)
		signal.Notify(qc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		for _ = range qc {
			fmt.Print("\r")
			logging.GetLogger("util").Trace("handle SIGINT")
			time.Sleep(duration)
			os.Exit(0)
		}
	}()
}
