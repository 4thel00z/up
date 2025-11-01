# GitHub Actions Workflows

This directory contains GitHub Actions workflows for automated CI/CD.

## Workflows

### CI (`ci.yml`)

**Triggers:**
- Push to `master` or `main` branches
- Pull requests to `master` or `main` branches

**What it does:**
1. **Test Job** (runs on Ubuntu and macOS):
   - Sets up Go 1.23
   - Installs required build tools (PEG parser, go-bindata)
   - Downloads and verifies dependencies
   - Generates code from templates
   - Builds all packages
   - Runs tests with race detection
   - Generates code coverage report
   - Uploads coverage to Codecov

2. **Lint Job** (runs on Ubuntu):
   - Sets up Go 1.23
   - Installs build tools
   - Generates code
   - Runs golangci-lint for code quality checks

### Release (`release.yml`)

**Triggers:**
- Push of tags matching `v*` (e.g., `v2.0.0`)

**What it does:**
1. Sets up Go 1.23
2. Installs build tools
3. Generates code
4. Runs GoReleaser to:
   - Build binaries for multiple platforms (Linux, macOS, Windows, BSD)
   - Create archives (tar.gz for Unix, zip for Windows)
   - Generate checksums
   - Create GitHub release with changelog
   - Upload release assets

## Local Testing with Act

You can test these workflows locally using [act](https://github.com/nektos/act).

### Installation

```bash
# macOS
brew install act

# Linux
curl https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash

# Or download from releases
# https://github.com/nektos/act/releases
```

### Usage

The project includes an `.actrc` configuration file at the root for act settings.

```bash
# Run default workflow (push event)
act

# Run pull request workflow
act pull_request

# List available workflows
act -l

# Run specific job
act -j test
act -j lint

# Dry run (don't execute, just show what would run)
act -n

# Run with verbose output
act -v
```

### Environment Setup

Act uses the settings from `.actrc`:
- Uses Docker images compatible with GitHub Actions
- Pulls `GITHUB_TOKEN` from your environment
- Can use `.secrets` file for additional secrets

**Create `.secrets` file (optional):**
```bash
GITHUB_TOKEN=your_token_here
CODECOV_TOKEN=your_codecov_token_here
```

**Or export environment variable:**
```bash
export GITHUB_TOKEN=$(gh auth token)  # If using GitHub CLI
```

## Making Changes

When adding new workflows:
1. Create or modify `.yml` files in this directory
2. Test locally with `act` before pushing
3. Ensure required secrets are documented

## Required Secrets

- `GITHUB_TOKEN`: Automatically provided by GitHub Actions (for releases)
- `CODECOV_TOKEN`: Optional, for code coverage reporting

## Platform Support

Built and tested on:
- Linux (amd64, arm64, 386)
- macOS (amd64, arm64)
- Windows (amd64, 386)
- FreeBSD, NetBSD, OpenBSD (amd64, 386)

