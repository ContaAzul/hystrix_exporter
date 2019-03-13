package hystrix

import (
	"encoding/json"
	"strings"
)

// Unmarshal the given string into a struct
func Unmarshal(line string) (data Data, err error) {
	line = strings.TrimPrefix(line, "data:")
	err = json.Unmarshal([]byte(line), &data)
	return
}

// Latencies of a circuit
type Latencies struct {
	L0   float64 `json:"0,omitempty"`
	L25  float64 `json:"25,omitempty"`
	L50  float64 `json:"50,omitempty"`
	L75  float64 `json:"75,omitempty"`
	L90  float64 `json:"90,omitempty"`
	L95  float64 `json:"95,omitempty"`
	L99  float64 `json:"99,omitempty"`
	L995 float64 `json:"99.5,omitempty"`
	L100 float64 `json:"100,omitempty"`
}

// Data Hystrix main data type
type Data struct {

	// Common properties
	Type           string  `json:"type,omitempty"`
	Name           string  `json:"name,omitempty"`
	ReportingHosts float64 `json:"reportingHosts,omitempty"`

	// Command properties
	Group                string  `json:"group,omitempty"`
	ThreadPool           string  `json:"threadPool,omitempty"`
	IsCircuitBreakerOpen bool    `json:"isCircuitBreakerOpen,omitempty"`
	ErrorPercentage      float64 `json:"errorPercentage,omitempty"`
	ErrorCount           float64 `json:"errorCount,omitempty"`
	RequestCount         float64 `json:"requestCount,omitempty"`

	RollingCountBadRequests        float64 `json:"rollingCountBadRequests,omitempty"`
	RollingCountCollapsedRequests  float64 `json:"rollingCountCollapsedRequests,omitempty"`
	RollingCountEmit               float64 `json:"rollingCountEmit,omitempty"`
	RollingCountExceptionsThrown   float64 `json:"rollingCountExceptionsThrown,omitempty"`
	RollingCountFailure            float64 `json:"rollingCountFailure,omitempty"`
	RollingCountFallbackEmit       float64 `json:"rollingCountFallbackEmit,omitempty"`
	RollingCountFallbackFailure    float64 `json:"rollingCountFallbackFailure,omitempty"`
	RollingCountFallbackMissing    float64 `json:"rollingCountFallbackMissing,omitempty"`
	RollingCountFallbackRejection  float64 `json:"rollingCountFallbackRejection,omitempty"`
	RollingCountFallbackSuccess    float64 `json:"rollingCountFallbackSuccess,omitempty"`
	RollingCountResponsesFromCache float64 `json:"rollingCountResponsesFromCache,omitempty"`
	RollingCountSemaphoreRejected  float64 `json:"rollingCountSemaphoreRejected,omitempty"`
	RollingCountShortCircuited     float64 `json:"rollingCountShortCircuited,omitempty"`
	RollingCountSuccess            float64 `json:"rollingCountSuccess,omitempty"`
	RollingCountThreadPoolRejected float64 `json:"rollingCountThreadPoolRejected,omitempty"`
	RollingCountTimeout            float64 `json:"rollingCountTimeout,omitempty"`

	CurrentConcurrentExecutionCount    float64 `json:"currentConcurrentExecutionCount,omitempty"`
	RollingMaxConcurrentExecutionCount float64 `json:"rollingMaxConcurrentExecutionCount,omitempty"`

	LatencyExecuteMean float64   `json:"latencyExecute_mean,omitempty"`
	LatencyTotalMean   float64   `json:"latencyTotal_mean,omitempty"`
	LatencyTotal       Latencies `json:"latencyTotal,omitempty"`
	LatencyExecute     Latencies `json:"latencyExecute,omitempty"`

	// Threadpool properties
	CurrentActiveCount            float64 `json:"currentActiveCount,omitempty"`
	CurrentCompletedTaskCount     float64 `json:"currentCompletedTaskCount,omitempty"`
	CurrentCorePoolSize           float64 `json:"currentCorePoolSize,omitempty"`
	CurrentLargestPoolSize        float64 `json:"currentLargestPoolSize,omitempty"`
	CurrentMaximumPoolSize        float64 `json:"currentMaximumPoolSize,omitempty"`
	CurrentPoolSize               float64 `json:"currentPoolSize,omitempty"`
	CurrentQueueSize              float64 `json:"currentQueueSize,omitempty"`
	CurrentTaskCount              float64 `json:"currentTaskCount,omitempty"`
	RollingCountThreadsExecuted   float64 `json:"rollingCountThreadsExecuted,omitempty"`
	RollingMaxActiveThreads       float64 `json:"rollingMaxActiveThreads,omitempty"`
	RollingCountCommandRejections float64 `json:"rollingCountCommandRejections,omitempty"`
}
