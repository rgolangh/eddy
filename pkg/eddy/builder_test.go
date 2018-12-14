package eddy_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rgolangh/eddy/pkg/eddy"
)

var _ = Describe("Builder", func() {

	Context("When creating basic unit ", func() {

		var underTest = eddy.Basic()

		It("has a basic description", func() {
			Expect(underTest.Unit.Description).To(Equal("desc"))
		})

		Specify("is required by no other", func() {
			Expect(underTest.Unit.Requires).To(BeEmpty())
		})

		Specify("is as After non specified", func() {
			Expect(underTest.Unit.After).To(BeEmpty())
		})

	})

	Context("When passing argument", func() {
		var underTest = eddy.UnitFile{
			Unit: eddy.Unit{
				Description: "test unit",
				Requires:    "parent-unitfile",
				After:       "sybling-unitfile",
			},
			Install: eddy.Install{
				WantedBy:   "wanted-by-other",
				RequiredBy: "required-by-other",
			},
			Section: struct{}{},
		}

		Specify("description is set", func() {
			Expect(underTest.Unit.Description).To(Equal("test unit"))
		})
		Specify("Requires is set", func() {
			Expect(underTest.Unit.Requires).To(Equal("parent-unitfile"))
		})
		Specify("After is set", func() {
			Expect(underTest.Unit.After).To(Equal("sybling-unitfile"))
		})

		Context("When converting to ini file", func() {

			iniFile, e := eddy.ToIniFile(underTest)

			It("returns an ini file", func() {
				Expect(iniFile).NotTo(BeNil())
			})
			It("didn't fail to convert", func() {
				Expect(e).NotTo(HaveOccurred())
			})

			Describe("Unit section tests", func() {

				unitSection, err := iniFile.GetSection("Unit")

				It("exists", func() {
					Expect(err).NotTo(HaveOccurred())
				})
				It("has matching Description ", func() {
					Expect(unitSection.Key("Description").Value()).To(Equal("test unit"))
				})
			})

			Describe("Install section tests", func() {

				unitSection, err := iniFile.GetSection("Install")

				It("exists", func() {
					Expect(err).NotTo(HaveOccurred())
				})
				It("has matching WantedBy ", func() {
					Expect(unitSection.Key("WantedBy").Value()).To(Equal("wanted-by-other"))
				})
				It("has matching RequiredBy ", func() {
					Expect(unitSection.Key("RequiredBy").Value()).To(Equal("required-by-other"))
				})
			})

		})
	})

})
