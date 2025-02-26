package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

const (
	xml01 string = `<?xml version="1.0" encoding="windows-1251"?>
<reg:register xmlns:reg="http://rsoc.ru" xmlns:tns="http://rsoc.ru" updateTime="2011-01-01T01:01:01+03:00" updateTimeUrgently="2010-02-02T02:02:01+03:00" formatVersion="2.4">

<content id="111" includeTime="2001-01-01T01:01:01" entryType="1" blockType="default" hash="XXXX">
        <decision date="2000-01-01" number="1/1/11-1111" org="ONE"/>
        <url><![CDATA[https://www.e01.tld/sex]]></url>
        <url><![CDATA[http://www.e01.tld/cheese]]></url>
        <url><![CDATA[http://www.e01.tld/slip]]></url>
        <domain><![CDATA[www.e01.tld]]></domain>
        <ip>192.168.1.11</ip>
        <ip>192.168.0.100</ip>
        <ip>10.1.1.1</ip>
        <ipv6>fd11:1::1</ipv6>
        <ipv6>fd11:11::1</ipv6>
        <ipv6>fdaa:f::100</ipv6>
</content>
<content id="222" includeTime="2001-01-01T02:02:02" entryType="1" blockType="domain" hash="YYYY">
        <decision date="2000-01-02" number="2/2/22-2222" org="TWO"/>
        <domain><![CDATA[www.e02.tld]]></domain>
        <ip>192.168.2.22</ip>
        <ip>192.168.0.100</ip>
        <ip>10.2.2.2</ip>
        <ipv6>fd22:2::2</ipv6>
        <ipv6>fd22:22::2</ipv6>
        <ipv6>fdaa:f::100</ipv6>
</content>
<content id="333" includeTime="2001-01-01T03:03:03" entryType="1" blockType="ip" hash="ZZZZ">
        <decision date="2001-01-03" number="3/3/33-3333" org="THREE"/>
        <ip>192.168.3.33</ip>
        <ip>192.168.0.100</ip>
        <ip>10.3.3.3</ip>
        <ipv6>fd33:3::3</ipv6>
        <ipv6>fd33:33::3</ipv6>
        <ipv6>fdaa:f::100</ipv6>
</content>
<content id="444" includeTime="2001-01-01T04:04:04" entryType="1" blockType="ip" hash="QQQQ">
        <decision date="2001-01-04" number="4/4/44-4444" org="FOUR"/>
        <ip>192.168.4.44</ip>
        <ip>192.168.4.100</ip>
        <ip>10.4.4.4</ip>
        <ipSubnet>10.4.0.0/16</ipSubnet>
        <ipv6>fd44:4::1</ipv6>
        <ipv6>fd44:44::1</ipv6>
        <ipv6>fdaa:f::100</ipv6>
</content>
<content id="555" includeTime="2001-01-01T05:05:05" entryType="1" blockType="domain" hash="PPPP">
        <decision date="2001-01-05" number="5/5/55-5555" org="FIVE"/>
        <domain><![CDATA[www.e02.tld]]></domain>
        <ip>192.168.5.55</ip>
        <ip>192.168.0.111</ip>
        <ip>10.5.5.5</ip>
        <ipv6>fd55:5::5</ipv6>
        <ipv6>fd55:55::5</ipv6>
        <ipv6>fdaa:f::100</ipv6>
</content>
</reg:register>`

	xml02 string = `<?xml version="1.0" encoding="windows-1251"?>
<reg:register xmlns:reg="http://rsoc.ru" xmlns:tns="http://rsoc.ru" updateTime="2013-03-03T03:03:03+03:00" updateTimeUrgently="2012-04-04T04:04:04+03:00" formatVersion="2.4">

<content id="111" includeTime="2009-10-11T23:00:00" entryType="1" blockType="default" hash="XXXX">
        <decision date="2001-02-17" number="1/1/11-1111" org="FSKN"/>
        <url><![CDATA[https://www.example01.com/sex]]></url>
        <url><![CDATA[http://www.example01.com/cheese]]></url>
        <domain><![CDATA[www.example01.com]]></domain>
        <ip>192.168.1.14</ip>
        <ip>192.168.12.100</ip>
        <ip>10.1.1.2</ip>
        <ipv6>fd11:beaf:7ea::1</ipv6>
        <ipv6>fd11:c01d:7ea::1</ipv6>
        <ipv6>fd12:c01d:7ea::100</ipv6>
</content>
<content id="222" includeTime="2009-10-11T12:00:00" entryType="1" blockType="domain" hash="YYYY">
        <decision date="2001-03-18" number="2/2/22-2222" org="RKN"/>
        <domain><![CDATA[www.example02.com]]></domain>
        <ip>192.168.2.11</ip>
        <ip>10.2.2.2</ip>
        <ipv6>fd12:beaf:7ea::1</ipv6>
        <ipv6>fd12:c01d:7ea::1</ipv6>
        <ipv6>fd12:c01d:7ea::100</ipv6>
</content>
<content id="333" includeTime="2009-12-11T06:00:00" entryType="1" blockType="ip" hash="ZZZZ">
        <decision date="2011-04-11" number="3/3/33-3333" org="MVD"/>
        <ip>192.168.3.11</ip>
        <ip>192.168.12.100</ip>
        <ip>10.0.3.2</ip>
        <ipv6>fd13:beaf:7ea::1</ipv6>
        <ipv6>fd13:c01d:7ea::1</ipv6>
        <ipv6>fd12:c01d:7ea::100</ipv6>
</content>
<content id="444" includeTime="2013-12-14T16:00:00" entryType="1" blockType="ip" hash="QQQQ">
        <decision date="2012-05-21" number="4/4/44-4444" org="MVD"/>
        <ip>192.168.4.11</ip>
        <ip>192.168.4.100</ip>
        <ip>10.0.4.2</ip>
        <ipSubnet>10.4.0.0/16</ipSubnet>
        <ipv6>fd14:beaf:7ea::1</ipv6>
        <ipv6>fd14:c01d:7ea::1</ipv6>
</content>
<content id="555" includeTime="2008-10-11T12:00:00" entryType="1" blockType="domain" hash="PPPP">
        <decision date="2002-03-18" number="2/2/22-2222" org="FSB"/>
        <domain><![CDATA[www.example02.com]]></domain>
        <ip>192.168.2.11</ip>
        <ip>192.168.12.111</ip>
        <ip>10.0.2.2</ip>
        <ipv6>fd12:beaf:7ea::2</ipv6>
        <ipv6>fd12:c01d:7ea::1</ipv6>
        <ipv6>fd12:c01d:7ea::100</ipv6>
</content>
</reg:register>`
)

