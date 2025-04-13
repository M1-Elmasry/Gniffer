
# 🕵️‍♂️ Gniffer (Go Sniffer)

`Gniffer` is a lightweight Go tool that passively monitors incoming TCP connections to specific ports on your system **without interfering with the actual services**. Think of it as a simple packet sniffer that listens silently, logging connection attempts like SSH, HTTP, or any custom port you choose.

---

## ⚙️ Features

- ✅ Passive sniffing using `gopacket` and `pcap`
- 🎯 Monitor only **specified TCP ports**
- 📋 Logs source IP, destination IP, port, and TCP flags (SYN/ACK)
- 🧩 Modular structure with separate `sniffer` and `logger` packages
- 🧠 Designed for learning, showcasing Go skills, or real-world usage

---

## 🗂️ Project Structure

```
portsniffer/
├── main.go              # CLI interface and app entrypoint
├── sniffer/
│   └── sniffer.go       # Packet capturing and filtering logic
├── logger/
│   └── logger.go        # Connection logging logic
├── go.mod               # Go module definition
```

---

## 🚀 Usage

### 1. Install Dependencies

```bash
go mod tidy
```

### 2. Run the Sniffer
Use the following command to list your network interfaces:

```bash
ip link
```
Common examples: `eth0`, `wlan0`, `ens33`, etc.  
  

```bash
sudo go run Gniffer -ports=22,80,443 -iface=eth0 --log=log.txt
```

### 3. Example Output

```
Sniffing on eth0 | Ports: 22,80,443
[12:35:19] 192.168.1.20 ➜ 192.168.1.5:22 | SYN=true ACK=false
[12:35:21] 192.168.1.11 ➜ 192.168.1.5:443 | SYN=true ACK=false
...
```


---

## 📦 Build Binary

```bash
go build -o gniffer
sudo ./gniffer -ports=22,80,443 -iface=eth0 --log=log.txt
```

---

## 📌 Notes

- Only monitors **TCP traffic**
- Focuses on **incoming packets**, not outbound
- Doesn’t block or interfere with existing services
