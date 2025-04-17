package util

import (
	"fmt"

	"github.com/litsea/log-slog"
)

func RecoverFn(from string) {
	if err := recover(); err != nil {
		log.Error(from+": recovery form panic", "err", fmt.Errorf("panic error: %v", err))
	}
}
