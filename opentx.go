package opentelx

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/phansivang/opentx/middleware"
	"net/http"

	"github.com/phansivang/opentx/trace"
)

func SetupOpenTxSDK(openTxTarget, serviceName string) {
	trace.Setup(openTxTarget, serviceName)
}

func Shutdown(ctx context.Context) error {
	return trace.Shutdown(ctx)
}

func StartSpan(ctx context.Context, spanName string) (context.Context, func()) {
	return trace.StartSpan(ctx, spanName)
}

func GoSpanMiddleware(h http.HandlerFunc, spanName string) http.HandlerFunc {
	return middleware.GoSpan(h, spanName)
}

func GinSpanMiddleware(spanName string) gin.HandlerFunc {
	return middleware.GinSpan(spanName)
}
