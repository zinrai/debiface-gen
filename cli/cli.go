package cli

import (
	"flag"
	"fmt"
	"strings"

	"github.com/zinrai/debiface-gen/config"
)

func Run() {
	bondingCmd := flag.NewFlagSet("bonding", flag.ExitOnError)
	dsrCmd := flag.NewFlagSet("dsr", flag.ExitOnError)
	standardCmd := flag.NewFlagSet("standard", flag.ExitOnError)

	// Bonding flags
	bondingAuto := bondingCmd.Bool("auto", false, "Up interface after reboot")
	bondingIface := bondingCmd.String("iface", "", "Interface name")
	bondingIP := bondingCmd.String("ip", "", "IP address")
	bondingNetmask := bondingCmd.String("netmask", "", "Netmask")
	bondingGateway := bondingCmd.String("gateway", "", "Gateway")
	bondingMaster := bondingCmd.String("bond-master", "", "Master interface")
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

	if len(flag.Args()) < 1 {
		fmt.Println("Expected 'bonding', 'dsr', or 'standard' subcommands")
		return
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
			BondMaster:  *bondingMaster,
			BondSlaves:  strings.Fields(*bondingSlaves),
			BondMiimon:  miimon,
			BondMode:    *bondingMode,
		}
		fmt.Println(config.GenerateBondingConfig(cfg))

	case "dsr":
		dsrCmd.Parse(flag.Args()[1:])
		cfg := config.DSRConfig{
			AutoIfaceUp: *dsrAuto,
			Iface:       *dsrIface,
			IP:          *dsrIP,
		}
		fmt.Println(config.GenerateDSRConfig(cfg))

	case "standard":
		standardCmd.Parse(flag.Args()[1:])
		cfg := config.StandardConfig{
			AutoIfaceUp: *standardAuto,
			Iface:       *standardIface,
			IP:          *standardIP,
			Netmask:     *standardNetmask,
			Gateway:     *standardGateway,
		}
		fmt.Println(config.GenerateStandardConfig(cfg))

	default:
		fmt.Println("Expected 'bonding', 'dsr', or 'standard' subcommands")
	}
}
