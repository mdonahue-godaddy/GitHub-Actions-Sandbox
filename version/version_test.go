package version_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/mdonahue-godaddy/GitHub-Actions-Sandbox/version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Setup Suite.
type VersionSuite struct {
	suite.Suite
}

func TestVersionSuite(t *testing.T) {
	versionSuite := VersionSuite{} //nolint:exhaustruct  // This is then normal way to instantiate a suite
	suite.Run(t, &versionSuite)
}

type testGetVersion struct {
	Description string
	AppVersion  string
	Branch      string
	BuildTime   string
	Commit      string
	GoVersion   string
}

func createGetVersionTestData() []testGetVersion {
	tests := []testGetVersion{
		{
			Description: "All are empty strings",
			AppVersion:  "",
			Branch:      "",
			BuildTime:   "",
			Commit:      "",
			GoVersion:   "",
		},
		{
			Description: "All have values",
			AppVersion:  "AppVersion",
			Branch:      "Branch",
			BuildTime:   "BuildTime",
			Commit:      "Commit",
			GoVersion:   "GoVersion",
		},
	}

	return tests
}

func (s *VersionSuite) TestGetLongVersion() {
	for _, tst := range createGetVersionTestData() {
		expected := fmt.Sprintf(
			"%s version: [%s]\n- Branch:     [%s]\n- Build Time: [%s]\n- Commit:     [%s]\n- Go Version: [%s]\n",
			os.Args[0], tst.AppVersion, tst.Branch, tst.BuildTime, tst.Commit, tst.GoVersion)

		AppVersion = tst.AppVersion
		Branch = tst.Branch
		BuildTime = tst.BuildTime
		Commit = tst.Commit
		GoVersion = tst.GoVersion

		actual := GetLongVersion()

		assert.Equal(s.T(), expected, actual, tst.Description+fmt.Sprintf(" expected '%s', actual '%s'", expected, actual))
	}
}

func (s *VersionSuite) TestGetShortVersion() {
	for _, tst := range createGetVersionTestData() {
		expected := tst.AppVersion

		AppVersion = tst.AppVersion
		Branch = tst.Branch
		BuildTime = tst.BuildTime
		Commit = tst.Commit
		GoVersion = tst.GoVersion

		actual := GetShortVersion()

		assert.Equal(s.T(), expected, actual, tst.Description+fmt.Sprintf(" expected '%s', actual '%s'", expected, actual))
	}
}

func (s *VersionSuite) TestGetVersion() {
	for _, tst := range createGetVersionTestData() {
		// override os.Args with test array
		os.Args = []string{"TestAppVersion"}

		expected := fmt.Sprintf("TestAppVersion version: %s", tst.AppVersion)

		AppVersion = tst.AppVersion
		Branch = tst.Branch
		BuildTime = tst.BuildTime
		Commit = tst.Commit
		GoVersion = tst.GoVersion
		actual := GetVersion()

		assert.Equal(s.T(), expected, actual, tst.Description+fmt.Sprintf(" expected '%s', actual '%s'", expected, actual))
	}
}

func (s *VersionSuite) TestGetVersionNoArgs() {
	for _, tst := range createGetVersionTestData() {
		// override os.Args with empty array
		os.Args = []string{}

		expected := fmt.Sprintf("version: %s", tst.AppVersion)

		AppVersion = tst.AppVersion
		Branch = tst.Branch
		BuildTime = tst.BuildTime
		Commit = tst.Commit
		GoVersion = tst.GoVersion

		actual := GetVersion()

		assert.Equal(s.T(), expected, actual, tst.Description+fmt.Sprintf(" expected '%s', actual '%s'", expected, actual))
	}
}
