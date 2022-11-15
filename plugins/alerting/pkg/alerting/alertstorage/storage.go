package alertstorage

import (
	"context"
	"encoding/json"
	"errors"
	"path"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rancher/opni/pkg/alerting/shared"
	alertingv1 "github.com/rancher/opni/pkg/apis/alerting/v1"
	"github.com/rancher/opni/pkg/logger"
	"github.com/rancher/opni/pkg/storage"
	"github.com/rancher/opni/pkg/util/future"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const conditionPrefix = "/alerting/conditions"
const endpointPrefix = "/alerting/endpoints"

type StorageAPIs struct {
	Conditions storage.KeyValueStoreT[*alertingv1.AlertCondition]
	Endpoints  storage.KeyValueStoreT[*alertingv1.AlertEndpoint]
	// key : conditionId
	// value  : AgentTracker
	SystemTrackerStorage future.Future[nats.KeyValue]
}

// Responsible for anything related to persistent storage
// external to AlertManager
type StorageNode struct {
	*StorageNodeOptions
	agentTrackerMu sync.Mutex
}
type StorageNodeOptions struct {
	Logger  *zap.SugaredLogger
	timeout time.Duration
	storage future.Future[*StorageAPIs]
}

func NewStorageNode(opts ...StorageNodeOption) *StorageNode {
	options := &StorageNodeOptions{
		timeout: time.Second * 60,
	}
	options.storage = future.New[*StorageAPIs]()
	options.apply(opts...)
	if options.Logger == nil {
		options.Logger = logger.NewPluginLogger().Named("alerting-storage-node")
	}
	return &StorageNode{
		StorageNodeOptions: options,
	}
}

func (s *StorageNode) SetSystemTrackerStorage(kv nats.KeyValue) {
	s.storage.Get().SystemTrackerStorage.Set(kv)
}

type StorageNodeOption func(*StorageNodeOptions)

func WithStorageTimeout(timeout time.Duration) StorageNodeOption {
	return func(o *StorageNodeOptions) {
		o.timeout = timeout
	}
}

func WithLogger(lg *zap.SugaredLogger) StorageNodeOption {
	return func(o *StorageNodeOptions) {
		o.Logger = lg
	}
}

func WithStorage(storage *StorageAPIs) StorageNodeOption {
	return func(o *StorageNodeOptions) {
		o.storage = future.New[*StorageAPIs]()
		o.storage.Set(storage)
	}
}

func (s *StorageNodeOptions) apply(opts ...StorageNodeOption) {
	for _, opt := range opts {
		opt(s)
	}
}

func (s *StorageNode) CreateCondition(ctx context.Context, conditionId string, condition *alertingv1.AlertCondition) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	storage, err := s.storage.GetContext(ctxTimeout)
	if err != nil {
		return err
	}
	return storage.Conditions.Put(ctx, path.Join(conditionPrefix, conditionId), condition)
}

func (s *StorageNode) GetCondition(ctx context.Context, conditionId string) (*alertingv1.AlertCondition, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	storage, err := s.storage.GetContext(ctxTimeout)
	if err != nil {
		return nil, err
	}
	return storage.Conditions.Get(ctx, path.Join(conditionPrefix, conditionId))
}

func (s *StorageNode) ListConditions(ctx context.Context) ([]*alertingv1.AlertCondition, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	storage, err := s.storage.GetContext(ctxTimeout)
	if err != nil {
		return nil, err
	}
	items, err := list(ctx, storage.Conditions, conditionPrefix)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *StorageNode) ListWithKeysConditions(ctx context.Context) ([]string, []*alertingv1.AlertCondition, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	storage, err := s.storage.GetContext(ctxTimeout)
	if err != nil {
		return nil, nil, err
	}
	keys, items, err := listWithKeys(ctx, storage.Conditions, conditionPrefix)
	if err != nil {
		return nil, nil, err
	}
	return keys, items, nil
}

func (s *StorageNode) UpdateCondition(ctx context.Context, conditionId string, newCondition *alertingv1.AlertCondition) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	storage, err := s.storage.GetContext(ctxTimeout)
	if err != nil {
		return err
	}
	_, err = storage.Conditions.Get(ctx, path.Join(conditionPrefix, conditionId))
	if err != nil {
		return shared.WithNotFoundErrorf("condition to update '%s' not found : %s", conditionId, err)
	}
	return storage.Conditions.Put(ctx, path.Join(conditionPrefix, conditionId), newCondition)
}

