package metrics

import (
	"strings"

	"github.com/ContaAzul/hystrix_exporter/hystrix"
	"github.com/apex/log"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	//
	// command gauges:
	//

	// LatencyTotal prometheus gauge
	latencyTotal = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_latency_total",
		Help: "latencies total",
	}, []string{"cluster", "name", "group", "statistic"})

	// LatencyExecute prometheus gauge
	latencyExecute = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_latency_Execute",
		Help: "latencies execute",
	}, []string{"cluster", "name", "group", "statistic"})

	// RollingCountCollapsedRequests prometheus gauge
	rollingCountCollapsedRequests = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_command_rollingCountCollapsedRequests",
		Help: "rollingCountCollapsedRequests",
	}, []string{"cluster", "name", "group"})

	// RollingCountShortCircuited prometheus gauge
	rollingCountShortCircuited = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_command_rollingCountShortCircuited",
		Help: "rollingCountShortCircuited",
	}, []string{"cluster", "name", "group"})

	// RollingCountThreadPoolRejected prometheus gauge
	rollingCountThreadPoolRejected = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_command_rollingCountThreadPoolRejected",
		Help: "rollingCountThreadPoolRejected",
	}, []string{"cluster", "name", "group"})

	// RollingCountFallbackEmit prometheus gauge
	rollingCountFallbackEmit = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_command_rollingCountFallbackEmit",
		Help: "rollingCountFallbackEmit",
	}, []string{"cluster", "name", "group"})

	// RollingCountSuccess prometheus gauge
	rollingCountSuccess = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_command_rollingCountSuccess",
		Help: "rollingCountSuccess",
	}, []string{"cluster", "name", "group"})

	// RollingCountTimeout prometheus gauge
	rollingCountTimeout = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_command_rollingCountTimeout",
		Help: "rollingCountTimeout",
	}, []string{"cluster", "name", "group"})

	// RollingCountFailure prometheus gauge
	rollingCountFailure = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_command_rollingCountFailure",
		Help: "rollingCountFailure",
	}, []string{"cluster", "name", "group"})

	// RollingCountExceptionsThrown prometheus gauge
	rollingCountExceptionsThrown = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_command_rollingCountExceptionsThrown",
		Help: "rollingCountExceptionsThrown",
	}, []string{"cluster", "name", "group"})

	// CircuitOpen prometheus gauge
	circuitOpen = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_command_circuit_open",
		Help: "circuit open, 1 means true",
	}, []string{"cluster", "name", "group"})

	//
	// thread pool gauges:
	//

	// ThreadPoolCurrentCorePoolSize prometheus gauge
	threadPoolCurrentCorePoolSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_threadpool_currentCorePoolSize",
		Help: "currentCorePoolSize",
	}, []string{"cluster", "name"})

	// ThreadPoolCurrentLargestPoolSize prometheus gauge
	threadPoolCurrentLargestPoolSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_threadpool_currentLargestPoolSize",
		Help: "currentLargestPoolSize",
	}, []string{"cluster", "name"})

	// ThreadPoolCurrentActiveCount prometheus gauge
	threadPoolCurrentActiveCount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_threadpool_currentActiveCount",
		Help: "currentActiveCount",
	}, []string{"cluster", "name"})

	// ThreadPoolCurrentMaximumPoolSize prometheus gauge
	threadPoolCurrentMaximumPoolSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_threadpool_currentMaximumPoolSize",
		Help: "currentMaximumPoolSize",
	}, []string{"cluster", "name"})

	// ThreadPoolCurrentQueueSize prometheus gauge
	threadPoolCurrentQueueSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_threadpool_currentQueueSize",
		Help: "currentQueueSize",
	}, []string{"cluster", "name"})

	// ThreadPoolCurrentTaskCount prometheus gauge
	threadPoolCurrentTaskCount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_threadpool_currentTaskCount",
		Help: "currentTaskCount",
	}, []string{"cluster", "name"})

	// ThreadPoolCurrentCompletedTaskCount prometheus gauge
	threadPoolCurrentCompletedTaskCount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_threadpool_currentCompletedTaskCount",
		Help: "currentCompletedTaskCount",
	}, []string{"cluster", "name"})

	// ThreadPoolRollingMaxActiveThreads prometheus gauge
	threadPoolRollingMaxActiveThreads = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_threadpool_rollingMaxActiveThreads",
		Help: "rollingMaxActiveThreads",
	}, []string{"cluster", "name"})

	// ThreadPoolRollingCountCommandRejections prometheus gauge
	threadPoolRollingCountCommandRejections = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_threadpool_rollingCountCommandRejections",
		Help: "rollingCountCommandRejections",
	}, []string{"cluster", "name"})

	// ThreadPoolReportingHosts prometheus gauge
	threadPoolReportingHosts = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_threadpool_reportingHosts",
		Help: "reportingHosts",
	}, []string{"cluster", "name"})

	// ThreadPoolCurrentPoolSize prometheus gauge
	threadPoolCurrentPoolSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_threadpool_currentPoolSize",
		Help: "currentPoolSize",
	}, []string{"cluster", "name"})

	// ThreadPoolRollingCountThreadsExecuted prometheus gauge
	threadPoolRollingCountThreadsExecuted = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_threadpool_rollingCountThreadsExecuted",
		Help: "rollingCountThreadsExecuted",
	}, []string{"cluster", "name"})
)

