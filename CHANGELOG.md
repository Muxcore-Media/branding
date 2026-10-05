# Changelog

## [0.1.0] - 2026-10-05

### Changed
- CI runs on GitHub-hosted runners from the umbrella template; retired-origin workflows removed.
- Dependencies resolve from published GitHub tags (no filesystem `replace`); requires core v0.6.0.

## 0.1.0 — 2026-09-05

- Initial bootstrap: Go `pkg/branding` package with brand and theme constants
- Embedded placeholder SVG logos (`logo.svg`, `logo-mark.svg`)
- `tokens/theme.json` for non-Go consumers
- Asset naming contract in `assets/README.md`