func (s *StorageNode) DeleteCondition(ctx context.Context, conditionId string) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	storage, err := s.storage.GetContext(ctxTimeout)
	if err != nil {
		return err
	}
	_, err = storage.Conditions.Get(ctx, path.Join(conditionPrefix, conditionId))
	if err != nil {
		return shared.WithNotFoundErrorf("condition to delete '%s' not found : %s", conditionId, err)
	}
	return storage.Conditions.Delete(ctx, path.Join(conditionPrefix, conditionId))
}

func (s *StorageNode) CreateEndpoint(ctx context.Context, endpointId string, endpoint *alertingv1.AlertEndpoint) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	storage, err := s.storage.GetContext(ctxTimeout)
	if err != nil {
		return err
	}
	return storage.Endpoints.Put(ctx, path.Join(endpointPrefix, endpointId), endpoint)
}

func (s *StorageNode) GetEndpoint(ctx context.Context, endpointId string) (*alertingv1.AlertEndpoint, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	storage, err := s.storage.GetContext(ctxTimeout)
	if err != nil {
		return nil, err
	}
	return storage.Endpoints.Get(ctx, path.Join(endpointPrefix, endpointId))
}

func (s *StorageNode) ListEndpoints(ctx context.Context) ([]*alertingv1.AlertEndpoint, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	storage, err := s.storage.GetContext(ctxTimeout)
	if err != nil {
		return nil, err
	}
	items, err := list(ctx, storage.Endpoints, endpointPrefix)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *StorageNode) ListWithKeysEndpoints(ctx context.Context) ([]string, []*alertingv1.AlertEndpoint, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	storage, err := s.storage.GetContext(ctxTimeout)
	if err != nil {
		return nil, nil, err
	}
	keys, items, err := listWithKeys(ctxTimeout, storage.Endpoints, endpointPrefix)
	if err != nil {
		return nil, nil, err
	}
	return keys, items, nil
}

func (s *StorageNode) UpdateEndpoint(ctx context.Context, endpointId string, newEndpoint *alertingv1.AlertEndpoint) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	storage, err := s.storage.GetContext(ctxTimeout)
	if err != nil {
		return err
	}
	_, err = storage.Endpoints.Get(ctxTimeout, path.Join(endpointPrefix, endpointId))
	if err != nil {
		return shared.WithNotFoundErrorf("condition to update '%s' not found : %s", endpointId, err)
	}
	return storage.Endpoints.Put(ctx, path.Join(endpointPrefix, endpointId), newEndpoint)
}

func (s *StorageNode) DeleteEndpoint(ctx context.Context, endpointId string) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	storage, err := s.storage.GetContext(ctxTimeout)
	if err != nil {
		return err
	}
	_, err = storage.Endpoints.Get(ctx, path.Join(endpointPrefix, endpointId))
	if err != nil {
		return shared.WithNotFoundErrorf("condition to delete '%s' not found : %s", endpointId, err)
	}
	return storage.Endpoints.Delete(ctx, path.Join(endpointPrefix, endpointId))
}

func (s *StorageNode) CreateIncidentTracker(
	ctx context.Context,
	conditionId string,
	initialValue InternalIncidentStep) error {
	// TODO : generic constructor
	// FIXME: input the type ourselves
	t := HandleInternalIncidentType(shared.AgentDisconnectStorageType)
	t.New(conditionId, initialValue)
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	storage, err := s.storage.GetContext(ctxTimeout)
	if err != nil {
		return err
	}
	sts, err := storage.SystemTrackerStorage.GetContext(ctxTimeout)
	if err != nil {
		return err
	}
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	_, err = sts.Create(conditionId, data)
	return err
}