func Test_Parse(t *testing.T) {
	logInit(os.Stderr, os.Stdout, os.Stderr, os.Stderr)
	dumpFile := strings.NewReader(xml01)
	err := Parse(dumpFile)
	if err != nil {
		t.Errorf(err.Error())
	}
	if Stats.MaxArrayIntSet != 5 ||
		Stats.Cnt != 5 ||
		Stats.CntAdd != 5 ||
		Stats.CntUpdate != 0 ||
		Stats.CntRemove != 0 {
		t.Errorf("Stat error")
	}
	if len(DumpSnap.ip) != 13 ||
		len(DumpSnap.ip6) != 11 ||
		len(DumpSnap.subnet) != 1 ||
		len(DumpSnap.subnet6) != 0 ||
		len(DumpSnap.url) != 3 ||
		len(DumpSnap.domain) != 2 {
		t.Errorf("Count error")
	}
	if len(DumpSnap.Content) != 5 ||
		len(DumpSnap.Content) != Stats.Cnt {
		t.Errorf("DumpSnap integrity error")
	}

	fmt.Println()
	dumpFile = strings.NewReader(xml02)
	err = Parse(dumpFile)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Printf("IP4:\n%v\n", DumpSnap.ip)
	for k := range DumpSnap.Content {
		fmt.Printf("%d ", k)
	}
	fmt.Println()
}
