Golang guidelines
=====================

These are the Digital Publishing guidelines for developing code using Golang. They can also be used as a code review
checklist.

#### General
* Start with the language level recommendations for Golang in [Effective Go](https://golang.org/doc/effective_go.html) 
and the [code review guidelines](https://github.com/golang/go/wiki/CodeReviewComments)
* [Listen for OS signals](https://golang.org/pkg/os/signal/) and gracefully shutdown the app when one is received

#### Tooling

* Use [gometalinter](https://github.com/alecthomas/gometalinter) to identify code issues
* Prefer using [GoConvey](https://github.com/smartystreets/goconvey) as a test tool
* Use [govendor](https://github.com/kardianos/govendor) for vendoring dependencies

#### Libraries

* Utilise the [go-ns](https://github.com/ONSdigital/go-ns) library for common functionality. Extract common code from projects and contribute it to go-ns
    * Use the common logging implementation
    * Default to the go-ns HTTP server implementation
        * By default the go-ns HTTP server will handle OS signals and gracefully shutdown. This can be disabled by setting `srv.HandleOSSignals = false`
        * If you don't, ensure your server utilises the common middleware handlers
            * log
            * requestID
        * Use the health check handler for all services. You may need to implement a more complex handler to check the health of database connections etc
* Read environment configuration using [gofigure](https://github.com/ian-kent/gofigure)



