//go:build windows && (386 || arm)
// +build windows
// +build 386 arm

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019-2022 WireGuard LLC. All Rights Reserved.
 */

package iphlpapi

// MibIPInterfaceRow structure stores interface management information for a particular IP address family on a network interface.
// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/ns-netioapi-_mib_ipinterface_row
type MibIPInterfaceRow struct {
	Family                               AddressFamily
	_                                    [4]byte
	InterfaceLUID                        LUID
	InterfaceIndex                       uint32
	MaxReassemblySize                    uint32
	InterfaceIdentifier                  uint64
	MinRouterAdvertisementInterval       uint32
	MaxRouterAdvertisementInterval       uint32
	AdvertisingEnabled                   bool
	ForwardingEnabled                    bool
	WeakHostSend                         bool
	WeakHostReceive                      bool
	UseAutomaticMetric                   bool
	UseNeighborUnreachabilityDetection   bool
	ManagedAddressConfigurationSupported bool
	OtherStatefulConfigurationSupported  bool
	AdvertiseDefaultRoute                bool
	RouterDiscoveryBehavior              RouterDiscoveryBehavior
	DadTransmits                         uint32
	BaseReachableTime                    uint32
	RetransmitTime                       uint32
	PathMTUDiscoveryTimeout              uint32
	LinkLocalAddressBehavior             LinkLocalAddressBehavior
	LinkLocalAddressTimeout              uint32
	ZoneIndices                          [ScopeLevelCount]uint32
	SitePrefixLength                     uint32
	Metric                               uint32
	NLMTU                                uint32
	Connected                            bool
	SupportsWakeUpPatterns               bool
	SupportsNeighborDiscovery            bool
	SupportsRouterDiscovery              bool
	ReachableTime                        uint32
	TransmitOffload                      OffloadRod
	ReceiveOffload                       OffloadRod
	DisableDefaultRoutes                 bool
}
