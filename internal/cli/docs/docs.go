package docs

import (
	"github.com/pkg/browser"
	"github.com/tj/kingpin"

	"github.com/4thel00z/up/internal/cli/root"
	"github.com/4thel00z/up/internal/stats"
)

var url = "https://up.docs.apex.sh"

func init() {
	cmd := root.Command("docs", "Open documentation website in the browser.")
	cmd.Example(`up docs`, "Open the documentation site.")

	cmd.Action(func(_ *kingpin.ParseContext) error {
		stats.Track("Open Docs", nil)
		return browser.OpenURL(url)
	})
}
