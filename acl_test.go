package libfw

import (
	"testing"

	"github.com/RTradeLtd/go-libp2p-fw/pb"
	"github.com/ipfs/go-datastore"
	dssync "github.com/ipfs/go-datastore/sync"
)

func Test_NewListManager(t *testing.T) {
	ds := dssync.MutexWrap(datastore.NewMapDatastore())
	manager, err := NewListManager(datastore.NewKey("lol"), ds)
	if err != nil {
		t.Fatal(err)
	}
	manager.StoreRule(newACL(), 0, false)
}

func newACL() *pb.ACL {
	return &pb.ACL{
		FilterType: pb.FILTER_TYPE_PEER_ID,
		Direction:  pb.DIRECTION_IN_OUT,
	}
}
