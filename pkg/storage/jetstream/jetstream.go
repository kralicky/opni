package jetstream

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/nats-io/nats.go"
	corev1 "github.com/rancher/opni/pkg/apis/core/v1"
	"github.com/rancher/opni/pkg/config/v1beta1"
	"github.com/rancher/opni/pkg/logger"
	"github.com/rancher/opni/pkg/storage"
	"go.uber.org/zap"
)

const (
	tokensBucket       = "tokens"
	clustersBucket     = "clusters"
	keyringsBucket     = "keyrings"
	rolesBucket        = "roles"
	roleBindingsBucket = "rolebindings"
	dynamicBucket      = "dynamic"
)

type JetStreamStore struct {
	JetStreamStoreOptions

	nc *nats.Conn
	js nats.JetStreamContext

	kv struct {
		Tokens       nats.KeyValue
		Clusters     nats.KeyValue
		Keyrings     nats.KeyValue
		Roles        nats.KeyValue
		RoleBindings nats.KeyValue
	}
	logger *zap.SugaredLogger
}

var _ storage.Backend = (*JetStreamStore)(nil)

type JetStreamStoreOptions struct {
	BucketPrefix string
}

type JetStreamStoreOption func(*JetStreamStoreOptions)

func (o *JetStreamStoreOptions) apply(opts ...JetStreamStoreOption) {
	for _, op := range opts {
		op(o)
	}
}

func NewJetStreamStore(ctx context.Context, conf *v1beta1.JetStreamStorageSpec, opts ...JetStreamStoreOption) (*JetStreamStore, error) {
	options := JetStreamStoreOptions{
		BucketPrefix: "gateway",
	}
	options.apply(opts...)

	lg := logger.New(logger.WithLogLevel(zap.WarnLevel)).Named("jetstream")

	nkeyOpt, err := nats.NkeyOptionFromSeed(conf.NkeySeedPath)
	if err != nil {
		return nil, err
	}
	nc, err := nats.Connect(conf.Endpoint, nkeyOpt, nats.MaxReconnects(-1))
	if err != nil {
		return nil, err
	}
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	store := &JetStreamStore{
		JetStreamStoreOptions: options,
		nc:                    nc,
		js:                    js,
		logger:                lg,
	}

	store.kv.Tokens = store.upsertBucket(tokensBucket)
	store.kv.Clusters = store.upsertBucket(clustersBucket)
	store.kv.Keyrings = store.upsertBucket(keyringsBucket)
	store.kv.Roles = store.upsertBucket(rolesBucket)
	store.kv.RoleBindings = store.upsertBucket(roleBindingsBucket)

	return store, nil
}

func (s *JetStreamStore) upsertBucket(name string) nats.KeyValue {
	bucketName := fmt.Sprintf("%s-%s", s.BucketPrefix, name)
	kv, err := s.js.KeyValue(bucketName)
	if err != nil {
		if errors.Is(err, nats.ErrBucketNotFound) {
			kv, err = s.js.CreateKeyValue(&nats.KeyValueConfig{
				Bucket: bucketName,
				Description: fmt.Sprintf("Opni %s %s Store",
					strcase.ToCamel(s.BucketPrefix),
					strcase.ToCamel(name)),
				Storage:  nats.FileStorage,
				History:  64,
				Replicas: 1,
			})
		}
	}
	if err != nil {
		s.logger.With(
			"bucket", bucketName,
			zap.Error(err),
		).Panic("failed to create bucket")
	}
	return kv
}

func (s *JetStreamStore) KeyringStore(prefix string, ref *corev1.Reference) storage.KeyringStore {
	return &jetstreamKeyringStore{
		kv:     s.kv.Keyrings,
		ref:    ref,
		prefix: prefix,
	}
}

func (s *JetStreamStore) KeyValueStore(prefix string) storage.KeyValueStore {
	// sanitize bucket name
	prefix = strings.ReplaceAll(strings.ReplaceAll(prefix, "/", "-"), ".", "_")
	bucket := s.upsertBucket(fmt.Sprintf("%s-%s", dynamicBucket, prefix))
	return &jetstreamKeyValueStore{
		kv: bucket,
	}
}
