# honeyvent

[![OSS Lifecycle](https://img.shields.io/osslifecycle/honeycombio/honeyvent?color=success)](https://github.com/honeycombio/home/blob/main/honeycomb-oss-lifecycle-and-practices.md)

**STATUS: this project is being sunset.** See https://github.com/honeycombio/honeyvent/issues/75

CLI for sending individual events in to [Honeycomb](https://docs.honeycomb.io)

## Installation

If you have a working go environment in your build, the easiest way to install `honeyvent` is via `go get`.

```
go get github.com/honeycombio/honeyvent/
```

If you're using go 1.17 or greater you should use `go install` instead, since `go get` is [deprecated](https://go.dev/doc/go-get-install-deprecation).

```
go install github.com/honeycombio/honeyvent@latest
```

## Usage

Call with a collection of names and values to send an event from the
command line:

```
honeyvent -k <writekey> -d <dataset> -n field -v val -n field -v val ...
```

If you are targeting a local instance of Honeycomb, use the `api_host` parameter, e.g: `--api_host=http://localhost:8888`

The tool will detect floats and ints and send them as numbers; everything else
turns in to strings.  Quote any values that have spaces.
