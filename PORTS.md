Port allocations
================

These are the port numbers used by DP services.

Using distinct port numbers for each services keeps it simple to start everything
in development without having to configure port numbers etc.

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
| 9000  | [dp-dd-database-loader](https://github.com/ONSdigital/dp-dd-database-loader) | Conflict - needs updating
| 9200  | [dp-compose](https://github.com/ONSdigital/dp-compose) : elasticsearch |
| 9300  | [dp-compose](https://github.com/ONSdigital/dp-compose) : elasticsearch |
| 9999  | [dp-compose](https://github.com/ONSdigital/dp-compose) : highcharts |
| 20000 | [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router) |
| 20001 | [dp-frontend-filter-dataset-controller](https://github.com/ONSdigital/dp-frontend-filter-dataset-controller)
| 20010 | [dp-frontend-renderer](https://github.com/ONSdigital/dp-frontend-renderer) |
| 20019 | [dp-dd-file-uploader](https://github.com/ONSdigital/dp-dd-file-uploader) |
| 20020 | [dp-content-resolver](https://github.com/ONSdigital/dp-content-resolver) |
| 20030 | [dp-dd-frontend-controller](https://github.com/ONSdigital/dp-dd-frontend-controller) |
| 20040 | [dp-dd-react-app](https://github.com/ONSdigital/dp-dd-react-app) |
| 20050 | [dp-dd-search-indexer](https://github.com/ONSdigital/dp-dd-search-indexer) |
| 20051 | [dp-dd-search-api](https://github.com/ONSdigital/dp-dd-search-api) |
| 20099 | [dp-dd-api-stub](https://github.com/ONSdigital/dp-dd-api-stub) |
| 20099 | [dp-dd-dimensional-metadata-api](https://github.com/ONSdigital/dp-dd-dimensional-metadata-api) | shares port with stub
| 20100 | [dp-dd-job-creator-api-stub](https://github.com/ONSdigital/dp-dd-job-creator-api-stub) |
| 20200 | [dp-frontend-dataset-controller](https://github.com/ONSdigital/dp-frontend-dataset-controller)
| 21000 | [dp-csv-splitter](https://github.com/ONSdigital/dp-csv-splitter) |
| 21100 | [dp-dd-csv-filter](https://github.com/ONSdigital/dp-dd-csv-filter) |
| 21200 | [dp-dd-csv-transformer](https://github.com/ONSdigital/dp-dd-csv-transformer) |
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
| 22900 | [dp-search-builder](https://github.com/ONSdigital/dp-search-builder) |
| 23000 | [dp-dd-metadata-editor](https://github.com/ONSdigital/dp-dd-metadata-editor) |
| 23100 | [dp-search-api](https://github.com/ONSdigital/dp-search-api) |
