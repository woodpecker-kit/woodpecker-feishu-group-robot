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
	//go:embed resource/feishu_card_build_status_tpl
	embedResourceContributingDocFiles embed.FS

	embedResourceCCDocPathList = []string{
		path.Join(resource.GroupFeishuCardTpl, resource.ItemFeishuCardBuildStatusTypeCommit),
		path.Join(resource.GroupFeishuCardTpl, resource.ItemFeishuCardBuildStatusTypePullRequest),
		path.Join(resource.GroupFeishuCardTpl, resource.ItemFeishuCardBuildStatusTypePullRequestClose),
		path.Join(resource.GroupFeishuCardTpl, resource.ItemFeishuCardBuildStatusTypeTag),
		path.Join(resource.GroupFeishuCardTpl, resource.ItemFeishuCardBuildStatusTypeRelease),
		path.Join(resource.GroupFeishuCardTpl, resource.ItemFeishuCardBuildStatusTypeCron),
	}
)

func CheckAllResource(root string) error {
	embed_source.SettingResourceRootPath(root)

	for _, resPathItem := range embedResourceCCDocPathList {
		err := embed_source.InitResourceGroupByLanguage(resource.PathResourceRoot, embedResourceContributingDocFiles, resPathItem, constant.SupportLanguage())
		if err != nil {
			return err
		}
	}

	return nil
}
