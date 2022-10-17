module github.com/SAP/stewardci-core

go 1.18

require (
	github.com/benbjohnson/clock v1.3.0
	github.com/davecgh/go-spew v1.1.1
	github.com/ghodss/yaml v1.0.0
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.3.0
	github.com/lithammer/dedent v1.1.0
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.13.0
	github.com/tektoncd/pipeline v0.40.2
	go.uber.org/zap v1.23.0
	gopkg.in/yaml.v2 v2.4.0
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.23.9
	k8s.io/apimachinery v0.23.9
	k8s.io/client-go v1.5.2
	k8s.io/klog/v2 v2.70.2-0.20220707122935-0990e81f1a8f
	knative.dev/pkg v0.0.0-20220818004048-4a03844c0b15
)

require (
	cloud.google.com/go/compute v1.5.0 // indirect
	contrib.go.opencensus.io/exporter/ocagent v0.7.1-0.20200907061046-05415f1de66d // indirect
	contrib.go.opencensus.io/exporter/prometheus v0.4.0 // indirect
	github.com/PuerkitoBio/purell v1.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blendle/zapdriver v1.3.1 // indirect
	github.com/census-instrumentation/opencensus-proto v0.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/emicklei/go-restful v2.16.0+incompatible // indirect
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/evanphx/json-patch/v5 v5.6.0 // indirect
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/go-containerregistry v0.8.1-0.20220216220642-00c59d91847c // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/openzipkin/zipkin-go v0.3.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/prometheus/statsd_exporter v0.22.4 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/net v0.0.0-20221014081412-f15817d10f9b // indirect
	golang.org/x/oauth2 v0.0.0-20220411215720-9780585627b5 // indirect
	golang.org/x/sync v0.0.0-20220601150217-0de741cfad7f // indirect
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.8 // indirect
	golang.org/x/time v0.0.0-20220224211638-0e9765cccd65 // indirect
	gomodules.xyz/jsonpatch/v2 v2.2.0 // indirect
	google.golang.org/api v0.70.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220303160752-862486edd9cc // indirect
	google.golang.org/grpc v1.44.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/kube-openapi v0.0.0-20220124234850-424119656bbf // indirect
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9 // indirect
	sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.23.9 // kubernetes-1.23.9
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.23.9
	k8s.io/apimachinery => k8s.io/apimachinery v0.23.9 // kubernetes-1.23.9
	k8s.io/apiserver => k8s.io/apiserver v0.23.9
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.23.9
	k8s.io/client-go => k8s.io/client-go v0.23.9 // kubernetes-1.23.9
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.23.9
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.23.9
	k8s.io/code-generator => k8s.io/code-generator v0.23.9
	k8s.io/component-base => k8s.io/component-base v0.23.9
	k8s.io/component-helpers => k8s.io/component-helpers v0.23.9
	k8s.io/controller-manager => k8s.io/controller-manager v0.23.9
	k8s.io/cri-api => k8s.io/cri-api v0.23.9
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.23.9
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.23.9
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.23.9
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.23.9
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.23.9
	k8s.io/kubectl => k8s.io/kubectl v0.23.9
	k8s.io/kubelet => k8s.io/kubelet v0.23.9
	k8s.io/kubernetes => k8s.io/kubernetes v1.23.9
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.23.9
	k8s.io/metrics => k8s.io/metrics v0.23.9
	k8s.io/mount-utils => k8s.io/mount-utils v0.23.9
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.23.9 // kubernetes-1.23.9
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.23.9
	knative.dev/pkg => knative.dev/pkg v0.0.0-20220818004048-4a03844c0b15 // release-1.7
)
