package config

import "testing"

func TestValidateRejectsBadInput(t *testing.T) {
	tests := []struct {
		name   string
		config interface{ Validate() error }
	}{
		{"bonding missing iface", BondingConfig{IP: "192.168.0.1", Netmask: "255.255.255.0", BondSlaves: []string{"eth0"}}},
		{"bonding bad ip", BondingConfig{Iface: "bond0", IP: "not-an-ip", Netmask: "255.255.255.0", BondSlaves: []string{"eth0"}}},
		{"bonding bad netmask", BondingConfig{Iface: "bond0", IP: "192.168.0.1", Netmask: "nope", BondSlaves: []string{"eth0"}}},
		{"bonding no slaves", BondingConfig{Iface: "bond0", IP: "192.168.0.1", Netmask: "255.255.255.0"}},
		{"bonding bad gateway", BondingConfig{Iface: "bond0", IP: "192.168.0.1", Netmask: "255.255.255.0", Gateway: "x", BondSlaves: []string{"eth0"}}},
		{"dsr missing iface", DSRConfig{IP: "10.0.0.1"}},
		{"dsr bad ip", DSRConfig{Iface: "dsr0", IP: "bad"}},
		{"standard bad ip", StandardConfig{Iface: "eth0", IP: "bad", Netmask: "255.255.255.0"}},
		{"bridge no ports", BridgeConfig{Iface: "br0"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.config.Validate(); err == nil {
				t.Errorf("expected validation error, got nil")
			}
		})
	}
}

func TestValidateAcceptsGoodInput(t *testing.T) {
	tests := []struct {
		name   string
		config interface{ Validate() error }
	}{
		{"bonding without gateway", BondingConfig{Iface: "bond0", IP: "192.168.0.1", Netmask: "255.255.255.0", BondSlaves: []string{"eth0", "eth1"}}},
		{"dsr", DSRConfig{Iface: "dsr0", IP: "10.0.0.1"}},
		{"standard without gateway", StandardConfig{Iface: "eth0", IP: "192.168.1.10", Netmask: "255.255.255.0"}},
		{"bridge", BridgeConfig{Iface: "br0", BridgePorts: []string{"eth0"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.config.Validate(); err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}
