API standards
=============

* `GET` for search endpoints

* Field names snake case, e.g. `items_per_page`

* Search/list endpoints should contain consistent fields:
  * Output: `count`, `items_per_page`, `start_index`, `total_count`
  * Input: `items_per_page`, `start_index`
  * `items` array containing results
  * Internally, set a maximum allowed `items_per_page` to prevent requests above
    the limit. Should return an error explaining the limit and providing the
    maximum values.

* Use appropriate HTTP status codes

* Resource IDs
    * Should use GUIDs not auto-incrementing counts

* Placeholder names
    * Root endpoint IDs should always be referred to as `id`, rather than
      name-spaced. All IDs under this level will need descriptive placeholder
      names.

* Field types
    * Use appropriate types for fields - numbers should be ints not strings
    * Timestamps should be UTC (optionally, with a timezone)

* Links
    * Resource should avoid a nested list by having a Links
    * Links should be fully qualified URLs
    * `Links` object contains descriptively named objects which each contain
      an `id` and `link` field. eg
      {
        id: thing1
        data: "i have stuff"
        links: {
          "descriptive 1" : {
            id: "otherthing17"
            link: "/otherapi/otherthing17"
          }
        }
      }

* Errors should return a status code and a JSON payload with the error message
    * `errors` array containing error messages

# TODO

* Consider versioning / base path
* What is time?!
* Resources (or rather, database items) should contain a `last_updated` field, which may not be exposed via the public interface
