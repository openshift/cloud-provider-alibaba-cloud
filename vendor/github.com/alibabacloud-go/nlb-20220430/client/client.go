// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddServersToServerGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** of each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck the request. Valid values:
	//
	// *   **true**: prechecks the request but does not add the servers to the server group. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the server group.
	ServerGroupId *string                                  `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	Servers       []*AddServersToServerGroupRequestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
}

func (s AddServersToServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupRequest) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupRequest) SetClientToken(v string) *AddServersToServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetDryRun(v bool) *AddServersToServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetRegionId(v string) *AddServersToServerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetServerGroupId(v string) *AddServersToServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetServers(v []*AddServersToServerGroupRequestServers) *AddServersToServerGroupRequest {
	s.Servers = v
	return s
}

type AddServersToServerGroupRequestServers struct {
	// The description of the servers.
	//
	// The description must be 2 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (\_), and hyphens (-).
	//
	// >  You can specify at most 40 servers in each call.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port used by the backend server. Valid values: **1** to **65535**.
	//
	// >  You can specify at most 40 servers in each call.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the server. You can specify at most 40 server IDs in each call.
	//
	// *   If the server group type is **Instance**, set the ServerId parameter to the ID of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance. These backend servers are specified by **Ecs**, **Eni**, or **Eci**.
	// *   If the server group type is **Ip**, set the ServerId parameter to an IP address.
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the server. If the server group type is **Ip**, set the ServerId parameter to an IP address.
	//
	// >  You can specify at most 40 server IP addresses in each call.
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of the backend server. Valid values:
	//
	// *   **Ecs**: an ECS instance
	// *   **Eni**: an ENI
	// *   **Eci**: an elastic container instance
	// *   **Ip**: an IP address
	//
	// >  You can specify at most 40 servers in each call.
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The weight of the backend server. Valid values: **0** to **100**. Default value: **100**. If the weight of a backend server is set to **0**, no requests are forwarded to the backend server.
	//
	// >  You can specify at most 40 servers in each call.
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s AddServersToServerGroupRequestServers) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupRequestServers) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupRequestServers) SetDescription(v string) *AddServersToServerGroupRequestServers {
	s.Description = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetPort(v int32) *AddServersToServerGroupRequestServers {
	s.Port = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetServerId(v string) *AddServersToServerGroupRequestServers {
	s.ServerId = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetServerIp(v string) *AddServersToServerGroupRequestServers {
	s.ServerIp = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetServerType(v string) *AddServersToServerGroupRequestServers {
	s.ServerType = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetWeight(v int32) *AddServersToServerGroupRequestServers {
	s.Weight = &v
	return s
}

type AddServersToServerGroupResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s AddServersToServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupResponseBody) SetJobId(v string) *AddServersToServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *AddServersToServerGroupResponseBody) SetRequestId(v string) *AddServersToServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddServersToServerGroupResponseBody) SetServerGroupId(v string) *AddServersToServerGroupResponseBody {
	s.ServerGroupId = &v
	return s
}

type AddServersToServerGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddServersToServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddServersToServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupResponse) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupResponse) SetHeaders(v map[string]*string) *AddServersToServerGroupResponse {
	s.Headers = v
	return s
}

func (s *AddServersToServerGroupResponse) SetStatusCode(v int32) *AddServersToServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddServersToServerGroupResponse) SetBody(v *AddServersToServerGroupResponseBody) *AddServersToServerGroupResponse {
	s.Body = v
	return s
}

type AssociateAdditionalCertificatesWithListenerRequest struct {
	AdditionalCertificateIds []*string `json:"AdditionalCertificateIds,omitempty" xml:"AdditionalCertificateIds,omitempty" type:"Repeated"`
	ClientToken              *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ListenerId               *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId                 *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AssociateAdditionalCertificatesWithListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerRequest) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetAdditionalCertificateIds(v []*string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.AdditionalCertificateIds = v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetClientToken(v string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetDryRun(v bool) *AssociateAdditionalCertificatesWithListenerRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetListenerId(v string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetRegionId(v string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.RegionId = &v
	return s
}

type AssociateAdditionalCertificatesWithListenerResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateAdditionalCertificatesWithListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerResponseBody) SetJobId(v string) *AssociateAdditionalCertificatesWithListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerResponseBody) SetRequestId(v string) *AssociateAdditionalCertificatesWithListenerResponseBody {
	s.RequestId = &v
	return s
}

type AssociateAdditionalCertificatesWithListenerResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssociateAdditionalCertificatesWithListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateAdditionalCertificatesWithListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerResponse) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerResponse) SetHeaders(v map[string]*string) *AssociateAdditionalCertificatesWithListenerResponse {
	s.Headers = v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerResponse) SetStatusCode(v int32) *AssociateAdditionalCertificatesWithListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerResponse) SetBody(v *AssociateAdditionalCertificatesWithListenerResponseBody) *AssociateAdditionalCertificatesWithListenerResponse {
	s.Body = v
	return s
}

type AttachCommonBandwidthPackageToLoadBalancerRequest struct {
	// The ID of the EIP bandwidth plan.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** of each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck the request. Valid values:
	//
	// *   **true**: prechecks the request but does not associate the EIP bandwidth plan with the NLB instance. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the NLB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachCommonBandwidthPackageToLoadBalancerRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachCommonBandwidthPackageToLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) SetBandwidthPackageId(v string) *AttachCommonBandwidthPackageToLoadBalancerRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) SetClientToken(v string) *AttachCommonBandwidthPackageToLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) SetDryRun(v bool) *AttachCommonBandwidthPackageToLoadBalancerRequest {
	s.DryRun = &v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) SetLoadBalancerId(v string) *AttachCommonBandwidthPackageToLoadBalancerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) SetRegionId(v string) *AttachCommonBandwidthPackageToLoadBalancerRequest {
	s.RegionId = &v
	return s
}

type AttachCommonBandwidthPackageToLoadBalancerResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachCommonBandwidthPackageToLoadBalancerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachCommonBandwidthPackageToLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *AttachCommonBandwidthPackageToLoadBalancerResponseBody) SetJobId(v string) *AttachCommonBandwidthPackageToLoadBalancerResponseBody {
	s.JobId = &v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerResponseBody) SetRequestId(v string) *AttachCommonBandwidthPackageToLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

type AttachCommonBandwidthPackageToLoadBalancerResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachCommonBandwidthPackageToLoadBalancerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachCommonBandwidthPackageToLoadBalancerResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachCommonBandwidthPackageToLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s *AttachCommonBandwidthPackageToLoadBalancerResponse) SetHeaders(v map[string]*string) *AttachCommonBandwidthPackageToLoadBalancerResponse {
	s.Headers = v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerResponse) SetStatusCode(v int32) *AttachCommonBandwidthPackageToLoadBalancerResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerResponse) SetBody(v *AttachCommonBandwidthPackageToLoadBalancerResponseBody) *AttachCommonBandwidthPackageToLoadBalancerResponse {
	s.Body = v
	return s
}

type CreateListenerRequest struct {
	// Specifies whether to enable Application-Layer Protocol Negotiation (ALPN). Valid values:
	//
	// *   **true**: yes
	// *   **false** (default): no
	AlpnEnabled *bool `json:"AlpnEnabled,omitempty" xml:"AlpnEnabled,omitempty"`
	// The ALPN policy.
	AlpnPolicy       *string   `json:"AlpnPolicy,omitempty" xml:"AlpnPolicy,omitempty"`
	CaCertificateIds []*string `json:"CaCertificateIds,omitempty" xml:"CaCertificateIds,omitempty" type:"Repeated"`
	// Specifies whether to enable mutual authentication. Valid values:
	//
	// *   **true**: yes
	// *   **false** (default): no
	CaEnabled      *bool     `json:"CaEnabled,omitempty" xml:"CaEnabled,omitempty"`
	CertificateIds []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId** as **ClientToken**. **RequestId** may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The maximum number of connections that can be created per second on the NLB instance. Valid values: **0** to **1000000**. **0** specifies that the number of connections is unlimited.
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// Specifies whether to only precheck this request. Valid values:
	//
	// *   **true**: prechecks the request without creating the resource. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The last port in the listening port range. Valid values: **0** to **65535**.
	//
	// The number of the last port must be larger than that of the first port.
	EndPort *int32 `json:"EndPort,omitempty" xml:"EndPort,omitempty"`
	// The timeout period of an idle connection. Unit: seconds.
	//
	// Valid values: **1** to **900**. Default value: **900**.
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The name of the listener.
	//
	// The name must be 2 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (\_), and hyphens (-).
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The listening port. Valid values: **0** to **65535**.
	//
	// If you set the value to **0**, the listener listens by port range. If you set the value to **0**, you must also set the **StartPort** and **EndPort** parameters.
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The listening protocol. Valid values: **TCP**, **UDP**, and **TCPSSL**.
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The ID of the NLB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The maximum size of a TCP segment. Unit: bytes. Valid values: **0** to **1500**.
	//
	// **0** specifies that the maximum segment size remains unchanged.
	//
	// >  This parameter is supported only by listeners that use SSL over TCP.
	Mss *int32 `json:"Mss,omitempty" xml:"Mss,omitempty"`
	// Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers. Valid values:
	//
	// *   **true**: yes
	// *   **false** (default): no
	ProxyProtocolEnabled *bool `json:"ProxyProtocolEnabled,omitempty" xml:"ProxyProtocolEnabled,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable fine-grained monitoring. Valid values:
	//
	// *   **true**: yes
	// *   **false** (default): no
	SecSensorEnabled *bool `json:"SecSensorEnabled,omitempty" xml:"SecSensorEnabled,omitempty"`
	// The ID of the security policy. System security policies and custom security policies are supported.
	//
	// Valid values: **tls_cipher_policy\_1\_0** (default), **tls_cipher_policy\_1\_1**, **tls_cipher_policy\_1\_2**, **tls_cipher_policy\_1\_2\_strict**, and **tls_cipher_policy\_1\_2\_strict_with\_1\_3**.
	//
	// >  This parameter takes effect only for listeners that use SSL over TCP.
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The first port in the listening port range. Valid values: **0** to **65535**.
	StartPort *int32                      `json:"StartPort,omitempty" xml:"StartPort,omitempty"`
	Tag       []*CreateListenerRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateListenerRequest) SetAlpnEnabled(v bool) *CreateListenerRequest {
	s.AlpnEnabled = &v
	return s
}

func (s *CreateListenerRequest) SetAlpnPolicy(v string) *CreateListenerRequest {
	s.AlpnPolicy = &v
	return s
}

func (s *CreateListenerRequest) SetCaCertificateIds(v []*string) *CreateListenerRequest {
	s.CaCertificateIds = v
	return s
}

func (s *CreateListenerRequest) SetCaEnabled(v bool) *CreateListenerRequest {
	s.CaEnabled = &v
	return s
}

func (s *CreateListenerRequest) SetCertificateIds(v []*string) *CreateListenerRequest {
	s.CertificateIds = v
	return s
}

func (s *CreateListenerRequest) SetClientToken(v string) *CreateListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateListenerRequest) SetCps(v int32) *CreateListenerRequest {
	s.Cps = &v
	return s
}

func (s *CreateListenerRequest) SetDryRun(v bool) *CreateListenerRequest {
	s.DryRun = &v
	return s
}

func (s *CreateListenerRequest) SetEndPort(v int32) *CreateListenerRequest {
	s.EndPort = &v
	return s
}

func (s *CreateListenerRequest) SetIdleTimeout(v int32) *CreateListenerRequest {
	s.IdleTimeout = &v
	return s
}

func (s *CreateListenerRequest) SetListenerDescription(v string) *CreateListenerRequest {
	s.ListenerDescription = &v
	return s
}

func (s *CreateListenerRequest) SetListenerPort(v int32) *CreateListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateListenerRequest) SetListenerProtocol(v string) *CreateListenerRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *CreateListenerRequest) SetLoadBalancerId(v string) *CreateListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateListenerRequest) SetMss(v int32) *CreateListenerRequest {
	s.Mss = &v
	return s
}

func (s *CreateListenerRequest) SetProxyProtocolEnabled(v bool) *CreateListenerRequest {
	s.ProxyProtocolEnabled = &v
	return s
}

func (s *CreateListenerRequest) SetRegionId(v string) *CreateListenerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateListenerRequest) SetSecSensorEnabled(v bool) *CreateListenerRequest {
	s.SecSensorEnabled = &v
	return s
}

func (s *CreateListenerRequest) SetSecurityPolicyId(v string) *CreateListenerRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *CreateListenerRequest) SetServerGroupId(v string) *CreateListenerRequest {
	s.ServerGroupId = &v
	return s
}

func (s *CreateListenerRequest) SetStartPort(v int32) *CreateListenerRequest {
	s.StartPort = &v
	return s
}

func (s *CreateListenerRequest) SetTag(v []*CreateListenerRequestTag) *CreateListenerRequest {
	s.Tag = v
	return s
}

type CreateListenerRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateListenerRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestTag) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestTag) SetKey(v string) *CreateListenerRequestTag {
	s.Key = &v
	return s
}

func (s *CreateListenerRequestTag) SetValue(v string) *CreateListenerRequestTag {
	s.Value = &v
	return s
}

type CreateListenerResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the listener.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateListenerResponseBody) SetJobId(v string) *CreateListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateListenerResponseBody) SetListenerId(v string) *CreateListenerResponseBody {
	s.ListenerId = &v
	return s
}

func (s *CreateListenerResponseBody) SetRequestId(v string) *CreateListenerResponseBody {
	s.RequestId = &v
	return s
}

type CreateListenerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerResponse) GoString() string {
	return s.String()
}

func (s *CreateListenerResponse) SetHeaders(v map[string]*string) *CreateListenerResponse {
	s.Headers = v
	return s
}

func (s *CreateListenerResponse) SetStatusCode(v int32) *CreateListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateListenerResponse) SetBody(v *CreateListenerResponseBody) *CreateListenerResponse {
	s.Body = v
	return s
}

type CreateLoadBalancerRequest struct {
	// The protocol version. Valid values:
	//
	// *   **ipv4:** IPv4. This is the default value.
	// *   **DualStack:** dual stack.
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The type of IPv4 address used by the NLB instance. Valid values:
	//
	// *   **Internet**: The NLB instance uses a public IP address. The domain name of the NLB instance is resolved to the public IP address. Therefore, the NLB instance can be accessed over the Internet.
	// *   **Intranet**: The NLB instance uses a private IP address. The domain name of the NLB instance is resolved to the private IP address. Therefore, the NLB instance can be accessed over the virtual private cloud (VPC) where the NLB instance is deployed.
	//
	// >  To enable a public IPv6 address for an NLB instance, call the [EnableLoadBalancerIpv6Internet](~~445878~~) operation.
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The ID of the EIP bandwidth plan that is associated with the Internet-facing NLB instance.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** for each API request is different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configuration of the deletion protection feature.
	DeletionProtectionConfig *CreateLoadBalancerRequestDeletionProtectionConfig `json:"DeletionProtectionConfig,omitempty" xml:"DeletionProtectionConfig,omitempty" type:"Struct"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// *   **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false**: performs a dry run and sends the request. This is the default value. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The billing settings of the NLB instance.
	LoadBalancerBillingConfig *CreateLoadBalancerRequestLoadBalancerBillingConfig `json:"LoadBalancerBillingConfig,omitempty" xml:"LoadBalancerBillingConfig,omitempty" type:"Struct"`
	// The name of the NLB instance.
	//
	// The value must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The value must start with a letter.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The type of the instance. Set the value to **network**, which specifies an NLB instance.
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" xml:"LoadBalancerType,omitempty"`
	// The configuration of the configuration read-only mode.
	ModificationProtectionConfig *CreateLoadBalancerRequestModificationProtectionConfig `json:"ModificationProtectionConfig,omitempty" xml:"ModificationProtectionConfig,omitempty" type:"Struct"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*CreateLoadBalancerRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC where the NLB instance is deployed.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The mappings between zones and vSwitches. You must add at least two zones. You can add a maximum of 10 zones.
	ZoneMappings []*CreateLoadBalancerRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s CreateLoadBalancerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequest) SetAddressIpVersion(v string) *CreateLoadBalancerRequest {
	s.AddressIpVersion = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetAddressType(v string) *CreateLoadBalancerRequest {
	s.AddressType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetBandwidthPackageId(v string) *CreateLoadBalancerRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetClientToken(v string) *CreateLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetDeletionProtectionConfig(v *CreateLoadBalancerRequestDeletionProtectionConfig) *CreateLoadBalancerRequest {
	s.DeletionProtectionConfig = v
	return s
}

func (s *CreateLoadBalancerRequest) SetDryRun(v bool) *CreateLoadBalancerRequest {
	s.DryRun = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerBillingConfig(v *CreateLoadBalancerRequestLoadBalancerBillingConfig) *CreateLoadBalancerRequest {
	s.LoadBalancerBillingConfig = v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerName(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerType(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetModificationProtectionConfig(v *CreateLoadBalancerRequestModificationProtectionConfig) *CreateLoadBalancerRequest {
	s.ModificationProtectionConfig = v
	return s
}

func (s *CreateLoadBalancerRequest) SetRegionId(v string) *CreateLoadBalancerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetResourceGroupId(v string) *CreateLoadBalancerRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetTag(v []*CreateLoadBalancerRequestTag) *CreateLoadBalancerRequest {
	s.Tag = v
	return s
}

func (s *CreateLoadBalancerRequest) SetVpcId(v string) *CreateLoadBalancerRequest {
	s.VpcId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetZoneMappings(v []*CreateLoadBalancerRequestZoneMappings) *CreateLoadBalancerRequest {
	s.ZoneMappings = v
	return s
}

type CreateLoadBalancerRequestDeletionProtectionConfig struct {
	// Specifies whether to enable deletion protection. Valid values:
	//
	// *   **true**: yes
	// *   **false** (default): no
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The reason why the deletion protection feature is enabled or disabled. The value must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The value must start with a letter.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s CreateLoadBalancerRequestDeletionProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequestDeletionProtectionConfig) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestDeletionProtectionConfig) SetEnabled(v bool) *CreateLoadBalancerRequestDeletionProtectionConfig {
	s.Enabled = &v
	return s
}

func (s *CreateLoadBalancerRequestDeletionProtectionConfig) SetReason(v string) *CreateLoadBalancerRequestDeletionProtectionConfig {
	s.Reason = &v
	return s
}

type CreateLoadBalancerRequestLoadBalancerBillingConfig struct {
	// The billing method of the NLB instance.
	//
	// Set the value to **PostPay**, which specifies the pay-as-you-go billing method.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
}

func (s CreateLoadBalancerRequestLoadBalancerBillingConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequestLoadBalancerBillingConfig) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestLoadBalancerBillingConfig) SetPayType(v string) *CreateLoadBalancerRequestLoadBalancerBillingConfig {
	s.PayType = &v
	return s
}

type CreateLoadBalancerRequestModificationProtectionConfig struct {
	// The reason why the configuration read-only mode is enabled. The value must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The value must start with a letter.
	//
	// >  This parameter takes effect only if the **Status** parameter is set to **ConsoleProtection**.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// Specifies whether to enable the configuration read-only mode. Valid values:
	//
	// *   **NonProtection**: does not enable the configuration read-only mode. You cannot set the **Reason** parameter. If the **Reason** parameter is set, the value is cleared.
	// *   **ConsoleProtection**: enables the configuration read-only mode. You can set the **Reason** parameter.
	//
	// >  If you set this parameter to **ConsoleProtection**, you cannot use the NLB console to modify instance configurations. However, you can call API operations to modify instance configurations.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateLoadBalancerRequestModificationProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequestModificationProtectionConfig) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestModificationProtectionConfig) SetReason(v string) *CreateLoadBalancerRequestModificationProtectionConfig {
	s.Reason = &v
	return s
}

func (s *CreateLoadBalancerRequestModificationProtectionConfig) SetStatus(v string) *CreateLoadBalancerRequestModificationProtectionConfig {
	s.Status = &v
	return s
}

type CreateLoadBalancerRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateLoadBalancerRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequestTag) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestTag) SetKey(v string) *CreateLoadBalancerRequestTag {
	s.Key = &v
	return s
}

func (s *CreateLoadBalancerRequestTag) SetValue(v string) *CreateLoadBalancerRequestTag {
	s.Value = &v
	return s
}

type CreateLoadBalancerRequestZoneMappings struct {
	// The ID of the elastic IP address (EIP) that is associated with the Internet-facing NLB instance. You can specify one EIP for each zone. You must add at least two zones. You can add a maximum of 10 zones.
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The private IP address. You must add at least two zones. You can add a maximum of 10 zones.
	PrivateIPv4Address *string `json:"PrivateIPv4Address,omitempty" xml:"PrivateIPv4Address,omitempty"`
	// The vSwitch in the zone. You can specify only one vSwitch (subnet) in each zone of an NLB instance. You must add at least two zones. You can add a maximum of 10 zones.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone of the NLB instance. You must add at least two zones. You can add a maximum of 10 zones.
	//
	// You can call the [DescribeZones](~~443890~~) operation to query the most recent zone list.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateLoadBalancerRequestZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestZoneMappings) SetAllocationId(v string) *CreateLoadBalancerRequestZoneMappings {
	s.AllocationId = &v
	return s
}

func (s *CreateLoadBalancerRequestZoneMappings) SetPrivateIPv4Address(v string) *CreateLoadBalancerRequestZoneMappings {
	s.PrivateIPv4Address = &v
	return s
}

func (s *CreateLoadBalancerRequestZoneMappings) SetVSwitchId(v string) *CreateLoadBalancerRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *CreateLoadBalancerRequestZoneMappings) SetZoneId(v string) *CreateLoadBalancerRequestZoneMappings {
	s.ZoneId = &v
	return s
}

type CreateLoadBalancerResponseBody struct {
	// The ID of the NLB instance.
	LoadbalancerId *string `json:"LoadbalancerId,omitempty" xml:"LoadbalancerId,omitempty"`
	// The ID of the order for the NLB instance.
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLoadBalancerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerResponseBody) SetLoadbalancerId(v string) *CreateLoadBalancerResponseBody {
	s.LoadbalancerId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetOrderId(v int64) *CreateLoadBalancerResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetRequestId(v string) *CreateLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

type CreateLoadBalancerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLoadBalancerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLoadBalancerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerResponse) SetHeaders(v map[string]*string) *CreateLoadBalancerResponse {
	s.Headers = v
	return s
}

func (s *CreateLoadBalancerResponse) SetStatusCode(v int32) *CreateLoadBalancerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoadBalancerResponse) SetBody(v *CreateLoadBalancerResponseBody) *CreateLoadBalancerResponse {
	s.Body = v
	return s
}

type CreateSecurityPolicyRequest struct {
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** for each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck the request. Valid values:
	//
	// *   **true**: checks the request but does not create the security policy. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the security policy.
	//
	// The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-).
	SecurityPolicyName *string                           `json:"SecurityPolicyName,omitempty" xml:"SecurityPolicyName,omitempty"`
	Tag                []*CreateSecurityPolicyRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TlsVersions        []*string                         `json:"TlsVersions,omitempty" xml:"TlsVersions,omitempty" type:"Repeated"`
}

func (s CreateSecurityPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityPolicyRequest) SetCiphers(v []*string) *CreateSecurityPolicyRequest {
	s.Ciphers = v
	return s
}

func (s *CreateSecurityPolicyRequest) SetClientToken(v string) *CreateSecurityPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetDryRun(v bool) *CreateSecurityPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetRegionId(v string) *CreateSecurityPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetResourceGroupId(v string) *CreateSecurityPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetSecurityPolicyName(v string) *CreateSecurityPolicyRequest {
	s.SecurityPolicyName = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetTag(v []*CreateSecurityPolicyRequestTag) *CreateSecurityPolicyRequest {
	s.Tag = v
	return s
}

func (s *CreateSecurityPolicyRequest) SetTlsVersions(v []*string) *CreateSecurityPolicyRequest {
	s.TlsVersions = v
	return s
}

type CreateSecurityPolicyRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSecurityPolicyRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityPolicyRequestTag) GoString() string {
	return s.String()
}

func (s *CreateSecurityPolicyRequestTag) SetKey(v string) *CreateSecurityPolicyRequestTag {
	s.Key = &v
	return s
}

func (s *CreateSecurityPolicyRequestTag) SetValue(v string) *CreateSecurityPolicyRequestTag {
	s.Value = &v
	return s
}

type CreateSecurityPolicyResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the TLS security policy.
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
}

func (s CreateSecurityPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecurityPolicyResponseBody) SetJobId(v string) *CreateSecurityPolicyResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateSecurityPolicyResponseBody) SetRequestId(v string) *CreateSecurityPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecurityPolicyResponseBody) SetSecurityPolicyId(v string) *CreateSecurityPolicyResponseBody {
	s.SecurityPolicyId = &v
	return s
}

type CreateSecurityPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSecurityPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSecurityPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateSecurityPolicyResponse) SetHeaders(v map[string]*string) *CreateSecurityPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateSecurityPolicyResponse) SetStatusCode(v int32) *CreateSecurityPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecurityPolicyResponse) SetBody(v *CreateSecurityPolicyResponseBody) *CreateSecurityPolicyResponse {
	s.Body = v
	return s
}

type CreateServerGroupRequest struct {
	// The protocol version. Valid values:
	//
	// *   **ipv4:** IPv4. This is the default value.
	// *   **DualStack:** dual stack.
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// Specifies whether to enable all-port forwarding. Valid values:
	//
	// *   **true:** yes.
	// *   **false:** no. This is the default value.
	AnyPortEnabled *bool `json:"AnyPortEnabled,omitempty" xml:"AnyPortEnabled,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system automatically uses the value of **RequestId** as the value of **ClientToken**. The value of **RequestId** may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable connection draining. Valid values:
	//
	// *   **true:** yes.
	// *   **false:** no. This is the default value.
	ConnectionDrainEnabled *bool `json:"ConnectionDrainEnabled,omitempty" xml:"ConnectionDrainEnabled,omitempty"`
	// The timeout period of connection draining. Unit: seconds.
	//
	// Valid values: **10** to **900**.
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// *   **true:** performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false:** performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun            *bool                                      `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	HealthCheckConfig *CreateServerGroupRequestHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// Specifies whether to enable client IP preservation. Valid values:
	//
	// *   **true:** yes.
	// *   **false:** no. This is the default value.
	PreserveClientIpEnabled *bool `json:"PreserveClientIpEnabled,omitempty" xml:"PreserveClientIpEnabled,omitempty"`
	// The protocol used to forward requests to the backend servers. Valid values:
	//
	// *   **TCP:** This is the default value.
	// *   **UDP**
	// *   **TCPSSL**
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the server group belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// *   **Wrr:** The weighted round-robin algorithm is used. Backend servers with higher weights receive more requests than backend servers with lower weights. This is the default value.
	// *   **rr:** The round-robin algorithm is used. Requests are forwarded to backend servers in sequence.
	// *   **sch:** Source IP hashing is used. Requests from the same source IP address are forwarded to the same backend server.
	// *   **tch:** Four-element hashing is used. It specifies consistent hashing that is based on four factors: source IP address, destination IP address, source port, and destination port. Requests that contain the same information based on the four factors are forwarded to the same backend server.
	// *   **qch:** QUIC ID hashing is used. Requests that contain the same QUIC ID are forwarded to the same backend server.
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The name of the server group.
	//
	// The name must be 2 to 128 characters in length and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The name must start with a letter.
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// The type of the server group. Valid values:
	//
	// *   **Instance:** allows you to add servers of the **Ecs**, **Ens**, or **Eci** type. This is the default value.
	// *   **Ip:** allows you to add servers by specifying IP addresses.
	ServerGroupType *string                        `json:"ServerGroupType,omitempty" xml:"ServerGroupType,omitempty"`
	Tag             []*CreateServerGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC to which the server group belongs.
	//
	// >  If **ServerGroupType** is set to **Instance**, only servers in the specified VPC can be added to the server group.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequest) SetAddressIPVersion(v string) *CreateServerGroupRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *CreateServerGroupRequest) SetAnyPortEnabled(v bool) *CreateServerGroupRequest {
	s.AnyPortEnabled = &v
	return s
}

