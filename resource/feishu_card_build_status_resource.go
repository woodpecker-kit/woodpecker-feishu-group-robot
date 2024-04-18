package resource

import (
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/internal/embed_source"
	"path"
)

const (
	GroupFeishuCardTpl = "feishu_card_build_status_tpl"

	// ItemFeishuCardBuildStatusTypeCommit
	// this template use wd_short_info.WoodpeckerInfoShort
	ItemFeishuCardBuildStatusTypeCommit = "build_status_type_commit.golden"

	// ItemFeishuCardBuildStatusTypePullRequest
	// this template use wd_short_info.WoodpeckerInfoShort
	ItemFeishuCardBuildStatusTypePullRequest = "build_status_type_pull_request.golden"

	// ItemFeishuCardBuildStatusTypePullRequestClose
	// this template use wd_short_info.WoodpeckerInfoShort
	ItemFeishuCardBuildStatusTypePullRequestClose = "build_status_type_pull_request_close.golden"

	// ItemFeishuCardBuildStatusTypeTag
	// this template use wd_short_info.WoodpeckerInfoShort
	ItemFeishuCardBuildStatusTypeTag = "build_status_type_tag.golden"

	// ItemFeishuCardBuildStatusTypeRelease
	// this template use wd_short_info.WoodpeckerInfoShort
	ItemFeishuCardBuildStatusTypeRelease = "build_status_type_release.golden"

	// ItemFeishuCardBuildStatusTypeCron
	// this template use wd_short_info.WoodpeckerInfoShort
	ItemFeishuCardBuildStatusTypeCron = "build_status_type_cron.golden"
)

func FetchFeishuCardTplByName(name string, lang string) (string, error) {
	conventionalTitleEmbed, err := embed_source.GetResourceByLanguage(PathResourceRoot,
		path.Join(GroupFeishuCardTpl, name), lang)
	if err != nil {
		return "", err
	}
	return conventionalTitleEmbed.Data(), nil
}
