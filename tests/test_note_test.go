package test

import (
	"SecondBrain/src/util"
	"testing"
)

// go test test_note_test.go

func TestNote(t *testing.T) {
	return
}

func TestFind(t *testing.T) {
	_, err := util.Exec("cd ..; ./jnoc checkout Playground")
	if err != nil {
		t.Errorf("%+v", err)
	}

	_, err = util.Exec("cd ..; ./jnoc find \"a\"")
	if err != nil {
		t.Errorf("%+v", err)
	}
}

func TestHistory(t *testing.T) {
	_, err := util.Exec("cd ..; ./jnoc history")
	if err != nil {
		t.Errorf("%+v", err)
	}
}

func TestList(t *testing.T) {
	_, err := util.Exec("cd ..; ./jnoc checkout Playground")
	if err != nil {
		t.Errorf("%+v", err)
	}

	_, err = util.Exec("cd ..; ./jnoc list")
	if err != nil {
		t.Errorf("%+v", err)
	}
}

func TestProject(t *testing.T) {
	_, err := util.Exec("cd ..; ./jnoc project")
	if err != nil {
		t.Errorf("%+v", err)
	}
}
