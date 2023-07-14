package ontap

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type SystemNodeInfo struct {
	SystemNodeAttributes *SystemNodeAttributes `xml:"node-details-info,omitempty"`
}

type SystemNodeGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		SystemNodeGetOptions
	}
}

type SystemNodeGetOptions struct {
	DesiredAttributes *SystemNodeInfo `xml:"desired-attributes"`
	MaxRecords        int             `xml:"max-records,omitempty"`
	Query             *SystemNodeInfo `xml:"query"`
	Tag               string          `xml:"tag,omitempty"`
}

type SystemNodeGetResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		SystemNodeAttributes *SystemNodeGetIterResponseResultAttributesList `xml:"attributes-list"`
		NextTag              string                                         `xml:"next-tag"`
		NumRecords           int                                            `xml:"num-records"`
	} `xml:"results"`
}

type SystemNodeGetIterResponseResultAttributesList struct {
	XMLName         xml.Name               `xml:"attributes-list"`
	NodeDetailsInfo []SystemNodeAttributes `xml:"node-details-info"`
}

type SystemNodeAttributes struct {
	XMLName                     xml.Name `xml:"node-details-info"`
	CpuBusytime                 int      `xml:"cpu-busytime"`
	CpuFirmwareRelease          string   `xml:"cpu-firmware-release"`
	EnvFailedFanCount           int      `xml:"env-failed-fan-count"`
	EnvFailedFanMessage         string   `xml:"env-failed-fan-message"`
	EnvFailedPowerSupplyCount   int      `xml:"env-failed-power-supply-count"`
	EnvFailedPowerSupplyMessage string   `xml:"env-failed-power-supply-message"`
	EnvOverTemperature          bool     `xml:"env-over-temperature"`
	IsAllFlashOptimized         bool     `xml:"is-all-flash-optimized"`
	IsAllFlashSelectOptimized   bool     `xml:"is-all-flash-select-optimized"`
	IsCapacityOptimized         bool     `xml:"is-capacity-optimized"`
	IsCloudOptimized            bool     `xml:"is-cloud-optimized"`
	IsDiffSvcs                  bool     `xml:"is-diff-svcs"`
	IsEpsilonNode               bool     `xml:"is-epsilon-node"`
	IsNodeClusterEligible       bool     `xml:"is-node-cluster-eligible"`
	IsNodeHealthy               bool     `xml:"is-node-healthy"`
	IsPerfOptimized             bool     `xml:"is-perf-optimized"`
	MaximumAggregateSize        int      `xml:"maximum-aggregate-size"`
	MaximumNumberOfVolumes      int      `xml:"maximum-number-of-volumes"`
	MaximumVolumeSize           int      `xml:"maximum-volume-size"`
	Node                        string   `xml:"node"`
	NodeAssetTag                string   `xml:"node-asset-tag"`
	NodeLocation                string   `xml:"node-location"`
	NodeModel                   string   `xml:"node-model"`
	NodeNvramId                 int      `xml:"node-nvram-id"`
	NodeOwner                   string   `xml:"node-owner"`
	NodeSerialNumber            string   `xml:"node-serial-number"`
	NodeStorageConfiguration    string   `xml:"node-storage-configuration"`
	NodeSystemId                string   `xml:"node-system-id"`
	NodeUptime                  int      `xml:"node-uptime"`
	NodeUuid                    string   `xml:"node-uuid"`
	NodeVendor                  string   `xml:"node-vendor"`
	NvramBatteryStatus          string   `xml:"nvram-battery-status"`
	ProductVersion              string   `xml:"product-version"`
	Sas2Sas3MixedStackSupport   string   `xml:"sas2-sas3-mixed-stack-support"`
	VmSystemDisks               string   `xml:"vm-system-disks"`
	VmhostInfo                  string   `xml:"vmhost-info"`
}

func (c *Client) SystemNodeGetAPI(options *SystemNodeGetOptions) (*SystemNodeGetResponse, *http.Response, error) {
	if c.SystemNodeGetIter == nil {
		c.SystemNodeGetIter = &SystemNodeGetIter{
			Base: Base{
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.SystemNodeGetIter.Params.XMLName = xml.Name{Local: "system-node-get-iter"}
	}
	c.SystemNodeGetIter.Params.SystemNodeGetOptions = *options
	r := SystemNodeGetResponse{}
	res, err := c.SystemNodeGetIter.get(c.SystemNodeGetIter, &r)
	if err == nil && r.Results.Passed() == false {
		err = fmt.Errorf("error(SystemNodeGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) SystemNodeGetIterAPI(options *SystemNodeGetOptions) (responseNodes []*SystemNodeGetResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.SystemNodeGetAPI(options)
		if err != nil {
			break
		} else {
			nextTag = r.Results.NextTag
			fmt.Printf("nextTag: %s", nextTag)
			fmt.Printf("%s", nextTag)
			responseNodes = append(responseNodes, r)
			if nextTag == "" {
				fmt.Print("nextTag is empty")
				break
			}
			options.Tag = nextTag
		}
	}

	return
}
