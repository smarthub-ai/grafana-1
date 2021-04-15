package state

import (
	"fmt"
	"testing"
	"time"

	"github.com/grafana/grafana/pkg/infra/log"

	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/grafana/pkg/services/ngalert/eval"
	"github.com/grafana/grafana/pkg/services/ngalert/models"
	"github.com/stretchr/testify/assert"
)

func TestProcessEvalResults(t *testing.T) {
	evaluationTime, err := time.Parse("2006-01-02", "2021-03-25")
	if err != nil {
		t.Fatalf("error parsing date format: %s", err.Error())
	}
	cacheId := "test_uid map[label1:value1 label2:value2]"
	testCases := []struct {
		desc                       string
		uid                        string
		evalResults                eval.Results
		condition                  models.Condition
		expectedState              eval.State
		expectedReturnedStateCount int
		expectedResultCount        int
		expectedCacheEntries       []AlertState
	}{
		{
			desc: "given a single evaluation result",
			uid:  "test_uid",
			evalResults: eval.Results{
				eval.Result{
					Instance:    data.Labels{"label1": "value1", "label2": "value2"},
					State:       eval.Normal,
					EvaluatedAt: evaluationTime,
				},
			},
			condition: models.Condition{
				Condition: "A",
				OrgID:     123,
			},
			expectedState:              eval.Normal,
			expectedReturnedStateCount: 0,
			expectedResultCount:        1,
			expectedCacheEntries: []AlertState{
				{
					UID:     "test_uid",
					OrgID:   123,
					CacheId: cacheId,
					Labels:  data.Labels{"label1": "value1", "label2": "value2"},
					State:   eval.Normal,
					Results: []StateEvaluation{
						{EvaluationTime: evaluationTime, EvaluationState: eval.Normal},
					},
					StartsAt:           time.Time{},
					EndsAt:             time.Time{},
					LastEvaluationTime: evaluationTime,
				},
			},
		},
		{
			desc: "given a state change from normal to alerting for a single entity",
			uid:  "test_uid",
			evalResults: eval.Results{
				eval.Result{
					Instance:    data.Labels{"label1": "value1", "label2": "value2"},
					State:       eval.Normal,
					EvaluatedAt: evaluationTime,
				},
				eval.Result{
					Instance:    data.Labels{"label1": "value1", "label2": "value2"},
					State:       eval.Alerting,
					EvaluatedAt: evaluationTime.Add(1 * time.Minute),
				},
			},
			condition: models.Condition{
				Condition: "A",
				OrgID:     123,
			},
			expectedState:              eval.Alerting,
			expectedReturnedStateCount: 1,
			expectedResultCount:        2,
			expectedCacheEntries: []AlertState{
				{
					UID:     "test_uid",
					OrgID:   123,
					CacheId: cacheId,
					Labels:  data.Labels{"label1": "value1", "label2": "value2"},
					State:   eval.Alerting,
					Results: []StateEvaluation{
						{EvaluationTime: evaluationTime, EvaluationState: eval.Normal},
						{EvaluationTime: evaluationTime.Add(1 * time.Minute), EvaluationState: eval.Alerting},
					},
					StartsAt:           evaluationTime.Add(1 * time.Minute),
					EndsAt:             evaluationTime.Add(100 * time.Second),
					LastEvaluationTime: evaluationTime.Add(1 * time.Minute),
				},
			},
		},
		{
			desc: "given a state change from alerting to normal for a single entity",
			uid:  "test_uid",
			evalResults: eval.Results{
				eval.Result{
					Instance:    data.Labels{"label1": "value1", "label2": "value2"},
					State:       eval.Alerting,
					EvaluatedAt: evaluationTime,
				},
				eval.Result{
					Instance:    data.Labels{"label1": "value1", "label2": "value2"},
					State:       eval.Normal,
					EvaluatedAt: evaluationTime.Add(1 * time.Minute),
				},
			},
			condition: models.Condition{
				Condition: "A",
				OrgID:     123,
			},
			expectedState:              eval.Normal,
			expectedReturnedStateCount: 1,
			expectedResultCount:        2,
			expectedCacheEntries: []AlertState{
				{
					UID:     "test_uid",
					OrgID:   123,
					CacheId: cacheId,
					Labels:  data.Labels{"label1": "value1", "label2": "value2"},
					State:   eval.Normal,
					Results: []StateEvaluation{
						{EvaluationTime: evaluationTime, EvaluationState: eval.Alerting},
						{EvaluationTime: evaluationTime.Add(1 * time.Minute), EvaluationState: eval.Normal},
					},
					StartsAt:           evaluationTime,
					EndsAt:             evaluationTime.Add(1 * time.Minute),
					LastEvaluationTime: evaluationTime.Add(1 * time.Minute),
				},
			},
		},
		{
			desc: "given a constant alerting state for a single entity",
			uid:  "test_uid",
			evalResults: eval.Results{
				eval.Result{
					Instance:    data.Labels{"label1": "value1", "label2": "value2"},
					State:       eval.Alerting,
					EvaluatedAt: evaluationTime,
				},
				eval.Result{
					Instance:    data.Labels{"label1": "value1", "label2": "value2"},
					State:       eval.Alerting,
					EvaluatedAt: evaluationTime.Add(1 * time.Minute),
				},
			},
			condition: models.Condition{
				Condition: "A",
				OrgID:     123,
			},
			expectedState:              eval.Alerting,
			expectedReturnedStateCount: 0,
			expectedResultCount:        2,
			expectedCacheEntries: []AlertState{
				{
					UID:     "test_uid",
					OrgID:   123,
					CacheId: cacheId,
					Labels:  data.Labels{"label1": "value1", "label2": "value2"},
					State:   eval.Alerting,
					Results: []StateEvaluation{
						{EvaluationTime: evaluationTime, EvaluationState: eval.Alerting},
						{EvaluationTime: evaluationTime.Add(1 * time.Minute), EvaluationState: eval.Alerting},
					},
					StartsAt:           evaluationTime,
					EndsAt:             evaluationTime.Add(100 * time.Second),
					LastEvaluationTime: evaluationTime.Add(1 * time.Minute),
				},
			},
		},
		{
			desc: "given a constant normal state for a single entity",
			uid:  "test_uid",
			evalResults: eval.Results{
				eval.Result{
					Instance:    data.Labels{"label1": "value1", "label2": "value2"},
					State:       eval.Normal,
					EvaluatedAt: evaluationTime,
				},
				eval.Result{
					Instance:    data.Labels{"label1": "value1", "label2": "value2"},
					State:       eval.Normal,
					EvaluatedAt: evaluationTime.Add(1 * time.Minute),
				},
			},
			condition: models.Condition{
				Condition: "A",
				OrgID:     123,
			},
			expectedState:              eval.Normal,
			expectedReturnedStateCount: 0,
			expectedResultCount:        2,
			expectedCacheEntries: []AlertState{
				{
					UID:     "test_uid",
					OrgID:   123,
					CacheId: cacheId,
					Labels:  data.Labels{"label1": "value1", "label2": "value2"},
					State:   eval.Normal,
					Results: []StateEvaluation{
						{evaluationTime, eval.Normal},
						{EvaluationTime: evaluationTime.Add(1 * time.Minute), EvaluationState: eval.Normal},
					},
					StartsAt:           time.Time{},
					EndsAt:             time.Time{},
					LastEvaluationTime: evaluationTime.Add(1 * time.Minute),
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run("all fields for a cache entry are set correctly", func(t *testing.T) {
			st := NewStateTracker(log.New("test_state_tracker"))
			_ = st.ProcessEvalResults(tc.uid, tc.evalResults, tc.condition)
			for _, entry := range tc.expectedCacheEntries {
				if !entry.Equals(st.Get(entry.CacheId)) {
					t.Log(tc.desc)
					printEntryDiff(entry, st.Get(entry.CacheId), t)
				}
				assert.True(t, entry.Equals(st.Get(entry.CacheId)))
			}
		})

		t.Run("the expected number of entries are added to the cache", func(t *testing.T) {
			st := NewStateTracker(log.New("test_state_tracker"))
			st.ProcessEvalResults(tc.uid, tc.evalResults, tc.condition)
			assert.Equal(t, len(tc.expectedCacheEntries), len(st.stateCache.cacheMap))
		})

		//This test, as configured, does not quite represent the behavior of the system.
		//It is expected that each batch of evaluation results will have only one result
		//for a unique set of labels.
		t.Run("the expected number of states are returned to the caller", func(t *testing.T) {
			st := NewStateTracker(log.New("test_state_tracker"))
			results := st.ProcessEvalResults(tc.uid, tc.evalResults, tc.condition)
			assert.Equal(t, len(tc.evalResults), len(results))
		})
	}
}

func printEntryDiff(a, b AlertState, t *testing.T) {
	if a.UID != b.UID {
		t.Log(fmt.Sprintf("%v \t %v\n", a.UID, b.UID))
	}
	if a.OrgID != b.OrgID {
		t.Log(fmt.Sprintf("%v \t %v\n", a.OrgID, b.OrgID))
	}
	if a.CacheId != b.CacheId {
		t.Log(fmt.Sprintf("%v \t %v\n", a.CacheId, b.CacheId))
	}
	if !a.Labels.Equals(b.Labels) {
		t.Log(fmt.Sprintf("%v \t %v\n", a.Labels, b.Labels))
	}
	if a.StartsAt != b.StartsAt {
		t.Log(fmt.Sprintf("%v \t %v\n", a.StartsAt, b.StartsAt))
	}
	if a.EndsAt != b.EndsAt {
		t.Log(fmt.Sprintf("%v \t %v\n", a.EndsAt, b.EndsAt))
	}
	if a.LastEvaluationTime != b.LastEvaluationTime {
		t.Log(fmt.Sprintf("%v \t %v\n", a.LastEvaluationTime, b.LastEvaluationTime))
	}
	if len(a.Results) != len(b.Results) {
		t.Log(fmt.Sprintf("a: %d b: %d", len(a.Results), len(b.Results)))
		t.Log("a")
		for i := 0; i < len(a.Results); i++ {
			t.Log(fmt.Sprintf("%v\n", a.Results[i]))
		}
		t.Log("b")
		for i := 0; i < len(b.Results); i++ {
			t.Log(fmt.Sprintf("%v\n", b.Results[i]))
		}
	}
}
