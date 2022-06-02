package RevenueCalculatorService_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRevenueCalculatorService(t *testing.T) {
	RegisterFailHandler(Fail) //glue code to connect ginkgo to gomega#telling our Gomega which function to call in the event of failure occured
	//above means gomega.RegisterFailHandler(ginkgo.Fail)
	RunSpecs(t, "RevenueCalculatorService Suite")
}
