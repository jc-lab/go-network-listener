//go:build windows
// +build windows

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019-2022 WireGuard LLC. All Rights Reserved.
 */

package iphlpapi

const (
	anySize                  = 1
	maxDNSSuffixStringLength = 256
	maxDHCPv6DUIDLength      = 130
	ifMaxStringSize          = 256
	ifMaxPhysAddressLength   = 32
)

// AddressFamily enumeration specifies protocol family and is one of the windows.AF_* constants.
type AddressFamily uint16

// IPAAFlags enumeration describes adapter addresses flags
// https://docs.microsoft.com/en-us/windows/desktop/api/iptypes/ns-iptypes-_ip_adapter_addresses_lh
type IPAAFlags uint32

const (
	IPAAFlagDdnsEnabled IPAAFlags = 1 << iota
	IPAAFlagRegisterAdapterSuffix
	IPAAFlagDhcpv4Enabled
	IPAAFlagReceiveOnly
	IPAAFlagNoMulticast
	IPAAFlagIpv6OtherStatefulConfig
	IPAAFlagNetbiosOverTcpipEnabled
	IPAAFlagIpv4Enabled
	IPAAFlagIpv6Enabled
	IPAAFlagIpv6ManagedAddressConfigurationSupported
)

// IfOperStatus enumeration specifies the operational status of an interface.
// https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ne-ifdef-if_oper_status
type IfOperStatus uint32

const (
	IfOperStatusUp IfOperStatus = iota + 1
	IfOperStatusDown
	IfOperStatusTesting
	IfOperStatusUnknown
	IfOperStatusDormant
	IfOperStatusNotPresent
	IfOperStatusLowerLayerDown
)

// IfType enumeration specifies interface type.
type IfType uint32

