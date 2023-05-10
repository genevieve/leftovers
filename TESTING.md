# Testing

We use [ginkgo](https://onsi.github.io/ginkgo/) for our acceptance and unit tests.

To install, you can follow [these instructions](https://onsi.github.io/ginkgo/#installing-ginkgo).

## Unit tests

Run `ginkgo -r -p -race .`.

### Fakes

We use [faux](https://github.com/ryanmoran/faux) to generate fakes.

```
brew tap ryanmoran/tools
brew install faux
```

To generate a fake, run: `go generate ./...`

## Acceptance tests

1. Export the IaaS-specific environment variables.
1. Export `LEFTOVERS_ACCEPTANCE` to be the value of the IaaS you want to test.
1. Run `ginkgo -r -p -race acceptance`.
