package commands

import (
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-client-go/artifactory/services"

	"github.com/jfrog/jfrog-client-go/config"
	"github.com/jfrog/jfrog-client-go/http/jfroghttpclient"
	ioutils "github.com/jfrog/jfrog-client-go/utils/io"
)

const Username = "admin"
const Password = "@World10"
const Url = "https://ramkannans-apac-sbx.dev.gcp.devopsacc.team/"
const CronExp = "30 5 * * *"
const RepoKey = "test-1-replicate"
const Proxy = ""
const EnableEventReplication = false
const SocketTimeoutMillis = 5
const Enabled = false
const SyncDeletes = false
const SyncProperties = false
const SyncStatistics = false
const LocalRepositoryRepoType = "local"

func StartReplicationCommandOLD() components.Command {
	return components.Command{
		Name:        "set-replication",
		Description: "Configure push replication.",
		Aliases:     []string{"sr"},
		Arguments:   migrateArguments(),
		Flags:       setReplicationFlags(),
		Action: func(c *components.Context) error {
			return CreateReplication(c)
		},
	}
}

type ArtifactoryServicesManagerImp struct {
	client   *jfroghttpclient.JfrogHttpClient
	config   config.Config
	progress ioutils.ProgressMgr
}

func migrateArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "server-id",
			Description: "Enter the Server ID",
		},
		{
			Name:        "repo-name",
			Description: "Enter the repo name to convert",
		},
	}
}

func setReplicationFlags() []components.Flag {
	return []components.Flag{
		components.StringFlag{
			Name:         "project",
			Description:  "filter repo name based on JFrog project name",
			DefaultValue: "",
		},
	}
}

func CreateReplication(c *components.Context) error {
	//CreateReplication(services.ReplicationParams)
	return nil
}

func (sm *ArtifactoryServicesManagerImp) CreateReplication(params services.CreateReplicationParams) error {
	replicationService := services.NewCreateReplicationService(sm.client)
	replicationService.ArtDetails = sm.config.GetServiceDetails()
	return replicationService.CreateReplication(params)
}
