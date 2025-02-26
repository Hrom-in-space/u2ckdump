package main

import (
	"net"
	"sync"

	"github.com/yl2chen/cidranger"
)

type (
	Nothing       struct{}
	Int32Map      map[int32]Nothing
	MinContentMap map[int32]*MinContent
)

var NothingV = Nothing{}

type Stat struct {
	Cnt            int
	CntAdd         int
	CntUpdate      int
	CntRemove      int
	MaxArrayIntSet int
	MaxContentSize int
}

var Stats Stat

type TDump struct {
	sync.RWMutex
	utime    int64
	ip       IP4Set
	ip6      StringIntSet
	subnet   StringIntSet
	subnet6  StringIntSet
	net      cidranger.Ranger
	url      StringIntSet
	domain   StringIntSet
	decision DecisionSet
	Content  MinContentMap
}

func NewTDump() *TDump {
	return &TDump{
		utime:    0,
		ip:       make(IP4Set),
		ip6:      make(StringIntSet),
		subnet:   make(StringIntSet),
		subnet6:  make(StringIntSet),
		url:      make(StringIntSet),
		domain:   make(StringIntSet),
		decision: make(DecisionSet),
		Content:  make(MinContentMap),
		net:      cidranger.NewPCTrieRanger(),
	}
}

func (t *TDump) AddIp(ip uint32, id int32) {
	t.ip.Add(ip, id)
}

func (t *TDump) DeleteIp(ip uint32, id int32) {
	t.ip.Delete(ip, id)
}

func (t *TDump) AddIp6(i string, id int32) {
	t.ip6.Add(i, id)
}
func (t *TDump) DeleteIp6(i string, id int32) {
	t.ip6.Delete(i, id)
}

func (t *TDump) AddSubnet(i string, id int32) {
	if t.subnet.Add(i, id) {
		_, network, err := net.ParseCIDR(i)
		if err != nil {
			Debug.Printf("Can't parse CIDR: %s: %s\n", i, err.Error())
		}
		err = t.net.Insert(cidranger.NewBasicRangerEntry(*network))
		if err != nil {
			Debug.Printf("Can't insert CIDR: %s: %s\n", i, err.Error())
		}
	}
}
func (t *TDump) DeleteSubnet(i string, id int32) {
	if t.subnet.Delete(i, id) {
		_, network, err := net.ParseCIDR(i)
		if err != nil {
			Debug.Printf("Can't parse CIDR: %s: %s\n", i, err.Error())
		}
		_, err = t.net.Remove(*network)
		if err != nil {
			Debug.Printf("Can't remove CIDR: %s: %s\n", i, err.Error())
		}
	}
}

func (t *TDump) AddSubnet6(i string, id int32) {
	if t.subnet6.Add(i, id) {
		_, network, err := net.ParseCIDR(i)
		if err != nil {
			Debug.Printf("Can't parse CIDR: %s: %s\n", i, err.Error())
		}
		err = t.net.Insert(cidranger.NewBasicRangerEntry(*network))
		if err != nil {
			Debug.Printf("Can't insert CIDR: %s: %s\n", i, err.Error())
		}
	}
}
func (t *TDump) DeleteSubnet6(i string, id int32) {
	if t.subnet6.Delete(i, id) {
		_, network, err := net.ParseCIDR(i)
		if err != nil {
			Debug.Printf("Can't parse CIDR: %s: %s\n", i, err.Error())
		}
		_, err = t.net.Remove(*network)
		if err != nil {
			Debug.Printf("Can't remove CIDR: %s: %s\n", i, err.Error())
		}
	}
}

func (t *TDump) AddUrl(i string, id int32) {
	t.url.Add(i, id)
}
func (t *TDump) DeleteUrl(i string, id int32) {
	t.url.Delete(i, id)
}

func (t *TDump) AddDomain(i string, id int32) {
	t.domain.Add(i, id)
}
func (t *TDump) DeleteDomain(i string, id int32) {
	t.domain.Delete(i, id)
}

func (t *TDump) AddDecision(i uint64, id int32) {
	t.decision.Add(i, id)
}
func (t *TDump) DeleteDecision(i uint64, id int32) {
	t.decision.Delete(i, id)
}

var DumpSnap = NewTDump()

type Reg struct {
	UpdateTime         int64
	UpdateTimeUrgently string
	FormatVersion      string
}

func Parse2(UpdateTime int64) {
	DumpSnap.Lock()
	for _, v := range DumpSnap.Content {
		v.RegistryUpdateTime = UpdateTime
	}
	DumpSnap.utime = UpdateTime
	DumpSnap.Unlock()
}
