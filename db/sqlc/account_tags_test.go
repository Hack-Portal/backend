package db

import "testing"

func createAccountTagsTest(t *testing.T) AccountTags {
	account := createAccountTest(t)
	tags := getAllTag(t)

}
