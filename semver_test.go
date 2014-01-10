package semver

import (
    "testing"
)

var (
    BASE_VERSION            = "1.0.0"
    BUMP_MAJOR_VERSION      = "2.0.0"
    BUMP_MINOR_VERSION      = "1.1.0"
    BUMP_PATCH_VERSION      = "1.0.1"
    BUMP_PRERELEASE_VERSION = "1.0.0-alpha.2"
    BUMP_BUILD_VERSION      = "1.0.0+build.2"
    ALL_FIELDS_VERSION      = "1.0.0-alpha.1+build.1"
)

func TestNewSemver(t *testing.T) {
    s, err := NewVersion(BASE_VERSION)
    if err != nil {
        t.Fatalf("Unable to allocate new Version object: %s", err)
    }

    if s == nil {
        t.Fatalf("Unable to allocate new version.")
    }
}

func TestSemverString(t *testing.T) {
    s, err := NewVersion(BASE_VERSION)
    if err != nil {
        t.Fatalf("Unable to allocate new Version object: %s.", err)
    }

    if s.String() != BASE_VERSION {
        t.Fatalf("Version doesn't match.  Got %s, looking for %s.", s.String(), BASE_VERSION)
    }

    s.Reset()
    s.ParseString(BUMP_MAJOR_VERSION)
    if s.String() != BUMP_MAJOR_VERSION {
        t.Fatalf("Version doesn't match.  Got %s, looking for %s.", s.String(), BUMP_MAJOR_VERSION)
    }

    s.Reset()
    s.ParseString(BUMP_MINOR_VERSION)
    if s.String() != BUMP_MINOR_VERSION {
        t.Fatalf("Version doesn't match.  Got %s, looking for %s.", s.String(), BUMP_MINOR_VERSION)
    }

    s.Reset()
    s.ParseString(BUMP_PATCH_VERSION)
    if s.String() != BUMP_PATCH_VERSION {
        t.Fatalf("Version doesn't match.  Got %s, looking for %s.", s.String(), BUMP_PATCH_VERSION)
    }

    s.Reset()
    s.ParseString(BUMP_PRERELEASE_VERSION)
    if s.String() != BUMP_PRERELEASE_VERSION {
        t.Fatalf("Version doesn't match.  Got %s, looking for %s.", s.String(), BUMP_PRERELEASE_VERSION)
    }

    s.Reset()
    s.ParseString(BUMP_BUILD_VERSION)
    if s.String() != BUMP_BUILD_VERSION {
        t.Fatalf("Version doesn't match.  Got %s, looking for %s.", s.String(), BUMP_BUILD_VERSION)
    }
}

func TestMajorVersion(t *testing.T) {
    s, err := NewVersion(BASE_VERSION)
    if err != nil {
        t.Fatalf("Unable to allocate new Version object: %s.", err)
    }

    s.BumpMajor()

    if s.String() != BUMP_MAJOR_VERSION {
        t.Fatalf("Unable to bump prerelease version.  Got %s, looking for %s.", s.String(), BUMP_MAJOR_VERSION)
    }
}

func TestMinorVersion(t *testing.T) {
    s, err := NewVersion(BASE_VERSION)
    if err != nil {
        t.Fatalf("Unable to allocate new Version object: %s.", err)
    }

    s.BumpMinor()

    if s.String() != BUMP_MINOR_VERSION {
        t.Fatalf("Unable to bump prerelease version.  Got %s, looking for %s.", s.String(), BUMP_MINOR_VERSION)
    }
}

func TestPatchVersion(t *testing.T) {
    s, err := NewVersion(BASE_VERSION)
    if err != nil {
        t.Fatalf("Unable to allocate new Version object: %s.", err)
    }

    s.BumpPatch()

    if s.String() != BUMP_PATCH_VERSION {
        t.Fatalf("Unable to bump prerelease version.  Got %s, looking for %s.", s.String(), BUMP_PATCH_VERSION)
    }
}

func TestPrereleaseVersion(t *testing.T) {
    s, err := NewVersion(BASE_VERSION)
    if err != nil {
        t.Fatalf("Unable to allocate new Version object: %s.", err)
    }

    s.SetPrerelease("alpha")
    s.BumpPrerelease()

    if s.String() != BUMP_PRERELEASE_VERSION {
        t.Fatalf("Unable to bump prerelease version.  Got %s, looking for %s.", s.String(), BUMP_PRERELEASE_VERSION)
    }
}

func TestBuildVersion(t *testing.T) {
    s, err := NewVersion(BASE_VERSION)
    if err != nil {
        t.Fatalf("Unable to allocate new Version object: %s.", err)
    }

    s.BumpBuild()

    if s.String() != BUMP_BUILD_VERSION {
        t.Fatalf("Unable to bump build version.  Got %s, looking for %s.", s.String(), BUMP_BUILD_VERSION)
    }
}

func TestAllFields(t *testing.T) {
    s, err := NewVersion(ALL_FIELDS_VERSION)
    if err != nil {
        t.Fatalf("Unable to allocate new Version object: %s.", err)
    }

    if s.String() != ALL_FIELDS_VERSION {
        t.Fatalf("Unable to parse version %s.", ALL_FIELDS_VERSION)
    }

    if s.Major != 1 {
        t.Fatalf("Major version should be 1, got %d", s.Major)
    }

    if s.Minor != 0 {
        t.Fatalf("Minor version should be 0, got %d", s.Minor)
    }

    if s.Patch != 0 {
        t.Fatalf("Patch version should be 0, got %d", s.Patch)
    }

    if s.PrereleaseType != "alpha" {
        t.Fatalf("PrereleaseType should be alpha, got %s", s.PrereleaseType)
    }

    if s.PrereleaseCount != 1 {
        t.Fatalf("PrereleaseCount should be 1, got %d", s.PrereleaseCount)
    }

    if s.Metadata != "build" {
        t.Fatalf("Metadata should be alpha, got %s", s.Metadata)
    }

    if s.MetadataCount != 1 {
        t.Fatalf("MetadataCount should be 1, got %d", s.MetadataCount)
    }
}