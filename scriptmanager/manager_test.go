package scriptmanager_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	. "github.com/JulzDiverse/resc/scriptmanager"
)

var _ = Describe("Manager", func() {

	Describe("", func() {

		var (
			scriptManager                         *ScriptManager
			fakeServer                            *ghttp.Server
			url, user, repo, scriptName, expected string
		)

		BeforeEach(func() {
			fakeServer = ghttp.NewServer()
			url = fakeServer.URL()
			user = "zeus"
			repo = "pandora"
			scriptName = "open-pandora"

		})

		JustBeforeEach(func() {
			scriptManager = New(url, user, repo)
		})

		Context("When requesting a script", func() {
			Context("When the request can be performed", func() {
				BeforeEach(func() {
					fakeServer.AppendHandlers(
						ghttp.CombineHandlers(
							ghttp.VerifyRequest(
								"GET",
								"/zeus/pandora/master/open-pandora/run.sh",
							),
							ghttp.RespondWith(
								http.StatusOK,
								expected,
							),
						),
					)

					expected = `!/bin/bash
echo "pandora opened"`
				})

				It("should send the request", func() {
					_, err := scriptManager.GetScript(scriptName)
					Expect(err).ToNot(HaveOccurred())
					Expect(fakeServer.ReceivedRequests()).Should(HaveLen(1))
				})

				It("should return the expected response", func() {
					script, err := scriptManager.GetScript(scriptName)
					Expect(err).ToNot(HaveOccurred())
					Expect(string(script)).To(Equal(expected))
				})
			})

			Context("When response status is other than 200 OK", func() {
				BeforeEach(func() {
					fakeServer.AppendHandlers(
						ghttp.RespondWith(
							http.StatusNotFound,
							"",
						),
					)
				})

				It("should error", func() {
					_, err := scriptManager.GetScript(scriptName)
					Expect(err).To(HaveOccurred())
				})

				It("should return an nil slice", func() {
					script, err := scriptManager.GetScript(scriptName)
					Expect(err).To(HaveOccurred())
					Expect(script).To(BeNil())
				})

				It("should return a meaningful message", func() {
					_, err := scriptManager.GetScript(scriptName)
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError(ContainSubstring("requesting file failed: 404 Not Found")))
				})
			})
		})

		Context("When requesting a README", func() {
			Context("When the request can be performed", func() {
				BeforeEach(func() {
					fakeServer.AppendHandlers(
						ghttp.CombineHandlers(
							ghttp.VerifyRequest(
								"GET",
								"/zeus/pandora/master/open-pandora/README.md",
							),
							ghttp.RespondWith(
								http.StatusOK,
								expected,
							),
						),
					)

					expected = "# Topic"
				})

				It("should send the request", func() {
					_, err := scriptManager.GetReadmeForScript(scriptName)
					Expect(err).ToNot(HaveOccurred())
					Expect(fakeServer.ReceivedRequests()).Should(HaveLen(1))
				})

				It("should return the expected response", func() {
					readme, err := scriptManager.GetReadmeForScript(scriptName)
					Expect(err).ToNot(HaveOccurred())
					Expect(string(readme)).To(Equal(expected))
				})
			})

			Context("When response status is other than 200 OK", func() {
				BeforeEach(func() {
					fakeServer.AppendHandlers(
						ghttp.RespondWith(
							http.StatusNotFound,
							"",
						),
					)
				})

				It("should error", func() {
					_, err := scriptManager.GetReadmeForScript(scriptName)
					Expect(err).To(HaveOccurred())
				})

				It("should return an nil slice", func() {
					readme, err := scriptManager.GetReadmeForScript(scriptName)
					Expect(err).To(HaveOccurred())
					Expect(readme).To(BeNil())
				})

				It("should return a meaningful message", func() {
					_, err := scriptManager.GetReadmeForScript(scriptName)
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError(ContainSubstring("requesting file failed: 404 Not Found")))
				})
			})
		})
	})
})
