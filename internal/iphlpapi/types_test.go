//go:build windows
// +build windows

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019-2022 WireGuard LLC. All Rights Reserved.
 */

package iphlpapi

import (
	"testing"
	"unsafe"
)

const (
	mibIPInterfaceRowSize                                       = 168
	mibIPInterfaceRowInterfaceLUIDOffset                        = 8
	mibIPInterfaceRowInterfaceIndexOffset                       = 16
	mibIPInterfaceRowMaxReassemblySizeOffset                    = 20
	mibIPInterfaceRowInterfaceIdentifierOffset                  = 24
	mibIPInterfaceRowMinRouterAdvertisementIntervalOffset       = 32
	mibIPInterfaceRowMaxRouterAdvertisementIntervalOffset       = 36
	mibIPInterfaceRowAdvertisingEnabledOffset                   = 40
	mibIPInterfaceRowForwardingEnabledOffset                    = 41
	mibIPInterfaceRowWeakHostSendOffset                         = 42
	mibIPInterfaceRowWeakHostReceiveOffset                      = 43
	mibIPInterfaceRowUseAutomaticMetricOffset                   = 44
	mibIPInterfaceRowUseNeighborUnreachabilityDetectionOffset   = 45
	mibIPInterfaceRowManagedAddressConfigurationSupportedOffset = 46
	mibIPInterfaceRowOtherStatefulConfigurationSupportedOffset  = 47
	mibIPInterfaceRowAdvertiseDefaultRouteOffset                = 48
	mibIPInterfaceRowRouterDiscoveryBehaviorOffset              = 52
	mibIPInterfaceRowDadTransmitsOffset                         = 56
	mibIPInterfaceRowBaseReachableTimeOffset                    = 60
	mibIPInterfaceRowRetransmitTimeOffset                       = 64
	mibIPInterfaceRowPathMTUDiscoveryTimeoutOffset              = 68
	mibIPInterfaceRowLinkLocalAddressBehaviorOffset             = 72
	mibIPInterfaceRowLinkLocalAddressTimeoutOffset              = 76
	mibIPInterfaceRowZoneIndicesOffset                          = 80
	mibIPInterfaceRowSitePrefixLengthOffset                     = 144
	mibIPInterfaceRowMetricOffset                               = 148
	mibIPInterfaceRowNLMTUOffset                                = 152
	mibIPInterfaceRowConnectedOffset                            = 156
	mibIPInterfaceRowSupportsWakeUpPatternsOffset               = 157
	mibIPInterfaceRowSupportsNeighborDiscoveryOffset            = 158
	mibIPInterfaceRowSupportsRouterDiscoveryOffset              = 159
	mibIPInterfaceRowReachableTimeOffset                        = 160
	mibIPInterfaceRowTransmitOffloadOffset                      = 164
	mibIPInterfaceRowReceiveOffloadOffset                       = 165
	mibIPInterfaceRowDisableDefaultRoutesOffset                 = 166
)

