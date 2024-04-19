package feishu_plugin

import (
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/resource"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_template"
	"github.com/woodpecker-kit/woodpecker-transfer-data/wd_share_file_browser_upload"
)

type OssCardFileBrowserRender struct {
	CiInfo    wd_short_info.WoodpeckerInfoShort
	CardOssFB CardOssFB
}

type CardOssFB struct {

	// file browser HostUrl
	HostUrl string

	// IsSendSuccess
	// send result
	IsSendSuccess bool

	// IsTagResult
	// tag result
	IsTagResult bool

	// IsPullRequestResult
	// pull request
	IsPullRequestResult bool

	InfoUser string

	PageUrl          string
	IsRenderPassword bool
	PagePasswd       string
}

// renderOssCardFileBrowser render oss card file browser
// use OssCardFileBrowserRender
func renderOssCardFileBrowser(r *OssCardFileBrowserRender, lang string) (string, error) {
	tpl, err := resource.FetchFeishuCardTransferFileBrowserTplByName(resource.ItemFeishuCardTransferFileBrowser, lang)
	if err != nil {
		return "", err
	}
	out, err := wd_template.Render(tpl, *r)
	if err != nil {
		return "", err
	}
	return out, nil
}

func parseOssCardFileBrowserRender(shareData wd_share_file_browser_upload.WdShareFileBrowserUpload, info wd_short_info.WoodpeckerInfoShort) (OssCardFileBrowserRender, error) {
	out := OssCardFileBrowserRender{
		CiInfo: info,
		CardOssFB: CardOssFB{
			HostUrl:       shareData.HostUrl,
			IsSendSuccess: shareData.IsSendSuccess,

			InfoUser: shareData.FileBrowserUserName,

			PageUrl:    shareData.DownloadPage,
			PagePasswd: shareData.DownloadPasswd,
		},
	}

	formatOssCardFileBrowser(&out)
	return out, nil
}

func formatOssCardFileBrowser(r *OssCardFileBrowserRender) {
	r.CardOssFB.IsTagResult = false
	if r.CiInfo.Commit.Tag != "" {
		r.CardOssFB.IsTagResult = true
	}
	r.CardOssFB.IsPullRequestResult = false
	if r.CiInfo.Commit.PR != "" {
		r.CardOssFB.IsPullRequestResult = true
	}

	r.CardOssFB.IsRenderPassword = false
	if r.CardOssFB.PagePasswd != "" {
		r.CardOssFB.IsRenderPassword = true
	}
}
