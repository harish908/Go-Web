package cmd

import (
	"PortalClient/configs"
	"PortalClient/pkg/tracing"
)

func Init() {
	configs.InitConfigFile()
	configs.InitMySql()
	tracing.InitTracer()
}
