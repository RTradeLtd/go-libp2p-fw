# go-libp2p-fw

`libfw` is an experimental go-libp2p firewalled swarm implementation. The end result should be a go-libp2p swarm, as a drop in replacement for `go-libp2p-swarm`, with the ability for fine-grained control over traffic.

Currently traffic rules are processed on a first-match wins basis. This allows for the ordering of higher priority rules that will trigger first. Rules start off at number 0, and increase in count. A rule with priority 10, will be matched before a rule with priority 0

# TODOs

* Enable different types of filters:
  * PeerID
  * Protocol ID
  * IPv4, IPv6 Addresses
* Enable flow control:
  * Inbound
  * Outbound
  * Inbound + Outbound
  * Relay
