package host

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/imdario/mergo"
	"net/http"
	"strconv"
	"time"
)

// 定义列表返回的字段
type HostSet struct {
	Total int     `json:"total"`
	Items []*Host `json:"items"`
}

func NewHostSet() *HostSet {
	return &HostSet{
		Items: []*Host{},
	}
}
func NewHost() *Host {
	return &Host{
		Resource: &Resource{},
		Describe: &Describe{},
	}
}

func (s *HostSet) Add(item *Host) {
	s.Items = append(s.Items, item)
}

// host模型的定义
type Host struct {
	// 资源公共属性部分
	*Resource
	// 资源独有属性部分
	*Describe
}

var (
	validate = validator.New()
)

func (h *Host) Validate() error {
	return validate.Struct(h)
}

func (h *Host) InjectDefault() {
	if h.CreateAt == 0 {
		h.CreateAt = time.Now().UnixMilli()
	}
}

type Vendor int

const (
	PRIVATE_IDC Vendor = iota
	ALIYUN
	TXYUN
)

type Resource struct {
	Id          string            `json:"id"  validate:"required"`     // 全局唯一Id
	Vendor      Vendor            `json:"vendor"`                      // 厂商
	Region      string            `json:"region"  validate:"required"` // 地域
	CreateAt    int64             `json:"create_at"`                   // 创建时间
	ExpireAt    int64             `json:"expire_at"`                   // 过期时间
	Type        string            `json:"type"  validate:"required"`   // 规格
	Name        string            `json:"name"  validate:"required"`   // 名称
	Description string            `json:"description"`                 // 描述
	Status      string            `json:"status"`                      // 服务商中的状态
	Tags        map[string]string `json:"tags"`                        // 标签
	UpdateAt    int64             `json:"update_at"`                   // 更新时间
	SyncAt      int64             `json:"sync_at"`                     // 同步时间
	Account     string            `json:"accout"`                      // 资源的所属账号
	PublicIP    string            `json:"public_ip"`                   // 公网IP
	PrivateIP   string            `json:"private_ip"`                  // 内网IP
}

type Describe struct {
	ResourceId   string `json:"id"  validate:"required"`    // 全局唯一Id
	CPU          int    `json:"cpu" validate:"required"`    // 核数
	Memory       int    `json:"memory" validate:"required"` // 内存
	GPUAmount    int    `json:"gpu_amount"`                 // GPU数量
	GPUSpec      string `json:"gpu_spec"`                   // GPU类型
	OSType       string `json:"os_type"`                    // 操作系统类型，分为Windows和Linux
	OSName       string `json:"os_name"`                    // 操作系统名称
	SerialNumber string `json:"serial_number"`              // 序列号
}

func NewQueryHostRequest() *QueryHostRequest {
	return &QueryHostRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}
func NewDescribeHostRequestWithId(id string) *DescribeHostRequest {
	return &DescribeHostRequest{
		Id: id,
	}
}

func NewPutUpdateHostRequest(id string) *UpdateHostRequest {
	h := NewHost()
	h.Id = id
	return &UpdateHostRequest{
		UpdateMode: UPDATE_PUT_PUT,
		Host:       h,
	}
}
func NewPatchUpdateHostRequest(id string) *UpdateHostRequest {
	h := NewHost()
	h.Id = id
	return &UpdateHostRequest{
		UpdateMode: UPDATE_PUT_PATCH,
		Host:       h,
	}
}

type QueryHostRequest struct {
	PageSize   int    `json:"page_size"`
	PageNumber int    `json:"page_number"`
	Keywords   string `json:"kws"`
}

func (req *QueryHostRequest) GetPageNumber() uint {
	return uint(req.PageNumber)
}

func (req *QueryHostRequest) Offset() int64 {
	return int64((req.PageNumber - 1) * req.PageSize)
}

// 这里新定义UPDATE_MODE string 是需要限制必须传put或者patch这两个值
// 如果用string的话，可以任意传值
type UPDATE_MODE string

const (
	// 全量更新
	UPDATE_PUT_PUT UPDATE_MODE = "PUT"
	// 局部更新
	UPDATE_PUT_PATCH UPDATE_MODE = "PATCH"
)

type UpdateHostRequest struct {
	UpdateMode UPDATE_MODE `json:"update_mode"`
	*Host
}

type DeleteHostRequest struct {
	Id string
}

func NewQueryHostFromHTTP(r *http.Request) *QueryHostRequest {
	req := NewQueryHostRequest()
	// query string
	qs := r.URL.Query()
	pss := qs.Get("page_size")
	if pss != "" {
		req.PageSize, _ = strconv.Atoi(pss)
	}

	pns := qs.Get("page_number")
	if pns != "" {
		req.PageNumber, _ = strconv.Atoi(pns)
	}

	req.Keywords = qs.Get("kws")
	return req
}

type DescribeHostRequest struct {
	Id string
}

// 全量更新
func (h *Host) Put(obj *Host) error {
	if obj.Id != h.Id {
		return fmt.Errorf("id not equal")
	}
	*h.Resource = *obj.Resource
	// 这是整个对象的拷贝，更新的时候只更新值，不更新对象，所以不建议这么写
	//h.Describe = obj
	// 这是值的拷贝
	*h.Describe = *obj.Describe

	return nil
}

// 部分更新
func (h *Host) Patch(obj *Host) error {
	//if obj.Name != "" {
	//	h.Name = obj.Name
	//}
	//if obj.CPU != 0 {
	//	h.CPU = obj.CPU
	//}
	//
	//return nil

	return mergo.MergeWithOverwrite(h, obj)
}
