# Release Process

1. Use [go-licenses](https://github.com/google/go-licenses) to ensure all project dependency licenses are correctly represented in this repository:
    - Install go-licenses (if not already installed) `go install github.com/google/go-licenses@latest`
    - Run and save licenses `go-licenses save github.com/honeyvent --save_path="./LICENSES"`
    - If there are any changes, submit PR to update licenses.
1. Update `CHANGELOG.md` with the changes since the last release.
1. Commit changes, push, and open a release preparation PR for review.
1. Once the pull request is merged, fetch the updated main branch.
1. Apply a tag for the new version on the merged commit (e.g. `git tag -a v1.4.1 -m "v1.4.1"`)
1. Push the tag upstream to kick off the release pipeline in CI (e.g. `git push origin v1.4.1`).
1. Ensure that there is a draft GitHub release created as part of CI publish steps
1. Edit the draft GitHub release: click the Generate Release Notes button and double-check the content against the CHANGELOG.
1. Publish the GitHub release
