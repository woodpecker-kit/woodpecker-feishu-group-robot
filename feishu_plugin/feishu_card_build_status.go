package feishu_plugin

import (
	"github.com/woodpecker-kit/woodpecker-feishu-group-robot/resource"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_template"
)

func renderBuildStatus(p FeishuPlugin, buildStatus string) (string, error) {
	shortInfo := p.ShortInfo()
	if shortInfo.Build.Status != buildStatus {
		wd_log.Warnf("build status not match, expect: %s, got: %s, just use settings status build status [ %s ]", buildStatus, shortInfo.Build.Status, buildStatus)
		shortInfo.Build.Status = buildStatus
	}

	switch shortInfo.Build.Event {
	default:
		return renderBuildStatusTypeCommit(shortInfo, p.Settings.I18nLangSet)
	case wd_info.EventPipelinePush:
		return renderBuildStatusTypeCommit(shortInfo, p.Settings.I18nLangSet)
	case wd_info.EventPipelinePullRequest:
		return renderBuildStatusTypePullRequest(shortInfo, p.Settings.I18nLangSet)
	case wd_info.EventPipelinePullRequestClose:
		return renderBuildStatusTypePullRequestClose(shortInfo, p.Settings.I18nLangSet)
	case wd_info.EventPipelineTag:
		return renderBuildStatusTypeTag(shortInfo, p.Settings.I18nLangSet)
	case wd_info.EventPipelineRelease:
		return renderBuildStatusTypeRelease(shortInfo, p.Settings.I18nLangSet)
	case wd_info.EventPipelineCron:
		return renderBuildStatusTypeCron(shortInfo, p.Settings.I18nLangSet)
	}
}

// renderBuildStatusTypeCommit
// this template use wd_short_info.WoodpeckerInfoShort
func renderBuildStatusTypeCommit(short wd_short_info.WoodpeckerInfoShort, lang string) (string, error) {

	tpl, err := resource.FetchFeishuCardBuildStatusTplByName(resource.ItemFeishuCardBuildStatusTypeCommit, lang)
	if err != nil {
		return "", err
	}

	return wd_template.Render(tpl, short)
}

// renderBuildStatusTypePullRequest
// this template use wd_short_info.WoodpeckerInfoShort
func renderBuildStatusTypePullRequest(short wd_short_info.WoodpeckerInfoShort, lang string) (string, error) {
	tpl, err := resource.FetchFeishuCardBuildStatusTplByName(resource.ItemFeishuCardBuildStatusTypePullRequest, lang)
	if err != nil {
		return "", err
	}

	return wd_template.Render(tpl, short)
}

// renderBuildStatusTypePullRequestClose
// this template use wd_short_info.WoodpeckerInfoShort
func renderBuildStatusTypePullRequestClose(short wd_short_info.WoodpeckerInfoShort, lang string) (string, error) {
	tpl, err := resource.FetchFeishuCardBuildStatusTplByName(resource.ItemFeishuCardBuildStatusTypePullRequestClose, lang)
	if err != nil {
		return "", err
	}

	return wd_template.Render(tpl, short)
}

// renderBuildStatusTypeTag
// this template use wd_short_info.WoodpeckerInfoShort
func renderBuildStatusTypeTag(short wd_short_info.WoodpeckerInfoShort, lang string) (string, error) {
	tpl, err := resource.FetchFeishuCardBuildStatusTplByName(resource.ItemFeishuCardBuildStatusTypeTag, lang)
	if err != nil {
		return "", err
	}

	return wd_template.Render(tpl, short)
}

// renderBuildStatusTypeRelease
// this template use wd_short_info.WoodpeckerInfoShort
func renderBuildStatusTypeRelease(short wd_short_info.WoodpeckerInfoShort, lang string) (string, error) {
	tpl, err := resource.FetchFeishuCardBuildStatusTplByName(resource.ItemFeishuCardBuildStatusTypeRelease, lang)
	if err != nil {
		return "", err
	}

	return wd_template.Render(tpl, short)
}

// renderBuildStatusTypeCron
// this template use wd_short_info.WoodpeckerInfoShort
func renderBuildStatusTypeCron(short wd_short_info.WoodpeckerInfoShort, lang string) (string, error) {
	tpl, err := resource.FetchFeishuCardBuildStatusTplByName(resource.ItemFeishuCardBuildStatusTypeCron, lang)
	if err != nil {
		return "", err
	}

	return wd_template.Render(tpl, short)

}