// MustRegister registers all metrics against a prometheus registerer
func MustRegister(registerer prometheus.Registerer) {
	registerer.MustRegister(
		latencyTotal,
		latencyExecute,
		rollingCountCollapsedRequests,
		rollingCountShortCircuited,
		rollingCountThreadPoolRejected,
		rollingCountFallbackEmit,
		rollingCountSuccess,
		rollingCountTimeout,
		rollingCountFailure,
		rollingCountExceptionsThrown,
		circuitOpen,

		threadPoolCurrentCorePoolSize,
		threadPoolCurrentLargestPoolSize,
		threadPoolCurrentActiveCount,
		threadPoolCurrentMaximumPoolSize,
		threadPoolCurrentQueueSize,
		threadPoolCurrentTaskCount,
		threadPoolCurrentCompletedTaskCount,
		threadPoolRollingMaxActiveThreads,
		threadPoolRollingCountCommandRejections,
		threadPoolReportingHosts,
		threadPoolCurrentPoolSize,
		threadPoolRollingCountThreadsExecuted,
	)
}

// ReportCommand reports metrics of a command
func ReportCommand(cluster string, data hystrix.Data) {
	log.WithField("cluster", cluster).WithField("data", data).Debug("reporting")

	var name = strings.ToLower(data.Name)
	var group = strings.ToLower(data.Group)

	latencyTotal.WithLabelValues(cluster, name, group, "0").Set(data.LatencyTotal.L0)
	latencyTotal.WithLabelValues(cluster, name, group, "25").Set(data.LatencyTotal.L25)
	latencyTotal.WithLabelValues(cluster, name, group, "50").Set(data.LatencyTotal.L50)
	latencyTotal.WithLabelValues(cluster, name, group, "75").Set(data.LatencyTotal.L75)
	latencyTotal.WithLabelValues(cluster, name, group, "90").Set(data.LatencyTotal.L90)
	latencyTotal.WithLabelValues(cluster, name, group, "95").Set(data.LatencyTotal.L95)
	latencyTotal.WithLabelValues(cluster, name, group, "99").Set(data.LatencyTotal.L99)
	latencyTotal.WithLabelValues(cluster, name, group, "99.5").Set(data.LatencyTotal.L995)
	latencyTotal.WithLabelValues(cluster, name, group, "100").Set(data.LatencyTotal.L100)
	latencyTotal.WithLabelValues(cluster, name, group, "mean").Set(data.LatencyTotalMean)
	latencyExecute.WithLabelValues(cluster, name, group, "0").Set(data.LatencyExecute.L0)
	latencyExecute.WithLabelValues(cluster, name, group, "25").Set(data.LatencyExecute.L25)
	latencyExecute.WithLabelValues(cluster, name, group, "50").Set(data.LatencyExecute.L50)
	latencyExecute.WithLabelValues(cluster, name, group, "75").Set(data.LatencyExecute.L75)
	latencyExecute.WithLabelValues(cluster, name, group, "90").Set(data.LatencyExecute.L90)
	latencyExecute.WithLabelValues(cluster, name, group, "95").Set(data.LatencyExecute.L95)
	latencyExecute.WithLabelValues(cluster, name, group, "99").Set(data.LatencyExecute.L99)
	latencyExecute.WithLabelValues(cluster, name, group, "99.5").Set(data.LatencyExecute.L995)
	latencyExecute.WithLabelValues(cluster, name, group, "100").Set(data.LatencyExecute.L100)
	latencyExecute.WithLabelValues(cluster, name, group, "mean").Set(data.LatencyExecuteMean)

	rollingCountCollapsedRequests.WithLabelValues(cluster, name, group).Set(data.RollingCountCollapsedRequests)
	rollingCountShortCircuited.WithLabelValues(cluster, name, group).Set(data.RollingCountShortCircuited)
	rollingCountThreadPoolRejected.WithLabelValues(cluster, name, group).Set(data.RollingCountThreadPoolRejected)
	rollingCountFallbackEmit.WithLabelValues(cluster, name, group).Set(data.RollingCountFallbackEmit)
	rollingCountSuccess.WithLabelValues(cluster, name, group).Set(data.RollingCountSuccess)
	rollingCountTimeout.WithLabelValues(cluster, name, group).Set(data.RollingCountTimeout)
	rollingCountFailure.WithLabelValues(cluster, name, group).Set(data.RollingCountFailure)
	rollingCountExceptionsThrown.WithLabelValues(cluster, name, group).Set(data.RollingCountExceptionsThrown)

	circuitOpen.WithLabelValues(cluster, name, group).Set(boolToFloat64(data.IsCircuitBreakerOpen))
}

