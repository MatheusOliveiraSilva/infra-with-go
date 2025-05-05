# Network Basics

Useful link for learning network: https://linuxjourney.com/lesson/network-basics

# The OSI Model for SRE/SWE with Go

The OSI (Open Systems Interconnection) model is a conceptual model that standardizes the functions of a telecommunication or computing system without regard to its underlying internal structure and technology. Its goal is the interoperability of diverse communication systems with standard protocols.

**Explaining the OSI Model for SRE and Go**

The OSI model is divided into 7 layers. For someone aiming to work with SRE and infrastructure using Go, it's interesting to understand how these layers impact network monitoring and communication. Layers like 3 (Network) and 4 (Transport) are crucial when dealing with network protocols and system stability. In Go, knowing how to work with these layers helps implement more efficient and robust infrastructure solutions.

## 1. Physical Layer (Layer 1)

*   **What it is:** Represents the physical means of transmitting raw data (bits). Includes cables (Ethernet, fiber optic), connectors, electrical signals, radio frequencies (Wi-Fi), etc. Defines electrical and mechanical specifications.
*   **Relevance for SRE/SWE with Go:** Low direct relevance. An SWE rarely interacts with this layer via Go code. The network APIs (`net` package) completely abstract this layer.
*   **Indirect Impact:** Physical problems are the root cause of many network failures that manifest in higher layers:
    *   Disconnected or damaged cable.
    *   Weak or interfered Wi-Fi signal.
    *   Problems with network hardware (network card, switch).
*   **Symptoms in your Go system/application:**
    *   **Connection Timeouts:** `net.Dialer` fails when trying to establish a TCP connection (Layer 4).
    *   **Packet Loss:** UDP (Layer 4) can lose data without direct notification; TCP will try to retransmit, causing latency.
    *   **High Latency:** Degraded signals can increase transmission time.
    *   **Intermittent Failures:** Connections that drop and come back.
*   **Diagnosis (Infra Tools):**
    *   Check LEDs on switches and network cards.
    *   Wi-Fi spectrum analysis tools.
    *   Commands like `ethtool` (Linux) to check physical link status (`Link detected: yes`).
    *   Infrastructure monitoring (e.g., SNMP on switches) can indicate port failures.
*   **In Go:** You deal with the *consequences* at Layer 4 (Transport) by implementing robust timeouts, retries, and health checks to detect and work around these physical problems.

---

## 2. Data Link Layer (Layer 2)

*   **What it is:** Responsible for transferring data (frames) between adjacent nodes on the same local network (LAN). Deals with physical addressing (MAC Address) and media access control (e.g., CSMA/CD in Ethernet). Common protocols: Ethernet, PPP, ARP.
*   **Relevance for SRE/SWE with Go:** Moderate, especially in virtualization and container environments.
    *   **Switches:** Operate at this layer, forwarding frames based on MAC addresses.
    *   **MAC Addresses:** Unique identifier for the network interface card.
    *   **ARP (Address Resolution Protocol):** Maps IP addresses (Layer 3) to MAC addresses (Layer 2) on the local network.
    *   **VLANs (Virtual LANs):** Logically segment L2 networks.
*   **Impact and Diagnosis:**
    *   **ARP Problems:** Can prevent hosts on the same local network from communicating, even if L3 connectivity seems to exist. Symptom: `ping` to a local IP fails. Tool: `arp -a` or `ip neigh` (Linux) to inspect the ARP table.
    *   **Duplicate MAC Addresses:** Causes conflicts and erratic network behavior.
    *   **VLAN Configuration:** Configuration errors can improperly isolate machines.
    *   **Switching Loops (Spanning Tree Protocol - STP):** If misconfigured, can cause "broadcast storms" and degrade the network.
*   **Relevance in Cloud/Containers:**
    *   **CNI (Container Network Interface) in Kubernetes:** CNI plugins (Calico, Flannel, Cilium) operate at L2/L3 to connect pods. Understanding L2 helps debug network issues between pods on the same node or between nodes.
    *   **Virtual Networks (AWS VPC, GCP VPC):** Although they abstract much, concepts like security groups (L3/L4 firewall) and network ACLs (L2/L3 firewall) are fundamental.
