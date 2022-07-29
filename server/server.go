package main

import (
	"crypto/sha256"
	"encoding/json"
	"faraway/pow"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"rsc.io/quote"
)

const (
	ERROR_EXIT_CODE = 1
)

func init() {
	log.SetPrefix("<Faraway Gaming Server>: ")
}

func main() {
	port := flag.Uint("port", 8080, "TCP Port Number for the gaming server")
	flag.Parse()

	li, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleNewConnection(conn)
	}
}

const (
	MAX_FAILED_ATTEMPTS_DDOS = 3
	CONN_DEADLINE_SEC        = 5
)

type Session struct {
	clientAddress net.Addr
	serverAddress net.Addr
	lastAttempt   int64
	attempts      int8
}

var sessions map[string]*Session = make(map[string]*Session, 0)

func handleNewConnection(conn net.Conn) {
	// store client details : IP and num of retries before block
	client := conn.RemoteAddr()
	server := conn.LocalAddr()
	log.Printf(strings.Repeat(">", 40))
	log.Printf("New connection: [%s] %s -> %s\n", client.Network(), client.String(), server.String())
	session_key := fmt.Sprintf("%s:%s", client.Network(), client.String())
	session, ok := sessions[session_key]
	if !ok {
		// first attempt
		session = &Session{
			clientAddress: client,
			serverAddress: server,
			lastAttempt:   time.Now().UnixNano(),
			attempts:      1,
		}
		sessions[session_key] = session
	} else {
		if sessions[session_key].attempts >= MAX_FAILED_ATTEMPTS_DDOS {
			// already blacklisted
			log.Printf("Client %s has been blacklisted. Ignore...\n", client.String())
			conn.Close()
			return
		}
	}

	timestampHash := sha256.Sum256([]byte(fmt.Sprintf("%x", session.lastAttempt)))

	connection := pow.NewConnection(session.clientAddress.String(), session.serverAddress.String(), uint64(session.lastAttempt))

	// create pow challenge
	block := pow.NewBlock(0, timestampHash, connection)
	// write pow challange: hash and complexity
	pow_paylaod, pow_err := json.Marshal(block)
	if pow_err != nil {
		log.Fatalln(pow_err)
	}
	sent_n, sent_err := conn.Write(pow_paylaod)
	if sent_err != nil || sent_n < len(pow_paylaod) {
		log.Fatalln(sent_err)
	}
	// receive pow answer
	recv_buf := make([]byte, len(pow_paylaod)*4)
	recv_n, recv_err := conn.Read(recv_buf)
	if recv_err != nil {
		log.Fatalln(recv_err)
	} else if recv_n < len(pow_paylaod) {
		log.Fatalf("Expected payload size %d, instead received %d", len(pow_paylaod), recv_n)
	}
	// verify pow
	recv_block := new(pow.Block)
	json.Unmarshal(recv_buf[:recv_n], recv_block)
	if !pow.ValidProof(recv_block.Nonce(), recv_block.Timestamp(), recv_block.Connection(), pow.POW_DIFFICULTY) {
		log.Printf("Invalid proof of work from %s. Close connection\n", client.String())
		sessions[session_key].attempts += 1
		conn.Close()
		return
	}
	// serve quote or disconnect if pow invalid
	msg := quote.Go()
	conn.Write([]byte(msg))
	defer conn.Close()

	fmt.Printf("***CONNECTION with %s TERMINATED***\n", session.clientAddress.String())
}