// ReportThreadPool reports metrics of a thread pool
func ReportThreadPool(cluster string, data hystrix.Data) {
	log.WithField("cluster", cluster).WithField("data", data).Debug("reporting")

	var name = strings.ToLower(data.Name)

	threadPoolCurrentCorePoolSize.WithLabelValues(cluster, name).Set(data.CurrentCorePoolSize)
	threadPoolCurrentLargestPoolSize.WithLabelValues(cluster, name).Set(data.CurrentLargestPoolSize)
	threadPoolCurrentActiveCount.WithLabelValues(cluster, name).Set(data.CurrentActiveCount)
	threadPoolCurrentMaximumPoolSize.WithLabelValues(cluster, name).Set(data.CurrentMaximumPoolSize)
	threadPoolCurrentQueueSize.WithLabelValues(cluster, name).Set(data.CurrentQueueSize)
	threadPoolCurrentTaskCount.WithLabelValues(cluster, name).Set(data.CurrentTaskCount)
	threadPoolCurrentCompletedTaskCount.WithLabelValues(cluster, name).Set(data.CurrentCompletedTaskCount)
	threadPoolRollingMaxActiveThreads.WithLabelValues(cluster, name).Set(data.RollingMaxActiveThreads)
	threadPoolRollingCountCommandRejections.WithLabelValues(cluster, name).Set(data.RollingCountCommandRejections)
	threadPoolReportingHosts.WithLabelValues(cluster, name).Set(data.ReportingHosts)
	threadPoolCurrentPoolSize.WithLabelValues(cluster, name).Set(data.CurrentPoolSize)
	threadPoolRollingCountThreadsExecuted.WithLabelValues(cluster, name).Set(data.RollingCountThreadsExecuted)
}

func boolToFloat64(b bool) float64 {
	if b {
		return 1.0
	}
	return 0.0
}
