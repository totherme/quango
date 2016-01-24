package quango_test

import (
	"math"

	"github.com/totherme/quango"
)

var _ = Describe("The quango matcher", func() {
	Context("when actual is not a function", func() {
		It("should error", func() {
			_, err := quango.Hold().Match("not a function")
			Expect(err).To(HaveOccurred())
		})
	})

	Context("when actual represents a property that holds", func() {
		Context("when actual is of type bool-> bool", func() {
			It("should succeed", func() {
				Expect(func(b bool) bool {
					return b || !b
				}).To(quango.Hold())
			})
		})

		Context("when actual is of type float->float->bool", func() {
			It("should succeed", func() {
				Expect(func(x, y float64) bool { return math.Abs(x)+math.Abs(y) >= math.Abs(x) }).To(quango.Hold())
			})
		})

		Context("when actual is of type bool->void", func() {
			It("should succeed", func() {
				Expect(func(b bool) {}).To(quango.Hold())
			})
		})

		Context("when actual is of type float->float->void", func() {
			It("should succeed", func() {
				Expect(func(a, b float64) {}).To(quango.Hold())
			})
		})
	})

	Context("when actual represents a property that doesn't hold", func() {
		It("should fail", func() {
			Expect(func(b bool) bool {
				return b && !b
			}).NotTo(quango.Hold())
		})

		It("should give a counterexample in the fail message", func() {
			matcher := quango.Hold()
			prop := func(b bool) bool { return b }
			pass, err := matcher.Match(prop)
			Expect(pass).To(BeFalse())
			Expect(err).NotTo(HaveOccurred())
			Expect(matcher.FailureMessage(prop)).To(ContainSubstring("false"))

			prop = func(b bool) bool { return !b }
			pass, err = matcher.Match(prop)
			Expect(pass).To(BeFalse())
			Expect(err).NotTo(HaveOccurred())
			Expect(matcher.FailureMessage(prop)).To(ContainSubstring("true"))
		})

		Context("when actual is of type bool->void", func() {
			// Pended for now, since this will require a gomega PR to make work.
			PIt("should still fail", func() {
				Expect(func(b bool) {
					Expect(false).To(BeTrue())
				}).NotTo(quango.Hold())
			})
		})
	})
})
