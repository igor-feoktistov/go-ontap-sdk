package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type VolumeGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		*VolumeGetOptions
	}
}

type VolumeGetOptions struct {
	DesiredAttributes *VolumeQuery `xml:"desired-attributes,omitempty"`
	Attributes        *VolumeQuery `xml:"attributes,omitempty"`
	MaxRecords        int          `xml:"max-records,omitempty"`
	Query             *VolumeQuery `xml:"query,omitempty"`
	Tag               string       `xml:"tag,omitempty"`
}

type VolumeQuery struct {
	VolumeInfo *VolumeInfo `xml:"volume-attributes,omitempty"`
}

type VolumeAntivirusAttributes struct {
	OnAccessPolicy string `xml:"on-access-policy"`
}

type VolumeAutobalanceAttributes struct {
	IsAutobalanceEligible string `xml:"is-autobalance-eligible"`
}

type VolumeAutosizeAttributes struct {
	GrowThresholdPercent   string `xml:"grow-threshold-percent"`
	IsEnabled              string `xml:"is-enabled"`
	MaximumSize            string `xml:"maximum-size"`
	MinimumSize            string `xml:"minimum-size"`
	Mode                   string `xml:"mode"`
	ShrinkThresholdPercent string `xml:"shrink-threshold-percent"`
}

type VolumeDirectoryAttributes struct {
	I2PEnabled string `xml:"i2p-enabled"`
	MaxDirSize string `xml:"max-dir-size"`
	RootDirGen string `xml:"root-dir-gen"`
}

type VolumeExportAttributes struct {
	Policy string `xml:"policy"`
}

type VolumeHybridCacheAttributes struct {
	CacheRetentionPriority string `xml:"cache-retention-priority"`
	CachingPolicy          string `xml:"caching-policy"`
	Eligibility            string `xml:"eligibility"`
}

type VolumeIDAttributes struct {
	AggrList                []string `xml:"aggr-list>aggr-name,omitempty"`
	Comment                 string   `xml:"comment,omitempty"`
	ContainingAggregateName string   `xml:"containing-aggregate-name,omitempty"`
	ContainingAggregateUUID string   `xml:"containing-aggregate-uuid,omitempty"`
	CreationTime            string   `xml:"creation-time,omitempty"`
	Dsid                    string   `xml:"dsid,omitempty"`
	Fsid                    string   `xml:"fsid,omitempty"`
	InstanceUUID            string   `xml:"instance-uuid,omitempty"`
	JunctionParentName      string   `xml:"junction-parent-name,omitempty"`
	JunctionPath            string   `xml:"junction-path,omitempty"`
	Msid                    string   `xml:"msid,omitempty"`
	Name                    string   `xml:"name,omitempty"`
	NameOrdinal             string   `xml:"name-ordinal,omitempty"`
	Node                    string   `xml:"node,omitempty"`
	Nodes                   []string `xml:"nodes>node-name,omitempty"`
	OwningVserverName       string   `xml:"owning-vserver-name,omitempty"`
	OwningVserverUUID       string   `xml:"owning-vserver-uuid,omitempty"`
	ProvenanceUUID          string   `xml:"provenance-uuid,omitempty"`
	Style                   string   `xml:"style,omitempty"`
	StyleExtended           string   `xml:"style-extended,omitempty"`
	Type                    string   `xml:"type,omitempty"`
	UUID                    string   `xml:"uuid,omitempty"`
}

type VolumeInodeAttributes struct {
	BlockType                string `xml:"block-type"`
	FilesPrivateUsed         string `xml:"files-private-used"`
	FilesTotal               string `xml:"files-total"`
	FilesUsed                string `xml:"files-used"`
	InodefilePrivateCapacity string `xml:"inodefile-private-capacity"`
	InodefilePublicCapacity  string `xml:"inodefile-public-capacity"`
	InofileVersion           string `xml:"inofile-version"`
}

