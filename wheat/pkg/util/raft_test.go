package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.etcd.io/etcd/client/pkg/v3/types"
	"go.etcd.io/etcd/server/v3/etcdserver/api/rafthttp"
	"go.etcd.io/raft/v3"
	"go.uber.org/zap"
)

func TestRaft(t *testing.T) {
	assert := require.New(t)

	// raft transport

	zlogger := zap.NewExample()
	assert.NotNil(zlogger)

	node1Id := 1
	transport := &rafthttp.Transport{
		Logger: zlogger,
		ID:     types.ID(node1Id),
	}
	assert.NotNil(transport)

	// test raft status
	rpeers := make([]raft.Peer, 3)
	raft.StartNode(nil, rpeers)
}
