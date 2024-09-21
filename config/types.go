package config

// https://wiki.debian.org/Bonding
// https://www.kernel.org/doc/Documentation/networking/bonding.txt
type BondingConfig struct {
	AutoIfaceUp bool
	Iface       string
	IP          string
	Netmask     string
	Gateway     string
	BondMaster  string
	BondSlaves  []string
	BondMiimon  *int
	BondMode    string
}

type DSRConfig struct {
	AutoIfaceUp bool
	Iface       string
	IP          string
}

type StandardConfig struct {
	AutoIfaceUp bool
	Iface       string
	IP          string
	Netmask     string
	Gateway     string
}
