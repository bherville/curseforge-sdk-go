---
"curseforge-sdk-go": patch
---

Remove duplicate release job from build workflow

- Release is now handled exclusively by auto-release.yml with changesets
- Prevents duplicate workflow runs on PR merge
