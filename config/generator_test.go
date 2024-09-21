package config

import (
	"strings"
	"testing"
)

func TestGenerateBondingConfig(t *testing.T) {
	tests := []struct {
		name     string
		config   BondingConfig
		expected []string
	}{
		{
			name: "Full configuration",
			config: BondingConfig{
				AutoIfaceUp: true,
				Iface:       "bond0",
				IP:          "192.168.0.1",
				Netmask:     "255.255.255.0",
				Gateway:     "192.168.0.254",
				BondMaster:  "eth0",
				BondSlaves:  []string{"eth0", "eth1"},
				BondMiimon:  intPtr(100),
				BondMode:    "active-backup",
			},
			expected: []string{
				"auto bond0",
				"iface bond0 inet static",
				"    address 192.168.0.1",
				"    netmask 255.255.255.0",
				"    gateway 192.168.0.254",
				"    bond-master eth0",
				"    bond-slaves eth0 eth1",
				"    bond-miimon 100",
				"    bond-mode active-backup",
			},
		},
		{
			name: "Minimal configuration",
			config: BondingConfig{
				Iface:      "bond0",
				IP:         "192.168.0.1",
				Netmask:    "255.255.255.0",
				Gateway:    "192.168.0.254",
				BondMaster: "eth0",
				BondSlaves: []string{"eth0", "eth1"},
			},
			expected: []string{
				"iface bond0 inet static",
				"    address 192.168.0.1",
				"    netmask 255.255.255.0",
				"    gateway 192.168.0.254",
				"    bond-master eth0",
				"    bond-slaves eth0 eth1",
				"    bond-miimon 100",
				"    bond-mode active-backup",
			},
		},
		{
			name: "Custom miimon and mode",
			config: BondingConfig{
				Iface:      "bond0",
				IP:         "192.168.0.1",
				Netmask:    "255.255.255.0",
				Gateway:    "192.168.0.254",
				BondMaster: "eth0",
				BondSlaves: []string{"eth0", "eth1"},
				BondMiimon: intPtr(200),
				BondMode:   "balance-rr",
			},
			expected: []string{
				"iface bond0 inet static",
				"    address 192.168.0.1",
				"    netmask 255.255.255.0",
				"    gateway 192.168.0.254",
				"    bond-master eth0",
				"    bond-slaves eth0 eth1",
				"    bond-miimon 200",
				"    bond-mode balance-rr",
			},
		},
		{
			name: "Slave order with master not first",
			config: BondingConfig{
				Iface:      "bond0",
				IP:         "192.168.0.1",
				Netmask:    "255.255.255.0",
				Gateway:    "192.168.0.254",
				BondMaster: "eth1",
				BondSlaves: []string{"eth0", "eth1", "eth2"},
			},
			expected: []string{
				"iface bond0 inet static",
				"    address 192.168.0.1",
				"    netmask 255.255.255.0",
				"    gateway 192.168.0.254",
				"    bond-master eth1",
				"    bond-slaves eth1 eth0 eth2",
				"    bond-miimon 100",
				"    bond-mode active-backup",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateBondingConfig(tt.config)
			lines := strings.Split(strings.TrimSpace(result), "\n")

			if len(lines) != len(tt.expected) {
				t.Errorf("Expected %d lines, got %d", len(tt.expected), len(lines))
				return
			}

			for i, expectedLine := range tt.expected {
				if strings.TrimSpace(lines[i]) != strings.TrimSpace(expectedLine) {
					t.Errorf("Line %d: expected '%s', got '%s'", i+1, expectedLine, lines[i])
				}
			}
		})
	}
}

// Helper function to create a pointer to an int
func intPtr(i int) *int {
	return &i
}
