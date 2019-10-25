/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 13:37 2019-10-25
 */
package main

import (
	"log"
	"net"
	"os"

	"fmt"
	"github.com/miekg/dns"
)

func main() {
	// 读取当前运行环境的 /etc/resolv.conf，获得 name server 的配置
	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")

	// 构造发起 DNS 请求的客户端
	c := new(dns.Client)

	// 构造 DNS 报文
	m := new(dns.Msg)

	// 设置问题字段，即查询命令行参数第一个参数的 A 记录
	m.SetQuestion(dns.Fqdn(os.Args[1]), dns.TypeA)
	m.RecursionDesired = true

	// client 发起 DNS 请求，其中 c 为上文创建的 client，m 为构造的 DNS 报文
	// config 为从 /etc/resolv.conf 构造出来的配置
	r, _, err := c.Exchange(m, net.JoinHostPort(config.Servers[0], config.Port))
	if r == nil {
		log.Fatalf("*** error: %s\n", err.Error())
	}

	if r.Rcode != dns.RcodeSuccess {
		log.Fatalf("*** invalid answer name %s after MX query for %s\n", os.Args[1], os.Args[1])
	}

	// 如果 DNS 查询成功
	for _, a := range r.Answer {
		fmt.Printf("%v\n", a)
	}
}
