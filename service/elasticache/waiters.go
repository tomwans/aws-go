// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elasticache

import (
	"github.com/aws/aws-sdk-go/private/waiter"
)

var waiterCacheClusterAvailable *waiter.Config

func (c *ElastiCache) WaitUntilCacheClusterAvailable(input *DescribeCacheClustersInput) error {
	if waiterCacheClusterAvailable == nil {
		waiterCacheClusterAvailable = &waiter.Config{
			Operation:   "DescribeCacheClusters",
			Delay:       15,
			MaxAttempts: 40,
			Acceptors: []waiter.WaitAcceptor{
				{
					State:    "success",
					Matcher:  "pathAll",
					Argument: "CacheClusters[].CacheClusterStatus",
					Expected: "available",
				},
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "CacheClusters[].CacheClusterStatus",
					Expected: "deleted",
				},
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "CacheClusters[].CacheClusterStatus",
					Expected: "deleting",
				},
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "CacheClusters[].CacheClusterStatus",
					Expected: "incompatible-network",
				},
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "CacheClusters[].CacheClusterStatus",
					Expected: "restore-failed",
				},
			},
		}
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCacheClusterAvailable,
	}
	return w.Wait()
}

var waiterCacheClusterDeleted *waiter.Config

func (c *ElastiCache) WaitUntilCacheClusterDeleted(input *DescribeCacheClustersInput) error {
	if waiterCacheClusterDeleted == nil {
		waiterCacheClusterDeleted = &waiter.Config{
			Operation:   "DescribeCacheClusters",
			Delay:       15,
			MaxAttempts: 40,
			Acceptors: []waiter.WaitAcceptor{
				{
					State:    "success",
					Matcher:  "pathAll",
					Argument: "CacheClusters[].CacheClusterStatus",
					Expected: "deleted",
				},
				{
					State:    "success",
					Matcher:  "error",
					Argument: "",
					Expected: "CacheClusterNotFound",
				},
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "CacheClusters[].CacheClusterStatus",
					Expected: "available",
				},
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "CacheClusters[].CacheClusterStatus",
					Expected: "creating",
				},
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "CacheClusters[].CacheClusterStatus",
					Expected: "incompatible-network",
				},
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "CacheClusters[].CacheClusterStatus",
					Expected: "modifying",
				},
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "CacheClusters[].CacheClusterStatus",
					Expected: "restore-failed",
				},
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "CacheClusters[].CacheClusterStatus",
					Expected: "snapshotting",
				},
			},
		}
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCacheClusterDeleted,
	}
	return w.Wait()
}

var waiterReplicationGroupAvailable *waiter.Config

func (c *ElastiCache) WaitUntilReplicationGroupAvailable(input *DescribeReplicationGroupsInput) error {
	if waiterReplicationGroupAvailable == nil {
		waiterReplicationGroupAvailable = &waiter.Config{
			Operation:   "DescribeReplicationGroups",
			Delay:       15,
			MaxAttempts: 40,
			Acceptors: []waiter.WaitAcceptor{
				{
					State:    "success",
					Matcher:  "pathAll",
					Argument: "ReplicationGroups[].Status",
					Expected: "available",
				},
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "ReplicationGroups[].Status",
					Expected: "deleted",
				},
			},
		}
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterReplicationGroupAvailable,
	}
	return w.Wait()
}

var waiterReplicationGroupDeleted *waiter.Config

func (c *ElastiCache) WaitUntilReplicationGroupDeleted(input *DescribeReplicationGroupsInput) error {
	if waiterReplicationGroupDeleted == nil {
		waiterReplicationGroupDeleted = &waiter.Config{
			Operation:   "DescribeReplicationGroups",
			Delay:       15,
			MaxAttempts: 40,
			Acceptors: []waiter.WaitAcceptor{
				{
					State:    "success",
					Matcher:  "pathAll",
					Argument: "ReplicationGroups[].Status",
					Expected: "deleted",
				},
				{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "ReplicationGroups[].Status",
					Expected: "available",
				},
				{
					State:    "success",
					Matcher:  "error",
					Argument: "",
					Expected: "ReplicationGroupNotFoundFault",
				},
			},
		}
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterReplicationGroupDeleted,
	}
	return w.Wait()
}