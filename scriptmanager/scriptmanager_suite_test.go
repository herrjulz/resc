package scriptmanager_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestScriptmanager(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scriptmanager Suite")
}
