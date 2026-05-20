# Changelog

## v0.8.0 - 2026-05-20

- Upgrade `github.com/hashicorp/terraform-plugin-sdk/v2` from `v2.25.0` to `v2.40.1`.
- Upgrade `github.com/transloadit/go-sdk` from `v1.5.1` to `v1.6.0`.
- Upgrade vulnerable transitive dependencies, including `google.golang.org/grpc` to
  `v1.81.1`, `golang.org/x/crypto` to `v0.51.0`, and `golang.org/x/net` to `v0.54.0`.
- Raise the provider build and development Go requirement to `1.25.8+`; CI now tests
  Go `1.25.x` and `1.26.x`.
- Modernize CI and release automation by moving GitHub Actions to current majors and
  migrating the GoReleaser config to v2.
- Remove the obsolete `go test -i` make step for compatibility with modern Go.
