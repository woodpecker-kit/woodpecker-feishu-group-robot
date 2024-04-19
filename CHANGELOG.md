# Changelog

All notable changes to this project will be documented in this file. See [convention-change-log](https://github.com/convention-change/convention-change-log) for commit guidelines.

## [1.4.0](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/compare/1.3.0...v1.4.0) (2024-04-19)

### ✨ Features

* add settings `feishu-msg-i18n-lang` to change i18n, and add check of this setings ([5343b070](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/5343b070c237291883e7bc2e3cdbf5e041a86efd)), feat [#5](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/issues/5)

* add basic i18n to support feishu_card_build_status and add basic tpl ([abb0bb7d](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/abb0bb7d7f5fdb50cc5186e5abb39ecbfbbaf03f))

### 📝 Documentation

* rEMDE.md Features item added ([b260657a](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/b260657a473cf98c7189a5488639a4b13eb38122))

### ♻ Refactor

* add card template test for different lang ([c6548fe4](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/c6548fe4105b33c7b8698acfc2650f8ff727ab7e))

* change render tast case more clear ([689ddc19](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/689ddc198c815fef1bf612c7d68e41f8231714fb))

* add zh-CN template and update test case ([97e1292f](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/97e1292f9fafc1944b55bd96e0b328d8f2f697b7))

* feishu card template management by embed Resource ([7f61f53c](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/7f61f53c2b16f7f321c14eef63851fd8427d13fd))

### 👷‍ Build System

* platform`linux/amd64 linux/386 linux/arm64/v8 linux/arm/v7 linux/ppc64le linux/s390x` ([06823610](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/068236108b377b0340e6e64e9a741833c4c4906a)), feat [#7](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/issues/7)

## [1.3.0](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/compare/1.2.0...v1.3.0) (2024-04-06)

### 🐛 Bug Fixes

* change plugin rander to support more info to show by `settings.feishu-notice-types` ([11f0b2a5](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/11f0b2a554ae666c0730bb46f7ff5619e9d0bb2c))

### ✨ Features

* change flag name with style `*-*`, and update usage of doc ([076cc10f](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/076cc10f04d22d426a3587f3994f1f2b8b316c78))

* support `NoticeTypeFileBrowser` and update unit test case more safe ([bc156278](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/bc156278ce9f3cd083404f84e5994a6bfe20022a))

### 📝 Documentation

* add woodpecker docs, and icon ([ba4e8154](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/ba4e815440f326b655f114da2242589db0f92eb9))

* add template usage at doc/README.md ([1345a216](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/1345a2160b523aafd90c7540e15050f519f4acd4))

* add docker bagdes and remove useless code of plugin ([785b68f3](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/785b68f3c7abb021d0f5dffa34b51e23edb96e54))

### ♻ Refactor

* change test code file and update dev doc at doc/README.md ([54892014](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/548920147b32c8a515d4273175e937ed7d2c8758))

## [1.2.0](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/compare/1.1.0...v1.2.0) (2024-03-06)

### ✨ Features

* update feishu_status_success_ignore and feishu_status_change_success logic ([5f199c7f](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/5f199c7f1c7a74c218a47dfd70ffe596ffb476bd))

## [1.1.0](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/compare/1.0.0...v1.1.0) (2024-03-06)

### ✨ Features

* github.com/woodpecker-kit/woodpecker-tools v1.14.0 ([c915aa91](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/c915aa91b1303dada1c497538b590563bc8811bf))

## 1.0.0 (2024-03-06)

### ✨ Features

* add final flag and add usage of doc ([e67e5d9d](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/e67e5d9dfc631637eafebcc234bf6a2a37823ad1))

* full notify of woodpecker for 2.0.+ ([77cb4534](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/77cb4534fbb5c7c4813a8b283e24d7540d92e75e))

### 👷‍ Build System

* github.com/woodpecker-kit/woodpecker-tools v1.13.0 ([88e409f2](https://github.com/woodpecker-kit/woodpecker-feishu-group-robot/commit/88e409f2b352931d0c102268590929b023a9c431))
