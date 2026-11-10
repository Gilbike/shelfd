# Contribution Guidelines

Thank you for taking interest in the project! Please read carefully and follow the guidelines, so we can ship the best possible experience.

> [!IMPORTANT]
>
> ### LLM Policy
>
> Using LLMs to aid your development and workflow (code completion, debugging, optimization) is _allowed_.
>
> Pull requests and commits with **clearly** LLM-generated, unreviewed code will be **automatically rejected** without warning.

## Getting Started with local development

### Requirements

- Go: `>= 1.27.0`
- Node.JS: `>= 22`
- pnpm: `>= 11`
- `golangci-lint` (for backend linting)

### Setting up

1. Clone the source code
2. Run `go mod download`
3. Download frontend dependencies with `cd web && pnpm install`
   - If you are not planning on changing the frontend it's better to just embed a static build. You can generate the static files with `go generate ./...`
4. To run the frontend development server run `pnpm dev` from the `web` folder
5. To run the backend development server run `go run ./cmd/server`

### Tests and linters

- Frontend commands: `pnpm test`, `pnpm check`, `pnpm lint`
- Backend commands: `golangci-lint run` (requires local setup), `go test ./...`

## How you can help

### Submitting issues

Feel free to suggest features and enhancements.

If you find any bug or security issue, please report it. Security issues should be reported privately on the GitHub Security Advisory page.

When submitting a bug report please make sure to check if it was already reported. When opening the issue please include steps on how to reproduce, the OS you are on and the version you are running.

### Submitting PRs

Make sure that a specific issue was not already assigned to another person. If you think you are able to contribute please ask for assignment under the issue you want to resolve.

#### What will be accepted

- Bug fixes
- Typo fixes
- Translations

#### What will be rejected

- Feature implementation without permission
- Features without open issue
- Fixes that were already assigned to someone else

## Making a pull request

Pull requests can be submitted onto the `master` branch

Before making any changes please be aware of the naming conventions

- Branches should be prefixed with `feat/` or `fix/` followed by a short description separated by hyphens. For example `fix/book-image-not-loading`, `feat/theme-toggle`
- Commits follow the [conventional commits](https://www.conventionalcommits.org/) format

Before submitting make sure you:

- Ran tests locally for both frontend and backend
- Ran linters and typecheckers locally for both frontend and backend
- Updated documentation and tests if needed

When submitting a PR:

- Keep the scope just to the issue
- Write a detailed description on the changes with carefully listing any breaking changes
- Mention any related issues (e.g., `Closes #123`)
- Ensure all tests and linters pass

## License

The project is licensed under the MIT license. Any PRs and commits submitted will also be under the MIT license.
