package main

import (
	"github.com/jfrog/jfrog-cli-core/v2/plugins"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"/Users/ramkannans/git_repos/frog_cli_plugins/ping-server/commands"
)

func main() {
	plugins.PluginMain(getApp())
}

func getApp() components.App {
	app := components.App{}
	app.Name = "ping-server"
	app.Description = "Ping the Server for Health Status"
	app.Version = "v1.1.2"
	app.Commands = getCommands()
	return app
}

func getCommands() []components.Command {
	return []components.Command{
		commands.GetPingCommand()}
}
