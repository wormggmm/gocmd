package shell

import (
	"github.com/wormggmm/gocmd/common"
	"github.com/wormggmm/gocmd/controller"
)

type BaseShell struct {
	*controller.BaseController
}

func NewBaseShell(hooker common.IKeyHooker) *BaseShell {
	return &BaseShell{
		BaseController: controller.NewBaseController(hooker),
	}
}
