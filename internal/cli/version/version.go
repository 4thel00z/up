package version

import (
	"fmt"

	"github.com/tj/kingpin"

	"github.com/4thel00z/up/internal/cli/root"
	"github.com/4thel00z/up/internal/stats"
)

func init() {
	cmd := root.Command("version", "Show version.")
	cmd.Action(func(_ *kingpin.ParseContext) error {
		stats.Track("Show Version", nil)
		fmt.Println(root.Cmd.GetVersion())
		return nil
	})
}
