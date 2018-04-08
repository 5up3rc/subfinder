// 
// state.go : Contains current program state
// Written By :  @codingo (Michael)
//				 @ice3man (Nizamul Rana)
// 
// Distributed Under MIT License
// Copyrights (C) 2018 Ice3man
//

package helper

// Holds the State read in from the CLI
type State struct {
	Color   		bool		// Whether to use color or not
	Threads 		int 		// Number of threads to use
	Timeout 		int 		// Timeout for requests to different passive sources
	Verbose 		bool 		// Show verbose information
	Domain  		string		// Domain name to find subdomains for
	Recursive 		bool		// Whether perform recursive subdomain discovery or not
	Output 			string		// Name of output file
	IsJSON			bool 		// Provide JSON output file
	Wordlist		string		// Wordlist file for subdomains bruteforcing
	Bruteforce 		bool 		// Flag to decide whether to bruteforce or not 
	WildcardIPs		[]string	// Wildcard IP Structure	
	IsWildcard 		bool 		// Does the host has wildcard subdomains, if yes parse them carefully

	ConfigState  	Config		// Current configuration file state
}

type Config struct {
	VirustotalAPIKey		string	`json:"virustotalApikey"`		// Virustotal API Key

	PassivetotalUsername 	string  `json:"passivetotalUsername"`	// PassiveTotal Username (Email Address)
	PassivetotalKey			string	`json:"passivetotalKey"`		// PassiveTotal api key 
}

func InitState() (state State, err error) {

	// Read the configuration file
	config, err := ReadConfigFile()
	if err != nil {
		return state, err
	}

	return State{true, 10, 180, false, "", false, "", false, "", false, []string{}, true, *config}, nil
}
