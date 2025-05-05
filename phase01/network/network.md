# Network Basics

Useful link for learning network: https://linuxjourney.com/lesson/network-basics

# O Modelo OSI para SRE/SWE com Go

O modelo OSI (Open Systems Interconnection) é um modelo conceitual que padroniza as funções de um sistema de telecomunicação ou de computação sem considerar sua estrutura interna e tecnologia subjacentes. Sua meta é a interoperabilidade de diversos sistemas de comunicação com protocolos padrão.

**Explicando o modelo OSI para SRE e Go**

O modelo OSI é dividido em 7 camadas, e para alguém com o objetivo de trabalhar com SRE e infraestrutura usando Go, é interessante entender como essas camadas impactam a monitoração e a comunicação de rede. Camadas como a 3 (Rede) e a 4 (Transporte) são cruciais quando lidamos com protocolos de rede e a estabilidade do sistema. Em Go, saber como trabalhar com essas camadas ajuda a implementar soluções de infraestrutura mais eficientes e robustas.

## 1. Camada Física (Layer 1)

*   **O que é:** Representa os meios físicos de transmissão de dados brutos (bits). Inclui cabos (Ethernet, fibra óptica), conectores, sinais elétricos, rádio frequência (Wi-Fi), etc. Define as especificações elétricas e mecânicas.
*   **Relevância para SRE/SWE com Go:** Baixa direta. Raramente um SWE interage com esta camada via código Go. As APIs de rede (`net` package) abstraem completamente esta camada.
*   **Impacto indireto:** Problemas físicos são a causa raiz de muitas falhas de rede que se manifestam em camadas superiores:
    *   Cabo desconectado ou danificado.
    *   Sinal Wi-Fi fraco ou com interferência.
    *   Problemas no hardware de rede (placa de rede, switch).
*   **Sintomas no seu sistema/aplicação Go:**
    *   **Timeouts de conexão:** `net.Dialer` falha ao tentar estabelecer uma conexão TCP (Layer 4).
    *   **Perda de pacotes:** UDP (Layer 4) pode perder dados sem notificação direta; TCP tentará retransmitir, causando latência.
    *   **Alta latência:** Sinais degradados podem aumentar o tempo de transmissão.
    *   **Falhas intermitentes:** Conexões que caem e voltam.
*   **Diagnóstico (Ferramentas de Infra):**
    *   Verificar LEDs nos switches e placas de rede.
    *   Ferramentas de análise de espectro Wi-Fi.
    *   Comandos como `ethtool` (Linux) para verificar status do link físico (`Link detected: yes`).
    *   Monitoramento de infraestrutura (e.g., SNMP em switches) pode indicar falhas de porta.
*   **Em Go:** Você lida com as *consequências* na camada 4 (Transporte), implementando timeouts, retries e health checks robustos para detectar e contornar esses problemas físicos.

---

## 2. Camada de Enlace (Layer 2)

*   **O que é:** Responsável pela transferência de dados (frames) entre nós adjacentes na mesma rede local (LAN). Lida com endereçamento físico (MAC Address) e controle de acesso ao meio (ex: CSMA/CD em Ethernet). Protocolos comuns: Ethernet, PPP, ARP.
*   **Relevância para SRE/SWE com Go:** Moderada, principalmente em ambientes de virtualização e contêineres.
    *   **Switches:** Operam nesta camada, encaminhando frames baseados em MAC addresses.
    *   **MAC Addresses:** Identificador único da placa de rede.
    *   **ARP (Address Resolution Protocol):** Mapeia endereços IP (Layer 3) para MAC addresses (Layer 2) na rede local.
    *   **VLANs (Virtual LANs):** Segmentam redes L2 logicamente.
