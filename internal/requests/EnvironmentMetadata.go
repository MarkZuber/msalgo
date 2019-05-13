package requests

import "strings"

type EnvironmentMetadata struct {
	webRequestManager IWebRequestManager
	aliasMap          map[string]*CloudEnvironmentInfo
}

func CreateEnvironmentMetadata(webRequestManager IWebRequestManager) *EnvironmentMetadata {

	idps := []*CloudEnvironmentInfo{}

	idps = append(idps, CreateCloudEnvironmentInfo("login.microsoftonline.com", "login.windows.net", []string{"login.microsoftonline.com", "login.windows.net", "login.microsoft.com", "sts.windows.net"}))
	idps = append(idps, CreateCloudEnvironmentInfo("login.partner.microsoftonline.cn", "login.partner.microsoftonline.cn", []string{"login.partner.microsoftonline.cn", "login.chinacloudapi.cn"}))

	var aliasMap map[string]*CloudEnvironmentInfo

	for _, idp := range idps {
		for alias := range idp.GetAliases() {
			aliasMap[alias] = idp
		}
	}

	em := &EnvironmentMetadata{webRequestManager, aliasMap}
	return em
}

func (em *EnvironmentMetadata) QueryCloudEnvironmentInfoFromServer(environment string) *CloudEnvironmentInfo {
	cloudEnvironmentInfo := em.aliasMap[strings.ToLower(environment)]
	return cloudEnvironmentInfo
}

func (em *EnvironmentMetadata) updateAliasMap() {
}
