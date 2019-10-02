package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type FileGetFileInfo struct {
	Base
	Params struct {
		XMLName xml.Name
		Path string `xml:"path"`
	}
}

type FileGetFileInfoResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		FileInfo FileInfo `xml:"file-info"`
	} `xml:"results"`
}

type FileInfo struct {
	AccessedTimestamp int    `xml:"accessed-timestamp"`
	AclType           string `xml:"acl-type"`
	BytesUsed         int    `xml:"bytes-used,omitempty"`
	ChangedTimestamp  int    `xml:"changed-timestamp"`
	CreationTimestamp int    `xml:"creation-timestamp"`
	Dsid              int    `xml:"dsid,omitempty"`
	FileSize          int    `xml:"file-size"`
	FileType          string `xml:"file-type"`
	GroupId           int    `xml:"group-id"`
	HardLinksCount    int    `xml:"hard-links-count"`
	InodeGenNumber    int    `xml:"inode-gen-number,omitempty"`
	InodeNumber       int    `xml:"inode-number"`
	IsEmpty           bool   `xml:"is-empty,omitempty"`
	IsJunction        bool   `xml:"is-junction,omitempty"`
	IsVmAligned       bool   `xml:"is-vm-aligned,omitempty"`
	ModifiedTimestamp int    `xml:"modified-timestamp"`
	Msid              int    `xml:"msid,omitempty"`
	Name              string `xml:"name,omitempty"`
	OwnerId           int    `xml:"owner-id"`
	Perm              string `xml:"perm"`
}

func (c *Client) FileGetFileInfoAPI(path string) (*FileGetFileInfoResponse, *http.Response, error) {
	if c.FileGetFileInfo == nil {
		c.FileGetFileInfo = &FileGetFileInfo {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.FileGetFileInfo.Params.XMLName = xml.Name{Local: "file-get-file-info"}
	}
	c.FileGetFileInfo.Base.Name = c.vserver
	c.FileGetFileInfo.Params.Path = path
	r := FileGetFileInfoResponse{}
	res, err := c.FileGetFileInfo.get(c.FileGetFileInfo, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(FileGetFileInfoAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
