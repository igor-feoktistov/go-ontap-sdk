package ontap

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	libraryVersion = "1"
	ServerURL      = `/servlets/netapp.servlets.admin.XMLrequest_filer`
	userAgent      = "go-ontap-sdk/" + libraryVersion
	XMLNs          = "http://www.netapp.com/filer/admin"
)

type Base struct {
	XMLName xml.Name `xml:"netapp"`
	Version string   `xml:"version,attr"`
	XMLNs   string   `xml:"xmlsns,attr"`
	Name    string   `xml:"vfiler,attr,omitempty"`
	client  *Client
}

type Result interface {
	Passed() bool
	Result() *SingleResultBase
}

type ResultBase struct {
	Status     string `xml:"status,attr"`
	Reason     string `xml:"reason,attr"`
	NumRecords int    `xml:"num-records"`
	ErrorNo    int    `xml:"errno,attr"`
}

type SingleResultBase struct {
	Status  string `xml:"status,attr"`
	Reason  string `xml:"reason,attr"`
	ErrorNo int    `xml:"errno,attr"`
}

type SingleResultResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		SingleResultBase
	} `xml:"results"`
}

type AsyncResultBase struct {
	SingleResultBase
	ErrorCode    int    `xml:"result-error-code"`
	ErrorMessage string `xml:"result-error-message"`
	JobID        int    `xml:"result-jobid"`
	JobStatus    string `xml:"result-status"`
}

func (r *ResultBase) Passed() bool {
	return r.Status == "passed"
}

func (r *ResultBase) Result() *SingleResultBase {
	return &SingleResultBase{
		Status:  r.Status,
		Reason:  r.Reason,
		ErrorNo: r.ErrorNo,
	}
}

func (r *SingleResultBase) Passed() bool {
	return r.Status == "passed"
}

func (r *SingleResultBase) Result() *SingleResultBase {
	return r
}

func (r *AsyncResultBase) Passed() bool {
	return r.Status == "passed"
}

func (r *AsyncResultBase) Result() *SingleResultBase {
	return &SingleResultBase{
		Status:  r.Status,
		Reason:  r.Reason,
		ErrorNo: r.ErrorNo,
	}
}

