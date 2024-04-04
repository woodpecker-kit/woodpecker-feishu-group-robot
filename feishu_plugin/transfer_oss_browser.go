package feishu_plugin

import (
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_template"
	"github.com/woodpecker-kit/woodpecker-transfer-data/wd_share_file_browser_upload"
)

// renderFileBrowser
// use OssCardFileBrowserRender
const renderFileBrowser = `      {{#if CardOssFB.IsSendSuccess }}
{{#if CardOssFB.IsTagResult }}
      {
        "tag": "markdown",
        "content": "📦 **Tag:** {{ CiInfo.Commit.Tag }}\nCommitCode: {{ CiInfo.Commit.Sha }}"
      },
{{else if CardOssFB.IsPullRequestResult}}
      {
        "tag": "markdown",
        "content": "🏗️ Pull Request: {{ CiInfo.Commit.SourceBranch }} -> {{ CiInfo.Commit.TargetBranch }} [#{{ CiInfo.Commit.PR }}]({{ CiInfo.Commit.Url }})"
      },
{{else}}
      {
        "tag": "markdown",
        "content": ":📝 Commit by [#{{ CiInfo.Commit.PR }}]({{ CiInfo.Commit.Url }})"
      },
{{/if}}
{{else}}
      {
        "tag": "markdown",
        "content": "send file browser to [{{ CardOssFB.HostUrl }}]({{ CardOssFB.HostUrl }}) failed, please check log"
      },
{{/if}}
`

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

	PageUrl    string
	PagePasswd string

	IsRenderResourceUrl bool
	ResourceUrl         string
}

func renderOssCardFileBrowser(r *OssCardFileBrowserRender) (string, error) {

	out, err := wd_template.Render(renderFileBrowser, &r)
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

			PageUrl:    shareData.ResourceUrl,
			PagePasswd: shareData.DownloadPasswd,

			ResourceUrl: shareData.ResourceUrl,
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

	r.CardOssFB.IsRenderResourceUrl = false
	if r.CardOssFB.PagePasswd == "" {
		r.CardOssFB.IsRenderResourceUrl = true
	}
}
