---
name: woodpecker-feishu-group-robot
description: woodpecker plugin template
author: woodpecker-kit
tags: [ feishu, robot, notice ]
containerImage: sinlov/woodpecker-feishu-group-robot
containerImageUrl: https://hub.docker.com/r/sinlov/woodpecker-feishu-group-robot
url: https://github.com/woodpecker-kit/woodpecker-feishu-group-robot
icon: https://raw.githubusercontent.com/woodpecker-kit/woodpecker-feishu-group-robot/main/doc/icon.png
---

woodpecker-feishu-group-robot

## Features

- [x] simple to set up and easy to use
- [x] Supports ignoring build success notifications in the same steps and comparing notifications after the last build
  failure.
- [x] internationalization support: en-US, zh-CN more support see --help (v1.4.+)
- [x] docker platform support
    - linux/amd64 linux/386 linux/arm64/v8 linux/arm/v7 linux/ppc64le linux/s390x (v1.4.+)

## before use

- sed doc at [feishu Custom bot guide](https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN?lang=en-US), to
  new group robot
- Configure webhook like `https://open.feishu.cn/open-apis/bot/v2/hook/{web_hook}` end `{web_hook}`
    - `{web_hook}` must settings at `settings.feishu-webhook`or `PLUGIN_FEISHU_WEBHOOK`
- Feishu security settings
    - if set `Custom keywords` you can change `settings.feishu-msg-title` or `PLUGIN_FEISHU_MSG_TITLE`
    - fi set `Signature validation` by `settings.feishu-secret` or `PLUGIN_FEISHU_SECRET`
- just add `.woodpecker_kit.steps.transfer` at git ignore

## Settings

| Name                           | Required | Default value     | Description                                                                                                       |
|--------------------------------|----------|-------------------|-------------------------------------------------------------------------------------------------------------------|
| `debug`                        | **no**   | *false*           | open debug log or open by env `PLUGIN_DEBUG`                                                                      |
| `force-status`                 | **no**   | *""*              | force status (v1.8+). If empty will use woodpecker ci pipeline status, only support `[success failure]`           |
| `feishu-enable-debug-notice`   | **no**   | *false*           | when debug open, will not send message, must enable it to notice under debug open                                 |
| `feishu-webhook`               | **yes**  | *none*            | feishu group robot webhook, end of feishu robot `https://open.feishu.cn/open-apis/bot/v2/hook/{web_hook}`         |
| `feishu-secret`                | **yes**  | *none*            | feishu robot secret, just `signature verification`, empty will not open.                                          |
| `feishu-msg-title`             | **yes**  | *CI Notification* | feishu group robot message title, most input `Security settings area` keywords                                    |
| `feishu-notice-types`          | **no**   | *none*            | feishu notice types, if empty will use `[ build_status ]`                                                         |
| `feishu-msg-i18n-lang`         | **no**   | *en-US*           | feishu group robot message i18n lang, support: en-US, zh-CN more support see --help (v1.4.+)                      |
| `feishu-status-success-ignore` | **no**   | *false*           | ignore this build success status                                                                                  |
| `feishu-status-change-success` | **no**   | *false*           | must open `[ feishu-status-success-ignore ]`, when status change to success, compare with CI_PREV_PIPELINE_STATUS |
| `feishu-enable-forward`        | **no**   | *false*           | let notification card change more info see https://open.feishu.cn/document/ukTMukTMukTM/uAjNwUjLwYDM14CM2ATN      |

**custom settings**

| Name                              | Required | Default value | Description                                           |
|-----------------------------------|----------|---------------|-------------------------------------------------------|
| `feishu-ntp-target`               | **no**   | *none*        | like "pool.ntp.org" if not set will not sync ntp time |
| `feishu-msg-powered-by-image-key` | **no**   | *none*        | card img by feishu-image-key                          |
| `feishu-msg-powered-by-image-alt` | **no**   | *none*        | card img alt tag name                                 |

**Hide Settings:**

