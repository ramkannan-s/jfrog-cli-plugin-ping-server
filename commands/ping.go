package commands

import (
	"errors"
	"fmt"
	"strconv"
	"os/exec"
//	"strings"

//	"github.com/jfrog/jfrog-cli-core/v2/artifactory/utils"
//	"github.com/jfrog/jfrog-cli-core/v2/common/commands"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
//	servicesManager "github.com/jfrog/jfrog-client-go/artifactory"
//	"github.com/jfrog/jfrog-cli-core/v2/utils/config"
)

func GetPingCommand() components.Command {
	return components.Command{
		Name:        "ping",
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

/*func GetPingCommand() components.Command {
	return components.Command{
		Name:         "ping",
		Flags:        cliutils.GetCommandFlags(cliutils.Ping),
		Aliases:      []string{"p"},
		Usage:        ping.GetDescription(),
		HelpName:     corecommon.CreateUsage("rt ping", ping.GetDescription(), ping.Usage),
		ArgsUsage:    common.CreateEnvVars(),
		BashComplete: corecommon.CreateBashCompletionFunc(),
		Action: func(c *cli.Context) error {
			return pingCmd(c)
		},
	}
}*/

func getCredentials() []components.Argument {
	return []components.Argument{
		{
			Name:        "user",
			Description: "username for artifactory",
			//Default: "admin",
		},
	}
}

func getFlags() []components.Flag {
	return []components.Flag{
		components.StringFlag{
			Name:        "url",
			Description: "Artifactory server URL configured using the config command.",
		},		
		components.StringFlag{
			Name:        "password",
			Description: "pwd for artifactory",
			//Default: "password",
		},
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
	conf.password = c.GetStringFlagValue("password")
	conf.url = c.GetStringFlagValue("url")
	cliquery := buildCLICmd(conf)

	//err := servicesManager.ConvertLocalToFederatedRepository("generic-local")

	test := doPing(conf)
	fmt.Println(doPing(conf))
	
	cmd := exec.Command(cliquery)
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return errors.New("failed to execute")
    }

    // Print the output
    fmt.Println(string(stdout))

	if len(test) != 1 {
		return errors.New("fdsf")
	}
	return errors.New("test")
}


func doPing(c *pingConfiguration) string {
	greet := "jf " + c.rt + " " + c.ping + " --user " + c.user + " --password" + " " + c.password + " --url" + " " + c.url + "\n"
	return greet
}

func buildCLICmd(c *pingConfiguration) (cliQuery string) {
	// Finds all artfacts that hasn't been downloaded or modified for at least noDownloadedTime
	cliQuery = "jf " + c.rt + " " + c.ping + " --user " + c.user + " --password" + " " + c.password + " --url" + " " + c.url + "\n"

	return fmt.Sprintf(cliQuery)
}