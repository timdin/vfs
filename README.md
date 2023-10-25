# vfs (virtual filesystem)

virtual file system assignment

## CLI usages

Note:

- The parameter followed by a `?` represent the parameter is optional.
- The parameter allows multiple words wrapped in quotation mark, i.e. `"user name"`

User management

- Register: `myvfs register [name]`

Folder management

- create folder: `myvfs create-folder [username] [foldername] [description]?`
- delete folder: `myvfs delete-folder [username] [foldername]`
- list folders: `myvfs list-folders [username] [--sort-name|--sort-created] [asc|desc]`
- rename folder: `myvfs rename-folder [username] [foldername] [new-folder-name]`

File management

- create file: `myvfs create-file [username] [foldername] [filename] [description]?`
- delete file: `myvfs delete-file [username] [foldername] [filename]`
- list file: `myvfs list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]`
- rename file: `myvfs rename-file [username] [foldername] [filename] [new-file-name]`

## running the app

1. Run

    ``` bash
    > make build
    ```

2. (Optional) run the follwing command

    ``` bash
    > alias myvfs="$PWD/main"
    ```

    This will set the alias to `myvfs`, and save some typing when running some commands

    It will only effect the current terminal session, so no need to worry about environment pollution
3. If you want to test out the remote storage mode, also run

    ```bash
    > docker compose up
    ```

    Which will get your db ready
4. Run

    ```bash
    > make test
    ```

    for unit tests

## configuration options

This application has multiple configuration options available in `config.yaml` and `constants.go`

`config.yaml`

- db_mode: This determines the logging mode for the gorm client
  - dev: verbose
  - prod: quiet
- db_type: This determines the database type used
  - remote: it will use mysql database with the connection defined in `database/conn`
  - local: it will use sqlite database with the path defined in `local_database/path`

`constants.go`

- ValidStringPattern: The character validation of input arguments
  - Currently set to accept space, underscore, hyphen, and characters from A-Z, a-z, 0-9
- ValidLength: The single argument length limitation
  - Currently set to 20
- TimeFormat: The time format applied to listing commands
  - Currently set to `YYYY-MM-DD HH-mm-SS`

## repo walkthrough

- cmd: all the commands implementation
- configs: the configuration structure and loader
- constants: cross package commonly used values
- helper: helper functions for unit tests
- mock: go generate mock
- storage: storage implementation
- validation: validation functions implementation for arguments

## tech stack

- cobra: the CLI interface engine
- gorm: the golang Object Relational Mapping library
- go-cmp: handy for printing diff in unit tests
