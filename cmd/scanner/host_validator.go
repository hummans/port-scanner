package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"unicode"
)

// validateHostnameOrIP validates a given hostname or IP address
// Shared (hostname, IP) validation is performed inline, before delegating
// to hostname and IP-specific methods.
func validateHostnameOrIP(host string) error {
	// Must not be empty
	if host == "" {
		return fmt.Errorf("invalid hostname or IP: empty string")
	}

	// Must not be all digits
	if _, err := strconv.Atoi(host); err == nil {
		return fmt.Errorf("invalid hostname or IP: must not be all digits")
	}

	// Trim any tailing dot
	if host[len(host)-1] == '.' {
		host = host[:len(host)-1]
	}

	// Must have atleast one dot (tld for IP, some octet for IP)
	if i := strings.Index(host, "."); i == -1 {
		return fmt.Errorf("invalid hostname or IP: must have atleast one dot")
	}

	if err := validateHost(host); err != nil {
		// Failed Hostname validation, try IP validation next
		if err := validateIP(host); err != nil {
			// Failed IP validation, too.
			return fmt.Errorf("invalid hostname or IP")
		} else {
			// Passed IP validation
			return nil
		}
	} else {
		// Passed Hostname validation
		return nil
	}
}

// validateHost validates a given host. A valid host is defined as follows:
// - Does not include a scheme
// - Less than 253 characters, including delimiting dots
// - Each label is between 1-63 characters long
// - Contains only ASCII characters
// - TLD is not all numbers
func validateHost(host string) error {
	if len(host) > 253 {
		return fmt.Errorf("invalid host: must be less than 253 characters")
	}

	labels := strings.Split(host, ".")
	for _, label := range labels {
		if len(label) < 1 || len(label) > 63 {
			return fmt.Errorf("invalid host: labels must be between 1-63 characters")
		}
	}

	tld := labels[len(labels)-1]
	if _, err := strconv.Atoi(tld); err == nil {
		return fmt.Errorf("invalid host: tld must not be all numbers")
	}

	if strings.HasPrefix(host, "http") {
		return fmt.Errorf("invalid host: must not include scheme")
	}

	if !isASCII(host) {
		return fmt.Errorf("invalid host: must only contain ASCII characters")
	}

	return nil
}

// validateIP validates an IP address.
// This implementation purposely leaves the definition of a valid IP address
// to the Go standard library.
func validateIP(ip string) error {
	if parsed := net.ParseIP(ip); parsed == nil {
		return fmt.Errorf("invalid IP")
	}

	return nil
}

// isASCII returns true if the given string contains only ASCII characters.
func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}

	return true
}
