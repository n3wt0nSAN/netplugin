package iotapi

// EPAddURL is the rest point for adding an endpoint type device IoT
const EPIoTAddURL = "/ContivIOT.AddIotDev"

// EPDelURL is the rest point for deleting an endpoint
//const EPDelURL = "/ContivCNI.DelPod"

// IOTDevAttr holds attributes of the pod to be attached or detached
type IOTDevAttr struct {
	Name          string `json:"IOT_DEV_NAME,omitempty"`
	Tenant        string `json:"IOT_DEV_TENANT,omitempty"`
	InfraIotDevID string `json:"IOT_DEV_INFRA_ID,omitempty"`
	IntfName      string `json:"IOT_DEV_IFNAME,omitempty"`
	Network       string `json:"IOT_DEV_NETWORK,omitempty"`
	Group         string `json:"IOT_DEV_GROUP,omitempty"`
}

// RspAddIotDev contains the response to the AddIotDev
type RspAddIotDev struct {
	EndpointID string `json:"endpointid,omitempty"`
	ErrMsg     string `json:"errmsg,omitempty"`
	ErrInfo    string `json:"errinfo,omitempty"`
	Attr       *Attr  `json:"epattr,omitempty"`
}

type Attr struct {
	IPAddress   string `json:"ipaddress,omitempty"`
	MacAddress  string `json:"macaddress,omitempty"`
	PortName    string `json:"portname,omitempty"`
	Gateway     string `json:"gateway,omitempty"`
	IPv6Address string `json:"ipv6address,omitempty"`
	IPv6Gateway string `json:"ipv6gateway,omitempty"`
}
