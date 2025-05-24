package logger

import (
	"context"
	"log/slog"
)

type SLogger struct {
	ctx context.Context
	lg  *slog.Logger
}

func NewSLogger(
	ctx context.Context,
	lg *slog.Logger,
) *SLogger {
	return &SLogger{
		ctx: ctx,
		lg:  lg,
	}
}

func (lg *SLogger) Log(
	msg string,
	args ...any,
) {
	lg.Lg().Log(lg.Ctx(), slog.LevelInfo, msg, args...)
}

func (lg *SLogger) Debug(
	msg string,
	args ...any,
) {
	lg.Lg().Log(lg.Ctx(), slog.LevelDebug, msg, args...)
}

func (lg *SLogger) Warn(
	msg string,
	args ...any,
) {
	lg.Lg().Log(lg.Ctx(), slog.LevelWarn, msg, args...)
}

func (lg *SLogger) Error(
	msg string,
	args ...any,
) {
	lg.Lg().Log(lg.Ctx(), slog.LevelError, msg, args...)
}

func (lg *SLogger) Ctx() context.Context {
	return lg.ctx
}

func (lg *SLogger) Lg() *slog.Logger {
	return lg.lg
}
