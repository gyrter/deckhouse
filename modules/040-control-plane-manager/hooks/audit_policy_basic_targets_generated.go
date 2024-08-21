// Code generated by "tools/audit_policy.go" DO NOT EDIT.
package hooks

var auditPolicyBasicNamespaces = []string{
	"d8-admission-policy-engine",
	"d8-ceph-csi",
	"d8-cert-manager",
	"d8-chrony",
	"d8-cloud-instance-manager",
	"d8-cloud-provider-aws",
	"d8-cloud-provider-azure",
	"d8-cloud-provider-gcp",
	"d8-cloud-provider-openstack",
	"d8-cloud-provider-vcd",
	"d8-cloud-provider-vsphere",
	"d8-cloud-provider-yandex",
	"d8-cloud-provider-zvirt",
	"d8-cni-cilium",
	"d8-cni-flannel",
	"d8-cni-simple-bridge",
	"d8-delivery",
	"d8-descheduler",
	"d8-flant-integration",
	"d8-ingress-nginx",
	"d8-istio",
	"d8-keepalived",
	"d8-l2-load-balancer",
	"d8-local-path-provisioner",
	"d8-log-shipper",
	"d8-metallb",
	"d8-monitoring",
	"d8-multitenancy-manager",
	"d8-network-gateway",
	"d8-okmeter",
	"d8-openvpn",
	"d8-operator-prometheus",
	"d8-operator-trivy",
	"d8-pod-reloader",
	"d8-runtime-audit-engine",
	"d8-snapshot-controller",
	"d8-static-routing-manager",
	"d8-system",
	"d8-upmeter",
	"d8-user-authn",
	"d8-user-authz",
	"kube-system",
}
var auditPolicyBasicServiceAccounts = []string{
	"system:serviceaccount:d8-cert-manager:cainjector",
	"system:serviceaccount:d8-cert-manager:cert-manager",
	"system:serviceaccount:d8-cert-manager:webhook",
	"system:serviceaccount:d8-cloud-instance-manager:caps-controller-manager",
	"system:serviceaccount:d8-cloud-instance-manager:cluster-autoscaler",
	"system:serviceaccount:d8-cloud-instance-manager:early-oom",
	"system:serviceaccount:d8-cloud-instance-manager:fencing-agent",
	"system:serviceaccount:d8-cloud-instance-manager:machine-controller-manager",
	"system:serviceaccount:d8-cloud-instance-manager:node-group",
	"system:serviceaccount:d8-cloud-instance-manager:registry-packages-proxy",
	"system:serviceaccount:d8-cloud-provider-aws:cloud-controller-manager",
	"system:serviceaccount:d8-cloud-provider-aws:cloud-data-discoverer",
	"system:serviceaccount:d8-cloud-provider-aws:node-termination-handler",
	"system:serviceaccount:d8-cloud-provider-azure:cloud-controller-manager",
	"system:serviceaccount:d8-cloud-provider-azure:cloud-data-discoverer",
	"system:serviceaccount:d8-cloud-provider-gcp:cloud-controller-manager",
	"system:serviceaccount:d8-cloud-provider-gcp:cloud-data-discoverer",
	"system:serviceaccount:d8-cloud-provider-openstack:cloud-controller-manager",
	"system:serviceaccount:d8-cloud-provider-openstack:cloud-data-discoverer",
	"system:serviceaccount:d8-cloud-provider-vcd:capcd-controller-manager",
	"system:serviceaccount:d8-cloud-provider-vcd:cloud-controller-manager",
	"system:serviceaccount:d8-cloud-provider-vcd:cloud-data-discoverer",
	"system:serviceaccount:d8-cloud-provider-vsphere:cloud-controller-manager",
	"system:serviceaccount:d8-cloud-provider-vsphere:cloud-data-discoverer",
	"system:serviceaccount:d8-cloud-provider-yandex:cloud-controller-manager",
	"system:serviceaccount:d8-cloud-provider-yandex:cloud-data-discoverer",
	"system:serviceaccount:d8-cloud-provider-yandex:cloud-metrics-exporter",
	"system:serviceaccount:d8-cloud-provider-zvirt:capz-controller-manager",
	"system:serviceaccount:d8-cloud-provider-zvirt:cloud-controller-manager",
	"system:serviceaccount:d8-cloud-provider-zvirt:cloud-data-discoverer",
	"system:serviceaccount:d8-cni-cilium:agent",
	"system:serviceaccount:d8-cni-cilium:dashboard",
	"system:serviceaccount:d8-cni-cilium:egress-gateway-agent",
	"system:serviceaccount:d8-cni-cilium:operator",
	"system:serviceaccount:d8-cni-cilium:relay",
	"system:serviceaccount:d8-cni-cilium:safe-agent-updater",
	"system:serviceaccount:d8-cni-cilium:ui",
	"system:serviceaccount:d8-cni-flannel:cni-flannel",
	"system:serviceaccount:d8-cni-simple-bridge:cni-simple-bridge",
	"system:serviceaccount:d8-descheduler:descheduler",
	"system:serviceaccount:d8-flant-integration:pricing",
	"system:serviceaccount:d8-ingress-nginx:ingress-nginx",
	"system:serviceaccount:d8-ingress-nginx:kruise",
	"system:serviceaccount:d8-istio:alliance-ingressgateway",
	"system:serviceaccount:d8-istio:alliance-metadata-exporter",
	"system:serviceaccount:d8-istio:cni",
	"system:serviceaccount:d8-istio:ingress-gateway-controller",
	"system:serviceaccount:d8-istio:kiali",
	"system:serviceaccount:d8-istio:multicluster-api-proxy",
	"system:serviceaccount:d8-l2-load-balancer:controller",
	"system:serviceaccount:d8-l2-load-balancer:speaker",
	"system:serviceaccount:d8-local-path-provisioner:local-path-provisioner",
	"system:serviceaccount:d8-log-shipper:log-shipper",
	"system:serviceaccount:d8-metallb:controller",
	"system:serviceaccount:d8-metallb:speaker",
	"system:serviceaccount:d8-monitoring:aggregating-proxy",
	"system:serviceaccount:d8-monitoring:alertmanager-internal",
	"system:serviceaccount:d8-monitoring:alerts-receiver",
	"system:serviceaccount:d8-monitoring:cert-exporter",
	"system:serviceaccount:d8-monitoring:control-plane-proxy",
	"system:serviceaccount:d8-monitoring:ebpf-exporter",
	"system:serviceaccount:d8-monitoring:events-exporter",
	"system:serviceaccount:d8-monitoring:extended-monitoring-exporter",
	"system:serviceaccount:d8-monitoring:grafana",
	"system:serviceaccount:d8-monitoring:image-availability-exporter",
	"system:serviceaccount:d8-monitoring:kube-state-metrics",
	"system:serviceaccount:d8-monitoring:loki",
	"system:serviceaccount:d8-monitoring:monitoring-ping",
	"system:serviceaccount:d8-monitoring:node-exporter",
	"system:serviceaccount:d8-monitoring:prometheus",
	"system:serviceaccount:d8-monitoring:trickster",
	"system:serviceaccount:d8-multitenancy-manager:multitenancy-manager",
	"system:serviceaccount:d8-okmeter:okmeter",
	"system:serviceaccount:d8-openvpn:openvpn",
	"system:serviceaccount:d8-operator-prometheus:operator-prometheus",
	"system:serviceaccount:d8-operator-trivy:operator-trivy",
	"system:serviceaccount:d8-operator-trivy:report-updater",
	"system:serviceaccount:d8-pod-reloader:pod-reloader",
	"system:serviceaccount:d8-runtime-audit-engine:k8s-metacollector",
	"system:serviceaccount:d8-runtime-audit-engine:runtime-audit-engine",
	"system:serviceaccount:d8-snapshot-controller:snapshot-controller",
	"system:serviceaccount:d8-static-routing-manager:agent",
	"system:serviceaccount:d8-system:deckhouse",
	"system:serviceaccount:d8-system:documentation",
	"system:serviceaccount:d8-system:network-policy-engine",
	"system:serviceaccount:d8-system:terraform-auto-converger",
	"system:serviceaccount:d8-system:terraform-state-exporter",
	"system:serviceaccount:d8-system:webhook-handler",
	"system:serviceaccount:d8-upmeter:smoke-mini",
	"system:serviceaccount:d8-upmeter:upmeter",
	"system:serviceaccount:d8-upmeter:upmeter-agent",
	"system:serviceaccount:d8-user-authn:dex",
	"system:serviceaccount:d8-user-authz:webhook",
	"system:serviceaccount:kube-system:d8-control-plane-manager",
	"system:serviceaccount:kube-system:d8-kube-dns",
	"system:serviceaccount:kube-system:d8-kube-proxy",
	"system:serviceaccount:kube-system:d8-node-local-dns",
	"system:serviceaccount:kube-system:d8-vertical-pod-autoscaler-admission-controller",
	"system:serviceaccount:kube-system:d8-vertical-pod-autoscaler-recommender",
	"system:serviceaccount:kube-system:d8-vertical-pod-autoscaler-updater",
	"system:serviceaccount:kube-system:stale-dns-connections-cleaner",
}
