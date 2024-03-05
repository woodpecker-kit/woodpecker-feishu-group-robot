package cli

import (
	"fmt"
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/constant"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/feishu_plugin"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/internal/pkgJson"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/internal/version_check"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_urfave_cli_v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_urfave_cli_v2/cli_exit_urfave"
	"os"
)

var wdPlugin *feishu_plugin.FeishuPlugin

// GlobalBeforeAction
// do command Action before flag global.
func GlobalBeforeAction(c *cli.Context) error {
	isDebug := wd_urfave_cli_v2.IsBuildDebugOpen(c)
	if isDebug {
		wd_log.OpenDebug()
		allEnvPrintStr := env_kit.FindAllEnv4PrintAsSortJust(36)
		wd_log.Debugf("==> feishu_plugin start with all env:\n%s", allEnvPrintStr)
	}
	namePlugin := pkgJson.GetPackageJsonName()
	cliVersion := pkgJson.GetPackageJsonVersionGoStyle(true)

	// bind woodpeckerInfo
	woodpeckerInfo := wd_urfave_cli_v2.UrfaveCliBindInfo(c)

	if namePlugin == "" {
		return cli_exit_urfave.ErrMsg("missing name, please set name")
	}

	if cliVersion == "" {
		return cli_exit_urfave.ErrMsg("missing version, please set version")
	}

	errVersionConstraint := version_check.SemverVersionConstraint(cliVersion, constant.VersionSupportMinimum, constant.VersionSupportMaximum)
	if errVersionConstraint != nil {
		return cli_exit_urfave.Err(errVersionConstraint)
	}

	errCheckVersion := wd_info.CiSystemVersionMinimumSupport(woodpeckerInfo)
	if errCheckVersion != nil {
		return cli_exit_urfave.Err(errCheckVersion)
	}

	wd_log.Debugf("cli version is %s\n", cliVersion)
	wd_log.Debugf("load woodpecker finish at repo link: %v\n", woodpeckerInfo.RepositoryInfo.CIRepoURL)

	rootPath, errRootPath := os.Getwd()
	if errRootPath != nil {
		return cli_exit_urfave.Err(errRootPath)
	}
	stepsTransferFilePath := c.String(constant.NameCliPluginStepsTransferFilePath)

	pluginBind, errBindPlugin := feishu_plugin.BindCliFlags(c, namePlugin, cliVersion, &woodpeckerInfo, rootPath, stepsTransferFilePath)
	if errBindPlugin != nil {
		return cli_exit_urfave.Err(errBindPlugin)
	}

	wdPlugin = pluginBind

	return nil
}

// GlobalAction
// do cli Action before flag.
func GlobalAction(c *cli.Context) error {
	if wdPlugin == nil {
		panic(fmt.Errorf("must success run GlobalBeforeAction then run GlobalAction"))
	}

	err := wdPlugin.Exec()
	if err != nil {
		return cli_exit_urfave.Err(err)
	}

	return nil
}

// GlobalAfterAction
//
//	do command Action after flag global.
//
//nolint:golint,unused
func GlobalAfterAction(c *cli.Context) error {
	if wdPlugin != nil {
		wd_log.Infof("=> finish run: %s, version: %s\n", wdPlugin.Name, wdPlugin.Version)
	}
	return nil
}
