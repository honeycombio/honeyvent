# Release Process

1. Use [go-licenses](https://github.com/google/go-licenses) to ensure all project dependency licenses are correclty represented in this repository:
  1. Install go-licenses (if not already installed) `go install github.com/google/go-licenses@latest`
  2. Run and save licenses `go-licenses save github.com/honeycombio/buildevents --save_path="./LICENSES"`
  3. If there are any changes, submit PR to update licenses.
2. Add release entry to [changelog](./CHANGELOG.md)
3. Open a PR with the above, and merge that into main
4. Create new tag on merged commit with the new version (e.g. `v1.4.1`)
5. Push the tag upstream (this will kick off the release pipeline in CI)
6. Copy change log entry for newest version into draft GitHub release created as part of CI publish steps
