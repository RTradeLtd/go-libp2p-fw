package libfw

import (
	"log"
	"os"
	"sync"

	"fmt"

	"strings"

	"io/ioutil"

	"context"
	"time"

	"github.com/RTradeLtd/go-libp2p-fw/pb"
	"github.com/ipfs/go-datastore"
	"github.com/multiformats/go-multihash"
)

// acl.go provides acl management tools

// ListManager allows managing the internal list
type ListManager struct {
	list       *pb.List
	store      datastore.Batching
	mux        *sync.RWMutex
	forceFlush chan bool
}

// NewListManager initializes a list manager, and loads the ruleset into memory
func NewListManager(key datastore.Key, store datastore.Batching) (*ListManager, error) {
	var list *pb.List
	value, err := store.Get(key)
	if err != nil && err != datastore.ErrNotFound {
		return nil, err
	} else if err != nil && err == datastore.ErrNotFound {
		list = new(pb.List)
		list.Acls = make(map[int64]*pb.ACL, 0)
	} else {
		if err := list.Unmarshal(value); err != nil {
			return nil, err
		}
	}
	return &ListManager{list, store, &sync.RWMutex{}, make(chan bool, 1)}, nil
}

// StoreRule is used to store an ACL within our list at the given priority.
// if force is set to true, and a rule is already occupying the priority slot
// then we overwrite it.
func (lm *ListManager) StoreRule(acl *pb.ACL, priority int64, force bool) error {
	if !lm.canWrite(priority, force) {
		return fmt.Errorf("unable to write rule with priortiy %v", priority)
	}
	lm.list.Acls[priority] = acl
	lm.forceFlush <- true
	return nil
}

// Start is used to start the internal list flusher
func (lm *ListManager) Start(ctx context.Context) {
	lm.flusher(ctx)
}

func (lm *ListManager) canWrite(priority int64, force bool) bool {
	lm.mux.RLock()
	var canWrite bool
	if lm.list.Acls[priority] != nil && !force {
		canWrite = false
	} else if lm.list.Acls[priority] != nil && force {
		canWrite = true
	} else if lm.list.Acls[priority] == nil {
		canWrite = true
	} else {
		canWrite = false
	}
	lm.mux.RUnlock()
	return canWrite
}

func (lm *ListManager) flusher(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 60)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-lm.forceFlush:
			lm.flush()
		case <-ticker.C:
			lm.flush()
		}
	}
}

func (lm *ListManager) flush() {
	lm.mux.RLock()
	listBytes, err := lm.list.Marshal()
	if err != nil {
		log.Println("failed to marshal list", err)
		return
	}
	lm.mux.RUnlock()
	sum, err := multihash.Sum(listBytes, multihash.Names[strings.ToLower("md5")], 16)
	if err != nil {
		log.Println("failed to sum list bytes", err)
		return
	}
	if err := lm.store.Put(datastore.NewKey(sum.String()), listBytes); err != nil {
		log.Println("failed to store updated list", err)
		return
	}
	if err := ioutil.WriteFile("acl.key", []byte(sum.String()), os.FileMode(0600)); err != nil {
		log.Println("failed to write acl datastore key to disk", err)
		return
	}
}
