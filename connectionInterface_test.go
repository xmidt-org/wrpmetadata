// SPDX-FileCopyrightText: 2026 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package wrpmetadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectionInterface_String(t *testing.T) {
	tests := []struct {
		name string
		ci   ConnectionInterface
		want string
	}{
		{name: "Unknown", ci: Unknown, want: "UNKNOWN"},
		{name: "Cellular", ci: Cellular, want: "CELLULAR"},
		{name: "CellularLte", ci: CellularLte, want: "CELLULAR-LTE"},
		{name: "Docsis", ci: Docsis, want: "Docsis"},
		{name: "Dsl", ci: Dsl, want: "Dsl"},
		{name: "Lan", ci: Lan, want: "LAN"},
		{name: "LanEthernet", ci: LanEthernet, want: "LAN-ETHERNET"},
		{name: "LanWifi", ci: LanWifi, want: "LAN-WIFI"},
		{name: "Other", ci: Other, want: "OTHER"},
		{name: "Wifi", ci: Wifi, want: "WIFI"},
		{name: "InvalidValue", ci: ConnectionInterface(999), want: "UNKNOWN"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.ci.String())
		})
	}
}

func TestConvertCpeInterface(t *testing.T) {
	tests := []struct {
		name     string
		iName    string
		hardware string
		want     ConnectionInterface
	}{
		// Unknown cases
		{name: "Empty", iName: "", hardware: "", want: Unknown},
		{name: "UnknownLiteral", iName: "unknown", hardware: "", want: Unknown},
		{name: "UnknownUpperCase", iName: "UNKNOWN", hardware: "", want: Unknown},
		{name: "UnknownMixedCase", iName: "Unknown", hardware: "", want: Unknown},
		{name: "WhitespaceOnly", iName: "   ", hardware: "", want: Unknown},
		{name: "WhitespaceAroundUnknown", iName: "  unknown  ", hardware: "", want: Unknown},

		// Docsis
		{name: "Erouter0", iName: "erouter0", hardware: "", want: Docsis},
		{name: "Erouter0UpperCase", iName: "EROUTER0", hardware: "", want: Docsis},

		// Dsl
		{name: "Vdsl0", iName: "vdsl0", hardware: "", want: Dsl},

		// Lan
		{name: "Brrwan", iName: "brrwan", hardware: "", want: Lan},

		// LanEthernet
		{name: "Brlan0", iName: "brlan0", hardware: "", want: LanEthernet},
		{name: "Eth0", iName: "eth0", hardware: "", want: LanEthernet},

		// LanWifi
		{name: "Wlan0", iName: "wlan0", hardware: "", want: LanWifi},
		{name: "BrHome", iName: "br-home", hardware: "", want: LanWifi},

		// Wifi
		{name: "Brww0", iName: "brww0", hardware: "", want: Wifi},

		// Cellular
		{name: "Wwan0", iName: "wwan0", hardware: "", want: Cellular},
		{name: "Wwan0UnknownHw", iName: "wwan0", hardware: "some-model", want: Cellular},

		// CellularLte
		{name: "Wwan0Lte", iName: "wwan0", hardware: "wnxl11bwl", want: CellularLte},
		{name: "Wwan0LteUpperCase", iName: "wwan0", hardware: "WNXL11BWL", want: CellularLte},

		// Other
		{name: "UnrecognizedInterface", iName: "something-else", hardware: "", want: Other},

		// Whitespace handling
		{name: "LeadingWhitespace", iName: "  erouter0", hardware: "", want: Docsis},
		{name: "TrailingWhitespace", iName: "erouter0  ", hardware: "", want: Docsis},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ConvertCpeInterface(tt.iName, tt.hardware))
		})
	}
}