func (s *CreateServerGroupRequest) SetClientToken(v string) *CreateServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServerGroupRequest) SetConnectionDrainEnabled(v bool) *CreateServerGroupRequest {
	s.ConnectionDrainEnabled = &v
	return s
}

func (s *CreateServerGroupRequest) SetConnectionDrainTimeout(v int32) *CreateServerGroupRequest {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *CreateServerGroupRequest) SetDryRun(v bool) *CreateServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServerGroupRequest) SetHealthCheckConfig(v *CreateServerGroupRequestHealthCheckConfig) *CreateServerGroupRequest {
	s.HealthCheckConfig = v
	return s
}

func (s *CreateServerGroupRequest) SetPreserveClientIpEnabled(v bool) *CreateServerGroupRequest {
	s.PreserveClientIpEnabled = &v
	return s
}

func (s *CreateServerGroupRequest) SetProtocol(v string) *CreateServerGroupRequest {
	s.Protocol = &v
	return s
}

func (s *CreateServerGroupRequest) SetRegionId(v string) *CreateServerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServerGroupRequest) SetResourceGroupId(v string) *CreateServerGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServerGroupRequest) SetScheduler(v string) *CreateServerGroupRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateServerGroupRequest) SetServerGroupName(v string) *CreateServerGroupRequest {
	s.ServerGroupName = &v
	return s
}

func (s *CreateServerGroupRequest) SetServerGroupType(v string) *CreateServerGroupRequest {
	s.ServerGroupType = &v
	return s
}

func (s *CreateServerGroupRequest) SetTag(v []*CreateServerGroupRequestTag) *CreateServerGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateServerGroupRequest) SetVpcId(v string) *CreateServerGroupRequest {
	s.VpcId = &v
	return s
}

type CreateServerGroupRequestHealthCheckConfig struct {
	// The backend port that is used for health checks.
	//
	// Valid values: **0** to **65535**.
	//
	// Default value: **0**. If you set the value to 0, the port of a backend server is used for health checks.
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The maximum timeout period of a health check response. Unit: seconds.
	//
	// Valid values: **1** to **300**.
	//
	// Default value: **5**.
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// *   **$SERVER_IP:** the private IP address of a backend server.
	// *   **domain:** the domain name you want to use for health checks. The domain name must be 1 to 80 characters in length and can contain lowercase letters, digits, hyphens (-), and periods (.).
	//
	// >  This parameter takes effect only when you set **HealthCheckType** to **HTTP**.
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// Specifies whether to enable the health check feature. Valid values:
	//
	// *   **true:** yes. This is the default value.
	// *   **false:** no.
	HealthCheckEnabled  *bool     `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckHttpCode []*string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty" type:"Repeated"`
	// The interval between two consecutive health checks. Unit: seconds.
	//
	// Valid values: **5** to **5000**.
	//
	// Default value: **10**.
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The protocol that is used for health checks. Valid values: **TCP** (default) and **HTTP**.
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The path to which health check requests are sent.
	//
	// The path must be 1 to 80 characters in length, and can contain only letters, digits, and the following special characters: `- / . % ? # & =`. It can also contain the following extended characters: `_ ; ~ ! ( ) * [ ] @ $ ^ : \" , +`. The path must start with a forward slash (/).
	//
	// >  This parameter takes effect only when you set **HealthCheckType** to **HTTP**.
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail** to **success**.
	//
	// Valid values: **2** to **10**.
	//
	// Default value: **2**.
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The HTTP method that is used for health checks. Valid values: **GET** (default) and **HEAD**.
	//
	// >  This parameter takes effect only when you set **HealthCheckType** to **HTTP**.
	HttpCheckMethod *string `json:"HttpCheckMethod,omitempty" xml:"HttpCheckMethod,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success** to **fail**.
	//
	// Valid values: **2** to **10**.
	//
	// Default value: **2**.
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s CreateServerGroupRequestHealthCheckConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupRequestHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckConnectPort(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckConnectTimeout(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckDomain(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckDomain = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckEnabled(v bool) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckEnabled = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckHttpCode(v []*string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckHttpCode = v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckInterval(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckType(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckType = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckUrl(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckUrl = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthyThreshold(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHttpCheckMethod(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HttpCheckMethod = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetUnhealthyThreshold(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.UnhealthyThreshold = &v
	return s
}

type CreateServerGroupRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServerGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequestTag) SetKey(v string) *CreateServerGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServerGroupRequestTag) SetValue(v string) *CreateServerGroupRequestTag {
	s.Value = &v
	return s
}

type CreateServerGroupResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s CreateServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServerGroupResponseBody) SetJobId(v string) *CreateServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateServerGroupResponseBody) SetRequestId(v string) *CreateServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServerGroupResponseBody) SetServerGroupId(v string) *CreateServerGroupResponseBody {
	s.ServerGroupId = &v
	return s
}

type CreateServerGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateServerGroupResponse) SetHeaders(v map[string]*string) *CreateServerGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateServerGroupResponse) SetStatusCode(v int32) *CreateServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServerGroupResponse) SetBody(v *CreateServerGroupResponseBody) *CreateServerGroupResponse {
	s.Body = v
	return s
}

type DeleteListenerRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** of each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck the request. Valid values:
	//
	// *   **true**: prechecks the request but does not delete the listener. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the listener.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerRequest) GoString() string {
	return s.String()
}

func (s *DeleteListenerRequest) SetClientToken(v string) *DeleteListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteListenerRequest) SetDryRun(v bool) *DeleteListenerRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteListenerRequest) SetListenerId(v string) *DeleteListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *DeleteListenerRequest) SetRegionId(v string) *DeleteListenerRequest {
	s.RegionId = &v
	return s
}

type DeleteListenerResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteListenerResponseBody) SetJobId(v string) *DeleteListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteListenerResponseBody) SetRequestId(v string) *DeleteListenerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteListenerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerResponse) GoString() string {
	return s.String()
}

func (s *DeleteListenerResponse) SetHeaders(v map[string]*string) *DeleteListenerResponse {
	s.Headers = v
	return s
}

func (s *DeleteListenerResponse) SetStatusCode(v int32) *DeleteListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteListenerResponse) SetBody(v *DeleteListenerResponseBody) *DeleteListenerResponse {
	s.Body = v
	return s
}

type DeleteLoadBalancerRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId** as **ClientToken**. **RequestId** may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck this request. Valid values:
	//
	// *   **true**: prechecks the request without deleting the NLB instance. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the NLB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteLoadBalancerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerRequest) SetClientToken(v string) *DeleteLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetDryRun(v bool) *DeleteLoadBalancerRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetLoadBalancerId(v string) *DeleteLoadBalancerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetRegionId(v string) *DeleteLoadBalancerRequest {
	s.RegionId = &v
	return s
}

type DeleteLoadBalancerResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLoadBalancerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerResponseBody) SetJobId(v string) *DeleteLoadBalancerResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteLoadBalancerResponseBody) SetRequestId(v string) *DeleteLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLoadBalancerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLoadBalancerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLoadBalancerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerResponse) SetHeaders(v map[string]*string) *DeleteLoadBalancerResponse {
	s.Headers = v
	return s
}

func (s *DeleteLoadBalancerResponse) SetStatusCode(v int32) *DeleteLoadBalancerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLoadBalancerResponse) SetBody(v *DeleteLoadBalancerResponseBody) *DeleteLoadBalancerResponse {
	s.Body = v
	return s
}

type DeleteSecurityPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can only contain ASCII characters.
	//
	// >  If you do not set this parameter, the system automatically uses the value of **RequestId** as the value of **ClientToken**. **RequestId** of each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// *   **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the available regions.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the TLS security policy.
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
}

func (s DeleteSecurityPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityPolicyRequest) SetClientToken(v string) *DeleteSecurityPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSecurityPolicyRequest) SetDryRun(v bool) *DeleteSecurityPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteSecurityPolicyRequest) SetRegionId(v string) *DeleteSecurityPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSecurityPolicyRequest) SetSecurityPolicyId(v string) *DeleteSecurityPolicyRequest {
	s.SecurityPolicyId = &v
	return s
}

type DeleteSecurityPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecurityPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityPolicyResponseBody) SetRequestId(v string) *DeleteSecurityPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSecurityPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSecurityPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSecurityPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityPolicyResponse) SetHeaders(v map[string]*string) *DeleteSecurityPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityPolicyResponse) SetStatusCode(v int32) *DeleteSecurityPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecurityPolicyResponse) SetBody(v *DeleteSecurityPolicyResponseBody) *DeleteSecurityPolicyResponse {
	s.Body = v
	return s
}

type DeleteServerGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can only contain ASCII characters.
	//
	// >  If you do not set this parameter, the system automatically uses the value of **RequestId** as the value of **ClientToken**. **RequestId** of each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// *   **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s DeleteServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteServerGroupRequest) SetClientToken(v string) *DeleteServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServerGroupRequest) SetDryRun(v bool) *DeleteServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteServerGroupRequest) SetRegionId(v string) *DeleteServerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServerGroupRequest) SetServerGroupId(v string) *DeleteServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

type DeleteServerGroupResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServerGroupResponseBody) SetJobId(v string) *DeleteServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteServerGroupResponseBody) SetRequestId(v string) *DeleteServerGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServerGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteServerGroupResponse) SetHeaders(v map[string]*string) *DeleteServerGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteServerGroupResponse) SetStatusCode(v int32) *DeleteServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServerGroupResponse) SetBody(v *DeleteServerGroupResponseBody) *DeleteServerGroupResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The supported natural language. Valid values:
	//
	// *   **zh-CN**: Chinese
	// *   **en-US** (default): English
	// *   **ja**: Japanese
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** for each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The service code. Set the value to **nlb**.
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetClientToken(v string) *DescribeRegionsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeRegionsRequest) SetServiceCode(v string) *DescribeRegionsRequest {
	s.ServiceCode = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// A list of regions.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	// The name of the region.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the region service.
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	// The supported natural language. Valid values:
	//
	// *   **zh-CN**: Chinese
	// *   **en-US** (default): English
	// *   **ja**: Japanese
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId** as **ClientToken**. **RequestId** may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the region to which the zone belongs. You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service code. Set the value to **nlb**.
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetAcceptLanguage(v string) *DescribeZonesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeZonesRequest) SetClientToken(v string) *DescribeZonesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesRequest) SetServiceCode(v string) *DescribeZonesRequest {
	s.ServiceCode = &v
	return s
}

type DescribeZonesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of zones.
	Zones []*DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v []*DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	// The name of the zone.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID of the zone.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetLocalName(v string) *DescribeZonesResponseBodyZones {
	s.LocalName = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetZoneId(v string) *DescribeZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

type DescribeZonesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetStatusCode(v int32) *DescribeZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type DetachCommonBandwidthPackageFromLoadBalancerRequest struct {
	// The ID of the EIP bandwidth plan.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** of each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck the request. Valid values:
	//
	// *   **true**: prechecks the request but does not disassociate the NLB instance from the EIP bandwidth plan. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the NLB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachCommonBandwidthPackageFromLoadBalancerRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachCommonBandwidthPackageFromLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) SetBandwidthPackageId(v string) *DetachCommonBandwidthPackageFromLoadBalancerRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) SetClientToken(v string) *DetachCommonBandwidthPackageFromLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) SetDryRun(v bool) *DetachCommonBandwidthPackageFromLoadBalancerRequest {
	s.DryRun = &v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) SetLoadBalancerId(v string) *DetachCommonBandwidthPackageFromLoadBalancerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) SetRegionId(v string) *DetachCommonBandwidthPackageFromLoadBalancerRequest {
	s.RegionId = &v
	return s
}

type DetachCommonBandwidthPackageFromLoadBalancerResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachCommonBandwidthPackageFromLoadBalancerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachCommonBandwidthPackageFromLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerResponseBody) SetJobId(v string) *DetachCommonBandwidthPackageFromLoadBalancerResponseBody {
	s.JobId = &v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerResponseBody) SetRequestId(v string) *DetachCommonBandwidthPackageFromLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

type DetachCommonBandwidthPackageFromLoadBalancerResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachCommonBandwidthPackageFromLoadBalancerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachCommonBandwidthPackageFromLoadBalancerResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachCommonBandwidthPackageFromLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerResponse) SetHeaders(v map[string]*string) *DetachCommonBandwidthPackageFromLoadBalancerResponse {
	s.Headers = v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerResponse) SetStatusCode(v int32) *DetachCommonBandwidthPackageFromLoadBalancerResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerResponse) SetBody(v *DetachCommonBandwidthPackageFromLoadBalancerResponseBody) *DetachCommonBandwidthPackageFromLoadBalancerResponse {
	s.Body = v
	return s
}

type DisableLoadBalancerIpv6InternetRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** of each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck the request. Valid values:
	//
	// *   **true**: prechecks the request but does not change the network type of the NLB instance. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the NLB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableLoadBalancerIpv6InternetRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableLoadBalancerIpv6InternetRequest) GoString() string {
	return s.String()
}

func (s *DisableLoadBalancerIpv6InternetRequest) SetClientToken(v string) *DisableLoadBalancerIpv6InternetRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableLoadBalancerIpv6InternetRequest) SetDryRun(v bool) *DisableLoadBalancerIpv6InternetRequest {
	s.DryRun = &v
	return s
}

func (s *DisableLoadBalancerIpv6InternetRequest) SetLoadBalancerId(v string) *DisableLoadBalancerIpv6InternetRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DisableLoadBalancerIpv6InternetRequest) SetRegionId(v string) *DisableLoadBalancerIpv6InternetRequest {
	s.RegionId = &v
	return s
}

type DisableLoadBalancerIpv6InternetResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableLoadBalancerIpv6InternetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableLoadBalancerIpv6InternetResponseBody) GoString() string {
	return s.String()
}

func (s *DisableLoadBalancerIpv6InternetResponseBody) SetRequestId(v string) *DisableLoadBalancerIpv6InternetResponseBody {
	s.RequestId = &v
	return s
}

type DisableLoadBalancerIpv6InternetResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableLoadBalancerIpv6InternetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableLoadBalancerIpv6InternetResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableLoadBalancerIpv6InternetResponse) GoString() string {
	return s.String()
}

func (s *DisableLoadBalancerIpv6InternetResponse) SetHeaders(v map[string]*string) *DisableLoadBalancerIpv6InternetResponse {
	s.Headers = v
	return s
}

func (s *DisableLoadBalancerIpv6InternetResponse) SetStatusCode(v int32) *DisableLoadBalancerIpv6InternetResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableLoadBalancerIpv6InternetResponse) SetBody(v *DisableLoadBalancerIpv6InternetResponseBody) *DisableLoadBalancerIpv6InternetResponse {
	s.Body = v
	return s
}

type DisassociateAdditionalCertificatesWithListenerRequest struct {
	AdditionalCertificateIds []*string `json:"AdditionalCertificateIds,omitempty" xml:"AdditionalCertificateIds,omitempty" type:"Repeated"`
	ClientToken              *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                   *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	ListenerId               *string   `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId                 *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisassociateAdditionalCertificatesWithListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DisassociateAdditionalCertificatesWithListenerRequest) GoString() string {
	return s.String()
}

func (s *DisassociateAdditionalCertificatesWithListenerRequest) SetAdditionalCertificateIds(v []*string) *DisassociateAdditionalCertificatesWithListenerRequest {
	s.AdditionalCertificateIds = v
	return s
}

func (s *DisassociateAdditionalCertificatesWithListenerRequest) SetClientToken(v string) *DisassociateAdditionalCertificatesWithListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *DisassociateAdditionalCertificatesWithListenerRequest) SetDryRun(v bool) *DisassociateAdditionalCertificatesWithListenerRequest {
	s.DryRun = &v
	return s
}

func (s *DisassociateAdditionalCertificatesWithListenerRequest) SetListenerId(v string) *DisassociateAdditionalCertificatesWithListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *DisassociateAdditionalCertificatesWithListenerRequest) SetRegionId(v string) *DisassociateAdditionalCertificatesWithListenerRequest {
	s.RegionId = &v
	return s
}

type DisassociateAdditionalCertificatesWithListenerResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociateAdditionalCertificatesWithListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisassociateAdditionalCertificatesWithListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateAdditionalCertificatesWithListenerResponseBody) SetJobId(v string) *DisassociateAdditionalCertificatesWithListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *DisassociateAdditionalCertificatesWithListenerResponseBody) SetRequestId(v string) *DisassociateAdditionalCertificatesWithListenerResponseBody {
	s.RequestId = &v
	return s
}

type DisassociateAdditionalCertificatesWithListenerResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisassociateAdditionalCertificatesWithListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisassociateAdditionalCertificatesWithListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s DisassociateAdditionalCertificatesWithListenerResponse) GoString() string {
	return s.String()
}

func (s *DisassociateAdditionalCertificatesWithListenerResponse) SetHeaders(v map[string]*string) *DisassociateAdditionalCertificatesWithListenerResponse {
	s.Headers = v
	return s
}

func (s *DisassociateAdditionalCertificatesWithListenerResponse) SetStatusCode(v int32) *DisassociateAdditionalCertificatesWithListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateAdditionalCertificatesWithListenerResponse) SetBody(v *DisassociateAdditionalCertificatesWithListenerResponseBody) *DisassociateAdditionalCertificatesWithListenerResponse {
	s.Body = v
	return s
}

type EnableLoadBalancerIpv6InternetRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** for each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck the request. Valid values:
	//
	// *   **true**: prechecks the request but does not change the network type of the NLB instance. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the NLB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableLoadBalancerIpv6InternetRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableLoadBalancerIpv6InternetRequest) GoString() string {
	return s.String()
}

func (s *EnableLoadBalancerIpv6InternetRequest) SetClientToken(v string) *EnableLoadBalancerIpv6InternetRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableLoadBalancerIpv6InternetRequest) SetDryRun(v bool) *EnableLoadBalancerIpv6InternetRequest {
	s.DryRun = &v
	return s
}

func (s *EnableLoadBalancerIpv6InternetRequest) SetLoadBalancerId(v string) *EnableLoadBalancerIpv6InternetRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *EnableLoadBalancerIpv6InternetRequest) SetRegionId(v string) *EnableLoadBalancerIpv6InternetRequest {
	s.RegionId = &v
	return s
}

type EnableLoadBalancerIpv6InternetResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableLoadBalancerIpv6InternetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableLoadBalancerIpv6InternetResponseBody) GoString() string {
	return s.String()
}

func (s *EnableLoadBalancerIpv6InternetResponseBody) SetRequestId(v string) *EnableLoadBalancerIpv6InternetResponseBody {
	s.RequestId = &v
	return s
}

type EnableLoadBalancerIpv6InternetResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableLoadBalancerIpv6InternetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableLoadBalancerIpv6InternetResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableLoadBalancerIpv6InternetResponse) GoString() string {
	return s.String()
}

func (s *EnableLoadBalancerIpv6InternetResponse) SetHeaders(v map[string]*string) *EnableLoadBalancerIpv6InternetResponse {
	s.Headers = v
	return s
}

func (s *EnableLoadBalancerIpv6InternetResponse) SetStatusCode(v int32) *EnableLoadBalancerIpv6InternetResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableLoadBalancerIpv6InternetResponse) SetBody(v *EnableLoadBalancerIpv6InternetResponseBody) *EnableLoadBalancerIpv6InternetResponse {
	s.Body = v
	return s
}

type GetJobStatusRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** of each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusRequest) GoString() string {
	return s.String()
}

func (s *GetJobStatusRequest) SetClientToken(v string) *GetJobStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *GetJobStatusRequest) SetJobId(v string) *GetJobStatusRequest {
	s.JobId = &v
	return s
}

type GetJobStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the task. Valid values:
	//
	// *   **Succeeded**: The task is successful.
	// *   **processing**: The ticket is being executed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetJobStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponseBody) SetRequestId(v string) *GetJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobStatusResponseBody) SetStatus(v string) *GetJobStatusResponseBody {
	s.Status = &v
	return s
}

type GetJobStatusResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusResponse) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponse) SetHeaders(v map[string]*string) *GetJobStatusResponse {
	s.Headers = v
	return s
}

func (s *GetJobStatusResponse) SetStatusCode(v int32) *GetJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobStatusResponse) SetBody(v *GetJobStatusResponseBody) *GetJobStatusResponse {
	s.Body = v
	return s
}

type GetListenerAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** of each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck the request. Valid values:
	//
	// *   **true**: checks the request but does not query the listener details. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the listener.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the region where the Network Load Balancer (NLB) instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeRequest) SetClientToken(v string) *GetListenerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *GetListenerAttributeRequest) SetDryRun(v bool) *GetListenerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *GetListenerAttributeRequest) SetListenerId(v string) *GetListenerAttributeRequest {
	s.ListenerId = &v
	return s
}

func (s *GetListenerAttributeRequest) SetRegionId(v string) *GetListenerAttributeRequest {
	s.RegionId = &v
	return s
}

type GetListenerAttributeResponseBody struct {
	// Indicates whether Application-Layer Protocol Negotiation (ALPN) is enabled. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	AlpnEnabled *bool `json:"AlpnEnabled,omitempty" xml:"AlpnEnabled,omitempty"`
	// The ALPN policy. Valid values:
	//
	// *   **HTTP1Only**
	// *   **HTTP2Only**
	// *   **HTTP2Preferred**
	// *   **HTTP2Optional**
	AlpnPolicy *string `json:"AlpnPolicy,omitempty" xml:"AlpnPolicy,omitempty"`
	// The CA certificates. Only one CA certificate is supported.
	//
	// >  This parameter takes effect only for listeners that use SSL over TCP.
	CaCertificateIds []*string `json:"CaCertificateIds,omitempty" xml:"CaCertificateIds,omitempty" type:"Repeated"`
	// Indicates whether mutual authentication is enabled. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	CaEnabled *bool `json:"CaEnabled,omitempty" xml:"CaEnabled,omitempty"`
	// The server certificates. Only one server certificate is supported.
	//
	// >  This parameter takes effect only for listeners that use SSL over TCP.
	CertificateIds []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	// The maximum number of connections that can be created per second on the NLB instance. Valid values: **0** to **1000000**. **0** specifies that the number of connections is unlimited.
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// The last port in the listening port range. Valid values: **0** to **65535**. The number of the last port must be smaller than that of the first port.
	EndPort *string `json:"EndPort,omitempty" xml:"EndPort,omitempty"`
	// The timeout period of an idle connection. Unit: seconds. Valid values: **1** to **900**.
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The name of the listener.
	//
	// The name must be 2 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (\_), and hyphens (-).
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The ID of the listener.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The listening port. Valid values: **0** to **65535**. A value of **0** specifies all ports. If you set the value to **0**, you must also set the **StartPort** and **EndPort** parameters.
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The listening protocol. Valid values: **TCP**, **UDP**, and **TCPSSL**.
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The status of the listener. Valid values:
	//
	// *   **Provisioning**: The listener is being created.
	// *   **Running**: The listener is running.
	// *   **Configuring**: The listener is being configured.
	// *   **Stopping**: The listener is being stopped.
	// *   **Stopped**: The listener is stopped.
	// *   **Starting**: The listener is being started.
	// *   **Deleting**: The listener is being deleted.
	// *   **Deleted**: The listener is deleted.
	ListenerStatus *string `json:"ListenerStatus,omitempty" xml:"ListenerStatus,omitempty"`
	// The ID of the NLB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The size of the largest TCP segment. Unit: bytes. Valid values: **0** to **1500**. **0** specifies that the maximum segment size remains unchanged.
	//
	// >  This parameter is supported only by listeners that use SSL over TCP.
	Mss *int32 `json:"Mss,omitempty" xml:"Mss,omitempty"`
	// Indicates whether the Proxy protocol is used to pass client IP addresses to backend servers. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	ProxyProtocolEnabled *bool `json:"ProxyProtocolEnabled,omitempty" xml:"ProxyProtocolEnabled,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether fine-grained monitoring is enabled. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	SecSensorEnabled *bool `json:"SecSensorEnabled,omitempty" xml:"SecSensorEnabled,omitempty"`
	// The ID of the security policy. System security policies and custom security policies are supported.
	//
	// Valid values: **tls_cipher_policy\_1\_0**, **tls_cipher_policy\_1\_1**, **tls_cipher_policy\_1\_2**, **tls_cipher_policy\_1\_2\_strict**, and **tls_cipher_policy\_1\_2\_strict_with\_1\_3**.
	//
	// >  This parameter takes effect only for listeners that use SSL over TCP.
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The first port in the listening port range. Valid values: **0** to **65535**.
	StartPort *string                                 `json:"StartPort,omitempty" xml:"StartPort,omitempty"`
	Tags      []*GetListenerAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBody) SetAlpnEnabled(v bool) *GetListenerAttributeResponseBody {
	s.AlpnEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetAlpnPolicy(v string) *GetListenerAttributeResponseBody {
	s.AlpnPolicy = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetCaCertificateIds(v []*string) *GetListenerAttributeResponseBody {
	s.CaCertificateIds = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetCaEnabled(v bool) *GetListenerAttributeResponseBody {
	s.CaEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetCertificateIds(v []*string) *GetListenerAttributeResponseBody {
	s.CertificateIds = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetCps(v int32) *GetListenerAttributeResponseBody {
	s.Cps = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetEndPort(v string) *GetListenerAttributeResponseBody {
	s.EndPort = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetIdleTimeout(v int32) *GetListenerAttributeResponseBody {
	s.IdleTimeout = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerDescription(v string) *GetListenerAttributeResponseBody {
	s.ListenerDescription = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerId(v string) *GetListenerAttributeResponseBody {
	s.ListenerId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerPort(v int32) *GetListenerAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerProtocol(v string) *GetListenerAttributeResponseBody {
	s.ListenerProtocol = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerStatus(v string) *GetListenerAttributeResponseBody {
	s.ListenerStatus = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetLoadBalancerId(v string) *GetListenerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetMss(v int32) *GetListenerAttributeResponseBody {
	s.Mss = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetProxyProtocolEnabled(v bool) *GetListenerAttributeResponseBody {
	s.ProxyProtocolEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetRegionId(v string) *GetListenerAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetRequestId(v string) *GetListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetSecSensorEnabled(v bool) *GetListenerAttributeResponseBody {
	s.SecSensorEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetSecurityPolicyId(v string) *GetListenerAttributeResponseBody {
	s.SecurityPolicyId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetServerGroupId(v string) *GetListenerAttributeResponseBody {
	s.ServerGroupId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetStartPort(v string) *GetListenerAttributeResponseBody {
	s.StartPort = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetTags(v []*GetListenerAttributeResponseBodyTags) *GetListenerAttributeResponseBody {
	s.Tags = v
	return s
}

type GetListenerAttributeResponseBodyTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetListenerAttributeResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyTags) SetTagKey(v string) *GetListenerAttributeResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *GetListenerAttributeResponseBodyTags) SetTagValue(v string) *GetListenerAttributeResponseBodyTags {
	s.TagValue = &v
	return s
}

type GetListenerAttributeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetListenerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponse) SetHeaders(v map[string]*string) *GetListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetListenerAttributeResponse) SetStatusCode(v int32) *GetListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetListenerAttributeResponse) SetBody(v *GetListenerAttributeResponseBody) *GetListenerAttributeResponse {
	s.Body = v
	return s
}

type GetListenerHealthStatusRequest struct {
	// The ID of the listener of the NLB instance.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The number of entries to return on each page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If this is your first query or no next query is to be sent, ignore this parameter.
	// *   If a next query is to be sent, set the parameter to the value of NextToken that is returned from the last call.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetListenerHealthStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusRequest) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusRequest) SetListenerId(v string) *GetListenerHealthStatusRequest {
	s.ListenerId = &v
	return s
}

func (s *GetListenerHealthStatusRequest) SetMaxResults(v int32) *GetListenerHealthStatusRequest {
	s.MaxResults = &v
	return s
}

func (s *GetListenerHealthStatusRequest) SetNextToken(v string) *GetListenerHealthStatusRequest {
	s.NextToken = &v
	return s
}

func (s *GetListenerHealthStatusRequest) SetRegionId(v string) *GetListenerHealthStatusRequest {
	s.RegionId = &v
	return s
}

type GetListenerHealthStatusResponseBody struct {
	// The health check status of the server groups that are associated with the listener.
	ListenerHealthStatus []*GetListenerHealthStatusResponseBodyListenerHealthStatus `json:"ListenerHealthStatus,omitempty" xml:"ListenerHealthStatus,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// - If **NextToken** is empty, it indicates that no next query is to be sent.
	// - If a value of **NextToken** is returned, the value is the token used for the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetListenerHealthStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBody) SetListenerHealthStatus(v []*GetListenerHealthStatusResponseBodyListenerHealthStatus) *GetListenerHealthStatusResponseBody {
	s.ListenerHealthStatus = v
	return s
}

func (s *GetListenerHealthStatusResponseBody) SetMaxResults(v int32) *GetListenerHealthStatusResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetListenerHealthStatusResponseBody) SetNextToken(v string) *GetListenerHealthStatusResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetListenerHealthStatusResponseBody) SetRequestId(v string) *GetListenerHealthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBody) SetTotalCount(v int32) *GetListenerHealthStatusResponseBody {
	s.TotalCount = &v
	return s
}

type GetListenerHealthStatusResponseBodyListenerHealthStatus struct {
	// The ID of the listener of the NLB instance.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The listening port.
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The listening protocol. Valid values: **TCP**, **UDP**, and **TCPSSL**.
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The information about the server groups.
	ServerGroupInfos []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos `json:"ServerGroupInfos,omitempty" xml:"ServerGroupInfos,omitempty" type:"Repeated"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatus) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatus) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) SetListenerId(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatus {
	s.ListenerId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) SetListenerPort(v int32) *GetListenerHealthStatusResponseBodyListenerHealthStatus {
	s.ListenerPort = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) SetListenerProtocol(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatus {
	s.ListenerProtocol = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) SetServerGroupInfos(v []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) *GetListenerHealthStatusResponseBodyListenerHealthStatus {
	s.ServerGroupInfos = v
	return s
}

type GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos struct {
	// Indicates whether the health check feature is enabled. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	HeathCheckEnabled *bool `json:"HeathCheckEnabled,omitempty" xml:"HeathCheckEnabled,omitempty"`
	// A list of unhealthy backend servers.
	NonNormalServers []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers `json:"NonNormalServers,omitempty" xml:"NonNormalServers,omitempty" type:"Repeated"`
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) SetHeathCheckEnabled(v bool) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	s.HeathCheckEnabled = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) SetNonNormalServers(v []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	s.NonNormalServers = v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) SetServerGroupId(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	s.ServerGroupId = &v
	return s
}

type GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers struct {
	// The backend port.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The cause of the health check failure.
	Reason *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason `json:"Reason,omitempty" xml:"Reason,omitempty" type:"Struct"`
	// The ID of the backend server.
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the backend server.
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The health check status. Valid values:
	//
	// *   **Initial**: indicates that health checks are configured for the NLB instance, but no data was found.
	// *   **Unhealthy**: indicates that the backend server consecutively fails health checks.
	// *   **Unused**: indicates that the weight of the backend server is 0.
	// *   **Unavailable**: indicates that health checks are disabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) SetPort(v int32) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers {
	s.Port = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) SetReason(v *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers {
	s.Reason = v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) SetServerId(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers {
	s.ServerId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) SetServerIp(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers {
	s.ServerIp = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers) SetStatus(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServers {
	s.Status = &v
	return s
}

type GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason struct {
	// The reason why the **status** is abnormal. Valid values:
	//
	// *   **CONNECT_TIMEOUT**: The NLB instance failed to connect to the backend server within the specified period of time.
	// *   **CONNECT_FAILED**: The NLB instance failed to connect to the backend server.
	// *   **RECV_RESPONSE_TIMEOUT**: The NLB instance failed to receive a response from the backend server within the specified period of time.
	// *   **CONNECT_INTERRUPT**: The connection between the health check and the backend servers was interrupted.
	// *   **HTTP_CODE_NOT_MATCH**: The HTTP status code from the backend servers was not the expected one.
	// *   **HTTP_INVALID_HEADER**: The format of the response from the backend servers is invalid.
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason) SetReasonCode(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosNonNormalServersReason {
	s.ReasonCode = &v
	return s
}

type GetListenerHealthStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetListenerHealthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetListenerHealthStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponse) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponse) SetHeaders(v map[string]*string) *GetListenerHealthStatusResponse {
	s.Headers = v
	return s
}

func (s *GetListenerHealthStatusResponse) SetStatusCode(v int32) *GetListenerHealthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetListenerHealthStatusResponse) SetBody(v *GetListenerHealthStatusResponseBody) *GetListenerHealthStatusResponse {
	s.Body = v
	return s
}

type GetLoadBalancerAttributeRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun         *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetLoadBalancerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeRequest) SetClientToken(v string) *GetLoadBalancerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *GetLoadBalancerAttributeRequest) SetDryRun(v bool) *GetLoadBalancerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *GetLoadBalancerAttributeRequest) SetLoadBalancerId(v string) *GetLoadBalancerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *GetLoadBalancerAttributeRequest) SetRegionId(v string) *GetLoadBalancerAttributeRequest {
	s.RegionId = &v
	return s
}

type GetLoadBalancerAttributeResponseBody struct {
	AddressIpVersion             *string                                                           `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	AddressType                  *string                                                           `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	BandwidthPackageId           *string                                                           `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	Cps                          *int32                                                            `json:"Cps,omitempty" xml:"Cps,omitempty"`
	CreateTime                   *string                                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CrossZoneEnabled             *bool                                                             `json:"CrossZoneEnabled,omitempty" xml:"CrossZoneEnabled,omitempty"`
	DNSName                      *string                                                           `json:"DNSName,omitempty" xml:"DNSName,omitempty"`
	DeletionProtectionConfig     *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig     `json:"DeletionProtectionConfig,omitempty" xml:"DeletionProtectionConfig,omitempty" type:"Struct"`
	Ipv6AddressType              *string                                                           `json:"Ipv6AddressType,omitempty" xml:"Ipv6AddressType,omitempty"`
	LoadBalancerBillingConfig    *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig    `json:"LoadBalancerBillingConfig,omitempty" xml:"LoadBalancerBillingConfig,omitempty" type:"Struct"`
	LoadBalancerBusinessStatus   *string                                                           `json:"LoadBalancerBusinessStatus,omitempty" xml:"LoadBalancerBusinessStatus,omitempty"`
	LoadBalancerId               *string                                                           `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerName             *string                                                           `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	LoadBalancerStatus           *string                                                           `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	LoadBalancerType             *string                                                           `json:"LoadBalancerType,omitempty" xml:"LoadBalancerType,omitempty"`
	ModificationProtectionConfig *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig `json:"ModificationProtectionConfig,omitempty" xml:"ModificationProtectionConfig,omitempty" type:"Struct"`
	OperationLocks               []*GetLoadBalancerAttributeResponseBodyOperationLocks             `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Repeated"`
	RegionId                     *string                                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId                    *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId              *string                                                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityGroupIds             []*string                                                         `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	Tags                         []*GetLoadBalancerAttributeResponseBodyTags                       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	VpcId                        *string                                                           `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneMappings                 []*GetLoadBalancerAttributeResponseBodyZoneMappings               `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s GetLoadBalancerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBody) SetAddressIpVersion(v string) *GetLoadBalancerAttributeResponseBody {
	s.AddressIpVersion = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetAddressType(v string) *GetLoadBalancerAttributeResponseBody {
	s.AddressType = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetBandwidthPackageId(v string) *GetLoadBalancerAttributeResponseBody {
	s.BandwidthPackageId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetCps(v int32) *GetLoadBalancerAttributeResponseBody {
	s.Cps = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetCreateTime(v string) *GetLoadBalancerAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetCrossZoneEnabled(v bool) *GetLoadBalancerAttributeResponseBody {
	s.CrossZoneEnabled = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetDNSName(v string) *GetLoadBalancerAttributeResponseBody {
	s.DNSName = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetDeletionProtectionConfig(v *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) *GetLoadBalancerAttributeResponseBody {
	s.DeletionProtectionConfig = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetIpv6AddressType(v string) *GetLoadBalancerAttributeResponseBody {
	s.Ipv6AddressType = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerBillingConfig(v *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerBillingConfig = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerBusinessStatus(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerBusinessStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerId(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerName(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerName = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerStatus(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerType(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerType = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetModificationProtectionConfig(v *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) *GetLoadBalancerAttributeResponseBody {
	s.ModificationProtectionConfig = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetOperationLocks(v []*GetLoadBalancerAttributeResponseBodyOperationLocks) *GetLoadBalancerAttributeResponseBody {
	s.OperationLocks = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetRegionId(v string) *GetLoadBalancerAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetRequestId(v string) *GetLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetResourceGroupId(v string) *GetLoadBalancerAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetSecurityGroupIds(v []*string) *GetLoadBalancerAttributeResponseBody {
	s.SecurityGroupIds = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetTags(v []*GetLoadBalancerAttributeResponseBodyTags) *GetLoadBalancerAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetVpcId(v string) *GetLoadBalancerAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetZoneMappings(v []*GetLoadBalancerAttributeResponseBodyZoneMappings) *GetLoadBalancerAttributeResponseBody {
	s.ZoneMappings = v
	return s
}

type GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig struct {
	Enabled     *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	EnabledTime *string `json:"EnabledTime,omitempty" xml:"EnabledTime,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) SetEnabled(v bool) *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig {
	s.Enabled = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) SetEnabledTime(v string) *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig {
	s.EnabledTime = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) SetReason(v string) *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig {
	s.Reason = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig struct {
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) SetPayType(v string) *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig {
	s.PayType = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyModificationProtectionConfig struct {
	EnabledTime *string `json:"EnabledTime,omitempty" xml:"EnabledTime,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) SetEnabledTime(v string) *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig {
	s.EnabledTime = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) SetReason(v string) *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig {
	s.Reason = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) SetStatus(v string) *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig {
	s.Status = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyOperationLocks struct {
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	LockType   *string `json:"LockType,omitempty" xml:"LockType,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyOperationLocks) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyOperationLocks) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyOperationLocks) SetLockReason(v string) *GetLoadBalancerAttributeResponseBodyOperationLocks {
	s.LockReason = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyOperationLocks) SetLockType(v string) *GetLoadBalancerAttributeResponseBodyOperationLocks {
	s.LockType = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyTags) SetTagKey(v string) *GetLoadBalancerAttributeResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyTags) SetTagValue(v string) *GetLoadBalancerAttributeResponseBodyTags {
	s.TagValue = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyZoneMappings struct {
	LoadBalancerAddresses []*GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses `json:"LoadBalancerAddresses,omitempty" xml:"LoadBalancerAddresses,omitempty" type:"Repeated"`
	Status                *string                                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchId             *string                                                                  `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                *string                                                                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappings) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetLoadBalancerAddresses(v []*GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.LoadBalancerAddresses = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetStatus(v string) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.Status = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetVSwitchId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetZoneId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.ZoneId = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses struct {
	AllocationId        *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	EniId               *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	Ipv6Address         *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	PrivateIPv4Address  *string `json:"PrivateIPv4Address,omitempty" xml:"PrivateIPv4Address,omitempty"`
	PrivateIPv4HcStatus *string `json:"PrivateIPv4HcStatus,omitempty" xml:"PrivateIPv4HcStatus,omitempty"`
	PrivateIPv6HcStatus *string `json:"PrivateIPv6HcStatus,omitempty" xml:"PrivateIPv6HcStatus,omitempty"`
	PublicIPv4Address   *string `json:"PublicIPv4Address,omitempty" xml:"PublicIPv4Address,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetAllocationId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.AllocationId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetEniId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.EniId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetIpv6Address(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.Ipv6Address = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetPrivateIPv4Address(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.PrivateIPv4Address = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetPrivateIPv4HcStatus(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.PrivateIPv4HcStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetPrivateIPv6HcStatus(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.PrivateIPv6HcStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetPublicIPv4Address(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.PublicIPv4Address = &v
	return s
}

type GetLoadBalancerAttributeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLoadBalancerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLoadBalancerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponse) SetHeaders(v map[string]*string) *GetLoadBalancerAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetLoadBalancerAttributeResponse) SetStatusCode(v int32) *GetLoadBalancerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoadBalancerAttributeResponse) SetBody(v *GetLoadBalancerAttributeResponseBody) *GetLoadBalancerAttributeResponse {
	s.Body = v
	return s
}

type ListListenerCertificatesRequest struct {
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The ID of the listener. Specify the ID of a listener that uses SSL over TCP.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The number of entries to return on each page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// *   You do not need to specify this parameter for the first request.
	// *   You must specify the token that is obtained from the previous query as the value of NextToken.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the Network Load Balancer (NLB) instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListListenerCertificatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListListenerCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesRequest) SetCertType(v string) *ListListenerCertificatesRequest {
	s.CertType = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetListenerId(v string) *ListListenerCertificatesRequest {
	s.ListenerId = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetMaxResults(v int32) *ListListenerCertificatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetNextToken(v string) *ListListenerCertificatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetRegionId(v string) *ListListenerCertificatesRequest {
	s.RegionId = &v
	return s
}

type ListListenerCertificatesResponseBody struct {
	// The server certificates.
	CertificateIds []*string                                           `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	Certificates   []*ListListenerCertificatesResponseBodyCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The number of entries returned per page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// *   You do not need to specify this parameter for the first request.
	// *   You must specify the token that is obtained from the previous query as the value of NextToken.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListListenerCertificatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListListenerCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesResponseBody) SetCertificateIds(v []*string) *ListListenerCertificatesResponseBody {
	s.CertificateIds = v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetCertificates(v []*ListListenerCertificatesResponseBodyCertificates) *ListListenerCertificatesResponseBody {
	s.Certificates = v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetMaxResults(v int32) *ListListenerCertificatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetNextToken(v string) *ListListenerCertificatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetRequestId(v string) *ListListenerCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetTotalCount(v int32) *ListListenerCertificatesResponseBody {
	s.TotalCount = &v
	return s
}

type ListListenerCertificatesResponseBodyCertificates struct {
	CertificateId   *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	IsDefault       *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListListenerCertificatesResponseBodyCertificates) String() string {
	return tea.Prettify(s)
}

func (s ListListenerCertificatesResponseBodyCertificates) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetCertificateId(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.CertificateId = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetCertificateType(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.CertificateType = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetIsDefault(v bool) *ListListenerCertificatesResponseBodyCertificates {
	s.IsDefault = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetStatus(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.Status = &v
	return s
}

type ListListenerCertificatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListListenerCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListListenerCertificatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListListenerCertificatesResponse) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesResponse) SetHeaders(v map[string]*string) *ListListenerCertificatesResponse {
	s.Headers = v
	return s
}

func (s *ListListenerCertificatesResponse) SetStatusCode(v int32) *ListListenerCertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListListenerCertificatesResponse) SetBody(v *ListListenerCertificatesResponseBody) *ListListenerCertificatesResponse {
	s.Body = v
	return s
}

type ListListenersRequest struct {
	ListenerIds []*string `json:"ListenerIds,omitempty" xml:"ListenerIds,omitempty" type:"Repeated"`
	// The listening protocol. Valid values: **TCP**, **UDP**, and **TCPSSL**.
	ListenerProtocol *string   `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	LoadBalancerIds  []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If this is your first query or no next query is to be sent, ignore this parameter.
	// *   If a next query is to be sent, set the parameter to the value of NextToken that is returned from the last call.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tag      []*ListListenersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListListenersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListListenersRequest) GoString() string {
	return s.String()
}

func (s *ListListenersRequest) SetListenerIds(v []*string) *ListListenersRequest {
	s.ListenerIds = v
	return s
}

func (s *ListListenersRequest) SetListenerProtocol(v string) *ListListenersRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *ListListenersRequest) SetLoadBalancerIds(v []*string) *ListListenersRequest {
	s.LoadBalancerIds = v
	return s
}

func (s *ListListenersRequest) SetMaxResults(v int32) *ListListenersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListListenersRequest) SetNextToken(v string) *ListListenersRequest {
	s.NextToken = &v
	return s
}

func (s *ListListenersRequest) SetRegionId(v string) *ListListenersRequest {
	s.RegionId = &v
	return s
}

func (s *ListListenersRequest) SetTag(v []*ListListenersRequestTag) *ListListenersRequest {
	s.Tag = v
	return s
}

type ListListenersRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListListenersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListListenersRequestTag) GoString() string {
	return s.String()
}

func (s *ListListenersRequestTag) SetKey(v string) *ListListenersRequestTag {
	s.Key = &v
	return s
}

func (s *ListListenersRequestTag) SetValue(v string) *ListListenersRequestTag {
	s.Value = &v
	return s
}

type ListListenersResponseBody struct {
	// The list of listeners.
	Listeners []*ListListenersResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If **NextToken** is empty, it indicates that no next query is to be sent.
	// *   If a value of **NextToken** is returned, the value is the token used for the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListListenersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBody) SetListeners(v []*ListListenersResponseBodyListeners) *ListListenersResponseBody {
	s.Listeners = v
	return s
}

func (s *ListListenersResponseBody) SetMaxResults(v int32) *ListListenersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListListenersResponseBody) SetNextToken(v string) *ListListenersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListListenersResponseBody) SetRequestId(v string) *ListListenersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenersResponseBody) SetTotalCount(v int32) *ListListenersResponseBody {
	s.TotalCount = &v
	return s
}

type ListListenersResponseBodyListeners struct {
	// Indicates whether Application-Layer Protocol Negotiation (ALPN) is enabled. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	AlpnEnabled *bool `json:"AlpnEnabled,omitempty" xml:"AlpnEnabled,omitempty"`
	// The ALPN policy. Valid values:
	//
	// *   **HTTP1Only**
	// *   **HTTP2Only**
	// *   **HTTP2Preferred**
	// *   **HTTP2Optional**
	AlpnPolicy *string `json:"AlpnPolicy,omitempty" xml:"AlpnPolicy,omitempty"`
	// The list of CA certificates.
	//
	// >  This parameter takes effect only for listeners that use SSL over TCP.
	CaCertificateIds []*string `json:"CaCertificateIds,omitempty" xml:"CaCertificateIds,omitempty" type:"Repeated"`
	// Indicates whether mutual authentication is enabled. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	CaEnabled *bool `json:"CaEnabled,omitempty" xml:"CaEnabled,omitempty"`
	// The list of server certificates.
	//
	// >  This parameter takes effect only for listeners that use SSL over TCP.
	CertificateIds []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	// The maximum number of connections that can be created per second on the NLB instance. Valid values: **0** to **1000000**. **0** indicates that the number of connections is unlimited.
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// The last port in the listening port range.
	EndPort *string `json:"EndPort,omitempty" xml:"EndPort,omitempty"`
	// The timeout period of an idle connection. Unit: seconds. Valid values: **1** to **900**. Default value: **900**.
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// The name of the listener.
	//
	// The name must be 2 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (\_), and hyphens (-).
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The ID of the listener.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The listening port.
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The listening protocol. Valid values: **TCP**, **UDP**, and **TCPSSL**.
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The status of the listener. Valid values:
	//
	// *   **Provisioning**
	// *   **Running**
	// *   **Configuring**
	// *   **Stopping**
	// *   **Stopped**
	// *   **Starting**
	// *   **Deleting**
	// *   **Deleted**
	ListenerStatus *string `json:"ListenerStatus,omitempty" xml:"ListenerStatus,omitempty"`
	// The ID of the NLB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The maximum size of a TCP segment. Unit: bytes. Valid values: **0** to **1500**. **0** indicates that the maximum segment size remains unchanged.
	//
	// >  This parameter is supported only by listeners that use SSL over TCP.
	Mss *int32 `json:"Mss,omitempty" xml:"Mss,omitempty"`
	// Indicates whether the Proxy protocol is used to pass client IP addresses to backend servers. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	ProxyProtocolEnabled *bool `json:"ProxyProtocolEnabled,omitempty" xml:"ProxyProtocolEnabled,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Indicates whether fine-grained monitoring is enabled. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	SecSensorEnabled *bool `json:"SecSensorEnabled,omitempty" xml:"SecSensorEnabled,omitempty"`
	// The ID of the security policy.
	//
	// >  This parameter takes effect only for listeners that use SSL over TCP.
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The first port in the listening port range.
	StartPort *string                                   `json:"StartPort,omitempty" xml:"StartPort,omitempty"`
	Tags      []*ListListenersResponseBodyListenersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListListenersResponseBodyListeners) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListeners) SetAlpnEnabled(v bool) *ListListenersResponseBodyListeners {
	s.AlpnEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetAlpnPolicy(v string) *ListListenersResponseBodyListeners {
	s.AlpnPolicy = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetCaCertificateIds(v []*string) *ListListenersResponseBodyListeners {
	s.CaCertificateIds = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetCaEnabled(v bool) *ListListenersResponseBodyListeners {
	s.CaEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetCertificateIds(v []*string) *ListListenersResponseBodyListeners {
	s.CertificateIds = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetCps(v int32) *ListListenersResponseBodyListeners {
	s.Cps = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetEndPort(v string) *ListListenersResponseBodyListeners {
	s.EndPort = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetIdleTimeout(v int32) *ListListenersResponseBodyListeners {
	s.IdleTimeout = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerDescription(v string) *ListListenersResponseBodyListeners {
	s.ListenerDescription = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerId(v string) *ListListenersResponseBodyListeners {
	s.ListenerId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerPort(v int32) *ListListenersResponseBodyListeners {
	s.ListenerPort = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerProtocol(v string) *ListListenersResponseBodyListeners {
	s.ListenerProtocol = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerStatus(v string) *ListListenersResponseBodyListeners {
	s.ListenerStatus = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetLoadBalancerId(v string) *ListListenersResponseBodyListeners {
	s.LoadBalancerId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetMss(v int32) *ListListenersResponseBodyListeners {
	s.Mss = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetProxyProtocolEnabled(v bool) *ListListenersResponseBodyListeners {
	s.ProxyProtocolEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetRegionId(v string) *ListListenersResponseBodyListeners {
	s.RegionId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetSecSensorEnabled(v bool) *ListListenersResponseBodyListeners {
	s.SecSensorEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetSecurityPolicyId(v string) *ListListenersResponseBodyListeners {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetServerGroupId(v string) *ListListenersResponseBodyListeners {
	s.ServerGroupId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetStartPort(v string) *ListListenersResponseBodyListeners {
	s.StartPort = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetTags(v []*ListListenersResponseBodyListenersTags) *ListListenersResponseBodyListeners {
	s.Tags = v
	return s
}

type ListListenersResponseBodyListenersTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListListenersResponseBodyListenersTags) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersTags) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersTags) SetKey(v string) *ListListenersResponseBodyListenersTags {
	s.Key = &v
	return s
}

func (s *ListListenersResponseBodyListenersTags) SetValue(v string) *ListListenersResponseBodyListenersTags {
	s.Value = &v
	return s
}

type ListListenersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListListenersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListListenersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponse) GoString() string {
	return s.String()
}

func (s *ListListenersResponse) SetHeaders(v map[string]*string) *ListListenersResponse {
	s.Headers = v
	return s
}

func (s *ListListenersResponse) SetStatusCode(v int32) *ListListenersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListListenersResponse) SetBody(v *ListListenersResponseBody) *ListListenersResponse {
	s.Body = v
	return s
}

type ListLoadBalancersRequest struct {
	AddressIpVersion           *string                        `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	AddressType                *string                        `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	DNSName                    *string                        `json:"DNSName,omitempty" xml:"DNSName,omitempty"`
	Ipv6AddressType            *string                        `json:"Ipv6AddressType,omitempty" xml:"Ipv6AddressType,omitempty"`
	LoadBalancerBusinessStatus *string                        `json:"LoadBalancerBusinessStatus,omitempty" xml:"LoadBalancerBusinessStatus,omitempty"`
	LoadBalancerIds            []*string                      `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	LoadBalancerNames          []*string                      `json:"LoadBalancerNames,omitempty" xml:"LoadBalancerNames,omitempty" type:"Repeated"`
	LoadBalancerStatus         *string                        `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	LoadBalancerType           *string                        `json:"LoadBalancerType,omitempty" xml:"LoadBalancerType,omitempty"`
	MaxResults                 *int32                         `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                  *string                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId                   *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId            *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag                        []*ListLoadBalancersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VpcIds                     []*string                      `json:"VpcIds,omitempty" xml:"VpcIds,omitempty" type:"Repeated"`
	ZoneId                     *string                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListLoadBalancersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersRequest) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersRequest) SetAddressIpVersion(v string) *ListLoadBalancersRequest {
	s.AddressIpVersion = &v
	return s
}

func (s *ListLoadBalancersRequest) SetAddressType(v string) *ListLoadBalancersRequest {
	s.AddressType = &v
	return s
}

func (s *ListLoadBalancersRequest) SetDNSName(v string) *ListLoadBalancersRequest {
	s.DNSName = &v
	return s
}

func (s *ListLoadBalancersRequest) SetIpv6AddressType(v string) *ListLoadBalancersRequest {
	s.Ipv6AddressType = &v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerBusinessStatus(v string) *ListLoadBalancersRequest {
	s.LoadBalancerBusinessStatus = &v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerIds(v []*string) *ListLoadBalancersRequest {
	s.LoadBalancerIds = v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerNames(v []*string) *ListLoadBalancersRequest {
	s.LoadBalancerNames = v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerStatus(v string) *ListLoadBalancersRequest {
	s.LoadBalancerStatus = &v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerType(v string) *ListLoadBalancersRequest {
	s.LoadBalancerType = &v
	return s
}

func (s *ListLoadBalancersRequest) SetMaxResults(v int32) *ListLoadBalancersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLoadBalancersRequest) SetNextToken(v string) *ListLoadBalancersRequest {
	s.NextToken = &v
	return s
}

func (s *ListLoadBalancersRequest) SetRegionId(v string) *ListLoadBalancersRequest {
	s.RegionId = &v
	return s
}

func (s *ListLoadBalancersRequest) SetResourceGroupId(v string) *ListLoadBalancersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLoadBalancersRequest) SetTag(v []*ListLoadBalancersRequestTag) *ListLoadBalancersRequest {
	s.Tag = v
	return s
}

func (s *ListLoadBalancersRequest) SetVpcIds(v []*string) *ListLoadBalancersRequest {
	s.VpcIds = v
	return s
}

func (s *ListLoadBalancersRequest) SetZoneId(v string) *ListLoadBalancersRequest {
	s.ZoneId = &v
	return s
}

type ListLoadBalancersRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListLoadBalancersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersRequestTag) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersRequestTag) SetKey(v string) *ListLoadBalancersRequestTag {
	s.Key = &v
	return s
}

func (s *ListLoadBalancersRequestTag) SetValue(v string) *ListLoadBalancersRequestTag {
	s.Value = &v
	return s
}

type ListLoadBalancersResponseBody struct {
	LoadBalancers []*ListLoadBalancersResponseBodyLoadBalancers `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Repeated"`
	MaxResults    *int32                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLoadBalancersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBody) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBody) SetLoadBalancers(v []*ListLoadBalancersResponseBodyLoadBalancers) *ListLoadBalancersResponseBody {
	s.LoadBalancers = v
	return s
}

func (s *ListLoadBalancersResponseBody) SetMaxResults(v int32) *ListLoadBalancersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetNextToken(v string) *ListLoadBalancersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetRequestId(v string) *ListLoadBalancersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetTotalCount(v int32) *ListLoadBalancersResponseBody {
	s.TotalCount = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancers struct {
	AddressIpVersion             *string                                                                 `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	AddressType                  *string                                                                 `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	BandwidthPackageId           *string                                                                 `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	CreateTime                   *string                                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CrossZoneEnabled             *bool                                                                   `json:"CrossZoneEnabled,omitempty" xml:"CrossZoneEnabled,omitempty"`
	DNSName                      *string                                                                 `json:"DNSName,omitempty" xml:"DNSName,omitempty"`
	DeletionProtectionConfig     *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig     `json:"DeletionProtectionConfig,omitempty" xml:"DeletionProtectionConfig,omitempty" type:"Struct"`
	Ipv6AddressType              *string                                                                 `json:"Ipv6AddressType,omitempty" xml:"Ipv6AddressType,omitempty"`
	LoadBalancerBillingConfig    *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig    `json:"LoadBalancerBillingConfig,omitempty" xml:"LoadBalancerBillingConfig,omitempty" type:"Struct"`
	LoadBalancerBusinessStatus   *string                                                                 `json:"LoadBalancerBusinessStatus,omitempty" xml:"LoadBalancerBusinessStatus,omitempty"`
	LoadBalancerId               *string                                                                 `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerName             *string                                                                 `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	LoadBalancerStatus           *string                                                                 `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	LoadBalancerType             *string                                                                 `json:"LoadBalancerType,omitempty" xml:"LoadBalancerType,omitempty"`
	ModificationProtectionConfig *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig `json:"ModificationProtectionConfig,omitempty" xml:"ModificationProtectionConfig,omitempty" type:"Struct"`
	OperationLocks               []*ListLoadBalancersResponseBodyLoadBalancersOperationLocks             `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Repeated"`
	RegionId                     *string                                                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId              *string                                                                 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityGroupIds             []*string                                                               `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	Tags                         []*ListLoadBalancersResponseBodyLoadBalancersTags                       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	VpcId                        *string                                                                 `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneMappings                 []*ListLoadBalancersResponseBodyLoadBalancersZoneMappings               `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s ListLoadBalancersResponseBodyLoadBalancers) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancers) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetAddressIpVersion(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.AddressIpVersion = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetAddressType(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.AddressType = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetBandwidthPackageId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.BandwidthPackageId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetCreateTime(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.CreateTime = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetCrossZoneEnabled(v bool) *ListLoadBalancersResponseBodyLoadBalancers {
	s.CrossZoneEnabled = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetDNSName(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.DNSName = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetDeletionProtectionConfig(v *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) *ListLoadBalancersResponseBodyLoadBalancers {
	s.DeletionProtectionConfig = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetIpv6AddressType(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.Ipv6AddressType = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerBillingConfig(v *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerBillingConfig = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerBusinessStatus(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerBusinessStatus = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerName(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerName = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerStatus(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerStatus = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerType(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerType = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetModificationProtectionConfig(v *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) *ListLoadBalancersResponseBodyLoadBalancers {
	s.ModificationProtectionConfig = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetOperationLocks(v []*ListLoadBalancersResponseBodyLoadBalancersOperationLocks) *ListLoadBalancersResponseBodyLoadBalancers {
	s.OperationLocks = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetRegionId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.RegionId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetResourceGroupId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetSecurityGroupIds(v []*string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.SecurityGroupIds = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetTags(v []*ListLoadBalancersResponseBodyLoadBalancersTags) *ListLoadBalancersResponseBodyLoadBalancers {
	s.Tags = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetVpcId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.VpcId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetZoneMappings(v []*ListLoadBalancersResponseBodyLoadBalancersZoneMappings) *ListLoadBalancersResponseBodyLoadBalancers {
	s.ZoneMappings = v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig struct {
	Enabled     *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	EnabledTime *string `json:"EnabledTime,omitempty" xml:"EnabledTime,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) SetEnabled(v bool) *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig {
	s.Enabled = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) SetEnabledTime(v string) *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig {
	s.EnabledTime = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) SetReason(v string) *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig {
	s.Reason = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig struct {
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) SetPayType(v string) *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig {
	s.PayType = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig struct {
	EnabledTime *string `json:"EnabledTime,omitempty" xml:"EnabledTime,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) SetEnabledTime(v string) *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig {
	s.EnabledTime = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) SetReason(v string) *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig {
	s.Reason = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) SetStatus(v string) *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig {
	s.Status = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersOperationLocks struct {
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	LockType   *string `json:"LockType,omitempty" xml:"LockType,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersOperationLocks) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersOperationLocks) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersOperationLocks) SetLockReason(v string) *ListLoadBalancersResponseBodyLoadBalancersOperationLocks {
	s.LockReason = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersOperationLocks) SetLockType(v string) *ListLoadBalancersResponseBodyLoadBalancersOperationLocks {
	s.LockType = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersTags) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersTags) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersTags) SetKey(v string) *ListLoadBalancersResponseBodyLoadBalancersTags {
	s.Key = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersTags) SetValue(v string) *ListLoadBalancersResponseBodyLoadBalancersTags {
	s.Value = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersZoneMappings struct {
	LoadBalancerAddresses []*ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses `json:"LoadBalancerAddresses,omitempty" xml:"LoadBalancerAddresses,omitempty" type:"Repeated"`
	Status                *string                                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchId             *string                                                                        `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                *string                                                                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersZoneMappings) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersZoneMappings) SetLoadBalancerAddresses(v []*ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses) *ListLoadBalancersResponseBodyLoadBalancersZoneMappings {
	s.LoadBalancerAddresses = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersZoneMappings) SetStatus(v string) *ListLoadBalancersResponseBodyLoadBalancersZoneMappings {
	s.Status = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersZoneMappings) SetVSwitchId(v string) *ListLoadBalancersResponseBodyLoadBalancersZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersZoneMappings) SetZoneId(v string) *ListLoadBalancersResponseBodyLoadBalancersZoneMappings {
	s.ZoneId = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses struct {
	AllocationId        *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	EniId               *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	Ipv6Address         *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	PrivateIPv4Address  *string `json:"PrivateIPv4Address,omitempty" xml:"PrivateIPv4Address,omitempty"`
	PrivateIPv4HcStatus *string `json:"PrivateIPv4HcStatus,omitempty" xml:"PrivateIPv4HcStatus,omitempty"`
	PrivateIPv6HcStatus *string `json:"PrivateIPv6HcStatus,omitempty" xml:"PrivateIPv6HcStatus,omitempty"`
	PublicIPv4Address   *string `json:"PublicIPv4Address,omitempty" xml:"PublicIPv4Address,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses) SetAllocationId(v string) *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses {
	s.AllocationId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses) SetEniId(v string) *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses {
	s.EniId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses) SetIpv6Address(v string) *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses {
	s.Ipv6Address = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses) SetPrivateIPv4Address(v string) *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses {
	s.PrivateIPv4Address = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses) SetPrivateIPv4HcStatus(v string) *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses {
	s.PrivateIPv4HcStatus = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses) SetPrivateIPv6HcStatus(v string) *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses {
	s.PrivateIPv6HcStatus = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses) SetPublicIPv4Address(v string) *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses {
	s.PublicIPv4Address = &v
	return s
}

type ListLoadBalancersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListLoadBalancersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLoadBalancersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponse) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponse) SetHeaders(v map[string]*string) *ListLoadBalancersResponse {
	s.Headers = v
	return s
}

func (s *ListLoadBalancersResponse) SetStatusCode(v int32) *ListLoadBalancersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLoadBalancersResponse) SetBody(v *ListLoadBalancersResponseBody) *ListLoadBalancersResponse {
	s.Body = v
	return s
}

type ListSecurityPolicyRequest struct {
	// The number of entries to return on each page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If this is your first query or no next query is to be sent, ignore this parameter.
	// *   If a next query is to be sent, set the parameter to the value of NextToken that is returned from the last call.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId     *string                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityPolicyIds   []*string                       `json:"SecurityPolicyIds,omitempty" xml:"SecurityPolicyIds,omitempty" type:"Repeated"`
	SecurityPolicyNames []*string                       `json:"SecurityPolicyNames,omitempty" xml:"SecurityPolicyNames,omitempty" type:"Repeated"`
	Tag                 []*ListSecurityPolicyRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListSecurityPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRequest) SetMaxResults(v int32) *ListSecurityPolicyRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSecurityPolicyRequest) SetNextToken(v string) *ListSecurityPolicyRequest {
	s.NextToken = &v
	return s
}

func (s *ListSecurityPolicyRequest) SetRegionId(v string) *ListSecurityPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ListSecurityPolicyRequest) SetResourceGroupId(v string) *ListSecurityPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSecurityPolicyRequest) SetSecurityPolicyIds(v []*string) *ListSecurityPolicyRequest {
	s.SecurityPolicyIds = v
	return s
}

func (s *ListSecurityPolicyRequest) SetSecurityPolicyNames(v []*string) *ListSecurityPolicyRequest {
	s.SecurityPolicyNames = v
	return s
}

func (s *ListSecurityPolicyRequest) SetTag(v []*ListSecurityPolicyRequestTag) *ListSecurityPolicyRequest {
	s.Tag = v
	return s
}

type ListSecurityPolicyRequestTag struct {
	// The tag keys. You can specify up to 10 tag keys.
	//
	// It can be at most 64 characters in length, and cannot contain `http://` or `https://`. It must not start with `aliyun` or `acs:`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values. You can specify up to 10 tag values.
	//
	// It can be at most 128 characters in length, and cannot contain `http://` or `https://`. It must not start with `aliyun` or `acs:`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSecurityPolicyRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyRequestTag) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRequestTag) SetKey(v string) *ListSecurityPolicyRequestTag {
	s.Key = &v
	return s
}

func (s *ListSecurityPolicyRequestTag) SetValue(v string) *ListSecurityPolicyRequestTag {
	s.Value = &v
	return s
}

type ListSecurityPolicyResponseBody struct {
	// The number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If this is your first query or no next query is to be sent, ignore this parameter.
	// *   If a next query is to be sent, set the parameter to the value of NextToken that is returned from the last call.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of TLS security policies.
	SecurityPolicies []*ListSecurityPolicyResponseBodySecurityPolicies `json:"SecurityPolicies,omitempty" xml:"SecurityPolicies,omitempty" type:"Repeated"`
	// The number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSecurityPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyResponseBody) SetMaxResults(v int32) *ListSecurityPolicyResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSecurityPolicyResponseBody) SetNextToken(v string) *ListSecurityPolicyResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSecurityPolicyResponseBody) SetRequestId(v string) *ListSecurityPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecurityPolicyResponseBody) SetSecurityPolicies(v []*ListSecurityPolicyResponseBodySecurityPolicies) *ListSecurityPolicyResponseBody {
	s.SecurityPolicies = v
	return s
}

func (s *ListSecurityPolicyResponseBody) SetTotalCount(v int32) *ListSecurityPolicyResponseBody {
	s.TotalCount = &v
	return s
}

type ListSecurityPolicyResponseBodySecurityPolicies struct {
	// The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
	//
	// TLS 1.0 and TLS 1.1 support the following cipher suites:
	//
	// *   **ECDHE-ECDSA-AES128-SHA**
	// *   **ECDHE-ECDSA-AES256-SHA**
	// *   **ECDHE-RSA-AES128-SHA**
	// *   **ECDHE-RSA-AES256-SHA**
	// *   **AES128-SHA**
	// *   **AES256-SHA**
	// *   **DES-CBC3-SHA**
	//
	// TLS 1.2 supports the following cipher suites:
	//
	// *   **ECDHE-ECDSA-AES128-SHA**
	// *   **ECDHE-ECDSA-AES256-SHA**
	// *   **ECDHE-RSA-AES128-SHA**
	// *   **ECDHE-RSA-AES256-SHA**
	// *   **AES128-SHA**
	// *   **AES256-SHA**
	// *   **DES-CBC3-SHA**
	// *   **ECDHE-ECDSA-AES128-GCM-SHA256**
	// *   **ECDHE-ECDSA-AES256-GCM-SHA384**
	// *   **ECDHE-ECDSA-AES128-SHA256**
	// *   **ECDHE-ECDSA-AES256-SHA384**
	// *   **ECDHE-RSA-AES128-GCM-SHA256**
	// *   **ECDHE-RSA-AES256-GCM-SHA384**
	// *   **ECDHE-RSA-AES128-SHA256**
	// *   **ECDHE-RSA-AES256-SHA384**
	// *   **AES128-GCM-SHA256**
	// *   **AES256-GCM-SHA384**
	// *   **AES128-SHA256**
	// *   **AES256-SHA256**
	//
	// TLS 1.3 supports the following cipher suites:
	//
	// *   **TLS_AES\_128\_GCM_SHA256**
	// *   **TLS_AES\_256\_GCM_SHA384**
	// *   **TLS_CHACHA20\_POLY1305\_SHA256**
	// *   **TLS_AES\_128\_CCM_SHA256**
	// *   **TLS_AES\_128\_CCM\_8\_SHA256**
	Ciphers *string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The listeners that are associated with the NLB instance.
	RelatedListeners []*ListSecurityPolicyResponseBodySecurityPoliciesRelatedListeners `json:"RelatedListeners,omitempty" xml:"RelatedListeners,omitempty" type:"Repeated"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the TLS security policy.
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The name of the TLS security policy.
	SecurityPolicyName *string `json:"SecurityPolicyName,omitempty" xml:"SecurityPolicyName,omitempty"`
	// The status of the TLS security policy. Valid values:
	//
	// *   **Configuring**: The security policy is being configured.
	// *   **Available**: The security policy is available.
	SecurityPolicyStatus *string `json:"SecurityPolicyStatus,omitempty" xml:"SecurityPolicyStatus,omitempty"`
	// The tags that are added to the NLB instance.
	Tags []*ListSecurityPolicyResponseBodySecurityPoliciesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The supported versions of the TLS protocol. Valid values: **TLSv1.0**, **TLSv1.1**, **TLSv1.2**, and **TLSv1.3**.
	TlsVersion *string `json:"TlsVersion,omitempty" xml:"TlsVersion,omitempty"`
}

func (s ListSecurityPolicyResponseBodySecurityPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyResponseBodySecurityPolicies) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyResponseBodySecurityPolicies) SetCiphers(v string) *ListSecurityPolicyResponseBodySecurityPolicies {
	s.Ciphers = &v
	return s
}

func (s *ListSecurityPolicyResponseBodySecurityPolicies) SetRegionId(v string) *ListSecurityPolicyResponseBodySecurityPolicies {
	s.RegionId = &v
	return s
}

func (s *ListSecurityPolicyResponseBodySecurityPolicies) SetRelatedListeners(v []*ListSecurityPolicyResponseBodySecurityPoliciesRelatedListeners) *ListSecurityPolicyResponseBodySecurityPolicies {
	s.RelatedListeners = v
	return s
}

func (s *ListSecurityPolicyResponseBodySecurityPolicies) SetResourceGroupId(v string) *ListSecurityPolicyResponseBodySecurityPolicies {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSecurityPolicyResponseBodySecurityPolicies) SetSecurityPolicyId(v string) *ListSecurityPolicyResponseBodySecurityPolicies {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListSecurityPolicyResponseBodySecurityPolicies) SetSecurityPolicyName(v string) *ListSecurityPolicyResponseBodySecurityPolicies {
	s.SecurityPolicyName = &v
	return s
}

func (s *ListSecurityPolicyResponseBodySecurityPolicies) SetSecurityPolicyStatus(v string) *ListSecurityPolicyResponseBodySecurityPolicies {
	s.SecurityPolicyStatus = &v
	return s
}

func (s *ListSecurityPolicyResponseBodySecurityPolicies) SetTags(v []*ListSecurityPolicyResponseBodySecurityPoliciesTags) *ListSecurityPolicyResponseBodySecurityPolicies {
	s.Tags = v
	return s
}

func (s *ListSecurityPolicyResponseBodySecurityPolicies) SetTlsVersion(v string) *ListSecurityPolicyResponseBodySecurityPolicies {
	s.TlsVersion = &v
	return s
}

type ListSecurityPolicyResponseBodySecurityPoliciesRelatedListeners struct {
	// The ID of the listener.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The listening port.
	ListenerPort *int64 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The listening protocol. Valid value: **TCPSSL**.
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The ID of the NLB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s ListSecurityPolicyResponseBodySecurityPoliciesRelatedListeners) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyResponseBodySecurityPoliciesRelatedListeners) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyResponseBodySecurityPoliciesRelatedListeners) SetListenerId(v string) *ListSecurityPolicyResponseBodySecurityPoliciesRelatedListeners {
	s.ListenerId = &v
	return s
}

func (s *ListSecurityPolicyResponseBodySecurityPoliciesRelatedListeners) SetListenerPort(v int64) *ListSecurityPolicyResponseBodySecurityPoliciesRelatedListeners {
	s.ListenerPort = &v
	return s
}

func (s *ListSecurityPolicyResponseBodySecurityPoliciesRelatedListeners) SetListenerProtocol(v string) *ListSecurityPolicyResponseBodySecurityPoliciesRelatedListeners {
	s.ListenerProtocol = &v
	return s
}

func (s *ListSecurityPolicyResponseBodySecurityPoliciesRelatedListeners) SetLoadBalancerId(v string) *ListSecurityPolicyResponseBodySecurityPoliciesRelatedListeners {
	s.LoadBalancerId = &v
	return s
}

type ListSecurityPolicyResponseBodySecurityPoliciesTags struct {
	// The tag keys. You can specify up to 10 tag keys.
	//
	// The tag key can be at most 64 characters in length, and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values. You can specify up to 10 tag values.
	//
	// It can be at most 128 characters in length, and cannot contain `http://` or `https://`. It must not start with `aliyun` or `acs:`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSecurityPolicyResponseBodySecurityPoliciesTags) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyResponseBodySecurityPoliciesTags) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyResponseBodySecurityPoliciesTags) SetKey(v string) *ListSecurityPolicyResponseBodySecurityPoliciesTags {
	s.Key = &v
	return s
}

func (s *ListSecurityPolicyResponseBodySecurityPoliciesTags) SetValue(v string) *ListSecurityPolicyResponseBodySecurityPoliciesTags {
	s.Value = &v
	return s
}

type ListSecurityPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSecurityPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSecurityPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyResponse) SetHeaders(v map[string]*string) *ListSecurityPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityPolicyResponse) SetStatusCode(v int32) *ListSecurityPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecurityPolicyResponse) SetBody(v *ListSecurityPolicyResponseBody) *ListSecurityPolicyResponse {
	s.Body = v
	return s
}

type ListServerGroupServersRequest struct {
	// The number of entries to return on each page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If this is your first query or no next query is to be sent, ignore this parameter.
	// *   If a next query is to be sent, set the parameter to the value of NextToken that is returned from the last call.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the server group.
	ServerGroupId *string   `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	ServerIds     []*string `json:"ServerIds,omitempty" xml:"ServerIds,omitempty" type:"Repeated"`
	ServerIps     []*string `json:"ServerIps,omitempty" xml:"ServerIps,omitempty" type:"Repeated"`
}

func (s ListServerGroupServersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersRequest) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersRequest) SetMaxResults(v int32) *ListServerGroupServersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupServersRequest) SetNextToken(v string) *ListServerGroupServersRequest {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupServersRequest) SetServerGroupId(v string) *ListServerGroupServersRequest {
	s.ServerGroupId = &v
	return s
}

func (s *ListServerGroupServersRequest) SetServerIds(v []*string) *ListServerGroupServersRequest {
	s.ServerIds = v
	return s
}

func (s *ListServerGroupServersRequest) SetServerIps(v []*string) *ListServerGroupServersRequest {
	s.ServerIps = v
	return s
}

type ListServerGroupServersResponseBody struct {
	// The number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// *   If this is your first query or no next query is to be sent, ignore this parameter.
	// *   If a next query is to be sent, set the parameter to the value of NextToken that is returned from the last call.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of backend servers.
	Servers []*ListServerGroupServersResponseBodyServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// The number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServerGroupServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponseBody) SetMaxResults(v int32) *ListServerGroupServersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetNextToken(v string) *ListServerGroupServersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetRequestId(v string) *ListServerGroupServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetServers(v []*ListServerGroupServersResponseBodyServers) *ListServerGroupServersResponseBody {
	s.Servers = v
	return s
}

func (s *ListServerGroupServersResponseBody) SetTotalCount(v int32) *ListServerGroupServersResponseBody {
	s.TotalCount = &v
	return s
}

type ListServerGroupServersResponseBodyServers struct {
	// The description of the backend server.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port used by the backend server. Valid values: **1** to **65535**.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The ID of the server.
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the backend server.
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of the backend server. Valid values:
	//
	// *   **Ecs**: an Elastic Compute Service (ECS) instance
	// *   **Eni**: an elastic network interface (ENI)
	// *   **Eci**: an elastic container instance
	// *   **Ip**: an IP address
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// Indicates the status of the backend server. Valid values:
	//
	// *   **Adding**: The backend server is being added.
	// *   **Available**: The backend server is added.
	// *   **Configuring**: The backend server is being configured.
	// *   **Removing**: The backend server is being removed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The weight of the backend server.
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// The zone ID of the server.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListServerGroupServersResponseBodyServers) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersResponseBodyServers) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponseBodyServers) SetDescription(v string) *ListServerGroupServersResponseBodyServers {
	s.Description = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetPort(v int32) *ListServerGroupServersResponseBodyServers {
	s.Port = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerGroupId(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerGroupId = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerId(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerId = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerIp(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerIp = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerType(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerType = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetStatus(v string) *ListServerGroupServersResponseBodyServers {
	s.Status = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetWeight(v int32) *ListServerGroupServersResponseBodyServers {
	s.Weight = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetZoneId(v string) *ListServerGroupServersResponseBodyServers {
	s.ZoneId = &v
	return s
}

type ListServerGroupServersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServerGroupServersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServerGroupServersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersResponse) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponse) SetHeaders(v map[string]*string) *ListServerGroupServersResponse {
	s.Headers = v
	return s
}

func (s *ListServerGroupServersResponse) SetStatusCode(v int32) *ListServerGroupServersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServerGroupServersResponse) SetBody(v *ListServerGroupServersResponseBody) *ListServerGroupServersResponse {
	s.Body = v
	return s
}

type ListServerGroupsRequest struct {
	// The number of entries to return on each page. Valid values: **1** to **100**. Default value: **20**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// *   If this is your first query and no next queries are to be sent, ignore this parameter.
	// *   If a next query is to be sent, set the value to the value of NextToken that is returned from the last call.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the server group belongs.
	ResourceGroupId  *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServerGroupIds   []*string `json:"ServerGroupIds,omitempty" xml:"ServerGroupIds,omitempty" type:"Repeated"`
	ServerGroupNames []*string `json:"ServerGroupNames,omitempty" xml:"ServerGroupNames,omitempty" type:"Repeated"`
	// The type of server group. Valid values:
	//
	// *   **Instance** : allows you to add servers of the **Ecs**, **Ens**, and **Eci** types.
	// *   **Ip**: allows you to add servers by specifying IP addresses.
	ServerGroupType *string                       `json:"ServerGroupType,omitempty" xml:"ServerGroupType,omitempty"`
	Tag             []*ListServerGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the server group belongs.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListServerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListServerGroupsRequest) SetMaxResults(v int32) *ListServerGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupsRequest) SetNextToken(v string) *ListServerGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupsRequest) SetRegionId(v string) *ListServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServerGroupsRequest) SetResourceGroupId(v string) *ListServerGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServerGroupsRequest) SetServerGroupIds(v []*string) *ListServerGroupsRequest {
	s.ServerGroupIds = v
	return s
}

func (s *ListServerGroupsRequest) SetServerGroupNames(v []*string) *ListServerGroupsRequest {
	s.ServerGroupNames = v
	return s
}

func (s *ListServerGroupsRequest) SetServerGroupType(v string) *ListServerGroupsRequest {
	s.ServerGroupType = &v
	return s
}

func (s *ListServerGroupsRequest) SetTag(v []*ListServerGroupsRequestTag) *ListServerGroupsRequest {
	s.Tag = v
	return s
}

func (s *ListServerGroupsRequest) SetVpcId(v string) *ListServerGroupsRequest {
	s.VpcId = &v
	return s
}

type ListServerGroupsRequestTag struct {
	// The tag key. You can specify up to 10 tag keys.
	//
	// The tag key can be up to 64 characters in length, and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify up to 10 tag values.
	//
	// The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServerGroupsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *ListServerGroupsRequestTag) SetKey(v string) *ListServerGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *ListServerGroupsRequestTag) SetValue(v string) *ListServerGroupsRequestTag {
	s.Value = &v
	return s
}

type ListServerGroupsResponseBody struct {
	// The number of entries returned per page. Valid values: **1** to **100**.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// *   If **NextToken** is empty, it indicates that no subsequent query is to be sent.
	// *   If a value of **NextToken** is returned, the value is the token used for the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The server groups.
	ServerGroups []*ListServerGroupsResponseBodyServerGroups `json:"ServerGroups,omitempty" xml:"ServerGroups,omitempty" type:"Repeated"`
	// The number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBody) SetMaxResults(v int32) *ListServerGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetNextToken(v string) *ListServerGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetRequestId(v string) *ListServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetServerGroups(v []*ListServerGroupsResponseBodyServerGroups) *ListServerGroupsResponseBody {
	s.ServerGroups = v
	return s
}

func (s *ListServerGroupsResponseBody) SetTotalCount(v int32) *ListServerGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListServerGroupsResponseBodyServerGroups struct {
	// The protocol version. Valid values:
	//
	// *   **ipv4**: IPv4
	// *   **DualStack**: dual stack
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// The UID of the Alibaba Cloud account.
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// Indicates whether the feature of forwarding requests to all ports is enabled. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	AnyPortEnabled *bool `json:"AnyPortEnabled,omitempty" xml:"AnyPortEnabled,omitempty"`
	// Indicates whether connection draining is enabled. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	ConnectionDrainEnabled *bool `json:"ConnectionDrainEnabled,omitempty" xml:"ConnectionDrainEnabled,omitempty"`
	// The timeout period of connection draining. Unit: seconds.
	//
	// Valid values: **10** to **900**.
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	// The configurations of health checks.
	HealthCheck *ListServerGroupsResponseBodyServerGroupsHealthCheck `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty" type:"Struct"`
	// Indicates whether client IP preservation is enabled. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	//
	// >  Note: If **AddressIPVersion** is set to **ipv4**, the default value is **true**. If **AddressIPVersion** is set to **ipv6**, the only valid value is **false**. **true** will be supported in later version.
	PreserveClientIpEnabled *bool `json:"PreserveClientIpEnabled,omitempty" xml:"PreserveClientIpEnabled,omitempty"`
	// The protocol used to forward requests to the backend servers. Valid values: **TCP**, **UDP**, and **TCPSSL**.
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The NLB instances.
	RelatedLoadBalancerIds []*string `json:"RelatedLoadBalancerIds,omitempty" xml:"RelatedLoadBalancerIds,omitempty" type:"Repeated"`
	// The ID of the resource group to which the server group belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// *   **Wrr**: Backend servers with higher weights receive more requests than backend servers with lower weights.
	// *   **rr**: Requests are forwarded to the backend servers in sequence. sch: Requests are forwarded to the backend servers based on source IP address hashing.
	// *   **sch**: Requests from the same source IP address are forwarded to the same backend server.
	// *   **tch**: Four-element hashing, which specifies consistent hashing that is based on four factors: source IP address, destination IP address, source port, and destination port. Requests that contain the same information based on the four factors are forwarded to the same backend server.
	// *   **qch**: QUIC ID hashing. Requests that contain the same QUIC ID are forwarded to the same backend server.
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The number of server groups associated with the NLB instances.
	ServerCount *int32 `json:"ServerCount,omitempty" xml:"ServerCount,omitempty"`
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The name of the server group.
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// The status of the server group. Valid values:
	//
	// *   **Creating**: The server group is being created.
	// *   **Available**: The server group is available.
	// *   **Configuring**: The server group is being configured.
	ServerGroupStatus *string `json:"ServerGroupStatus,omitempty" xml:"ServerGroupStatus,omitempty"`
	// The type of server group. Valid values:
	//
	// *   **Instance** : allows you to add servers of the **Ecs**, **Ens**, and **Eci** types.
	// *   **Ip**: allows you to add servers by specifying IP addresses.
	ServerGroupType *string `json:"ServerGroupType,omitempty" xml:"ServerGroupType,omitempty"`
	// The tags that are added to the NLB instance.
	Tags []*ListServerGroupsResponseBodyServerGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the VPC to which the server group belongs.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroups) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroups) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroups) SetAddressIPVersion(v string) *ListServerGroupsResponseBodyServerGroups {
	s.AddressIPVersion = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetAliUid(v int64) *ListServerGroupsResponseBodyServerGroups {
	s.AliUid = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetAnyPortEnabled(v bool) *ListServerGroupsResponseBodyServerGroups {
	s.AnyPortEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetConnectionDrainEnabled(v bool) *ListServerGroupsResponseBodyServerGroups {
	s.ConnectionDrainEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetConnectionDrainTimeout(v int32) *ListServerGroupsResponseBodyServerGroups {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetHealthCheck(v *ListServerGroupsResponseBodyServerGroupsHealthCheck) *ListServerGroupsResponseBodyServerGroups {
	s.HealthCheck = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetPreserveClientIpEnabled(v bool) *ListServerGroupsResponseBodyServerGroups {
	s.PreserveClientIpEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetProtocol(v string) *ListServerGroupsResponseBodyServerGroups {
	s.Protocol = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetRegionId(v string) *ListServerGroupsResponseBodyServerGroups {
	s.RegionId = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetRelatedLoadBalancerIds(v []*string) *ListServerGroupsResponseBodyServerGroups {
	s.RelatedLoadBalancerIds = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetResourceGroupId(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetScheduler(v string) *ListServerGroupsResponseBodyServerGroups {
	s.Scheduler = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerCount(v int32) *ListServerGroupsResponseBodyServerGroups {
	s.ServerCount = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupId(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupId = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupName(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupName = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupStatus(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupStatus = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupType(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupType = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetTags(v []*ListServerGroupsResponseBodyServerGroupsTags) *ListServerGroupsResponseBodyServerGroups {
	s.Tags = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetVpcId(v string) *ListServerGroupsResponseBodyServerGroups {
	s.VpcId = &v
	return s
}

type ListServerGroupsResponseBodyServerGroupsHealthCheck struct {
	// The backend port that is used for health checks.
	//
	// Valid values: **0** to **65535**.
	//
	// A value of **0** indicates that the port on a backend server is used for health checks.
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The maximum timeout period of a health check. Unit: seconds.
	//
	// Valid values: **1** to **300**.
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// *   **$SERVER_IP**: the private IP address of a backend server.
	// *   **domain**: a specified domain name. The domain name must be 1 to 80 characters in length, and can contain lowercase letters, digits, hyphens (-), and periods (.).
	//
	// >  This parameter takes effect only if **HealthCheckType** is set to **HTTP**.
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// Indicates whether the health check feature is enabled. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The HTTP status codes returned for health checks. Multiple HTTP status codes are separated by commas (,).
	//
	// Valid values: **http\_2xx**, **http\_3xx**, **http\_4xx**, and **http\_5xx**.
	//
	// >  This parameter takes effect only if **HealthCheckType** is set to **HTTP**.
	HealthCheckHttpCode []*string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty" type:"Repeated"`
	// The interval between two consecutive health checks. Unit: seconds.
	//
	// Valid values: **5** to **50**.
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The protocol that is used for health checks. Valid values: **TCP** and **HTTP**.
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The path to which health check requests are sent.
	//
	// >  This parameter takes effect only if **HealthCheckType** is set to **HTTP**.
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail** to **success**.
	//
	// Valid values: **2** to **10**.
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The HTTP method that is used for health checks. Valid values: **GET** and **HEAD**.
	//
	// >  This parameter takes effect only if **HealthCheckType** is set to **HTTP**.
	HttpCheckMethod *string `json:"HttpCheckMethod,omitempty" xml:"HttpCheckMethod,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success** to **fail**.
	//
	// Valid values: **2** to **10**.
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsHealthCheck) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsHealthCheck) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheck) SetHealthCheckConnectPort(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheck {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheck) SetHealthCheckConnectTimeout(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheck {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheck) SetHealthCheckDomain(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheck {
	s.HealthCheckDomain = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheck) SetHealthCheckEnabled(v bool) *ListServerGroupsResponseBodyServerGroupsHealthCheck {
	s.HealthCheckEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheck) SetHealthCheckHttpCode(v []*string) *ListServerGroupsResponseBodyServerGroupsHealthCheck {
	s.HealthCheckHttpCode = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheck) SetHealthCheckInterval(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheck {
	s.HealthCheckInterval = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheck) SetHealthCheckType(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheck {
	s.HealthCheckType = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheck) SetHealthCheckUrl(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheck {
	s.HealthCheckUrl = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheck) SetHealthyThreshold(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheck {
	s.HealthyThreshold = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheck) SetHttpCheckMethod(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheck {
	s.HttpCheckMethod = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheck) SetUnhealthyThreshold(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheck {
	s.UnhealthyThreshold = &v
	return s
}

type ListServerGroupsResponseBodyServerGroupsTags struct {
	// The tag key. At most 10 tag keys are returned.
	//
	// The tag key can be up to 64 characters in length, and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. At most 10 tag values are returned.
	//
	// The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsTags) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsTags) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsTags) SetKey(v string) *ListServerGroupsResponseBodyServerGroupsTags {
	s.Key = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsTags) SetValue(v string) *ListServerGroupsResponseBodyServerGroupsTags {
	s.Value = &v
	return s
}

type ListServerGroupsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServerGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponse) SetHeaders(v map[string]*string) *ListServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListServerGroupsResponse) SetStatusCode(v int32) *ListServerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServerGroupsResponse) SetBody(v *ListServerGroupsResponseBody) *ListServerGroupsResponse {
	s.Body = v
	return s
}

type ListSystemSecurityPolicyRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListSystemSecurityPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSystemSecurityPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPolicyRequest) SetRegionId(v string) *ListSystemSecurityPolicyRequest {
	s.RegionId = &v
	return s
}

type ListSystemSecurityPolicyResponseBody struct {
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityPolicies []*ListSystemSecurityPolicyResponseBodySecurityPolicies `json:"SecurityPolicies,omitempty" xml:"SecurityPolicies,omitempty" type:"Repeated"`
}

func (s ListSystemSecurityPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSystemSecurityPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPolicyResponseBody) SetRequestId(v string) *ListSystemSecurityPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSystemSecurityPolicyResponseBody) SetSecurityPolicies(v []*ListSystemSecurityPolicyResponseBodySecurityPolicies) *ListSystemSecurityPolicyResponseBody {
	s.SecurityPolicies = v
	return s
}

type ListSystemSecurityPolicyResponseBodySecurityPolicies struct {
	Ciphers            *string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty"`
	SecurityPolicyId   *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	SecurityPolicyName *string `json:"SecurityPolicyName,omitempty" xml:"SecurityPolicyName,omitempty"`
	TlsVersion         *string `json:"TlsVersion,omitempty" xml:"TlsVersion,omitempty"`
}

func (s ListSystemSecurityPolicyResponseBodySecurityPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListSystemSecurityPolicyResponseBodySecurityPolicies) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPolicyResponseBodySecurityPolicies) SetCiphers(v string) *ListSystemSecurityPolicyResponseBodySecurityPolicies {
	s.Ciphers = &v
	return s
}

func (s *ListSystemSecurityPolicyResponseBodySecurityPolicies) SetSecurityPolicyId(v string) *ListSystemSecurityPolicyResponseBodySecurityPolicies {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListSystemSecurityPolicyResponseBodySecurityPolicies) SetSecurityPolicyName(v string) *ListSystemSecurityPolicyResponseBodySecurityPolicies {
	s.SecurityPolicyName = &v
	return s
}

func (s *ListSystemSecurityPolicyResponseBodySecurityPolicies) SetTlsVersion(v string) *ListSystemSecurityPolicyResponseBodySecurityPolicies {
	s.TlsVersion = &v
	return s
}

type ListSystemSecurityPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSystemSecurityPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSystemSecurityPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSystemSecurityPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPolicyResponse) SetHeaders(v map[string]*string) *ListSystemSecurityPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListSystemSecurityPolicyResponse) SetStatusCode(v int32) *ListSystemSecurityPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSystemSecurityPolicyResponse) SetBody(v *ListSystemSecurityPolicyResponseBody) *ListSystemSecurityPolicyResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	MaxResults   *int32                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetMaxResults(v int32) *ListTagResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	MaxResults   *int32                                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
	TotalCount   *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetMaxResults(v int32) *ListTagResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesResponseBody) SetTotalCount(v int32) *ListTagResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	AliUid       *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Category     *string `json:"Category,omitempty" xml:"Category,omitempty"`
	RegionNo     *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Scope        *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetAliUid(v int64) *ListTagResourcesResponseBodyTagResources {
	s.AliUid = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetCategory(v string) *ListTagResourcesResponseBodyTagResources {
	s.Category = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetRegionNo(v string) *ListTagResourcesResponseBodyTagResources {
	s.RegionNo = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetScope(v string) *ListTagResourcesResponseBodyTagResources {
	s.Scope = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type LoadBalancerJoinSecurityGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system automatically uses the value of **RequestId** as the value of **ClientToken**. The value of **RequestId** may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// *   **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error code is returned based on the cause of the failure. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the NLB instance to be associated with the security group.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s LoadBalancerJoinSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s LoadBalancerJoinSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *LoadBalancerJoinSecurityGroupRequest) SetClientToken(v string) *LoadBalancerJoinSecurityGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *LoadBalancerJoinSecurityGroupRequest) SetDryRun(v bool) *LoadBalancerJoinSecurityGroupRequest {
	s.DryRun = &v
	return s
}

func (s *LoadBalancerJoinSecurityGroupRequest) SetLoadBalancerId(v string) *LoadBalancerJoinSecurityGroupRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *LoadBalancerJoinSecurityGroupRequest) SetRegionId(v string) *LoadBalancerJoinSecurityGroupRequest {
	s.RegionId = &v
	return s
}

func (s *LoadBalancerJoinSecurityGroupRequest) SetSecurityGroupIds(v []*string) *LoadBalancerJoinSecurityGroupRequest {
	s.SecurityGroupIds = v
	return s
}

type LoadBalancerJoinSecurityGroupResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LoadBalancerJoinSecurityGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoadBalancerJoinSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *LoadBalancerJoinSecurityGroupResponseBody) SetJobId(v string) *LoadBalancerJoinSecurityGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *LoadBalancerJoinSecurityGroupResponseBody) SetRequestId(v string) *LoadBalancerJoinSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

type LoadBalancerJoinSecurityGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LoadBalancerJoinSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LoadBalancerJoinSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s LoadBalancerJoinSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *LoadBalancerJoinSecurityGroupResponse) SetHeaders(v map[string]*string) *LoadBalancerJoinSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *LoadBalancerJoinSecurityGroupResponse) SetStatusCode(v int32) *LoadBalancerJoinSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *LoadBalancerJoinSecurityGroupResponse) SetBody(v *LoadBalancerJoinSecurityGroupResponseBody) *LoadBalancerJoinSecurityGroupResponse {
	s.Body = v
	return s
}

type LoadBalancerLeaveSecurityGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system automatically uses the value of **RequestId** as the value of **ClientToken**. The value of **RequestId** may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// *   **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error code is returned based on the cause of the failure. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the NLB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s LoadBalancerLeaveSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s LoadBalancerLeaveSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *LoadBalancerLeaveSecurityGroupRequest) SetClientToken(v string) *LoadBalancerLeaveSecurityGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *LoadBalancerLeaveSecurityGroupRequest) SetDryRun(v bool) *LoadBalancerLeaveSecurityGroupRequest {
	s.DryRun = &v
	return s
}

func (s *LoadBalancerLeaveSecurityGroupRequest) SetLoadBalancerId(v string) *LoadBalancerLeaveSecurityGroupRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *LoadBalancerLeaveSecurityGroupRequest) SetRegionId(v string) *LoadBalancerLeaveSecurityGroupRequest {
	s.RegionId = &v
	return s
}

func (s *LoadBalancerLeaveSecurityGroupRequest) SetSecurityGroupIds(v []*string) *LoadBalancerLeaveSecurityGroupRequest {
	s.SecurityGroupIds = v
	return s
}

type LoadBalancerLeaveSecurityGroupResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LoadBalancerLeaveSecurityGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoadBalancerLeaveSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *LoadBalancerLeaveSecurityGroupResponseBody) SetJobId(v string) *LoadBalancerLeaveSecurityGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *LoadBalancerLeaveSecurityGroupResponseBody) SetRequestId(v string) *LoadBalancerLeaveSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

type LoadBalancerLeaveSecurityGroupResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LoadBalancerLeaveSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LoadBalancerLeaveSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s LoadBalancerLeaveSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *LoadBalancerLeaveSecurityGroupResponse) SetHeaders(v map[string]*string) *LoadBalancerLeaveSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *LoadBalancerLeaveSecurityGroupResponse) SetStatusCode(v int32) *LoadBalancerLeaveSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *LoadBalancerLeaveSecurityGroupResponse) SetBody(v *LoadBalancerLeaveSecurityGroupResponseBody) *LoadBalancerLeaveSecurityGroupResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId         *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetNewResourceGroupId(v string) *MoveResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetRegionId(v string) *MoveResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type MoveResourceGroupResponseBody struct {
	Data           *MoveResourceGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) SetData(v *MoveResourceGroupResponseBodyData) *MoveResourceGroupResponseBody {
	s.Data = v
	return s
}

func (s *MoveResourceGroupResponseBody) SetHttpStatusCode(v int32) *MoveResourceGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveResourceGroupResponseBody) SetSuccess(v bool) *MoveResourceGroupResponseBody {
	s.Success = &v
	return s
}

type MoveResourceGroupResponseBodyData struct {
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s MoveResourceGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBodyData) SetResourceId(v string) *MoveResourceGroupResponseBodyData {
	s.ResourceId = &v
	return s
}

type MoveResourceGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponse) SetHeaders(v map[string]*string) *MoveResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceGroupResponse) SetStatusCode(v int32) *MoveResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type RemoveServersFromServerGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** of each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck the request. Valid values:
	//
	// *   **true**: prechecks the request but does not remove the backend servers. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the server group.
	ServerGroupId *string                                       `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	Servers       []*RemoveServersFromServerGroupRequestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
}

func (s RemoveServersFromServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupRequest) SetClientToken(v string) *RemoveServersFromServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetDryRun(v bool) *RemoveServersFromServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetRegionId(v string) *RemoveServersFromServerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetServerGroupId(v string) *RemoveServersFromServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetServers(v []*RemoveServersFromServerGroupRequestServers) *RemoveServersFromServerGroupRequest {
	s.Servers = v
	return s
}

type RemoveServersFromServerGroupRequestServers struct {
	// The port used by the backend server. Valid values: **1** to **65535**.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the server.
	//
	// *   If the server group type is **Instance**, set the ServerId parameter to the ID of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance. These backend servers are specified by **Ecs**, **Eni**, or **Eci**.
	// *   If the server group type is **Ip**, set the ServerId parameter to an IP address.
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the server. If the server group type is **Ip**, set the ServerId parameter to an IP address.
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of the backend server. Valid values:
	//
	// *   **Ecs**: an ECS instance
	// *   **Eni**: an ENI
	// *   **Eci**: an elastic container instance
	// *   **Ip**: an IP address
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
}

func (s RemoveServersFromServerGroupRequestServers) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupRequestServers) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupRequestServers) SetPort(v int32) *RemoveServersFromServerGroupRequestServers {
	s.Port = &v
	return s
}

func (s *RemoveServersFromServerGroupRequestServers) SetServerId(v string) *RemoveServersFromServerGroupRequestServers {
	s.ServerId = &v
	return s
}

func (s *RemoveServersFromServerGroupRequestServers) SetServerIp(v string) *RemoveServersFromServerGroupRequestServers {
	s.ServerIp = &v
	return s
}

func (s *RemoveServersFromServerGroupRequestServers) SetServerType(v string) *RemoveServersFromServerGroupRequestServers {
	s.ServerType = &v
	return s
}

type RemoveServersFromServerGroupResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s RemoveServersFromServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupResponseBody) SetJobId(v string) *RemoveServersFromServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *RemoveServersFromServerGroupResponseBody) SetRequestId(v string) *RemoveServersFromServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveServersFromServerGroupResponseBody) SetServerGroupId(v string) *RemoveServersFromServerGroupResponseBody {
	s.ServerGroupId = &v
	return s
}

type RemoveServersFromServerGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveServersFromServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveServersFromServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupResponse) SetHeaders(v map[string]*string) *RemoveServersFromServerGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveServersFromServerGroupResponse) SetStatusCode(v int32) *RemoveServersFromServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveServersFromServerGroupResponse) SetBody(v *RemoveServersFromServerGroupResponseBody) *RemoveServersFromServerGroupResponse {
	s.Body = v
	return s
}

type StartListenerRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** for each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck the request. Valid values:
	//
	// *   **true**: prechecks the request but does not enable the listener. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the listener.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s StartListenerRequest) GoString() string {
	return s.String()
}

func (s *StartListenerRequest) SetClientToken(v string) *StartListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *StartListenerRequest) SetDryRun(v bool) *StartListenerRequest {
	s.DryRun = &v
	return s
}

func (s *StartListenerRequest) SetListenerId(v string) *StartListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *StartListenerRequest) SetRegionId(v string) *StartListenerRequest {
	s.RegionId = &v
	return s
}

type StartListenerResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartListenerResponseBody) GoString() string {
	return s.String()
}

func (s *StartListenerResponseBody) SetJobId(v string) *StartListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *StartListenerResponseBody) SetRequestId(v string) *StartListenerResponseBody {
	s.RequestId = &v
	return s
}

type StartListenerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s StartListenerResponse) GoString() string {
	return s.String()
}

func (s *StartListenerResponse) SetHeaders(v map[string]*string) *StartListenerResponse {
	s.Headers = v
	return s
}

func (s *StartListenerResponse) SetStatusCode(v int32) *StartListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *StartListenerResponse) SetBody(v *StartListenerResponseBody) *StartListenerResponse {
	s.Body = v
	return s
}

type StopListenerRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can only contain ASCII characters.
	//
	// >  If you do not set this parameter, the system automatically uses the value of **RequestId** as the value of **ClientToken**. **RequestId** of each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// *   **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the listener.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s StopListenerRequest) GoString() string {
	return s.String()
}

func (s *StopListenerRequest) SetClientToken(v string) *StopListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *StopListenerRequest) SetDryRun(v bool) *StopListenerRequest {
	s.DryRun = &v
	return s
}

func (s *StopListenerRequest) SetListenerId(v string) *StopListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *StopListenerRequest) SetRegionId(v string) *StopListenerRequest {
	s.RegionId = &v
	return s
}

type StopListenerResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopListenerResponseBody) GoString() string {
	return s.String()
}

func (s *StopListenerResponseBody) SetJobId(v string) *StopListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *StopListenerResponseBody) SetRequestId(v string) *StopListenerResponseBody {
	s.RequestId = &v
	return s
}

type StopListenerResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s StopListenerResponse) GoString() string {
	return s.String()
}

func (s *StopListenerResponse) SetHeaders(v map[string]*string) *StopListenerResponse {
	s.Headers = v
	return s
}

func (s *StopListenerResponse) SetStatusCode(v int32) *StopListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *StopListenerResponse) SetBody(v *StopListenerResponseBody) *StopListenerResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system automatically uses the value of **RequestId** as the value of **ClientToken**. The value of **RequestId** is different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// *   **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false**: performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed. This is the default value.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the resource.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource. You can specify up to 50 resource IDs in each call.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Valid values:
	//
	// *   **loadbalancer**: a Network Load Balancer (NLB) instance
	// *   **securitypolicy**: a security policy
	// *   **servergroup**: a server group
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetClientToken(v string) *TagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *TagResourcesRequest) SetDryRun(v bool) *TagResourcesRequest {
	s.DryRun = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The tag key. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// You can add up to 20 tags in each call.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// You can add up to 20 tags in each call.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetJobId(v string) *TagResourcesResponseBody {
	s.JobId = &v
	return s
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the specified resource. Valid values:
	//
	// *   **true**: removes all tags from the specified resource.
	// *   **false**: does not remove all tags from the specified resource. This is the default value.
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system automatically uses the value of **RequestId** as the value of **ClientToken**. The value of **RequestId** is different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// *   **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false**: performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed. This is the default value.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the resource.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource. You can specify up to 50 resource IDs in each call.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource from which you want to remove tags. Valid values:
	//
	// *   **loadbalancer**: a Network Load Balancer (NLB) instance
	// *   **securitypolicy**: a security policy
	// *   **servergroup**: a server group
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag that you want to remove. You can remove up to 20 tags in each call.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetClientToken(v string) *UntagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *UntagResourcesRequest) SetDryRun(v bool) *UntagResourcesRequest {
	s.DryRun = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetJobId(v string) *UntagResourcesResponseBody {
	s.JobId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateListenerAttributeRequest struct {
	// Specifies whether to enable Application-Layer Protocol Negotiation (ALPN). Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	AlpnEnabled *bool `json:"AlpnEnabled,omitempty" xml:"AlpnEnabled,omitempty"`
	// The ALPN policy.
	AlpnPolicy       *string   `json:"AlpnPolicy,omitempty" xml:"AlpnPolicy,omitempty"`
	CaCertificateIds []*string `json:"CaCertificateIds,omitempty" xml:"CaCertificateIds,omitempty" type:"Repeated"`
	// Specifies whether to enable mutual authentication. Valid values:
	//
	// *   **true**: yes
	// *   **false** (default): no
	CaEnabled      *bool     `json:"CaEnabled,omitempty" xml:"CaEnabled,omitempty"`
	CertificateIds []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** of each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The maximum number of connections that can be created per second on the NLB instance. Valid values: **0** to **1000000**. **0** specifies that the number of connections is unlimited.
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// Specifies whether only to precheck the request. Valid values:
	//
	// *   **true**: prechecks the request but does not update the configurations of the listener. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The timeout period of an idle connection. Unit: seconds. Valid values: **1** to **900**.
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// Enter a name for the listener.
	//
	// The description must be 2 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (\_), and hyphens (-).
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The ID of the listener.
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The size of the largest TCP segment. Unit: bytes. Valid values: **0** to **1500**. **0** specifies that the maximum segment size remains unchanged. This parameter is supported only by listeners that use SSL over TCP.
	Mss *int32 `json:"Mss,omitempty" xml:"Mss,omitempty"`
	// Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	ProxyProtocolEnabled *bool `json:"ProxyProtocolEnabled,omitempty" xml:"ProxyProtocolEnabled,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable fine-grained monitoring. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	SecSensorEnabled *bool `json:"SecSensorEnabled,omitempty" xml:"SecSensorEnabled,omitempty"`
	// The ID of the security policy.
	//
	// >  This parameter takes effect only for listeners that use SSL over TCP.
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequest) SetAlpnEnabled(v bool) *UpdateListenerAttributeRequest {
	s.AlpnEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetAlpnPolicy(v string) *UpdateListenerAttributeRequest {
	s.AlpnPolicy = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetCaCertificateIds(v []*string) *UpdateListenerAttributeRequest {
	s.CaCertificateIds = v
	return s
}

func (s *UpdateListenerAttributeRequest) SetCaEnabled(v bool) *UpdateListenerAttributeRequest {
	s.CaEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetCertificateIds(v []*string) *UpdateListenerAttributeRequest {
	s.CertificateIds = v
	return s
}

func (s *UpdateListenerAttributeRequest) SetClientToken(v string) *UpdateListenerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetCps(v int32) *UpdateListenerAttributeRequest {
	s.Cps = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetDryRun(v bool) *UpdateListenerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetIdleTimeout(v int32) *UpdateListenerAttributeRequest {
	s.IdleTimeout = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetListenerDescription(v string) *UpdateListenerAttributeRequest {
	s.ListenerDescription = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetListenerId(v string) *UpdateListenerAttributeRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetMss(v int32) *UpdateListenerAttributeRequest {
	s.Mss = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetProxyProtocolEnabled(v bool) *UpdateListenerAttributeRequest {
	s.ProxyProtocolEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetRegionId(v string) *UpdateListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetSecSensorEnabled(v bool) *UpdateListenerAttributeRequest {
	s.SecSensorEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetSecurityPolicyId(v string) *UpdateListenerAttributeRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetServerGroupId(v string) *UpdateListenerAttributeRequest {
	s.ServerGroupId = &v
	return s
}

type UpdateListenerAttributeResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeResponseBody) SetJobId(v string) *UpdateListenerAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateListenerAttributeResponseBody) SetRequestId(v string) *UpdateListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateListenerAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateListenerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeResponse) SetHeaders(v map[string]*string) *UpdateListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateListenerAttributeResponse) SetStatusCode(v int32) *UpdateListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateListenerAttributeResponse) SetBody(v *UpdateListenerAttributeResponseBody) *UpdateListenerAttributeResponse {
	s.Body = v
	return s
}

type UpdateLoadBalancerAddressTypeConfigRequest struct {
	// The new network type. Valid values:
	//
	// *   **Internet**: The NLB instance uses a public IP address. The domain name of the NLB instance is resolved to the public IP address. Therefore, the NLB instance can be accessed over the Internet.
	// *   **Intranet**: The NLB instance uses a private IP address. The domain name of the NLB instance is resolved to the private IP address. Therefore, the NLB instance can be accessed over the virtual private cloud (VPC) where the NLB instance is deployed.
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. The token can only contain ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically sets **ClientToken** to the value of **RequestId**. **RequestId** of each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// *   **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the NLB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId     *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneMappings []*UpdateLoadBalancerAddressTypeConfigRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s UpdateLoadBalancerAddressTypeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAddressTypeConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetAddressType(v string) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.AddressType = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetClientToken(v string) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetDryRun(v bool) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetRegionId(v string) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetZoneMappings(v []*UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.ZoneMappings = v
	return s
}

type UpdateLoadBalancerAddressTypeConfigRequestZoneMappings struct {
	// The ID of the elastic IP address (EIP).
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The type of the EIP. Valid values:
	//
	// *   **Common**: EIP
	// *   **Anycast**: Anycast EIP
	//
	// >  Only NLB instances in the China (Hong Kong) region can be associated with Anycast EIPs. This parameter is required if you set the **AddressType** parameter to **Internet**.
	EipType *string `json:"EipType,omitempty" xml:"EipType,omitempty"`
	// The ID of the vSwitch in the zone. You can specify only one vSwitch (subnet) in each zone of an NLB instance.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone of the NLB instance.
	//
	// You can call the [DescribeZones](~~443890~~) operation to query the most recent zone list.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) SetAllocationId(v string) *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings {
	s.AllocationId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) SetEipType(v string) *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings {
	s.EipType = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) SetVSwitchId(v string) *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) SetZoneId(v string) *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings {
	s.ZoneId = &v
	return s
}

type UpdateLoadBalancerAddressTypeConfigResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerAddressTypeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAddressTypeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAddressTypeConfigResponseBody) SetJobId(v string) *UpdateLoadBalancerAddressTypeConfigResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigResponseBody) SetRequestId(v string) *UpdateLoadBalancerAddressTypeConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLoadBalancerAddressTypeConfigResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateLoadBalancerAddressTypeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLoadBalancerAddressTypeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAddressTypeConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAddressTypeConfigResponse) SetHeaders(v map[string]*string) *UpdateLoadBalancerAddressTypeConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigResponse) SetStatusCode(v int32) *UpdateLoadBalancerAddressTypeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigResponse) SetBody(v *UpdateLoadBalancerAddressTypeConfigResponseBody) *UpdateLoadBalancerAddressTypeConfigResponse {
	s.Body = v
	return s
}

type UpdateLoadBalancerAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** of each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The maximum number of connections that can be created per second on the NLB instance. Valid values: **1** to **1000000**.
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// Specifies whether to enable cross-zone load balancing for the NLB instance. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	CrossZoneEnabled *bool `json:"CrossZoneEnabled,omitempty" xml:"CrossZoneEnabled,omitempty"`
	// Specifies whether only to precheck this request. Valid values:
	//
	// *   **true**: prechecks the request but does not modify the name or status of the NLB instance. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the NLB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the NLB instance.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The name must start with a letter.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateLoadBalancerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeRequest) SetClientToken(v string) *UpdateLoadBalancerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetCps(v int32) *UpdateLoadBalancerAttributeRequest {
	s.Cps = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetCrossZoneEnabled(v bool) *UpdateLoadBalancerAttributeRequest {
	s.CrossZoneEnabled = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetDryRun(v bool) *UpdateLoadBalancerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetLoadBalancerName(v string) *UpdateLoadBalancerAttributeRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetRegionId(v string) *UpdateLoadBalancerAttributeRequest {
	s.RegionId = &v
	return s
}

type UpdateLoadBalancerAttributeResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetJobId(v string) *UpdateLoadBalancerAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetRequestId(v string) *UpdateLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLoadBalancerAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateLoadBalancerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLoadBalancerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeResponse) SetHeaders(v map[string]*string) *UpdateLoadBalancerAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoadBalancerAttributeResponse) SetStatusCode(v int32) *UpdateLoadBalancerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLoadBalancerAttributeResponse) SetBody(v *UpdateLoadBalancerAttributeResponseBody) *UpdateLoadBalancerAttributeResponse {
	s.Body = v
	return s
}

type UpdateLoadBalancerProtectionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID** as the **client token**. The **request ID** may be different for each request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable deletion protection. Valid values:
	//
	// *   **true**
	// *   **false**
	DeletionProtectionEnabled *bool `json:"DeletionProtectionEnabled,omitempty" xml:"DeletionProtectionEnabled,omitempty"`
	// The reason for enabling deletion protection. The reason must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The reason must start with a letter.
	//
	// >  This parameter is valid only if you set **DeletionProtectionEnabled** to **true**.
	DeletionProtectionReason *string `json:"DeletionProtectionReason,omitempty" xml:"DeletionProtectionReason,omitempty"`
	// Specifies whether to perform only a dry run without performing the actual request. Valid values:
	//
	// *   **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the NLB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The reason for enabling the configuration read-only mode. The reason must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The reason must start with a letter.
	//
	// >  This parameter takes effect only if you set **Status** to **ConsoleProtection**.
	ModificationProtectionReason *string `json:"ModificationProtectionReason,omitempty" xml:"ModificationProtectionReason,omitempty"`
	// Specifies whether to enable the configuration read-only mode. Valid values:
	//
	// *   **NonProtection**: disables the configuration read-only mode. In this case, you cannot specify **ModificationProtectionReason**. If you specify **ModificationProtectionReason**, the value is cleared.
	// *   **ConsoleProtection**: enables the configuration read-only mode. In this case, you can specify **ModificationProtectionReason**.
	//
	// >  If you set this parameter to **ConsoleProtection**, you cannot modify instance configurations in the NLB console. However, you can modify instance configurations by calling API operations.
	ModificationProtectionStatus *string `json:"ModificationProtectionStatus,omitempty" xml:"ModificationProtectionStatus,omitempty"`
	// The region ID of the NLB instance.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateLoadBalancerProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerProtectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerProtectionRequest) SetClientToken(v string) *UpdateLoadBalancerProtectionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerProtectionRequest) SetDeletionProtectionEnabled(v bool) *UpdateLoadBalancerProtectionRequest {
	s.DeletionProtectionEnabled = &v
	return s
}

func (s *UpdateLoadBalancerProtectionRequest) SetDeletionProtectionReason(v string) *UpdateLoadBalancerProtectionRequest {
	s.DeletionProtectionReason = &v
	return s
}

func (s *UpdateLoadBalancerProtectionRequest) SetDryRun(v bool) *UpdateLoadBalancerProtectionRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerProtectionRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerProtectionRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerProtectionRequest) SetModificationProtectionReason(v string) *UpdateLoadBalancerProtectionRequest {
	s.ModificationProtectionReason = &v
	return s
}

func (s *UpdateLoadBalancerProtectionRequest) SetModificationProtectionStatus(v string) *UpdateLoadBalancerProtectionRequest {
	s.ModificationProtectionStatus = &v
	return s
}

func (s *UpdateLoadBalancerProtectionRequest) SetRegionId(v string) *UpdateLoadBalancerProtectionRequest {
	s.RegionId = &v
	return s
}

type UpdateLoadBalancerProtectionResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerProtectionResponseBody) SetRequestId(v string) *UpdateLoadBalancerProtectionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLoadBalancerProtectionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateLoadBalancerProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLoadBalancerProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerProtectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerProtectionResponse) SetHeaders(v map[string]*string) *UpdateLoadBalancerProtectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoadBalancerProtectionResponse) SetStatusCode(v int32) *UpdateLoadBalancerProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLoadBalancerProtectionResponse) SetBody(v *UpdateLoadBalancerProtectionResponseBody) *UpdateLoadBalancerProtectionResponse {
	s.Body = v
	return s
}

type UpdateLoadBalancerZonesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. The token can only contain ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically sets **ClientToken** to the value of **RequestId**. **RequestId** of each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// *   **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false**: sends the request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed. This is the default value.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the NLB instance.
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId     *string                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneMappings []*UpdateLoadBalancerZonesRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s UpdateLoadBalancerZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerZonesRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesRequest) SetClientToken(v string) *UpdateLoadBalancerZonesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) SetDryRun(v bool) *UpdateLoadBalancerZonesRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerZonesRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) SetRegionId(v string) *UpdateLoadBalancerZonesRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) SetZoneMappings(v []*UpdateLoadBalancerZonesRequestZoneMappings) *UpdateLoadBalancerZonesRequest {
	s.ZoneMappings = v
	return s
}

type UpdateLoadBalancerZonesRequestZoneMappings struct {
	// The ID of the elastic IP address (EIP) or Anycast EIP.
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The type of the EIP. Valid values:
	//
	// *   **Common**: EIP
	// *   **Anycast**: Anycast EIP
	//
	// >  Only NLB instances in the China (Hong Kong) region can be associated with Anycast EIPs. This parameter is required if you set the **AddressType** parameter to **Internet**.
	EipType *string `json:"EipType,omitempty" xml:"EipType,omitempty"`
	// The private IP addresses.
	PrivateIPv4Address *string `json:"PrivateIPv4Address,omitempty" xml:"PrivateIPv4Address,omitempty"`
	// The ID of the vSwitch in the zone. By default, each zone contains one vSwitch and one subnet.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone. You can call the [DescribeZones](~~443890~~) operation to query the zones.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateLoadBalancerZonesRequestZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerZonesRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) SetAllocationId(v string) *UpdateLoadBalancerZonesRequestZoneMappings {
	s.AllocationId = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) SetEipType(v string) *UpdateLoadBalancerZonesRequestZoneMappings {
	s.EipType = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) SetPrivateIPv4Address(v string) *UpdateLoadBalancerZonesRequestZoneMappings {
	s.PrivateIPv4Address = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) SetVSwitchId(v string) *UpdateLoadBalancerZonesRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) SetZoneId(v string) *UpdateLoadBalancerZonesRequestZoneMappings {
	s.ZoneId = &v
	return s
}

type UpdateLoadBalancerZonesResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerZonesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesResponseBody) SetJobId(v string) *UpdateLoadBalancerZonesResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateLoadBalancerZonesResponseBody) SetRequestId(v string) *UpdateLoadBalancerZonesResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLoadBalancerZonesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateLoadBalancerZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLoadBalancerZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerZonesResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesResponse) SetHeaders(v map[string]*string) *UpdateLoadBalancerZonesResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoadBalancerZonesResponse) SetStatusCode(v int32) *UpdateLoadBalancerZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLoadBalancerZonesResponse) SetBody(v *UpdateLoadBalancerZonesResponseBody) *UpdateLoadBalancerZonesResponse {
	s.Body = v
	return s
}

type UpdateSecurityPolicyAttributeRequest struct {
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken** is set to the value of **RequestId**. The value of **RequestId** for each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck the request. Valid values:
	//
	// *   **true**: prechecks the request but does not modify the configurations of the security policy. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the TLS security policy.
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The name of the security policy.
	//
	// The name must be 1 to 200 characters in length, and can contain letters, digits, periods (.), underscores (\_), and hyphens (-).
	SecurityPolicyName *string   `json:"SecurityPolicyName,omitempty" xml:"SecurityPolicyName,omitempty"`
	TlsVersions        []*string `json:"TlsVersions,omitempty" xml:"TlsVersions,omitempty" type:"Repeated"`
}

func (s UpdateSecurityPolicyAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecurityPolicyAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecurityPolicyAttributeRequest) SetCiphers(v []*string) *UpdateSecurityPolicyAttributeRequest {
	s.Ciphers = v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetClientToken(v string) *UpdateSecurityPolicyAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetDryRun(v bool) *UpdateSecurityPolicyAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetRegionId(v string) *UpdateSecurityPolicyAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetSecurityPolicyId(v string) *UpdateSecurityPolicyAttributeRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetSecurityPolicyName(v string) *UpdateSecurityPolicyAttributeRequest {
	s.SecurityPolicyName = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetTlsVersions(v []*string) *UpdateSecurityPolicyAttributeRequest {
	s.TlsVersions = v
	return s
}

type UpdateSecurityPolicyAttributeResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the TLS security policy.
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
}

func (s UpdateSecurityPolicyAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecurityPolicyAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecurityPolicyAttributeResponseBody) SetJobId(v string) *UpdateSecurityPolicyAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponseBody) SetRequestId(v string) *UpdateSecurityPolicyAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponseBody) SetSecurityPolicyId(v string) *UpdateSecurityPolicyAttributeResponseBody {
	s.SecurityPolicyId = &v
	return s
}

type UpdateSecurityPolicyAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSecurityPolicyAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSecurityPolicyAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecurityPolicyAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecurityPolicyAttributeResponse) SetHeaders(v map[string]*string) *UpdateSecurityPolicyAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponse) SetStatusCode(v int32) *UpdateSecurityPolicyAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponse) SetBody(v *UpdateSecurityPolicyAttributeResponseBody) *UpdateSecurityPolicyAttributeResponse {
	s.Body = v
	return s
}

type UpdateServerGroupAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically sets **ClientToken** to the value of **RequestId**. **RequestId** may be different for each API request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable connection draining. Valid values:
	//
	// *   **true**: enables connection draining.
	// *   **false**: disables connection draining.
	ConnectionDrainEnabled *bool `json:"ConnectionDrainEnabled,omitempty" xml:"ConnectionDrainEnabled,omitempty"`
	// The timeout period of connection draining. Unit: seconds. Valid values: **10** to **900**.
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// *   **true**: checks the request without performing the UpdateServerGroupAttribute operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	// *   **false** (default): sends the request. If the request passes the check, a 2xx HTTP status code is returned and the UpdateServerGroupAttribute operation is performed.
	DryRun            *bool                                               `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	HealthCheckConfig *UpdateServerGroupAttributeRequestHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// Specifies whether to enable client IP preservation. Valid values:
	//
	// *   **true**: enables client IP preservation.
	// *   **false**: disables client IP preservation.
	PreserveClientIpEnabled *bool `json:"PreserveClientIpEnabled,omitempty" xml:"PreserveClientIpEnabled,omitempty"`
	// The ID of the region where the NLB instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the available regions.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// *   **Wrr**: Backend servers with higher weights receive more requests than backend servers with lower weights.
	// *   **rr**: Requests are forwarded to backend servers in sequence.
	// *   **sch**: Requests from the same source IP address are forwarded to the same backend server.
	// *   **tch**: Four-element hashing is used. This value specifies consistent hashing that is based on four factors: source IP address, destination IP address, source port, and destination port. Requests that contain the same information based on the four factors are forwarded to the same backend server.
	// *   **qch**: Requests that contain the same QUIC ID are forwarded to the same backend server.
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The new name of the server group.
	//
	// The name must be 2 to 128 characters in length and can contain letters, digits, periods (.), underscores (\_), and hyphens (-). The name must start with a letter.
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
}

func (s UpdateServerGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequest) SetClientToken(v string) *UpdateServerGroupAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetConnectionDrainEnabled(v bool) *UpdateServerGroupAttributeRequest {
	s.ConnectionDrainEnabled = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetConnectionDrainTimeout(v int32) *UpdateServerGroupAttributeRequest {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetDryRun(v bool) *UpdateServerGroupAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetHealthCheckConfig(v *UpdateServerGroupAttributeRequestHealthCheckConfig) *UpdateServerGroupAttributeRequest {
	s.HealthCheckConfig = v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetPreserveClientIpEnabled(v bool) *UpdateServerGroupAttributeRequest {
	s.PreserveClientIpEnabled = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetRegionId(v string) *UpdateServerGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetScheduler(v string) *UpdateServerGroupAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetServerGroupId(v string) *UpdateServerGroupAttributeRequest {
	s.ServerGroupId = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetServerGroupName(v string) *UpdateServerGroupAttributeRequest {
	s.ServerGroupName = &v
	return s
}

type UpdateServerGroupAttributeRequestHealthCheckConfig struct {
	// The backend port that is used for health checks. Valid values: **0** to **65535**. If you set the value to **0**, the ports of backend servers are used for health checks.
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The maximum timeout period of a health check response. Unit: seconds. Valid values: **1** to **300**.
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// *   **$SERVER_IP**: the private IP address of a backend server.
	// *   **domain**: a specific domain name. The domain name must be 1 to 80 characters in length and can contain lowercase letters, digits, hyphens (-), and periods (.).
	//
	// >  This parameter takes effect only when you set **HealthCheckType** to **HTTP**.
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// Specifies whether to enable the health check feature. Valid values:
	//
	// *   **true**: enables the health check feature.
	// *   **false**: disables the health check feature.
	HealthCheckEnabled  *bool     `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckHttpCode []*string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty" type:"Repeated"`
	// The interval between two consecutive health checks. Unit: seconds.
	//
	// Valid values: **5** to **50**.
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The protocol that is used for health checks. Valid values: **TCP** and **HTTP**.
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The path to which health check requests are sent.
	//
	// The path must be 1 to 80 characters in length, and can contain only letters, digits, and the following special characters: `- / . % ? # & =`. The path can also contain the following extended characters: `_ ; ~ ! ( ) * [ ] @ $ ^ : \" , +`. The path must start with a forward slash (/).
	//
	// >  This parameter takes effect only when you set **HealthCheckType** to **HTTP**.
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail** to **success**. Valid values: **2** to **10**.
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The HTTP method that is used for health checks. Valid values: **GET** and **HEAD**.
	//
	// >  This parameter takes effect only when you set **HealthCheckType** to **HTTP**.
	HttpCheckMethod *string `json:"HttpCheckMethod,omitempty" xml:"HttpCheckMethod,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success** to **fail**. Valid values: **2** to **10**.
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s UpdateServerGroupAttributeRequestHealthCheckConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeRequestHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckConnectPort(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckConnectTimeout(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckDomain(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckDomain = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckEnabled(v bool) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckEnabled = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckHttpCode(v []*string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckHttpCode = v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckInterval(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckType(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckType = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckUrl(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckUrl = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthyThreshold(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHttpCheckMethod(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HttpCheckMethod = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetUnhealthyThreshold(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.UnhealthyThreshold = &v
	return s
}

type UpdateServerGroupAttributeResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateServerGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeResponseBody) SetJobId(v string) *UpdateServerGroupAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateServerGroupAttributeResponseBody) SetRequestId(v string) *UpdateServerGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServerGroupAttributeResponseBody) SetServerGroupId(v string) *UpdateServerGroupAttributeResponseBody {
	s.ServerGroupId = &v
	return s
}

type UpdateServerGroupAttributeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateServerGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServerGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeResponse) SetHeaders(v map[string]*string) *UpdateServerGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateServerGroupAttributeResponse) SetStatusCode(v int32) *UpdateServerGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServerGroupAttributeResponse) SetBody(v *UpdateServerGroupAttributeResponseBody) *UpdateServerGroupAttributeResponse {
	s.Body = v
	return s
}

type UpdateServerGroupServersAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can only contain ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically sets **ClientToken** to the value of **RequestId**. **RequestId** of each API request may be different.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// *   **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	// *   **false** (default): performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the Network Load Balancer (NLB) instance is deployed.
	//
	// You can call the [DescribeRegions](~~443657~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the server group.
	ServerGroupId *string                                            `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	Servers       []*UpdateServerGroupServersAttributeRequestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
}

func (s UpdateServerGroupServersAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupServersAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeRequest) SetClientToken(v string) *UpdateServerGroupServersAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) SetDryRun(v bool) *UpdateServerGroupServersAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) SetRegionId(v string) *UpdateServerGroupServersAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) SetServerGroupId(v string) *UpdateServerGroupServersAttributeRequest {
	s.ServerGroupId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) SetServers(v []*UpdateServerGroupServersAttributeRequestServers) *UpdateServerGroupServersAttributeRequest {
	s.Servers = v
	return s
}