type VolumeLanguageAttributes struct {
	IsConvertUcodeEnabled string `xml:"is-convert-ucode-enabled"`
	IsCreateUcodeEnabled  string `xml:"is-create-ucode-enabled"`
	Language              string `xml:"language"`
	LanguageCode          string `xml:"language-code"`
	NfsCharacterSet       string `xml:"nfs-character-set"`
	OemCharacterSet       string `xml:"oem-character-set"`
}

type VolumeMirrorAttributes struct {
	IsDataProtectionMirror   string `xml:"is-data-protection-mirror"`
	IsLoadSharingMirror      string `xml:"is-load-sharing-mirror"`
	IsMoveMirror             string `xml:"is-move-mirror"`
	IsReplicaVolume          string `xml:"is-replica-volume"`
	MirrorTransferInProgress string `xml:"mirror-transfer-in-progress"`
	RedirectSnapshotID       string `xml:"redirect-snapshot-id"`
}

type VolumePerformanceAttributes struct {
	ExtentEnabled        string `xml:"extent-enabled"`
	FcDelegsEnabled      string `xml:"fc-delegs-enabled"`
	IsAtimeUpdateEnabled string `xml:"is-atime-update-enabled"`
	MaxWriteAllocBlocks  string `xml:"max-write-alloc-blocks"`
	MinimalReadAhead     string `xml:"minimal-read-ahead"`
	ReadRealloc          string `xml:"read-realloc"`
}

type VolumeQosAttributes struct {
	AdaptivePolicyGroupName string `xml:"adaptive-policy-group-name,omitempty"`
	PolicyGroupName         string `xml:"policy-group-name"`
}

type VolumeSecurityAttributes struct {
	Style                        string `xml:"style"`
	VolumeSecurityUnixAttributes struct {
		GroupID     string `xml:"group-id"`
		Permissions string `xml:"permissions"`
		UserID      string `xml:"user-id"`
	} `xml:"volume-security-unix-attributes"`
}

type VolumeSisAttributes struct {
	CompressionSpaceSaved             string `xml:"compression-space-saved"`
	DeduplicationSpaceSaved           string `xml:"deduplication-space-saved"`
	DeduplicationSpaceShared          string `xml:"deduplication-space-shared"`
	IsSisLoggingEnabled               string `xml:"is-sis-logging-enabled"`
	IsSisStateEnabled                 string `xml:"is-sis-state-enabled"`
	IsSisVolume                       string `xml:"is-sis-volume"`
	PercentageCompressionSpaceSaved   string `xml:"percentage-compression-space-saved"`
	PercentageDeduplicationSpaceSaved string `xml:"percentage-deduplication-space-saved"`
	PercentageTotalSpaceSaved         string `xml:"percentage-total-space-saved"`
	TotalSpaceSaved                   string `xml:"total-space-saved"`
}

type VolumeSnaplockAttributes struct {
	SnaplockType string `xml:"snaplock-type"`
}

type VolumeSnapshotAttributes struct {
	AutoSnapshotsEnabled           string `xml:"auto-snapshots-enabled,omitempty"`
	SnapdirAccessEnabled           bool   `xml:"snapdir-access-enabled,omitempty"`
	SnapshotCloneDependencyEnabled string `xml:"snapshot-clone-dependency-enabled,omitempty"`
	SnapshotCount                  string `xml:"snapshot-count,omitempty"`
	SnapshotPolicy                 string `xml:"snapshot-policy,omitempty"`
}

type VolumeSnapshotAutodeleteAttributes struct {
	Commitment          string `xml:"commitment"`
	DeferDelete         string `xml:"defer-delete"`
	DeleteOrder         string `xml:"delete-order"`
	DestroyList         string `xml:"destroy-list"`
	IsAutodeleteEnabled string `xml:"is-autodelete-enabled"`
	Prefix              string `xml:"prefix"`
	TargetFreeSpace     string `xml:"target-free-space"`
	Trigger             string `xml:"trigger"`
}

