package main

import (
	"log"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/atlas"
	"github.com/mongodb/mongocli/internal/cli/cloudmanager"
	"github.com/mongodb/mongocli/internal/cli/config"
	"github.com/mongodb/mongocli/internal/cli/iam"
	"github.com/mongodb/mongocli/internal/cli/opsmanager"
	"github.com/spf13/cobra/doc"
)

func main() {
	mongocli := cli.Builder()
	mongocli.AddCommand(
		config.Builder(),
		atlas.Builder(),
		cloudmanager.Builder(),
		opsmanager.Builder(),
		iam.Builder(),
	)

	err := doc.GenReSTTree(mongocli, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}
