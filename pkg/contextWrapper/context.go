package contextWrapper

import (
	"context"
)

var Ctx context.Context
var Cancel context.CancelFunc
func Start() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	Ctx = ctx
	Cancel = cancel
}
