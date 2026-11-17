# Changelog

## 0.2.0 - [UNRELEASED]

### Breaking changes

- Books now can only be deleted if its not added to a collection by any user

### Added

- Collections
  - Users can add a book to the following collections:
    - Want to Read
    - Currently Reading
    - Backlogged
    - Completed
  - Books added by the user can be listed and filtered in a separate page
- Ability to drag'n'drop files from any source, not just from local filesystem
- Ability to update book covers
- Titles for browser tabs

### Fixed

- Files that were drag'n'dropped were not uploaded
- Book description filed not flagged as optional on the UI
- Sidebar not responsize on mobile devices

## [0.1.0] - 2026. 11. 10.

### Added

- User authentication, with multi user support per instance
- Book creation, update and deletion

[0.1.0]: https://github.com/Gilbike/shelfd/commits/v0.1.0
