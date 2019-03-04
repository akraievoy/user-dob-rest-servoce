package service_utils

import (
	"context"
	"os"
	"os/signal"
)

func CancelOnTermSignal(cancelFunc context.CancelFunc) {
	signalsChan := make(chan os.Signal, 1)
	signal.Notify(signalsChan, os.Interrupt, os.Kill)

	switch <-signalsChan {
	case os.Interrupt:
		cancelFunc()
	case os.Kill:
		cancelFunc()
	}
}
