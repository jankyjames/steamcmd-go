package steamcmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppInfo_Do(t *testing.T) {
	steamCmd := New()
	man, err := steamCmd.GetAppInfo(AppIDArkSurvivalAscendedDedicatedServer)
	assert.NoError(t, err)
	assert.Equal(t, man.Common.Name, "ARK: Survival Ascended Dedicated Server")
	assert.Equal(t, man.Common.Gameid, "2430930")
	assert.Greater(t, len(man.Depots.Branches), 0)
	assert.Greater(t, len(man.GetBuildIDForBranch("public")), 0)
}
