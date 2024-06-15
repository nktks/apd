package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AppendCSV(t *testing.T) {
	kf := NewKeyFinder("", "")
	t.Run("csv", func(t *testing.T) {
		t.Run("with header", func(t *testing.T) {
			t.Run("can find keys", func(t *testing.T) {
				input := "created_at,updated_at\n2020-01-01 00:00:00,2020-01-01 00:00:01"
				output, err := AppendCSV(kf, input, NoIndex, NoIndex)
				require.NoError(t, err)
				require.Equal(t, "created_at,updated_at,duration\n2020-01-01 00:00:00,2020-01-01 00:00:01,1s\n", output)
			})
			t.Run("cant find from key", func(t *testing.T) {
				input := "hoge,updated_at\n2020-01-01 00:00:00,2020-01-01 00:00:01"
				output, err := AppendCSV(kf, input, NoIndex, NoIndex)
				require.NoError(t, err)
				require.Equal(t, "hoge,updated_at\n2020-01-01 00:00:00,2020-01-01 00:00:01\n", output)
			})
			t.Run("cant find to key", func(t *testing.T) {
				input := "created_at,hoge\n2020-01-01 00:00:00,2020-01-01 00:00:01"
				output, err := AppendCSV(kf, input, NoIndex, NoIndex)
				require.NoError(t, err)
				require.Equal(t, "created_at,hoge\n2020-01-01 00:00:00,2020-01-01 00:00:01\n", output)
			})
			t.Run("cant parse time", func(t *testing.T) {
				input := "created_at,updated_at\nhoge,2020-01-01 00:00:01\n2020-01-01 00:00:00,2020-01-01 00:00:01"
				output, err := AppendCSV(kf, input, NoIndex, NoIndex)
				require.NoError(t, err)
				require.Equal(t, "created_at,updated_at,duration\nhoge,2020-01-01 00:00:01,\n2020-01-01 00:00:00,2020-01-01 00:00:01,1s\n", output)
			})
		})
		t.Run("without header", func(t *testing.T) {
			t.Run("can find index", func(t *testing.T) {
				input := "2020-01-01 00:00:00,2020-01-01 00:00:01"
				output, err := AppendCSV(kf, input, 0, 1)
				require.NoError(t, err)
				require.Equal(t, "2020-01-01 00:00:00,2020-01-01 00:00:01,1s\n", output)
			})
			t.Run("cant find from index", func(t *testing.T) {
				input := "2020-01-01 00:00:00,2020-01-01 00:00:01"
				output, err := AppendCSV(kf, input, 2, 1)
				require.NoError(t, err)
				require.Equal(t, "2020-01-01 00:00:00,2020-01-01 00:00:01\n", output)
			})
			t.Run("cant find to index", func(t *testing.T) {
				input := "2020-01-01 00:00:00,2020-01-01 00:00:01"
				output, err := AppendCSV(kf, input, 0, 2)
				require.NoError(t, err)
				require.Equal(t, "2020-01-01 00:00:00,2020-01-01 00:00:01\n", output)
			})
		})
	})
	t.Run("tsv", func(t *testing.T) {
		t.Run("with header", func(t *testing.T) {
			t.Run("can find keys", func(t *testing.T) {
				input := "created_at\tupdated_at\n2020-01-01 00:00:00\t2020-01-01 00:00:01"
				output, err := AppendCSV(kf, input, NoIndex, NoIndex)
				require.NoError(t, err)
				require.Equal(t, "created_at\tupdated_at\tduration\n2020-01-01 00:00:00\t2020-01-01 00:00:01\t1s\n", output)
			})
			t.Run("cant find from key", func(t *testing.T) {
				input := "hoge\tupdated_at\n2020-01-01 00:00:00\t2020-01-01 00:00:01"
				output, err := AppendCSV(kf, input, NoIndex, NoIndex)
				require.NoError(t, err)
				require.Equal(t, "hoge\tupdated_at\n2020-01-01 00:00:00\t2020-01-01 00:00:01\n", output)
			})
			t.Run("cant find to key", func(t *testing.T) {
				input := "created_at\thoge\n2020-01-01 00:00:00\t2020-01-01 00:00:01"
				output, err := AppendCSV(kf, input, NoIndex, NoIndex)
				require.NoError(t, err)
				require.Equal(t, "created_at\thoge\n2020-01-01 00:00:00\t2020-01-01 00:00:01\n", output)
			})
			t.Run("cant parse time", func(t *testing.T) {
				input := "created_at\tupdated_at\nhoge\t2020-01-01 00:00:01\n2020-01-01 00:00:00\t2020-01-01 00:00:01"
				output, err := AppendCSV(kf, input, NoIndex, NoIndex)
				require.NoError(t, err)
				require.Equal(t, "created_at\tupdated_at\tduration\nhoge\t2020-01-01 00:00:01\t\n2020-01-01 00:00:00\t2020-01-01 00:00:01\t1s\n", output)
			})
		})
		t.Run("without header", func(t *testing.T) {
			t.Run("can find index", func(t *testing.T) {
				input := "2020-01-01 00:00:00\t2020-01-01 00:00:01"
				output, err := AppendCSV(kf, input, 0, 1)
				require.NoError(t, err)
				require.Equal(t, "2020-01-01 00:00:00\t2020-01-01 00:00:01\t1s\n", output)
			})
			t.Run("cant find from index", func(t *testing.T) {
				input := "2020-01-01 00:00:00\t2020-01-01 00:00:01"
				output, err := AppendCSV(kf, input, 2, 1)
				require.NoError(t, err)
				require.Equal(t, "2020-01-01 00:00:00\t2020-01-01 00:00:01\n", output)
			})
			t.Run("cant find to index", func(t *testing.T) {
				input := "2020-01-01 00:00:00\t2020-01-01 00:00:01"
				output, err := AppendCSV(kf, input, 0, 2)
				require.NoError(t, err)
				require.Equal(t, "2020-01-01 00:00:00\t2020-01-01 00:00:01\n", output)
			})
		})
	})
}
