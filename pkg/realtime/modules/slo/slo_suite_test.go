package slo_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var ctrl *gomock.Controller

func TestSlo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Slo Suite")
}

var _ = BeforeSuite(func() {
	ctrl = gomock.NewController(GinkgoT())
})
