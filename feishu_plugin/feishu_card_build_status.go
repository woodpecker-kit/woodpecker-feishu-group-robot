package feishu_plugin

import (
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/resource"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_template"
)

func renderBuildStatus(p FeishuPlugin, buildStatus string, lang string) (string, error) {
	shortInfo := p.ShortInfo()
	if shortInfo.Build.Status != buildStatus {
		wd_log.Warnf("build status not match, expect: %s, got: %s, just use settings status build status [ %s ]", buildStatus, shortInfo.Build.Status, buildStatus)
		shortInfo.Build.Status = buildStatus
	}

	switch shortInfo.Build.Event {
	default:
		return renderBuildStatusTypeCommit(shortInfo, lang)
	case wd_info.EventPipelinePush:
		return renderBuildStatusTypeCommit(shortInfo, lang)
	case wd_info.EventPipelinePullRequest:
		return renderBuildStatusTypePullRequest(shortInfo, lang)
	case wd_info.EventPipelinePullRequestClose:
		return renderBuildStatusTypePullRequestClose(shortInfo, lang)
	case wd_info.EventPipelineTag:
		return renderBuildStatusTypeTag(shortInfo, lang)
	case wd_info.EventPipelineRelease:
		return renderBuildStatusTypeRelease(shortInfo, lang)
	case wd_info.EventPipelineCron:
		return renderBuildStatusTypeCron(shortInfo, lang)
	}
}

// renderBuildStatusTypeCommit
// this template use wd_short_info.WoodpeckerInfoShort
func renderBuildStatusTypeCommit(short wd_short_info.WoodpeckerInfoShort, lang string) (string, error) {

	tpl, err := resource.FetchFeishuCardTplByName(resource.ItemFeishuCardBuildStatusTypeCommit, lang)
	if err != nil {
		return "", err
	}

	return wd_template.Render(tpl, short)
}

// renderBuildStatusTypePullRequest
// this template use wd_short_info.WoodpeckerInfoShort
func renderBuildStatusTypePullRequest(short wd_short_info.WoodpeckerInfoShort, lang string) (string, error) {
	tpl, err := resource.FetchFeishuCardTplByName(resource.ItemFeishuCardBuildStatusTypePullRequest, lang)
	if err != nil {
		return "", err
	}

	return wd_template.Render(tpl, short)
}

// renderBuildStatusTypePullRequestClose
// this template use wd_short_info.WoodpeckerInfoShort
func renderBuildStatusTypePullRequestClose(short wd_short_info.WoodpeckerInfoShort, lang string) (string, error) {
	tpl, err := resource.FetchFeishuCardTplByName(resource.ItemFeishuCardBuildStatusTypePullRequestClose, lang)
	if err != nil {
		return "", err
	}

	return wd_template.Render(tpl, short)
}

// renderBuildStatusTypeTag
// this template use wd_short_info.WoodpeckerInfoShort
func renderBuildStatusTypeTag(short wd_short_info.WoodpeckerInfoShort, lang string) (string, error) {
	tpl, err := resource.FetchFeishuCardTplByName(resource.ItemFeishuCardBuildStatusTypeTag, lang)
	if err != nil {
		return "", err
	}

	return wd_template.Render(tpl, short)
}

// renderBuildStatusTypeRelease
// this template use wd_short_info.WoodpeckerInfoShort
func renderBuildStatusTypeRelease(short wd_short_info.WoodpeckerInfoShort, lang string) (string, error) {
	tpl, err := resource.FetchFeishuCardTplByName(resource.ItemFeishuCardBuildStatusTypeRelease, lang)
	if err != nil {
		return "", err
	}

	return wd_template.Render(tpl, short)
}

// renderBuildStatusTypeCron
// this template use wd_short_info.WoodpeckerInfoShort
func renderBuildStatusTypeCron(short wd_short_info.WoodpeckerInfoShort, lang string) (string, error) {
	tpl, err := resource.FetchFeishuCardTplByName(resource.ItemFeishuCardBuildStatusTypeCron, lang)
	if err != nil {
		return "", err
	}

	return wd_template.Render(tpl, short)

}
