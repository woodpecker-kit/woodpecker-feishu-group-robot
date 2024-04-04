package feishu_plugin

import (
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
		return renderBuildStatusTypeCommit(shortInfo)
	case wd_info.EventPipelinePush:
		return renderBuildStatusTypeCommit(shortInfo)
	case wd_info.EventPipelinePullRequest:
		return renderBuildStatusTypePullRequest(shortInfo)
	case wd_info.EventPipelinePullRequestClose:
		return renderBuildStatusTypePullRequestClose(shortInfo)
	case wd_info.EventPipelineTag:
		return renderBuildStatusTypeTag(shortInfo)
	case wd_info.EventPipelineRelease:
		return renderBuildStatusTypeRelease(shortInfo)
	case wd_info.EventPipelineCron:
		return renderBuildStatusTypeCron(shortInfo)
	}
}

// tplFeishuCardBuildStatusTypeCommit
// this template use wd_short_info.WoodpeckerInfoShort
const tplFeishuCardBuildStatusTypeCommit = `      {
        "tag": "markdown",
        "content": "📝 Commit by {{ Commit.CommitAuthor.Username }} on **{{ Commit.CommitBranch }}**\nCommitCode: {{ Commit.Sha }}"
      },
      {
        "tag": "markdown",
        "content": "{{#success Build.Status }}✅{{/success}}{{#failure Build.Status }}❌{{/failure}} Build [#{{ Build.Number }}]({{ Build.LinkCi }}) {{ Build.Status }}"
      },
      {
        "tag": "markdown",
        "content": "**Commit:**\n{{ Commit.Message }}"
      },
      {
        "tag": "markdown",
        "content": "[See Git Link]({{ Commit.Link }}) | [See Build Details]({{ Build.LinkCi }})"
      },
      {
        "tag": "hr"
      },
`

func renderBuildStatusTypeCommit(short wd_short_info.WoodpeckerInfoShort) (string, error) {
	return wd_template.Render(tplFeishuCardBuildStatusTypeCommit, short)
}

// tplFeishuCardBuildStatusTypePullRequest
// this template use wd_short_info.WoodpeckerInfoShort
const tplFeishuCardBuildStatusTypePullRequest = `      {
        "tag": "markdown",
        "content": "🏗️ Pull Request: {{ Commit.SourceBranch }} -> {{ Commit.TargetBranch }}\nAuthor: **{{ Commit.CommitAuthor.Username }}** on **{{ Commit.CommitBranch }}** PR Url: [PR Link]({{ Commit.Link }})"
      },
      {
        "tag": "markdown",
        "content": "{{#success Build.Status }}✅{{/success}}{{#failure Build.Status }}❌{{/failure}} Build [#{{ Build.Number }}]({{ Build.LinkCi }}) {{ Build.Status }}"
      },
      {
        "tag": "markdown",
        "content": "**Commit:**\n{{ Commit.Message }}"
      },
      {
        "tag": "markdown",
        "content": "[See Git Link]({{ Commit.Link }}) | [See Build Details]({{ Build.LinkCi }})"
      },
      {
        "tag": "hr"
      },
`

func renderBuildStatusTypePullRequest(short wd_short_info.WoodpeckerInfoShort) (string, error) {
	return wd_template.Render(tplFeishuCardBuildStatusTypePullRequest, short)
}

// tplFeishuCardBuildStatusTypePullRequestClose
// this template use wd_short_info.WoodpeckerInfoShort
const tplFeishuCardBuildStatusTypePullRequestClose = `      {
        "tag": "markdown",
        "content": "🏗️ Pull Request close: {{ Commit.SourceBranch }} -> {{ Commit.TargetBranch }}\nAuthor: **{{ Commit.CommitAuthor.Username }}** on **{{ Commit.CommitBranch }}** PR Url: [PR Link]({{ Commit.Link }})"
      },
      {
        "tag": "markdown",
        "content": "{{#success Build.Status }}✅{{/success}}{{#failure Build.Status }}❌{{/failure}} Build [#{{ Build.Number }}]({{ Build.LinkCi }}) {{ Build.Status }}"
      },
      {
        "tag": "markdown",
        "content": "**Commit:**\n{{ Commit.Message }}"
      },
      {
        "tag": "markdown",
        "content": "[See Git Link]({{ Commit.Link }}) | [See Build Details]({{ Build.LinkCi }})"
      },
      {
        "tag": "hr"
      },
`

