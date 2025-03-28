package main

import (
	"fmt"
)

type MeasuredType string

type GoogleMetric struct {
	MeasuredType MeasuredType
}

func partitionMetrics(googleMetrics []GoogleMetric) [][]GoogleMetric {
	if hasDuplicateMeasuredTypes(googleMetrics) {
		return [][]GoogleMetric{googleMetrics}
	}

	metricsByMeasuredType := make(map[MeasuredType][]GoogleMetric)
	for _, metric := range googleMetrics {
		metricsByMeasuredType[metric.MeasuredType] = append(metricsByMeasuredType[metric.MeasuredType], metric)
	}

	var partitionedMetrics [][]GoogleMetric
	for len(metricsByMeasuredType) > 0 {
		var partition []GoogleMetric
		for measuredType, metrics := range metricsByMeasuredType {
			if len(metrics) > 0 {
				partition = append(partition, metrics[0])
				metricsByMeasuredType[measuredType] = metrics[1:]
				if len(metricsByMeasuredType[measuredType]) == 0 {
					delete(metricsByMeasuredType, measuredType)
				}
			}
		}
		partitionedMetrics = append(partitionedMetrics, partition)
	}
	return partitionedMetrics
}

func hasDuplicateMeasuredTypes(googleMetrics []GoogleMetric) bool {
	measuredTypes := make(map[MeasuredType]struct{})
	for _, metric := range googleMetrics {
		measuredTypes[metric.MeasuredType] = struct{}{}
	}
	return len(measuredTypes) < len(googleMetrics)
}

func main() {
	metrics := []GoogleMetric{
		{MeasuredType: "READ_IO"},
		{MeasuredType: "WRITE_IO"},
		{MeasuredType: "READ_IO"},
	}

	partitions := partitionMetrics(metrics)
	for i, partition := range partitions {
		fmt.Printf("Partition %d: %v\n", i, partition)
	}
}
