//go:build windows
// +build windows

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019-2022 WireGuard LLC. All Rights Reserved.
 */

package iphlpapi

import (
	"golang.org/x/sys/windows"
)

// LUID represents a network interface.
type LUID uint64

// GUID method converts a locally unique identifier (LUID) for a network interface to a globally unique identifier (GUID) for the interface.
// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-convertinterfaceluidtoguid
func (luid LUID) GUID() (*windows.GUID, error) {
	guid := &windows.GUID{}
	err := convertInterfaceLUIDToGUID(&luid, guid)
	if err != nil {
		return nil, err
	}
	return guid, nil
}

// LUIDFromGUID function converts a globally unique identifier (GUID) for a network interface to the locally unique identifier (LUID) for the interface.
// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-convertinterfaceguidtoluid
func LUIDFromGUID(guid *windows.GUID) (LUID, error) {
	var luid LUID
	err := convertInterfaceGUIDToLUID(guid, &luid)
	if err != nil {
		return 0, err
	}
	return luid, nil
}

// LUIDFromIndex function converts a local index for a network interface to the locally unique identifier (LUID) for the interface.
// https://docs.microsoft.com/en-us/windows/desktop/api/netioapi/nf-netioapi-convertinterfaceindextoluid
func LUIDFromIndex(index uint32) (LUID, error) {
	var luid LUID
	err := convertInterfaceIndexToLUID(index, &luid)
	if err != nil {
		return 0, err
	}
	return luid, nil
}
