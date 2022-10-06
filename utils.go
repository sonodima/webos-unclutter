package main

import (
	"fmt"
	"net"
	"os"
	"runtime"

	"github.com/fatih/color"
)

// Checks if the program is running in a docker container.
func isDockerized() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	return false
}

// Gets the local IP address of the machine.
func getLocalIP() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP, nil
			}
		}
	}

	return nil, fmt.Errorf("no local IP found")
}

// Prints a colored banner to the console, with information about the runtime.
func PrintHeader() {
	// If the program is running as a standalone binary, print the local IP address of the machine.
	// This would not work if the program is running in a docker container, as the IP address would be
	// the internal IP address of the container.
	ip_string := "unknown"
	if !isDockerized() {
		ip, err := getLocalIP()
		if err == nil {
			ip_string = ip.String()
		}
	}

	fmt.Print(
		color.HiBlackString(`           `)+color.HiRedString(`  ____   _____`)+"\n",
		color.HiBlackString(`           `)+color.HiRedString(` / __ \ / ____|`)+"\n",
		color.HiBlackString(`  __      _`)+color.HiRedString(`| |  | | (___ `)+color.HiYellowString(` _   _`)+"\n",
		color.HiBlackString(`  \ \ /\ / `)+color.HiRedString(`| |  | |\___ \`)+color.HiYellowString(`| | | |`)+"\n",
		color.HiBlackString(`   \ V  V /`)+color.HiRedString(`| |__| |____) `)+color.HiYellowString(`| |_| |`)+"\n",
		color.HiBlackString(`    \_/\_/  `)+color.HiRedString(`\____/|_____/ `)+color.HiYellowString(`\__,_|`)+"\n\n",
		color.HiBlackString("       OS | "+runtime.GOOS)+"\n",
		color.HiBlackString("     ARCH | "+runtime.GOARCH)+"\n",
		color.HiBlackString("       IP | "+ip_string)+"\n\n",
	)
}
