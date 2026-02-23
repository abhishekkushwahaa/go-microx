# Contributing to go-microx

First off, thank you for considering contributing to **go-microx**! It's people like you who make this project such a great tool for the Go community.

This document provides guidelines and best practices to help you contribute to the project effectively.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How Can I Contribute?](#how-can-i-contribute)
  - [Reporting Bugs](#reporting-bugs)
  - [Suggesting Enhancements](#suggesting-enhancements)
  - [Code Contributions](#code-contributions)
- [Development Workflow](#-development-workflow)
  - [Prerequisites](#prerequisites)
  - [Local Setup](#local-setup)
- [Coding Standards](#-coding-standards)
- [Pull Request Process](#-pull-request-process)
- [Commit Message Guidelines](#-commit-message-guidelines)
- [License](#-license)

## Code of Conduct

We are committed to providing a welcoming and inspiring community for all. By participating in this project, you agree to abide by our Code of Conduct (based on the [Contributor Covenant](https://www.contributor-covenant.org/version/2/1/code_of_conduct/)).

## How Can I Contribute?

### Reporting Bugs

If you find a bug, please help us by [opening an issue](https://github.com/abhishekkushwahaa/go-microx/issues). A good bug report includes:

- A clear, descriptive title.
- Steps to reproduce the problem.
- Expected vs. actual behavior.
- Screenshots or log outputs if applicable.
- Your environment details (OS, Go version).

### Suggesting Enhancements

We're always looking for ways to improve! To suggest an enhancement, please open an issue and describe:

- The problem this enhancement solves.
- A detailed description of the proposed feature.
- Potential implementation ideas (if any).

### Code Contributions

If you're looking to dive into the code:

1. Browse the [issues](https://github.com/abhishekkushwahaa/go-microx/issues) for any tasks labeled `help wanted` or `good first issue`.
2. Comment on the issue to let us know you're working on it.
3. If it's a new feature, please open an issue first to discuss the design before starting work.

## Development Workflow

### Prerequisites

- **Go**: v1.22 or higher.
- **Git**: For version control.
- **Make**: For running project commands (optional but recommended).

### Local Setup

1. **Fork the repository** on GitHub.
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/go-microx.git
   cd go-microx
   ```
3. **Add the upstream remote**:
   ```bash
   git remote add upstream https://github.com/abhishekkushwahaa/go-microx.git
   ```
4. **Create a branch** for your work:
   ```bash
   git checkout -b feature/your-feature-name
   # or
   git checkout -b bugfix/your-bug-name
   ```

## Pull Request Process

1. Ensure your code satisfies the **Coding Standards**.
2. Update the `README.md` or `docs/` if you've added new features or changed behaviors.
3. Verify that all tests pass: `go test ./...`.
4. Push your changes to your fork.
5. Submit a [Pull Request](https://github.com/abhishekkushwahaa/go-microx/pulls) against the `main` branch.
6. A maintainer will review your PR. Be prepared to address feedback!

## Commit Message Guidelines

We prefer clear, concise commit messages following the [Conventional Commits](https://www.conventionalcommits.org/) specification:

- `feat:` for new features.
- `fix:` for bug fixes.
- `docs:` for documentation changes.
- `style:` for formatting/styling (no logic changes).
- `refactor:` for code changes that neither fix a bug nor add a feature.
- `test:` for adding or updating tests.
- `chore:` for updating build tasks, dependencies, etc.

Example: `feat(cli): add interactive auth selection`

## License

By contributing to **go-microx**, you agree that your contributions will be licensed under its [MIT License](LICENSE).

<div align="center">
Thank you for your contribution!
</div>
