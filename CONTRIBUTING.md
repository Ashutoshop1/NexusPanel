# Contributing to NexusPanel

Thank you for your interest in contributing to NexusPanel! This document provides guidelines and instructions for contributing.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How to Contribute](#how-to-contribute)
- [Reporting Issues](#reporting-issues)
- [Submitting Pull Requests](#submitting-pull-requests)
- [Development Setup](#development-setup)
- [Coding Standards](#coding-standards)
- [Commit Message Guidelines](#commit-message-guidelines)
- [Testing](#testing)
- [Documentation](#documentation)

---

## Code of Conduct

By participating in this project, you agree to maintain a respectful and inclusive environment for everyone.

### Our Standards

- Be respectful and considerate
- Welcome newcomers and help them learn
- Focus on constructive feedback
- Accept responsibility for mistakes
- Prioritize the community's best interest

---

## How to Contribute

There are many ways to contribute to NexusPanel:

### 1. **Report Bugs**
   - Search existing issues to avoid duplicates
   - Use the bug report template
   - Provide detailed reproduction steps
   - Include system information and logs

### 2. **Suggest Features**
   - Check if the feature already exists or is planned
   - Use the feature request template
   - Explain the use case and benefits
   - Provide examples if possible

### 3. **Improve Documentation**
   - Fix typos and unclear explanations
   - Add examples and tutorials
   - Translate documentation
   - Update outdated information

### 4. **Submit Code**
   - Fix bugs
   - Implement new features
   - Improve performance
   - Refactor code
   - Write tests

---

## Reporting Issues

### Before Submitting an Issue

1. **Search existing issues** to check if it's already reported
2. **Update to the latest version** to see if the issue persists
3. **Gather information**:
   - NexusPanel version
   - Operating system and version
   - Browser and version (for UI issues)
   - Steps to reproduce
   - Expected vs actual behavior
   - Error messages and logs

### Issue Template

```markdown
**Description**
A clear description of the issue.

**Steps to Reproduce**
1. Go to '...'
2. Click on '...'
3. See error

**Expected Behavior**
What you expected to happen.

**Actual Behavior**
What actually happened.

**Environment**
- NexusPanel version: vX.Y.Z
- OS: Ubuntu 22.04 / macOS 13 / Windows 11
- Browser: Chrome 120 (if applicable)

**Logs**
```
Paste relevant logs here
```

**Screenshots**
If applicable, add screenshots.

**Additional Context**
Any other relevant information.
```

---

## Submitting Pull Requests

### Before You Start

1. **Create or comment on an issue** to discuss your proposed changes
2. **Fork the repository** to your GitHub account
3. **Create a feature branch** from `main`:
   ```bash
   git checkout -b feature/your-feature-name
   # or
   git checkout -b fix/your-bug-fix
   ```

### Making Changes

1. **Follow coding standards** (see below)
2. **Write clear, focused commits**
3. **Add tests** for new features or bug fixes
4. **Update documentation** as needed
5. **Ensure all tests pass**
6. **Run linters and formatters**

### Submitting Your PR

1. **Push your branch** to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```

2. **Create a Pull Request** on GitHub with:
   - Clear title summarizing the change
   - Description explaining what and why
   - Reference to related issues (e.g., "Fixes #123")
   - Screenshots for UI changes
   - Checklist of completed items

3. **Respond to feedback** promptly and respectfully

### PR Template

```markdown
**Description**
Brief description of the changes.

**Related Issues**
Fixes #123, Related to #456

**Type of Change**
- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to change)
- [ ] Documentation update

**Checklist**
- [ ] My code follows the project's coding standards
- [ ] I have performed a self-review of my code
- [ ] I have commented my code where necessary
- [ ] I have updated the documentation
- [ ] I have added tests that prove my fix/feature works
- [ ] All new and existing tests pass
- [ ] I have run go fmt and go vet (for Go code)
- [ ] I have run ESLint (for frontend code)

**Screenshots** (if applicable)
Add screenshots here.

**Additional Notes**
Any additional information.
```

---

## Development Setup

### Prerequisites

- **Go**: 1.21 or higher
- **Node.js**: 18.x or higher
- **PostgreSQL**: 13+ (or use SQLite for development)
- **Git**: Latest version

### Backend Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/2670044605/NexusPanel.git
   cd NexusPanel
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Set up configuration**:
   ```bash
   cp configs/config.example.yaml configs/config.yaml
   # Edit config.yaml with your settings
   ```

4. **Run database migrations**:
   ```bash
   # Migrations run automatically on first start
   ```

5. **Build and run**:
   ```bash
   make build
   ./bin/nexuspanel-server
   # or
   make run
   ```

### Frontend Setup

1. **Navigate to web directory**:
   ```bash
   cd web
   ```

2. **Install dependencies**:
   ```bash
   npm install
   ```

3. **Run development server**:
   ```bash
   npm run dev
   ```

4. **Build for production**:
   ```bash
   npm run build
   ```

### Running Tests

**Backend tests**:
```bash
go test ./...
go test -race ./...  # With race detector
go test -cover ./... # With coverage
```

**Frontend tests**:
```bash
cd web
npm run test
npm run test:coverage
```

---

## Coding Standards

### Go Code

1. **Follow Go conventions**:
   - Use `gofmt` to format code
   - Use `go vet` to check for issues
   - Follow [Effective Go](https://golang.org/doc/effective_go.html)

2. **Naming conventions**:
   - Use `camelCase` for private functions/variables
   - Use `PascalCase` for exported functions/types
   - Use meaningful, descriptive names

3. **Comments**:
   - Add package documentation
   - Document all exported functions/types
   - Use complete sentences
   - Start with the name being documented

4. **Error handling**:
   - Always check and handle errors
   - Provide context in error messages
   - Use `fmt.Errorf` with `%w` for wrapping errors

5. **Project structure**:
   - Keep `internal/` for private code
   - Keep `pkg/` for reusable libraries
   - Follow clean architecture principles

### TypeScript/Vue Code

1. **Use TypeScript** for type safety
2. **Follow Vue 3 Composition API** best practices
3. **Use ESLint** and Prettier for consistency:
   ```bash
   npm run lint
   npm run format
   ```

4. **Component structure**:
   - Use `<script setup>` syntax
   - Define props with TypeScript
   - Extract reusable logic to composables

5. **Naming conventions**:
   - Use `PascalCase` for components
   - Use `camelCase` for functions/variables
   - Use `kebab-case` for file names

### General Principles

- **KISS**: Keep It Simple, Stupid
- **DRY**: Don't Repeat Yourself
- **YAGNI**: You Aren't Gonna Need It
- **Write self-documenting code**
- **Prefer readability over cleverness**
- **Test your code**

---

## Commit Message Guidelines

We follow the [Conventional Commits](https://www.conventionalcommits.org/) specification.

### Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

- **feat**: New feature
- **fix**: Bug fix
- **docs**: Documentation changes
- **style**: Code style changes (formatting, etc.)
- **refactor**: Code refactoring
- **perf**: Performance improvements
- **test**: Adding or updating tests
- **chore**: Maintenance tasks
- **ci**: CI/CD changes

### Examples

```
feat(auth): add JWT token refresh mechanism

Implement automatic token refresh to improve user experience
and reduce login prompts.

Fixes #123
```

```
fix(server): resolve memory leak in metrics collector

The metrics collector was not properly closing SSH connections,
causing memory to accumulate over time.

Closes #456
```

```
docs(api): update authentication endpoint examples

Add more examples and clarify token expiration behavior.
```

### Guidelines

- Use imperative mood ("add" not "added")
- Don't capitalize first letter of subject
- No period at the end of subject
- Limit subject line to 50 characters
- Wrap body at 72 characters
- Reference issues in footer

---

## Testing

### Writing Tests

1. **Unit tests**: Test individual functions/methods
2. **Integration tests**: Test component interactions
3. **E2E tests**: Test complete user workflows

### Test Guidelines

- **Write tests first** (TDD) when possible
- **Test edge cases** and error conditions
- **Keep tests focused** and independent
- **Use descriptive test names**
- **Mock external dependencies**
- **Aim for high coverage** (>80%)

### Running Tests

```bash
# Backend tests
make test

# Frontend tests
cd web && npm run test

# Integration tests
make test-integration

# All tests
make test-all
```

---

## Documentation

### What to Document

- **Code**: Functions, types, complex logic
- **APIs**: Endpoints, parameters, responses
- **Features**: How to use new functionality
- **Setup**: Installation and configuration
- **Architecture**: System design decisions

### Documentation Guidelines

1. **Keep it up to date** with code changes
2. **Use clear, simple language**
3. **Provide examples** where helpful
4. **Include diagrams** for complex concepts
5. **Support multiple languages** (English and Chinese)

### Documentation Files

- `README.md` - Project overview and quick start
- `docs/en-US/` - English documentation
- `docs/zh-CN/` - Chinese documentation
- Code comments - Inline documentation

---

## Questions?

If you have questions or need help:

1. Check existing documentation
2. Search closed issues
3. Open a new issue with the question label
4. Join our community discussions (if available)

---

## License

By contributing to NexusPanel, you agree that your contributions will be licensed under the AGPL-3.0 license.

---

**Thank you for contributing to NexusPanel!**