const (
	IfTypeOther                         IfType = 1 // None of the below
	IfTypeRegular1822                          = 2
	IfTypeHdh1822                              = 3
	IfTypeDdnX25                               = 4
	IfTypeRfc877X25                            = 5
	IfTypeEthernetCSMACD                       = 6
	IfTypeISO88023CSMACD                       = 7
	IfTypeISO88024Tokenbus                     = 8
	IfTypeISO88025Tokenring                    = 9
	IfTypeISO88026Man                          = 10
	IfTypeStarlan                              = 11
	IfTypeProteon10Mbit                        = 12
	IfTypeProteon80Mbit                        = 13
	IfTypeHyperchannel                         = 14
	IfTypeFddi                                 = 15
	IfTypeLapB                                 = 16
	IfTypeSdlc                                 = 17
	IfTypeDs1                                  = 18 // DS1-MIB
	IfTypeE1                                   = 19 // Obsolete; see DS1-MIB
	IfTypeBasicISDN                            = 20
	IfTypePrimaryISDN                          = 21
	IfTypePropPoint2PointSerial                = 22 // proprietary serial
	IfTypePPP                                  = 23
	IfTypeSoftwareLoopback                     = 24
	IfTypeEon                                  = 25 // CLNP over IP
	IfTypeEthernet3Mbit                        = 26
	IfTypeNsip                                 = 27 // XNS over IP
	IfTypeSlip                                 = 28 // Generic Slip
	IfTypeUltra                                = 29 // ULTRA Technologies
	IfTypeDs3                                  = 30 // DS3-MIB
	IfTypeSip                                  = 31 // SMDS, coffee
	IfTypeFramerelay                           = 32 // DTE only
	IfTypeRs232                                = 33
	IfTypePara                                 = 34 // Parallel port
	IfTypeArcnet                               = 35
	IfTypeArcnetPlus                           = 36
	IfTypeAtm                                  = 37 // ATM cells
	IfTypeMioX25                               = 38
	IfTypeSonet                                = 39 // SONET or SDH
	IfTypeX25Ple                               = 40
	IfTypeIso88022LLC                          = 41
	IfTypeLocaltalk                            = 42
	IfTypeSmdsDxi                              = 43
	IfTypeFramerelayService                    = 44 // FRNETSERV-MIB
	IfTypeV35                                  = 45
	IfTypeHssi                                 = 46
	IfTypeHippi                                = 47
	IfTypeModem                                = 48 // Generic Modem
	IfTypeAal5                                 = 49 // AAL5 over ATM
	IfTypeSonetPath                            = 50
	IfTypeSonetVt                              = 51
	IfTypeSmdsIcip                             = 52 // SMDS InterCarrier Interface
	IfTypePropVirtual                          = 53 // Proprietary virtual/internal
	IfTypePropMultiplexor                      = 54 // Proprietary multiplexing
	IfTypeIEEE80212                            = 55 // 100BaseVG
	IfTypeFibrechannel                         = 56
	IfTypeHippiinterface                       = 57
	IfTypeFramerelayInterconnect               = 58 // Obsolete, use 32 or 44
	IfTypeAflane8023                           = 59 // ATM Emulated LAN for 802.3
	IfTypeAflane8025                           = 60 // ATM Emulated LAN for 802.5
	IfTypeCctemul                              = 61 // ATM Emulated circuit
	IfTypeFastether                            = 62 // Fast Ethernet (100BaseT)
	IfTypeISDN                                 = 63 // ISDN and X.25
	IfTypeV11                                  = 64 // CCITT V.11/X.21
	IfTypeV36                                  = 65 // CCITT V.36
	IfTypeG703_64k                             = 66 // CCITT G703 at 64Kbps
	IfTypeG703_2mb                             = 67 // Obsolete; see DS1-MIB
	IfTypeQllc                                 = 68 // SNA QLLC
	IfTypeFastetherFX                          = 69 // Fast Ethernet (100BaseFX)
	IfTypeChannel                              = 70
	IfTypeIEEE80211                            = 71  // Radio spread spectrum
	IfTypeIBM370parchan                        = 72  // IBM System 360/370 OEMI Channel
	IfTypeEscon                                = 73  // IBM Enterprise Systems Connection
	IfTypeDlsw                                 = 74  // Data Link Switching
	IfTypeISDNS                                = 75  // ISDN S/T interface
	IfTypeISDNU                                = 76  // ISDN U interface
	IfTypeLapD                                 = 77  // Link Access Protocol D
	IfTypeIpswitch                             = 78  // IP Switching Objects
	IfTypeRsrb                                 = 79  // Remote Source Route Bridging
	IfTypeAtmLogical                           = 80  // ATM Logical Port
	IfTypeDs0                                  = 81  // Digital Signal Level 0
	IfTypeDs0Bundle                            = 82  // Group of ds0s on the same ds1
	IfTypeBsc                                  = 83  // Bisynchronous Protocol
	IfTypeAsync                                = 84  // Asynchronous Protocol
	IfTypeCnr                                  = 85  // Combat Net Radio
	IfTypeIso88025rDtr                         = 86  // ISO 802.5r DTR
	IfTypeEplrs                                = 87  // Ext Pos Loc Report Sys
	IfTypeArap                                 = 88  // Appletalk Remote Access Protocol
	IfTypePropCnls                             = 89  // Proprietary Connectionless Proto
	IfTypeHostpad                              = 90  // CCITT-ITU X.29 PAD Protocol
	IfTypeTermpad                              = 91  // CCITT-ITU X.3 PAD Facility
	IfTypeFramerelayMpi                        = 92  // Multiproto Interconnect over FR
	IfTypeX213                                 = 93  // CCITT-ITU X213
	IfTypeAdsl                                 = 94  // Asymmetric Digital Subscrbr Loop
	IfTypeRadsl                                = 95  // Rate-Adapt Digital Subscrbr Loop
	IfTypeSdsl                                 = 96  // Symmetric Digital Subscriber Loop
	IfTypeVdsl                                 = 97  // Very H-Speed Digital Subscrb Loop
	IfTypeIso88025Crfprint                     = 98  // ISO 802.5 CRFP
	IfTypeMyrinet                              = 99  // Myricom Myrinet
	IfTypeVoiceEm                              = 100 // Voice recEive and transMit
	IfTypeVoiceFxo                             = 101 // Voice Foreign Exchange Office
	IfTypeVoiceFxs                             = 102 // Voice Foreign Exchange Station
	IfTypeVoiceEncap                           = 103 // Voice encapsulation
	IfTypeVoiceOverip                          = 104 // Voice over IP encapsulation
	IfTypeAtmDxi                               = 105 // ATM DXI
	IfTypeAtmFuni                              = 106 // ATM FUNI
	IfTypeAtmIma                               = 107 // ATM IMA
	IfTypePPPmultilinkbundle                   = 108 // PPP Multilink Bundle
	IfTypeIpoverCdlc                           = 109 // IBM ipOverCdlc
	IfTypeIpoverClaw                           = 110 // IBM Common Link Access to Workstn
	IfTypeStacktostack                         = 111 // IBM stackToStack
	IfTypeVirtualipaddress                     = 112 // IBM VIPA
	IfTypeMpc                                  = 113 // IBM multi-proto channel support
	IfTypeIpoverAtm                            = 114 // IBM ipOverAtm
	IfTypeIso88025Fiber                        = 115 // ISO 802.5j Fiber Token Ring
	IfTypeTdlc                                 = 116 // IBM twinaxial data link control
	IfTypeGigabitethernet                      = 117
	IfTypeHdlc                                 = 118
	IfTypeLapF                                 = 119
	IfTypeV37                                  = 120
	IfTypeX25Mlp                               = 121 // Multi-Link Protocol
	IfTypeX25Huntgroup                         = 122 // X.25 Hunt Group
	IfTypeTransphdlc                           = 123
	IfTypeInterleave                           = 124 // Interleave channel
	IfTypeFast                                 = 125 // Fast channel
	IfTypeIP                                   = 126 // IP (for APPN HPR in IP networks)
	IfTypeDocscableMaclayer                    = 127 // CATV Mac Layer
	IfTypeDocscableDownstream                  = 128 // CATV Downstream interface
	IfTypeDocscableUpstream                    = 129 // CATV Upstream interface
	IfTypeA12mppswitch                         = 130 // Avalon Parallel Processor
	IfTypeTunnel                               = 131 // Encapsulation interface
	IfTypeCoffee                               = 132 // Coffee pot
	IfTypeCes                                  = 133 // Circuit Emulation Service
	IfTypeAtmSubinterface                      = 134 // ATM Sub Interface
	IfTypeL2Vlan                               = 135 // Layer 2 Virtual LAN using 802.1Q
	IfTypeL3Ipvlan                             = 136 // Layer 3 Virtual LAN using IP
	IfTypeL3Ipxvlan                            = 137 // Layer 3 Virtual LAN using IPX
	IfTypeDigitalpowerline                     = 138 // IP over Power Lines
	IfTypeMediamailoverip                      = 139 // Multimedia Mail over IP
	IfTypeDtm                                  = 140 // Dynamic syncronous Transfer Mode
	IfTypeDcn                                  = 141 // Data Communications Network
	IfTypeIpforward                            = 142 // IP Forwarding Interface
	IfTypeMsdsl                                = 143 // Multi-rate Symmetric DSL
	IfTypeIEEE1394                             = 144 // IEEE1394 High Perf Serial Bus
	IfTypeIfGsn                                = 145
	IfTypeDvbrccMaclayer                       = 146
	IfTypeDvbrccDownstream                     = 147
	IfTypeDvbrccUpstream                       = 148
	IfTypeAtmVirtual                           = 149
	IfTypeMplsTunnel                           = 150
	IfTypeSrp                                  = 151
	IfTypeVoiceoveratm                         = 152
	IfTypeVoiceoverframerelay                  = 153
	IfTypeIdsl                                 = 154
	IfTypeCompositelink                        = 155
	IfTypeSs7Siglink                           = 156
	IfTypePropWirelessP2P                      = 157
	IfTypeFrForward                            = 158
	IfTypeRfc1483                              = 159
	IfTypeUsb                                  = 160
	IfTypeIEEE8023adLag                        = 161
	IfTypeBgpPolicyAccounting                  = 162
	IfTypeFrf16MfrBundle                       = 163
	IfTypeH323Gatekeeper                       = 164
	IfTypeH323Proxy                            = 165
	IfTypeMpls                                 = 166
	IfTypeMfSiglink                            = 167
	IfTypeHdsl2                                = 168
	IfTypeShdsl                                = 169
	IfTypeDs1Fdl                               = 170
	IfTypePos                                  = 171
	IfTypeDvbAsiIn                             = 172
	IfTypeDvbAsiOut                            = 173
	IfTypePlc                                  = 174
	IfTypeNfas                                 = 175
	IfTypeTr008                                = 176
	IfTypeGr303Rdt                             = 177
	IfTypeGr303Idt                             = 178
	IfTypeIsup                                 = 179
	IfTypePropDocsWirelessMaclayer             = 180
	IfTypePropDocsWirelessDownstream           = 181
	IfTypePropDocsWirelessUpstream             = 182
	IfTypeHiperlan2                            = 183
	IfTypePropBwaP2MP                          = 184
	IfTypeSonetOverheadChannel                 = 185
	IfTypeDigitalWrapperOverheadChannel        = 186
	IfTypeAal2                                 = 187
	IfTypeRadioMac                             = 188
	IfTypeAtmRadio                             = 189
	IfTypeImt                                  = 190
	IfTypeMvl                                  = 191
	IfTypeReachDsl                             = 192
	IfTypeFrDlciEndpt                          = 193
	IfTypeAtmVciEndpt                          = 194
	IfTypeOpticalChannel                       = 195
	IfTypeOpticalTransport                     = 196
	IfTypeIEEE80216Wman                        = 237
	IfTypeWwanpp                               = 243 // WWAN devices based on GSM technology
	IfTypeWwanpp2                              = 244 // WWAN devices based on CDMA technology
	IfTypeIEEE802154                           = 259 // IEEE 802.15.4 WPAN interface
	IfTypeXboxWireless                         = 281
)

