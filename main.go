package main

import (
	"log"

	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/adapter"
	config_i "gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/config/implementation"
	"gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/framework"
)

func main() {
	conf := config_i.NewEnvConfigStorage()
	core, err := framework.NewCore(conf, adapter.BuildColvirAdapter)
	if err != nil {
		log.Fatal(err)
	}
	core.Start(true)
}
