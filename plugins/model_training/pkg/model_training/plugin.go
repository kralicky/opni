package model_training

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	opensearch "github.com/opensearch-project/opensearch-go"
	managementv1 "github.com/rancher/opni/pkg/apis/management/v1"
	"github.com/rancher/opni/pkg/logger"
	managementext "github.com/rancher/opni/pkg/plugins/apis/apiextensions/management"
	"github.com/rancher/opni/pkg/plugins/apis/system"
	"github.com/rancher/opni/pkg/plugins/meta"
	"github.com/rancher/opni/pkg/util"
	"github.com/rancher/opni/pkg/util/future"
	"github.com/rancher/opni/plugins/model_training/pkg/apis/model_training"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ModelTrainingPlugin struct {
	model_training.UnsafeModelTrainingServer
	system.UnimplementedSystemPluginClient
	ctx            context.Context
	Logger         *zap.SugaredLogger
	k8sClient      client.Client
	osClient       future.Future[*opensearch.Client]
	natsConnection future.Future[*nats.Conn]
	kv             future.Future[nats.KeyValue]
}

func (s *ModelTrainingPlugin) UseManagementAPI(api managementv1.ManagementClient) {
	lg := s.Logger
	lg.Info("querying management API...")
	var list *managementv1.APIExtensionInfoList
	for {
		var err error
		list, err = api.APIExtensions(context.Background(), &emptypb.Empty{})
		if err != nil {
			log.Fatal(err)
		}
		if len(list.Items) == 0 {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		break
	}
	for _, ext := range list.Items {
		lg.Info("found API extension service", "name", ext.ServiceDesc.GetName())
	}
	nc, _ := newNatsConnection()
	mgr, _ := nc.JetStream()
	keyValue, err := mgr.CreateKeyValue(&nats.KeyValueConfig{
		Bucket:      "os-workload-aggregation",
		Description: "Storing aggregation of workload logs from Opensearch.",
	})
	if err != nil {
		os.Exit(1)
	}
	client, err := newOpensearchConnection()
	if err != nil {
		os.Exit(1)
	}
	s.natsConnection.Set(nc)
	s.kv.Set(keyValue)
	s.osClient.Set(client)
	go s.run_aggregation()

}

var _ model_training.ModelTrainingServer = (*ModelTrainingPlugin)(nil)

func Scheme(ctx context.Context) meta.Scheme {
	scheme := meta.NewScheme()
	p := &ModelTrainingPlugin{
		Logger:         logger.NewPluginLogger().Named("model_training"),
		natsConnection: future.New[*nats.Conn](),
		kv:             future.New[nats.KeyValue](),
		osClient:       future.New[*opensearch.Client](),
	}
	scheme.Add(system.SystemPluginID, system.NewPlugin(p))
	scheme.Add(managementext.ManagementAPIExtensionPluginID,
		managementext.NewPlugin(util.PackService(&model_training.ModelTraining_ServiceDesc, p)))
	return scheme
}
