package libfw

// TODO(postables): flesh out FireSwarm before enabling
// import "github.com/libp2p/go-libp2p-core/network"
// var _ network.Network = (*FireSwarm)(nil)

// FireSwarm is a firewalled libp2p swarm
type FireSwarm struct {
	lm *ListManager
}
