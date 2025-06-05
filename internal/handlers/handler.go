package handlers

import (
	"context"

	"github.com/vxdiazdel/rest-api/internal/db/stores"
	"github.com/vxdiazdel/rest-api/internal/logger"
)

type HandlerContext struct {
	ctx   context.Context
	store stores.IStore
	lg    logger.ILogger
}

func NewHandlerContext(
	ctx context.Context,
	store stores.IStore,
	lg logger.ILogger,
) *HandlerContext {
	return &HandlerContext{
		ctx:   ctx,
		store: store,
		lg:    lg,
	}
}

func (h *HandlerContext) Ctx() context.Context {
	return h.ctx
}

func (h *HandlerContext) Store() stores.IStore {
	return h.store
}

func (h *HandlerContext) Lg() logger.ILogger {
	return h.lg
}
