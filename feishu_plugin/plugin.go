package feishu_plugin

import (
	feishuMessageApi "github.com/woodpecker-kit/woodpecker-feishu-group-robot/feishu_open_api"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"github.com/woodpecker-kit/woodpecker-transfer-data/wd_share_file_browser_upload"
)

type (
	// FeishuPlugin feishu_plugin all config
	FeishuPlugin struct {
		Name           string
		Version        string
		woodpeckerInfo *wd_info.WoodpeckerInfo
		wdShortInfo    *wd_short_info.WoodpeckerInfoShort
		onlyArgsCheck  bool

		Settings               Settings
		FeishuRobotMsgTemplate feishuMessageApi.FeishuRobotMsgTemplate

		FuncPlugin FuncPlugin `json:"-"`

		innerData innerData
	}

	innerData struct {
		// by wd_share_file_browser_upload.WdShareFileBrowserUpload
		renderOssCardFileBrowser *OssCardFileBrowserRender

		// copyRenderFeishuCard
		// this data will remove timestamp and sign
		copyRenderFeishuCard string
	}
)

type FuncPlugin interface {
	ShortInfo() wd_short_info.WoodpeckerInfoShort

	SetWoodpeckerInfo(info wd_info.WoodpeckerInfo)
	GetWoodPeckerInfo() wd_info.WoodpeckerInfo

	OnlyArgsCheck()

	Exec() error

	loadStepsTransfer() error
	checkArgs() error
	saveStepsTransfer() error

	saveDeepCopyInnerData() error
	GetInnerCopyRenderFeishuCardContext() string

	// loadStepsFileBrowser
	// for NoticeTypeFileBrowser data load
	loadStepsFileBrowser() error
	// addOssCardFileBrowserRender
	// for NoticeTypeFileBrowser data render
	addRenderFileBrowser(shareData wd_share_file_browser_upload.WdShareFileBrowserUpload) error
}
