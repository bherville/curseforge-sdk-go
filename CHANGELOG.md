# curseforge-sdk-go

## 0.2.1

### Patch Changes

- e0df944: Remove duplicate release job from build workflow

  - Release is now handled exclusively by auto-release.yml with changesets
  - Prevents duplicate workflow runs on PR merge

## 0.2.0

### Minor Changes

- e58d90b: Add changesets for automated version management and releases

  - Integrate changesets for version bumping and changelog generation
  - Add auto-release workflow that triggers on merged PRs with changesets
  - Add changeset check workflow to enforce changesets on PRs

## 0.1.0

### Minor Changes

- 483a155: Add changesets for automated version management and CI

  - Integrate changesets CLI for semver-based version management
  - Add GitHub Actions workflow with build, test, and release jobs
  - Run tests with race detector for concurrency safety
  - Tags follow format `v{SEMVER}` (e.g., `v0.1.0`)
