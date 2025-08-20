// libXray is an Xray wrapper focusing on improving the experience of Xray-core mobile development.
package libHope

import (
	"encoding/base64"
	"encoding/json"

	"github.com/asimov1234/libnewv/nodep"
	"github.com/asimov1234/libnewv/xray"
)

type loadGeoDataRequest struct {
	DatDir  string `json:"datDir,omitempty"`
	Name    string `json:"name,omitempty"`
	GeoType string `json:"geoType,omitempty"`
}

// Read geo data and write all codes to text file.
func LoadGeoData(base64Text string) string {
	var response nodep.CallResponse[string]
	req, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	var request loadGeoDataRequest
	err = json.Unmarshal(req, &request)
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	err = xray.LoadGeoData(request.DatDir, request.Name, request.GeoType)
	return response.EncodeToBase64("", err)
}

type pingRequest struct {
	DatDir     string `json:"datDir,omitempty"`
	ConfigPath string `json:"configPath,omitempty"`
	Timeout    int    `json:"timeout,omitempty"`
	Url        string `json:"url,omitempty"`
	Proxy      string `json:"proxy,omitempty"`
}

// Ping Xray config and get the delay of its outbound.
func Ping(base64Text string) string {
	var response nodep.CallResponse[int64]
	req, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return response.EncodeToBase64(nodep.PingDelayError, err)
	}
	var request pingRequest
	err = json.Unmarshal(req, &request)
	if err != nil {
		return response.EncodeToBase64(nodep.PingDelayError, err)
	}
	delay, err := xray.Ping(request.DatDir, request.ConfigPath, request.Timeout, request.Url, request.Proxy)
	return response.EncodeToBase64(delay, err)
}

// query inbound and outbound stats.
func QueryStats(base64Text string) string {
	var response nodep.CallResponse[string]
	server, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return response.EncodeToBase64("", err)
	}

	stats, err := xray.QueryStats(string(server))
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	return response.EncodeToBase64(stats, nil)
}

// convert text to uuid
func CustomUUID(base64Text string) string {
	var response nodep.CallResponse[string]
	text, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	uuid := xray.CustomUUID(string(text))
	return response.EncodeToBase64(uuid, nil)
}

type TestHopeRequest struct {
	DatDir     string `json:"datDir,omitempty"`
	ConfigPath string `json:"configPath,omitempty"`
}

// Test Xray Config.
func TestHope(base64Text string) string {
	var response nodep.CallResponse[string]
	req, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	var request TestHopeRequest
	err = json.Unmarshal(req, &request)
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	err = xray.TestXray(request.DatDir, request.ConfigPath)
	return response.EncodeToBase64("", err)
}

type RunHopeRequest struct {
	DatDir     string `json:"datDir,omitempty"`
	ConfigPath string `json:"configPath,omitempty"`
}

// Run Xray instance.
func RunHope(base64Text string) string {
	var response nodep.CallResponse[string]
	req, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	var request RunHopeRequest
	err = json.Unmarshal(req, &request)
	if err != nil {
		return response.EncodeToBase64("", err)
	}
	err = xray.RunXray(request.DatDir, request.ConfigPath)
	return response.EncodeToBase64("", err)
}

// Stop Xray instance.
func StopHope() string {
	var response nodep.CallResponse[string]
	err := xray.StopXray()
	return response.EncodeToBase64("", err)
}

// Xray's version
func XrayVersion() string {
	var response nodep.CallResponse[string]
	return response.EncodeToBase64(xray.XrayVersion(), nil)
}
