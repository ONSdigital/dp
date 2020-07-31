API standards
=============

### Fields

  * Field names snake case, e.g. `total_count`
  * Use appropriate types for fields, e.g. numbers should be ints not strings
  * Resource IDs should use GUIDs not auto-incrementing counts
  * Timestamps should be UTC (optionally, with a timezone)
  * Use the `GET` method for search endpoints
  * Search/list endpoints should contain consistent fields:
    * Response: `count`, `limit`, `offset`, `total_count`
    * Request: `limit`, `offset`
    * `items` array containing results
    * Internally, set a maximum allowed `limit` to prevent requests above
        the limit. Should return an error explaining the limit and providing the
        maximum values.


### Links
* Resources should avoid a nested list by having a links to list endpoints
* Links should be fully qualified URLs
* `Links` object contains descriptively named objects which each contain
  an `id` (optional) and `url` field e.g.
  ```
  {  
      "id":"1234",
      "data":"i have stuff",
      "links": {  
          "latest_resource": {  
              "id":"5678",
              "url":"/thisapi/1234/subresources/5678"
          },
          "external": {  
              "id":"3456",
              "url":"/extresources/3456"
          },
          "subresources":{
              "url":"/thisapi/1234/subresources"
          }
      }
  }
  ```
* Resources may link to their direct parent, a child list, or any relevant
  external resources. Resources which are many layers deep in an API do not need
  to contain links to all higher elements e.g. the resource at
  `/datasets/1234/editions/2345/versions/1` would contain a link to
  `/datasets/1234/editions/2345` but not to `/datasets/1234`
  
### Spec
  * Placeholder names
      * Root endpoint IDs should always be referred to as `id`, rather than
        name-spaced. All IDs under this level will need descriptive placeholder
        names.
  * Concrete path elements should be plural words e.g. /**datasets**/{id} or /**animals**


### Responses
  * Use appropriate HTTP status codes
  * **Errors** should return a status code and a JSON payload with the error message
      * `errors` array containing error messages


#### TODO

* Consider versioning / base path
* Consider licensing
* Resources (or rather, database items) should contain a `last_updated` field, which might not be exposed via the public interface
