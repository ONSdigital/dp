# Port allocations

These are the port numbers used by dissemination services when running locally.

Using distinct port numbers for each service means we can run everything
in development without having to configure them individually.

## Default port numbers

| Port  | Service                                                                                                        | Notes |
|-------|----------------------------------------------------------------------------------------------------------------|-------|
| 2181  | [dp-compose](https://github.com/ONSdigital/dp-compose) : zookeeper                                             |       |
| 5432  | [dp-compose](https://github.com/ONSdigital/dp-compose) : postgres                                              |       |
| 8000  | [dis-wagtail](https://github.com/ONSdigital/dis-wagtail)                                                       |       |
| 8080  | [babbage](https://github.com/ONSdigital/babbage)                                                               |       |
| 8081  | [florence](https://github.com/ONSdigital/florence)                                                             |       |
| 8082  | [zebedee](https://github.com/ONSdigital/zebedee)                                                               |       |
| 8083  | [project-brian](https://github.com/ONSdigital/project-brian)                                                   |       |
| 8084  | [The-Train](https://github.com/ONSdigital/The-Train)                                                           |       |
| 8888  | [dp-compose](https://github.com/ONSdigital/dp-compose) : mathjax-api                                           |       |
| 9000  | [sixteens](https://github.com/ONSdigital/sixteens)                                                             |       |
| 9001  | [dp-compose](https://github.com/ONSdigital/dp-compose) : minio                                                 |       |
| 9002  | [dp-design-system](https://github.com/ONSdigital/dp-design-system)                                             |       |
| 9092  | [dp-compose](https://github.com/ONSdigital/dp-compose) : kafka                                                 |       |
| 9200  | [dp-compose](https://github.com/ONSdigital/dp-compose) : elasticsearch                                         |       |
| 9300  | [dp-compose](https://github.com/ONSdigital/dp-compose) : elasticsearch                                         |       |
| 9999  | [dp-compose](https://github.com/ONSdigital/dp-compose) : highcharts                                            |       |
| 10200 | [dp-compose](https://github.com/ONSdigital/dp-compose) : cmd elasticsearch                                     |       |
| 10300 | [dp-compose](https://github.com/ONSdigital/dp-compose) : cmd elasticsearch                                     |       |
| 11200 | [dp-compose](https://github.com/ONSdigital/dp-compose) : site wide elasticsearch                               |       |
| 11300 | [dp-compose](https://github.com/ONSdigital/dp-compose) : site wide elasticsearch                               |       |
| 20000 | [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router)                                         |       |
| 20001 | [dp-frontend-filter-dataset-controller](https://github.com/ONSdigital/dp-frontend-filter-dataset-controller)   |       |
| 20020 | [dp-content-resolver](https://github.com/ONSdigital/dp-content-resolver)                                       |       |
| 20100 | [dp-frontend-filter-flex-dataset](https://github.com/ONSdigital/dp-frontend-filter-flex-dataset)               |       |
| 20200 | [dp-frontend-dataset-controller](https://github.com/ONSdigital/dp-frontend-dataset-controller)                 |       |
| 21000 | [dp-csv-splitter](https://github.com/ONSdigital/dp-csv-splitter)                                               |       |
| 21300 | [dp-import-tracker](https://github.com/ONSdigital/dp-import-tracker)                                           |       |
| 21400 | [dp-dimension-extractor](https://github.com/ONSdigital/dp-dimension-extractor)                                 |       |
| 21500 | [dp-dimension-importer](https://github.com/ONSdigital/dp-dimension-importer)                                   |       |
| 21600 | [dp-observation-extractor](https://github.com/ONSdigital/dp-observation-extractor)                             |       |
| 21700 | [dp-observation-importer](https://github.com/ONSdigital/dp-observation-importer)                               |       |
| 21800 | [dp-import-api](https://github.com/ONSdigital/dp-import-api)                                                   |       |
| 21900 | [dp-code-list-api](https://github.com/ONSdigital/dp-code-list-api)                                             |       |
| 22000 | [dp-dataset-api](https://github.com/ONSdigital/dp-dataset-api)                                                 |       |
| 22100 | [dp-filter-api](https://github.com/ONSdigital/dp-filter-api)                                                   |       |
| 22200 | [dp-import-reporter](https://github.com/ONSdigital/dp-import-reporter)                                         |       |
| 22300 | [dp-recipe-api](https://github.com/ONSdigital/dp-recipe-api)                                                   |       |
| 22400 | [dp-code-list-api](https://github.com/ONSdigital/dp-code-list-api)                                             |       |
| 22500 | [dp-dataset-exporter](https://github.com/ONSdigital/dp-dataset-exporter)                                       |       |
| 22600 | [dp-hierarchy-api](https://github.com/ONSdigital/dp-hierarchy-api)                                             |       |
| 22700 | [dp-hierarchy-builder](https://github.com/ONSdigital/dp-hierarchy-builder)                                     |       |
| 22800 | [dp-dataset-exporter-xlsx](https://github.com/ONSdigital/dp-dataset-exporter-xlsx)                             |       |
| 22900 | [dp-dimension-search-builder](https://github.com/ONSdigital/dp-dimension-search-builder)                       |       |
| 23100 | [dp-dimension-search-api](https://github.com/ONSdigital/dp-dimension-search-api)                               |       |
| 23200 | [dp-api-router](https://github.com/ONSdigital/dp-api-router)                                                   |       |
| 23300 | [dp-table-renderer](https://github.com/ONSdigital/dp-table-renderer)                                           |       |
| 23400 | [dp-file-downloader](https://github.com/ONSdigital/dp-file-downloader)                                         |       |
| 23600 | [dp-developer-site](http://github.com/ONSdigital/dp-developer-site)                                            |       |
| 23900 | [dp-search-api](https://github.com/ONSdigital/dp-search-api)                                                   |       |
| 24000 | [dp-publishing-dataset-controller](https://github.com/ONSdigital/dp-publishing-dataset-controller)             |       |
| 24100 | [dp-frontend-cookie-controller](https://github.com/ONSdigital/dp-frontend-cookie-controller)                   |       |
| 24300 | [dp-deployer](https://github.com/ONSdigital/dp-deployer)                                                       |       |
| 24400 | [dp-frontend-homepage-controller](https://github.com/ONSdigital/dp-frontend-homepage-controller)               |       |
| 24500 | [dp-observation-api](https://github.com/ONSdigital/dp-observation-api)                                         |       |
| 24600 | [dp-legacy-redirector](https://github.com/ONSdigital/dp-legacy-redirector)                                     |       |
| 24700 | [dp-image-api](https://github.com/ONSdigital/dp-image-api)                                                     |       |
| 24800 | [dp-image-importer](https://github.com/ONSdigital/dp-image-importer)                                           |       |
| 24900 | [dp-static-file-publisher](https://github.com/ONSdigital/dp-static-file-publisher)                             |       |
| 25000 | [dp-frontend-search-controller](https://github.com/ONSdigital/dp-frontend-search-controller)                   |       |
| 25100 | [dp-upload-service](https://github.com/ONSdigital/dp-upload-service)                                           |       |
| 25200 | [dp-frontend-feedback-controller](https://github.com/ONSdigital/dp-frontend-feedback-controller)               |       |
| 25300 | [dp-topic-api](https://github.com/ONSdigital/dp-topic-api)                                                     |       |
| 25400 | [dp-permissions-api](https://github.com/ONSdigital/dp-permissions-api)                                         |       |
| 25600 | [dp-identity-api](https://github.com/ONSdigital/dp-identity-api)                                               |       |
| 25800 | [dp-search-data-extractor](https://github.com/ONSdigital/dp-search-data-extractor)                             |       |
| 25900 | [dp-search-data-importer](https://github.com/ONSdigital/dp-search-data-importer)                               |       |
| 26100 | [dp-import-cantabular-dataset](https://github.com/ONSdigital/dp-import-cantabular-dataset)                     |       |
| 26200 | [dp-import-cantabular-dimension-options](https://github.com/ONSdigital/dp-import-cantabular-dimension-options) |       |
| 26300 | [dp-cantabular-csv-exporter](https://github.com/ONSdigital/dp-cantabular-csv-exporter)                         |       |
| 26400 | [dp-content-api](https://github.com/ONSdigital/dp-content-api)                                                 |       |
| 26700 | [dp-cantabular-metadata-exporter](https://github.com/ONSdigital/dp-cantabular-metadata-exporter)               |       |
| 26800 | [dp-cantabular-xlsx-exporter](https://github.com/ONSdigital/dp-cantabular-xlsx-exporter)                       |       |
| 26900 | [dp-files-api](https://github.com/ONSdigital/dp-files-api)                                                     |       |
| 27017 | [dp-compose](https://github.com/ONSdigital/dp-compose) : mongo                                                 |       |
| 27100 | [dp-cantabular-filter-flex-api](https://github.com/ONSdigital/dp-cantabular-filter-flex-api)                   |       |
| 27200 | [dp-cantabular-dimension-api](https://github.com/ONSdigital/dp-cantabular-dimension-api)                       |       |
| 27300 | [dp-population-types-api](https://github.com/ONSdigital/dp-population-types-api)                               |       |
| 27700 | [dp-frontend-release-calendar](https://github.com/ONSdigital/dp-frontend-release-calendar)                     |       |
| 27800 | [dp-release-calendar-api](https://github.com/ONSdigital/dp-release-calendar-api)                               |       |
| 28100 | [dp-census-atlas](https://github.com/onSdigital/dp-census-atlas)                                               |       |
| 28300 | [dp-cantabular-metadata-extractor-api](https://github.com/ONSdigital/dp-cantabular-metadata-extractor-api)     |       |
| 28600 | [dp-feedback-api](https://github.com/ONSdigital/dp-feedback-api)                                               |       |
| 28700 | [dp-search-scrubber-api](https://github.com/ONSdigital/dp-search-scrubber-api)                                 |       |
| 28800 | [dp-nlp-category-api](https://github.com/ONSdigital/dp-nlp-category-api)                                       |       |
| 28900 | [dp-nlp-berlin-api](https://github.com/ONSdigital/dp-nlp-berlin-api)                                           |       |
| 29000 | [dp-cpih-export](https://github.com/ONSdigital/dp-cpih-export)                                                 |       |
| 29100 | [dp-legacy-cache-api](https://github.com/ONSdigital/dp-legacy-cache-api)                                       |       |
| 29200 | [dp-legacy-cache-proxy](https://github.com/ONSdigital/dp-legacy-cache-proxy)                                   |       |
| 29300 | [dp-data-api](https://github.com/ONSdigital/dp-data-api)                                                       |       |
| 29400 | [dis-data-admin-ui](https://github.com/ONSdigital/dis-data-admin-ui)                                           |       |
| 29500 | [dis-authentication-stub](https://github.com/ONSdigital/dis-authentication-stub)                               |       |
| 29600 | [dis-search-upstream-stub](https://github.com/ONSdigital/dis-search-upstream-stub)                             |       |
| 29700 | [dis-routing-api-poc](https://github.com/ONSdigital/dis-routing-api-poc)                                       |       |
| 29800 | [dis-bundle-api](https://github.com/ONSdigital/dis-bundle-api)                                                 |       |
| 29900 | [dis-redirect-api](https://github.com/ONSdigital/dis-redirect-api)                                             |       |
| 30000 | [dis-redirect-proxy](https://github.com/ONSdigital/dis-redirect-proxy)                                         |       |
