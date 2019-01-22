Logging standards
=================

## Libraries

We have libraries for Go, Java and Javascript which match the logging standards defined below.

* Go - [github.com/ONSdigital/go-ns/log](https://github.com/ONSdigital/go-ns/tree/master/log)
* Java - [com.github.onsdigital.logging](https://github.com/ONSdigital/dp-logging)
* Javascript - [app/utilities/log.js](https://github.com/ONSdigital/florence/blob/develop/src/app/utilities/log.js)

## What we should be logging

All services should log important events, especially failures, and any information which adds value
is worth including in the event. If you're unsure, log more rather than less.

Always consider the performance impact of logging:

* loops
* log-specific conditionals ("should I log this")
* JSON serialisation
* deeply nested data structures

## Output formatting

- [JSON Lines](http://jsonlines.org/) - one JSON object per line
- Use UTF8
- Use consistently named fields
- Field names should be snake_cased and all lowercase

## Log data

We have defined the data structures we use for logging, including field names and types.

This is important to allow Elasticsearch to properly index the data and make it searchable.

### Common fields

| Field name | Type        | Example                 | Description
| ---------- | ----------- | ----------------------- | -----------
| created    | Date string | "2019-01-22T08:39:15Z"  | ISO8601 date string
| namespace  | String      | "dp-frontend-router"    | Service name or other identifier
| trace_id   | String      |                         | Trace ID from OpenCensus
| span_id    | String      |                         | Span ID from OpenCensus
| event      | String      | "Connecting to MongoDB" | The event being logged (see below for a full explanation)
| level      | String      | "info"                  | The log level (see below for a full explanation)
| http       | Object      | {}                      | HTTP event data (see below for a full explanation)
| data       | Object      | {}                      | Arbitrary key-value pairs (see below for a full explanation)

#### HTTP event data

HTTP events are so common that we've defined a specific top-level field to capture it. The `http` field also
defines it's own common fields.

This can be used to log inbound or outbound HTTP event data.

| Field name  | Type        | Example                 | Description
| ----------- | ----------- | ----------------------- | -----------
| status_code | Number      | 200                     | The HTTP status code
| start       | Date string | "2019-01-22T08:39:15Z"  | ISO8601 date string of the request start time
| end         | Date string | "2019-01-22T08:39:15Z"  | ISO8601 date string of the request end time
| duration    | Number      | 236578                  | Difference between `start` and `end` in nanoseconds
| method      | String      | "GET"                   | The HTTP method used
| path        | String      | "/healthcheck"          | The request path from the URL
| query       | String      | "x=1&y=1"               | The unparsed query string from the URL
| scheme      | String      | "https"                 | The scheme or protocol from the URL
| host        | String      | "localhost"             | The hostname from the URL
| port        | Number      | 443                     | The port number from the URL


- `created` - timestamps should be present in all log messages, and should be ISO-8601 formatted
- `namespace` - identifier for the service - should match the repo name
  - `service_instance` if running multiple instances of a service, an instance-specific identifier - perhaps an *allocation ID*
- `event` - the message explaining the error
  - should be kept consistent where possible, and should never be a formatted string
  - not sentences (no starting capital, no full-stop) e.g.
    - Good: `'event': 'file could not be read from s3'`
    - Bad: `'event': 'file [%s] could not be read from [%s]', $filename, $database'`
- `context` - an identifier to marry up all messages relating to an individual request across multiple services
  - can be multiple (comma-separated)
- `data` - event-specific details in a sub-document
  - key-value pairs for details like `'error':<error from called service>, 'dataset_id':<id>`

### Important events

- On startup, log config
- Callers should log the errors they receive
- HTTP Requests
    - On request
        - `url` (including query string)
        - `method`
    - Response failures, log:
        - Response (code)
        - Expected response

### Apps

Logical blocks should use logs as checkpoints
- `starting x process` is useful at the start of logical blocks, `finished x` may not be explicitly needed, if implied by the next `starting y` block
 - `finished x` may be useful where a process is time-consuming, to get a clear idea of duration.
- avoid back-to-back log statements unless they are related to different dataset which is relevant

##### Kafka

- Events should refer to an ID to tie the events across services together - whether this is the same ID as used in context, or another identifier, which has previously been logged alongside the context
- log `topic` and (if consuming) `group` and `offset`

### What not to log?

- log *levels* are not so valuable, default to `info` if library insists
- do not log queries (e.g. SQL strings) - log the individual parameters in a structured way
- never log passwords, credentials, auth headers
  - mask or take sub-strings allows for small portions of credentials more securely
- libraries should return (not log) errors
