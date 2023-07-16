# hap-switch-cmd

This app provides homekit virtual switch devices executing local commands.

## Usage

### Install

#### binary

Download from Releases.

#### macOS(WIP)

Please install by homebrew.

```sh
brew install hap-switch-command
```

#### Ubuntu(WIP)

```sh
apt install hap-switch-command
```

#### Docker

```sh
docker pull 
```

## Config

All configs are provided from `~/.hap-switch-command/config.yml` file.
Create a configuration file with the following command and edit it.

```sh
hap-switch-command init
vi ~/.hap-switch-command/config.yml
```

Sample: [`sample/configs.yml`](./cmd/sample/configs.yml).


## start

### binary

```sh
hap-switch-command serve
```

### macOS(WIP)

It can be start by homebrew services.

```sh
brew services start hap-switch-command
```

### Ubuntu(WIP)

```sh
systemctl start hap-switch-command
```

### Docker(WIP)

```sh
docker run -v ""
```
