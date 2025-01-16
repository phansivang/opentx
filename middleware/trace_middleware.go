package middleware

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"log"
	"net/http"
)

func GoSpan(h http.HandlerFunc, spanName string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := otel.Tracer("opentx").Start(r.Context(), spanName)
		defer span.End()

		log.Println("Root span - Trace ID:", span.SpanContext().TraceID())

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GinSpan(spanName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, span := otel.Tracer("opentx").Start(c.Request.Context(), spanName)
		defer span.End()

		log.Println("Root span - Trace ID:", span.SpanContext().TraceID())

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
