/*
Package deferred provides helpers to close `io.Closer``.
This helpers can be used as one-liner with defer keyword to call `Close` method and check error if occurs.

example:
```go
    ...
    resp, err := http.Get("http://example.com/")
    if err != nil {
        // Handle error
    }

    defer deferred.CloseOrLog(resp.Body, logger)
    ...
```
*/
package deferred

import (
	"context"
	"io"
	"log"
)

type (
	logger interface {
		Errorf(template string, args ...interface{})
	}

	closerWithContext interface {
		Close(ctx context.Context) error
	}
)

// CloseOrLog tries to close `cl`. If error returned, log it.
// Supports loggers that have `Errorf` method. Like :
// 		zap
//		logrus
// 		apex/log
// 		etc...
func CloseOrLog(cl io.Closer, log logger) {
	err := cl.Close()
	if err != nil {
		log.Errorf("can't close: %v", err)
	}
}

// CloseOrLogCtx tries to close `cl`. If error returned, log it.
// Supports loggers that have `Errorf` method. Like :
// 		zap
//		logrus
// 		apex/log
// 		etc...
func CloseOrLogCtx(ctx context.Context, cl closerWithContext, log logger) {
	err := cl.Close(ctx)
	if err != nil {
		log.Errorf("can't close: %v", err)
	}
}

// CloseOrLogStd tries to close `cl`. If error returned, log it with standard `log`.
func CloseOrLogStd(cl io.Closer) {
	err := cl.Close()
	if err != nil {
		log.Printf("can't close: %v", err)
	}
}

// CloseOrLogCtx tries to close `cl`. If error returned, log it with standard `log`.
func CloseOrLogStdCtx(ctx context.Context, cl closerWithContext) {
	err := cl.Close(ctx)
	if err != nil {
		log.Printf("can't close: %v", err)
	}
}
