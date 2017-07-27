API standards
=============

* `GET` for search endpoints
* Field names snake case, e.g. `items_per_page`
* Search/list endpoints:
  * `count`, `items_per_page`, `start_index`, `total_count`
  * `items` array containing results
* Use correct HTTP status codes
* Resource IDs
    * should use GUIDs not auto-incrementing

# TODO

* Consider versioning / base path
* Resource IDs
    * naming - e.g. should ID on filter API be `filter_id` or just `id`?