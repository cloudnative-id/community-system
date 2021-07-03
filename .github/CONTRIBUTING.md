# Contributing
By participating to this project, you agree to abide our [code of conduct](https://github.com/cloudnative-id/community-system/blob/master/.github/CODE_OF_CONDUCT.md).

## Development
For small things like fixing typos in documentation, you can [make edits through GitHub](https://help.github.com/articles/editing-files-in-another-user-s-repository/), which will handle forking and making a pull request (PR) for you. For anything bigger or more complex, you'll probably want to set up a development environment on your machine, a quick procedure for which is as folows:

### Setup your machine
Prerequisites:
- make
- [Go 1.15+](https://golang.org/doc/install)
- [Docker & docker-compose]

- Setup your environment
```
make dev.environment.up
source .env.local
```

- Run locally
```
make dev.run
```
