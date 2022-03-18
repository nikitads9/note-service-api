package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitSlice(t *testing.T) {
	t.Run("last batch not full", func(t *testing.T) {
		numbers := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
		expected := [][]int64{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9}}
		res, err := SplitSlice(numbers, int64(2))
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("last batch not full", func(t *testing.T) {
		numbers := []int64{1, 2, 3, 4, 5, 6, 7, 8}
		expected := [][]int64{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
		res, err := SplitSlice(numbers, int64(2))
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("batch size is equal to slice size", func(t *testing.T) {
		numbers := []int64{1, 2, 3, 4, 5, 6, 7, 8}
		expected := [][]int64{{1, 2, 3, 4, 5, 6, 7, 8}}
		res, err := SplitSlice(numbers, int64(8))
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("empty slice passed", func(t *testing.T) {
		numbers := []int64{}
		_, err := SplitSlice(numbers, int64(2))
		require.Error(t, err)
	})

	t.Run("negative batchSize passed", func(t *testing.T) {
		numbers := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
		_, err := SplitSlice(numbers, int64(-2))
		require.Error(t, err)
	})
}

func TestSwapKeyValue(t *testing.T) {
	t.Run("swap successful", func(t *testing.T) {
		someMap := map[int64]string{1: "адын", 2: "цвай", 3: "драй"}
		expected := map[string]int64{"адын": 1, "цвай": 2, "драй": 3}
		res, err := SwapKeyValue(someMap)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("empty map passed", func(t *testing.T) {
		someMap := map[int64]string{}
		_, err := SwapKeyValue(someMap)
		require.Error(t, err)
	})
}

func TestSliceFilter(t *testing.T) {
	t.Run("filter values that exist in base slice", func(t *testing.T) {
		someSlice := []int64{1, 2, 3}
		someFilter := []int64{1, 2}
		expected := []int64{3}
		res, err := SliceFilter(someSlice, someFilter)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("filter values that do not exist in base slice", func(t *testing.T) {
		someSlice := []int64{3, 4, 5, 6}
		someFilter := []int64{1, 2}
		expected := []int64{3, 4, 5, 6}
		res, err := SliceFilter(someSlice, someFilter)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("filter values in random order", func(t *testing.T) {
		someSlice := []int64{1, 2, 3, 4, 5, 6}
		someFilter := []int64{2, 5}
		expected := []int64{1, 3, 4, 6}
		res, err := SliceFilter(someSlice, someFilter)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("empty base slice passed", func(t *testing.T) {
		someSlice := []int64{}
		someFilter := []int64{1, 2}
		_, err := SliceFilter(someSlice, someFilter)
		require.Error(t, err)
	})

	t.Run("empty filter slice passed", func(t *testing.T) {
		someSlice := []int64{1, 2}
		someFilter := []int64{}
		_, err := SliceFilter(someSlice, someFilter)
		require.Error(t, err)
	})
}

func TestSliceToMap(t *testing.T) {
	t.Run("cloned slice to map", func(t *testing.T) {
		someSlice := []int64{1, 2, 3, 4}
		expected := map[int64]struct{}{int64(1): {}, int64(2): {}, int64(3): {}, int64(4): {}}
		res, err := SliceToMap(someSlice)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("empty slice passed", func(t *testing.T) {
		someSlice := []int64{}
		_, err := SliceToMap(someSlice)
		require.Error(t, err)
	})
}