type VolumeSpaceAttributes struct {
	FilesystemSize                  string `xml:"filesystem-size,omitempty"`
	IsFilesysSizeFixed              string `xml:"is-filesys-size-fixed,omitempty"`
	IsSpaceGuaranteeEnabled         string `xml:"is-space-guarantee-enabled,omitempty"`
	IsSpaceSloEnabled               string `xml:"is-space-slo-enabled,omitempty"`
	OverwriteReserve                string `xml:"overwrite-reserve,omitempty"`
	OverwriteReserveRequired        string `xml:"overwrite-reserve-required,omitempty"`
	OverwriteReserveUsed            string `xml:"overwrite-reserve-used,omitempty"`
	OverwriteReserveUsedActual      string `xml:"overwrite-reserve-used-actual,omitempty"`
	PercentageFractionalReserve     string `xml:"percentage-fractional-reserve,omitempty"`
	PercentageSizeUsed              string `xml:"percentage-size-used,omitempty"`
	PercentageSnapshotReserve       string `xml:"percentage-snapshot-reserve,omitempty"`
	PercentageSnapshotReserveUsed   string `xml:"percentage-snapshot-reserve-used,omitempty"`
	PhysicalUsed                    string `xml:"physical-used,omitempty"`
	PhysicalUsedPercent             string `xml:"physical-used-percent,omitempty"`
	Size                            int    `xml:"size,omitempty"`
	SizeAvailable                   string `xml:"size-available,omitempty"`
	SizeAvailableForSnapshots       string `xml:"size-available-for-snapshots,omitempty"`
	SizeTotal                       string `xml:"size-total,omitempty"`
	SizeUsed                        string `xml:"size-used,omitempty"`
	SizeUsedBySnapshots             string `xml:"size-used-by-snapshots,omitempty"`
	SnapshotReserveSize             string `xml:"snapshot-reserve-size,omitempty"`
	SpaceFullThresholdPercent       string `xml:"space-full-threshold-percent,omitempty"`
	SpaceGuarantee                  string `xml:"space-guarantee,omitempty"`
	SpaceMgmtOptionTryFirst         string `xml:"space-mgmt-option-try-first,omitempty"`
	SpaceNearlyFullThresholdPercent string `xml:"space-nearly-full-threshold-percent,omitempty"`
	SpaceSlo                        string `xml:"space-slo,omitempty"`
}

type VolumeStateAttributes struct {
	BecomeNodeRootAfterReboot string `xml:"become-node-root-after-reboot"`
	ForceNvfailOnDr           string `xml:"force-nvfail-on-dr"`
	IgnoreInconsistent        string `xml:"ignore-inconsistent"`
	InNvfailedState           string `xml:"in-nvfailed-state"`
	IsClusterVolume           string `xml:"is-cluster-volume"`
	IsConstituent             string `xml:"is-constituent"`
	IsFlexgroup               string `xml:"is-flexgroup"`
	IsInconsistent            string `xml:"is-inconsistent"`
	IsInvalid                 string `xml:"is-invalid"`
	IsJunctionActive          string `xml:"is-junction-active"`
	IsMoving                  string `xml:"is-moving"`
	IsNodeRoot                string `xml:"is-node-root"`
	IsNvfailEnabled           string `xml:"is-nvfail-enabled"`
	IsQuiescedInMemory        string `xml:"is-quiesced-in-memory"`
	IsQuiescedOnDisk          string `xml:"is-quiesced-on-disk"`
	IsUnrecoverable           string `xml:"is-unrecoverable"`
	IsVolumeInCutover         string `xml:"is-volume-in-cutover"`
	IsVserverRoot             string `xml:"is-vserver-root"`
	State                     string `xml:"state"`
}

type VolumeTransitionAttributes struct {
	IsCftPrecommit        string `xml:"is-cft-precommit"`
	IsCopiedForTransition string `xml:"is-copied-for-transition"`
	IsTransitioned        string `xml:"is-transitioned"`
	TransitionBehavior    string `xml:"transition-behavior"`
}

