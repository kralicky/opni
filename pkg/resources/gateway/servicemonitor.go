package gateway

import (
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/rancher/opni/pkg/resources"
	"github.com/rancher/opni/pkg/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *Reconciler) serviceMonitor() resources.Resource {
	publicSvcLabels := resources.NewGatewayLabels()
	publicSvcLabels["service-type"] = "public"
	svcMonitor := &monitoringv1.ServiceMonitor{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "opni-gateway",
			Namespace: r.gw.Namespace,
			Labels:    resources.NewGatewayLabels(),
		},
		Spec: monitoringv1.ServiceMonitorSpec{
			Selector: metav1.LabelSelector{
				MatchLabels: publicSvcLabels,
			},
			NamespaceSelector: monitoringv1.NamespaceSelector{
				MatchNames: []string{r.gw.Namespace},
			},
			Endpoints: []monitoringv1.Endpoint{
				{
					TargetPort: util.Pointer(intstr.FromInt(8086)),
					Path:       "/metrics",
					Scheme:     "http",
				},
			},
		},
	}
	controllerutil.SetOwnerReference(r.gw, svcMonitor, r.client.Scheme())
	return resources.Present(svcMonitor)
}
