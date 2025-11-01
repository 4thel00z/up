package run

import (
	"github.com/4thel00z/up/internal/cli/root"
	"github.com/4thel00z/up/internal/stats"
	"github.com/4thel00z/up/internal/util"
	"github.com/pkg/errors"
	"github.com/tj/kingpin"
)

func init() {
	cmd := root.Command("run", "Run a hook.")
	cmd.Example(`up run build`, "Run build hook.")
	cmd.Example(`up run clean`, "Run clean hook.")

	hook := cmd.Arg("hook", "Name of the hook to run.").Required().String()
	stage := cmd.Flag("stage", "Target stage name.").Short('s').Default("staging").String()

	cmd.Action(func(_ *kingpin.ParseContext) error {
		_, p, err := root.Init()
		if err != nil {
			return errors.Wrap(err, "initializing")
		}

		defer util.Pad()()

		stats.Track("Hook", map[string]interface{}{
			"name": *hook,
		})

		if err := p.Init(*stage); err != nil {
			return errors.Wrap(err, "initializing")
		}

		return p.RunHook(*hook)
	})
}
