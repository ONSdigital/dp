Creating a new API microservice
==================

The following steps will guide you through adding a new API to our platform.

**NOTE:** For the specifics of getting a new service up and running follow our [new app](NEW_APP.md) guide first, and come back here for the specific *additional* steps required for APIs.

Prerequisites
-------------

If any of the following Github links 404, you will likely need to be added to the ONSdigital organisation.

To follow this guide you will need:

* A Github account with access to the ONSdigital organisation

* The [dp-api-router](https://github.com/ONSdigital/dp-api-router) repo cloned locally

* [dp-configs](https://github.com/ONSdigital/dp-configs) repo cloned locally

* The ['magic port'](https://github.com/ONSdigital/dp-setup/tree/awsb/PORTS.md) assigned to your microservice API

Getting started
---------------

1. Identify the top level route your API will be responsible for. This likely can be taken from the repository name e.g. dp-dataset-api has the top level route **/datasets**. Make sure it doesn't clash with any existing routes [in the router](https://github.com/ONSdigital/dp-api-router/blob/develop/service/service.go#L126)

2. Add a new API URL variable to the dp-api-router [config.go](https://github.com/ONSdigital/dp-api-router/blob/develop/config/config.go#L19) and set a [default value](https://github.com/ONSdigital/dp-api-router/blob/develop/config/config.go#L71) using the localhost:port combination the app will default to when running locally. The default ports [should be documented](https://github.com/ONSdigital/dp/blob/main/guides/PORTS.md) if not already.

3. Update the dp-api-router [secrets for relevant environments](https://github.com/ONSdigital/dp-configs/tree/master/secrets) to contain your new API URL environment variable, listing the ['magic port'](https://github.com/ONSdigital/dp-setup/tree/awsb/PORTS.md) instead of the default value.

3. Set up a new API Proxy in [service.go](https://github.com/ONSdigital/dp-api-router/blob/develop/service/service.go#L138), referencing the variable you added to config.go above.

4. Determine where best to add the [handler line](https://github.com/ONSdigital/dp-api-router/blob/develop/service/service.go#L145) - replacing the values for your new API Proxy and the top level route you identified in step 1.

    1. If your API is still in development, you likely want to add it behind a feature flag [as shown by the Observation API](https://github.com/ONSdigital/dp-api-router/blob/develop/service/service.go#L130) so that the route doesn't start to forward in production until you are ready. This will also mean adding an extra feature flag to the config.go file.
    2. If your API should only be available privately (for use interally/by the publishing system) it should be created [inside the private endpoints list](https://github.com/ONSdigital/dp-api-router/blob/develop/service/service.go#L155). It is acceptable to further control an API behind a feature flag while in development within this section.

5. If some features of your API should only be available internally, update your new service to include a feature flag for 'privateEndpoints' and only attach those routes to your service when enabled [as shown here](https://github.com/ONSdigital/dp-dimension-search-api/blob/develop/api/api.go#L120)

6. Update the [secrets](https://github.com/ONSdigital/dp-configs/tree/master/secrets) for the dp-api-router and your new application to reflect all feature flags your may have added an enable them only in develop while you get your new API ready for production use.

7. Merge your changes to both applications and deploy to develop!

You should now see your API reachable through https://api.dp.aws.onsdigital.uk/v1/<your-route>