// MibIfEntryLevel enumeration specifies level of interface information to retrieve in GetIfTable2Ex function call.
// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-getifentry2ex
type MibIfEntryLevel uint32

const (
	MibIfEntryNormal                  MibIfEntryLevel = 0
	MibIfEntryNormalWithoutStatistics                 = 2
)

// NdisMedium enumeration type identifies the medium types that NDIS drivers support.
// https://docs.microsoft.com/en-us/windows-hardware/drivers/ddi/content/ntddndis/ne-ntddndis-_ndis_medium
type NdisMedium uint32

const (
	NdisMedium802_3 NdisMedium = iota
	NdisMedium802_5
	NdisMediumFddi
	NdisMediumWan
	NdisMediumLocalTalk
	NdisMediumDix // defined for convenience, not a real medium
	NdisMediumArcnetRaw
	NdisMediumArcnet878_2
	NdisMediumAtm
	NdisMediumWirelessWan
	NdisMediumIrda
	NdisMediumBpc
	NdisMediumCoWan
	NdisMedium1394
	NdisMediumInfiniBand
	NdisMediumTunnel
	NdisMediumNative802_11
	NdisMediumLoopback
	NdisMediumWiMAX
	NdisMediumIP
	NdisMediumMax
)

// NdisPhysicalMedium describes NDIS physical medium type.
// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_if_row2
type NdisPhysicalMedium uint32

