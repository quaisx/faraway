# POW guard against DDOS
Simplified demo version which does not contain a fully guarded against DDoS atacks code. It simply shows how a _pow_ request can be used to protect TCP servers from abusing client requests.
---

Goal: create a server/client with PoW to guard against DDoS attacks

This code was written for demo purpose - to demonstate how a _pow_ challenge may be used.

##Details
Design and implement _“Word of Wisdom”_ tcp server with _PoW_ challenge. 
 - TCP server should be protected from DDOS attacks with the [Proof of Work](https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used. 
 - The choice of the POW algorithm - simplified for demo purpose. 
 - After Prof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes. 
 - Docker files for the server and for the client to quickly build docker images
 
 
 By default the server runs on the port 8080 and the client picks an ephemeral port and connects to the server.
 
 
 PoW Wisdom Quote Server:
 ![PoW Quote Server](/images/server.jpeg)
 
 
 PoW Wisdom Quote Client:
 ![PoW Quote Client](/images/client.jpeg)
 
 