*   **Impacto e Diagnóstico:**
    *   **Problemas de ARP:** Podem impedir que hosts na mesma rede local se comuniquem, mesmo que a conectividade L3 pareça existir. Sintoma: `ping` para IP local falha. Ferramenta: `arp -a` ou `ip neigh` (Linux) para inspecionar a tabela ARP.
    *   **Duplicação de MAC Address:** Causa conflitos e comportamento errático na rede.
    *   **Configuração de VLAN:** Erros de configuração podem isolar máquinas indevidamente.
    *   **Loops de Switching (Spanning Tree Protocol - STP):** Mal configurado, pode causar "broadcast storms" e degradar a rede.
*   **Relevância em Cloud/Containers:**
    *   **CNI (Container Network Interface) no Kubernetes:** Plugins CNI (Calico, Flannel, Cilium) operam em L2/L3 para conectar pods. Entender L2 ajuda a depurar problemas de rede entre pods no mesmo nó ou entre nós.
    *   **Redes Virtuais (AWS VPC, GCP VPC):** Embora abstraiam muito, conceitos como grupos de segurança (firewall L3/L4) e ACLs de rede (firewall L2/L3) são fundamentais.
*   **Em Go:**
    *   **Interação Direta Rara:** Geralmente não se manipula frames Ethernet diretamente em Go, a menos que se esteja construindo ferramentas de rede de baixo nível (e.g., usando `gopacket`).
    *   **Observabilidade:** Logs de sistema ou de switches podem indicar problemas L2 ("MAC flapping", "STP loop"). Erros de conexão em Go podem ter causa raiz em L2 (ex: falha de ARP).
    *   **Exemplo Indireto:** Ao configurar um `NetworkPolicy` no Kubernetes (que pode ser implementado por um CNI usando regras L2/L3), você está influenciando o comportamento desta camada.

---

## 3. Camada de Rede (Layer 3)

*   **O que é:** Responsável pelo endereçamento lógico (IP - IPv4/IPv6), roteamento de pacotes entre diferentes redes e fragmentação/remontagem de pacotes. Define como os dados são enviados através da "internetwork". Protocolo principal: IP (Internet Protocol). Protocolos auxiliares: ICMP, OSPF, BGP.
*   **Relevância para SRE/SWE com Go:** Alta. Fundamental para conectividade entre serviços, máquinas e redes.
    *   **Endereçamento IP:** Cada host/serviço precisa de um IP para ser alcançado.
    *   **Sub-redes e CIDR:** Planejamento de rede em VPCs (AWS, GCP, Azure) ou data centers. Essencial para Kubernetes (Cluster CIDR, Service CIDR).
    *   **Roteamento:** Como pacotes encontram o caminho até o destino. Tabelas de roteamento nos hosts e roteadores.
    *   **Firewalls/ACLs:** Regras baseadas em IP/porta que permitem ou bloqueiam tráfego.
    *   **NAT (Network Address Translation):** Permite que múltiplos dispositivos em uma rede privada compartilhem um único IP público.
*   **Impacto e Diagnóstico:**
    *   **Configuração de IP:** IP estático incorreto, falha no DHCP. Ferramenta: `ip a` (Linux), `ifconfig` (macOS/Linux antigo).
    *   **Problemas de Roteamento:** Rotas faltando ou incorretas impedem a comunicação entre redes. Sintoma: `ping` ou `traceroute` falha com "Destination Host Unreachable" ou "Request timed out". Ferramenta: `ip route` (Linux), `netstat -nr` (macOS/Linux), `traceroute`/`tracert`.
    *   **Firewalls Bloqueando:** Regras de firewall (iptables, nftables, AWS Security Groups, GCP Firewall Rules) podem impedir conexões esperadas. Diagnóstico: Verificar regras, usar `telnet <host> <port>` ou `nc -zv <host> <port>` para testar conectividade em portas específicas.
    *   **Problemas de NAT:** Configuração incorreta pode impedir conexões de entrada ou saída.