const (
	NdisPhysicalMediumUnspecified NdisPhysicalMedium = iota
	NdisPhysicalMediumWirelessLan
	NdisPhysicalMediumCableModem
	NdisPhysicalMediumPhoneLine
	NdisPhysicalMediumPowerLine
	NdisPhysicalMediumDSL // includes ADSL and UADSL (G.Lite)
	NdisPhysicalMediumFibreChannel
	NdisPhysicalMedium1394
	NdisPhysicalMediumWirelessWan
	NdisPhysicalMediumNative802_11
	NdisPhysicalMediumBluetooth
	NdisPhysicalMediumInfiniband
	NdisPhysicalMediumWiMax
	NdisPhysicalMediumUWB
	NdisPhysicalMedium802_3
	NdisPhysicalMedium802_5
	NdisPhysicalMediumIrda
	NdisPhysicalMediumWiredWAN
	NdisPhysicalMediumWiredCoWan
	NdisPhysicalMediumOther
	NdisPhysicalMediumNative802_15_4
	NdisPhysicalMediumMax
)

// NetIfAccessType enumeration type specifies the NDIS network interface access type.
// https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ne-ifdef-_net_if_access_type
type NetIfAccessType uint32

const (
	NetIfAccessLoopback NetIfAccessType = iota + 1
	NetIfAccessBroadcast
	NetIfAccessPointToPoint
	NetIfAccessPointToMultiPoint
	NetIfAccessMax
)

// NetIfAdminStatus enumeration type specifies the NDIS network interface administrative status, as described in RFC 2863.
// https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ne-ifdef-net_if_admin_status
type NetIfAdminStatus uint32

const (
	NetIfAdminStatusUp NetIfAdminStatus = iota + 1
	NetIfAdminStatusDown
	NetIfAdminStatusTesting
)

// NetIfConnectionType enumeration type specifies the NDIS network interface connection type.
// https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ne-ifdef-_net_if_connection_type
type NetIfConnectionType uint32

