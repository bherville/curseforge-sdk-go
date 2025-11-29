---
"curseforge-sdk-go": minor
---

Add changesets for automated version management and CI

- Integrate changesets CLI for semver-based version management
- Add GitHub Actions workflow with build, test, and release jobs
- Run tests with race detector for concurrency safety
- Tags follow format `v{SEMVER}` (e.g., `v0.1.0`)
