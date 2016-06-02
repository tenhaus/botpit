package useraccounts

import (
  "testing"
)

func TestCreateDeleteUserAccount(t *testing.T) {
  var account UserAccount

  if err := Create("TestyTesterson1134", &account); err != nil {
    t.Error(err)
  }

  if err := Delete(account.Key); err != nil {
    t.Error(err)
  }
}

// Make sure short names produce an error
func TestRejectShortName(t *testing.T) {
  t.Error()
}

// Make sure long names produce an error
func TestRejectLongName(t *testing.T) {
  t.Error()
}