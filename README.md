[![ci](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/workflows/ci/badge.svg)](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/actions/workflows/ci.yml)

[![docker hub version semver](https://img.shields.io/docker/v/sinlov/woodpecker-feishu-group-robot?sort=semver)](https://hub.docker.com/r/sinlov/woodpecker-feishu-group-robot/tags?page=1&ordering=last_updated)
[![docker hub image size](https://img.shields.io/docker/image-size/sinlov/woodpecker-feishu-group-robot)](https://hub.docker.com/r/sinlov/woodpecker-feishu-group-robot)
[![docker hub image pulls](https://img.shields.io/docker/pulls/sinlov/woodpecker-feishu-group-robot)](https://hub.docker.com/r/sinlov/woodpecker-feishu-group-robot/tags?page=1&ordering=last_updated)

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

## before use

- sed doc at [feishu Custom bot guide](https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN?lang=en-US), to new group robot
- Configure webhook like `https://open.feishu.cn/open-apis/bot/v2/hook/{web_hook}` end `{web_hook}`
- `{web_hook}` must settings at `settings.feishu_webhook`or `PLUGIN_FEISHU_WEBHOOK`
- if set `Custom keywords` you can change `settings.feishu_msg_title` or `PLUGIN_FEISHU_MSG_TITLE`
- or set `Signature validation` by `settings.feishu_secret` or `PLUGIN_FEISHU_SECRET`
- use `wd_steps_transfer` just add `.woodpecker_kit.steps.transfer` at git ignore

### workflow usage

- workflow with backend `docker`

```yml
labels:
  backend: docker
steps:
  notification-feishu-group-robot:
    image: sinlov/woodpecker-feishu-group-robot:latest
    pull: false
    settings:
      # debug: true # plugin debug switch
      # feishu_ntp_target: "pool.ntp.org" # if not set will not sync ntp time
      feishu_webhook:
        # https://woodpecker-ci.org/docs/usage/secrets
        from_secret: feishu_group_bot_token
      feishu_secret:
        from_secret: feishu_group_secret_bot
      feishu_msg_title: "CI Notification" # default [CI Notification]
      # let notification card change more info see https://open.feishu.cn/document/ukTMukTMukTM/uAjNwUjLwYDM14CM2ATN
      feishu_enable_forward: true
      feishu_status_success_ignore: false # ignore this build success status
      feishu_status_change_success: false # must open [ status_success_ignore ], when status change to success, compare with CI_PREV_PIPELINE_STATUS
    when:
      status: # only support failure/success,  both open will send anything
        - failure
        - success
```

- workflow with backend `local`, must install at local and effective at evn `PATH`

```bash
go install -a github.com/woodpecker-kit/woodpecker-feishu-group-robot/cmd/woodpecker-feishu-group-robot@latest
```

- install at ${GOPATH}/bin, v1.0.0

```bash
go install -v github.com/woodpecker-kit/woodpecker-feishu-group-robot/cmd/woodpecker-feishu-group-robot@v1.0.0
```

```yml
labels:
  backend: local
steps:
  notification-feishu-group-robot:
    image: woodpecker-feishu-group-robot
    settings:
      # debug: true # plugin debug switch
      # feishu_ntp_target: "pool.ntp.org" # if not set will not sync ntp time
      feishu_webhook:
        # https://woodpecker-ci.org/docs/usage/secrets
        from_secret: feishu_group_bot_token
      feishu_secret:
        from_secret: feishu_group_secret_bot
      feishu_msg_title: "CI Notification" # default [CI Notification]
      # let notification card change more info see https://open.feishu.cn/document/ukTMukTMukTM/uAjNwUjLwYDM14CM2ATN
      feishu_enable_forward: true
      feishu_status_success_ignore: false # ignore this build success status
      feishu_status_change_success: false # must open [ status_success_ignore ], when status change to success, compare with CI_PREV_PIPELINE_STATUS
    when:
      status: # only support failure/success,  both open will send anything
        - failure
        - success
```

### steps transfer

#### In step by mark: oss_send_transfer

- with ResourceUrl

```json
{
    "InfoSendResult": "success",
    "OssHost": "https://docs.aws.amazon.com/s3/index.html",
    "OssPath": "dist/demo/pass.tar.gz",
    "OssUserProfile": {
        "UserName": "",
        "UserToken": ""
    },
    "PagePasswd": "",
    "PageUrl": "",
    "ResourceUrl": "https://docs.aws.amazon.com/s/dist/demo/pass.tar.gz"
}
```

- with share by password

```json
{
    "InfoSendResult": "success",
    "OssHost": "https://docs.aws.amazon.com/s3/index.html",
    "OssPath": "dist/demo/pass.tar.gz",
    "OssUserProfile": {
        "UserName": "",
        "UserToken": ""
    },
    "PagePasswd": "abc-zxy",
    "PageUrl": "https://docs.aws.amazon.com/p/dist/demo/page-xyz.html",
    "ResourceUrl": ""
}
```

### custom settings

- `settings.debug` or `PLUGIN_DEBUG` can open plugin debug mode
- `settings.timeout_second` or `PLUGIN_TIMEOUT_SECOND` can set send message timeout

- `settings.feishu_ntp_target` or `PLUGIN_FEISHU_NTP_TARGET` set ntp server to sync time for `Signature validation` by error code 19021
- `settings.feishu_msg_title` or `PLUGIN_FEISHU_MSG_TITLE` can change message card title
- `settings.feishu_enable_forward` or `PLUGIN_FEISHU_ENABLE_FORWARD` can change message share way
- `settings.feishu_msg_powered_by_image_key` or `PLUGIN_FEISHU_MSG_POWERED_BY_IMAGE_KEY` can change card img by feishu-image-key
- `settings.feishu_msg_powered_by_image_alt` or `PLUGIN_FEISHU_MSG_POWERED_BY_IMAGE_ALT` can change card img alt tag name

---

- want dev this project, see [doc](doc/README.md)