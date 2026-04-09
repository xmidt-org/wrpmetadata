// SPDX-FileCopyrightText: 2026 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package wrpmetadata

import "strings"

// ConnectionInterface represents the type of network interface a CPE (Customer
// Premises Equipment) uses to connect to the Xmidt cloud.  The CPE reports a
// raw interface name (e.g. "erouter0") that is then mapped to one of these
// canonical types.
type ConnectionInterface int

const (
	// Unknown indicates the interface name was empty or unrecognized
	// as a known CPE interface.
	Unknown ConnectionInterface = iota

	// Cellular represents a generic cellular/mobile data connection.
	Cellular

	// CellularLte is a cellular connection known to be LTE.
	CellularLte

	// Docsis represents a Data Over Cable Service Interface
	// Specification (DOCSIS) connection.
	Docsis

	// Dsl represents a Digital Subscriber Line (DSL)connection.
	Dsl

	// Lan represents a generic LAN connection.
	Lan

	// LanEthernet represents an Ethernet-based LAN connection.
	LanEthernet

	// LanWifi represents a WiFi-based LAN connection.
	LanWifi

	// Other represents a recognized but unmapped interface name.
	Other

	// Wifi represents a direct WiFi connection.
	Wifi
)

var inameMap = map[ConnectionInterface]string{
	Cellular:    "CELLULAR",
	CellularLte: "CELLULAR-LTE",
	Docsis:      "DOCSIS",
	Dsl:         "DSL",
	Lan:         "LAN",
	LanEthernet: "LAN-ETHERNET",
	LanWifi:     "LAN-WIFI",
	Other:       "OTHER",
	Wifi:        "WIFI",
}

func (ci ConnectionInterface) String() string {
	if s, ok := inameMap[ci]; ok {
		return s
	}

	return "UNKNOWN"
}

// ConvertCpeInterface maps a raw CPE network interface name and hardware model
// to a canonical ConnectionInterface.  The interface name is the value the CPE
// sets in the "webpa-interface-used" metadata field (part of X-Xmidt-Metadata).
// The model parameter is the "hw-model" metadata value and is used to
// further refine the mapping (e.g. distinguishing CellularLte from Cellular).
//
// Empty or "unknown" interface names return Unknown.  Unrecognized non-empty
// names return Other.
func ConvertCpeInterface(iName string, model string) ConnectionInterface {
	iName = strings.TrimSpace(iName)
	iName = strings.ToLower(iName)

	if (iName == "") || iName == "unknown" {
		return Unknown
	}

	rv := Other

	switch iName {
	case "brrwan":
		rv = Lan
	case "brlan0":
		rv = LanEthernet
	case "brww0":
		rv = Wifi
	case "erouter0":
		rv = Docsis
	case "eth0":
		rv = LanEthernet
	case "vdsl0":
		rv = Dsl
	case "wlan0", "br-home":
		rv = LanWifi
	case "wwan0":
		rv = Cellular
		switch strings.ToLower(model) {
		case "wnxl11bwl":
			rv = CellularLte
		}
	}

	return rv
}
