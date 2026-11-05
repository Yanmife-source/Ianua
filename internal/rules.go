package waf

import (
	"net/http"
	"strings"
	"log"
)
type Rule struct {
	Name    string
	Pattern string
}

var rules = []Rule{
	// SQL injection
	{"sqli", "' or "}, {"sqli", "\" or "}, {"sqli", "' and "},
	{"sqli", "or 1=1"}, {"sqli", "or '1'='1"}, {"sqli", "' or ''='"},
	{"sqli", "'--"}, {"sqli", "'#"}, {"sqli", "; --"},
	{"sqli", "union select"}, {"sqli", "union all select"},
	{"sqli", "information_schema"}, {"sqli", "sqlite_master"},
	{"sqli", "pg_catalog"}, {"sqli", "sys.tables"}, {"sqli", "v$version"},
	{"sqli", "sleep("}, {"sqli", "benchmark("}, {"sqli", "pg_sleep("},
	{"sqli", "waitfor delay"}, {"sqli", "extractvalue("}, {"sqli", "updatexml("},
	{"sqli", "load_file("}, {"sqli", "into outfile"}, {"sqli", "into dumpfile"},
	{"sqli", "xp_cmdshell"}, {"sqli", "drop table"}, {"sqli", "insert into"},
	{"sqli", "delete from"}, {"sqli", "group_concat("}, {"sqli", "@@version"},
	{"sqli", "version()"}, {"sqli", "database()"},

	// Cross-site scripting
	{"xss", "<script"}, {"xss", "</script"}, {"xss", "javascript:"},
	{"xss", "vbscript:"}, {"xss", "data:text/html"}, {"xss", "srcdoc="},
	{"xss", "<iframe"}, {"xss", "<svg"}, {"xss", "<img"}, {"xss", "<object"},
	{"xss", "<embed"}, {"xss", "<body"}, {"xss", "<link"}, {"xss", "<meta"},
	{"xss", "<style"}, {"xss", "<base"}, {"xss", "<math"}, {"xss", "<details"},
	{"xss", "<marquee"}, {"xss", "<video"}, {"xss", "<audio"},
	{"xss", "onerror="}, {"xss", "onload="}, {"xss", "onclick="},
	{"xss", "onmouseover="}, {"xss", "onfocus="}, {"xss", "onmouseenter="},
	{"xss", "document.cookie"}, {"xss", "document.write"},
	{"xss", "window.location"}, {"xss", "eval("}, {"xss", "alert("},
	{"xss", "prompt("}, {"xss", "confirm("}, {"xss", "fromcharcode"},
	{"xss", "expression("},

	// Path traversal / file inclusion
	{"traversal", "../"}, {"traversal", "..\\"}, {"traversal", "..%2f"},
	{"traversal", "%2e%2e"}, {"traversal", "/etc/passwd"},
	{"traversal", "/etc/shadow"}, {"traversal", "/proc/self"},
	{"traversal", "c:\\windows"}, {"traversal", "boot.ini"}, {"traversal", "win.ini"},
	{"lfi", "php://"}, {"lfi", "file://"}, {"lfi", "phar://"},
	{"lfi", "zip://"}, {"lfi", "expect://"}, {"lfi", "data://"},
	{"nullbyte", "%00"}, {"nullbyte", "\x00"},

	// Command injection (optional for v1)
	{"cmdi", "/bin/sh"}, {"cmdi", "/bin/bash"}, {"cmdi", "cmd.exe"},
	{"cmdi", "powershell"}, {"cmdi", "$("}, {"cmdi", "&& whoami"},
	{"cmdi", "; ls"}, {"cmdi", "| cat "}, {"cmdi", "nc -e"},
}

// Check against the User-Agent header only
var badAgents = []string{"sqlmap", "nikto", "acunetix", "nessus", "dirbuster", "gobuster", "wpscan", "nuclei", "masscan"}

func Inspect(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// WAF LOGIC HERE: Inspect request
		// Example: Block a specific User-Agent or malicious path
		for _,values:=range r.URL.Query(){
			for _,value:=range values{
				for _,rule:=range rules{
					if strings.Contains(value,rule.Pattern){
						http.Error(w,"Forbidden by Web Applcation Firewall",http.StatusForbidden)
						log.Fatalf("Attack detected: %s",rule.Name)
						return
					}
				}	
			}
		}

		// If the request is safe, pass it to the Reverse Proxy
		next.ServeHTTP(w, r)
	})
}