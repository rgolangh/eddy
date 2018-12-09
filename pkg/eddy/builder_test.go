package eddy_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rgolangh/eddy/pkg/eddy"
)

var _ = Describe("Builder", func() {
	var (
		underTest eddy.UnitFile
	)

	Context("When creating basic unit ", func() {

		underTest = eddy.Basic()


		It("has a basic description", func() {
			Expect(underTest.Unit.Description).To(Equal("desc"))
		})

	})
})