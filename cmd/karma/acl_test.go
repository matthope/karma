package main

import (
	"regexp"
	"testing"

	"github.com/prymitive/karma/internal/models"
)

func TestAclSilenceMatcher(t *testing.T) {
	type testCaseT struct {
		requiredMatcher silenceMatcher
		silenceMatcher  models.SilenceMatcher
		isMatch         bool
	}

	testCases := []testCaseT{
		{
			requiredMatcher: silenceMatcher{
				Name:  "foo",
				Value: "bar",
			},
			silenceMatcher: models.SilenceMatcher{
				Name:  "foo",
				Value: "bar",
			},
			isMatch: true,
		},
		{
			requiredMatcher: silenceMatcher{
				Name:  "foo",
				Value: "bar",
			},
			silenceMatcher: models.SilenceMatcher{
				Name:  "bar",
				Value: "bar",
			},
			isMatch: false,
		},
		{
			requiredMatcher: silenceMatcher{
				Name:  "foo",
				Value: "bar",
			},
			silenceMatcher: models.SilenceMatcher{
				Name:  "foo",
				Value: "foo",
			},
			isMatch: false,
		},
		{
			requiredMatcher: silenceMatcher{
				Name:    "foo",
				Value:   "bar",
				IsRegex: true,
			},
			silenceMatcher: models.SilenceMatcher{
				Name:  "foo",
				Value: "bar",
			},
			isMatch: false,
		},
		{
			requiredMatcher: silenceMatcher{
				Name:    "foo",
				Value:   "bar",
				IsRegex: false,
			},
			silenceMatcher: models.SilenceMatcher{
				Name:    "foo",
				Value:   "bar",
				IsRegex: true,
			},
			isMatch: false,
		},
		{
			requiredMatcher: silenceMatcher{
				NameRegex: regexp.MustCompile("^.+$"),
				Value:     "bar",
			},
			silenceMatcher: models.SilenceMatcher{
				Name:  "foo",
				Value: "bar",
			},
			isMatch: true,
		},
		{
			requiredMatcher: silenceMatcher{
				NameRegex: regexp.MustCompile("^notfoo$"),
				Value:     "bar",
			},
			silenceMatcher: models.SilenceMatcher{
				Name:  "foo",
				Value: "bar",
			},
			isMatch: false,
		},
		{
			requiredMatcher: silenceMatcher{
				NameRegex:  regexp.MustCompile("^.+$"),
				ValueRegex: regexp.MustCompile("^.+$"),
			},
			silenceMatcher: models.SilenceMatcher{
				Name:  "foo",
				Value: "bar",
			},
			isMatch: true,
		},
		{
			requiredMatcher: silenceMatcher{
				NameRegex:  regexp.MustCompile("^.+$"),
				ValueRegex: regexp.MustCompile("^bar.+$"),
			},
			silenceMatcher: models.SilenceMatcher{
				Name:  "foo",
				Value: "notbar",
			},
			isMatch: false,
		},
		{
			requiredMatcher: silenceMatcher{
				NameRegex:  regexp.MustCompile("^.+$"),
				ValueRegex: regexp.MustCompile("^.+$"),
				IsRegex:    true,
			},
			silenceMatcher: models.SilenceMatcher{
				Name:    "foo",
				Value:   "bar",
				IsRegex: false,
			},
			isMatch: false,
		},
		{
			requiredMatcher: silenceMatcher{
				Name:    "foo",
				Value:   "bar",
				IsRegex: true,
				IsEqual: false,
			},
			silenceMatcher: models.SilenceMatcher{
				Name:    "foo",
				Value:   "bar",
				IsRegex: true,
				IsEqual: false,
			},
			isMatch: true,
		},
		{
			requiredMatcher: silenceMatcher{
				Name:    "foo",
				Value:   "bar",
				IsRegex: true,
				IsEqual: true,
			},
			silenceMatcher: models.SilenceMatcher{
				Name:    "foo",
				Value:   "bar",
				IsRegex: true,
				IsEqual: true,
			},
			isMatch: true,
		},
		{
			requiredMatcher: silenceMatcher{
				Name:    "foo",
				Value:   "bar",
				IsEqual: false,
			},
			silenceMatcher: models.SilenceMatcher{
				Name:    "foo",
				Value:   "bar",
				IsEqual: true,
			},
			isMatch: false,
		},
		{
			requiredMatcher: silenceMatcher{
				Name:    "foo",
				Value:   "bar",
				IsEqual: true,
			},
			silenceMatcher: models.SilenceMatcher{
				Name:    "foo",
				Value:   "bar",
				IsEqual: false,
			},
			isMatch: false,
		},
	}

	for _, testCase := range testCases {
		isMatch := testCase.requiredMatcher.isMatch(testCase.silenceMatcher)
		if isMatch != testCase.isMatch {
			t.Errorf("isMatch() returned %v, expected %v, requiredMatcher=%v silenceMatcher=%v", isMatch, testCase.isMatch, testCase.requiredMatcher, testCase.silenceMatcher)
		}
	}
}
