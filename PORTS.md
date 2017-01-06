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
| 8080  | [dp-dd-dimensional-metadata-api](https://github.com/ONSdigital/dp-dd-dimensional-metadata-api) | Conflict - needs updating
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
| 20010 | [dp-frontend-renderer](https://github.com/ONSdigital/dp-frontend-renderer) |
| 20019 | [dp-dd-file-uploader](https://github.com/ONSdigital/dp-dd-file-uploader) |
| 20020 | [dp-content-resolver](https://github.com/ONSdigital/dp-content-resolver) |
| 20030 | [dp-dd-frontend-controller](https://github.com/ONSdigital/dp-dd-frontend-controller) |
| 20040 | [dp-dd-react-app](https://github.com/ONSdigital/dp-dd-react-app) |
| 20099 | [dp-dd-api-stub](https://github.com/ONSdigital/dp-dd-api-stub) |
| 20100 | [dp-dd-job-creator-api-stub](https://github.com/ONSdigital/dp-dd-job-creator-api-stub) |
| 21000 | [dp-csv-splitter](https://github.com/ONSdigital/dp-csv-splitter) |
|       | [dp-dd-csv-filter](https://github.com/ONSdigital/dp-dd-csv-filter) | No code
