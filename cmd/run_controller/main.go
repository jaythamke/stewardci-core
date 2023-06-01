package main

import (
	"flag"
	"time"

	"github.com/SAP/stewardci-core/pkg/k8s"
	"github.com/SAP/stewardci-core/pkg/metrics"
	"github.com/SAP/stewardci-core/pkg/runctl"
	"github.com/SAP/stewardci-core/pkg/signals"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	klog "k8s.io/klog/v2"
	"knative.dev/pkg/system"
)

const (
	// resyncPeriod is the period between full resyncs performed
	// by the controller.
	resyncPeriod = 30 * time.Second

	// metricsPort is the TCP port number to be used by the metrics
	// HTTP server.
	metricsPort = 9090

	// errorExitCode is the exit code sent if error occurs during
	// startup.
	errorExitCode = 1
)

var (
	kubeconfig              string
	burst, qps, threadiness int

	heartbeatInterval time.Duration
	heartbeatLogging  bool
	heartbeatLogLevel int

	k8sAPIRequestTimeout time.Duration
)

func init() {
	klog.InitFlags(nil)

	flag.StringVar(
		&kubeconfig,
		"kubeconfig",
		"",
		"The path to a kubeconfig file configuring access to the Kubernetes cluster."+
			" If not specified or empty, assume running in-cluster.",
	)
	flag.IntVar(
		&qps,
		"qps",
		5,
		"The queries per seconds (QPS) for Kubernetes API client-side rate limiting.",
	)
	flag.IntVar(
		&burst,
		"burst",
		10,
		"The size of the burst bucket for Kubernetes API client-side rate limiting.",
	)
	flag.IntVar(
		&threadiness,
		"threadiness",
		2,
		"The maximum number of reconciliations performed by the controller in parallel.",
	)
	flag.DurationVar(
		&heartbeatInterval,
		"heartbeat-interval",
		1*time.Minute,
		"The interval of controller heartbeats.",
	)
	flag.BoolVar(
		&heartbeatLogging,
		"heartbeat-logging",
		true,
		"Whether controller heartbeats should be logged.",
	)
	flag.IntVar(
		&heartbeatLogLevel,
		"heartbeat-log-level",
		3,
		"The log level to be used for controller heartbeats.",
	)
	flag.DurationVar(
		&k8sAPIRequestTimeout,
		"k8s-api-request-timeout",
		15*time.Minute,
		"The maximum length of time to wait before giving up on a server request. A value of zero means no timeout.",
	)

	flag.Parse()
}

func main() {
	defer klog.Flush()

	system.Namespace() // ensure that namespace is set in environment

	var config *rest.Config
	var err error

	if kubeconfig == "" {
		klog.InfoS("In cluster")
		config, err = rest.InClusterConfig()
		if err != nil {
			klog.ErrorS(err, "Failed to load kubeconfig. Hint: You can use parameter '-kubeconfig' for local testing")
			timeoutAndExit(errorExitCode)
		}
	} else {
		klog.InfoS("Outside cluster")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			klog.ErrorS(err, "Failed to create kubeconfig from command line flag", "flag", "-kubeconfig", "path", kubeconfig)
			timeoutAndExit(errorExitCode)
		}
	}

	klog.V(3).Infof("Create Factory (resync period: %s, QPS: %d, burst: %d, k8s-api-request-timeout: %s)", resyncPeriod.String(), qps, burst, k8sAPIRequestTimeout.String())
	config.QPS = float32(qps)
	config.Burst = burst
	config.Timeout = k8sAPIRequestTimeout
	factory := k8s.NewClientFactory(config, resyncPeriod)

	klog.V(2).Infof("Provide metrics on http://0.0.0.0:%d/metrics", metricsPort)
	metrics.StartServer(metricsPort)

	klog.V(3).Infof("Create Controller")
	controllerOpts := runctl.ControllerOpts{
		HeartbeatInterval: heartbeatInterval,
	}
	if heartbeatLogging {
		tmp := klog.Level(heartbeatLogLevel)
		controllerOpts.HeartbeatLogLevel = &tmp
	}
	controller := runctl.NewController(factory, controllerOpts)

	klog.V(3).Infof("Create Signal Handlers")
	stopCh := signals.SetupShutdownSignalHandler()
	signals.SetupThreadDumpSignalHandler()

	klog.V(2).Infof("Start Informer")
	factory.StewardInformerFactory().Start(stopCh)
	factory.TektonInformerFactory().Start(stopCh)

	klog.V(2).Infof("Run controller (threadiness=%d)", threadiness)
	if err = controller.Run(threadiness, stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}

func timeoutAndExit(exitCode int) {
	klog.FlushAndExit(klog.ExitFlushTimeout, exitCode)
}
