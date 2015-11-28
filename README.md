# QUANGO: QUickcheck And giNkGO

A tiny [gomega](https://github.com/onsi/gomega) matcher for using
[quickcheck](https://golang.org/pkg/testing/quick) with
[ginkgo](https://github.com/onsi/ginkgo). Lets you write property-based tests
like this:


```golang
Describe("maths", func() {
  It("should work", func() {
    property := func(x, y float64) bool {
      return math.Abs(x)+math.Abs(y) >= math.Abs(x)
    }
    Expect(property).To(quango.Hold())
  })
})
```

