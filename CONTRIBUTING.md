# Contributing to mus-go

First off, thank you for considering contributing to **mus-go**! It's 
people like you who make the open-source community such an amazing place to 
learn, inspire, and create.

All types of contributions are welcome: from reporting bugs and suggesting new 
features to improving documentation or submitting code changes.

## Code of Conduct

By participating in this project, you agree to abide by the [Contributor Covenant](https://www.contributor-covenant.org/version/2/1/code_of_conduct.html)
to ensure a welcoming and inclusive environment for everyone.

## Getting Started

### Prerequisites

- **Go 1.24+**: The project uses modern Go features (like the latest loop 
  variable semantics).
- **golangci-lint**: We use this for static analysis.

### Fork and Clone

1. Fork the repository on GitHub.
2. Clone your fork locally:

   ```bash
   git clone https://github.com/YOUR_USERNAME/mus-go.git
   cd mus-go
   ```

3. Add the upstream repository as a remote:

   ```bash
   git remote add upstream https://github.com/mus-format/mus-go.git
   ```

## Reporting Issues

- **Bugs**: Use the [GitHub Issues](https://github.com/mus-format/mus-go/issues)
  page. Please include a clear description and, if possible, a minimal 
  reproducible example.
- **Features**: We're open to suggestions! Open an issue to discuss your idea 
  before starting implementation.

## Development Workflow

1. Create a descriptive branch for your changes:

   ```bash
   git checkout -b feature/my-new-feature
   ```

2. Make your changes, following the [Coding Standards](#coding-standards).
3. Ensure all tests pass and there are no linting issues.
4. Commit your changes using [Conventional Commits](https://www.conventionalcommits.org/):

   ```bash
   git commit -m "feat: add support for X"
   ```

## Coding Standards

- **Effective Go**: Follow the official [Effective Go](https://golang.org/doc/effective_go)
  and [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments).
- **Documentation**: All exported types, interfaces, and functions **must** have
  professional documentation comments. This is enforced by our linter.
- **Performance**: As a high-performance library, we prioritize efficiency. 
  Avoid unnecessary allocations and use tools like `prealloc` and `perfsprint` 
  (included in our linting suite).
- **Naming**: Avoid "stuttering" in names. For example, use `group.Client` 
  instead of `group.GroupClient`.

## Testing and Linting

We maintain strict quality standards to ensure the library remains reliable.

### Linting

Always run the linter before submitting a PR:

```bash
golangci-lint run
```

Our configuration focuses on performance markers, resource safety (body closing),
and idiomatic style.

### Testing

Run the full test suite:

```bash
go test ./...
```

New features should always include unit tests. For complex integrations, consider
adding cases to the `test/` directory.

### Coverage

We monitor code coverage. To run tests with coverage:

```bash
go test -coverprofile=coverage.txt $(go list ./... | grep -v "test")
```

*(Note: Core logic coverage is prioritized; helpers in `test/` are excluded.)*

## Pull Request Process

1. Push your branch to your fork.
2. Open a Pull Request against the `main` branch.
3. Ensure the CI/CD pipeline (Static Analysis, Security Scan, and Coverage) passes.
4. Once reviewed and approved, your changes will be merged.

## License

By contributing to **mus-go**, you agree that your contributions will be 
licensed under the project's [MIT License](LICENSE).

---

Thank you for your contribution!
