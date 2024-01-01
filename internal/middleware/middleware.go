package middleware

import "github.com/labstack/echo/v4"

func WithMiddleware(h echo.HandlerFunc, middleware echo.MiddlewareFunc) echo.HandlerFunc {
	return middleware(h)
}