*   **Em Go:**
    *   **`net` package:**
        *   `net.ParseIP`, `net.IPNet`: Manipulação de endereços IP e redes CIDR. Útil para validação de configuração, ferramentas de rede.
        *   `net.Dial`, `net.DialTimeout`, `net.Dialer`: Estabelece conexões (geralmente L4), mas depende do roteamento L3 funcionar. Timeouts aqui podem indicar problemas L3.
        *   `net.LookupIP`: Realiza consultas DNS (Layer 7) para obter IPs (Layer 3).
    *   **ICMP:** Usado por `ping`. Pode-se criar "raw sockets" em Go para enviar pacotes ICMP customizados (ex: para health checks mais avançados), usando `golang.org/x/net/icmp`.
        ```go
        // Exemplo conceitual de ping ICMP em Go
        import "golang.org/x/net/icmp"
        import "golang.org/x/net/ipv4"
        // ...
        conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
        // // ... construir e enviar mensagem echo ICMP ...
        // // ... ler resposta ...
        ```
    *   **Configuração de Infra:** Ao definir Services do tipo LoadBalancer ou Ingress no Kubernetes, ou configurar VPCs/Subnets na Cloud, você está trabalhando diretamente com conceitos L3. Ferramentas de IaC (Terraform, Pulumi) usam APIs que manipulam esses recursos.

---

## 4. Camada de Transporte (Layer 4)

*   **O que é:** Provê comunicação fim-a-fim entre processos em hosts diferentes. Lida com segmentação de dados, controle de fluxo, controle de erros e multiplexação usando números de porta. Protocolos principais: TCP (Transmission Control Protocol) e UDP (User Datagram Protocol).
*   **Relevância para SRE/SWE com Go:** Crítica. Quase toda comunicação de rede em aplicações (HTTP, gRPC, bancos de dados, etc.) usa TCP ou UDP.
    *   **TCP:**
        *   Orientado à conexão (requer handshake de 3 vias).
        *   Garante entrega ordenada e confiável dos segmentos (retransmissão em caso de perda).
        *   Controle de fluxo (evita sobrecarregar o receptor) e controle de congestionamento (evita sobrecarregar a rede).
        *   Ideal para: HTTP, gRPC, conexões de banco de dados, transferência de arquivos.
    *   **UDP:**
        *   Sem conexão (envia datagramas diretamente).
        *   Não garante entrega, ordem ou ausência de duplicatas.
        *   Baixa latência (sem handshake), menor overhead.
        *   Ideal para: DNS, streaming de vídeo/áudio, jogos online, telemetria (StatsD, logs via syslog), alguns tipos de descoberta de serviço.
    *   **Portas:** Identificam o processo/aplicação específico no host (0-65535). Ex: HTTP na porta 80, HTTPS na 443.
*   **Impacto e Diagnóstico:**
    *   **Falha no Handshake TCP:** Conexão não estabelecida. Sintoma: `net.Dial` falha com timeout ou "connection refused" (se o processo não está ouvindo na porta). Ferramenta: `telnet <host> <port>`, `nc -zv <host> <port>`.
    *   **Conexões TCP Interrompidas:** "Connection reset by peer", "broken pipe". Pode ser causado por crash do servidor, firewall no meio do caminho, timeouts.
    *   **Perda de Pacotes UDP:** Aplicação precisa lidar com isso (ex: retransmitir em nível de aplicação ou aceitar perda).
    *   **Esgotamento de Portas:** Em cenários de alta carga ou conexões de curta duração (sem `TIME_WAIT` reuse), pode-se esgotar portas efêmeras. Ferramenta: `ss -s` (Linux).
    *   **Latência:** Handshake TCP adiciona RTT (Round Trip Time). Controle de fluxo/congestionamento pode limitar throughput.
