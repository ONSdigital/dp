## Logging standards

**Does it add value?**
All apps should log important events, especially failures, and any information
which adds value is worth including in the event.

### Formatting

- JSON formatted
- Use consistently named fields
- Keys should be snake_cased and all lowercase

### Important fields

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
