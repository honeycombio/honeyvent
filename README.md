# honeyvent

[![OSS Lifecycle](https://img.shields.io/osslifecycle/honeycombio/honeyvent)](https://github.com/honeycombio/home/blob/main/honeycomb-oss-lifecycle-and-practices.md)

CLI for sending individual events in to [Honeycomb](https://honeycomb.io/docs)

## Installation

If you have a working go environment in your build, the easiest way to install `honeyvent` is via `go get`.

```
go get github.com/honeycombio/honeyvent/
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
