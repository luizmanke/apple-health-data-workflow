# Apple Health Data Workflow

This project aims to extract, transform and load Apple health data to ease further data analysis.

## How to use

1. Export the Apple Health metrics to a CSV file using the [Health Auto Export](https://www.healthexportapp.com/) app.

2. Place the exported CSV files into the [./data](./data/) folder.

3. Run the application:

```sh
make run
```

## Development

Run tests:

```sh
make tests
```
