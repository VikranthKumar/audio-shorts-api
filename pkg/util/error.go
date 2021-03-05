package util

import (
	"context"
	"github.com/nooble/task/audio-short-api/pkg/logging"
)

// for fatal error on server initialization
func ExitOnErr(ctx context.Context, err error) {
	if err != nil {
		logging.WithContext(ctx).Fatal(err.Error())
	}
}
