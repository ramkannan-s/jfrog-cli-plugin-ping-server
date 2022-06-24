package commands

import (
	"errors"
	"fmt"

	"github.com/jfrog/jfrog-cli-core/v2/artifactory/utils"
	"github.com/jfrog/jfrog-cli-core/v2/common/commands"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-client-go/artifactory/services"

	//	"github.com/jfrog/jfrog-client-go/artifactory"
	//	"github.com/jfrog/jfrog-client-go/artifactory/services"
	"github.com/jfrog/jfrog-cli-core/v2/utils/config"

	clientutils "github.com/jfrog/jfrog-client-go/utils"
)

func ListRepositoriesCommand() components.Command {
	return components.Command{
		Name:        "list-repositories",
		Description: "get list of all repositories",
		Arguments:   getAllReposArguments(),
		Aliases:     []string{"lr"},
		Action: func(c *components.Context) error {
			return getRepositories(c)
		},
	}
}

func getAllReposArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "server-id",
			Description: "Enter the Server ID",
		},
	}
}

func getRepositories(c *components.Context) error {
	rtDetails, err := getRtDetails2(c)
	servicesManager, err := utils.CreateServiceManager(rtDetails, -1, -1, false)
	params := services.NewRepositoriesFilterParams()
	repoDetails, err := servicesManager.GetAllRepositoriesFiltered(params)
	//	repoDetails, err := servicesManager.GetAllRepositories()
	//	url := fmt.Sprintf(params.RepoType, params.PackageType)
	//	fmt.Println(url)
	fmt.Println(repoDetails)
	fmt.Println(params.RepoType)
	if err != nil {
		return nil
	}
	return nil
}

// Returns the Artifactory Details of the provided server-id, or the default one.
func getRtDetails2(c *components.Context) (*config.ServerDetails, error) {
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
