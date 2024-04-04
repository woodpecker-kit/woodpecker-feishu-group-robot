package feishu_plugin_transfer

const OssSendTransferKey = "oss_send_transfer"

type OssUserProfile struct {
	UserName  string
	UserToken string
}

type OssSendTransfer struct {

	// OssHost
	// must Oss HostUrl to access oss resource
	OssHost string

	// InfoSendResult
	// send result [ success or failure]
	InfoSendResult string

	OssUserProfile OssUserProfile

	// OssPath
	// oss path at remote path
	OssPath string
	// ResourceUrl
	// oss resource url, this is absolute url to access oss resource
	ResourceUrl string
	// PagePasswd
	// if page_passwd is not empty, must show PageUrl
	PagePasswd string
	// PageUrl
	// if page_passwd is empty, use ResourceUrl
	PageUrl string
}
