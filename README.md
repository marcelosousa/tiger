# tiger :tiger:

This is an experimental CLI that uses Reviewpad configurations to automate workflows locally.

Git hooks with reviewpad :tiger:

- Enforce conventional commits.

## Installation

Install `tiger` with:

```console
go install github.com/marcelosousa/tiger@latest
```

Install git hooks in current repository with:

```console
tiger init
```

You can check the available hooks [here](cmd/init.go).