*   **In Go:**
    *   **Rare Direct Interaction:** Ethernet frames are generally not manipulated directly in Go unless building low-level network tools (e.g., using `gopacket`).
    *   **Observability:** System or switch logs may indicate L2 problems ("MAC flapping," "STP loop"). Connection errors in Go can have a root cause in L2 (e.g., ARP failure).
    *   **Indirect Example:** When configuring a `NetworkPolicy` in Kubernetes (which might be implemented by a CNI using L2/L3 rules), you are influencing the behavior of this layer.

---

## 3. Network Layer (Layer 3)

*   **What it is:** Responsible for logical addressing (IP - IPv4/IPv6), routing packets between different networks, and packet fragmentation/reassembly. Defines how data is sent across the "internetwork." Main protocol: IP (Internet Protocol). Auxiliary protocols: ICMP, OSPF, BGP.
*   **Relevance for SRE/SWE with Go:** High. Fundamental for connectivity between services, machines, and networks.
    *   **IP Addressing:** Each host/service needs an IP to be reached.
    *   **Subnets and CIDR:** Network planning in VPCs (AWS, GCP, Azure) or data centers. Essential for Kubernetes (Cluster CIDR, Service CIDR).
    *   **Routing:** How packets find the path to the destination. Routing tables on hosts and routers.
    *   **Firewalls/ACLs:** IP/port-based rules that allow or block traffic.
    *   **NAT (Network Address Translation):** Allows multiple devices on a private network to share a single public IP.
*   **Impact and Diagnosis:**
    *   **IP Configuration:** Incorrect static IP, DHCP failure. Tool: `ip a` (Linux), `ifconfig` (macOS/older Linux).
    *   **Routing Problems:** Missing or incorrect routes prevent communication between networks. Symptom: `ping` or `traceroute` fails with "Destination Host Unreachable" or "Request timed out." Tool: `ip route` (Linux), `netstat -nr` (macOS/Linux), `traceroute`/`tracert`.
    *   **Firewalls Blocking:** Firewall rules (iptables, nftables, AWS Security Groups, GCP Firewall Rules) can prevent expected connections. Diagnosis: Check rules, use `telnet <host> <port>` or `nc -zv <host> <port>` to test connectivity on specific ports.
    *   **NAT Problems:** Incorrect configuration can prevent inbound or outbound connections.
*   **In Go:**
    *   **`net` package:**
        *   `net.ParseIP`, `net.IPNet`: Manipulation of IP addresses and CIDR networks. Useful for configuration validation, network tools.
        *   `net.Dial`, `net.DialTimeout`, `net.Dialer`: Establishes connections (usually L4), but depends on L3 routing working. Timeouts here can indicate L3 problems.
        *   `net.LookupIP`: Performs DNS lookups (Layer 7) to get IPs (Layer 3).
    *   **ICMP:** Used by `ping`. You can create "raw sockets" in Go to send custom ICMP packets (e.g., for more advanced health checks), using `golang.org/x/net/icmp`.
        ```go
        // Conceptual example of ICMP ping in Go
        import "golang.org/x/net/icmp"
        import "golang.org/x/net/ipv4"
        // ...
        conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
        // // ... build and send ICMP echo message ...
        // // ... read response ...
        ```
    *   **Infra Configuration:** When defining LoadBalancer or Ingress Services in Kubernetes, or configuring VPCs/Subnets in the Cloud, you are working directly with L3 concepts. IaC tools (Terraform, Pulumi) use APIs that manipulate these resources.

---

## 4. Transport Layer (Layer 4)

