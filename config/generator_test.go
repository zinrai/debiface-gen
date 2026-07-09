package config

import (
	"strings"
	"testing"
)

func assertLines(t *testing.T, result string, expected []string) {
	t.Helper()
	lines := strings.Split(strings.TrimSpace(result), "\n")
	if len(lines) != len(expected) {
		t.Fatalf("Expected %d lines, got %d:\n%s", len(expected), len(lines), result)
	}
	for i, expectedLine := range expected {
		if strings.TrimSpace(lines[i]) != strings.TrimSpace(expectedLine) {
			t.Errorf("Line %d: expected '%s', got '%s'", i+1, expectedLine, lines[i])
		}
	}
}

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
				"    bond-slaves eth0 eth1",
				"    bond-miimon 100",
				"    bond-mode active-backup",
			},
		},
		{
			name: "Defaults for miimon and mode",
			config: BondingConfig{
				Iface:      "bond0",
				IP:         "192.168.0.1",
				Netmask:    "255.255.255.0",
				Gateway:    "192.168.0.254",
				BondSlaves: []string{"eth0", "eth1"},
			},
			expected: []string{
				"iface bond0 inet static",
				"    address 192.168.0.1",
				"    netmask 255.255.255.0",
				"    gateway 192.168.0.254",
				"    bond-slaves eth0 eth1",
				"    bond-miimon 100",
				"    bond-mode active-backup",
			},
		},
		{
			name: "No gateway omits the gateway line",
			config: BondingConfig{
				Iface:      "bond0",
				IP:         "192.168.0.1",
				Netmask:    "255.255.255.0",
				BondSlaves: []string{"eth0", "eth1"},
			},
			expected: []string{
				"iface bond0 inet static",
				"    address 192.168.0.1",
				"    netmask 255.255.255.0",
				"    bond-slaves eth0 eth1",
				"    bond-miimon 100",
				"    bond-mode active-backup",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertLines(t, GenerateBondingConfig(tt.config), tt.expected)
		})
	}
}

func TestGenerateDSRConfig(t *testing.T) {
	cfg := DSRConfig{
		AutoIfaceUp: true,
		Iface:       "dsr0",
		IP:          "10.0.0.1",
	}
	expected := []string{
		"auto dsr0",
		"iface dsr0 inet static",
		"    pre-up ip link add dsr0 type dummy",
		"    post-down ip link del dsr0",
		"    address 10.0.0.1",
		"    netmask 255.255.255.255",
	}
	assertLines(t, GenerateDSRConfig(cfg), expected)
}

func TestGenerateStandardConfig(t *testing.T) {
	tests := []struct {
		name     string
		config   StandardConfig
		expected []string
	}{
		{
			name: "With gateway",
			config: StandardConfig{
				AutoIfaceUp: true,
				Iface:       "eth0",
				IP:          "192.168.1.10",
				Netmask:     "255.255.255.0",
				Gateway:     "192.168.1.1",
			},
			expected: []string{
				"auto eth0",
				"iface eth0 inet static",
				"    address 192.168.1.10",
				"    netmask 255.255.255.0",
				"    gateway 192.168.1.1",
			},
		},
		{
			name: "No gateway omits the gateway line",
			config: StandardConfig{
				Iface:   "eth0",
				IP:      "192.168.1.10",
				Netmask: "255.255.255.0",
			},
			expected: []string{
				"iface eth0 inet static",
				"    address 192.168.1.10",
				"    netmask 255.255.255.0",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertLines(t, GenerateStandardConfig(tt.config), tt.expected)
		})
	}
}

func TestGenerateBridgeConfig(t *testing.T) {
	cfg := BridgeConfig{
		AutoIfaceUp: true,
		Iface:       "br0",
		BridgePorts: []string{"eth0", "eth1"},
	}
	expected := []string{
		"auto br0",
		"iface br0 inet manual",
		"    bridge_ports eth0 eth1",
	}
	assertLines(t, GenerateBridgeConfig(cfg), expected)
}

// Helper function to create a pointer to an int
func intPtr(i int) *int {
	return &i
}
