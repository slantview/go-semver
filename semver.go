// Package semver provides a semantic versioning library for Go.
package semver

import (
    "errors"
    "fmt"
    "regexp"
    "strconv"
)

const (
    MAJOR_VERSION      = 1
    MINOR_VERSION      = 2
    PATCH_VERSION      = 4
    PRERELEASE_VERSION = 8
    BUILD_VERSION      = 16
    VERSION_FORMAT     = "^(?P<major>\\d+)\\.(?P<minor>\\d+)\\.(?P<patch>\\d+)(?:-(?P<prerelease>[0-9A-Za-z-]+)(?:\\.(?P<prerelease_count>[0-9A-Za-z-]+))*)?(?:\\+(?P<metadata>[0-9A-Za-z-]+)\\.?(?P<metadata_count>[0-9A-Za-z-]+)*)?$"
)

type Version struct {
    Major           int64
    Minor           int64
    Patch           int64
    PrereleaseType  string
    PrereleaseCount int64
    Metadata        string
    MetadataCount   int64
}

// Create a new version object from a string.
func NewVersion(version string) (*Version, error) {
    v := &Version{0, 0, 1, "", 0, "", 0}

    err := v.ParseString(version)
    if err != nil {
        return nil, err
    }

    return v, nil
}

// Reset resets the version to a default version of 0.0.1.
func (v *Version) Reset() {
    v.Major = 0
    v.Minor = 0
    v.Patch = 1
    v.PrereleaseType = ""
    v.PrereleaseCount = 0
    v.Metadata = ""
    v.MetadataCount = 0
}

// Bump increases the version by a specified version type.  There are helper
// functions for bumping specific types as well.
func (v *Version) Bump(bump int) error {
    switch bump {
    case MAJOR_VERSION:
        v.Major += 1
        v.Minor = 0
        v.Patch = 0
    case MINOR_VERSION:
        v.Minor += 1
        v.Patch = 0
    case PATCH_VERSION:
        v.Patch += 1
    case PRERELEASE_VERSION:
        v.PrereleaseCount += 1
    case BUILD_VERSION:
        v.MetadataCount += 1
    }

    return nil
}

// Bumps the Major number of the version.
func (v *Version) BumpMajor() error {
    return v.Bump(MAJOR_VERSION)
}

// Bumps the Minor number of the version.
func (v *Version) BumpMinor() error {
    return v.Bump(MINOR_VERSION)
}

// Bumps the Patch number of the version.
func (v *Version) BumpPatch() error {
    return v.Bump(PATCH_VERSION)
}

// Bumps the Pre-release number of the version.  Additionally if you don't set
// the pre-release type (e.g. alpha, beta, etc), it will default the pre-release
// type to alpha.
func (v *Version) BumpPrerelease() error {
    if v.PrereleaseType == "" {
        v.SetPrerelease("alpha")
    }
    return v.Bump(PRERELEASE_VERSION)
}

// Bumps the build version by setting the metadata to "build" and increasing the
// metadata count.
func (v *Version) BumpBuild() error {
    v.SetMetadata("build")
    return v.Bump(BUILD_VERSION)
}

// Sets the metadata type.  Default metadata type is build, but any ascii
// alphanumeric data may be used.
func (v *Version) SetMetadata(metadata string) {
    if metadata != "" {
        v.Metadata = metadata
    } else {
        v.Metadata = "build"
    }

    if v.MetadataCount == 0 {
        v.MetadataCount = 1
    }
}

// Sets the prerelease type.  Should be alpha, beta, etc.  Defaults to "alpha"
func (v *Version) SetPrerelease(prerelease string) {
    if prerelease != "" {
        v.PrereleaseType = prerelease
    } else {
        v.PrereleaseType = "alpha"
    }

    if v.PrereleaseCount == 0 {
        v.PrereleaseCount = 1
    }
}

// Parses a version string.  Returns error on invalid version.
func (v *Version) ParseString(version string) error {
    re := regexp.MustCompile(VERSION_FORMAT)

    result := re.FindStringSubmatch(version)
    if len(result) == 0 {
        return errors.New(fmt.Sprintf("Unable to parse version string: %s", version))
    } else {
        names := re.SubexpNames()

        for i := 1; i < len(result); i++ {
            m, n := result[i], names[i]

            switch n {
            case "major":
                v.Major, _ = strconv.ParseInt(m, 10, 0)
            case "minor":
                v.Minor, _ = strconv.ParseInt(m, 10, 0)
            case "patch":
                v.Patch, _ = strconv.ParseInt(m, 10, 0)
            case "prerelease":
                if m != "" {
                    v.SetPrerelease(m)
                }
            case "prerelease_count":
                if m != "" {
                    v.PrereleaseCount, _ = strconv.ParseInt(m, 10, 0)
                }
            case "metadata":
                if m != "" {
                    v.SetMetadata(m)
                }
            case "metadata_count":
                if m != "" {
                    v.MetadataCount, _ = strconv.ParseInt(m, 10, 0)
                }
            }
        }
    }
    return nil
}

// Returns a string representation of the version.
func (v *Version) String() string {
    out := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)

    if v.PrereleaseType != "" {
        out += fmt.Sprintf("-%s.%d", v.PrereleaseType, v.PrereleaseCount)
    }
    if v.Metadata != "" {
        out += fmt.Sprintf("+%s.%d", v.Metadata, v.MetadataCount)
    }

    return out
}

// Returns a bool of whether this version is less than the version argument.
func (v *Version) LessThan(v2 *Version) bool {
    var major = (v.Major == v2.Major)
    var minor = (v.Minor == v2.Minor)
    var patch = (v.Patch == v2.Patch)
    var base = (major && minor && patch)

    if v.Major > v2.Major {
        return false
    } else if major && v.Minor > v2.Minor {
        return false
    } else if major && minor && v.Patch > v2.Patch {
        return false
    } else if base && (v.PrereleaseType == "" && v2.PrereleaseType != "") {
        return false
    } else if base && (v.PrereleaseType > v2.PrereleaseType && (v.PrereleaseType != "" && v2.PrereleaseType != "")) {
        return false
    } else if base && (v.PrereleaseType == v2.PrereleaseType && v.PrereleaseCount > v2.PrereleaseCount) {
        return false
    } else {
        return true
    }
}

// Returns a bool of whether this version is greater than the version argument.
func (v *Version) GreaterThan(v2 *Version) bool {
    if !v.LessThan(v2) && !v.Equals(v2) {
        return true
    } else {
        return false
    }
}

// Returns a bool of whether this version is equal to the version argument.
func (v *Version) Equals(v2 *Version) bool {
    if (v.Major == v2.Major && v.Minor == v2.Minor && v.Patch == v2.Patch) &&
        (v.PrereleaseType == v2.PrereleaseType && v.PrereleaseCount == v2.PrereleaseCount) {
        return true
    } else {
        return false
    }
}
