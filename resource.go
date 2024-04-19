package woodpecker_feishu_group_robot

import (
	"embed"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/constant"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/internal/embed_source"
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/resource"
	"path"
)

//go:embed package.json
var PackageJson string

var (
	//go:embed resource/feishu_card_template_tpl
	embedFeishuCardTemplateFiles embed.FS

	embedFeishuCardTemplateList = []string{
		path.Join(resource.GroupFeishuCardTemplateTpl, resource.ItemFeishuCardTemplateHeadTpl),
		path.Join(resource.GroupFeishuCardTemplateTpl, resource.ItemFeishuCardTemplateTailTpl),
	}

	//go:embed resource/feishu_card_build_status_tpl
	embedResourceFeishuCardBuildStatusFiles embed.FS

	embedFeishuCardBuildStatusList = []string{
		path.Join(resource.GroupFeishuCardBuildStatusTpl, resource.ItemFeishuCardBuildStatusTypeCommit),
		path.Join(resource.GroupFeishuCardBuildStatusTpl, resource.ItemFeishuCardBuildStatusTypePullRequest),
		path.Join(resource.GroupFeishuCardBuildStatusTpl, resource.ItemFeishuCardBuildStatusTypePullRequestClose),
		path.Join(resource.GroupFeishuCardBuildStatusTpl, resource.ItemFeishuCardBuildStatusTypeTag),
		path.Join(resource.GroupFeishuCardBuildStatusTpl, resource.ItemFeishuCardBuildStatusTypeRelease),
		path.Join(resource.GroupFeishuCardBuildStatusTpl, resource.ItemFeishuCardBuildStatusTypeCron),
	}

	//go:embed resource/feishu_card_transfer_file_browser_tpl
	embedResourceFeishuCardTransferFileBrowserFiles embed.FS

	embedFeishuCardTransferFileBrowserList = []string{
		path.Join(resource.GroupFeishuCardTransferFileBrowserTpl, resource.ItemFeishuCardTransferFileBrowser),
	}
)

func CheckAllResource(root string) error {
	embed_source.SettingResourceRootPath(root)

	for _, resPathItem := range embedFeishuCardTemplateList {
		err := embed_source.InitResourceGroupByLanguage(resource.PathResourceRoot, embedFeishuCardTemplateFiles, resPathItem, constant.SupportLanguage())
		if err != nil {
			return err
		}
	}

	for _, resPathItem := range embedFeishuCardBuildStatusList {
		err := embed_source.InitResourceGroupByLanguage(resource.PathResourceRoot, embedResourceFeishuCardBuildStatusFiles, resPathItem, constant.SupportLanguage())
		if err != nil {
			return err
		}
	}

	for _, resPathItem := range embedFeishuCardTransferFileBrowserList {
		err := embed_source.InitResourceGroupByLanguage(resource.PathResourceRoot, embedResourceFeishuCardTransferFileBrowserFiles, resPathItem, constant.SupportLanguage())
		if err != nil {
			return err
		}
	}

	return nil
}