const (
	NetIfConnectionDedicated NetIfConnectionType = iota + 1
	NetIfConnectionPassive
	NetIfConnectionDemand
	NetIfConnectionMaximum
)

// NetIfDirectionType enumeration type specifies the NDIS network interface direction type.
// https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ne-ifdef-net_if_direction_type
type NetIfDirectionType uint32

const (
	NetIfDirectionSendReceive NetIfDirectionType = iota
	NetIfDirectionSendOnly
	NetIfDirectionReceiveOnly
	NetIfDirectionMaximum
)

// NetIfMediaConnectState enumeration type specifies the NDIS network interface connection state.
// https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ne-ifdef-_net_if_media_connect_state
type NetIfMediaConnectState uint32

const (
	MediaConnectStateUnknown NetIfMediaConnectState = iota
	MediaConnectStateConnected
	MediaConnectStateDisconnected
)

// DadState enumeration specifies information about the duplicate address detection (DAD) state for an IPv4 or IPv6 address.
// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ne-nldef-nl_dad_state
type DadState uint32

const (
	DadStateInvalid DadState = iota
	DadStateTentative
	DadStateDuplicate
	DadStateDeprecated
	DadStatePreferred
)

// PrefixOrigin enumeration specifies the origin of an IPv4 or IPv6 address prefix, and is used with the IP_ADAPTER_UNICAST_ADDRESS structure.
// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ne-nldef-nl_prefix_origin
type PrefixOrigin uint32

const (
	PrefixOriginOther PrefixOrigin = iota
	PrefixOriginManual
	PrefixOriginWellKnown
	PrefixOriginDHCP
	PrefixOriginRouterAdvertisement
	PrefixOriginUnchanged = 1 << 4
)

// LinkLocalAddressBehavior enumeration type defines the link local address behavior.
// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ne-nldef-_nl_link_local_address_behavior
type LinkLocalAddressBehavior int32

const (
	LinkLocalAddressAlwaysOff LinkLocalAddressBehavior = iota // Never use link locals.
	LinkLocalAddressDelayed                                   // Use link locals only if no other addresses. (default for IPv4). Legacy mapping: IPAutoconfigurationEnabled.
	LinkLocalAddressAlwaysOn                                  // Always use link locals (default for IPv6).
	LinkLocalAddressUnchanged = -1
)

// OffloadRod enumeration specifies a set of flags that indicate the offload capabilities for an IP interface.
// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ns-nldef-_nl_interface_offload_rod
type OffloadRod uint8

const (
	ChecksumSupported OffloadRod = 1 << iota
	OptionsSupported
	DatagramChecksumSupported
	StreamChecksumSupported
	StreamOptionsSupported
	FastPathCompatible
	LargeSendOffloadSupported
	GiantSendOffloadSupported
)

// RouteOrigin enumeration type defines the origin of the IP route.
// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ne-nldef-nl_route_origin
type RouteOrigin uint32

const (
	RouteOriginManual RouteOrigin = iota
	RouteOriginWellKnown
	RouteOriginDHCP
	RouteOriginRouterAdvertisement
	RouteOrigin6to4
)

// RouteProtocol enumeration type defines the routing mechanism that an IP route was added with, as described in RFC 4292.
// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ne-nldef-nl_route_protocol
type RouteProtocol uint32

const (
	RouteProtocolOther RouteProtocol = iota + 1
	RouteProtocolLocal
	RouteProtocolNetMgmt
	RouteProtocolIcmp
	RouteProtocolEgp
	RouteProtocolGgp
	RouteProtocolHello
	RouteProtocolRip
	RouteProtocolIsIs
	RouteProtocolEsIs
	RouteProtocolCisco
	RouteProtocolBbn
	RouteProtocolOspf
	RouteProtocolBgp
	RouteProtocolIdpr
	RouteProtocolEigrp
	RouteProtocolDvmrp
	RouteProtocolRpl
	RouteProtocolDHCP
	RouteProtocolNTAutostatic   = 10002
	RouteProtocolNTStatic       = 10006
	RouteProtocolNTStaticNonDOD = 10007
)

// RouterDiscoveryBehavior enumeration type defines the router discovery behavior, as described in RFC 2461.
// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ne-nldef-_nl_router_discovery_behavior
type RouterDiscoveryBehavior int32

