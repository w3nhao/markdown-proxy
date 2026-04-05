package main

import (
	"flag"
	"reflect"
	"testing"
)

func TestCollectServerArgs_Basic(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	fs.String("port", "9080", "")
	fs.String("theme", "github", "")
	fs.Parse([]string{"-port=8080", "-theme=dark"})

	args := collectServerArgs(fs.Visit)
	want := []string{"-port=8080", "-theme=dark"}
	if !reflect.DeepEqual(args, want) {
		t.Errorf("got %v, want %v", args, want)
	}
}

func TestCollectServerArgs_SkipsVersion(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	fs.Bool("version", false, "")
	fs.String("port", "9080", "")
	fs.Parse([]string{"-version", "-port=8080"})

	args := collectServerArgs(fs.Visit)
	want := []string{"-port=8080"}
	if !reflect.DeepEqual(args, want) {
		t.Errorf("got %v, want %v", args, want)
	}
}

func TestCollectServerArgs_NoFlags(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	fs.String("port", "9080", "")
	fs.Parse([]string{})

	args := collectServerArgs(fs.Visit)
	if args != nil {
		t.Errorf("got %v, want nil", args)
	}
}
