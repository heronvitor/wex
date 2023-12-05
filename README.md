# wex

A web application to store purchases. It provides an endpoint to save a purchase and other to retrieve the purchase with the amount converted to a specific currency.

The application consists of two parts: one is a job to sync the exchange rates table with the api provided, the other is the web application that provides the api rest. 

## Update Job
The update job can be deployed as a kubernetes cronjob to run periodicaly. The interval between updates and retry attemps can be chosen using command flags and the execution will be skiped when is not necessary. This strategy was chosen because this data is stable and sufficiently small to be stored locally, avoind the dependence of an external api for each request.
```bash
$ wex update -h
Usage: wex update

Update Exchange Rates

Flags:
  -h, --help                   Show context-sensitive help.
      --db-url=STRING          ($DB_URL)

      --update-interval=24h
      --retry-interval=1h
```

## Api Rest
The api rest provides a swagger documentation in the root path.

```bash
Usage: wex api

Run api

Flags:
  -h, --help             Show context-sensitive help.
      --db-url=STRING    ($DB_URL)

      --port=8080        ($PORT)
```

## Run with docker-compose
```bash
$ docker-compose up -d # run all: db, migration, app
# then you can access http://localhost:8080/swagger/index.html
```
```bash
$ docker-compose run wex update # run job to update exchange rates
```
### Code Architecture
It has been made using a tree layer architecture: data access layer, business and presentation layer. The architecture was chosen because is allow the separation of responsabilities, facilitates tests using dependecy injections and is not too complex.


## Known issues
The currencies names repeats between countries, so I'm going to change the primary key and the endpoint to use the country information too. 