const (
	RouterDiscoveryDisabled RouterDiscoveryBehavior = iota
	RouterDiscoveryEnabled
	RouterDiscoveryDHCP
	RouterDiscoveryUnchanged = -1
)

// SuffixOrigin enumeration specifies the origin of an IPv4 or IPv6 address suffix, and is used with the IP_ADAPTER_UNICAST_ADDRESS structure.
// https://docs.microsoft.com/en-us/windows/desktop/api/nldef/ne-nldef-nl_suffix_origin
type SuffixOrigin uint32

const (
	SuffixOriginOther SuffixOrigin = iota
	SuffixOriginManual
	SuffixOriginWellKnown
	SuffixOriginDHCP
	SuffixOriginLinkLayerAddress
	SuffixOriginRandom
	SuffixOriginUnchanged = 1 << 4
)

// MibNotificationType enumeration defines the notification type passed to a callback function when a notification occurs.
// https://learn.microsoft.com/en-us/windows-hardware/drivers/network/mib-notification-type
type MibNotificationType uint32

const (
	MibParameterNotification MibNotificationType = iota // Parameter change
	MibAddInstance                                      // Addition
	MibDeleteInstance                                   // Deletion
	MibInitialNotification                              // Initial notification
)

type ChangeCallback interface {
	Unregister() error
}

// TunnelType enumeration type defines the encapsulation method used by a tunnel, as described by the Internet Assigned Names Authority (IANA).
// https://docs.microsoft.com/en-us/windows/desktop/api/ifdef/ne-ifdef-tunnel_type
type TunnelType uint32

const (
	TunnelTypeNone    TunnelType = 0
	TunnelTypeOther              = 1
	TunnelTypeDirect             = 2
	TunnelType6to4               = 11
	TunnelTypeIsatap             = 13
	TunnelTypeTeredo             = 14
	TunnelTypeIPHTTPS            = 15
)

// InterfaceAndOperStatusFlags enumeration type defines interface and operation flags
// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_if_row2
type InterfaceAndOperStatusFlags uint8

const (
	IAOSFHardwareInterface InterfaceAndOperStatusFlags = 1 << iota
	IAOSFFilterInterface
	IAOSFConnectorPresent
	IAOSFNotAuthenticated
	IAOSFNotMediaConnected
	IAOSFPaused
	IAOSFLowPower
	IAOSFEndPointInterface
)

// GAAFlags enumeration defines flags used in GetAdaptersAddresses calls
// https://docs.microsoft.com/en-us/windows/desktop/api/iphlpapi/nf-iphlpapi-getadaptersaddresses
type GAAFlags uint32

const (
	GAAFlagSkipUnicast GAAFlags = 1 << iota
	GAAFlagSkipAnycast
	GAAFlagSkipMulticast
	GAAFlagSkipDNSServer
	GAAFlagIncludePrefix
	GAAFlagSkipFriendlyName
	GAAFlagIncludeWinsInfo
	GAAFlagIncludeGateways
	GAAFlagIncludeAllInterfaces
	GAAFlagIncludeAllCompartments
	GAAFlagIncludeTunnelBindingOrder
	GAAFlagSkipDNSInfo

	GAAFlagDefault    GAAFlags = 0
	GAAFlagSkipAll             = GAAFlagSkipUnicast | GAAFlagSkipAnycast | GAAFlagSkipMulticast | GAAFlagSkipDNSServer | GAAFlagSkipFriendlyName | GAAFlagSkipDNSInfo
	GAAFlagIncludeAll          = GAAFlagIncludePrefix | GAAFlagIncludeWinsInfo | GAAFlagIncludeGateways | GAAFlagIncludeAllInterfaces | GAAFlagIncludeAllCompartments | GAAFlagIncludeTunnelBindingOrder
)

// ScopeLevel enumeration is used with the IP_ADAPTER_ADDRESSES structure to identify scope levels for IPv6 addresses.
// https://docs.microsoft.com/en-us/windows/desktop/api/ws2def/ne-ws2def-scope_level
type ScopeLevel uint32

const (
	ScopeLevelInterface    ScopeLevel = 1
	ScopeLevelLink                    = 2
	ScopeLevelSubnet                  = 3
	ScopeLevelAdmin                   = 4
	ScopeLevelSite                    = 5
	ScopeLevelOrganization            = 8
	ScopeLevelGlobal                  = 14
	ScopeLevelCount                   = 16
)
