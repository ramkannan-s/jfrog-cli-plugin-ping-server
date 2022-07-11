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

func MoveRepositoryContents() components.Command {
	return components.Command{
		Name:        "move-contents",
		Description: "Move the repository contents.",
		Arguments:   getRepoContentArguments(),
		Aliases:     []string{"mrc"},
		Action: func(c *components.Context) error {
			return MoveRepoContents(c)
		},
	}
}

func getRepoContentArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "server-id",
			Description: "Enter the Server ID",
		},
	}
}

func MoveRepoContents(c *components.Context) error {
	params := services.NewMoveCopyParams()
	// Source replication repository.
	params.Pattern = "test-1-replicate"
	params.Target = "test-1-replicate-tmp"
	params.Recursive = true
	params.Flat = false
	//params.PathPrefix = "/path/to/repo"

	rtDetails, err := getRtDetails3(c)
	servicesManager, err := utils.CreateServiceManager(rtDetails, -1, -1, false)
	servicesManager.Move(params)
	if err != nil {
		return nil
	}
	return nil
}

// Returns the Artifactory Details of the provided server-id, or the default one.
func getRtDetails3(c *components.Context) (*config.ServerDetails, error) {
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