func renderBuildStatusTypePullRequestClose(short wd_short_info.WoodpeckerInfoShort) (string, error) {
	return wd_template.Render(tplFeishuCardBuildStatusTypePullRequestClose, short)
}

// tplFeishuCardBuildStatusTypeTag
// this template use wd_short_info.WoodpeckerInfoShort
const tplFeishuCardBuildStatusTypeTag = `      {
        "tag": "markdown",
        "content": "📦 **Tag:** {{ Commit.Tag }}\nCommitCode: {{ Commit.Sha }}"
      },
      {
        "tag": "markdown",
        "content": "{{#success Build.Status }}✅{{/success}}{{#failure Build.Status }}❌{{/failure}} Build [#{{ Build.Number }}]({{ Build.LinkCi }}) {{ Build.Status }}"
      },
      {
        "tag": "markdown",
        "content": "**Commit:**\n{{ Commit.Message }}"
      },
      {
        "tag": "markdown",
        "content": "[See Git Link]({{ Commit.Link }}) | [See Build Details]({{ Build.LinkCi }})"
      },
      {
        "tag": "hr"
      },
`

func renderBuildStatusTypeTag(short wd_short_info.WoodpeckerInfoShort) (string, error) {
	return wd_template.Render(tplFeishuCardBuildStatusTypeTag, short)
}

// tplFeishuCardBuildStatusTypeRelease
// this template use wd_short_info.WoodpeckerInfoShort
const tplFeishuCardBuildStatusTypeRelease = `      {
        "tag": "markdown",
        "content": "🔖 **Release:** {{ Build.Number }}\nCommitCode: {{ Commit.Sha }}"
      },
      {
        "tag": "markdown",
        "content": "{{#success Build.Status }}✅{{/success}}{{#failure Build.Status }}❌{{/failure}} Build [#{{ Build.Number }}]({{ Build.LinkCi }}) {{ Build.Status }}"
      },
      {
        "tag": "markdown",
        "content": "**Commit:**\n{{ Commit.Message }}"
      },
      {
        "tag": "markdown",
        "content": "[See Git Link]({{ Commit.Link }}) | [See Build Details]({{ Build.LinkCi }})"
      },
      {
        "tag": "hr"
      },
`

func renderBuildStatusTypeRelease(short wd_short_info.WoodpeckerInfoShort) (string, error) {
	return wd_template.Render(tplFeishuCardBuildStatusTypeRelease, short)
}

// tplFeishuCardBuildStatusTypeCron
// this template use wd_short_info.WoodpeckerInfoShort
const tplFeishuCardBuildStatusTypeCron = `      {
        "tag": "markdown",
        "content": "🤖 **Cron:** {{ Build.Number }}\nCommitCode: {{ Commit.Sha }}"
      },
      {
        "tag": "markdown",
        "content": "{{#success Build.Status }}✅{{/success}}{{#failure Build.Status }}❌{{/failure}} Build [#{{ Build.Number }}]({{ Build.LinkCi }}) {{ Build.Status }}"
      },
      {
        "tag": "markdown",
        "content": "[See Git Link]({{ Commit.Link }}) | [See Build Details]({{ Build.LinkCi }})"
      },
      {
        "tag": "hr"
      },
`

func renderBuildStatusTypeCron(short wd_short_info.WoodpeckerInfoShort) (string, error) {
	return wd_template.Render(tplFeishuCardBuildStatusTypeCron, short)
}
