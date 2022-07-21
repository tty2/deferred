deferred
---

## idea:

This is a tiny library that contains helpers to close any closer that returns error with `defer`.

The general idea is simple: make it with one line.

## example:

```go
    ...
    resp, err := http.Get("http://example.com/")
    if err != nil {
        // Handle error
    }

    defer deferred.CloseOrLog(resp.Body, logger)
    ...
```

## list of methods:

    // pass logger you use in your code
    - CloseOrLog(cl io.Closer, log logger)  
    - CloseOrLogCtx(ctx context.Context, cl io.Closer, log logger)

    // only if you don't mind using standard log
    - CloseOrLogStd(cl io.Closer)              
    - CloseOrLogStdCtx(ctx context.Context, cl io.Closer)