# Go Install With Github Releases Guide
This is a brief guide of how to create, build and release a go app to be installed with `go install` using a github release.

## Prerequisites
1. Make sure your `$GOPATH` environment variable exists
2. Make sure the `$GOPATH/bin` directory exists and its part of the `$PATH` environment variable. This is required to the bin file to be used in your terminal.

## To create the bin files as a part of a github release
1. First, it creates a new release everytime that a new tag is pushed using: [softprops/action-gh-release](https://github.com/softprops/action-gh-release).
2. Then it creates the binary file using [ wangyoucao577/go-release-action](https://github.com/wangyoucao577/go-release-action) and attach that (or those) bin file to the release that was just created.

## Install

### With `go install`

```bash
$ go install github.com/juancsr/go-install-with-release/cmd/sky@latest
```

### Releases binary

Go to the releases page and downalod the corresponding binary file for your OS and Arch.