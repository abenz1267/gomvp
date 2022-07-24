# gomvp - move packages, rename modules

gomvp is a simple CLI tool that helps you refactor packages and rename modules.

## Usage

`gomvp <src pkg> <dst pkg>`, for example `gomvp somepackage movedpackage`

If the src is equivalent to the current module name, then instead the module is renamed, so f.e. `gomvp github.com/abenz1267/gomvp github.com/abenz1267/newmodule` will rename the module.
