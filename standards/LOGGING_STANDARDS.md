Logging standards
=================

Libraries
---------

We have libraries the following libaries which match the logging standards defined below:

* Go - [github.com/ONSdigital/log.go](https://github.com/ONSdigital/log.go)
* Java - [com.github.onsdigital.logging](https://github.com/ONSdigital/dp-logging)
* Javascript - [app/utilities/log.js](https://github.com/ONSdigital/florence/blob/develop/src/app/utilities/log.js)

What we should be logging
-------------------------

Apps should log all important events, especially failures, as well as enough events to be able to trace the execution logic through each major step without being overly noisy.  Any information which adds value is worth including in the event. **If you're unsure, log more rather than less.**

A few ideas of useful things to log are:

* application startup (with sanitised application config output under `data`)
* log all failures
* http requests
* requests to data stores and other third party services (e.g. mongo, postgres, kafka, neptune, etc.)
* log events at significant steps throughout the execution flow
  * sometimes it is useful to have `started x` and `finished x` for very long running events, but usually a single event can suffice
  * avoid back to back log statements as a general rule, especially if the second event implies the first

Always consider the performance impact of logging:

* loops
* log-specific conditionals ("should I log this")
* JSON serialisation
* deeply nested data structures

Sensitive information
---------------------

**NEVER** log passwords, credentials, auth headers, or other sensitive information in any form.

Logging specification
---------------------

We have defined the formatting and data structures, including field names and types, we use for logging.

This is important to allow the centralised logging service to properly index the data and make it searchable.  **Any logs not adhering to the specification may be lost.**

### Changes to the spec

Any changes, including the addition of new fields, must be agreed by the team as a whole before being added to this spec.  Any changes to this spec will require changes to all of the [logging libraries](#libraries) as well as the centralised logging service.

### Output formatting

Although "human readable" formatting is useful for local development, all apps deployed to an environment must comply to the following formatting rules.

* All logs **must** be [JSON Lines](http://jsonlines.org/)
* All logs **must** be UTF8 encoded
* All logs **must not** be colour formatted
* All logs **must not** contain newlines
* All logs **must not** contain tabs
* All logs **must not** be string escaped

### Log structure

#### Common fields

The following are the only top level fields. If you have app specific details you wish to add, it is likely that they should be added under `data` rather than creating a new field.

| Field name   | Required       | Type                           | Example                      | Description                                                             |
| ------------ | -------------- | ------------------------------ | ---------------------------- | ----------------------------------------------------------------------- |
| `created_at` | Yes            | `datetime`<sup>1</sup>         | `"2019-01-21T16:19:12.356Z"` | Date and time of the log event (ISO8601 extended date and time string)  |
| `namespace`  | Yes            | `string`                       | `"dp-frontend-router"`       | Service name or other identifier                                        |
| `event`      | Yes            | `string`                       | `"connecting to mongodb"`    | [The event being logged](#effective-event-types)                        |
| `trace_id`   | No<sup>2</sup> | `string`                       |                              | [Trace ID from OpenCensus](https://opencensus.io/tracing/span/traceid/) |
| `span_id`    | No             | `string`                       |                              | [Span ID from OpenCensus](https://opencensus.io/tracing/span/spanid/)   |
| `severity`   | [v2](#v2) - Yes            | `int8`                         | `2`                          | [The event severity level code](#severity-levels)                       |
| `http`       | No             | [`http`](#http-event-data)     |                              | [HTTP event data](#http-event-data)                                     |
| `auth`       | No             | [`auth`](#auth-event-data)     |                              | [Authentication event data](#auth-event-data)                           |
| `errors`     | No             | [`[]error`](#error-event-data) |                              | [Error event data](#error-event-data)                                   |
| `raw`        | No             | `string`                       |                              | [Log message captured from a third party library](#third-party-logs)    |
| `data`       | No             | `object`<sup>3</sup>           |                              | [Arbitrary key-value pairs](#arbitrary-data-fields)                     |

<sup>1</sup> All dates must be UTC and in ISO8601 extended date time format with at least millisecond precision: `yyyy-MM-dd'T'HH:mm:ss.SSSZZ` (e.g. `2019-01-21T16:19:12.356Z`).

<sup>2</sup> Not mandatory, but must be included for all events created during the handling of a request.

<sup>3</sup> The event-specific details in `data` are input as an object/map in code, but are stored as a text string by the centralised logging service to allow additional detail to be added to an event without needing to worry about key name collision.  These details can still be searched using a general text search on the `data` field.

#### HTTP event data

HTTP events are so common that we've defined a specific top-level field to capture it. The `http` field defines the common http event data.

This can be used to log inbound or outbound HTTP event data.

| Field name                | Required | Type                   | Example                      | Description                                                   |
| ------------------------- | -------- | ---------------------- | ---------------------------- | ------------------------------------------------------------- |
| `method`                  | Yes      | `string`               | `"GET"`                      | The HTTP method used                                          |
| `scheme`                  | Yes      | `string`               | `"https"`                    | The scheme or protocol from the URL                           |
| `host`                    | Yes      | `string`               | `"10.100.20.1"`              | The hostname or ip from the URL                               |
| `port`                    | Yes      | `int32`                | `443`                        | The port number from the URL                                  |
| `path`                    | Yes      | `string`               | `"/healthcheck"`             | The request path from the URL                                 |
| `query`                   | No       | `string`               | `"x=1&y=1"`                  | The unparsed query string from the URL                        |
| `status_code`             | No       | `int16`                | `200`                        | The HTTP status code                                          |
| `started_at`              | Yes      | `datetime`<sup>1</sup> | `"2019-01-21T16:19:12.356Z"` | ISO8601 date string of the request start time                 |
| `ended_at`                | No       | `datetime`<sup>1</sup> | `"2019-01-21T16:19:12.356Z"` | ISO8601 date string of the request end time                   |
| `duration`                | No       | `int64`                | `236578`                     | Difference between `started_at` and `ended_at` in nanoseconds |
| `response_content_length` | No       | `int64`                | `34`                         | The length of the response                                    |

<sup>1</sup> All dates must be UTC and in ISO8601 extended date time format with at least millisecond precision: `yyyy-MM-dd'T'HH:mm:ss.SSSZZ` (e.g. `2019-01-21T16:19:12.356Z`).

#### Auth event data

The `auth` field defines the common auth event data.

| Field name      | Required | Type     | Example               | Description                                  |
| --------------- | -------- | -------- | --------------------- | -------------------------------------------- |
| `identity`      | Yes      | `string` | `"dp-import-tracker"` | The user id or service id                    |
| `identity_type` | Yes      | `string` | `"service"`           | The identity type (i.e. `user` or `service`) |

#### Error event data

Errors are logged as an array. Each element of the array represents a single error and has the following fields.

| Field name    | Required | Type                                    | Example                | Description                                         |
| ------------- | -------- | --------------------------------------- | ---------------------- | --------------------------------------------------- |
| `message`     | Yes      | `string`                                | `"connection refused"` | The error cause                                     |
| `stack_trace` | No       | [`[]stack_trace`](#stack-trace-element) |                        | [Stack trace as an array](#stack-trace-element)     |
| `data`        | No       | `object`<sup>1</sup>                    |                        | [Arbitrary key-value pairs](#arbitrary-data-fields) |

<sup>1</sup> The error-specific details in `data` are input as an object/map in code, but are stored as a text string by the centralised logging service to allow additional detail to be added to an event without needing to worry about key name collision.  These details can still be searched using a general text search on the `data` field.

##### Stack trace element

The stack trace is logged as an array.  Each element of the array has the following fields.

| Field name | Required | Type     | Example                | Description                                            |
| ---------- | -------- | -------- | ---------------------- | ------------------------------------------------------ |
| `file`     | No       | `string` | `"/some/path/main.go"` | The source file containing the code throwing an error  |
| `function` | No       | `string` | `"main.main"`          | The function in which the error is occurring           |
| `line`     | No       | `int16`  | `18`                   | The line in the source file that is throwing the error |

### Effective `event` types

The `event` field is a brief description of the type of event.

A good `event` description is:

* lowercase
* **not** a sentence (i.e. is brief and without a full stop)
* generic (e.g. `http request`)
* kept consistent where possible
* **not** a formatted string (additional information should be provided via other fields)

For example:

```text
'event': 'file could not be read from s3'
```

These are guidelines only - there are some exceptions where it's ok to ignore them, for example when
using uppercase to refer to a constant or environment variable:

```text
'event': 'DEBUG is not a boolean'
```

### Severity levels

The event severity levels are used to identify how critical an event is, especially failure events.

There are four severity levels:

| Code | Severity | Description                                                                             |
| ---- | -------- | --------------------------------------------------------------------------------------- |
| 0    | FATAL    | An catastrophic failure event resulting in the death of the application                 |
| 1    | ERROR    | An unrecoverable failure event that halts the current flow of execution                 |
| 2    | WARN     | A failure event that can be managed within the current flow of execution (e.g. retried) |
| 3    | INFO     | All non-failure events                                                                  |

For example, if in the process of serving an HTTP request an action is retried three times, the first two failures would have a severity of `WARN`, whereas the third time would have an `ERROR`.

### Third party logs

The [log libraries](#libraries) attempt to intercept all logs from third party libraries in order to prevent collision with our logging specification.  These intercepted log messages should be output as a string under the `raw` field with the `event` as `"third party log"` and `severity` of `3` (INFO).  All other relevant fields (e.g. `created_at`, `namespace`, `trace_id`, `span_id`, etc) should be set as they would normally.

For example:

```json
{
  "created_at" : "2019-02-01T13:45:24.157Z",
  "namespace" : "my-app",
  "trace_id": "1105cb0c04f86a4b6a1abaf74246b87f",
  "severity" : 3,
  "event": "third party log",
  "raw" : "Started ServerConnector@7f4fedd{HTTP/1.1,[http/1.1]}{0.0.0.0:4567}"
}
```

### Arbitrary data fields

There are two arbitrary `data` fields in the logging spec.  One at the top level for general event specific data and one under `error` for error specific data.  These fields are a map/object in code, but will be stored as a string field by the centralised logging service to allow additional detail to be added to an event without needing to worry about key name collision.  These details can still be searched using a general text search on the `data` field.

The `data` fields are a means of adding event specific details that help in understanding the event, but are not common enough to justify a field in the spec. These fields are an important part of the event data and it is important that the data be useful and well structured.

### v2

Following creation of the v2 library in [log.go](https://github.com/ONSdigital/log.go), `severity` is now a mandatory field when creating log events. In addition to this, it is not possible to pass `severity` as an optional field, therefore removing the possibility of duplicate `severity` logs.

Four additional functions have been added to the log library: `Info`, `Warn`, `Error` and `Fatal`. These functions wrap the existing `Event` function and have been named to correspond with their severity level (shown [here](#severity-levels)). So, for example, the `Warn` function corresponds with the WARN severity. These functions can be used to provide information about a specific severity level rather than using the more generic `Event` function. 

As we already had a public `Error` function that is used in our apps to display error messages as an option in the logs, the original `Error` function has now been renamed to `FormatError`, which also ensures that the wrapper function names (`Info`, `Warn`, `Error` & `Fatal`) are consistent with their error level. The functionality of `FormatError` remains unchanged.

The `Error` and `Fatal` functions also require that an error is passed as an argument. An example of this for the `Fatal` function would be: 

```go
log.Fatal(ctx, "failed to shutdown http server", errors.New("testfatal"))
``` 

`FormatError` is called within the `Error` and `Fatal` functions, which ensures that the `StackTrace` information is also shown in the log.


##### Kafka

TODO: Add kafka log data to the logging spec

- Events should refer to an ID to tie the events across services together - whether this is the same ID as used in context, or another identifier, which has previously been logged alongside the context
- log `topic` and (if consuming) `group` and `offset`
