package semver

import (
    "errors"
    "fmt"
    "regexp"
    "strconv"
)

var (
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

func NewVersion(version string) (*Version, error) {
    v := &Version{0, 0, 1, "", 0, "", 0}

    err := v.ParseString(version)
    if err != nil {
        return nil, err
    }

    return v, nil
}

func (v *Version) Reset() {
    v.Major = 0
    v.Minor = 0
    v.Patch = 1
    v.PrereleaseType = ""
    v.PrereleaseCount = 0
    v.Metadata = ""
    v.MetadataCount = 0
}

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

func (v *Version) BumpMajor() error {
    return v.Bump(MAJOR_VERSION)
}

func (v *Version) BumpMinor() error {
    return v.Bump(MINOR_VERSION)
}

func (v *Version) BumpPatch() error {
    return v.Bump(PATCH_VERSION)
}

func (v *Version) BumpPrerelease() error {
    if v.PrereleaseType == "" {
        v.SetPrerelease("alpha")
    }
    return v.Bump(PRERELEASE_VERSION)
}

func (v *Version) BumpBuild() error {
    err := v.SetMetadata("build")
    if err != nil {
        return err
    }
    return v.Bump(BUILD_VERSION)
}

func (v *Version) SetMetadata(metadata string) error {
    if metadata != "" {
        v.Metadata = metadata
    } else {
        v.Metadata = "build"
    }

    if v.MetadataCount == 0 {
        v.MetadataCount = 1
    }
    return nil
}

func (v *Version) SetPrerelease(prerelease string) error {
    if prerelease != "" {
        v.PrereleaseType = prerelease
    } else {
        v.PrereleaseType = "alpha"
    }

    if v.PrereleaseCount == 0 {
        v.PrereleaseCount = 1
    }

    return nil
}

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