*   **Em Go:**
    *   **`net` package:** Coração da interação L4.
        *   `net.Dial("tcp", "host:port")`, `net.Listen("tcp", ":port")`: Para TCP.
        *   `net.Dial("udp", "host:port")`, `net.ListenPacket("udp", ":port")`: Para UDP.
        *   `net.Conn` interface: Abstração para conexões (TCP, Unix sockets). Métodos `Read()`, `Write()`, `Close()`, `SetDeadline()`, `SetReadDeadline()`, `SetWriteDeadline()`. **Usar deadlines é crucial** para evitar goroutines bloqueadas indefinidamente.
        *   `net.TCPConn`, `net.UDPConn`: Tipos concretos com opções específicas (e.g., `SetKeepAlive`, `SetNoDelay` para TCP).
    *   **Bibliotecas de Alto Nível:** `net/http`, `google.golang.org/grpc` usam `net` por baixo dos panos. Elas oferecem configurações de timeout específicas (e.g., `http.Server.ReadTimeout`, `http.Transport.DialContext`).
    *   **Exemplo de Timeout em Conexão TCP:**
        ```go
        package main

        import (
        	"fmt"
        	"net"
        	"time"
        )

        func main() {
        	// Tentar conectar a uma porta inexistente com timeout
        	target := "google.com:81" // Porta 81 provavelmente fechada
        	timeout := 2 * time.Second

        	conn, err := net.DialTimeout("tcp", target, timeout)
        	if err != nil {
        		// O erro pode ser um timeout ou outro problema de rede (DNS, roteamento)
        		fmt.Printf("Erro ao conectar a %s: %v", target, err)
        		// Verificar se foi especificamente um timeout
        		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
        			fmt.Println("A conexão atingiu o timeout.")
        		}
        		return
        	}
        	defer conn.Close()
        	fmt.Println("Conexão bem-sucedida (improvável neste exemplo)")
        }
        ```
    *   **Monitoramento:** Verificar estado das conexões (`netstat`, `ss`), métricas de erros de conexão, latência TCP, perda de pacotes UDP.

---

## 5. Camada de Sessão (Layer 5)

*   **O que é:** Gerencia o estabelecimento, manutenção e encerramento de sessões (diálogos) entre aplicações. Lida com controle de diálogo (quem fala quando), sincronização e recuperação de falhas em sessões longas.
*   **Relevância Prática Hoje:** As funcionalidades desta camada são frequentemente integradas nas camadas de Transporte (L4) ou Aplicação (L7). Não há protocolos L5 dedicados amplamente usados como TCP (L4) ou HTTP (L7).
*   **Analogias no Mundo Moderno/Go:**
    *   **`context.Context` em Go:** Embora não seja um protocolo de rede, o `context` serve a um propósito análogo *dentro* de uma aplicação Go para gerenciar o ciclo de vida, cancelamento e deadlines de operações (incluindo chamadas de rede). Propagar um `context` com timeout ou cancelamento através de várias chamadas de função/goroutines é similar ao controle de uma "sessão" de trabalho.
        ```go
        // Exemplo conceitual com context
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        //
        // // Passa o ctx para funções que fazem I/O (rede, DB, etc.)
        result, err := performNetworkRequest(ctx, requestData)
        if err != nil {
            if errors.Is(err, context.DeadlineExceeded) {
                fmt.Println("Operação cancelada devido a timeout do contexto (sessão)")
            }
        }
        ```
    *   **gRPC Streams:** O conceito de streams bidirecionais em gRPC (que roda sobre HTTP/2, que roda sobre TCP) permite um diálogo contínuo entre cliente e servidor, gerenciado pela biblioteca gRPC, que pode ser visto como uma forma de gerenciamento de sessão. Deadlines em gRPC também se encaixam aqui.
    *   **Keep-Alive em TCP/HTTP:** A reutilização de conexões TCP (via `Keep-Alive`) para múltiplas requisições HTTP pode ser vista como uma forma de otimização de sessão, embora gerenciada em L4/L7.
    *   **WebSockets:** Mantêm uma conexão persistente para comunicação bidirecional, estabelecendo uma sessão longa.
