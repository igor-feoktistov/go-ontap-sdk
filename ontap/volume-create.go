package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type VolumeCreate struct {
	Base
	Params struct {
		XMLName    xml.Name
		VolumeName *volumeName
		VolumeCreateOptions
	}
}

type volumeName struct {
	XMLName xml.Name
	Name    string `xml:",innerxml"`
}

type VolumeCreateOptions struct {
	AntivirusOnAccessPolicy    string `xml:"antivirus-on-access-policy,omitempty"`
	CacheRetentionPriority     string `xml:"cache-retention-priority,omitempty"`
	CachingPolicy              string `xml:"caching-policy,omitempty"`
	ConstituentRole            string `xml:"constituent-role,omitempty"`
	ContainingAggregateName    string `xml:"containing-aggr-name,omitempty"`
	EfficiencyPolicy           string `xml:"efficiency-policy,omitempty"`
	Encrypt                    bool   `xml:"encrypt,omitempty"`
	ExcludedFromAutobalance    bool   `xml:"excluded-from-autobalance,omitempty"`
	ExportPolicy               string `xml:"export-policy,omitempty"`
	ExtentSize                 string `xml:"extent-size,omitempty"`
	FlexcachePolicy            string `xml:"flexcache-cache-policy,omitempty"`
	FlexcacheFillPolicy        string `xml:"flexcache-fill-policy,omitempty"`
	FlexcacheOriginVolumeName  string `xml:"flexcache-origin-volume-name,omitempty"`
	GroupID                    int    `xml:"group-id,omitempty"`
	IsJunctionActive           bool   `xml:"is-junction-active,omitempty"`
	IsNvfailEnabled            string `xml:"is-nvfail-enabled,omitempty"`
	IsspaceEnforcementLogical  bool   `xml:"is-space-enforcement-logical,omitempty"`
	IsSpaceReportingLogical    bool   `xml:"is-space-reporting-logical,omitempty"`
	IsVserverRoot              bool   `xml:"is-vserver-root,omitempty"`
	JunctionPath               string `xml:"junction-path,omitempty"`
	LanguageCode               string `xml:"language-code,omitempty"`
	MaxDirSize                 int    `xml:"max-dir-size,omitempty"`
	MaxWriteAllocBlocks        int    `xml:"max-write-alloc-blocks,omitempty"`
	PercentageSnapshotReserve  int    `xml:"percentage-snapshot-reserve"`
	QosAdaptivePolicyGroupName string `xml:"qos-adaptive-policy-group-name,omitempty"`
	QosPolicyGroupName         string `xml:"qos-policy-group-name,omitempty"`
	Size                       string `xml:"size,omitempty"`
	SnapshotPolicy             string `xml:"snapshot-policy,omitempty"`
	SpaceReserve               string `xml:"space-reserve,omitempty"`
	SpaceSlo                   string `xml:"space-slo,omitempty"`
	StorageService             string `xml:"storage-service,omitempty"`
	TieringPolicy              string `xml:"tiering-policy,omitempty"`
	UnixPermissions            string `xml:"unix-permissions,omitempty"`
	UserID                     int    `xml:"user-id,omitempty"`
	VMAlignSector              int    `xml:"vm-align-sector,omitempty"`
	VMAlignSuffix              string `xml:"vm-align-suffix,omitempty"`
	Volume                     string `xml:"volume,omitempty"`
	VolumeComment              string `xml:"volume-comment,omitempty"`
	VolumeSecurityStyle        string `xml:"volume-security-style,omitempty"`
	VolumeState                string `xml:"volume-state,omitempty"`
	VolumeType                 string `xml:"volume-type,omitempty"`
	VserverDrProtection        string `xml:"vserver-dr-protection,omitempty"`
}

func (c *Client) VolumeCreateAPI(options *VolumeCreateOptions) (*SingleResultResponse, *http.Response, error) {
	if c.VolumeCreate == nil {
		c.VolumeCreate = &VolumeCreate {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.VolumeCreate.Params.XMLName = xml.Name{Local: "volume-create"}
	}
	c.VolumeCreate.Base.Name = c.vserver
	c.VolumeCreate.Params.VolumeCreateOptions = *options
	r := SingleResultResponse{}
	res, err := c.VolumeCreate.get(c.VolumeCreate, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(VolumeCreateAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