type VolumeInfo struct {
	Encrypt                            string                              `xml:"encrypt,omitempty"`
	KeyID                              string                              `xml:"key-id,omitempty"`
	VolumeAntivirusAttributes          *VolumeAntivirusAttributes          `xml:"volume-antivirus-attributes,omitempty"`
	VolumeAutobalanceAttributes        *VolumeAutobalanceAttributes        `xml:"volume-autobalance-attributes,omitempty"`
	VolumeAutosizeAttributes           *VolumeAutosizeAttributes           `xml:"volume-autosize-attributes"`
	VolumeDirectoryAttributes          *VolumeDirectoryAttributes          `xml:"volume-directory-attributes"`
	VolumeExportAttributes             *VolumeExportAttributes             `xml:"volume-export-attributes,omitempty"`
	VolumeHybridCacheAttributes        *VolumeHybridCacheAttributes        `xml:"volume-hybrid-cache-attributes,omitempty"`
	VolumeIDAttributes                 *VolumeIDAttributes                 `xml:"volume-id-attributes,omitempty"`
	VolumeInodeAttributes              *VolumeInodeAttributes              `xml:"volume-inode-attributes,omitempty"`
	VolumeLanguageAttributes           *VolumeLanguageAttributes           `xml:"volume-language-attributes,omitempty"`
	VolumeMirrorAttributes             *VolumeMirrorAttributes             `xml:"volume-mirror-attributes,omitempty"`
	VolumePerformanceAttributes        *VolumePerformanceAttributes        `xml:"volume-performance-attributes,omitempty"`
	VolumeQosAttributes                *VolumeQosAttributes                `xml:"volume-qos-attributes,omitempty"`
	VolumeSecurityAttributes           *VolumeSecurityAttributes           `xml:"volume-security-attributes,omitempty"`
	VolumeSisAttributes                *VolumeSisAttributes                `xml:"volume-sis-attributes,omitempty"`
	VolumeSnaplockAttributes           *VolumeSnaplockAttributes           `xml:"volume-snaplock-attributes,omitempty"`
	VolumeSnapshotAttributes           *VolumeSnapshotAttributes           `xml:"volume-snapshot-attributes,omitempty"`
	VolumeSnapshotAutodeleteAttributes *VolumeSnapshotAutodeleteAttributes `xml:"volume-snapshot-autodelete-attributes,omitempty"`
	VolumeSpaceAttributes              *VolumeSpaceAttributes              `xml:"volume-space-attributes,omitempty"`
	VolumeStateAttributes              *VolumeStateAttributes              `xml:"volume-state-attributes,omitempty"`
	VolumeTransitionAttributes         *VolumeTransitionAttributes         `xml:"volume-transition-attributes,omitempty"`
}

type VolumeGetResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList []VolumeInfo `xml:"attributes-list>volume-attributes"`
		NextTag        string       `xml:"next-tag"`
		NumRecords     int          `xml:"num-records"`
	} `xml:"results"`
}

func (c *Client) VolumeGetAPI(options *VolumeGetOptions) (*VolumeGetResponse, *http.Response, error) {
	if c.VolumeGetIter == nil {
		c.VolumeGetIter = &VolumeGetIter {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.VolumeGetIter.Params.XMLName = xml.Name{Local: "volume-get-iter"}
	}
	c.VolumeGetIter.Base.Name = c.vserver
	c.VolumeGetIter.Params.VolumeGetOptions = options
	r := VolumeGetResponse{}
	res, err := c.VolumeGetIter.get(c.VolumeGetIter, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(VolumeGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) VolumeGetIterAPI(options *VolumeGetOptions) (responseVolumes []*VolumeGetResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.VolumeGetAPI(options)
		if err == nil {
			nextTag = r.Results.NextTag
			responseVolumes = append(responseVolumes, r)
			if nextTag == "" {
				break
			}
			options.Tag = nextTag
		} else {
			break
		}
	}
	return
}
