package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeyFinder_DetectFromKey(t *testing.T) {
	t.Run("force from empty", func(t *testing.T) {
		kf := NewKeyFinder("", "")
		t.Run("match default from keys", func(t *testing.T) {
			require.Equal(t, "created_at", kf.DetectFromKey([]string{"created_at"}))
		})
		t.Run("doesn't match default from keys", func(t *testing.T) {
			require.Equal(t, "", kf.DetectFromKey([]string{"hoge"}))
		})
	})
	t.Run("force from not empty", func(t *testing.T) {
		kf := NewKeyFinder("test_from", "")
		t.Run("match force from key", func(t *testing.T) {
			t.Run("match default from keys", func(t *testing.T) {
				require.Equal(t, "test_from", kf.DetectFromKey([]string{"created_at", "test_from"}))
			})
			t.Run("doesn't match default from keys", func(t *testing.T) {
				require.Equal(t, "", kf.DetectFromKey([]string{"created_at"}))
			})
		})
	})
}
func TestKeyFinder_DetectToKey(t *testing.T) {
	t.Run("force to empty", func(t *testing.T) {
		kf := NewKeyFinder("", "")
		t.Run("match default to keys", func(t *testing.T) {
			require.Equal(t, "updated_at", kf.DetectToKey([]string{"updated_at"}))
		})
		t.Run("doesn't match default to keys", func(t *testing.T) {
			require.Equal(t, "", kf.DetectToKey([]string{"hoge"}))
		})
	})
	t.Run("force to not empty", func(t *testing.T) {
		kf := NewKeyFinder("", "test_to")
		t.Run("match force to key", func(t *testing.T) {
			t.Run("match default to keys", func(t *testing.T) {
				require.Equal(t, "test_to", kf.DetectToKey([]string{"updated_at", "test_to"}))
			})
			t.Run("doesn't match default to keys", func(t *testing.T) {
				require.Equal(t, "", kf.DetectToKey([]string{"updated_at"}))
			})
		})
	})
}

func TestIndexFinder_DetectFromIndex(t *testing.T) {
	t.Run("force from empty", func(t *testing.T) {
		kf := NewKeyFinder("", "")
		t.Run("match default from keys", func(t *testing.T) {
			i, found := kf.DetectFromIndex([]string{"created_at"})
			require.Equal(t, 0, i)
			require.True(t, found)
		})
		t.Run("doesn't match default from keys", func(t *testing.T) {
			i, found := kf.DetectFromIndex([]string{"hoge"})
			require.Equal(t, 0, i)
			require.False(t, found)
		})
	})
	t.Run("force from not empty", func(t *testing.T) {
		kf := NewKeyFinder("test_from", "")
		t.Run("match force from key", func(t *testing.T) {
			t.Run("match default from keys", func(t *testing.T) {
				i, found := kf.DetectFromIndex([]string{"created_at", "test_from"})
				require.Equal(t, 1, i)
				require.True(t, found)
			})
			t.Run("doesn't match default from keys", func(t *testing.T) {
				i, found := kf.DetectFromIndex([]string{"created_at"})
				require.Equal(t, 0, i)
				require.False(t, found)
			})
		})
	})
}
func TestIndexFinder_DetectToIndex(t *testing.T) {
	t.Run("force to empty", func(t *testing.T) {
		kf := NewKeyFinder("", "")
		t.Run("match default to keys", func(t *testing.T) {
			i, found := kf.DetectToIndex([]string{"updated_at"})
			require.Equal(t, 0, i)
			require.True(t, found)
		})
		t.Run("doesn't match default to keys", func(t *testing.T) {
			i, found := kf.DetectToIndex([]string{"hoge"})
			require.Equal(t, 0, i)
			require.False(t, found)
		})
	})
	t.Run("force to not empty", func(t *testing.T) {
		kf := NewKeyFinder("", "test_to")
		t.Run("match force to key", func(t *testing.T) {
			t.Run("match default to keys", func(t *testing.T) {
				i, found := kf.DetectToIndex([]string{"updated_at", "test_to"})
				require.Equal(t, 1, i)
				require.True(t, found)
			})
			t.Run("doesn't match default to keys", func(t *testing.T) {
				i, found := kf.DetectToIndex([]string{"updated_at"})
				require.Equal(t, 0, i)
				require.False(t, found)
			})
		})
	})
}
