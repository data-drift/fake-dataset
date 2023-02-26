# Introduction

This repository holds csv files of randomly generated data for experimentation on historical data mutation, or data-drift.

In our case, the historical data is all our organisation pricing plans of the past, month by month.

Ex: At the end of january, we print all the organisations that had a pricing plan in january and store them in the `mart/organisation-mrr` table.

The [fake-data-generator.go](fake-data-generator.go) is a script to generate that table in a csv format.

[WIP] The other script job is to randomly modify historical data:

```sh
go run main.go introduceDataDrift mart/organisation-mrr.csv
```

- **Modify a historical pricing plan**. For instance, let's say we are in March. Due to a human error `organisationA` was set to have a 50€ pricing plan in January, but in reality they had a 30€ pricing plan.
- **Add or remove an organisation**. For instance, due to a bug in the pipeline, we discover that 10 organisations were not stored in the `mart/organisation-mrr` table in january. So they are added later.

TODO:

- [x] Add a github action every day that introduces a data drift
- [x] implement data drift: Deleted Line
- [ ] implement data drift: Added Line
- [ ] implement data drift: Modified Line
