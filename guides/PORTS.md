Port allocations
================

These are the port numbers used by DP services.

Using distinct port numbers for each service means we can run everything
in development without having to configure them individually.

### Default port numbers

| Port  | Service | Notes
| ----- | ------- | -----
| 2222  | [Dylan](https://github.com/ONSdigital/Dylan) : ssh |
| 5432  | [dp-compose](https://github.com/ONSdigital/dp-compose) : postgres |
| 8080  | [babbage](https://github.com/ONSdigital/babbage) |
| 8081  | [florence](https://github.com/ONSdigital/florence) |
| 8082  | [zebedee](https://github.com/ONSdigital/zebedee) |
| 8083  | [project-brian](https://github.com/ONSdigital/project-brian) |
| 8084  | [The-Train](https://github.com/ONSdigital/The-Train) |
| 8085  | [Dylan](https://github.com/ONSdigital/Dylan) : http |
| 8086  | [ermintrude](https://github.com/ONSdigital/ermintrude) |
| 9000  | [sixteens](https://github.com/ONSdigital/sixteens) |
| 9200  | [dp-compose](https://github.com/ONSdigital/dp-compose) : elasticsearch |
| 9300  | [dp-compose](https://github.com/ONSdigital/dp-compose) : elasticsearch |
| 9999  | [dp-compose](https://github.com/ONSdigital/dp-compose) : highcharts |
| 20000 | [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router) |
| 20001 | [dp-frontend-filter-dataset-controller](https://github.com/ONSdigital/dp-frontend-filter-dataset-controller)
| 20010 | [dp-frontend-renderer](https://github.com/ONSdigital/dp-frontend-renderer) |
| 20020 | [dp-content-resolver](https://github.com/ONSdigital/dp-content-resolver) |
| 20200 | [dp-frontend-dataset-controller](https://github.com/ONSdigital/dp-frontend-dataset-controller)
| 21000 | [dp-csv-splitter](https://github.com/ONSdigital/dp-csv-splitter) |
| 21300 | [dp-import-tracker](https://github.com/ONSdigital/dp-import-tracker) |
| 21400 | [dp-dimension-extractor](https://github.com/ONSdigital/dp-dimension-extractor) |
| 21500 | [dp-dimension-importer](https://github.com/ONSdigital/dp-dimension-importer) |
| 21600 | [dp-observation-extractor](https://github.com/ONSdigital/dp-observation-extractor) |
| 21700 | [dp-observation-importer](https://github.com/ONSdigital/dp-observation-importer) |
| 21800 | [dp-import-api](https://github.com/ONSdigital/dp-import-api) |
| 21900 | [dp-code-list-api](https://github.com/ONSdigital/dp-code-list-api) |
| 22000 | [dp-dataset-api](https://github.com/ONSdigital/dp-dataset-api) |
| 22100 | [dp-filter-api](https://github.com/ONSdigital/dp-filter-api) |
| 22200 | [dp-import-reporter](https://github.com/ONSdigital/dp-import-reporter) |
| 22300 | [dp-recipe-api](https://github.com/ONSdigital/dp-recipe-api) |
| 22400 | [dp-code-list-api](https://github.com/ONSdigital/dp-code-list-api) |
| 22500 | [dp-dataset-exporter](https://github.com/ONSdigital/dp-dataset-exporter) |
| 22600 | [dp-hierarchy-api](https://github.com/ONSdigital/dp-hierarchy-api) |
| 22700 | [dp-hierarchy-builder](https://github.com/ONSdigital/dp-hierarchy-builder) |
| 22800 | [dp-dataset-exporter-xlsx](https://github.com/ONSdigital/dp-dataset-exporter-xlsx) |
| 22900 | [dp-dimension-search-builder](https://github.com/ONSdigital/dp-dimension-search-builder) |
| 23100 | [dp-dimension-search-api](https://github.com/ONSdigital/dp-dimension-search-api) |
| 23200 | [dp-api-router](https://github.com/ONSdigital/dp-api-router) |
| 23300 | [dp-table-renderer](https://github.com/ONSdigital/dp-table-renderer) |
| 23400 | [dp-file-downloader](https://github.com/ONSdigital/dp-file-downloader) |
| 23500 | [dp-map-renderer](https://github.com/ONSdigital/dp-map-renderer) |
| 23600 | [dp-developer-site](http://github.com/ONSdigital/dp-developer-site) |
| 23700 | [dp-frontend-geography-controller](https://github.com/ONSdigital/dp-frontend-geography-controller) |
| 23900 | [dp-search-api](https://github.com/ONSdigital/dp-search-api) |
| 24000 | [dp-publishing-dataset-controller](https://github.com/ONSdigital/dp-publishing-dataset-controller) |
| 24100 | [dp-frontend-cookie-controller](https://github.com/ONSdigital/dp-frontend-cookie-controller) |
| 24200 | [dp-bulletin-api](https://github.com/ONSdigital/dp-bulletin-api) |
| 24300 | [dp-deployer](https://github.com/ONSdigital/dp-deployer) |
| 24500 | [dp-observation-api](https://github.com/ONSdigital/dp-observation-api) |
| 24600 | [dp-legacy-redirector](https://github.com/ONSdigital/dp-legacy-redirector) |
| 24700 | [dp-image-api](https://github.com/ONSdigital/dp-image-api) |
| 24800 | [dp-image-importer](https://github.com/ONSdigital/dp-image-importer) |
| 24900 | [dp-static-file-publisher](https://github.com/ONSdigital/dp-static-file-publisher) |
| 25000 | [dp-frontend-search-controller](https://github.com/ONSdigital/dp-frontend-search-controller) |
| 25100 | [dp-upload-service](https://github.com/ONSdigital/dp-upload-service) |
| 25200 | [dp-frontend-feedback-controller](https://github.com/ONSdigital/dp-frontend-feedback-controller) |
| 25300 | [dp-topic-api](https://github.com/ONSdigital/dp-topic-api) |
| 25400 | [dp-permissions-api](https://github.com/ONSdigital/dp-permissions-api) |
| 25500 | [dp-areas-api](https://github.com/ONSdigital/dp-areas-api) |
| 25600 | [dp-identity-api](https://github.com/ONSdigital/dp-identity-api) |
| 25700 | [dp-search-reindex-api](https://github.com/ONSdigital/dp-search-reindex-api) |
| 25800 | [dp-search-data-extractor](https://github.com/ONSdigital/dp-search-data-extractor) |
| 25900 | [dp-search-data-importer](https://github.com/ONSdigital/dp-search-data-importer) |
| 26000 | [dp-collection-api](https://github.com/ONSdigital/dp-collection-api) |
| 26100 | [dp-import-cantabular-dataset](https://github.com/ONSdigital/dp-import-cantabular-dataset) |
| 26200 | [dp-import-cantabular-dimension-options](https://github.com/ONSdigital/dp-import-cantabular-dimension-options) |
| 26300 | [dp-cantabular-csv-exporter](https://github.com/ONSdigital/dp-cantabular-csv-exporter) |
