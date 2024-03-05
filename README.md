[![ci](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/workflows/ci/badge.svg)](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/woodpecker-kit/woodpecker-feishu-group-robot?label=go.mod)](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot)
[![GoDoc](https://godoc.org/github.com/woodpecker-kit/woodpecker-feishu-group-robot?status.png)](https://godoc.org/github.com/woodpecker-kit/woodpecker-feishu-group-robot)
[![goreportcard](https://goreportcard.com/badge/github.com/woodpecker-kit/woodpecker-feishu-group-robot)](https://goreportcard.com/report/github.com/woodpecker-kit/woodpecker-feishu-group-robot)

[![GitHub license](https://img.shields.io/github/license/woodpecker-kit/woodpecker-feishu-group-robot)](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot)
[![codecov](https://codecov.io/gh/woodpecker-kit/woodpecker-feishu-group-robot/branch/main/graph/badge.svg)](https://codecov.io/gh/woodpecker-kit/woodpecker-feishu-group-robot)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/woodpecker-kit/woodpecker-feishu-group-robot)](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/tags)
[![GitHub release)](https://img.shields.io/github/v/release/woodpecker-kit/woodpecker-feishu-group-robot)](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/releases)

## for what

- this project used to woodpecker plugin

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/woodpecker-kit/woodpecker-feishu-group-robot)](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## Features

- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## usage

- use this template, replace list below and add usage
    - `github.com/woodpecker-kit/woodpecker-feishu-group-robot` to your package name
    - `woodpecker-kit` to your owner name
    - `woodpecker-feishu-group-robot` to your project name

- use github action for this workflow push to docker hub, must add at github secrets 
    - `DOCKERHUB_OWNER` user of docker hub
    - `DOCKERHUB_REPO_NAME` repo name of docker hub
    - `DOCKERHUB_TOKEN` token of docker hub user

- if use `wd_steps_transfer` just add `.woodpecker_kit.steps.transfer` at git ignore

### workflow usage

- workflow with backend `docker`

```yml
labels:
  backend: docker
steps:
  env:
    image: sinlov/woodpecker-feishu-group-robot:latest
    pull: false
    settings:
      # debug: true
      env_printer_print_keys: # print env keys
        - GOPATH
        - GOPRIVATE
        - GOBIN
      env_printer_padding_left_max: 36 # padding left max
      steps_transfer_demo: false # open this show steps transfer demo
```

- workflow with backend `local`, must install at local and effective at evn `PATH`

```bash
# install at ${GOPATH}/bin
$ go install -v github.com/woodpecker-kit/woodpecker-feishu-group-robot/cmd/woodpecker-feishu-group-robot@latest
# install version v1.0.0
$ go install -v github.com/woodpecker-kit/woodpecker-feishu-group-robot/cmd/woodpecker-feishu-group-robot@v1.0.0
```

```yml
labels:
  backend: local
steps:
  env:
    image: woodpecker-feishu-group-robot
    settings:
      # debug: false
      env_printer_print_keys: # print env keys
        - GOPATH
        - GOPRIVATE
        - GOBIN
      env_printer_padding_left_max: 36 # padding left max
      steps_transfer_demo: false # open this show steps transfer demo
```

---

- want dev this project, see [doc](doc/README.md)