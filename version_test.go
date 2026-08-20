package clic4go_test

import (
	clic4go "github.com/synesissoftware/CliC4.Go"

	"github.com/stretchr/testify/require"

	"testing"
)

const (
	Expected_VersionMajor uint16 = 0
	Expected_VersionMinor uint16 = 0
	Expected_VersionPatch uint16 = 1
	Expected_VersionAB    uint16 = 0xFFFF
)

func Test_Version_Elements(t *testing.T) {
	require.Equal(t, Expected_VersionMajor, clic4go.VersionMajor)
	require.Equal(t, Expected_VersionMinor, clic4go.VersionMinor)
	require.Equal(t, Expected_VersionPatch, clic4go.VersionPatch)
	require.Equal(t, Expected_VersionAB, clic4go.VersionAB)
}

func Test_Version(t *testing.T) {
	require.Equal(t, uint64(0x0000_0000_0001_FFFF), clic4go.Version())
}

func Test_Version_String(t *testing.T) {
	require.Equal(t, "0.0.1", clic4go.VersionString())
}
