// SPDX-License-Identifier: AGPL-3.0-only
// Copyright (C) 2024-2026 Mikhail Kantur

package util

import (
	"fmt"
	"go.slink.ws/logging"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func HandleStopSignal(duration time.Duration) {
	go func() {
		qc := make(chan os.Signal, 1)
		signal.Notify(qc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		for _ = range qc {
			fmt.Print("\r")
			logging.GetLogger("util").Trace("handle SIGINT")
			time.Sleep(duration)
			os.Exit(0)
		}
	}()
}
