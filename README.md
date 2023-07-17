# hap-switch-cmd

[![Static Badge](https://img.shields.io/badge/homebrew-legnoh%2Fetc%2Fhap--switch--command-orange?logo=apple)](https://github.com/legnoh/homebrew-etc/blob/main/Formula/hap-switch-command.rb)
[![Static Badge](https://img.shields.io/badge/image-ghcr.io%2Flegnoh%2Fhap--switch--command-blue?logo=github)](https://github.com/legnoh/hap-switch-command/pkgs/container/hap-switch-command)

This app provides homekit virtual switch devices executing local commands.

## Usage

Install, init, and start. That's it !

All configs are provided from `~/.hap-switch-command/config.yml` file.  
Create a configuration file with the following command and edit it.

- Config sample: [`sample/configs.yml`](./cmd/sample/configs.yml).

### macOS

```sh
# install
brew install legnoh/etc/hap-switch-command

# init & edit
hap-switch-command init
vi ~/.hap-switch-command/config.yml

# start
brew services start hap-switch-command
```

### Docker

> **Warning**
> This app does not work when running in Docker for Mac or Docker for Windows due to [this](https://github.com/docker/for-mac/issues/68) and [this](https://github.com/docker/for-win/issues/543).

```sh
# pull
docker pull ghcr.io/legnoh/hap-switch-command

# init
docker run \
    -v .:/root/.hap-switch-command/ \
    ghcr.io/legnoh/hap-switch-command init

# edit
vi config.yml

# start
docker run \
    --network host \
    -v "./config.yml:/root/.hap-switch-command/config.yml" \
    ghcr.io/legnoh/hap-switch-command
```
