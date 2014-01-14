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
    INVALID_VERSION         = "zxv8d"
    EMPTY_VERSION           = ""
)

func TestNewSemver(t *testing.T) {
    s, err := NewVersion(BASE_VERSION)
    if err != nil {
        t.Fatalf("Unable to allocate new Version object: %s", err)
    }

    if s == nil {
        t.Fatalf("Unable to allocate new version.")
    }

    _, err2 := NewVersion(INVALID_VERSION)
    if err2 == nil {
        t.Fatalf("Should throw error on invalid version.")
    }

    _, err3 := NewVersion(EMPTY_VERSION)
    if err3 == nil {
        t.Fatalf("Should throw error on empty version.")
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

func TestNonExplicitPrereleaseVersion(t *testing.T) {
    s, err := NewVersion(BASE_VERSION)
    if err != nil {
        t.Fatalf("Unable to allocate new Version object: %s.", err)
    }

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

func TestMetadata(t *testing.T) {
    s, err := NewVersion(BASE_VERSION)
    if err != nil {
        t.Fatalf("Unable to allocate new Version object: %s.", err)
    }

    s.SetMetadata("dangerzone")

    if s.Metadata != "dangerzone" {
        t.Fatalf("Unable to set Metadata to 'dangerzone'")
    }

    s.SetMetadata("")

    if s.Metadata != "build" {
        t.Fatalf("Unable to default Metadata to 'build'")
    }
}

func TestPrerelease(t *testing.T) {
    s, err := NewVersion(BASE_VERSION)
    if err != nil {
        t.Fatalf("Unable to allocate new Version object: %s.", err)
    }

    s.SetPrerelease("beta")

    if s.PrereleaseType != "beta" {
        t.Fatalf("Unable to set PrereleaseType to 'beta'")
    }

    s.SetPrerelease("")

    if s.PrereleaseType != "alpha" {
        t.Fatalf("Unable to default PrereleaseType to 'alpha'")
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
        t.Fatalf("Metadata should be build, got %s", s.Metadata)
    }

    if s.MetadataCount != 1 {
        t.Fatalf("MetadataCount should be 1, got %d", s.MetadataCount)
    }
}

func TestEquals(t *testing.T) {
    var versions = []string{
        "1.0.1",                // Patch
        "1.1.0",                // Minor
        "2.0.0",                // Major
        "1.0.0-alpha.1",        // Prerelease
        "1.0.0-beta.1+build.3", // Prerelease + Metadata
    }

    for i := range versions {
        v1, _ := NewVersion(versions[i])
        v2, _ := NewVersion(versions[i])

        if !v1.Equals(v2) {
            t.Fatalf("v1 (%s) should equal v2 (%s).", v1.String(), v2.String())
        }
    }

    v1, _ := NewVersion(BASE_VERSION)

    for i := range versions {
        v2, _ := NewVersion(versions[i])

        if v1.Equals(v2) {
            t.Fatalf("v1 (%s) should not equal v2 (%s).", v1.String(), v2.String())
        }
    }

    for i := range versions {
        for j := len(versions) - 1; j >= 0; j-- {
            if i != j {
                v1, _ := NewVersion(versions[i])
                v2, _ := NewVersion(versions[j])

                if v1.Equals(v2) {
                    t.Fatalf("v1 (%s) should not equal v2 (%s).", v1.String(), v2.String())
                }
            }
        }
    }
}

func TestEqualsWithMetadata(t *testing.T) {
    v1, _ := NewVersion(BASE_VERSION)
    v2, _ := NewVersion("1.0.0+build.1")

    if !v1.Equals(v2) {
        t.Fatalf("v1 (%s) should equal v2 (%s) with differing metadata.", v1.String(), v2.String())
    }
}

func TestComparison(t *testing.T) {
    var versions = []string{
        "0.0.9",         // Patch release
        "0.1.0",         // Minor release
        "1.0.0-alpha.1", // Prerelease alpha
        "1.0.0-beta.1",  // Prerelease beta
        "1.0.0",         // Major release
        "1.0.1-rc.1",    // Release candidate
        "1.0.1-rc.2",    // Release candidate
        "1.0.1",         // Patch
        "1.1.0",         // Minor
        "2.0.0",         // Major

    }

    for i := range versions {
        for j := len(versions) - 1; j >= 0; j-- {
            v1, _ := NewVersion(versions[i])
            v2, _ := NewVersion(versions[j])

            if i < j {
                if !v1.LessThan(v2) {
                    t.Fatalf("v1 (%s) should be less than v2 (%s).", v1.String(), v2.String())
                }
                if v2.LessThan(v1) {
                    t.Fatalf("v1 (%s) should be less than v2 (%s).", v1.String(), v2.String())
                }
            } else if i == j {
                if !v1.Equals(v2) {
                    t.Fatalf("v1 (%s) should equal v2 (%s).", v1.String(), v2.String())
                }
                if !v2.Equals(v1) {
                    t.Fatalf("v1 (%s) should equal v2 (%s).", v1.String(), v2.String())
                }
            } else if i > j {
                if !v1.GreaterThan(v2) {
                    t.Fatalf("v1 (%s) should be greater than v2 (%s).", v1.String(), v2.String())
                }
                if v2.GreaterThan(v1) {
                    t.Fatalf("v1 (%s) should be greater than v2 (%s).", v1.String(), v2.String())
                }
            }
        }
    }
}
