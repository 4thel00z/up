package reporter

import (
	"github.com/4thel00z/up/reporter/discard"
	"github.com/4thel00z/up/reporter/plain"
	"github.com/4thel00z/up/reporter/text"
)

var (
	// Discard reporter.
	Discard = discard.Report

	// Plain reporter.
	Plain = plain.Report

	// Text reporter.
	Text = text.Report
)
