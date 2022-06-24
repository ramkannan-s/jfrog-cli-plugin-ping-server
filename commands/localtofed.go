package commands

import (
	"errors"
	//	"fmt"
	"strconv"
	//	"strings"

	"github.com/jfrog/jfrog-cli-core/v2/artifactory/utils"
	"github.com/jfrog/jfrog-cli-core/v2/common/commands"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"

	//	"github.com/jfrog/jfrog-client-go/artifactory"
	//	"github.com/jfrog/jfrog-client-go/artifactory/services"
	"github.com/jfrog/jfrog-cli-core/v2/utils/config"

	clientutils "github.com/jfrog/jfrog-client-go/utils"
)

func PerformLocalFederatedCommand() components.Command {
	return components.Command{
		Name:        "localfed",
		Description: "convert local to federated",
		Aliases:     []string{"lf"},
		Arguments:   getLFArguments(),
		Flags:       getLFFlags(),
		EnvVars:     getCleanEnvVar(),
		Action: func(c *components.Context) error {
			return LFCmd(c)
		},
	}
}

// Returns the Artifactory Details of the provided server-id, or the default one.
func getRtDetails(c *components.Context) (*config.ServerDetails, error) {
	//print(c.Arguments[0])
	//	print(config.ServerId)

	configuration, err := config.GetAllServersConfigs()
	var serverIds []string
	for _, serverConfig := range configuration {
		serverIds = append(serverIds, serverConfig.ServerId)
		print(serverIds)
		//	print(serverConfig.ServerId)
	}

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

func getLFArguments() []components.Argument {
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

func getLFFlags() []components.Flag {
	return []components.Flag{
		components.StringFlag{
			Name:         "project",
			Description:  "filter repo name based on JFrog project name",
			DefaultValue: "",
		},
	}
}

func LFCmd(c *components.Context) error {

	if len(c.Arguments) != 2 {
		return errors.New("Wrong number of arguments. Expected: 2, " + "Received: " + strconv.Itoa(len(c.Arguments)))
	}

	rtDetails, err := getRtDetails(c)

	if err != nil {
		return err
	}

	servicesManager, err := utils.CreateServiceManager(rtDetails, -1, -1, false)

	servicesManager.ConvertLocalToFederatedRepository(c.Arguments[1])

	return nil

}