type UpdateServerGroupServersAttributeRequestServers struct {
	// The description of the backend server.
	//
	// The description must be 2 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (\_), and hyphens (-).
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port used by the backend server. Valid values: **1** to **65535**. You can specify at most 40 backend servers in each call.
	//
	// >  The value of this parameter cannot be modified.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the backend server. You can specify at most 40 backend servers in each call.
	//
	// *   If the server group type is **Instance**, set the ServerId parameter to the ID of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance. These backend servers are specified by **Ecs**, **Eni**, or **Eci**.
	// *   If the server group type is **Ip**, set the ServerId parameter to an IP address.
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the server. If the server group type is **Ip**, you must specify an IP address.
	//
	// >  You can specify at most 40 backend servers in each call.
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of the backend server. Valid values:
	//
	// *   **Ecs**: an ECS instance
	// *   **Eni**: an ENI
	// *   **Eci**: an elastic container instance
	// *   **Ip**: an IP address
	//
	// >  You can specify at most 40 backend servers in each call.
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The weight of the backend server. Valid values: **0** to **100**. Default value: **100**. If the weight of a backend server is set to **0**, no requests are forwarded to the backend server.
	//
	// >  You can specify at most 40 backend servers in each call.
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateServerGroupServersAttributeRequestServers) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupServersAttributeRequestServers) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetDescription(v string) *UpdateServerGroupServersAttributeRequestServers {
	s.Description = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetPort(v int32) *UpdateServerGroupServersAttributeRequestServers {
	s.Port = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetServerId(v string) *UpdateServerGroupServersAttributeRequestServers {
	s.ServerId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetServerIp(v string) *UpdateServerGroupServersAttributeRequestServers {
	s.ServerIp = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetServerType(v string) *UpdateServerGroupServersAttributeRequestServers {
	s.ServerType = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetWeight(v int32) *UpdateServerGroupServersAttributeRequestServers {
	s.Weight = &v
	return s
}

type UpdateServerGroupServersAttributeResponseBody struct {
	// The ID of the asynchronous task.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the server group.
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateServerGroupServersAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupServersAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeResponseBody) SetJobId(v string) *UpdateServerGroupServersAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeResponseBody) SetRequestId(v string) *UpdateServerGroupServersAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeResponseBody) SetServerGroupId(v string) *UpdateServerGroupServersAttributeResponseBody {
	s.ServerGroupId = &v
	return s
}

type UpdateServerGroupServersAttributeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateServerGroupServersAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServerGroupServersAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupServersAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeResponse) SetHeaders(v map[string]*string) *UpdateServerGroupServersAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateServerGroupServersAttributeResponse) SetStatusCode(v int32) *UpdateServerGroupServersAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServerGroupServersAttributeResponse) SetBody(v *UpdateServerGroupServersAttributeResponseBody) *UpdateServerGroupServersAttributeResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("nlb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddServersToServerGroupWithOptions(request *AddServersToServerGroupRequest, runtime *util.RuntimeOptions) (_result *AddServersToServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Servers)) {
		bodyFlat["Servers"] = request.Servers
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddServersToServerGroup"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddServersToServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddServersToServerGroup(request *AddServersToServerGroupRequest) (_result *AddServersToServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddServersToServerGroupResponse{}
	_body, _err := client.AddServersToServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateAdditionalCertificatesWithListenerWithOptions(request *AssociateAdditionalCertificatesWithListenerRequest, runtime *util.RuntimeOptions) (_result *AssociateAdditionalCertificatesWithListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdditionalCertificateIds)) {
		body["AdditionalCertificateIds"] = request.AdditionalCertificateIds
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		body["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateAdditionalCertificatesWithListener"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateAdditionalCertificatesWithListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateAdditionalCertificatesWithListener(request *AssociateAdditionalCertificatesWithListenerRequest) (_result *AssociateAdditionalCertificatesWithListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateAdditionalCertificatesWithListenerResponse{}
	_body, _err := client.AssociateAdditionalCertificatesWithListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachCommonBandwidthPackageToLoadBalancerWithOptions(request *AttachCommonBandwidthPackageToLoadBalancerRequest, runtime *util.RuntimeOptions) (_result *AttachCommonBandwidthPackageToLoadBalancerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BandwidthPackageId)) {
		body["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachCommonBandwidthPackageToLoadBalancer"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachCommonBandwidthPackageToLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachCommonBandwidthPackageToLoadBalancer(request *AttachCommonBandwidthPackageToLoadBalancerRequest) (_result *AttachCommonBandwidthPackageToLoadBalancerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachCommonBandwidthPackageToLoadBalancerResponse{}
	_body, _err := client.AttachCommonBandwidthPackageToLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateListenerWithOptions(request *CreateListenerRequest, runtime *util.RuntimeOptions) (_result *CreateListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlpnEnabled)) {
		body["AlpnEnabled"] = request.AlpnEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.AlpnPolicy)) {
		body["AlpnPolicy"] = request.AlpnPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.CaCertificateIds)) {
		body["CaCertificateIds"] = request.CaCertificateIds
	}

	if !tea.BoolValue(util.IsUnset(request.CaEnabled)) {
		body["CaEnabled"] = request.CaEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateIds)) {
		body["CertificateIds"] = request.CertificateIds
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Cps)) {
		body["Cps"] = request.Cps
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EndPort)) {
		body["EndPort"] = request.EndPort
	}

	if !tea.BoolValue(util.IsUnset(request.IdleTimeout)) {
		body["IdleTimeout"] = request.IdleTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerDescription)) {
		body["ListenerDescription"] = request.ListenerDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		body["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		body["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.Mss)) {
		body["Mss"] = request.Mss
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyProtocolEnabled)) {
		body["ProxyProtocolEnabled"] = request.ProxyProtocolEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecSensorEnabled)) {
		body["SecSensorEnabled"] = request.SecSensorEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyId)) {
		body["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartPort)) {
		body["StartPort"] = request.StartPort
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateListener"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateListener(request *CreateListenerRequest) (_result *CreateListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateListenerResponse{}
	_body, _err := client.CreateListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   When you create an NLB instance, the service-linked role AliyunServiceRoleForNlb is automatically created and assigned to you.
 * *   **CreateLoadBalancer** is an asynchronous operation. After you send a request, the system returns an instance ID and runs the task in the background. You can call [GetLoadBalancerAttribute](~~445873~~) to query the status of an NLB instance.
 *     *   If an NLB instance is in the **Provisioning** state, the NLB instance is being created.
 *     *   If an NLB instance is in the **Active** state, the NLB instance is created.
 *
 * @param request CreateLoadBalancerRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateLoadBalancerResponse
 */
func (client *Client) CreateLoadBalancerWithOptions(request *CreateLoadBalancerRequest, runtime *util.RuntimeOptions) (_result *CreateLoadBalancerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressIpVersion)) {
		body["AddressIpVersion"] = request.AddressIpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		body["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.BandwidthPackageId)) {
		body["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeletionProtectionConfig)) {
		bodyFlat["DeletionProtectionConfig"] = request.DeletionProtectionConfig
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerBillingConfig)) {
		bodyFlat["LoadBalancerBillingConfig"] = request.LoadBalancerBillingConfig
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerName)) {
		body["LoadBalancerName"] = request.LoadBalancerName
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerType)) {
		body["LoadBalancerType"] = request.LoadBalancerType
	}

	if !tea.BoolValue(util.IsUnset(request.ModificationProtectionConfig)) {
		bodyFlat["ModificationProtectionConfig"] = request.ModificationProtectionConfig
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneMappings)) {
		bodyFlat["ZoneMappings"] = request.ZoneMappings
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLoadBalancer"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   When you create an NLB instance, the service-linked role AliyunServiceRoleForNlb is automatically created and assigned to you.
 * *   **CreateLoadBalancer** is an asynchronous operation. After you send a request, the system returns an instance ID and runs the task in the background. You can call [GetLoadBalancerAttribute](~~445873~~) to query the status of an NLB instance.
 *     *   If an NLB instance is in the **Provisioning** state, the NLB instance is being created.
 *     *   If an NLB instance is in the **Active** state, the NLB instance is created.
 *
 * @param request CreateLoadBalancerRequest
 * @return CreateLoadBalancerResponse
 */
func (client *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (_result *CreateLoadBalancerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLoadBalancerResponse{}
	_body, _err := client.CreateLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSecurityPolicyWithOptions(request *CreateSecurityPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateSecurityPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ciphers)) {
		body["Ciphers"] = request.Ciphers
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyName)) {
		body["SecurityPolicyName"] = request.SecurityPolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TlsVersions)) {
		body["TlsVersions"] = request.TlsVersions
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSecurityPolicy"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSecurityPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSecurityPolicy(request *CreateSecurityPolicyRequest) (_result *CreateSecurityPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSecurityPolicyResponse{}
	_body, _err := client.CreateSecurityPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   **protocol** specifies the protocol used to forward requests to the backend servers.
 * *   NLB instances support only backend server groups that use TCP, UDP, or SSL over TCP.
 * *   **CreateServerGroup** is an asynchronous operation. After you send the request, the system returns a request ID even though the operation is still being performed in the background. You can call the [GetJobStatus](~~445904~~) operation to query the creation status of a server group.
 *     *   If the task is in the **Succeeded** status, the server group is created.
 *     *   If the task is in the **Processing** status, the server group is being created.
 *
 * @param request CreateServerGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateServerGroupResponse
 */
func (client *Client) CreateServerGroupWithOptions(request *CreateServerGroupRequest, runtime *util.RuntimeOptions) (_result *CreateServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressIPVersion)) {
		body["AddressIPVersion"] = request.AddressIPVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AnyPortEnabled)) {
		body["AnyPortEnabled"] = request.AnyPortEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionDrainEnabled)) {
		body["ConnectionDrainEnabled"] = request.ConnectionDrainEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionDrainTimeout)) {
		body["ConnectionDrainTimeout"] = request.ConnectionDrainTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HealthCheckConfig)) {
		bodyFlat["HealthCheckConfig"] = request.HealthCheckConfig
	}

	if !tea.BoolValue(util.IsUnset(request.PreserveClientIpEnabled)) {
		body["PreserveClientIpEnabled"] = request.PreserveClientIpEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		body["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		body["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupName)) {
		body["ServerGroupName"] = request.ServerGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupType)) {
		body["ServerGroupType"] = request.ServerGroupType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServerGroup"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   **protocol** specifies the protocol used to forward requests to the backend servers.
 * *   NLB instances support only backend server groups that use TCP, UDP, or SSL over TCP.
 * *   **CreateServerGroup** is an asynchronous operation. After you send the request, the system returns a request ID even though the operation is still being performed in the background. You can call the [GetJobStatus](~~445904~~) operation to query the creation status of a server group.
 *     *   If the task is in the **Succeeded** status, the server group is created.
 *     *   If the task is in the **Processing** status, the server group is being created.
 *
 * @param request CreateServerGroupRequest
 * @return CreateServerGroupResponse
 */
func (client *Client) CreateServerGroup(request *CreateServerGroupRequest) (_result *CreateServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServerGroupResponse{}
	_body, _err := client.CreateServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteListenerWithOptions(request *DeleteListenerRequest, runtime *util.RuntimeOptions) (_result *DeleteListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		body["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteListener"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteListener(request *DeleteListenerRequest) (_result *DeleteListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteListenerResponse{}
	_body, _err := client.DeleteListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLoadBalancerWithOptions(request *DeleteLoadBalancerRequest, runtime *util.RuntimeOptions) (_result *DeleteLoadBalancerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLoadBalancer"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (_result *DeleteLoadBalancerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLoadBalancerResponse{}
	_body, _err := client.DeleteLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSecurityPolicyWithOptions(request *DeleteSecurityPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteSecurityPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyId)) {
		body["SecurityPolicyId"] = request.SecurityPolicyId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSecurityPolicy"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSecurityPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSecurityPolicy(request *DeleteSecurityPolicyRequest) (_result *DeleteSecurityPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecurityPolicyResponse{}
	_body, _err := client.DeleteSecurityPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can delete server groups that are not associated with listeners.
 *
 * @param request DeleteServerGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteServerGroupResponse
 */
func (client *Client) DeleteServerGroupWithOptions(request *DeleteServerGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServerGroup"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can delete server groups that are not associated with listeners.
 *
 * @param request DeleteServerGroupRequest
 * @return DeleteServerGroupResponse
 */
func (client *Client) DeleteServerGroup(request *DeleteServerGroupRequest) (_result *DeleteServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServerGroupResponse{}
	_body, _err := client.DeleteServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		query["ServiceCode"] = request.ServiceCode
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		query["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeZones"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachCommonBandwidthPackageFromLoadBalancerWithOptions(request *DetachCommonBandwidthPackageFromLoadBalancerRequest, runtime *util.RuntimeOptions) (_result *DetachCommonBandwidthPackageFromLoadBalancerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BandwidthPackageId)) {
		body["BandwidthPackageId"] = request.BandwidthPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachCommonBandwidthPackageFromLoadBalancer"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachCommonBandwidthPackageFromLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachCommonBandwidthPackageFromLoadBalancer(request *DetachCommonBandwidthPackageFromLoadBalancerRequest) (_result *DetachCommonBandwidthPackageFromLoadBalancerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachCommonBandwidthPackageFromLoadBalancerResponse{}
	_body, _err := client.DetachCommonBandwidthPackageFromLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableLoadBalancerIpv6InternetWithOptions(request *DisableLoadBalancerIpv6InternetRequest, runtime *util.RuntimeOptions) (_result *DisableLoadBalancerIpv6InternetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableLoadBalancerIpv6Internet"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableLoadBalancerIpv6InternetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableLoadBalancerIpv6Internet(request *DisableLoadBalancerIpv6InternetRequest) (_result *DisableLoadBalancerIpv6InternetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableLoadBalancerIpv6InternetResponse{}
	_body, _err := client.DisableLoadBalancerIpv6InternetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisassociateAdditionalCertificatesWithListenerWithOptions(request *DisassociateAdditionalCertificatesWithListenerRequest, runtime *util.RuntimeOptions) (_result *DisassociateAdditionalCertificatesWithListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdditionalCertificateIds)) {
		body["AdditionalCertificateIds"] = request.AdditionalCertificateIds
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		body["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisassociateAdditionalCertificatesWithListener"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisassociateAdditionalCertificatesWithListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisassociateAdditionalCertificatesWithListener(request *DisassociateAdditionalCertificatesWithListenerRequest) (_result *DisassociateAdditionalCertificatesWithListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisassociateAdditionalCertificatesWithListenerResponse{}
	_body, _err := client.DisassociateAdditionalCertificatesWithListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableLoadBalancerIpv6InternetWithOptions(request *EnableLoadBalancerIpv6InternetRequest, runtime *util.RuntimeOptions) (_result *EnableLoadBalancerIpv6InternetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableLoadBalancerIpv6Internet"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableLoadBalancerIpv6InternetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableLoadBalancerIpv6Internet(request *EnableLoadBalancerIpv6InternetRequest) (_result *EnableLoadBalancerIpv6InternetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableLoadBalancerIpv6InternetResponse{}
	_body, _err := client.EnableLoadBalancerIpv6InternetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobStatusWithOptions(request *GetJobStatusRequest, runtime *util.RuntimeOptions) (_result *GetJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobStatus"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobStatus(request *GetJobStatusRequest) (_result *GetJobStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobStatusResponse{}
	_body, _err := client.GetJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetListenerAttributeWithOptions(request *GetListenerAttributeRequest, runtime *util.RuntimeOptions) (_result *GetListenerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetListenerAttribute"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetListenerAttribute(request *GetListenerAttributeRequest) (_result *GetListenerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetListenerAttributeResponse{}
	_body, _err := client.GetListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetListenerHealthStatusWithOptions(request *GetListenerHealthStatusRequest, runtime *util.RuntimeOptions) (_result *GetListenerHealthStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetListenerHealthStatus"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetListenerHealthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetListenerHealthStatus(request *GetListenerHealthStatusRequest) (_result *GetListenerHealthStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetListenerHealthStatusResponse{}
	_body, _err := client.GetListenerHealthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLoadBalancerAttributeWithOptions(request *GetLoadBalancerAttributeRequest, runtime *util.RuntimeOptions) (_result *GetLoadBalancerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLoadBalancerAttribute"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLoadBalancerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLoadBalancerAttribute(request *GetLoadBalancerAttributeRequest) (_result *GetLoadBalancerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLoadBalancerAttributeResponse{}
	_body, _err := client.GetLoadBalancerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListListenerCertificatesWithOptions(request *ListListenerCertificatesRequest, runtime *util.RuntimeOptions) (_result *ListListenerCertificatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertType)) {
		body["CertType"] = request.CertType
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		body["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListListenerCertificates"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListListenerCertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListListenerCertificates(request *ListListenerCertificatesRequest) (_result *ListListenerCertificatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListListenerCertificatesResponse{}
	_body, _err := client.ListListenerCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListListenersWithOptions(request *ListListenersRequest, runtime *util.RuntimeOptions) (_result *ListListenersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerIds)) {
		query["ListenerIds"] = request.ListenerIds
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerIds)) {
		query["LoadBalancerIds"] = request.LoadBalancerIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListListeners"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListListenersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListListeners(request *ListListenersRequest) (_result *ListListenersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListListenersResponse{}
	_body, _err := client.ListListenersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLoadBalancersWithOptions(request *ListLoadBalancersRequest, runtime *util.RuntimeOptions) (_result *ListLoadBalancersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressIpVersion)) {
		query["AddressIpVersion"] = request.AddressIpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		query["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.DNSName)) {
		query["DNSName"] = request.DNSName
	}

	if !tea.BoolValue(util.IsUnset(request.Ipv6AddressType)) {
		query["Ipv6AddressType"] = request.Ipv6AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerBusinessStatus)) {
		query["LoadBalancerBusinessStatus"] = request.LoadBalancerBusinessStatus
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerIds)) {
		query["LoadBalancerIds"] = request.LoadBalancerIds
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerNames)) {
		query["LoadBalancerNames"] = request.LoadBalancerNames
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerStatus)) {
		query["LoadBalancerStatus"] = request.LoadBalancerStatus
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerType)) {
		query["LoadBalancerType"] = request.LoadBalancerType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VpcIds)) {
		query["VpcIds"] = request.VpcIds
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLoadBalancers"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLoadBalancersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLoadBalancers(request *ListLoadBalancersRequest) (_result *ListLoadBalancersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLoadBalancersResponse{}
	_body, _err := client.ListLoadBalancersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSecurityPolicyWithOptions(request *ListSecurityPolicyRequest, runtime *util.RuntimeOptions) (_result *ListSecurityPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyIds)) {
		body["SecurityPolicyIds"] = request.SecurityPolicyIds
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyNames)) {
		body["SecurityPolicyNames"] = request.SecurityPolicyNames
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSecurityPolicy"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSecurityPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSecurityPolicy(request *ListSecurityPolicyRequest) (_result *ListSecurityPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSecurityPolicyResponse{}
	_body, _err := client.ListSecurityPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServerGroupServersWithOptions(request *ListServerGroupServersRequest, runtime *util.RuntimeOptions) (_result *ListServerGroupServersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerIds)) {
		body["ServerIds"] = request.ServerIds
	}

	if !tea.BoolValue(util.IsUnset(request.ServerIps)) {
		body["ServerIps"] = request.ServerIps
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServerGroupServers"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServerGroupServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServerGroupServers(request *ListServerGroupServersRequest) (_result *ListServerGroupServersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServerGroupServersResponse{}
	_body, _err := client.ListServerGroupServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServerGroupsWithOptions(request *ListServerGroupsRequest, runtime *util.RuntimeOptions) (_result *ListServerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupIds)) {
		body["ServerGroupIds"] = request.ServerGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupNames)) {
		body["ServerGroupNames"] = request.ServerGroupNames
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupType)) {
		body["ServerGroupType"] = request.ServerGroupType
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServerGroups"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServerGroups(request *ListServerGroupsRequest) (_result *ListServerGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServerGroupsResponse{}
	_body, _err := client.ListServerGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSystemSecurityPolicyWithOptions(request *ListSystemSecurityPolicyRequest, runtime *util.RuntimeOptions) (_result *ListSystemSecurityPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSystemSecurityPolicy"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSystemSecurityPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSystemSecurityPolicy(request *ListSystemSecurityPolicyRequest) (_result *ListSystemSecurityPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSystemSecurityPolicyResponse{}
	_body, _err := client.ListSystemSecurityPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		bodyFlat["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		bodyFlat["Tag"] = request.Tag
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Make sure that you have created a security group. For more information about how to create a security group, see [CreateSecurityGroup](~~25553~~).
 * *   An NLB instance can be associated with up to four security groups.
 * *   You can query the security groups that are associated with an NLB instance by calling the [GetLoadBalancerAttribute](~~214362~~) operation.
 * *   LoadBalancerJoinSecurityGroup is an asynchronous operation. After you call the operation, the system returns a request ID and runs the task in the background. You can call the [GetJobStatus](~~445904~~) operation to query the status of a task.
 *     *   If the task is in the **Succeeded** state, the security group is associated.
 *     *   If the task is in the **Processing** state, the security group is being associated. In this case, you can perform only query operations.
 *
 * @param request LoadBalancerJoinSecurityGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return LoadBalancerJoinSecurityGroupResponse
 */
func (client *Client) LoadBalancerJoinSecurityGroupWithOptions(request *LoadBalancerJoinSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *LoadBalancerJoinSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupIds)) {
		body["SecurityGroupIds"] = request.SecurityGroupIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LoadBalancerJoinSecurityGroup"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LoadBalancerJoinSecurityGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Make sure that you have created a security group. For more information about how to create a security group, see [CreateSecurityGroup](~~25553~~).
 * *   An NLB instance can be associated with up to four security groups.
 * *   You can query the security groups that are associated with an NLB instance by calling the [GetLoadBalancerAttribute](~~214362~~) operation.
 * *   LoadBalancerJoinSecurityGroup is an asynchronous operation. After you call the operation, the system returns a request ID and runs the task in the background. You can call the [GetJobStatus](~~445904~~) operation to query the status of a task.
 *     *   If the task is in the **Succeeded** state, the security group is associated.
 *     *   If the task is in the **Processing** state, the security group is being associated. In this case, you can perform only query operations.
 *
 * @param request LoadBalancerJoinSecurityGroupRequest
 * @return LoadBalancerJoinSecurityGroupResponse
 */
func (client *Client) LoadBalancerJoinSecurityGroup(request *LoadBalancerJoinSecurityGroupRequest) (_result *LoadBalancerJoinSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LoadBalancerJoinSecurityGroupResponse{}
	_body, _err := client.LoadBalancerJoinSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * LoadBalancerLeaveSecurityGroup is an asynchronous operation. After you call the operation, the system returns a request ID and runs the task in the background. You can call the [GetJobStatus](~~445904~~) operation to query the status of a task.
 * *   If the task is in the **Succeeded** state, the security group is disassociated.
 * *   If the task is in the **Processing** state, the security group is being disassociated. In this case, you can perform only query operations.
 *
 * @param request LoadBalancerLeaveSecurityGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return LoadBalancerLeaveSecurityGroupResponse
 */
func (client *Client) LoadBalancerLeaveSecurityGroupWithOptions(request *LoadBalancerLeaveSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *LoadBalancerLeaveSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupIds)) {
		body["SecurityGroupIds"] = request.SecurityGroupIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LoadBalancerLeaveSecurityGroup"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LoadBalancerLeaveSecurityGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * LoadBalancerLeaveSecurityGroup is an asynchronous operation. After you call the operation, the system returns a request ID and runs the task in the background. You can call the [GetJobStatus](~~445904~~) operation to query the status of a task.
 * *   If the task is in the **Succeeded** state, the security group is disassociated.
 * *   If the task is in the **Processing** state, the security group is being disassociated. In this case, you can perform only query operations.
 *
 * @param request LoadBalancerLeaveSecurityGroupRequest
 * @return LoadBalancerLeaveSecurityGroupResponse
 */
func (client *Client) LoadBalancerLeaveSecurityGroup(request *LoadBalancerLeaveSecurityGroupRequest) (_result *LoadBalancerLeaveSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LoadBalancerLeaveSecurityGroupResponse{}
	_body, _err := client.LoadBalancerLeaveSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		body["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResourceGroup"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveServersFromServerGroupWithOptions(request *RemoveServersFromServerGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveServersFromServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Servers)) {
		body["Servers"] = request.Servers
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveServersFromServerGroup"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveServersFromServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveServersFromServerGroup(request *RemoveServersFromServerGroupRequest) (_result *RemoveServersFromServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveServersFromServerGroupResponse{}
	_body, _err := client.RemoveServersFromServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartListenerWithOptions(request *StartListenerRequest, runtime *util.RuntimeOptions) (_result *StartListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		body["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartListener"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartListener(request *StartListenerRequest) (_result *StartListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartListenerResponse{}
	_body, _err := client.StartListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopListenerWithOptions(request *StopListenerRequest, runtime *util.RuntimeOptions) (_result *StopListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		body["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopListener"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopListener(request *StopListenerRequest) (_result *StopListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopListenerResponse{}
	_body, _err := client.StopListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		bodyFlat["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		bodyFlat["Tag"] = request.Tag
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		body["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		bodyFlat["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		bodyFlat["TagKey"] = request.TagKey
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateListenerAttributeWithOptions(request *UpdateListenerAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateListenerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlpnEnabled)) {
		body["AlpnEnabled"] = request.AlpnEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.AlpnPolicy)) {
		body["AlpnPolicy"] = request.AlpnPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.CaCertificateIds)) {
		body["CaCertificateIds"] = request.CaCertificateIds
	}

	if !tea.BoolValue(util.IsUnset(request.CaEnabled)) {
		body["CaEnabled"] = request.CaEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateIds)) {
		body["CertificateIds"] = request.CertificateIds
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Cps)) {
		body["Cps"] = request.Cps
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IdleTimeout)) {
		body["IdleTimeout"] = request.IdleTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerDescription)) {
		body["ListenerDescription"] = request.ListenerDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		body["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.Mss)) {
		body["Mss"] = request.Mss
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyProtocolEnabled)) {
		body["ProxyProtocolEnabled"] = request.ProxyProtocolEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecSensorEnabled)) {
		body["SecSensorEnabled"] = request.SecSensorEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyId)) {
		body["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateListenerAttribute"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateListenerAttribute(request *UpdateListenerAttributeRequest) (_result *UpdateListenerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateListenerAttributeResponse{}
	_body, _err := client.UpdateListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Make sure that an NLB instance is created. For more information, see [CreateLoadBalancer](~~445868~~).
 * *   You can call the [GetLoadBalancerAttribute](~~445873~~) operation to query the **AddressType** value of an NLB instance after you change the network type.
 * *   **UpdateLoadBalancerAddressTypeConfig** is an asynchronous operation. After you send a request, the request ID is returned but the operation is still being performed in the system background. You can call the [GetJobStatus](~~445904~~) operation to query the task status:
 *     *   If the task is in the **Succeeded** state, the network type of the IPv4 address of the NLB instance is changed.
 *     *   If the task is in the **Processing** state, the network type of the IPv4 address of the NLB instance is being changed. In this case, you can perform only query operations.
 *
 * @param request UpdateLoadBalancerAddressTypeConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateLoadBalancerAddressTypeConfigResponse
 */
func (client *Client) UpdateLoadBalancerAddressTypeConfigWithOptions(request *UpdateLoadBalancerAddressTypeConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateLoadBalancerAddressTypeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		body["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneMappings)) {
		body["ZoneMappings"] = request.ZoneMappings
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLoadBalancerAddressTypeConfig"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLoadBalancerAddressTypeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Make sure that an NLB instance is created. For more information, see [CreateLoadBalancer](~~445868~~).
 * *   You can call the [GetLoadBalancerAttribute](~~445873~~) operation to query the **AddressType** value of an NLB instance after you change the network type.
 * *   **UpdateLoadBalancerAddressTypeConfig** is an asynchronous operation. After you send a request, the request ID is returned but the operation is still being performed in the system background. You can call the [GetJobStatus](~~445904~~) operation to query the task status:
 *     *   If the task is in the **Succeeded** state, the network type of the IPv4 address of the NLB instance is changed.
 *     *   If the task is in the **Processing** state, the network type of the IPv4 address of the NLB instance is being changed. In this case, you can perform only query operations.
 *
 * @param request UpdateLoadBalancerAddressTypeConfigRequest
 * @return UpdateLoadBalancerAddressTypeConfigResponse
 */
func (client *Client) UpdateLoadBalancerAddressTypeConfig(request *UpdateLoadBalancerAddressTypeConfigRequest) (_result *UpdateLoadBalancerAddressTypeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLoadBalancerAddressTypeConfigResponse{}
	_body, _err := client.UpdateLoadBalancerAddressTypeConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLoadBalancerAttributeWithOptions(request *UpdateLoadBalancerAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateLoadBalancerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Cps)) {
		body["Cps"] = request.Cps
	}

	if !tea.BoolValue(util.IsUnset(request.CrossZoneEnabled)) {
		body["CrossZoneEnabled"] = request.CrossZoneEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerName)) {
		body["LoadBalancerName"] = request.LoadBalancerName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLoadBalancerAttribute"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLoadBalancerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLoadBalancerAttribute(request *UpdateLoadBalancerAttributeRequest) (_result *UpdateLoadBalancerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLoadBalancerAttributeResponse{}
	_body, _err := client.UpdateLoadBalancerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * >  You can call the [GetLoadBalancerAttribute](~~445873~~) operation to query the details of deletion protection and the configuration read-only mode.
 *
 * @param request UpdateLoadBalancerProtectionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateLoadBalancerProtectionResponse
 */
func (client *Client) UpdateLoadBalancerProtectionWithOptions(request *UpdateLoadBalancerProtectionRequest, runtime *util.RuntimeOptions) (_result *UpdateLoadBalancerProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeletionProtectionEnabled)) {
		body["DeletionProtectionEnabled"] = request.DeletionProtectionEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.DeletionProtectionReason)) {
		body["DeletionProtectionReason"] = request.DeletionProtectionReason
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.ModificationProtectionReason)) {
		body["ModificationProtectionReason"] = request.ModificationProtectionReason
	}

	if !tea.BoolValue(util.IsUnset(request.ModificationProtectionStatus)) {
		body["ModificationProtectionStatus"] = request.ModificationProtectionStatus
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLoadBalancerProtection"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLoadBalancerProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * >  You can call the [GetLoadBalancerAttribute](~~445873~~) operation to query the details of deletion protection and the configuration read-only mode.
 *
 * @param request UpdateLoadBalancerProtectionRequest
 * @return UpdateLoadBalancerProtectionResponse
 */
func (client *Client) UpdateLoadBalancerProtection(request *UpdateLoadBalancerProtectionRequest) (_result *UpdateLoadBalancerProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLoadBalancerProtectionResponse{}
	_body, _err := client.UpdateLoadBalancerProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * - Make sure that an NLB instance is created. For more information, see [CreateLoadBalancer](/help/en/server-load-balancer/latest/createloadbalancer).
 * - You can call the [GetLoadBalancerAttribute](/help/en/server-load-balancer/latest/getloadbalancerattribute-nlb) operation to query the zones and zone attributes of an NLB instance.
 * - **UpdateLoadBalancerZones** is an asynchronous operation. After you send a request, the request ID is returned but the operation is still being performed in the system background. You can call the [GetJobStatus](/help/en/server-load-balancer/latest/getjobstatus) operation to query the status of a task:
 *          - If the task is in the **Succeeded** state, the zones and zone attributes are modified.
 *   - If the task is in the **Processing** state, the zones and zone attributes are being modified. In this case, you can perform only query operations.
 * ## Precautions
 * When you call this operation, make sure that you specify all the zones of the NLB instance, including the existing zones and new zones. If you do not specify the existing zones, the existing zones are removed.
 *
 * @param request UpdateLoadBalancerZonesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateLoadBalancerZonesResponse
 */
func (client *Client) UpdateLoadBalancerZonesWithOptions(request *UpdateLoadBalancerZonesRequest, runtime *util.RuntimeOptions) (_result *UpdateLoadBalancerZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneMappings)) {
		body["ZoneMappings"] = request.ZoneMappings
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLoadBalancerZones"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLoadBalancerZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * - Make sure that an NLB instance is created. For more information, see [CreateLoadBalancer](/help/en/server-load-balancer/latest/createloadbalancer).
 * - You can call the [GetLoadBalancerAttribute](/help/en/server-load-balancer/latest/getloadbalancerattribute-nlb) operation to query the zones and zone attributes of an NLB instance.
 * - **UpdateLoadBalancerZones** is an asynchronous operation. After you send a request, the request ID is returned but the operation is still being performed in the system background. You can call the [GetJobStatus](/help/en/server-load-balancer/latest/getjobstatus) operation to query the status of a task:
 *          - If the task is in the **Succeeded** state, the zones and zone attributes are modified.
 *   - If the task is in the **Processing** state, the zones and zone attributes are being modified. In this case, you can perform only query operations.
 * ## Precautions
 * When you call this operation, make sure that you specify all the zones of the NLB instance, including the existing zones and new zones. If you do not specify the existing zones, the existing zones are removed.
 *
 * @param request UpdateLoadBalancerZonesRequest
 * @return UpdateLoadBalancerZonesResponse
 */
func (client *Client) UpdateLoadBalancerZones(request *UpdateLoadBalancerZonesRequest) (_result *UpdateLoadBalancerZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLoadBalancerZonesResponse{}
	_body, _err := client.UpdateLoadBalancerZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSecurityPolicyAttributeWithOptions(request *UpdateSecurityPolicyAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateSecurityPolicyAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ciphers)) {
		body["Ciphers"] = request.Ciphers
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyId)) {
		body["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyName)) {
		body["SecurityPolicyName"] = request.SecurityPolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.TlsVersions)) {
		body["TlsVersions"] = request.TlsVersions
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSecurityPolicyAttribute"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSecurityPolicyAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSecurityPolicyAttribute(request *UpdateSecurityPolicyAttributeRequest) (_result *UpdateSecurityPolicyAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSecurityPolicyAttributeResponse{}
	_body, _err := client.UpdateSecurityPolicyAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServerGroupAttributeWithOptions(request *UpdateServerGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateServerGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionDrainEnabled)) {
		body["ConnectionDrainEnabled"] = request.ConnectionDrainEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionDrainTimeout)) {
		body["ConnectionDrainTimeout"] = request.ConnectionDrainTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HealthCheckConfig)) {
		bodyFlat["HealthCheckConfig"] = request.HealthCheckConfig
	}

	if !tea.BoolValue(util.IsUnset(request.PreserveClientIpEnabled)) {
		body["PreserveClientIpEnabled"] = request.PreserveClientIpEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		body["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupName)) {
		body["ServerGroupName"] = request.ServerGroupName
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServerGroupAttribute"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServerGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServerGroupAttribute(request *UpdateServerGroupAttributeRequest) (_result *UpdateServerGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServerGroupAttributeResponse{}
	_body, _err := client.UpdateServerGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The **UpdateServerGroupServersAttribute** operation is asynchronous. After you send a request, the system returns the request ID, but the operation is still being performed in the system background.
 * 1\\. You can call [ListServerGroups](~~445895~~) to query the status of a server group.
 * *   If a server group is in the **Configuring** state, it indicates that the server group is being modified.
 * *   If a server group is in the **Available** state, it indicates that the server group is running.
 * 2\\. You can call [ListServerGroupServers](~~445896~~) to query the status of a backend server.
 * *   If a backend server is in the **Configuring** state, it indicates that the backend server is being modified.
 * *   If a backend server is in the **Available** state, it indicates that the backend server is running.
 *
 * @param request UpdateServerGroupServersAttributeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateServerGroupServersAttributeResponse
 */
func (client *Client) UpdateServerGroupServersAttributeWithOptions(request *UpdateServerGroupServersAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateServerGroupServersAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Servers)) {
		body["Servers"] = request.Servers
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServerGroupServersAttribute"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServerGroupServersAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The **UpdateServerGroupServersAttribute** operation is asynchronous. After you send a request, the system returns the request ID, but the operation is still being performed in the system background.
 * 1\\. You can call [ListServerGroups](~~445895~~) to query the status of a server group.
 * *   If a server group is in the **Configuring** state, it indicates that the server group is being modified.
 * *   If a server group is in the **Available** state, it indicates that the server group is running.
 * 2\\. You can call [ListServerGroupServers](~~445896~~) to query the status of a backend server.
 * *   If a backend server is in the **Configuring** state, it indicates that the backend server is being modified.
 * *   If a backend server is in the **Available** state, it indicates that the backend server is running.
 *
 * @param request UpdateServerGroupServersAttributeRequest
 * @return UpdateServerGroupServersAttributeResponse
 */
func (client *Client) UpdateServerGroupServersAttribute(request *UpdateServerGroupServersAttributeRequest) (_result *UpdateServerGroupServersAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServerGroupServersAttributeResponse{}
	_body, _err := client.UpdateServerGroupServersAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