*   **What it is:** Provides end-to-end communication between processes on different hosts. Handles data segmentation, flow control, error control, and multiplexing using port numbers. Main protocols: TCP (Transmission Control Protocol) and UDP (User Datagram Protocol).
*   **Relevance for SRE/SWE with Go:** Critical. Almost all network communication in applications (HTTP, gRPC, databases, etc.) uses TCP or UDP.
    *   **TCP:**
        *   Connection-oriented (requires 3-way handshake).
        *   Guarantees ordered and reliable delivery of segments (retransmission in case of loss).
        *   Flow control (prevents overwhelming the receiver) and congestion control (prevents overwhelming the network).
        *   Ideal for: HTTP, gRPC, database connections, file transfers.
    *   **UDP:**
        *   Connectionless (sends datagrams directly).
        *   Does not guarantee delivery, order, or absence of duplicates.
        *   Low latency (no handshake), lower overhead.
        *   Ideal for: DNS, video/audio streaming, online games, telemetry (StatsD, logs via syslog), some types of service discovery.
    *   **Ports:** Identify the specific process/application on the host (0-65535). E.g., HTTP on port 80, HTTPS on 443.
*   **Impact and Diagnosis:**
    *   **TCP Handshake Failure:** Connection not established. Symptom: `net.Dial` fails with timeout or "connection refused" (if the process isn't listening on the port). Tool: `telnet <host> <port>`, `nc -zv <host> <port>`.
    *   **Dropped TCP Connections:** "Connection reset by peer," "broken pipe." Can be caused by server crash, intermediate firewall, timeouts.
    *   **UDP Packet Loss:** Application needs to handle this (e.g., retransmit at the application level or accept loss).
    *   **Port Exhaustion:** In high-load scenarios or with short-lived connections (without `TIME_WAIT` reuse), ephemeral ports can be exhausted. Tool: `ss -s` (Linux).
    *   **Latency:** TCP handshake adds RTT (Round Trip Time). Flow/congestion control can limit throughput.
*   **In Go:**
    *   **`net` package:** The heart of L4 interaction.
        *   `net.Dial("tcp", "host:port")`, `net.Listen("tcp", ":port")`: For TCP.
        *   `net.Dial("udp", "host:port")`, `net.ListenPacket("udp", ":port")`: For UDP.
        *   `net.Conn` interface: Abstraction for connections (TCP, Unix sockets). Methods `Read()`, `Write()`, `Close()`, `SetDeadline()`, `SetReadDeadline()`, `SetWriteDeadline()`. **Using deadlines is crucial** to prevent indefinitely blocked goroutines.
        *   `net.TCPConn`, `net.UDPConn`: Concrete types with specific options (e.g., `SetKeepAlive`, `SetNoDelay` for TCP).
    *   **High-Level Libraries:** `net/http`, `google.golang.org/grpc` use `net` underneath. They offer specific timeout configurations (e.g., `http.Server.ReadTimeout`, `http.Transport.DialContext`).
    *   **TCP Connection Timeout Example:**
        ```go
        package main

        import (
        	"fmt"
        	"net"
        	"time"
        )

        func main() {
        	// Try to connect to a non-existent port with a timeout
        	target := "google.com:81" // Port 81 likely closed
        	timeout := 2 * time.Second

        	conn, err := net.DialTimeout("tcp", target, timeout)
        	if err != nil {
        		// The error could be a timeout or another network issue (DNS, routing)
        		fmt.Printf("Error connecting to %s: %v
", target, err)
        		// Check if it was specifically a timeout
        		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
        			fmt.Println("The connection timed out.")
        		}
        		return
        	}
        	defer conn.Close()
        	fmt.Println("Connection successful (unlikely in this example)")
        }
        ```
    *   **Monitoring:** Check connection states (`netstat`, `ss`), connection error metrics, TCP latency, UDP packet loss.

---

## 5. Session Layer (Layer 5)

*   **What it is:** Manages the establishment, maintenance, and termination of sessions (dialogs) between applications. Handles dialog control (who talks when), synchronization, and failure recovery in long sessions.
*   **Practical Relevance Today:** The functionalities of this layer are often integrated into the Transport (L4) or Application (L7) layers. There are no widely used dedicated L5 protocols like TCP (L4) or HTTP (L7).
*   **Analogies in the Modern World/Go:**
    *   **`context.Context` in Go:** Although not a network protocol, `context` serves an analogous purpose *within* a Go application to manage the lifecycle, cancellation, and deadlines of operations (including network calls). Propagating a `context` with a timeout or cancellation through various function calls/goroutines is similar to controlling a work "session."
        ```go
        // Conceptual example with context
        import (
            "context"
            "errors"
            "fmt"
            "time"
        )

        func performNetworkRequest(ctx context.Context, requestData string) (string, error) {
            // Simulate a network call that respects context cancellation
            select {
            case <-time.After(1 * time.Second): // Simulate work
                fmt.Println("Network request finished")
                return "result data", nil
            case <-ctx.Done(): // Context cancelled or deadline exceeded
                fmt.Println("Network request cancelled")
                return "", ctx.Err()
            }
        }


        func main() {
            ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond) // Shorter timeout
            defer cancel()

            // Pass the ctx to functions that perform I/O (network, DB, etc.)
            result, err := performNetworkRequest(ctx, "some data")
            if err != nil {
                if errors.Is(err, context.DeadlineExceeded) {
                    fmt.Println("Operation cancelled due to context (session) timeout")
                } else {
                    fmt.Printf("Network request failed: %v
", err)
                }
            } else {
                fmt.Printf("Received result: %s
", result)
            }
        }

        ```
    *   **gRPC Streams:** The concept of bidirectional streams in gRPC (running over HTTP/2, running over TCP) allows continuous dialog between client and server, managed by the gRPC library, which can be seen as a form of session management. Deadlines in gRPC also fit here.
    *   **Keep-Alive in TCP/HTTP:** Reusing TCP connections (via `Keep-Alive`) for multiple HTTP requests can be seen as a form of session optimization, although managed at L4/L7.
    *   **WebSockets:** Maintain a persistent connection for bidirectional communication, establishing a long session.
*   **Relevance for SRE/SWE with Go:** Indirect, but conceptually important. Understanding the need to manage the lifecycle of complex or long-running operations is crucial.
    *   **Focus in Go:** Using `context.Context` correctly is the main practical manifestation of this concept for controlling timeouts, cancellations, and propagating values related to a request "session."
    *   **Diagnosis:** Related problems can manifest as operations that never finish (lacking appropriate timeouts), resource leaks (goroutines stuck waiting on I/O without cancellation), or failure to clean up state after communication is interrupted.

---

## 6. Presentation Layer (Layer 6)

*   **What it is:** Responsible for ensuring that data is presented in a format understandable by the receiving application. Handles:
    *   **Data Translation/Formatting:** Conversion between different data representations (e.g., ASCII to EBCDIC, though rarely relevant today).
    *   **Encoding/Serialization:** Transforming application data structures (Go structs) into transmission formats (JSON, XML, Protobuf, Gob) and vice-versa.
    *   **Data Compression:** Reducing data size to save bandwidth (gzip, brotli).
    *   **Encryption/Decryption (TLS/SSL):** Ensuring data confidentiality and integrity in transit. Often associated with this layer, although TLS operates "between" L4 and L7.
*   **Relevance for SRE/SWE with Go:** Very high. You deal with this constantly.
    *   **Serialization/Deserialization:** Choosing the right format (JSON, Protobuf, etc.) impacts performance, data size, and interoperability.
        *   **Go:** Packages `encoding/json`, `encoding/xml`, `encoding/gob`, `google.golang.org/protobuf/proto`.
    *   **Compression:** Configuring compression on HTTP servers or using libraries like `compress/gzip`.
        *   **Go:** Middleware for `net/http` or direct use of `compress/*` packages.
    *   **TLS (Transport Layer Security):** Essential for security. Configuring certificates, ciphers, TLS versions.
        *   **Go:** Package `crypto/tls`. Configuration in `http.Server`, `http.Transport`, gRPC/DB clients.
*   **Impact and Diagnosis:**
    *   **Serialization/Deserialization Errors:** "Bad Request" (400) in APIs, failure to process messages. Cause: invalid format, schema incompatibility (Protobuf).
    *   **Performance Issues:** JSON can be slow for large volumes; Protobuf is more efficient. Compression consumes CPU but saves network bandwidth.
    *   **TLS Handshake Errors:** Failure to establish a secure connection. Causes: expired/invalid certificate, cipher/protocol mismatch, configuration issues on client or server. Tools: `openssl s_client -connect host:port`.
    *   **Vulnerabilities:** Poorly configured TLS (old versions, weak ciphers) exposes security risks.
*   **In Go:**
    *   **Working with JSON:**
        ```go
        package main

        import (
        	"encoding/json"
        	"fmt"
        	"log"
        )

        type Request struct {
        	UserID int    `json:"user_id"`
        	Data   string `json:"data"`
        }

        func main() {
        	// Serialization (Go struct -> JSON)
        	req := Request{UserID: 123, Data: "payload"}
        	jsonData, err := json.Marshal(req)
        	if err != nil {
        		log.Fatalf("Error marshaling JSON: %v", err)
        	}
        	fmt.Printf("JSON: %s
", jsonData)

        	// Deserialization (JSON -> Go struct)
        	var receivedReq Request
        	err = json.Unmarshal(jsonData, &receivedReq)
        	if err != nil {
        		log.Fatalf("Error unmarshaling JSON: %v", err)
        	}
        	fmt.Printf("Received struct: %+v
", receivedReq)
        }
        ```
    *   **Basic TLS Configuration on HTTP Server:**
        ```go
        package main

        import (
            "fmt"
            "log"
            "net/http"
        )

        func handler(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintf(w, "Hello, TLS!")
        }

        func main() {
            http.HandleFunc("/", handler)
            // Assume cert.pem and key.pem exist
            log.Println("HTTPS server listening on port 8443...")
            err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil) // Provide paths to cert and key files
            if err != nil {
                log.Fatal("ListenAndServeTLS: ", err)
            }
        }

        ```
    *   **Monitoring:** Metrics for parsing errors, serialization latency, TLS handshake success rate.

---

## 7. Application Layer (Layer 7)

*   **What it is:** The layer closest to the end-user. Provides interfaces and protocols that applications use to communicate over the network. Defines the rules for specific interactions.
*   **Common Protocols:** HTTP/HTTPS, gRPC, DNS, SMTP (email), FTP (file transfer), SSH (secure remote access), MQTT (IoT messaging), LDAP (directory), NTP (time).
*   **Relevance for SRE/SWE with Go:** Maximum. This is where business logic and direct interaction with other services occur.
    *   **Development:** Building APIs (RESTful with HTTP, RPC with gRPC), clients for other services, interacting with DNS, sending emails.
    *   **Operation (SRE):** Monitoring application performance (latency, HTTP/gRPC error rates), L7 health checks (checking if `/health` returns 200 OK), configuring L7 Load Balancers, WAFs (Web Application Firewalls).
*   **Impact and Diagnosis:**
    *   **HTTP Errors:** 4xx codes (client errors: Bad Request, Not Found, Unauthorized) and 5xx codes (server errors: Internal Server Error, Bad Gateway, Service Unavailable).
    *   **gRPC Errors:** Specific status codes (Unavailable, DeadlineExceeded, Unauthenticated).
    *   **DNS Failures:** Inability to resolve hostnames to IPs. Symptom: "host not found" error when trying to connect. Tool: `dig`, `nslookup`.
    *   **Application Latency:** Slow logic on the server, heavy processing, slow dependencies.
    *   **Authentication/Authorization Issues:** Failure to access protected resources.
*   **In Go:**
    *   **`net/http`:** Standard package for building HTTP clients and servers. Popular routers (`gorilla/mux`, `chi`, `gin-gonic`) facilitate creating REST APIs.
    *   **`google.golang.org/grpc`:** Standard library for gRPC in Go.
    *   **`net` (DNS):** `net.LookupHost`, `net.LookupCNAME`, etc., for DNS queries. Specialized libraries like `miekg/dns` for finer control.
    *   **Client Libraries:** SDKs for AWS, GCP, databases (e.g., `database/sql`), etc., operate at this layer.
    *   **Basic HTTP Server Example:**
        ```go
        package main

        import (
        	"fmt"
        	"log"
        	"net/http"
        	"time"
        )

        func helloHandler(w http.ResponseWriter, r *http.Request) {
        	log.Printf("Received request for: %s %s", r.Method, r.URL.Path)
        	// Simulate some processing
        	time.Sleep(50 * time.Millisecond)
        	fmt.Fprintf(w, "Hello, application world!")
        }

        func healthHandler(w http.ResponseWriter, r *http.Request) {
        	// Basic L7 health check
        	w.WriteHeader(http.StatusOK)
        	fmt.Fprintf(w, "OK")
        }

        func main() {
        	mux := http.NewServeMux()
        	mux.HandleFunc("/hello", helloHandler)
        	mux.HandleFunc("/health", healthHandler) // Health check endpoint

        	server := &http.Server{
        		Addr:         ":8080",
        		Handler:      mux,
        		ReadTimeout:  5 * time.Second,  // L4/L7 Protection
        		WriteTimeout: 10 * time.Second, // L4/L7 Protection
        		IdleTimeout:  15 * time.Second, // L4/L7 Protection
        	}

        	log.Println("HTTP server listening on port 8080...")
        	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        		log.Fatalf("Error starting server: %v", err)
        	}
        }
        ```
    *   **Monitoring:** Key metrics (RED - Rate, Errors, Duration) for each endpoint/service. Distributed tracing to follow requests across multiple services. Detailed logging.

---

### How This Fits into the SWE ↔ SRE Flow in Go

1.  **Development (SWE)**
    *   You write code primarily operating at **L7** (HTTP/gRPC handlers, business logic) and **L6** (JSON/Protobuf serialization, TLS config).
    *   You use libraries interacting with **L4** (`net.Dial`, `net.Listen`) and indirectly relying on **L3** (DNS resolution, IP routing).
    *   You implement resilience (timeouts, retries, circuit breakers) based on failures observed at **L4** and **L7**.
    *   You use `context.Context` (**L5** analog) to manage lifecycle.

2.  **Operation (SRE)**
    *   Monitors health and performance at **L7** (application metrics, logs, tracing) and **L4** (TCP connections, errors).
    *   Configures health checks at **L7** (HTTP `/health`) and **L4** (TCP connect) for Load Balancers and orchestrators (Kubernetes).
    *   Diagnoses problems using tools inspecting **L3** (`ping`, `traceroute`), **L4** (`telnet`, `nc`, `ss`), and **L7** (`curl`, `grpcurl`).
    *   Manages network configuration at **L3** (IPs, routes, firewalls/Security Groups) and **L6** (TLS certificates).
    *   Responds to alerts about failures that can originate at any layer but often manifest as errors at **L4** or **L7**.

3.  **Infrastructure in Go**
    *   Creates CLI tools/agents using `net` to interact with **L3/L4** (port scanners, connectivity tests).
    *   Develops Prometheus exporters collecting **L7** metrics (e.g., `promhttp`).
    *   Writes Kubernetes Operators manipulating network resources like Services (**L4**), Ingress (**L7**), NetworkPolicies (**L3/L4**).
    *   Implements service discovery logic based on DNS (**L7**) or orchestrator APIs.

---

**Quick Summary:**

*   **Main Focus (Direct Coding/Configuration):** Layers **7 (Application)**, **6 (Presentation)**, **4 (Transport)**, and **3 (Network)**.
*   **Layer 4 (TCP/UDP):** Critical decisions about reliability vs. performance, connection management, timeouts. It's where "the rubber meets the road" for most applications.
*   **Layers 5-7:** Where application logic, security (TLS), data format, and specific protocols reside. Go offers excellent standard and third-party libraries here (`net/http`, `crypto/tls`, `context`, `grpc`).
*   **Layers 1-3:** Fundamental for connectivity to exist, but usually managed by the infrastructure (physical, cloud provider, network team) and diagnosed with specific tools when problems arise. SREs need to understand how failures here impact upper layers.

Understanding how these layers interact helps build more robust, secure, and observable Go applications, and diagnose problems more effectively, whether in code (SWE) or operation (SRE).

