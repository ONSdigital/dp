Instrumenting ONS services for Open Telemetry
===============

# Instrumenting Java services for OT
These lines to be added in the Dockerfile(s), and to any run.sh scripts that are used locally:

`-javaagent:target/dependency/aws-opentelemetry-agent-1.31.0.jar \`

`-Dotel.propagators=tracecontext,baggage \`

For instance:
```
ENTRYPOINT java $JAVA_OPTS \
-Drestolino.realm=$REALM \
-Drestolino.files=$RESTOLINO_STATIC \
-Drestolino.classes=$RESTOLINO_CLASSES \
-Drestolino.packageprefix=$PACKAGE_PREFIX \
-javaagent:target/dependency/aws-opentelemetry-agent-1.31.0.jar \
-Dotel.propagators=tracecontext,baggage \
-cp "target/dependency/*:target/classes/" \
com.github.davidcarboni.restolino.Main
```


The following entry needs to be added to the pom:
```
<!-- OpenTelemetry-->
<dependency>
    <groupId>software.amazon.opentelemetry</groupId>
    <artifactId>aws-opentelemetry-agent</artifactId>
    <version>1.31.0</version>
</dependency>
```


The following environment variables need to be set on the instance:

`OTEL_SERVICE_NAME=<service name>`

`OTEL_EXPORTER_OTLP_ENDPOINT=http://$NOMAD_IP_http:4317` (http at present, moving to https as TLS is enabled)**

The URL specified has an identifier of http. In fact the protocol in use is GRPC (the collector is only listening
on GRPC). However, the java URL parsing lib doesn't recognise 'grpc://' as a valid protocol, so the configuration
requires it to be specified as above.


Spans can be created around individual calls within the code as follows:
```
import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.context.Scope;

...

Tracer tracer = GlobalOpenTelemetry.getTracer("<package name>", "<package version>");

Span span = tracer.spanBuilder("<name of your span>").startSpan();

try (Scope scope = span.makeCurrent()) {
    ...
    ...
  } catch(Throwable t) {
    span.recordException(t);
    throw t;
  } finally {
 span.end();
}
```


(NB - using the tracer pulled from the global scope this way works, but it's preferable to instantiate the tracer in the init code for the service then inject it into your method. See here for full guidance: [Open Telemetry Docs](https://opentelemetry.io/docs/instrumentation/java/manual/#:~:text=To%20create%20Spans%2C%20you%20only,set%20by%20the%20OpenTelemetry%20SDK.&text=It's%20required%20to%20call%20end,you%20want%20it%20to%20end)).


## Logging Implementation

The existing logging library has been modified to extract the traceId from the traceparent header (if it exists) and add it to the TraceId section of the log message. This will enable log entries to be correlated with trace ids which will allow engineers to zero in on problems quickly and accurately.


The original logging library was modified in order to manage the change centrally and avoid the need for code changes across multiple applications.


# Instrumenting Go services for OT
The following environment variables need to be created:
```
OTServiceName              string        `envconfig:"OTEL_SERVICE_NAME"`
OTExporterOTLPEndpoint     string        `envconfig:"OTEL_EXPORTER_OTLP_ENDPOINT"`
OTBatchTimeout             time.Duration `envconfig:"OTEL_BATCH_TIMEOUT"`
```
These can then be set in the config:
```
cfg = &Config{
    OTExporterOTLPEndpoint:     "localhost:4317",
    OTServiceName:              "service-name",
    OTBatchTimeout:              5 * time.Second,
}
```
Note that the exporter endpoint is `<hostname>:<port>`, unlike the java configuration there is no protocol identifier

Import the shared init library for go dp-otel-go

`import "github.com/ONSdigital/dp-otel-go"`

From the init code of the library initialise the otel services:
```
//Set up OpenTelemetry
cfg, err := config.Get()

otelConfig := dpotelgo.Config{
    OtelServiceName:          cfg.OTServiceName,
    OtelExporterOtlpEndpoint: cfg.OTExporterOTLPEndpoint,
    OtelBatchTimeout:         cfg.OTBatchTimeout,
}

otelShutdown, oErr := dpotelgo.SetupOTelSDK(ctx, otelConfig)

if oErr != nil {
    log.Fatal(ctx, "error setting up OpenTelemetry - hint: ensure OTEL_EXPORTER_OTLP_ENDPOINT is set", oErr)
    return
}
// Handle shutdown properly so nothing leaks.
defer func() {
    err = errors.Join(err, otelShutdown(context.Background()))
}()
```
NB: if this isn't done any calls to the otel service will fail silently. If you find that traces are not coming through, ensure this code is getting called.


