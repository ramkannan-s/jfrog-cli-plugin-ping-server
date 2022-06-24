# ping-server

## About this plugin
This plugin is a simple Artifactory cleanup plugin.
It can be used to delete all artifacts that have not been downloaded or modified for the past n time units (both can be configured)
from a given repository.

**Note:**
If you're planning to clean Docker repositories, this plugin may lead to unexpectedly partial or broken images. It is currently recommended to instead use the following Artifactory [cleanDockerImages](https://github.com/jfrog/artifactory-user-plugins/tree/master/cleanup/cleanDockerImages) user plugin for this purpose.

## Installation with JFrog CLI
Installing the latest version:

`$ jf plugin install ping-server`

Installing a specific version:

`$ jf plugin install ping-server@version`

Uninstalling a plugin

`$ jf plugin uninstall ping-server`

## Usage
### Commands
* clean
    - Arguments:
        - repository - The name of the repository you would like to clean.
    - Flags:
        - server-id: The Artifactory server ID configured using the config command.
        - time-unit: The time unit of the no-dl time. year, month and day are the allowed values. **[Default: month]**
        - no-dl: Artifacts that have not been downloaded or modified for at least no-dl will be deleted. **[Default: 1]**
    - Examples:
    ```
    $ jf ping-server --url https://ramkannans-sbx.dev.gcp.devopsacc.team/artifactory/ --user admin --password ******

    ```

### Environment variables
None.

## Additional info
None.

## Release Notes
The release notes are available [here](RELEASE.md).
