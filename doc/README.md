# woodpecker-feishu-group-robot

how to dev

## env

- minimum go version: go 1.19
- change `go 1.19`, `^1.19`, `1.19.13` to new go version

### libs

| lib                                                | version |
|:---------------------------------------------------|:--------|
| https://github.com/stretchr/testify                | v1.8.4  |
| https://github.com/sebdah/goldie                   | v2.5.3  |
| https://github.com/gookit/color                    | v1.5.3  |
| https://github.com/urfave/cli/                     | v2.23.7 |
| https://github.com/woodpecker-kit/woodpecker-tools | v1.5.0  |

- more libs see [go.mod](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/blob/main/go.mod)

## depends

in go mod project

```bash
# warning use private git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "https://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q https://github.com/woodpecker-kit/woodpecker-feishu-group-robot.git

# test depends see full version
$ go list -mod readonly -v -m -versions github.com/woodpecker-kit/woodpecker-feishu-group-robot
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/woodpecker-kit/woodpecker-feishu-group-robot | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## local dev

```bash
# It needs to be executed after the first use or update of dependencies.
$ make init dep
```

- test code

```bash
$ make test testBenchmark
```

add main.go file and run

```bash
# run and shell help
$ make devHelp

# run at PLUGIN_DEBUG=true
$ make dev

# run at ordinary mode
$ make run
```

- ci to fast check

```bash
# check style at local
$ make style

# run ci at local
$ make ci
```

### docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# more info see
$ make helpDocker
```

### EngineeringStructure

```
.
├── Dockerfile                     # ci docker build
├── build.dockerfile               # local docker build
├── Makefile                       # make entry
├── README.md
├── build                          # build output folder
├── dist                           # dist output folder
├── cmd
│     ├── cli
│     │     ├── app.go             # cli entry
│     │     ├── cli_aciton_test.go # cli action test
│     │     └── cli_action.go      # cli action
│     └── woodpecker-feishu-group-robot    # command line main package install and dev entrance
│         ├── main.go                   # command line entry
│         └── main_test.go              # integrated test entry
├── constant                       # constant package
│         ├── common_flag.go         # common environment variable
│         ├── platform_flag.go       # platform environment variable
│         └── version.go             # semver version constraint set
├── doc
│         ├── README.md              # command line tools documentation
│         └── docs.md                # woodpecker documentation
├── go.mod
├── go.sum
├── z-MakefileUtils                # make toolkit
├── package.json                   # command line profile information for embed
├── resource.go                    # embed resource
├── internal                          # toolkit package
│         ├── pkgJson                 # package.json toolkit
│         └── version_check           # version check by semver
├── feishu_open_api                # feishu api
├── feishu_plugin                  # feishu_plugin package
│         ├── flag.go                 # plugin flag
│         ├── plugin.go               # plugin entry
│         ├── impl.go                 # plugin implement
│         ├── impl_inner_data.go      # implement for unit test
│         ├── render_config.go        # config of render
│         ├── feishu_card_template.go # render entrance
│         ├── feishu_card_root.go     # render root
│         ├── feishu_card_build_status.go # render of notyce type `build_status`
│         ├── impl_file_browser.go    # implement notyce type `file_browser`
│         ├── transfer_file_browser   # render of notyce type `file_browser`
│         └── settings.go             # plugin settings
├── feishu_plugin_test             # feishu_plugin test
│         ├── init_test.go            # each test init
│         ├── feishu_card_template_test.go # test of render `build_status`
│         ├── feishu_card_ignore_test.go # test of render ignore with `build_status`
│         ├── transfer_oss_browser_test.go # test of render ignore with `file_browser`
│         └── plugin_test.go          # plugin test
└── zymosis                         # resource mark by https://github.com/convention-change/zymosis

```

### log

- open debug log by env `PLUGIN_DEBUG=true` or global flag `--plugin.debug true`

```go
package foo

func GlobalBeforeAction(c *cli.Context) error {
  isDebug := wd_urfave_cli_v2.IsBuildDebugOpen(c)
  if isDebug {
    wd_log.OpenDebug()
  }
  return nil
}
```