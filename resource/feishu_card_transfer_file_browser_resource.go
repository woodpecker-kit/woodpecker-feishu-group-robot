package resource

import (
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/internal/embed_source"
	"path"
)

const (
	GroupFeishuCardTransferFileBrowserTpl = "feishu_card_transfer_file_browser_tpl"

	// ItemFeishuCardTransferFileBrowser
	// this template use feishu_plugin.OssCardFileBrowserRender
	ItemFeishuCardTransferFileBrowser = "feishu_card_transfer_file_browser.golden"
)

func FetchFeishuCardTransferFileBrowserTplByName(name string, lang string) (string, error) {
	conventionalTitleEmbed, err := embed_source.GetResourceByLanguage(PathResourceRoot,
		path.Join(GroupFeishuCardTransferFileBrowserTpl, name), lang)
	if err != nil {
		return "", err
	}
	return conventionalTitleEmbed.Data(), nil
}
