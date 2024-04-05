package feishu_plugin

import (
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_steps_transfer"
	"github.com/woodpecker-kit/woodpecker-transfer-data/wd_share_file_browser_upload"
)

// loadStepsFileBrowser
// for NoticeTypeFileBrowser data load
func (p *FeishuPlugin) loadStepsFileBrowser() error {
	var wdShareFileBrowserUpload wd_share_file_browser_upload.WdShareFileBrowserUpload
	errShareFileBrowserUpload := wd_steps_transfer.In(
		p.Settings.RootPath, p.Settings.StepsTransferPath, p.GetWoodPeckerInfo(),
		wd_share_file_browser_upload.WdShareKeyFileBrowserUpload,
		&wdShareFileBrowserUpload,
	)
	if errShareFileBrowserUpload != nil {
		wd_log.Warnf("load steps transfer by notice type [ %s ] err: %v", NoticeTypeFileBrowser, errShareFileBrowserUpload)
		return nil
	}

	errAddOssCardFileBrowserRender := p.addRenderFileBrowser(wdShareFileBrowserUpload)
	if errAddOssCardFileBrowserRender != nil {
		return errAddOssCardFileBrowserRender
	}
	return nil
}

// addOssCardFileBrowserRender
// for NoticeTypeFileBrowser data render
func (p *FeishuPlugin) addRenderFileBrowser(shareData wd_share_file_browser_upload.WdShareFileBrowserUpload) error {

	cardFileBrowserRender, err := parseOssCardFileBrowserRender(shareData, p.ShortInfo())
	if err != nil {
		wd_log.Warnf("addRenderFileBrowser err: %v", err)
		return err
	}
	p.innerData.renderOssCardFileBrowser = &cardFileBrowserRender

	return nil
}
