package fasthttp

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/vafinvr/go-admin/tests/common"
)

func TestFasthttp(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewFastBinder(newHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}
