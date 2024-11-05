// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package metrics

// Metrics keys
const (
	WorkflowCompletedCounter     = "temporal_workflow_completed"
	WorkflowCanceledCounter      = "temporal_workflow_canceled"
	WorkflowFailedCounter        = "temporal_workflow_failed"
	WorkflowContinueAsNewCounter = "temporal_workflow_continue_as_new"
	WorkflowEndToEndLatency      = "temporal_workflow_endtoend_latency" // measure workflow execution from start to close

	WorkflowTaskReplayLatency           = "temporal_workflow_task_replay_latency"
	WorkflowTaskQueuePollEmptyCounter   = "temporal_workflow_task_queue_poll_empty"
	WorkflowTaskQueuePollSucceedCounter = "temporal_workflow_task_queue_poll_succeed"
	WorkflowTaskScheduleToStartLatency  = "temporal_workflow_task_schedule_to_start_latency"
	WorkflowTaskExecutionLatency        = "temporal_workflow_task_execution_latency"
	WorkflowTaskExecutionFailureCounter = "temporal_workflow_task_execution_failed"
	WorkflowTaskNoCompletionCounter     = "temporal_workflow_task_no_completion"

	ActivityPollNoTaskCounter             = "temporal_activity_poll_no_task"
	ActivityScheduleToStartLatency        = "temporal_activity_schedule_to_start_latency"
	ActivityExecutionFailedCounter        = "temporal_activity_execution_failed"
	UnregisteredActivityInvocationCounter = "temporal_unregistered_activity_invocation"
	ActivityExecutionLatency              = "temporal_activity_execution_latency"
	ActivitySucceedEndToEndLatency        = "temporal_activity_succeed_endtoend_latency"
	ActivityTaskErrorCounter              = "temporal_activity_task_error"

	LocalActivityTotalCounter             = "temporal_local_activity_total"
	LocalActivityCanceledCounter          = "temporal_local_activity_canceled" // Deprecated: Use LocalActivityExecutionCanceledCounter instead.
	LocalActivityExecutionCanceledCounter = "temporal_local_activity_execution_cancelled"
	LocalActivityFailedCounter            = "temporal_local_activity_failed" // Deprecated: Use LocalActivityExecutionFailedCounter instead.
	LocalActivityExecutionFailedCounter   = "temporal_local_activity_execution_failed"
	LocalActivityErrorCounter             = "temporal_local_activity_error"
	LocalActivityExecutionLatency         = "temporal_local_activity_execution_latency"
	LocalActivitySucceedEndToEndLatency   = "temporal_local_activity_succeed_endtoend_latency"

	CorruptedSignalsCounter = "temporal_corrupted_signals"

	WorkerStartCounter       = "temporal_worker_start"
	WorkerTaskSlotsAvailable = "temporal_worker_task_slots_available"
	WorkerTaskSlotsUsed      = "temporal_worker_task_slots_used"
	PollerStartCounter       = "temporal_poller_start"
	NumPoller                = "temporal_num_pollers"

	TemporalRequest                      = "temporal_request"
	TemporalRequestFailure               = "temporal_request_failure"
	TemporalRequestLatency               = "temporal_request_latency"
	TemporalLongRequest                  = "temporal_long_request"
	TemporalLongRequestFailure           = "temporal_long_request_failure"
	TemporalLongRequestLatency           = "temporal_long_request_latency"
	TemporalRequestResourceExhausted     = "temporal_request_resource_exhausted"
	TemporalLongRequestResourceExhausted = "temporal_long_request_resource_exhausted"

	StickyCacheHit                 = "temporal_sticky_cache_hit"
	StickyCacheMiss                = "temporal_sticky_cache_miss"
	StickyCacheTotalForcedEviction = "temporal_sticky_cache_total_forced_eviction"
	StickyCacheSize                = "temporal_sticky_cache_size"

	WorkflowActiveThreadCount = "temporal_workflow_active_thread_count"

	NexusPollNoTaskCounter          = "temporal_nexus_poll_no_task"
	NexusTaskScheduleToStartLatency = "temporal_nexus_task_schedule_to_start_latency"
	NexusTaskExecutionFailedCounter = "temporal_nexus_task_execution_failed"
	NexusTaskExecutionLatency       = "temporal_nexus_task_execution_latency"
	NexusTaskEndToEndLatency        = "temporal_nexus_task_endtoend_latency"
)

// Metric tag keys
const (
	NamespaceTagName        = "namespace"
	ClientTagName           = "client_name"
	PollerTypeTagName       = "poller_type"
	WorkerTypeTagName       = "worker_type"
	WorkflowTypeNameTagName = "workflow_type"
	ActivityTypeNameTagName = "activity_type"
	NexusServiceTagName     = "nexus_service"
	NexusOperationTagName   = "nexus_operation"
	FailureReasonTagName    = "failure_reason"
	TaskQueueTagName        = "task_queue"
	OperationTagName        = "operation"
	CauseTagName            = "cause"
	RequestFailureCode      = "status_code"
)

// Metric tag values
const (
	NoneTagValue                 = "none"
	ClientTagValue               = "temporal_go"
	PollerTypeWorkflowTask       = "workflow_task"
	PollerTypeWorkflowStickyTask = "workflow_sticky_task"
	PollerTypeActivityTask       = "activity_task"
	PollerTypeNexusTask          = "nexus_task"
)
