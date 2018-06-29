package runner_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/JulzDiverse/resc/runner"
)

var _ = Describe("Exec", func() {

	var (
		runner *Runner
		script string
	)

	BeforeEach(func() {
		runner = New(Test)
		script = `#!/bin/bash

echo "$1"
echo "$2"`
	})

	It("should return the expected ouput", func() {
		output, err := runner.Run("echo hello")
		Expect(err).ToNot(HaveOccurred())
		Expect(strings.TrimSpace(string(output))).To(Equal("hello"))
	})

	It("should be able to pass parameters to the script", func() {
		output, err := runner.Run(script, "Hello", "World")
		Expect(err).ToNot(HaveOccurred())
		Expect(strings.TrimSpace(string(output))).To(Equal(`Hello
World`))
	})

	Context("When an empty script is provided", func() {
		It("it shouldn't error", func() {
			output, err := runner.Run("")
			Expect(err).ToNot(HaveOccurred())
			Expect(string(output)).To(Equal(""))
		})
	})
})
