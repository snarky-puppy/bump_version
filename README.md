# bump_version

This is a tool for bumping version numbers in Go files.

The current repo is a fork of github.com/Shyp/bump_version .

Changes:
 - go mod support
 - add a 'v' prefix to the generated version (https://go.dev/ref/mod#versions)
 - changes commit message to "Version %s [skip ci]" for use in github, circleci, etc.

Planned changes:
 - check current commit log message for "fix, feat" prefix and automatically bump minor and patch versions.

## Installation

For the moment, you'll need a working Go installation.

```
go get github.com/snarky-puppy/bump_version
```

That will install the `bump_version` binary to your `$GOPATH`.

## Usage

```
bump_version <major|minor|patch> <filename>
```

This will:

1. Look for a `const` named `version`, `VERSION`, or `Version` in that file.
   Here's an example:

    ```go
    package main

    const VERSION = "v0.2.1"
    ```

2. Apply the version bump - `bump_version major` will increment the major
version number, `bump_version minor` will increment the middle version number,
`bump_version patch` will increment the last version number. If your version is
"v0.3" and you ask for `bump_version minor`, the new version will be "v0.4".

3. Write the new file to disk, with the bumped version.

4. Add the file with `git add <filename>`.

5. Add a commit with the message "Version vX.Y.Z [skip ci]"

6. Tag the new version.

If any of these steps fail, `bump_version` will abort.

## Notes

The VERSION should be a string in one of these formats: "3", "0.3",
"0.3.4". Any suffixes like "0.3.3-beta" will be
stripped or generate an error.

- `bump_version("0.1", "minor")` -> "v0.2"
- `bump_version("0.1", "patch")` -> "v0.1.1"
- `bump_version("0.1", "major")` -> "v1.1"
- `bump_version("0.1-beta", "major")` -> "v1.1"
- `bump_version("devel", "major")` -> parse error.

We use the VERSION in code exclusively - any existing git tags are ignored.

Alan Shreve would like to note that you probably shouldn't store version
numbers in code - instead, check in `const VERSION = "devel"`, then build your
project via:

```
go build -ldflags="-X main.VERSION=0.2"
```

Which you are welcome to do!
