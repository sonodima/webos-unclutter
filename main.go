package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/fatih/color"
	"github.com/miekg/dns"
)

// Object that represents the configuration of the program.
var config *Config

// Generated list of domains regular expressions to block
var blacklist *BlackList

// Sends an error response to the client.
func sendError(w dns.ResponseWriter, r *dns.Msg) error {
	msg := &dns.Msg{}
	msg.SetReply(r)
	msg.Rcode = dns.RcodeNameError
	return w.WriteMsg(msg)
}

// Function that handles DNS requests, and selectively blocks them or forwards them to the upstream DNS server.
func handler(w dns.ResponseWriter, r *dns.Msg) {
	if len(r.Question) == 0 {
		log.Println("Received empty DNS request")
		sendError(w, r)
		return
	}

	domain := r.Question[0].Name

	// Check if domain name matches any of the regular expressions in the blocklist
	for _, entry := range blacklist.Domains {
		exp, err := regexp.Compile(entry.Expression)
		if err != nil {
			log.Println("Error compiling", entry.Expression, "expression:", err)
		}

		// If the domain name matches the regular expression, send an error response
		match := exp.MatchString(domain)
		if err == nil && match {
			if config.Logging.Blocked {
				log.Println(color.RedString("[BLOCKED]"), color.HiBlackString("["+entry.Comment+"]"), domain)
			}

			// Send an error response
			sendError(w, r)
			return
		}
	}

	// Fallthrough to Google's DNS server
	if config.Logging.Allowed {
		log.Println(color.GreenString("[ALLOWED]"), domain)
	}

	// Forward the request to the upstream DNS server
	c := &dns.Client{}
	msg, _, err := c.Exchange(r, "8.8.8.8:53")
	if err != nil {
		log.Println("Error resolving request:", err)

		// Send an error response
		sendError(w, r)
		return
	}

	// Send the response to the client
	err = w.WriteMsg(msg)
	if err != nil {
		log.Println("Error writing response:", err)
	}
}

func main() {
	// Print a cool startup banner
	PrintHeader()

	// Set the configuration to the default values
	config = DefaultConfig()

	// Get the name of the configuration file from the 'CONFIG' environment variable
	config_name := "config.yml"
	if env := os.Getenv("CONFIG"); env != "" {
		config_name = env
	}

	// Load the configuration from a file
	if config.Exists(config_name) {
		err := config.Load(config_name)
		if err != nil {
			log.Println("Error loading configuration:", err)
		} else {
			log.Println("Loaded configuration from", config_name)
		}
	}

	// Generate the list of domains to block according to the configuration
	log.Println("Generating blacklist from configuration")
	blacklist = NewBlackList(config)

	log.Println("Using resolver:", config.Network.Resolver)

	// Handle all the requests in the same function
	dns.HandleFunc(".", handler)

	// Starts a DNS server listening on port 53
	listen_address := fmt.Sprintf(":%d", config.Network.ListenPort)
	server := &dns.Server{Addr: listen_address, Net: "udp"}
	log.Println("Starting DNS server at port [" + listen_address + "]")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Failed to start listener:", err)
	}
}