*   **Relevância para SRE/SWE com Go:** Indireta, mas importante conceitualmente. Entender a necessidade de gerenciar o ciclo de vida de operações complexas ou de longa duração é crucial.
    *   **Foco em Go:** Usar `context.Context` corretamente é a principal manifestação prática deste conceito para controlar timeouts, cancelamentos e propagar valores relacionados a uma "sessão" de requisição.
    *   **Diagnóstico:** Problemas relacionados podem se manifestar como operações que não terminam (sem timeout apropriado), recursos vazando (goroutines presas esperando em I/O sem cancelamento), ou falha em limpar estados após uma comunicação ser interrompida.

---

## 6. Camada de Apresentação (Layer 6)

*   **O que é:** Responsável por garantir que os dados sejam apresentados em um formato compreensível pela aplicação receptora. Lida com:
    *   **Tradução/Formatação de Dados:** Conversão entre diferentes representações de dados (ex: ASCII para EBCDIC, embora raramente relevante hoje).
    *   **Codificação/Serialização:** Transformar estruturas de dados da aplicação (structs Go) em formatos de transmissão (JSON, XML, Protobuf, Gob) e vice-versa.
    *   **Compressão de Dados:** Reduzir o tamanho dos dados para economizar largura de banda (gzip, brotli).
    *   **Criptografia/Decriptografia (TLS/SSL):** Garantir a confidencialidade e integridade dos dados em trânsito. Frequentemente associada a esta camada, embora TLS opere "entre" L4 e L7.
*   **Relevância para SRE/SWE com Go:** Muito alta. Você lida com isso constantemente.
    *   **Serialização/Desserialização:** Escolher o formato certo (JSON, Protobuf, etc.) impacta performance, tamanho dos dados e interoperabilidade.
        *   **Go:** Packages `encoding/json`, `encoding/xml`, `encoding/gob`, `google.golang.org/protobuf/proto`.
    *   **Compressão:** Configurar compressão em servidores HTTP ou usar bibliotecas como `compress/gzip`.
        *   **Go:** Middleware para `net/http` ou uso direto dos packages `compress/*`.
    *   **TLS (Transport Layer Security):** Essencial para segurança. Configurar certificados, cifras, versões de TLS.
        *   **Go:** Package `crypto/tls`. Configuração em `http.Server`, `http.Transport`, clientes gRPC/DB.
*   **Impacto e Diagnóstico:**
    *   **Erros de Serialização/Desserialização:** "Bad Request" (400) em APIs, falha ao processar mensagens. Causa: formato inválido, incompatibilidade de schema (Protobuf).
    *   **Problemas de Performance:** JSON pode ser lento para grandes volumes; Protobuf é mais eficiente. Compressão consome CPU mas economiza rede.
    *   **Erros de TLS Handshake:** Falha ao estabelecer conexão segura. Causas: certificado expirado/inválido, incompatibilidade de cifra/protocolo, problemas de configuração no cliente ou servidor. Ferramentas: `openssl s_client -connect host:port`.
    *   **Vulnerabilidades:** TLS mal configurado (versões antigas, cifras fracas) expõe a riscos de segurança.
