# Contributing to go-microx

We appreciate your interest in contributing to **go-microx**! Whether you're submitting a bug report, suggesting a new feature, or improving the documentation, your contribution is valuable.

Before you begin, please take a moment to read through the guidelines below to ensure a smooth and consistent contribution process.

## Table of Contents

- [Contributing to go-microx](#contributing-to-go-microx)
  - [Table of Contents](#table-of-contents)
  - [Design Principles](#design-principles)
  - [Reporting Issues](#reporting-issues)
  - [Contributing Code with Pull Requests](#contributing-code-with-pull-requests)
    - [Requirements](#requirements)
  - [Licensing](#licensing)

## Design Principles

Contributions to **go-microx** should align with the project’s design principles:

- **Maintain backwards compatibility** whenever possible. Ensure that changes do not break existing features or functionality unless absolutely necessary.
- **Follow clean code practices**. Code should be readable, maintainable, and well-documented.
- **Keep the scope clear and focused**. If you're implementing a feature, try to ensure it’s narrowly scoped so it can be merged and tested quickly.

## Reporting Issues

If you encounter a bug, have questions, or want to discuss new features, please [file an issue](https://github.com/abhishekkushwahaa/go-microx/issues). When creating an issue, please provide the following:

- A clear description of the problem or feature request.
- Steps to reproduce the bug, if applicable.
- Any logs, screenshots, or error messages that can help us understand the issue.

This will help maintainers triage and address issues more effectively.

## Contributing Code with Pull Requests

**go-microx** uses [GitHub pull requests](https://github.com/abhishekkushwahaa/go-microx/pulls) for contributions. If you have a feature or bug fix you’d like to contribute, feel free to fork the repository, make your changes, and submit a pull request (PR).

### Requirements

- **Documentation**: All commands, features, and updates should be properly documented. This ensures that other developers understand how to use the feature or functionality.
- **Unit Testing**: Any new functionality should have accompanying unit tests. This ensures that the new feature works as expected and does not break existing functionality.

- **Style and Consistency**: Follow the existing code style. Pay attention to naming conventions, indentation, and function structure. This helps keep the codebase clean and readable.

- **Testing**: Ensure that the code passes all tests before submitting a PR. Use the provided `Makefile` to run tests and build the project:
  - Run `make test` to ensure all tests pass.
  - Run `make build` to verify the binary compiles correctly.

- **Commit Messages**: Write clear, concise commit messages. Follow this convention:
  - Use the present tense (e.g., "Add new feature" instead of "Added new feature").
  - Provide a short, meaningful description of what the commit does.

Once your PR is ready, submit it and we'll review your contribution. If everything looks good, it will be merged.

## Licensing

By contributing to **go-microx**, you agree to license your contributions under the terms of the project's [MIT License](https://github.com/abhishekkushwahaa/go-microx/blob/main/LICENSE).

You will be asked to confirm this licensing when submitting your pull request.

---

Thank you for contributing! We look forward to working with you to make **go-microx** even better.
