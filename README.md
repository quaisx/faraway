# faraway technical assignment
Task: create a server/client with PoW to guard against DDoS attacks
##Details
Design and implement _“Word of Wisdom”_ tcp server. 
 - TCP server should be protected from DDOS attacks with the [Proof of Work](https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used. 
 - The choice of the POW algorithm should be explained. 
 - After Prof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes. 
 - Docker file should be provided both for the server and for the client that solves the POW challenge
 
 By default the server runs on the port 8080 and the client picks an ephemeral port and connects to the server.
