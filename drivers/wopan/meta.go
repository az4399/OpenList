package template

import (
	"github.com/OpenListTeam/OpenList/v4/internal/driver"
	"github.com/OpenListTeam/OpenList/v4/internal/op"
)

type Addition struct {
	// Usually one of two
	driver.RootID
	// define other
	RefreshToken string `json:"refresh_token" help:"Provide a refresh token or an access token (at least 16 characters)"`
	FamilyID     string `json:"family_id" help:"Keep it empty if you want to use your personal drive"`
	SortRule     string `json:"sort_rule" type:"select" options:"name_asc,name_desc,time_asc,time_desc,size_asc,size_desc" default:"name_asc"`

	AccessToken string `json:"access_token" help:"Optional when a refresh token is provided; it is refreshed automatically when missing or too short"`

	// BypassUploadSNI enables SNI bypass for upload requests.
	// Enable this if you are on a network (e.g. China Mobile broadband) that
	// resets TLS connections to the upload endpoint based on SNI inspection.
	// When enabled, the SDK resolves the upload hostname via DoH (223.5.5.5)
	// and connects directly to the resolved IP without sending an SNI field.
	BypassUploadSNI bool `json:"bypass_upload_sni" help:"Enable SNI bypass for uploads (for users on networks that block the 联通云盘 upload endpoint via SNI inspection, e.g. China Mobile broadband)"`
}

var config = driver.Config{
	Name:              "WoPan",
	DefaultRoot:       "0",
	NoOverwriteUpload: true,
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &Wopan{}
	})
}
