package libHope


type getFreePortsResponse struct {
	Ports []int `json:"ports,omitempty"`
}

// Wrapper of nodep.GetFreePorts
// count means how many ports you need.
func GetFreePorts(count int) string {
	return ""
}

// Convert share text to XrayJson
// support XrayJson, v2rayN plain text, v2rayN base64 text, Clash.Meta yaml
func ConvertShareLinksToXrayJson(base64Text string) string {
	return ""
}

// Convert XrayJson to share links.
// VMess will generate VMessAEAD link.
func ConvertXrayJsonToShareLinks(base64Text string) string {
	return ""
}
