package grifts

import (
  "github.com/gobuffalo/buffalo"
	"github.com/joscelynjames/birds/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
