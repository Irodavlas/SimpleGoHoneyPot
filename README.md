# SSH Honeypot Server (Go)

A simple SSH honeypot server written in Go that logs password attempts and connection metadata.
Project OnGoing.


## 🔹 Overview

- Listens on **TCP port 4123** for incoming SSH connections.  
- Logs attacker IP addresses, usernames, and attempted passwords.  
- Simulates an SSH server without granting access.  
## 🔹 Features

- Logs connection attempts via `PasswordCallback` and `AuthLogCallback`.  
- Custom SSH server version (`SSH-2.0-OpenSSH_8.2p1 Ubuntu-4ubuntu0.5`) for realism.  
- Concurrent handling of multiple connections with goroutines.  
- Fully written in Go with minimal dependencies (only `golang.org/x/crypto/ssh`).  

unfinished Project
