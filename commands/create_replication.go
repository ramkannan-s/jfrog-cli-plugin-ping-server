package commands

import (
	"errors"

	"github.com/jfrog/jfrog-cli-core/v2/artifactory/utils"
	"github.com/jfrog/jfrog-cli-core/v2/common/commands"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-client-go/artifactory/services"

	//	"github.com/jfrog/jfrog-client-go/artifactory"
	//	"github.com/jfrog/jfrog-client-go/artifactory/services"
	"github.com/jfrog/jfrog-cli-core/v2/utils/config"

	clientutils "github.com/jfrog/jfrog-client-go/utils"
)

func StartReplicationCommand() components.Command {
	return components.Command{
		Name:        "set-replication",
		Description: "Configure push replication.",
		Arguments:   getReplicationArguments(),
		Aliases:     []string{"sr"},
		Action: func(c *components.Context) error {
			return CreateReplication1(c)
		},
	}
}

func getReplicationArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "server-id",
			Description: "Enter the Server ID",
		},
	}
}

func CreateReplication1(c *components.Context) error {
	params := services.NewCreateReplicationParams()
	// Source replication repository.
	params.RepoKey = "test-2-replicate"
	params.CronExp = "0 0 12 * * ?"
	params.Username = "admin"
	params.Password = "@World10"
	params.Url = "https://ramkannans-apac-sbx.dev.gcp.devopsacc.team/artifactory/apac-test-1"
	params.Enabled = true
	params.SocketTimeoutMillis = 15000
	params.EnableEventReplication = true
	params.SyncDeletes = true
	params.SyncProperties = true
	params.SyncStatistics = true
	//params.PathPrefix = "/path/to/repo"

	rtDetails, err := getRtDetails1(c)
	servicesManager, err := utils.CreateServiceManager(rtDetails, -1, -1, false)
	err = servicesManager.CreateReplication(params)
	if err != nil {
		return nil
	}
	return nil
}

// Returns the Artifactory Details of the provided server-id, or the default one.
func getRtDetails1(c *components.Context) (*config.ServerDetails, error) {
	details, err := commands.GetConfig(c.Arguments[0], false)

	if err != nil {
		return nil, err
	}
	if details.ArtifactoryUrl == "" {
		return nil, errors.New("no server-id was found, or the server-id has no Artifactory url.")
	}
	details.ArtifactoryUrl = clientutils.AddTrailingSlashIfNeeded(details.ArtifactoryUrl)
	err = config.CreateInitialRefreshableTokensIfNeeded(details)
	if err != nil {
		return nil, err
	}
	return details, nil
}
