package pytest

import (
	"github.com/lottetian/hrp/internal/builtin"
	"github.com/lottetian/hrp/internal/sdk"
)

func RunPytest(args []string) error {
	sdk.SendEvent(sdk.EventTracking{
		Category: "RunAPITests",
		Action:   "hrp pytest",
	})

	python3, err := builtin.EnsurePython3Venv("httprunner")
	if err != nil {
		return err
	}

	args = append([]string{"-m", "httprunner", "run"}, args...)
	return builtin.ExecCommand(python3, args...)
}
