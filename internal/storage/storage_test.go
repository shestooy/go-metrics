package storage

import (
	"github.com/shestooy/go-musthave-metrics-tpl.git/internal/flags"
	"github.com/shestooy/go-musthave-metrics-tpl.git/internal/server/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemStorage_Init(t *testing.T) {
	flags.Restore = false
	flags.StorageInterval = 5000
	err := MStorage.Init()
	require.NoError(t, err)
}

func TestStorage_UpdateMetric(t *testing.T) {
	flags.Restore = false
	flags.StorageInterval = 5000
	err := MStorage.Init()
	require.NoError(t, err)

	metric := model.Metrics{
		ID:    "test_gauge",
		MType: "gauge",
		Value: func(v float64) *float64 { return &v }(10),
	}

	err = MStorage.UpdateMetric(metric)
	assert.NoError(t, err)

	storedMetric, err := MStorage.GetMetricID(metric.ID)
	assert.NoError(t, err)
	assert.Equal(t, metric, storedMetric)
}

func TestMemStorage_GetMetricID(t *testing.T) {
	flags.Restore = false
	flags.StorageInterval = 5000
	err := MStorage.Init()
	require.NoError(t, err)

	metric := model.Metrics{
		ID:    "test_gauge",
		MType: "gauge",
		Value: func(v float64) *float64 { return &v }(20),
	}

	err = MStorage.UpdateMetric(metric)
	require.NoError(t, err)

	retrievedMetric, err := MStorage.GetMetricID(metric.ID)
	assert.NoError(t, err)
	assert.Equal(t, metric, retrievedMetric)

	_, err = MStorage.GetMetricID("testErr")
	assert.Error(t, err)
}

func TestStorage_GetAllMetrics(t *testing.T) {
	flags.Restore = false
	flags.StorageInterval = 5000
	err := MStorage.Init()
	require.NoError(t, err)

	metric1 := model.Metrics{
		ID:    "test_gauge",
		MType: "gauge",
		Value: func(v float64) *float64 { return &v }(450.1),
	}
	metric2 := model.Metrics{
		ID:    "test_counter",
		MType: "counter",
		Delta: func(v int64) *int64 { return &v }(432),
	}

	err = MStorage.UpdateMetric(metric1)
	require.NoError(t, err)
	err = MStorage.UpdateMetric(metric2)
	require.NoError(t, err)

	allMetrics := MStorage.GetAllMetrics()
	assert.Equal(t, 2, len(allMetrics))
	assert.Contains(t, allMetrics, metric1.ID)
	assert.Contains(t, allMetrics, metric2.ID)
}
