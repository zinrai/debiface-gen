package cli

import (
	"flag"
	"fmt"
	"strings"

	"github.com/zinrai/debiface-gen/config"
)

func Run() error {
	bondingCmd := flag.NewFlagSet("bonding", flag.ExitOnError)
	dsrCmd := flag.NewFlagSet("dsr", flag.ExitOnError)
	standardCmd := flag.NewFlagSet("standard", flag.ExitOnError)
	bridgeCmd := flag.NewFlagSet("bridge", flag.ExitOnError)

	// Bonding flags
	bondingAuto := bondingCmd.Bool("auto", false, "Up interface after reboot")
	bondingIface := bondingCmd.String("iface", "", "Interface name")
	bondingIP := bondingCmd.String("ip", "", "IP address")
	bondingNetmask := bondingCmd.String("netmask", "", "Netmask")
	bondingGateway := bondingCmd.String("gateway", "", "Gateway")
	bondingSlaves := bondingCmd.String("bond-slaves", "", "Slave interfaces (space-separated)")
	bondingMiimon := bondingCmd.Int("bond-miimon", -1, "MII link monitoring interval (default: 100)")
	bondingMode := bondingCmd.String("bond-mode", "", "Bonding mode (default: active-backup)")

	// DSR flags
	dsrAuto := dsrCmd.Bool("auto", false, "Up interface after reboot")
	dsrIface := dsrCmd.String("iface", "", "Interface name")
	dsrIP := dsrCmd.String("ip", "", "IP address")

	// Standard flags
	standardAuto := standardCmd.Bool("auto", false, "Up interface after reboot")
	standardIface := standardCmd.String("iface", "", "Interface name")
	standardIP := standardCmd.String("ip", "", "IP address")
	standardNetmask := standardCmd.String("netmask", "", "Netmask")
	standardGateway := standardCmd.String("gateway", "", "Gateway")

	// Bridge flags
	bridgeAuto := bridgeCmd.Bool("auto", false, "Up interface after reboot")
	bridgeIface := bridgeCmd.String("iface", "", "Interface name")
	bridgePorts := bridgeCmd.String("bridge-ports", "", "Bridge ports (space-separated)")

	if len(flag.Args()) < 1 {
		return fmt.Errorf("expected 'bonding', 'dsr', 'standard', or 'bridge' subcommands")
	}

	switch flag.Arg(0) {
	case "bonding":
		bondingCmd.Parse(flag.Args()[1:])

		var miimon *int
		if *bondingMiimon != -1 {
			miimon = bondingMiimon
		}

		cfg := config.BondingConfig{
			AutoIfaceUp: *bondingAuto,
			Iface:       *bondingIface,
			IP:          *bondingIP,
			Netmask:     *bondingNetmask,
			Gateway:     *bondingGateway,
			BondSlaves:  strings.Fields(*bondingSlaves),
			BondMiimon:  miimon,
			BondMode:    *bondingMode,
		}
		if err := cfg.Validate(); err != nil {
			return err
		}
		fmt.Print(config.GenerateBondingConfig(cfg))

	case "dsr":
		dsrCmd.Parse(flag.Args()[1:])
		cfg := config.DSRConfig{
			AutoIfaceUp: *dsrAuto,
			Iface:       *dsrIface,
			IP:          *dsrIP,
		}
		if err := cfg.Validate(); err != nil {
			return err
		}
		fmt.Print(config.GenerateDSRConfig(cfg))

	case "standard":
		standardCmd.Parse(flag.Args()[1:])
		cfg := config.StandardConfig{
			AutoIfaceUp: *standardAuto,
			Iface:       *standardIface,
			IP:          *standardIP,
			Netmask:     *standardNetmask,
			Gateway:     *standardGateway,
		}
		if err := cfg.Validate(); err != nil {
			return err
		}
		fmt.Print(config.GenerateStandardConfig(cfg))

	case "bridge":
		bridgeCmd.Parse(flag.Args()[1:])
		cfg := config.BridgeConfig{
			AutoIfaceUp: *bridgeAuto,
			Iface:       *bridgeIface,
			BridgePorts: strings.Fields(*bridgePorts),
		}
		if err := cfg.Validate(); err != nil {
			return err
		}
		fmt.Print(config.GenerateBridgeConfig(cfg))

	default:
		return fmt.Errorf("expected 'bonding', 'dsr', 'standard', or 'bridge' subcommands")
	}

	return nil
}
