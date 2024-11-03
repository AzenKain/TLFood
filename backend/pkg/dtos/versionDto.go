package dtos

type CreateVersionStarRailDto struct {
	VersionName         string `json:"version_name" validate:"required"`
	AssetBundleURL      string `json:"asset_bundle_url" validate:"required" `
	ExResourceURL       string `json:"ex_resource_url" validate:"required"`
	LuaURL              string `json:"lua_url" validate:"required"`
	IfixURL             string `json:"ifixUrl"`
	LuaVersion          string `json:"lua_version" `
	CustomMdkResVersion int    `json:"customMdkResVersion"`
	CustomIfixVersion   int    `json:"customIfixVersion"`
}

type UpdateVersionStarRailDto struct {
	Id                  string `json:"version_id" validate:"required"`
	VersionName         string `json:"version_name" `
	AssetBundleURL      string `json:"asset_bundle_url"`
	ExResourceURL       string `json:"ex_resource_url" `
	LuaURL              string `json:"lua_url" `
	IfixURL             string `json:"ifixUrl"`
	LuaVersion          string `json:"lua_version" `
	CustomMdkResVersion int    `json:"customMdkResVersion"`
	CustomIfixVersion   int    `json:"customIfixVersion"`
}
