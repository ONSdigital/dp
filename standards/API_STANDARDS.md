API standards
=============

## Core Principles

We use APIs to represent our data in machine-readable, predictable ways.

Our APIs follow [RESTful principles](https://restfulapi.net/) which are worth understanding more fully but highlights we're particularly focussed on are:
- Endpoints must return JSON responses and correct/meaningful HTTP status codes
- Endpoints must use appropriate requests methods (GET/PUT/POST etc). Changes must not be made via GET requests.
- Concrete path elements must be pluralised e.g. datasets not dataset in /datasets/<id>
- Endpoints must be documented using the Swagger/OpenAPI standard
- Updates to data must be idempotent

More details of precise behaviour we expect including guidance on field names, timestamps, pagination and responses are given below.

## Our API Landscape

Our [API training](../training/architecture/API.md) lays this out in more detail, but in order to allow each microservice API to focus on managing its data well some behaviour has been abstracted from the API services. This behaviour instead sits at the [API router](https://github.com/ONSdigital/dp-api-router).

Every request coming in on our api.ons.gov.uk domain must pass through our API router, which will perform a number of common checks and then determine which microservice to route traffic through to. When this API router is running internally, these checks include checking for relevant authentication headers and ensuring the request is audited for security reasons.

All API requests must travel through the API router, including app-to-app traffic. This is referred to as 'eating our own dog food' in our [Architecture Standards](ARCHITECTURE_STANDARDS.md). Not only does this ensure that all internal traffic is subject to the same security checks, it serves as quality assurance that our APIs are usable.

We use feature flags to limit what behaviour is available on our public facing API microservices. Typically, only GET endpoints are exposed on public APIs, though user transaction services such as filter and flex data journeys are an exception to this.

To add a new API microservice [follow our guide](../guides/NEW_API.md) to ensure the API router is updated correctly.

## API Structure and Syntax

### Syntax & types

  * Field names snake case, e.g. `total_count`
  * Use appropriate types for fields, e.g. numbers should be ints not strings
  * Timestamps should be UTC (optionally, with a timezone)
  * Root endpoint IDs should always be referred to as `id`, rather than name-spaced. All IDs under this level will need descriptive placeholder names. e.g. /datasets/<id>/editions/<edition> where /datasets is the root endpoint

### Methods & Behaviours
 * Use the `GET` method for search endpoints
 * `POST` endpoints should typically generate IDs, and these should use GUIDs not auto-incrementing counts

 ### Links
* All endpoints must return a `links.self` subdoc in their response,
* Resources should avoid a nested list by having a links to list endpoints
* Responses must contain fully qualified URLs, though this may be managed by the API router
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
          },
          "self":{
              "id":"1234",
              "url":"/thisapi/1234"
          }
      }
  }
  ```
* Resources may link to their direct parent, a child list, or any relevant
  external resources. Resources which are many layers deep in an API do not need
  to contain links to all higher elements e.g. the resource at
  `/datasets/1234/editions/2345/versions/1` would contain a link to
  `/datasets/1234/editions/2345` but not to `/datasets/1234`


### Returning errors
  * **Errors** should return a status code and a JSON payload with the error message
      * `errors` array containing error messages


### Returning lists or search results
Search/list endpoints should contain the following fields: `count`, `limit`, `offset`, `total_count`, `items`

The fields work as:

* `limit` - max number of items we're returning in this response.
    * If value is not set, a configurable default limit should be used, unless otherwise stated this should be set to a value of `20`.
    * The limit value can be set to `0` to return `0` items, so an API user can obtain the metadata for the list endpoint.

* `count` - how many items are actually present in the response.

* `total_count` - how many total items there may be (so the full list size, maybe thousands).

* `offset` - the number of documents into the full list that this particular response is starting at, this should default to `0` if not set. 
    * For example, in a list that has a total_count of 511, we might set a limit of 100, an offset of 500, and get a response whose count is 11, because it's the last 11 documents in the list.*
	  
* `items` - array containing results. 
    * Should only return items which match the offset and limit criteria, e.g. using the example above, you would expect the items array to only contain 11 documents.
	  
When pagination is required, it is also required the data be in a sorted order before it can be paginated. We should always configure defaults to help protect the service from performance problems due to allowing large limit and offset query values being set: `DEFAULT_MAXIMUM_LIMIT`, `DEFAULT_MAXIMUM_OFFSET`

* `DEFAULT_MAXIMUM_LIMIT` - Environment variable to cap the number of items to be returned.
	  * Should be set in the region of 500 to 1000 items. It is good practice to return a few results at a time, so requests don’t tie up resources for too long by trying to get all the requested data at once.
	  * Should return 400 (bad request) status code and an error explaining the maximum limit on limit value and providing the maximum value.

* `DEFAULT_MAXIMUM_OFFSET` - **Optional** environment variable to cap how far one can access a list of items.
	  * This is optional and is dependent on the underlying database. *Example an API relying on Elasticsearch will need a maximum offset value see why in the [example below](#elasticsearch-example-of-performance-issues-using-large-offsets).*
	  * Should return 400 (bad request) status code and an error explaining the maximum limit on offset value and providing the maximum value.


#### Elasticsearch example of performance issues using large offsets

An API accepting large offset values to paginate data returned from elasticsearch will result in slow responses. This is due to data being spread across multiple shards in an Elasticsearch cluster.

The processing in Elasticsearch requires it to sort the data on each shard based on the request (search query) and store the requested hits and hits for previous pages into memory; then accumulate all the information from each shard to work out the actual hits (search document) to return. This gets more resource intensive as one pages deeper and can result in slower responses between the requesing service (API) to Elasticsearch, [to read more on Elasticsearch pagination here](https://www.elastic.co/guide/en/elasticsearch/reference/current/paginate-search-results.html#paginate-search-results).


#### TODO

* Consider versioning / base path
* Consider licensing
* Resources (or rather, database items) should contain a `last_updated` field, which might not be exposed via the public interface
