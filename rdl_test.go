package main

import (
	"fmt"
	"github.com/l-pdufour/pokedexcli/internal/cache"
	"testing"
	"time"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "HELLO world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)

		if len(actual) != len(cs.expected) {
			t.Errorf("Lengths are not equal: %d vs %d", len(actual), len(cs.expected))
		}

		for i := 0; i < len(actual); i++ {
			if actual[i] != cs.expected[i] {
				t.Errorf("Mismatch at index %d: got %s, expected %s", i, actual[i], cs.expected[i])
			}
		}
	}
}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := cache.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := cache.NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestAddAndGet(t *testing.T) {
	expirationInterval := 1 * time.Second
	cache := cache.NewCache(expirationInterval)

	key := "example"
	value := []byte("test value")

	cache.Add(key, value)

	got, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected to find key: %s", key)
	}

	if string(got) != string(value) {
		t.Errorf("expected value: %s, got: %s", value, got)
	}
}

func TestExpiredEntry(t *testing.T) {
	expirationInterval := 100 * time.Millisecond
	cache := cache.NewCache(expirationInterval)

	key := "example"
	value := []byte("test value")

	cache.Add(key, value)

	time.Sleep(2 * expirationInterval)

	_, ok := cache.Get(key)
	if ok {
		t.Errorf("expected key to be expired: %s", key)
	}
}