func (b *Base) get(base interface{}, r interface{}) (*http.Response, error) {
	req, err := b.client.NewRequest("POST", &base)
	if err != nil {
		return nil, err
	}
	res, err := b.client.Do(req, r)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type Client struct {
	client          *http.Client
	BaseURL         *url.URL
	UserAgent       string
	options         *ClientOptions
	ResponseTimeout time.Duration
	vserver         string

	AggregateGetIter       *AggregateGetIter
	ClusterIdentityGet     *ClusterIdentityGet
	ExportPolicyGetIter    *ExportPolicyGetIter
	ExportPolicyCreate     *ExportPolicyCreate
	ExportPolicyDestroy    *ExportPolicyDestroy
	ExportRuleCreate       *ExportRuleCreate
	ExportRuleDestroy      *ExportRuleDestroy
	ExportRuleGetIter      *ExportRuleGetIter
	ExportRuleModify       *ExportRuleModify
	FileCreateDirectory    *FileCreateDirectory
	FileDeleteDirectory    *FileDeleteDirectory
	FileDeleteFile         *FileDeleteFile
	FileListDirectoryIter  *FileListDirectoryIter
	FileGetFileInfo        *FileGetFileInfo
	FileReadFile           *FileReadFile
	FileTruncateFile       *FileTruncateFile
	FileWriteFile          *FileWriteFile
	IgroupAdd              *IgroupAdd
	IgroupCreate           *IgroupCreate
	IgroupDestroy          *IgroupDestroy
	IgroupGetIter          *IgroupGetIter
	IgroupLookupLun        *IgroupLookupLun
	IgroupRemove           *IgroupRemove
	IscsiConnectionGetIter *IscsiConnectionGetIter
	IscsiInitiatorGetIter  *IscsiInitiatorGetIter
	IscsiInterfaceGetIter  *IscsiInterfaceGetIter
	IscsiNodeGetName       *IscsiNodeGetName
	IscsiServiceStatus     *IscsiServiceStatus
	LunCopyGetIter         *LunCopyGetIter
	LunCopyStart           *LunCopyStart
	LunCreateBySize        *LunCreateBySize
	LunCreateFromFile      *LunCreateFromFile
	LunDestroy             *LunDestroy
	LunInitiatorLoggedIn   *LunInitiatorLoggedIn
	LunGetAttribute        *LunGetAttribute
	LunGetAttributes       *LunGetAttributes
	LunGetIter             *LunGetIter
	LunMap                 *LunMap
	LunMapListInfo         *LunMapListInfo
	LunOffline             *LunOffline
	LunOnline              *LunOnline
	LunUnmap               *LunUnmap
	LunResize              *LunResize
	LunSetAttribute        *LunSetAttribute
	NetInterfaceGetIter    *NetInterfaceGetIter
	NetRoutesGetIter       *NetRoutesGetIter
	SnapshotCreate         *SnapshotCreate
	SnapshotDelete         *SnapshotDelete
	SnapshotGetIter        *SnapshotGetIter
	SnapshotListInfo       *SnapshotListInfo
	SnapshotRestoreVolume  *SnapshotRestoreVolume
	SystemNodeGetIter      *SystemNodeGetIter
	VolumeAutosizeSet      *VolumeAutosizeSet
	VolumeContainer        *VolumeContainer
	VolumeCreate           *VolumeCreate
	VolumeDestroy          *VolumeDestroy
	VolumeGetIter          *VolumeGetIter
	VolumeMount            *VolumeMount
	VolumeOffline          *VolumeOffline
	VolumeOnline           *VolumeOnline
	VolumeSetOption        *VolumeSetOption
	VolumeSize             *VolumeSize
	VolumeUnmount          *VolumeUnmount
	VserverGetIter         *VserverGetIter
	VserverShowAggrGetIter *VserverShowAggrGetIter
}

type ClientOptions struct {
	BasicAuthUser     string
	BasicAuthPassword string
	SSLVerify         bool
	Debug             bool
	Timeout           time.Duration
	Version           string
}

func DefaultOptions() *ClientOptions {
	return &ClientOptions{
		SSLVerify: true,
		Debug:     false,
		Timeout:   60 * time.Second,
		Version:   "1.15",
	}
}

func NewClient(endpoint string, options *ClientOptions) *Client {
	if options == nil {
		options = DefaultOptions()
	}
	httpClient := &http.Client{
		Timeout: options.Timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: !options.SSLVerify,
			},
		},
	}
	if !strings.HasSuffix(endpoint, "/") {
		endpoint = endpoint + "/"
	}
	baseURL, _ := url.Parse(endpoint)
	c := &Client{
		client:          httpClient,
		BaseURL:         baseURL,
		UserAgent:       userAgent,
		options:         options,
		ResponseTimeout: options.Timeout,
	}
	return c
}

func (c *Client) NewRequest(method string, body interface{}) (*http.Request, error) {
	u, _ := c.BaseURL.Parse(ServerURL)
	buf, err := xml.MarshalIndent(body, "", "  ")
	if err != nil {
		return nil, err
	}
	if c.options.Debug {
		log.Printf("[DEBUG] request xml: \n%v\n", string(buf))
	}
	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "text/xml")
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if c.options.BasicAuthUser != "" && c.options.BasicAuthPassword != "" {
		req.SetBasicAuth(c.options.BasicAuthUser, c.options.BasicAuthPassword)
	}
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (resp *http.Response, err error) {
	var bs []byte
	ctx, cncl := context.WithTimeout(context.Background(), c.ResponseTimeout)
	defer cncl()
	if resp, err = checkResp(c.client.Do(req.WithContext(ctx))); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			resp, err = checkResp(c.client.Do(req.WithContext(ctx)))
		}
	}
	if err != nil {
		return
	}
	if bs, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bs))
	if c.options.Debug {
		log.Printf("[DEBUG] response xml \n%v\n", string(bs))
	}
	if v != nil {
		defer resp.Body.Close()
		err = xml.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			return
		}
	}
	return
}

func (c *Client) SetVserver(vserver string) {
	c.vserver = vserver
}

func checkResp(resp *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return resp, err
	}
	switch resp.StatusCode {
	case 200, 201, 202, 204, 205, 206:
		return resp, nil
	default:
		return resp, newHTTPError(resp)
	}
}

func newHTTPError(resp *http.Response) error {
	return fmt.Errorf("Http Error status %d, Message: %s", resp.StatusCode, resp.Body)
}
