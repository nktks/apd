package main

import "strings"

var (
	normalizedFromKeys = []string{
		"createdat",
		"creationtime",
		"starttime",
		"start",
		"startedat",
		"startat",
		"from",
		"fromat",
		"fromtime",
	}
	normalizedToKeys = []string{
		"updatedat",
		"statetime",
		"completedtime",
		"completedat",
		"end",
		"endedat",
		"endat",
		"closeat",
		"closedat",
		"endtime",
		"to",
	}
	defaultDurationKey = "duration"
)

type KeyFinder struct {
	forceFrom   string
	forceTo     string
	DurationKey string
}

func NewKeyFinder(from, to string) *KeyFinder {
	return &KeyFinder{
		forceFrom:   from,
		forceTo:     to,
		DurationKey: defaultDurationKey,
	}
}

func (kf *KeyFinder) DetectFromKey(keys []string) string {
	if kf.forceFrom != "" {
		for _, k := range keys {
			if normalize(kf.forceFrom) == normalize(k) {
				return k
			}
		}
		return ""
	} else {
		for _, k := range keys {
			for _, nk := range normalizedFromKeys {
				if nk == normalize(k) {
					return k
				}
			}
		}
	}

	return ""
}
func (kf *KeyFinder) DetectToKey(keys []string) string {
	if kf.forceTo != "" {
		for _, k := range keys {
			if normalize(kf.forceTo) == normalize(k) {
				return k
			}
		}
		return ""
	} else {
		for _, k := range keys {
			for _, nk := range normalizedToKeys {
				if nk == normalize(k) {
					return k
				}
			}
		}
	}
	return ""
}

func (kf *KeyFinder) DetectFromIndex(keys []string) (int, bool) {
	if kf.forceFrom != "" {
		for i, k := range keys {
			if normalize(kf.forceFrom) == normalize(k) {
				return i, true
			}
		}
		return 0, false
	} else {
		for i, k := range keys {
			for _, nk := range normalizedFromKeys {
				if nk == normalize(k) {
					return i, true
				}
			}
		}
	}

	return 0, false
}
func (kf *KeyFinder) DetectToIndex(keys []string) (int, bool) {
	if kf.forceTo != "" {
		for i, k := range keys {
			if normalize(kf.forceTo) == normalize(k) {
				return i, true
			}
		}
		return 0, false
	} else {
		for i, k := range keys {
			for _, nk := range normalizedToKeys {
				if nk == normalize(k) {
					return i, true
				}
			}
		}
	}

	return 0, false
}

func normalize(k string) string {
	k = strings.Replace(k, "-", "", -1)
	k = strings.Replace(k, "_", "", -1)
	return strings.ToLower(k)
}