| Name                                        | Required | Default value                    | Description                                                                      |
|---------------------------------------------|----------|----------------------------------|----------------------------------------------------------------------------------|
| `timeout_second`                            | **no**   | *10*                             | command timeout setting by second                                                |
| `woodpecker-kit-steps-transfer-file-path`   | **no**   | `.woodpecker_kit.steps.transfer` | Steps transfer file path, default by `wd_steps_transfer.DefaultKitStepsFileName` |
| `woodpecker-kit-steps-transfer-disable-out` | **no**   | *false*                          | Steps transfer write disable out                                                 |

## Example

- workflow with backend `docker`

[![docker hub version semver](https://img.shields.io/docker/v/sinlov/woodpecker-feishu-group-robot?sort=semver)](https://hub.docker.com/r/sinlov/woodpecker-feishu-group-robot/tags?page=1&ordering=last_updated)
[![docker hub image size](https://img.shields.io/docker/image-size/sinlov/woodpecker-feishu-group-robot)](https://hub.docker.com/r/sinlov/woodpecker-feishu-group-robot)
[![docker hub image pulls](https://img.shields.io/docker/pulls/sinlov/woodpecker-feishu-group-robot)](https://hub.docker.com/r/sinlov/woodpecker-feishu-group-robot/tags?page=1&ordering=last_updated)

```yml
labels:
  backend: docker
steps:
  notification-feishu-group-robot:
    image: sinlov/woodpecker-feishu-group-robot:latest
    pull: false
    settings:
      # debug: true # plugin debug switch
      ## force status (v1.8+). If empty will use woodpecker ci pipeline status, only support [success failure]
      # force-status: "failure"
      # feishu-enable-debug-notice: true when debug open, will not send message, must enable it to notice under debug open
      # feishu-ntp-target: "pool.ntp.org" # if not set will not sync ntp time
      feishu-webhook:
        # https://woodpecker-ci.org/docs/usage/secrets
        from_secret: feishu_group_bot_token
      feishu-secret:
        from_secret: feishu_group_secret_bot
      feishu-msg-i18n-lang: en-US # support: en-US, zh-CN more support see --help (v1.4.+)
      feishu-msg-title: "CI Notification" # default [CI Notification]
      # let notification card change more info see https://open.feishu.cn/document/ukTMukTMukTM/uAjNwUjLwYDM14CM2ATN
      feishu-enable-forward: true
      feishu-status-success-ignore: false # ignore this build success status
      feishu-status-change-success: false # must open [ feishu-status-success-ignore ], when status change to success, compare with CI_PREV_PIPELINE_STATUS
    when:
      status: # only support failure/success,  both open will send anything
        - failure
        - success
```

- workflow with backend `local`, must install at local and effective at evn `PATH`

```bash
go install -a github.com/woodpecker-kit/woodpecker-feishu-group-robot/cmd/woodpecker-feishu-group-robot@latest
```

[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/woodpecker-kit/woodpecker-feishu-group-robot)](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/tags)
[![GitHub release)](https://img.shields.io/github/v/release/woodpecker-kit/woodpecker-feishu-group-robot)](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/releases)

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
      ## force status (v1.8+). If empty will use woodpecker ci pipeline status, only support [success failure]
      # force-status: "failure"
      # feishu-enable-debug-notice: true when debug open, will not send message, must enable it to notice under debug open
      # feishu-ntp-target: "pool.ntp.org" # if not set will not sync ntp time
      feishu-webhook:
        # https://woodpecker-ci.org/docs/usage/secrets
        from_secret: feishu_group_bot_token
      feishu-secret:
        from_secret: feishu_group_secret_bot
      feishu-msg-i18n-lang: en-US # support: en-US, zh-CN more support see --help (v1.4.+)
      feishu-msg-title: "CI Notification" # default [CI Notification]
      # let notification card change more info see https://open.feishu.cn/document/ukTMukTMukTM/uAjNwUjLwYDM14CM2ATN
      feishu-enable-forward: true
      feishu-status-success-ignore: false # ignore this build success status
      feishu-status-change-success: false # must open [ feishu-status-success-ignore ], when status change to success, compare with CI_PREV_PIPELINE_STATUS
    when:
      status: # only support failure/success,  both open will send anything
        - failure
        - success
```

### force notice failure

```yaml
labels: # https://woodpecker-ci.org/docs/usage/workflow-syntax#labels
  platform: linux/amd64
  backend: docker

steps:
  notification-feishu-failure:
    image: sinlov/woodpecker-feishu-group-robot:latest
    pull: true
    settings:
      # debug: true # plugin debug switch
      # force status (v1.8+). If empty will use woodpecker ci pipeline status, only support [success failure]
      force-status: "failure"
      feishu-webhook:
        # https://woodpecker-ci.org/docs/usage/secrets
        from_secret: feishu_group_bot_token
      feishu-secret:
        from_secret: feishu_group_secret_bot
      feishu-msg-i18n-lang: en-US # support: en-US, zh-CN more support see --help (v1.4.+)
      feishu-msg-title: "CI Notification" # default [CI Notification]
      # let notification card change more info see https://open.feishu.cn/document/ukTMukTMukTM/uAjNwUjLwYDM14CM2ATN
      feishu-enable-forward: true
    when:
      status: # only support failure/success, both open will send anything
        - failure
```

### steps transfer

- The response needs to enable support for the type corresponding to `settings.feishu-notice-types`
- less support version is `v1.3.0`

#### With notice type: `file_browser`

- support plugin [woodpecker-file-browser-upload](https://github.com/woodpecker-kit/woodpecker-file-browser-upload)
    - support version `v1.5.+`

[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/woodpecker-kit/woodpecker-file-browser-upload)](https://github.com/woodpecker-kit/woodpecker-file-browser-upload/tags)
[![GitHub release)](https://img.shields.io/github/v/release/woodpecker-kit/woodpecker-file-browser-upload)](https://github.com/woodpecker-kit/woodpecker-file-browser-upload/releases)
[![docker hub version semver](https://img.shields.io/docker/v/sinlov/woodpecker-file-browser-upload?sort=semver)](https://hub.docker.com/r/sinlov/woodpecker-file-browser-upload/tags?page=1&ordering=last_updated)

```yml
labels:
  backend: docker
steps:
  woodpecker-file-browser-upload:
    image: sinlov/woodpecker-file-browser-upload:latest
    pull: false
    settings:
      # debug: false # plugin debug switch
      file-browser-host: "http://127.0.0.1:80" # must set args, file_browser host like http://127.0.0.1:80
      file-browser-username: # must set args, file_browser username
        # https://woodpecker-ci.org/docs/usage/secrets
        from_secret: file_browser_user_name
      file-browser-user-password: # must set args, file_browser user password
        from_secret: file_browser_user_passwd
      file-browser-remote-root-path: dist/ # must set args, send to file_browser base path
      file-browser-dist-type: git # must set args, type of dist file graph only can use: git, custom
      file-browser-file-glob: # must set args, globs list of send to file_browser under file-browser-target-dist-root-path
        - "**/*.tar.gz"
        - "**/*.sha256"

  notification-feishu-group-robot:
    image: sinlov/woodpecker-feishu-group-robot:latest
    pull: false
    depends_on:
      - woodpecker-file-browser-upload # must depend on woodpecker-file-browser-upload plugin
    settings:
      feishu-webhook:
        # https://woodpecker-ci.org/docs/usage/secrets
        from_secret: feishu_group_bot_token
      feishu-secret:
        from_secret: feishu_group_secret_bot
      feishu-msg-title: "CI Notification" # default [CI Notification]
      # let notification card change more info see https://open.feishu.cn/document/ukTMukTMukTM/uAjNwUjLwYDM14CM2ATN
      feishu-enable-forward: true
      feishu-notice-types:
        - file_browser # after file_browser upload will show file_browser info
        - build_status # also show build status
    when:
      status: # only support failure/success,  both open will send anything
        - failure
        - success
```
