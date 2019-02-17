# Emulating-TCP_UDP-Client - Protocol Fuzzing

This is a TCP or UDP client that can emulate a network stack client calling a TCP or UDP server. It prompts on the screen on the message you want to send to the Server. 

config.json provides base configuration for the emulated client, in json format. Current Config parameter takes in

1. Address - the address of the server, either IP address or hostname or FQDN
2. Protocol - TCP or UDP
3. Port - TCP or UDP port
A sampe of config.json is provided as example. Values have to be embedded within the quotes "". You can bake the config into the code but for configuration flexibility, added an external file config.json for configuration parameters.

You can test the GO program without building or installing the binaries, by using go run ./tcpclient.go

Added the func init() which gets loaded after all the imports, for declaration of the config.json variables to be used. The main func main() is loaded after this, client runs in the main package.

The tcp/udp emulated client does similar to what you can do with NetCat utilities except, it does not close the connection (for TCP) after the message is send. NetCat closes the connection immediately after sending the message.

Sample command of NC
Establishing tcp port 80 session with content “test out the server” to a remote server
echo -n "test out the server" | nc -v www.akamai.com 80

versus running tcpclient.go with punted text content

tcpdump on tcp port 80 with snaplength of 1500 for 20 packet, verbose (vvv)
sudo tcpdump -i en0 -vvv -s 1500 -X -c 20 port 80

The differences are highlighted in the tcpdump packet capture. You can view the pcap using wireshark. The connection emulate port 80 but it is not actually sending any HTTP. With tcpclient.go, because the session is preserve, the remote server responded with a HTTP error. A HTTP server will expect HTTP messages. Some HTTP servers from testing, closes the connection upon receiving an invalid message, while others could still be waiting. The former could work to remove exhausting TCP sockets/memory and avoids invalid TCP session dangling. 

This demonstrate it is possible to send messages using standard wellknown ports, a firewall could let it through without inspecting further if it is a valid packet. Only parsing it deeper will reveal.

tcpclient.go and tcpserver.go can be used together and you will be able to see the connection flows between the client and the server for educational purposes, all within the same machine.

Numerous help refer and infer including Stackoverflow, @Sean repositories and articles on Internet. Learn about the func init() and how to use it with func main()