*   **Em Go:**
    *   **Trabalhando com JSON:**
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
        	// Serialização (Go struct -> JSON)
        	req := Request{UserID: 123, Data: "payload"}
        	jsonData, err := json.Marshal(req)
        	if err != nil {
        		log.Fatalf("Erro ao serializar JSON: %v", err)
        	}
        	fmt.Printf("JSON: %s", jsonData)

        	// Desserialização (JSON -> Go struct)
        	var receivedReq Request
        	err = json.Unmarshal(jsonData, &receivedReq)
        	if err != nil {
        		log.Fatalf("Erro ao desserializar JSON: %v", err)
        	}
        	fmt.Printf("Struct recebida: %+v", receivedReq)
        }
        ```
    *   **Configurando TLS Básico em Servidor HTTP:**
        ```go
        import "net/http"
        import "log"
        
        func handler(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintf(w, "Hello, TLS!")
        }
        
        func main() {
            http.HandleFunc("/", handler)
            // Assume cert.pem e key.pem existem
            log.Println("Servidor HTTPS ouvindo na porta 8443...")
            err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
            if err != nil {
                log.Fatal("ListenAndServeTLS: ", err)
            }
        }
        ```
    *   **Monitoramento:** Métricas de erros de parse, latência de serialização, taxa de sucesso de handshake TLS.

---

## 7. Camada de Aplicação (Layer 7)

*   **O que é:** A camada mais próxima do usuário final. Fornece interfaces e protocolos que as aplicações usam para se comunicar pela rede. Define as regras para interações específicas.
*   **Protocolos Comuns:** HTTP/HTTPS, gRPC, DNS, SMTP (email), FTP (transferência de arquivos), SSH (acesso remoto seguro), MQTT (mensageria IoT), LDAP (diretório), NTP (tempo).
*   **Relevância para SRE/SWE com Go:** Máxima. É aqui que a lógica de negócio e a interação direta com outros serviços ocorrem.
    *   **Desenvolvimento:** Construção de APIs (RESTful com HTTP, RPC com gRPC), clientes para outros serviços, interação com DNS, envio de emails.
    *   **Operação (SRE):** Monitoramento de performance de aplicações (latência, taxa de erros HTTP/gRPC), health checks L7 (verificar se `/health` retorna 200 OK), configuração de Load Balancers L7, WAFs (Web Application Firewalls).
*   **Impacto e Diagnóstico:**
    *   **Erros HTTP:** Códigos 4xx (erros do cliente: Bad Request, Not Found, Unauthorized) e 5xx (erros do servidor: Internal Server Error, Bad Gateway, Service Unavailable).
    *   **Erros gRPC:** Status codes específicos (Unavailable, DeadlineExceeded, Unauthenticated).
    *   **Falhas de DNS:** Incapacidade de resolver nomes de host para IPs. Sintoma: Erro "host not found" ao tentar conectar. Ferramenta: `dig`, `nslookup`.
    *   **Latência da Aplicação:** Lógica lenta no servidor, processamento pesado, dependências lentas.
    *   **Problemas de Autenticação/Autorização:** Falha ao acessar recursos protegidos.
*   **Em Go:**
    *   **`net/http`:** Pacote padrão para construir clientes e servidores HTTP. Roteadores populares (`gorilla/mux`, `chi`, `gin-gonic`) facilitam a criação de APIs REST.
    *   **`google.golang.org/grpc`:** Biblioteca padrão para gRPC em Go.
    *   **`net` (DNS):** `net.LookupHost`, `net.LookupCNAME`, etc., para consultas DNS. Bibliotecas especializadas como `miekg/dns` para controle mais fino.
    *   **Bibliotecas Cliente:** SDKs para AWS, GCP, bancos de dados (e.g., `database/sql`), etc., operam nesta camada.
    *   **Exemplo Servidor HTTP Básico:**
        ```go
        package main

        import (
        	"fmt"
        	"log"
        	"net/http"
        	"time"
        )

        func helloHandler(w http.ResponseWriter, r *http.Request) {
        	log.Printf("Recebida requisição para: %s %s", r.Method, r.URL.Path)
        	// Simula algum processamento
        	time.Sleep(50 * time.Millisecond)
        	fmt.Fprintf(w, "Olá, mundo da aplicação!")
        }

        func healthHandler(w http.ResponseWriter, r *http.Request) {
        	// Health check L7 básico
        	w.WriteHeader(http.StatusOK)
        	fmt.Fprintf(w, "OK")
        }

        func main() {
        	mux := http.NewServeMux()
        	mux.HandleFunc("/hello", helloHandler)
        	mux.HandleFunc("/health", healthHandler) // Endpoint de health check

        	server := &http.Server{
        		Addr:         ":8080",
        		Handler:      mux,
        		ReadTimeout:  5 * time.Second,  // Proteção L4/L7
        		WriteTimeout: 10 * time.Second, // Proteção L4/L7
        		IdleTimeout:  15 * time.Second, // Proteção L4/L7
        	}

        	log.Println("Servidor HTTP ouvindo na porta 8080...")
        	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        		log.Fatalf("Erro ao iniciar servidor: %v", err)
        	}
        }
        ```
    *   **Monitoramento:** Métricas chave (RED - Rate, Errors, Duration) para cada endpoint/serviço. Tracing distribuído para seguir requisições através de múltiplos serviços. Logging detalhado.

---

### Como isso se encaixa no fluxo SWE ↔ SRE em Go

1.  **Desenvolvimento (SWE)**
    *   Você escreve código que opera primariamente em **L7** (handlers HTTP/gRPC, lógica de negócio) e **L6** (serialização JSON/Protobuf, configuração TLS).
    *   Usa bibliotecas que interagem com **L4** (`net.Dial`, `net.Listen`) e indiretamente dependem de **L3** (resolução DNS, roteamento IP).
    *   Implementa resiliência (timeouts, retries, circuit breakers) baseados em falhas observadas em **L4** e **L7**.
    *   Usa `context.Context` (**L5** análogo) para gerenciar ciclo de vida.

2.  **Operação (SRE)**
    *   Monitora a saúde e performance em **L7** (métricas de aplicação, logs, tracing) e **L4** (conexões TCP, erros).
    *   Configura health checks em **L7** (HTTP `/health`) e **L4** (TCP connect) para Load Balancers e orquestradores (Kubernetes).
    *   Diagnostica problemas usando ferramentas que inspecionam **L3** (`ping`, `traceroute`), **L4** (`telnet`, `nc`, `ss`), e **L7** (`curl`, `grpcurl`).
    *   Gerencia configuração de rede em **L3** (IPs, rotas, firewalls/Security Groups) e **L6** (certificados TLS).
    *   Responde a alertas sobre falhas que podem originar em qualquer camada, mas frequentemente se manifestam como erros em **L4** ou **L7**.

3.  **Infraestrutura em Go**
    *   Cria ferramentas CLI/agentes que usam `net` para interagir com **L3/L4** (scanners de porta, testes de conectividade).
    *   Desenvolve exporters Prometheus que coletam métricas de **L7** (e.g., `promhttp`).
    *   Escreve Operadores Kubernetes que manipulam recursos de rede como Services (**L4**), Ingress (**L7**), NetworkPolicies (**L3/L4**).
    *   Implementa lógica de descoberta de serviço baseada em DNS (**L7**) ou APIs de orquestrador.

---

**Resumo rápido:**

*   **Foco Principal (Codificação/Configuração Direta):** Layers **7 (Aplicação)**, **6 (Apresentação)**, **4 (Transporte)**, e **3 (Rede)**.
*   **Camada 4 (TCP/UDP):** Decisões críticas sobre confiabilidade vs performance, gerenciamento de conexões, timeouts. É onde "a borracha encontra a estrada" para a maioria das aplicações.
*   **Camadas 5-7:** Onde a lógica da aplicação, segurança (TLS), formato de dados e protocolos específicos residem. Go oferece excelentes bibliotecas padrão e de terceiros aqui (`net/http`, `crypto/tls`, `context`, `grpc`).
*   **Camadas 1-3:** Fundamentais para a conectividade existir, mas geralmente gerenciadas pela infraestrutura (física, cloud provider, equipe de rede) e diagnosticadas com ferramentas específicas quando problemas surgem. O SRE precisa entender como falhas aqui impactam as camadas superiores.

Compreender como essas camadas interagem ajuda a construir aplicações Go mais robustas, seguras e observáveis, e a diagnosticar problemas de forma mais eficaz, seja no código (SWE) ou na operação (SRE).

