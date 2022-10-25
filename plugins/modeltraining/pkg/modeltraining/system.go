package modeltraining

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rancher/opni/apis"
	"github.com/rancher/opni/pkg/auth/cluster"
	"github.com/rancher/opni/pkg/resources"
	util "github.com/rancher/opni/pkg/util/k8sutil"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/cenkalti/backoff"
	backoffv2 "github.com/lestrrat-go/backoff/v2"
	"github.com/nats-io/nats.go"
	opensearch "github.com/opensearch-project/opensearch-go"
)

func newNatsConnection() (*nats.Conn, error) {
	natsURL := os.Getenv("NATS_SERVER_URL")
	natsSeedPath := os.Getenv("NKEY_SEED_FILENAME")

	opt, err := nats.NkeyOptionFromSeed(natsSeedPath)
	if err != nil {
		return nil, err
	}

	retryBackoff := backoff.NewExponentialBackOff()
	return nats.Connect(
		natsURL,
		opt,
		nats.MaxReconnects(-1),
		nats.CustomReconnectDelay(
			func(i int) time.Duration {
				if i == 1 {
					retryBackoff.Reset()
				}
				return retryBackoff.NextBackOff()
			},
		),
	)
}

func getOpensearchCredentials(ctx context.Context) (username string, password string) {
	id, ok := cluster.AuthorizedIDFromIncomingContext(ctx)
	if !ok {
		return "admin", "admin"
	}
	labels := map[string]string{
		resources.OpniClusterID: id,
	}
	k8sClient, err := util.NewK8sClient(util.ClientOptions{Scheme: apis.NewScheme()})
	if err != nil {
		return "admin", "admin"
	}
	secrets := &corev1.SecretList{}
	namespace, ok := os.LookupEnv("POD_NAMESPACE")
	if !ok {
		namespace = "opni-cluster-system"
	}
	if err := k8sClient.List(ctx, secrets, client.InNamespace(namespace), client.MatchingLabels(labels)); err != nil {
		return "admin", "admin"
	}

	if len(secrets.Items) != 1 {
		return "admin", "admin"
	}

	username = secrets.Items[0].Name
	password = string(secrets.Items[0].Data["password"])
	return username, password
}

func getOpensearchEndpoint() string {

	namespace, ok := os.LookupEnv("POD_NAMESPACE")
	if !ok {
		namespace = "opni-cluster-system"
	}
	endpoint := fmt.Sprintf("https://opni-opensearch-svc.%s.svc:9200", namespace)
	return endpoint
}

func (s *ModelTrainingPlugin) newOpensearchConnection() (*opensearch.Client, error) {
	esEndpoint := getOpensearchEndpoint()
	s.Logger.Debug(esEndpoint)
	esUsername, esPassword := getOpensearchCredentials(s.ctx)
	s.Logger.Debug(esUsername)

	retrier := backoffv2.Exponential(
		backoffv2.WithMaxRetries(0),
		backoffv2.WithMinInterval(5*time.Second),
		backoffv2.WithMaxInterval(1*time.Minute),
		backoffv2.WithMultiplier(1.1),
	)
	b := retrier.Start(s.ctx)
	var (
		client *opensearch.Client
	)
	for backoffv2.Continue(b) {
		osClient, err := opensearch.NewClient(opensearch.Config{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Addresses: []string{esEndpoint},
			Username:  esUsername,
			Password:  esPassword,
		})
		if err == nil {
			client = osClient
			break
		}
	}
	return client, nil
}
