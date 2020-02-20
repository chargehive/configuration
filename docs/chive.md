# ChargeHive Tool - `chive`
All the json configuration files for ChargeHive are applied using the `chive` command line tool. 
With this tool you can create, remove, update and delete policies and configurations. 

## Installation

You will need a `project-id` and an `access-token` from ChargeHive. If you do not have these, get in touch with support.

Download `chive` for your system:
+ [Windows](https://cdn.chargehive.com/tools/chive/dist/windows/chive.exe)
+ [Mac](https://cdn.chargehive.com/tools/chive/dist/mac/chive)

To use the `chive` tool, you can either call it directly with your credentials:
```
chive --project-id="your-project-id" --access-token="your-access-token"
```

Or you can create a credentials file `.chive.yaml` with the following structure:
```yaml
projectID: your-project-id
accessToken: your-access-token
```

You can use the `.chive.yaml` config by any of these methods:
- storing it in the same directory as the `chive` executable
- storing it in the `$HOME` directory
- set where `chive` looks for the file: `chive --config-file="/path/to/your/config/.chive.yaml` and put it there

Once this is done, you should run `chive health` to check that your credentials are correct.
If you receive an error, please get in contact for support.

Options for the `chive` command are as follows:
```
❯ chive --help
chive allows communication with the Charge Hive API via CLI

Usage:
  chive [command]

Available Commands:
  apply       Apply a configuration file
  backup      Backup current project configurations
  delete      Delete a specific configuration
  get         Get a configuration
  health      Check your project and access token details are correct
  help        Help about any command
  list        List stored configurations
  verify      Verify a configuration

Flags:
      --access-token string   Access Token
      --api-host string       API Host
      --config-file string    configuration file (default is $HOME/.chive.yaml)
  -h, --help                  help for chive
      --project-id string     project ID

Use "chive [command] --help" for more information about a command.
```

## Usage

#### Apply a file or directory of configuration files to your project.
```
Usage:
  chive apply -f <file> [flags]

Flags:
  -d, --dir string     directory location for config files
  -f, --file strings   config-file.json file location
  -h, --help           help for apply
  -R, --recursive      process sub directories (when using -d)
```
#### Backup all your stored chargehive configuration files
```
Usage:
  chive backup <kind> [flags]
```
#### Delete a configuration based on its configuration kind and ID
```
Usage:
  chive delete <kind> <id> [flags]
```
#### Retrieve a specific ChargeHive configuration JSON file
```
Usage:
  chive get <kind> <id> [flags]
```

#### Connect to the ChargeHive API, and verify your credentials
```   
Usage:
     chive health [flags]
```

#### List all your stored ChargeHive configuration files
```
Usage:
  chive list <kind> [flags]
```

#### Verify a configuration
```
Usage:
  chive verify <kind> <id> [flags]
```