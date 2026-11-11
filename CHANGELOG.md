# Changelog

## v1.0.0 - 2026-11-11

### ✨ Features
- Native L4D2 mod manager for Linux (and Windows).
- Automatic parsing of both Workshop and local VPKs upon selecting the L4D2 directory.
- Auto-generated primary and secondary categories via the `Tools` page, with dynamic VPK counts per category.
- Customizable VPK display and sorting rules.
- Toggle switches to easily enable or disable VPKs.
- Batch selection tools (select/deselect up to 50 VPKs per page).
- View detailed metadata for individual VPKs.

### 🐛 Known Issues
- **Local VPK images fail to display:** Due to a limitation in the upstream UI framework's `Image` component. As a workaround, place a fallback image at `assets/imgs/l4d2.png` next to the executable.

### ⚠️ Note for Windows Users
- Windows Defender / SmartScreen may flag the executable because it is unsigned. Click **"More info" → "Run anyway"** if you trust the release, or build from source.