# gomvp - move packages, rename modules

gomvp is a simple CLI tool that helps you refactor packages and rename modules.

## Installation

`go install github.com/abenz1267/gomvp@latest`

... or download the binary from the releases.

## Usage

`gomvp <src pkg> <dst pkg>`, for example `gomvp somepackage movedpackage`

If the src is equivalent to the current module name, then instead the module is renamed, so f.e. `gomvp github.com/abenz1267/gomvp github.com/abenz1267/newmodule` will rename the module.

Moving files is performed via `git mv`, so they have to be added via git before usage.

## Notes

-   will abort if there's a conflicting package present
-   named imports won't be altered, as it's just replacing the actual import string
-   usage of named imports won't be altered, as it's just replacing the package name
