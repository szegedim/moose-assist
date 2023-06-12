# Moose - Studio Quality Audio Call
## This is an example for an audio call package implemented in Golang and JavaScript

```
Licensed under Creative Commons CC0.

To the extent possible under law, the author(s) have dedicated all copyright and related and
neighboring rights to this software to the public domain worldwide.
This software is distributed without any warranty.
You should have received a copy of the CC0 Public Domain Dedication along with this software.
If not, see <https:#creativecommons.org/publicdomain/zero/1.0/legalcode>.
```

## Abstract

This is an example audio call package implemented in golang and html/js.
It allows a peer to peer connection.


It reduces the risks of privacy.
You run a simple container yourself in your cloud or on premises environment.

## Security

The connection is over HTML websockets.
We advise you to follow your organization's security best practices.
It does not do TLS. You can secure it by your own standards and infrastructure.

We generate en end-to-end encrypted channel.
Also, the administrator sets an apikey for the line with APIKEY environment variable.

You can use NGINX to secure the socket. https://www.nginx.com/blog/using-free-ssltls-certificates-from-lets-encrypt-with-nginx
This is important. By having a central way to manage TLS parameters, you do not expose yourself to upgrade issues, if microservices need security updates. Only nginx has read permission to your private keys. This is an important best practice.

## Protocol

We do not use webrtc to give you full control over the packets. We basically pass a 200kHz Pulse Code Modulated audio of 16 bits per sample. **The bandwidth requirement is about 3.4Mbps**. The power lies in the simplicity. You do not need to worry about supporting different hardware, etc. This protocol is always supported.

## Design considerations

The Moose audio call package follows these principles
- Simplicity supports the use of microcontrollers on one or both sides. **The client code is 200 lines**
- Remove any special codec and hardware requirements for the broadest support. We support plain PCM.
- Use studio quality 200 kHz and 16 bits for each pulse sample. Studios originally used 192 kHz audio sampling and 24 bits.
- Simplify the codebase, so that you can run the server on your own.
- Let your Ops team deal with security. This helps to follow organizational standards and it reduces long term support costs.
- Authorization is established by an `apikey` that is shared by participants privately over their own environment (email, corporate chat).
- API key should be at least 16 characters containing a-zA-Z and -. This helps to prevent 'moosebombing' attacks with 53^16 variants.
- High bandwidth connections over 50Mbps can easily handle PCM audio with the lowest latency. There is no extra buffering needed for convolutional codecs like MP3 ensuring the lowest codec latency possible. (Martucci, 1994)
- We implement websockets privately to avoid security risks of third party components. There are **no dependencies**.

## Requirements

Please use a connection of at least 7Mbps on each client and the server to run this example.

## Usage

2. Establish at a TLS 1.3 or better closure using the documentation below.

```dockerfile
https://www.nginx.com/blog/using-free-ssltls-certificates-from-lets-encrypt-with-nginx
```

2. Build te docker file.
```
docker build . -t mooseassist.example.com
```


3. Run the server side in docker. Set APIKEY to a unique cryptographically strong random value.

```
docker run -d -p 7777:7777 -e APIKEY='8f711a8f43a3d6fbf9c367a8cd1b68b14db2781273c56e206040ecd31761f9d3' mooseassist.example.com
```

4. Connect to the server from the peer, who hosts the meeting. This will generate a leaf key for end to end encryption

```dockerfile
https://mooseassist.example.com/?apikey=8f711a8f43a3d6fbf9c367a8cd1b68b14db2781273c56e206040ecd31761f9d3#generate_leaf
```

5. Connect the second peer using the same URL.

```dockerfile
https://mooseassist.example.com/?apikey=8f711a8f43a3d6fbf9c367a8cd1b68b14db2781273c56e206040ecd31761f9d3#leaf_sywICvC1tnE7D5IOdPahv0NzQC-C3mpNLhluhEEW0vA
```

6. Example run logs

```
connection received
connection received
-32768.00 -32263.82 32418.00
-32768.00 -32237.32 32347.00
```

## References

Martucci, S. A. (May 1994). "Symmetric convolution and the discrete sine and cosine transforms". IEEE Transactions on Signal Processing.

## Contact and support

Do you need servicing from us?

- support agreement
- warranty, and reliability guarantee
- patent licensing and third party patent licensing
- security certification
- legal and privacy compliance certificates
- integration (Kubernetes, Helm, etc.)

**Please contact miklos.szegedi@eper.io at Schmied Enterprises LLC.**