## Instrumenting http handlers
There are a wide range of different facilities for instrumenting http calls. The simplest (taken here from dp-search-api) simply creates a new opentelemetry handler to pass to the server and attaches otelmux middlewarer to the router:
```
import "go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
import "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
...
router := mux.NewRouter()
mux.Use(otelmux.Middleware(cfg.OTServiceName))
otelHandler := otelhttp.NewHandler(router, "/")
router.Use(otelmux.Middleware(cfg.OTServiceName))

server := serviceList.GetHTTPServer(cfg.BindAddr, otelHandler)
```


Where gorillamux or Chi is not being used for the router, it may be necessary to instrument individual routes as follows:

```
func routes(router *mux.Router, hc *healthcheck.HealthCheck) *RendererAPI {
    api := RendererAPI{router: router}

    handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
        // Configure the "http.route" for the HTTP instrumentation.
        handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
        api.router.Handle(pattern, handler)
    }

    api.router.StrictSlash(true).Path("/health").HandlerFunc(hc.Handler)
    handleFunc("/render/{render_type}", api.renderTable)
    handleFunc("/parse/html", api.parseHTML)
    return &api
}
```

Where the server is configured with an api field:
```
return &Service{
	api:                 searchAPI,
}
```
Handlers can be wrapped as below: 
```
func (a *SearchAPI) RegisterGetSearch(...) *SearchAPI {
	a.Router.Handle(
		"/search",
		otelhttp.NewHandler(
			SearchHandlerFunc(
                ...
			), "/search"),
	).Methods(http.MethodGet)
	return a
}
```


The following shows an alternative way to instrument:
```
func CreateRendererAPI(ctx context.Context, bindAddr string, allowedOrigins string, errorChan chan error, hc *healthcheck.HealthCheck) {
    router := mux.NewRouter()
    routes(router, hc)
    otelhandler := otelhttp.NewHandler(router, "/")

    httpServer := dphttp.NewServer(bindAddr, otelhandler)
    // Disable this here to allow main to manage graceful shutdown of the entire app.
    httpServer.HandleSignals = false

    go func() {
        log.Info(ctx, "starting table renderer")
        if err := httpServer.ListenAndServe(); err != nil {
            errorChan <- err
            log.Error(ctx, "error occurred when running ListenAndServe", err)
        }
    }()
}
```


A purely middleware approach can also be taken where Alice is already in place chaining middleware. Both otelmux and otelhttp are used here to capture all requests with sufficient detail. Here you can see an example instrumentation:
```
func New(cfg Config) http.Handler {
    router := mux.NewRouter()
	router.Use(otelmux.Middleware(cfg.OTServiceName))
	middleware := []alice.Constructor{
		otelhttp.NewMiddleware(cfg.OTServiceName),
        ...
	}
	newAlice := alice.New(middleware...).Then(router)
}
```

## Instrumenting http calls
Outgoing service calls need to be instrumented to include the traceparent header when the handler itself is not instrumented. This can be done as follows:
```
import ("go.opentelemetry.io/otel"
"go.opentelemetry.io/otel/propagation")
...

otel.GetTextMapPropagator().Inject(req.Context(), propagation.HeaderCarrier(req.Header))
```


## Manually adding spans:

Similarly to the Java approach, you can create a span manually as follows:
```
import ("go.opentelemetry.io/otel")
...
...
tracer := otel.GetTracerProvider().Tracer("tablerenderer")
ctx, span := tracer.Start(r.Context(), "table render span")

defer span.End()
```

## Mongo Instrumentation:
The `dp-mongodb` package has been instrumented centrally as of version 3.7.0. This means that there is no need for additional instrumentation in services that import this package version or above. Make sure this version or above is imported!


## Kafka Instrumentation:
The `dp-kafka` package has also be instrumented centerally as of version 4, this requires the context to be passed in order to work:
```
kafkaProducer.Channels().Output <- kafka.BytesMessage{Value: bytes, Context: ctx}
```


## Logging Implementation

Go http request middleware was created to extract the traceId from the traceparent header and insert into the expected place in the request context (as controlled by the RequestIdKey in the github.com/ONSdigital/dp-net/v2/request package)

This will enable log entries to be correlated with trace ids, to zero in on problems quickly and accurately.

The original logging library was modified in order to manage the change centrally and avoid the need for code changes across multiple applications. Update to >= v2.4.3 to include this functionality in your service.
