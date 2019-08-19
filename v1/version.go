package v1

import (
	"github.com/chargehive/configuration/v1/pool"
	"github.com/chargehive/configuration/v1/scheduler"
)

func AddHandlers() {
	pool.AddHandlers()
	scheduler.AddHandlers()
}
