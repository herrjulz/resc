package scriptmanager_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	. "github.com/JulzDiverse/resc/scriptmanager"
)

var _ = Describe("List", func() {

	var (
		scriptManager   *ScriptManager
		fakeServer      *ghttp.Server
		url, user, repo string
	)

	BeforeEach(func() {
		fakeServer = ghttp.NewServer()
		url = fakeServer.URL()
		user = "zeus"
		repo = "pandora"
	})

	JustBeforeEach(func() {
		scriptManager = New(url, user, repo)
	})

	Context("listing available scripts in a repository", func() {
		Context("When all directories are script directories", func() {

			BeforeEach(func() {
				fakeServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(
							"GET",
							"/repos/zeus/pandora/contents",
						),
						ghttp.RespondWith(
							http.StatusOK,
							`[{
						   "name": "script-one",
							 "type": "dir"
							},
							{
							  "name": "script-two",
								"type": "dir"
							},
							{
							  "name": "script-three",
								"type": "dir"
							}]`,
						),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(
							"GET",
							"/repos/zeus/pandora/contents/script-one/.resc",
						),
						ghttp.RespondWith(
							http.StatusOK,
							"{}",
						),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(
							"GET",
							"/repos/zeus/pandora/contents/script-two/.resc",
						),
						ghttp.RespondWith(
							http.StatusOK,
							"{}",
						),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(
							"GET",
							"/repos/zeus/pandora/contents/script-three/.resc",
						),
						ghttp.RespondWith(
							http.StatusOK,
							"{}",
						),
					),
				)

			})

			It("should list all available repositories", func() {
				scriptList, err := scriptManager.List()
				Expect(err).ToNot(HaveOccurred())
				Expect(scriptList).To(HaveLen(3))
				Expect(scriptList).To(Equal([]string{
					"script-one",
					"script-two",
					"script-three",
				}))
			})
		})

		Context("When some directories are no script directories", func() {

			BeforeEach(func() {
				fakeServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(
							"GET",
							"/repos/zeus/pandora/contents",
						),
						ghttp.RespondWith(
							http.StatusOK,
							`[{
						   "name": "script-one",
							 "type": "dir"
							},
							{
							  "name": "no-script",
								"type": "dir"
							},
							{
							  "name": "script-three",
								"type": "dir"
							}]`,
						),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(
							"GET",
							"/repos/zeus/pandora/contents/script-one/.resc",
						),
						ghttp.RespondWith(
							http.StatusOK,
							"{}",
						),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(
							"GET",
							"/repos/zeus/pandora/contents/no-script/.resc",
						),
						ghttp.RespondWith(
							http.StatusNotFound,
							"{}",
						),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(
							"GET",
							"/repos/zeus/pandora/contents/script-three/.resc",
						),
						ghttp.RespondWith(
							http.StatusOK,
							"{}",
						),
					),
				)
			})

			It("should list all available repositories", func() {
				scriptList, err := scriptManager.List()
				Expect(err).ToNot(HaveOccurred())
				Expect(scriptList).To(HaveLen(2))
				Expect(scriptList).To(Equal([]string{
					"script-one",
					"script-three",
				}))
			})
		})

		Context("When repository contains files", func() {
			BeforeEach(func() {
				fakeServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(
							"GET",
							"/repos/zeus/pandora/contents",
						),
						ghttp.RespondWith(
							http.StatusOK,
							`[{
						   "name": "script-one",
							 "type": "dir"
							},
							{
							  "name": "no-script",
								"type": "file"
							},
							{
							  "name": "script-three",
								"type": "dir"
							}]`,
						),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(
							"GET",
							"/repos/zeus/pandora/contents/script-one/.resc",
						),
						ghttp.RespondWith(
							http.StatusOK,
							"{}",
						),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(
							"GET",
							"/repos/zeus/pandora/contents/script-three/.resc",
						),
						ghttp.RespondWith(
							http.StatusOK,
							"{}",
						),
					),
				)

			})

			It("should list all available repositories", func() {
				scriptList, err := scriptManager.List()
				Expect(err).ToNot(HaveOccurred())
				Expect(scriptList).To(HaveLen(2))
				Expect(scriptList).To(Equal([]string{
					"script-one",
					"script-three",
				}))
			})
		})

		Context("When repository contains no valid script directories", func() {
			BeforeEach(func() {
				fakeServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(
							"GET",
							"/repos/zeus/pandora/contents",
						),
						ghttp.RespondWith(
							http.StatusOK,
							`[{
						   "name": "no-script-dir",
							 "type": "dir"
							},
							{
							  "name": "no-script",
								"type": "file"
							}]`,
						),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(
							"GET",
							"/repos/zeus/pandora/contents/no-script-dir/.resc",
						),
						ghttp.RespondWith(
							http.StatusNotFound,
							"{}",
						),
					),
				)
			})

			It("should list all available repositories", func() {
				scriptList, err := scriptManager.List()
				Expect(err).ToNot(HaveOccurred())
				Expect(scriptList).To(HaveLen(0))
				Expect(scriptList).To(Equal([]string{}))
			})
		})
	})
})
