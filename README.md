[![ci](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/workflows/ci/badge.svg)](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/actions/workflows/ci.yml)

[![docker hub version semver](https://img.shields.io/docker/v/sinlov/woodpecker-feishu-group-robot?sort=semver)](https://hub.docker.com/r/sinlov/woodpecker-feishu-group-robot/tags?page=1&ordering=last_updated)
[![docker hub image size](https://img.shields.io/docker/image-size/sinlov/woodpecker-feishu-group-robot)](https://hub.docker.com/r/sinlov/woodpecker-feishu-group-robot)
[![docker hub image pulls](https://img.shields.io/docker/pulls/sinlov/woodpecker-feishu-group-robot)](https://hub.docker.com/r/sinlov/woodpecker-feishu-group-robot/tags?page=1&ordering=last_updated)

[![go mod version](https://img.shields.io/github/go-mod/go-version/woodpecker-kit/woodpecker-feishu-group-robot?label=go.mod)](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot)
[![GoDoc](https://godoc.org/github.com/woodpecker-kit/woodpecker-feishu-group-robot?status.png)](https://godoc.org/github.com/woodpecker-kit/woodpecker-feishu-group-robot)
[![goreportcard](https://goreportcard.com/badge/github.com/woodpecker-kit/woodpecker-feishu-group-robot)](https://goreportcard.com/report/github.com/woodpecker-kit/woodpecker-feishu-group-robot)

[![GitHub license](https://img.shields.io/github/license/woodpecker-kit/woodpecker-feishu-group-robot)](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot)
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

- [x] simple to set up and easy to use
- [x] Supports ignoring build success notifications in the same steps and comparing notifications after the last build failure.
- [x] internationalization support: en-US, zh-CN more support see --help (v1.4.+)
- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## usage

- see [doc](doc/docs.md)

## before use

- sed doc at [feishu Custom bot guide](https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN?lang=en-US), to new group robot
- Configure webhook like `https://open.feishu.cn/open-apis/bot/v2/hook/{web_hook}` end `{web_hook}`
  - `{web_hook}` must settings at `settings.feishu-webhook`or `PLUGIN_FEISHU_WEBHOOK`
- Feishu security settings
  - if set `Custom keywords` you can change `settings.feishu-msg-title` or `PLUGIN_FEISHU_MSG_TITLE`
  - fi set `Signature validation` by `settings.feishu-secret` or `PLUGIN_FEISHU_SECRET`
- just add `.woodpecker_kit.steps.transfer` at git ignore

## dev

- want dev this project, see [dev doc](doc/README.md)