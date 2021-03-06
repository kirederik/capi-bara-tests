package v3_helpers

import (
	"strings"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
)

func GetAuthToken() string {
	session := cf.Cf("oauth-token")
	bytes := session.Wait().Out.Contents()
	return strings.TrimSpace(string(bytes))
}
