package quango_test

import (
	. "github.com/totherme/quango"
)

var _ = Describe("Quango quantifiers", func() {
	Describe("for all", func() {
		Context("when called on an empty list", func() {

			It("doesn't call the body", func() {
				list := make([]int, 0)
				ForAll(list, func(_ int) bool {
					Fail("This shouldn't be called")
					return false
				})
			})

			It("returns true", func() {
				list := make([]int, 0)
				Expect(ForAll(list, func(_ int) bool {
					return false
				})).To(BeTrue())
			})
		})

		Context("when called on a list of length n ", func() {
			Context("with a true body", func() {
				It("calls the body n times", func() {
					Expect(func(list []int) bool {
						n := len(list)
						count := 0
						ForAll(list, func(_ int) bool {
							count++
							return true
						})
						return count == n
					}).To(Hold())
				})
			})

			Context("with a false body", func() {
				It("calls the body only once", func() {
					Expect(func(list []int) bool {
						if len(list) == 0 {
							return true
						}
						count := 0
						ForAll(list, func(_ int) bool {
							count++
							return false
						})
						return count == 1
					}).To(Hold())
				})
			})
		})

		Context("when called with a body that returns true", func() {
			It("returns true", func() {
				list := make([]int, 1)
				Expect(ForAll(list, func(_ int) bool {
					return true
				})).To(BeTrue())
			})
		})

		Context("when called with a body that returns false", func() {
			It("returns false", func() {
				list := make([]int, 1)
				Expect(ForAll(list, func(_ int) bool {
					return false
				})).To(BeFalse())
			})
		})

		Context("when called with a body that sometimes returns true, and sometimes false", func() {
			It("returns false", func() {
				list := []int{1, 0, 1, 0, 1, 0}
				Expect(ForAll(list, func(n int) bool {
					return n == 0
				})).To(BeFalse())
			})
		})
	})
})
