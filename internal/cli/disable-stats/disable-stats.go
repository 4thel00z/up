package disablestats

import (
	"github.com/pkg/errors"
	"github.com/tj/kingpin"

	"github.com/4thel00z/up/internal/cli/root"
	"github.com/4thel00z/up/internal/stats"
)

func init() {
	cmd := root.Command("disable-stats", "Disable anonymized usage stats").Hidden()
	cmd.Action(func(_ *kingpin.ParseContext) error {
		err := stats.Client.Disable()
		if err != nil {
			return errors.Wrap(err, "disabling")
		}
		return nil
	})
}