func TestMibIPInterfaceRow(t *testing.T) {
	s := MibIPInterfaceRow{}
	sp := uintptr(unsafe.Pointer(&s))
	const actualTestMibIPInterfaceRowSize = unsafe.Sizeof(s)

	if actualTestMibIPInterfaceRowSize != mibIPInterfaceRowSize {
		t.Errorf("Size of MibIPInterfaceRow is %d, although %d is expected.", actualTestMibIPInterfaceRowSize, mibIPInterfaceRowSize)
	}

	offset := uintptr(unsafe.Pointer(&s.InterfaceLUID)) - sp
	if offset != mibIPInterfaceRowInterfaceLUIDOffset {
		t.Errorf("MibIPInterfaceRow.InterfaceLUID offset is %d although %d is expected", offset, mibIPInterfaceRowInterfaceLUIDOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.InterfaceIndex)) - sp
	if offset != mibIPInterfaceRowInterfaceIndexOffset {
		t.Errorf("MibIPInterfaceRow.InterfaceIndex offset is %d although %d is expected", offset, mibIPInterfaceRowInterfaceIndexOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.MaxReassemblySize)) - sp
	if offset != mibIPInterfaceRowMaxReassemblySizeOffset {
		t.Errorf("mibIPInterfaceRow.MaxReassemblySize offset is %d although %d is expected", offset, mibIPInterfaceRowMaxReassemblySizeOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.InterfaceIdentifier)) - sp
	if offset != mibIPInterfaceRowInterfaceIdentifierOffset {
		t.Errorf("MibIPInterfaceRow.InterfaceIdentifier offset is %d although %d is expected", offset, mibIPInterfaceRowInterfaceIdentifierOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.MinRouterAdvertisementInterval)) - sp
	if offset != mibIPInterfaceRowMinRouterAdvertisementIntervalOffset {
		t.Errorf("MibIPInterfaceRow.MinRouterAdvertisementInterval offset is %d although %d is expected", offset, mibIPInterfaceRowMinRouterAdvertisementIntervalOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.MaxRouterAdvertisementInterval)) - sp
	if offset != mibIPInterfaceRowMaxRouterAdvertisementIntervalOffset {
		t.Errorf("MibIPInterfaceRow.MaxRouterAdvertisementInterval offset is %d although %d is expected", offset, mibIPInterfaceRowMaxRouterAdvertisementIntervalOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.AdvertisingEnabled)) - sp
	if offset != mibIPInterfaceRowAdvertisingEnabledOffset {
		t.Errorf("MibIPInterfaceRow.AdvertisingEnabled offset is %d although %d is expected", offset, mibIPInterfaceRowAdvertisingEnabledOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.ForwardingEnabled)) - sp
	if offset != mibIPInterfaceRowForwardingEnabledOffset {
		t.Errorf("MibIPInterfaceRow.ForwardingEnabled offset is %d although %d is expected", offset, mibIPInterfaceRowForwardingEnabledOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.WeakHostSend)) - sp
	if offset != mibIPInterfaceRowWeakHostSendOffset {
		t.Errorf("MibIPInterfaceRow.WeakHostSend offset is %d although %d is expected", offset, mibIPInterfaceRowWeakHostSendOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.WeakHostReceive)) - sp
	if offset != mibIPInterfaceRowWeakHostReceiveOffset {
		t.Errorf("MibIPInterfaceRow.WeakHostReceive offset is %d although %d is expected", offset, mibIPInterfaceRowWeakHostReceiveOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.UseAutomaticMetric)) - sp
	if offset != mibIPInterfaceRowUseAutomaticMetricOffset {
		t.Errorf("MibIPInterfaceRow.UseAutomaticMetric offset is %d although %d is expected", offset, mibIPInterfaceRowUseAutomaticMetricOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.UseNeighborUnreachabilityDetection)) - sp
	if offset != mibIPInterfaceRowUseNeighborUnreachabilityDetectionOffset {
		t.Errorf("MibIPInterfaceRow.UseNeighborUnreachabilityDetection offset is %d although %d is expected", offset, mibIPInterfaceRowUseNeighborUnreachabilityDetectionOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.ManagedAddressConfigurationSupported)) - sp
	if offset != mibIPInterfaceRowManagedAddressConfigurationSupportedOffset {
		t.Errorf("MibIPInterfaceRow.ManagedAddressConfigurationSupported offset is %d although %d is expected", offset, mibIPInterfaceRowManagedAddressConfigurationSupportedOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.OtherStatefulConfigurationSupported)) - sp
	if offset != mibIPInterfaceRowOtherStatefulConfigurationSupportedOffset {
		t.Errorf("MibIPInterfaceRow.OtherStatefulConfigurationSupported offset is %d although %d is expected", offset, mibIPInterfaceRowOtherStatefulConfigurationSupportedOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.AdvertiseDefaultRoute)) - sp
	if offset != mibIPInterfaceRowAdvertiseDefaultRouteOffset {
		t.Errorf("MibIPInterfaceRow.AdvertiseDefaultRoute offset is %d although %d is expected", offset, mibIPInterfaceRowAdvertiseDefaultRouteOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.RouterDiscoveryBehavior)) - sp
	if offset != mibIPInterfaceRowRouterDiscoveryBehaviorOffset {
		t.Errorf("MibIPInterfaceRow.RouterDiscoveryBehavior offset is %d although %d is expected", offset, mibIPInterfaceRowRouterDiscoveryBehaviorOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.DadTransmits)) - sp
	if offset != mibIPInterfaceRowDadTransmitsOffset {
		t.Errorf("MibIPInterfaceRow.DadTransmits offset is %d although %d is expected", offset, mibIPInterfaceRowDadTransmitsOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.BaseReachableTime)) - sp
	if offset != mibIPInterfaceRowBaseReachableTimeOffset {
		t.Errorf("MibIPInterfaceRow.BaseReachableTime offset is %d although %d is expected", offset, mibIPInterfaceRowBaseReachableTimeOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.RetransmitTime)) - sp
	if offset != mibIPInterfaceRowRetransmitTimeOffset {
		t.Errorf("MibIPInterfaceRow.RetransmitTime offset is %d although %d is expected", offset, mibIPInterfaceRowRetransmitTimeOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.PathMTUDiscoveryTimeout)) - sp
	if offset != mibIPInterfaceRowPathMTUDiscoveryTimeoutOffset {
		t.Errorf("MibIPInterfaceRow.PathMTUDiscoveryTimeout offset is %d although %d is expected", offset, mibIPInterfaceRowPathMTUDiscoveryTimeoutOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.LinkLocalAddressBehavior)) - sp
	if offset != mibIPInterfaceRowLinkLocalAddressBehaviorOffset {
		t.Errorf("MibIPInterfaceRow.LinkLocalAddressBehavior offset is %d although %d is expected", offset, mibIPInterfaceRowLinkLocalAddressBehaviorOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.LinkLocalAddressTimeout)) - sp
	if offset != mibIPInterfaceRowLinkLocalAddressTimeoutOffset {
		t.Errorf("MibIPInterfaceRow.LinkLocalAddressTimeout offset is %d although %d is expected", offset, mibIPInterfaceRowLinkLocalAddressTimeoutOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.ZoneIndices)) - sp
	if offset != mibIPInterfaceRowZoneIndicesOffset {
		t.Errorf("MibIPInterfaceRow.ZoneIndices offset is %d although %d is expected", offset, mibIPInterfaceRowZoneIndicesOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.SitePrefixLength)) - sp
	if offset != mibIPInterfaceRowSitePrefixLengthOffset {
		t.Errorf("MibIPInterfaceRow.SitePrefixLength offset is %d although %d is expected", offset, mibIPInterfaceRowSitePrefixLengthOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.Metric)) - sp
	if offset != mibIPInterfaceRowMetricOffset {
		t.Errorf("MibIPInterfaceRow.Metric offset is %d although %d is expected", offset, mibIPInterfaceRowMetricOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.NLMTU)) - sp
	if offset != mibIPInterfaceRowNLMTUOffset {
		t.Errorf("MibIPInterfaceRow.NLMTU offset is %d although %d is expected", offset, mibIPInterfaceRowNLMTUOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.Connected)) - sp
	if offset != mibIPInterfaceRowConnectedOffset {
		t.Errorf("MibIPInterfaceRow.Connected offset is %d although %d is expected", offset, mibIPInterfaceRowConnectedOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.SupportsWakeUpPatterns)) - sp
	if offset != mibIPInterfaceRowSupportsWakeUpPatternsOffset {
		t.Errorf("MibIPInterfaceRow.SupportsWakeUpPatterns offset is %d although %d is expected", offset, mibIPInterfaceRowSupportsWakeUpPatternsOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.SupportsNeighborDiscovery)) - sp
	if offset != mibIPInterfaceRowSupportsNeighborDiscoveryOffset {
		t.Errorf("MibIPInterfaceRow.SupportsNeighborDiscovery offset is %d although %d is expected", offset, mibIPInterfaceRowSupportsNeighborDiscoveryOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.SupportsRouterDiscovery)) - sp
	if offset != mibIPInterfaceRowSupportsRouterDiscoveryOffset {
		t.Errorf("MibIPInterfaceRow.SupportsRouterDiscovery offset is %d although %d is expected", offset, mibIPInterfaceRowSupportsRouterDiscoveryOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.ReachableTime)) - sp
	if offset != mibIPInterfaceRowReachableTimeOffset {
		t.Errorf("MibIPInterfaceRow.ReachableTime offset is %d although %d is expected", offset, mibIPInterfaceRowReachableTimeOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.TransmitOffload)) - sp
	if offset != mibIPInterfaceRowTransmitOffloadOffset {
		t.Errorf("MibIPInterfaceRow.TransmitOffload offset is %d although %d is expected", offset, mibIPInterfaceRowTransmitOffloadOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.ReceiveOffload)) - sp
	if offset != mibIPInterfaceRowReceiveOffloadOffset {
		t.Errorf("MibIPInterfaceRow.ReceiveOffload offset is %d although %d is expected", offset, mibIPInterfaceRowReceiveOffloadOffset)
	}

	offset = uintptr(unsafe.Pointer(&s.DisableDefaultRoutes)) - sp
	if offset != mibIPInterfaceRowDisableDefaultRoutesOffset {
		t.Errorf("MibIPInterfaceRow.DisableDefaultRoutes offset is %d although %d is expected", offset, mibIPInterfaceRowDisableDefaultRoutesOffset)
	}
}