func (s *StorageNode) GetIncidentTracker(ctx context.Context, conditionId string) (InternalIncident, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	storage, err := s.storage.GetContext(ctxTimeout)
	if err != nil {
		return nil, err
	}
	sts, err := storage.SystemTrackerStorage.GetContext(ctxTimeout)
	if err != nil {
		return nil, err
	}
	entry, err := sts.Get(conditionId)
	if err != nil {
		return nil, err
	}
	var st *AgentIncident
	err = json.Unmarshal(entry.Value(), &st)
	if err != nil {
		return nil, err
	}
	return st, nil
}

func (s *StorageNode) GetActiveWindowsFromIncidentTracker(
	ctx context.Context,
	conditionId string,
	start,
	end *timestamppb.Timestamp,
) ([]*alertingv1.ActiveWindow, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	incident, err := s.GetIncidentTracker(ctxTimeout, conditionId)
	if err != nil {
		return nil, err
	}
	res := []*alertingv1.ActiveWindow{}
	steps := incident.GetSteps()
	if len(steps) == 0 {
		return res, nil
	}
	risingEdge := true
	for _, step := range steps {
		if step.GetTimestamp() == nil {
			continue
		}
		if step.IsFiring() && risingEdge {
			res = append(res, &alertingv1.ActiveWindow{
				Start: step.GetTimestamp(),
				End:   timestamppb.Now(), // overwritten if it is found later
				Type:  alertingv1.TimelineType_Timeline_Alerting,
			})
			risingEdge = false
		} else if !step.IsFiring() && !risingEdge {
			res[len(res)-1].End = step.GetTimestamp()
			risingEdge = true
		}
	}
	pruneIdx := 0
	for _, window := range res {
		if window.End.AsTime().Before(start.AsTime()) {
			pruneIdx++
		}
	}
	res = slices.Delete(res, 0, pruneIdx)
	return res, nil
}

func (s *StorageNode) ListIncidentTrackers(ctx context.Context) ([]InternalIncident, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	storage, err := s.storage.GetContext(ctxTimeout)
	if err != nil {
		return nil, err
	}
	sts, err := storage.SystemTrackerStorage.GetContext(ctxTimeout)
	if err != nil {
		return nil, err
	}
	ids, err := sts.Keys()
	if err != nil {
		return nil, err
	}
	res := make([]InternalIncident, 0)
	for _, id := range ids {
		entry, err := sts.Get(id)
		if err != nil {
			return nil, err
		}
		var st *AgentIncident
		err = json.Unmarshal(entry.Value(), &st)
		if err != nil {
			continue
		}
		if st == nil {
			panic("st should not unmarshal to nil")
		}
		err = json.Unmarshal(entry.Value(), &st)
		if err != nil {
			continue
		}
		res = append(res, st)
	}
	return res, nil
}

func (s *StorageNode) AddToIncidentTracker(ctx context.Context, conditionId string, updateValue InternalIncidentStep) error {
	s.agentTrackerMu.Lock()
	defer s.agentTrackerMu.Unlock()
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	storage, err := s.storage.GetContext(ctxTimeout)
	if err != nil {
		return err
	}
	sts, err := storage.SystemTrackerStorage.GetContext(ctxTimeout)
	if err != nil {
		return err
	}
	var agentEntry nats.KeyValueEntry
	entry, err := sts.Get(conditionId)
	if errors.Is(err, nats.ErrKeyNotFound) {
		err := s.CreateIncidentTracker(ctx, conditionId, updateValue)
		if err != nil {
			return err
		}
		entry, err := sts.Get(conditionId)
		if err != nil {
			return err
		}
		agentEntry = entry
	} else if err != nil {
		return err
	} else {
		agentEntry = entry
	}
	var prev AgentIncident
	err = json.Unmarshal(agentEntry.Value(), &prev)
	if err != nil {
		return err
	}
	if prev.isEquivalentState(updateValue) { // prevent filling up K,V with duplicate states
		return nil
	}
	prev.AddStep(updateValue)
	data, err := json.Marshal(prev)
	if err != nil {
		return err
	}
	_, err = sts.Put(conditionId, data)
	return err
}

func (s *StorageNode) DeleteIncidentTracker(ctx context.Context, conditionId string) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	storage, err := s.storage.GetContext(ctxTimeout)
	if err != nil {
		return err
	}
	sts, err := storage.SystemTrackerStorage.GetContext(ctxTimeout)
	if err != nil {
		return err
	}
	return sts.Delete(conditionId)
}
