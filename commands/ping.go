package commands

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jfrog/jfrog-cli-core/v2/artifactory/utils"
	"github.com/jfrog/jfrog-cli-core/v2/common/commands"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-cli-core/v2/utils/config"
	searchutils "github.com/jfrog/jfrog-client-go/artifactory/services/utils"
	clientutils "github.com/jfrog/jfrog-client-go/utils"
)

func GetPingCommand() components.Command {
	return components.Command{
		Name:        "ping_server",
		Description: "ping the server for health",
		Aliases:     []string{"c"},
		Arguments:   getCredentials(),
		Flags:       getFlags(),
		EnvVars:     getCleanEnvVar(),
		Action: func(c *components.Context) error {
			return pingCmd(c)
		},
	}
}

func getCredentials() []components.Argument {
	return []components.Argument{
		{
			Name:        "user",
			Description: "username for artifactory",
			DefaultValue: "admin"
		},
		{
			Name:        "password",
			Description: "pwd for artifactory",
			DefaultValue: "password"
		}
	}
}

func getFlags() []components.Flag {
	return []components.Flag{
		components.StringFlag{
			Name:        "url",
			Description: "Artifactory server URL configured using the config command.",
		}
	}
}

func getCleanEnvVar() []components.EnvVar {
	return []components.EnvVar{
		{},
	}
}

type pingConfiguration struct {
	rt string
	ping    string
	user string
	password string
	url string
}

func pingCmd(c *components.Context) error {
	if len(c.Arguments) != 1 {
		return errors.New("Wrong number of arguments. Expected: 1, " + "Received: " + strconv.Itoa(len(c.Arguments)))
	}
	var conf = new(pingConfiguration)
	conf.rt = "rt"
	conf.ping = "ping"
	conf.user = c.Arguments[0]
	conf.password = c.Arguments[1]
	conf.url = c.GetStringFlagValue("url")
	doPing(conf)
}


func doPing(c *pingConfiguration) string {
	greet := "jf" + c.rt + c.ping + "--user" + conf.user "--password" + conf.password + "--url" + c.url + "\n"

	return strings.TrimSpace(greet)
}
