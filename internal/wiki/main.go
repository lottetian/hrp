package wiki

import (
	"os/exec"

	"github.com/rs/zerolog/log"

	"github.com/lottetian/hrp/internal/sdk"
)

func OpenWiki() error {
	sdk.SendEvent(sdk.EventTracking{
		Category: "OpenWiki",
		Action:   "hrp wiki",
	})
	log.Info().Msgf("%s https://httprunner.com", openCmd)
	return exec.Command(openCmd, "https://httprunner.com").Run()
}
