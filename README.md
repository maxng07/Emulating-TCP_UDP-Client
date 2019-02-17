# Emulating-TCP_UDP-Client - Protocol Fuzzing

This is a TCP or UDP client that can emulate a network stack client calling a TCP or UDP server. It prompts on the screen on the message you want to send to the Server. 

config.json provides base configuration for the client in json format. Current Config takes in
Address - the address of the server, either IP address or hostname or FQDN 
Protocol - TCP or UDP
Port - TCP or UDP port
A sampe of config.json is provided for example. Values have to be embedded with quotes "". You can bake the config into the code but for configuration flexibility, added an external file config.json for configuration.

You can test the GO program without building or installing the binaries, using go run ./tcpclient.go

Numerous help refer and infer including Stackoverflow, @Sean repositories and articles on Internet. Learn about the func init() and how to use it with func main()

The tcpclient does similar to what you can do with NetCat utilities except, it does not close the connection (for TCP) after the message is send. NetCat closes the connection.

Sample command of NC 
Establishing tcp port 80 session with content “test out the server”
echo -n "test out the server" | nc -v www.akamai.com 80

versus tcpclient.go
The differences are in the packet capture 

tcpdump on tcp port 80 with snaplength of 1500 for 20 packet, verbose (vvv)
sudo tcpdump -i en0 -vvv -s 1500 -X -c 20 port 80

tcpclient.go and tcpserver.go can be used together and you will be able to see the connections flow between the client and the server for educational purposes, all within the same machine.
