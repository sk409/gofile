package gofile_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGofile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gofile Suite")
}
