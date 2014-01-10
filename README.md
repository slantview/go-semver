# go-semver

A library for using (semantic versioning)[http://semver.org/] in Go.

## Example Usage

```go

s := semver.NewVersion("1.0.0")

fmt.Printf("%s", s.String())
// "1.0.0"

s.BumpPatch()
fmt.Printf("%s", s.String())
// "1.0.1"

s.BumpMinor()
fmt.Printf("%s", s.String())
// "1.1.0"

s.BumpMajor()
fmt.Printf("%s", s.String())
// "2.0.0"

s.SetPrerelease("alpha")
fmt.Printf("%s", s.String())
// "2.0.0-alpha.1"

s.BumpPrerelease()
fmt.Printf("%s", s.String())
// "2.0.0-alpha.2"

s.SetBuild("build")
fmt.Printf("%s", s.String())
// "2.0.0+build.1"

s.BumpBuild()
fmt.Printf("%s", s.String())
// "2.0.0+build.2"

```
