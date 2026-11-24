# Changelog

## 0.2.0 - [UNRELEASED]

### Breaking changes

- Books now can only be deleted if its not added to a collection by any user

### Added

- Shelfs
  - Users can add a book to their shelfs:
    - Want to Read
    - Currently Reading
    - Backlogged
    - Completed
  - Each book can only be at one shelf at a time
  - Books added by the user can be listed and filtered in a separate page
- Ability to drag'n'drop files from any source, not just from local filesystem
- Ability to update book covers
- Titles for browser tabs
- SHELFD_SESSION_SECURE configuration option (see fixed section)

### Fixed

- Authentication was only possible on encrypted connections. A new configuration option was added to disable the issuing of secure cookies (see .env.example)
- Files that were drag'n'dropped were not uploaded
- Book description filed not flagged as optional on the UI
- Sidebar not responsize on mobile devices
- [Internal] Invalid routes after a trailing slash still registered as a valid route
- [Internal] 404 routes rerouted the user to an empty Svelte SPA page

## [0.1.0] - 2026. 11. 10.

### Added

- User authentication, with multi user support per instance
- Book creation, update and deletion

[0.1.0]: https://github.com/Gilbike/shelfd/commits/v0.1.0
