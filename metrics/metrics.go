package metrics

import (
	"strings"

	"github.com/ContaAzul/hystrix_exporter/hystrix"
	"github.com/apex/log"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	//
	// Command gauges:
	//

	// CircuitOpen prometheus gauge
	commandCircuitOpen = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_command_circuit_open",
		Help: "circuit open, 1 means true",
	}, []string{"cluster", "name", "group"})

	// ErrorPercentage prometheus gauge
	commandErrorPercentage = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_command_error_percentage",
		Help: "error percentage",
	}, []string{"cluster", "name", "group"})

	// ErrorCount prometheus gauge
	commandErrors = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_command_errors",
		Help: "number of errors",
	}, []string{"cluster", "name", "group"})

	// RequestCount prometheus gauge
	commandRequests = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_command_requests",
		Help: "number of requests",
	}, []string{"cluster", "name", "group"})

	// RollingEventCounts prometheus gauge
	commandRollingEvents = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_command_rolling_events",
		Help: "rolling event counts",
	}, []string{"cluster", "name", "group", "event"})

	// commandConcurrentExecutions prometheus gauge
	commandConcurrentExecutions = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_command_concurrent_executions",
		Help: "number of concurrent executions",
	}, []string{"cluster", "name", "group"})

	// MaxConcurrentExecutionCount prometheus gauge
	commandMaxConcurrentExecutions = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_command_max_concurrent_executions",
		Help: "max number of concurrent executions",
	}, []string{"cluster", "name", "group"})

	// ReportingHosts prometheus gauge
	commandReportingHosts = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "hystrix_command_reporting_hosts",
		Help: "number of reporting hosts",
	}, []string{"cluster", "name", "group"})

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

	//
	// Threadpool gauges:
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
		commandCircuitOpen,
		commandErrorPercentage,
		commandErrors,
		commandRequests,
		commandRollingEvents,
		commandConcurrentExecutions,
		commandMaxConcurrentExecutions,
		commandReportingHosts,

		latencyTotal,
		latencyExecute,

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

	commandCircuitOpen.WithLabelValues(cluster, name, group).Set(boolToFloat64(data.IsCircuitBreakerOpen))
	commandErrorPercentage.WithLabelValues(cluster, name, group).Set(data.ErrorPercentage)
	commandErrors.WithLabelValues(cluster, name, group).Set(data.ErrorCount)
	commandRequests.WithLabelValues(cluster, name, group).Set(data.RequestCount)
	commandConcurrentExecutions.WithLabelValues(cluster, name, group).Set(data.CurrentConcurrentExecutionCount)
	commandMaxConcurrentExecutions.WithLabelValues(cluster, name, group).Set(data.RollingMaxConcurrentExecutionCount)
	commandReportingHosts.WithLabelValues(cluster, name, group).Set(data.ReportingHosts)

	commandRollingEvents.WithLabelValues(cluster, name, group, "bad_request").Set(data.RollingCountBadRequests)
	commandRollingEvents.WithLabelValues(cluster, name, group, "collapsed").Set(data.RollingCountCollapsedRequests)
	commandRollingEvents.WithLabelValues(cluster, name, group, "emit").Set(data.RollingCountEmit)
	commandRollingEvents.WithLabelValues(cluster, name, group, "exception_thrown").Set(data.RollingCountExceptionsThrown)
	commandRollingEvents.WithLabelValues(cluster, name, group, "failure").Set(data.RollingCountFailure)
	commandRollingEvents.WithLabelValues(cluster, name, group, "fallback_emit").Set(data.RollingCountFallbackEmit)
	commandRollingEvents.WithLabelValues(cluster, name, group, "fallback_failure").Set(data.RollingCountFallbackFailure)
	commandRollingEvents.WithLabelValues(cluster, name, group, "fallback_missing").Set(data.RollingCountFallbackMissing)
	commandRollingEvents.WithLabelValues(cluster, name, group, "fallback_rejection").Set(data.RollingCountFallbackRejection)
	commandRollingEvents.WithLabelValues(cluster, name, group, "fallback_success").Set(data.RollingCountFallbackSuccess)
	commandRollingEvents.WithLabelValues(cluster, name, group, "response_from_cache").Set(data.RollingCountResponsesFromCache)
	commandRollingEvents.WithLabelValues(cluster, name, group, "semaphore_rejected").Set(data.RollingCountSemaphoreRejected)
	commandRollingEvents.WithLabelValues(cluster, name, group, "short_circuited").Set(data.RollingCountShortCircuited)
	commandRollingEvents.WithLabelValues(cluster, name, group, "success").Set(data.RollingCountSuccess)
	commandRollingEvents.WithLabelValues(cluster, name, group, "thread_pool_rejected").Set(data.RollingCountThreadPoolRejected)
	commandRollingEvents.WithLabelValues(cluster, name, group, "timeout").Set(data.RollingCountTimeout)

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
