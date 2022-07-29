package main

import (
	"encoding/json"
	"faraway/pow"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"time"
)

func init() {
	log.SetPrefix("<Faraway Gaming Client>: ")
}

func checkErr(e error) {
	if e != nil {
		log.Println(e)
		panic(e)
	}
}

func main() {
	server_address := flag.String("server-address", "localhost", "IPv4 address for the gaming server")
	server_port := flag.Uint("server-port", 8080, "TCP Port Number for the gaming server")

	flag.Parse()

	server_ips, e := net.LookupHost(*server_address)
	checkErr(e)

	var server_ip string

	for _, ip_addr := range server_ips {
		ip := net.ParseIP(ip_addr)
		ipv4 := ip.To4()
		if ipv4 == nil {
			continue
		}
		server_ip = fmt.Sprintf("%s:%d", ipv4, *server_port)
		break
	}
	log.Println(strings.Repeat(">", 40))
	d := net.Dialer{Timeout: 2 * time.Second}
	log.Println(" Connect to server", server_ip)
	conn, err := d.Dial("tcp", server_ip)
	checkErr(err)
	defer conn.Close()
	log.Println(" Connected to", server_ip)
	block := new(pow.Block)
	block_payload, err := json.Marshal(block)
	checkErr(err)
	block_buf := make([]byte, len(block_payload)*4)
	n, err := conn.Read(block_buf)
	// need to unmarchal result to Block and perform proof of work
	log.Printf(" Received %d bytes. Trying to unmarshal...", n)
	json.Unmarshal(block_buf[:n], block)
	log.Printf(" Unmarshaleled block -> \n%s\n", block)
	nonce := pow.ProofOfWork(block.Timestamp(), block.Connection(), pow.POW_DIFFICULTY)
	log.Println(" Proof of work: nonce=", nonce)
	my_pow_block := pow.NewBlock(nonce, block.Timestamp(), block.Connection())
	valid := pow.ValidProof(my_pow_block.Nonce(), my_pow_block.Timestamp(), my_pow_block.Connection(), pow.POW_DIFFICULTY)
	log.Printf(" Validate POW locally. Result %t", valid)
	if valid {
		pow_paylaod, pow_err := json.Marshal(my_pow_block)
		if pow_err != nil {
			log.Println(" Failed misarably with", pow_err)
			return
		}
		sent_n, sent_err := conn.Write(pow_paylaod)
		if sent_err != nil || sent_n < len(pow_paylaod) {
			log.Println(" Failed on write to server with", sent_err)
		}

		read_n, read_err := ioutil.ReadAll(conn)
		checkErr(read_err)
		if len(read_n) > 0 {
			log.Println(string(read_n))
		}
	} else {
		log.Println(" I failed to calculate POW")
		return
	}
	log.Println(strings.Repeat("<", 40))
}
