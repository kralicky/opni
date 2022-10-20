package agent

import (
	"context"
	"sync"

	"github.com/rancher/opni/pkg/health"
	"github.com/rancher/opni/plugins/topology/pkg/apis/node"
	"github.com/rancher/opni/plugins/topology/pkg/apis/remote"
	"go.uber.org/zap"
)

type TopologyStreamer struct {
	logger     *zap.SugaredLogger
	conditions health.ConditionTracker

	remoteWriteClientMu sync.Mutex
	remoteWriteClient   remote.RemoteTopologyClient
}

func NewTopologyStreamer(ct health.ConditionTracker, lg *zap.SugaredLogger) *TopologyStreamer {
	return &TopologyStreamer{
		logger:     lg,
		conditions: ct,
	}
}

func (s *TopologyStreamer) SetRemoteWriteClient(client remote.RemoteTopologyClient) {
	s.remoteWriteClientMu.Lock()
	defer s.remoteWriteClientMu.Unlock()
	s.remoteWriteClient = client
}

func (s *TopologyStreamer) Run(ctx context.Context, spec *node.TopologyCapabilitySpec) error {
	lg := s.logger
	if spec == nil {
		lg.With("stream", "topology").Warn("no topology capability spec provided, setting defaults")

		// set some sensible defaults
	}

	// blocking action
	for {
		select {
		case <-ctx.Done():
			lg.With(
				zap.Error(ctx.Err()),
			).Warn("topology stream closing")
			return nil
		}
	}
}