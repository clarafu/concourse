package secretsmanager_test

import (
	"github.com/concourse/concourse/atc/atccmd"
	"github.com/concourse/concourse/atc/creds/secretsmanager"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Manager", func() {
	var manager secretsmanager.Manager

	Describe("Validate()", func() {
		BeforeEach(func() {
			manager = atccmd.CmdDefaults.CredentialManagers.SecretsManager
			manager.AwsRegion = "test-region"

			Expect(manager.PipelineSecretTemplate).To(Equal(secretsmanager.DefaultPipelineSecretTemplate))
			Expect(manager.TeamSecretTemplate).To(Equal(secretsmanager.DefaultTeamSecretTemplate))
		})

		It("passes on default parameters", func() {
			Expect(manager.Validate()).To(BeNil())
		})

		Context("when region is not set", func() {
			BeforeEach(func() {
				manager.AwsRegion = ""
			})

			It("fails validation", func() {
				Expect(manager.Validate()).ToNot(BeNil())
			})
		})

		DescribeTable("passes if all aws credentials are specified",
			func(accessKey, secretKey, sessionToken string) {
				manager.AwsAccessKeyID = accessKey
				manager.AwsSecretAccessKey = secretKey
				manager.AwsSessionToken = sessionToken
				Expect(manager.Validate()).To(BeNil())
			},
			Entry("all values", "access", "secret", "token"),
			Entry("access & secret", "access", "secret", ""),
		)

		DescribeTable("fails on partial AWS credentials",
			func(accessKey, secretKey, sessionToken string) {
				manager.AwsAccessKeyID = accessKey
				manager.AwsSecretAccessKey = secretKey
				manager.AwsSessionToken = sessionToken
				Expect(manager.Validate()).ToNot(BeNil())
			},
			Entry("only access", "access", "", ""),
			Entry("access & token", "access", "", "token"),
			Entry("only secret", "", "secret", ""),
			Entry("secret & token", "", "secret", "token"),
			Entry("only token", "", "", "token"),
		)

		It("passes on pipe secret template containing less specialization", func() {
			manager.PipelineSecretTemplate = "{{.Secret}}"
			Expect(manager.Validate()).To(BeNil())
		})

		It("passes on pipe secret template containing no specialization", func() {
			manager.PipelineSecretTemplate = "var"
			Expect(manager.Validate()).To(BeNil())
		})

		It("fails on empty pipe secret template", func() {
			manager.PipelineSecretTemplate = ""
			Expect(manager.Validate()).ToNot(BeNil())
		})

		It("fails on pipe secret template containing invalid parameters", func() {
			manager.PipelineSecretTemplate = "{{.Teams}}"
			Expect(manager.Validate()).ToNot(BeNil())
		})

		It("passes on team secret template containing less specialization", func() {
			manager.TeamSecretTemplate = "{{.Secret}}"
			Expect(manager.Validate()).To(BeNil())
		})

		It("passes on team secret template containing no specialization", func() {
			manager.TeamSecretTemplate = "var"
			Expect(manager.Validate()).To(BeNil())
		})

		It("fails on empty team secret template", func() {
			manager.TeamSecretTemplate = ""
			Expect(manager.Validate()).ToNot(BeNil())
		})

		It("fails on team secret template containing invalid parameters", func() {
			manager.TeamSecretTemplate = "{{.Teams}}"
			Expect(manager.Validate()).ToNot(BeNil())
		})
	})
})
