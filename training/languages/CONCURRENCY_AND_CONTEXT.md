Concurrency in Go (with `context`)
=================
Go provides concurrency through light weight `go routines`.

To effectively manage `go routines`, go has the `context` package.

### Concurrency Tutorials
Links that give a good introduction:

[Concurrency With Golang Goroutines](https://tutorialedge.net/golang/concurrency-with-golang-goroutines/)

[Golang Concurrency](https://www.golangprograms.com/go-language/concurrency.html)

### Concurrency - Best practices
Safe Channel usage:

[ assume the channel element type is T ]:

```go
    func SafeClose(ch chan T) (justClosed bool) {
    	defer func() {
    		  if recover() != nil {
    			// The return result can be altered
    			// in a defer function call.
    			justClosed = false
    		}
    	}()

    	// assume ch != nil here.
    	close(ch)   // panic if ch is closed
    	return true // <=> justClosed = true; return
    }

    func SafeSend(ch chan T, value T) (closed bool) {
    	defer func() {
    		if recover() != nil {
    			closed = true
    		}
    	}()

    	ch <- value  // panic if ch is closed
    	return false // <=> closed = false; return
    }
```

### Context Tutorials
Links that cover `context` well from slightly different angles:

[Getting Started with Go Context](https://dev.to/gopher/getting-started-with-go-context-l7g)

[Chapter 37: Context](https://www.practical-go-lessons.com/chap-37-context)

### Context - Best practices
`context.Background` should be used only at the highest level, as the root of all derived contexts.

On this page, see the section: Best practices

[Understanding the context package in golang](http://p.agnihotry.com/post/understanding_the_context_package_in_golang/)

Further resources
----------------------------
Books that cover the above topics in order of difficulty:

*Head First Go*

*The Go Programming Language*

*Concurrency in Go*

## Concurrency example
Concurrent file read ... using `errgroup.WithContext`

(see this article: [Why you should be using errgroup.WithContext() in your Golang server handlers](https://www.fullstory.com/blog/why-errgroup-withcontext-in-golang-server-handlers/)) for explanation.

```go
    package main

    import (
        "context"
        "fmt"
        "io/ioutil"
        "os"

        "golang.org/x/sync/errgroup"
    )

    func readFiles(ctx context.Context, files []string) ([]string, error) {
        g, ctx := errgroup.WithContext(ctx)

        results := make([]string, len(files))

        for i, file := range files {
            i, file := i, file
            g.Go(func() error {
                data, err := ioutil.ReadFile(file)
                if err == nil {
                    results[i] = string(data)
                }
                return err
            })
        }

        if err := g.Wait(); err != nil {
            return nil, err
        }
        return results, nil
    }

    func main() {
        var files = []string{
            "file1",
            "file2",
        }

        results, err := readFiles(context.Background(), files)
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            return
        }
        for _, result := range results {
            fmt.Println(result)
        }
    }
```

## Pseudo code examples of context

```go
    example 1:
    
        func main() {
            ctx := context.Background() // this is the root
            ctx, cancel := context.WithTimeout(ctx, time.Second) // add a child to the tree (of possible contex's)
            defer cancel() // Even though the following line will be canceled with the above timeout
                           // the cancel() function must still be called to release the allocated timer(s)
                           // that could be added further down any additional function call depth / path.
                           // from: justforfunc #9: The Context Package
                           // and in: justforfunc #10: implementing the context package (~34 mins in)
                           // the WithcDeadline() function calls t.stop() to release the resources
                           // allocated to the timer.
                           // The 'real' context.go file has the stop timer in the timer's cancel() function.
                           
            sleepAndTalk(ctx, 5*time.Second, "hello")
        }
        
        ... when run, after 1 second this produces: "context canceled"
            
    example 2:
    
        func main() {
            ctx := context.Background()
            ctx, cancel := context.WithTimeout(ctx, time.Second)
            defer cancel()
                           
            mySleepAndTalk(ctx, 5*time.Second, "hello")
        }

        func mySleepAndTalk(ctx context.Context, d time.Duration, msg string) {
            select {
            case <-time.After(d):
                fmt.Println(msg)
            case <-ctx.Done():
                log.Print(ctx.Err())
            }
        }
        
        ... when run, after 1 second this produces: "context deadline exceeded"

    example 3:
        
        Server:
        =======
        
        package main
        
        import (
            "flag"
            "fmt"
            "log"
            "net/http"
            "time"
        )
        
        func main() {
            http.HandleFunc("/", Handler)
            log.Fatal((http.ListenAndServe("127.0.0.1:8080", nil)))
        }
        
        func Handler(w http.ResponseWriter, r *http.Request) {
            ctx := r.Context()
            log.Printf("handler started")
            defer log.Printf("Handler ended")
            
            select {
            case <-time.After(5 * time.Second):
                fmt.Fprintln(w, "hello")
            case <-ctx.Done():
                err := ctx.Err()
                log.Print(err)
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
        }
        
        Client (simple, no context):
        ============================
        
        package main
        
        import (
            "io"
            "log"
            "net/http"
            "os
        }
        
        func main() {
            res, err := httpGet("http://localhost:8080")
            if err != nil {
                log.Fatal(err)
            }
            defer res.Body.Close()
            if res.StatusCode != http.StatusOK {
                log.Fatal(res.Status)
            }
            io.Copy(os.Stdout, res.Body)
        }
        
        ==>>
            now run the server, and then in another terminal send a request with:
                curl localhost:8080
                
            (and wait for full 5 seconds) the server terminal will show:
            
                handler started
                handler ended
                
            *** Then, doing the request again and after 2 seconds cancelling it with CTRL-C,
                the server terminal will now show:
                
                handler started
                context canceled
                handler ended
                
            thus the client cancelled the request and the HTTP server got the "context" cancel and
            passed it on to what might be a long running process - which is looking "ctx.Done" in a "select {}"
            and it can then stop what its doing as it now knows that the client will not receive the response.
            
        Client (with context):
        ======================
        
        package main
        
        import (
            "io"
            "log"
            "net/http"
            "os
        }
        
        func main() {
            ctx := context.Background()
            ctx, cancel := context.WithTimeout(ctx, time.Second)
            defer cancel()
            
            req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
            if err != nil {
                log.Fatal(err)
            }
            req = req.WithContext(ctx)    // create a new request that has a 'context' inside it
                    
            res, err := http.DefaultClient.Do(req)
            if err != nil {
                log.Fatal(err)
            }
            defer res.Body.Close()
            if res.StatusCode != http.StatusOK {
                log.Fatal(res.Status)
            }
            io.Copy(os.Stdout, res.Body)
        }


        ==>> have the server running and running the above client, it outputs:
        
            Get http://localhost:8080: context deadline exceeded
            
        and the server also outputs:
        
                handler started
                context canceled
                handler ended

        ==>> All of this allows us to cancel an operation as soon as they are not needed anymore.
            This all being done by adding extra code to achive whats needed and very often will
            require the adding of channels.
            
            
    example 4:
        
        Passing values in context (not a good idea, but if you have to):
        Whatever is in the context should be request specific.
        Whatever is in the context should add extra information that is useful but does not impact
        the flow of your program.
        Do NOT pass much as it makes code hard to follow.
        Parameters should not be passed, and idealy belong elsewhere - in flags for instance.

        An example of one thing that is typically passed as a value is the "request ID" ...
        
        Log (with context):
        ===================
        
        package log
        
        import (
            "context"
            "log"
            "math/rand"
            "net/http"
        )
        
        type key int
        
        const requestIDKey = key(42)  // only the log package uses this 'type' which is not exported
        // which stops extrnal 'same' value of 42 being used / clashing with log package in a line like this:
        //    ctx.context.WithValue(ctx, int(42), int64(100)) 
        
        func Println(ctx context.Context, msg string) {
            id, ok := ctx.Value(requestIDKey).(int64)
            if !ok {
                log.Println("could not find request ID in context")
                return
            }
            log.Printf("[%d] %s", id, msg)
        }
        
        // now we create a HTTP decorator / handler func
        func Decorate(f http.HandlerFunc) http.HandlerFunc {
            return func (w.http.ResponseWriter, r*http.Request) {
                ctx := r.Context()  // we receive a context
                id := rand.Int63()
                ctx = context.WithValue(ctx, requestIDKey, id)  // we add a value to the context
                f(w, r.WithContext(ctx))  // and we send the context back
            }
        }
        
        Server (new that uses our log):
        ===============================
        
        package main
        
        import (
            "flag"
            "fmt"
            "net/http"
            "time"
            
            "github.com/campoy/justforfunc/context/log"  *** where the above "log" package has been placed
        )
        
        func main() {
            http.HandleFunc("/", log.Decorate(Handler))
            log.Fatal((http.ListenAndServe("127.0.0.1:8080", nil)))
        }
        
        func Handler(w http.ResponseWriter, r *http.Request) {
            ctx := r.Context()
            log.Println(ctx, "handler started")
            defer log.Println(ctx, "Handler ended")
            
            select {
            case <-time.After(5 * time.Second):
                fmt.Fprintln(w, "hello")
            case <-ctx.Done():
                err := ctx.Err()
                log.Println(ctx, err.Error())
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
        }

        ==>> now running the server and doing a couple of requests to it, each request has a different
             request ID:
             
             2017/03/30 16:56:01 [5577006791947779411] handler started
             2017/03/30 16:56:06 [5577006791947779411] handler ended
             2017/03/30 16:56:09 [8674665230821535522] handler started
             2017/03/30 16:56:14 [8674665230821535522] handler ended
```
