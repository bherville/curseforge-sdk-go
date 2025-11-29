# Changesets

This folder contains changesets - files that describe changes to the project.

## Adding a Changeset

When you make a change that should be released, run:

```bash
npx changeset
```

This will prompt you to:
1. Select the type of change (patch, minor, major)
2. Write a summary of the change

The changeset file will be created in this folder and should be committed with your PR.

## Changeset Types

- **patch**: Bug fixes, small improvements (0.0.X)
- **minor**: New features, non-breaking changes (0.X.0)
- **major**: Breaking changes (X.0.0)

## How it Works

1. Contributors add changesets with their PRs
2. When PRs are merged, the changesets accumulate
3. The "Version Packages" PR is automatically created/updated
4. When that PR is merged, versions are bumped and tags are created

## Example Changeset

```markdown
---
"curseforge-sdk-go": minor
---

Added new endpoint for fetching mod categories
```
