package config

import (
	"fmt"
	"strings"
)

func GenerateBondingConfig(cfg BondingConfig) string {
	var sb strings.Builder

	if cfg.AutoIfaceUp {
		sb.WriteString(fmt.Sprintf("auto %s\n", cfg.Iface))
	}

	sb.WriteString(fmt.Sprintf("iface %s inet static\n", cfg.Iface))
	sb.WriteString(fmt.Sprintf("    address %s\n", cfg.IP))
	sb.WriteString(fmt.Sprintf("    netmask %s\n", cfg.Netmask))
	sb.WriteString(fmt.Sprintf("    gateway %s\n", cfg.Gateway))

	// Ensure bond-master is at the beginning of bond-slaves
	slaves := make([]string, 0, len(cfg.BondSlaves)+1)
	slaves = append(slaves, cfg.BondMaster)
	for _, slave := range cfg.BondSlaves {
		if slave != cfg.BondMaster {
			slaves = append(slaves, slave)
		}
	}

	sb.WriteString(fmt.Sprintf("    bond-master %s\n", cfg.BondMaster))
	sb.WriteString(fmt.Sprintf("    bond-slaves %s\n", strings.Join(slaves, " ")))

	// Use default value 100 if BondMiimon is nil
	miimon := 100
	if cfg.BondMiimon != nil {
		miimon = *cfg.BondMiimon
	}
	sb.WriteString(fmt.Sprintf("    bond-miimon %d\n", miimon))

	// Add bond-mode with default value "active-backup"
	bondMode := "active-backup"
	if cfg.BondMode != "" {
		bondMode = cfg.BondMode
	}
	sb.WriteString(fmt.Sprintf("    bond-mode %s\n", bondMode))

	return sb.String()
}

func GenerateDSRConfig(cfg DSRConfig) string {
	var sb strings.Builder

	if cfg.AutoIfaceUp {
		sb.WriteString(fmt.Sprintf("auto %s\n", cfg.Iface))
	}

	sb.WriteString(fmt.Sprintf("iface %s inet static\n", cfg.Iface))
	sb.WriteString(fmt.Sprintf("    pre-up ip link add %s type dummy\n", cfg.Iface))
	sb.WriteString(fmt.Sprintf("    pre-down ip link add %s type dummy\n", cfg.Iface))
	sb.WriteString(fmt.Sprintf("    address %s\n", cfg.IP))
	sb.WriteString("    netmask 255.255.255.255\n")

	return sb.String()
}

func GenerateStandardConfig(cfg StandardConfig) string {
	var sb strings.Builder

	if cfg.AutoIfaceUp {
		sb.WriteString(fmt.Sprintf("auto %s\n", cfg.Iface))
	}

	sb.WriteString(fmt.Sprintf("iface %s inet static\n", cfg.Iface))
	sb.WriteString(fmt.Sprintf("    address %s\n", cfg.IP))
	sb.WriteString(fmt.Sprintf("    netmask %s\n", cfg.Netmask))
	sb.WriteString(fmt.Sprintf("    gateway %s\n", cfg.Gateway))

	return sb.String()
}

func GenerateBridgeConfig(cfg BridgeConfig) string {
	var sb strings.Builder

	if cfg.AutoIfaceUp {
		sb.WriteString(fmt.Sprintf("auto %s\n", cfg.Iface))
	}

	sb.WriteString(fmt.Sprintf("iface %s inet manual\n", cfg.Iface))
	sb.WriteString(fmt.Sprintf("    bridge_ports %s\n", strings.Join(cfg.BridgePorts, " ")))

	return sb.String()
}
