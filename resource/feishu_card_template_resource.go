package resource

import (
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/internal/embed_source"
	"path"
)

const (
	GroupFeishuCardTemplateTpl = "feishu_card_template_tpl"

	// ItemFeishuCardTemplateHeadTpl
	// this template use feishu_plugin.FeishuCardRenderRoot
	ItemFeishuCardTemplateHeadTpl = "feishu_card_template_head_tpl.golden"

	// ItemFeishuCardTemplateTailTpl
	// this template use feishu_plugin.FeishuCardRenderRoot
	ItemFeishuCardTemplateTailTpl = "feishu_card_template_tail_tpl.golden"
)

func FetchFeishuCardTemplateTplByName(name string, lang string) (string, error) {
	conventionalTitleEmbed, err := embed_source.GetResourceByLanguage(PathResourceRoot,
		path.Join(GroupFeishuCardTemplateTpl, name), lang)
	if err != nil {
		return "", err
	}
	return conventionalTitleEmbed.Data(), nil
}
