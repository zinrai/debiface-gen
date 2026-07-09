package config

import (
	"fmt"
	"net"
)

// validateIP returns an error if value is not a parseable IP address.
// Netmasks in dotted-quad form (e.g. 255.255.255.0) parse as IPv4 here too,
// which is enough to catch empty or typo'd values before they reach a
// not-yet-networked host.
func validateIP(field, value string) error {
	if net.ParseIP(value) == nil {
		return fmt.Errorf("%s: invalid IP address %q", field, value)
	}
	return nil
}

func (c BondingConfig) Validate() error {
	if c.Iface == "" {
		return fmt.Errorf("iface is required")
	}
	if err := validateIP("ip", c.IP); err != nil {
		return err
	}
	if err := validateIP("netmask", c.Netmask); err != nil {
		return err
	}
	if c.Gateway != "" {
		if err := validateIP("gateway", c.Gateway); err != nil {
			return err
		}
	}
	if len(c.BondSlaves) == 0 {
		return fmt.Errorf("at least one bond-slave is required")
	}
	return nil
}

func (c DSRConfig) Validate() error {
	if c.Iface == "" {
		return fmt.Errorf("iface is required")
	}
	return validateIP("ip", c.IP)
}

func (c StandardConfig) Validate() error {
	if c.Iface == "" {
		return fmt.Errorf("iface is required")
	}
	if err := validateIP("ip", c.IP); err != nil {
		return err
	}
	if err := validateIP("netmask", c.Netmask); err != nil {
		return err
	}
	if c.Gateway != "" {
		if err := validateIP("gateway", c.Gateway); err != nil {
			return err
		}
	}
	return nil
}

func (c BridgeConfig) Validate() error {
	if c.Iface == "" {
		return fmt.Errorf("iface is required")
	}
	if len(c.BridgePorts) == 0 {
		return fmt.Errorf("at least one bridge-port is required")
	}
	return nil
}
