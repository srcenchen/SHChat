package main

import (
	_ "SHChat/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"SHChat/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
