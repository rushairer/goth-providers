package wechat

import (
	"os"
	"testing"

	"github.com/markbates/goth"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	p := provider()

	a.Equal(p.ClientId, os.Getenv("WECHAT_KEY"))
	a.Equal(p.ClientSecret, os.Getenv("WECHAT_SECRET"))
	a.Equal(p.RedirectUrl, "/foo")
}

func Test_Implements_Provider(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	a.Implements((*goth.Provider)(nil), provider())
}

func Test_BeginAuth(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	p := provider()
	session, err := p.BeginAuth("test_state")
	s := session.(*Session)
	a.NoError(err)
	a.Contains(s.AuthUrl, "open.weixin.qq.com/connect/qrconnect")
}

func Test_SessionFromJSON(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	p := provider()
	session, err := p.UnmarshalSession(`{"AuthUrl":"https://open.weixin.qq.com/connect/qrconnect","AccessToken":"1234567890"}`)
	a.NoError(err)

	s := session.(*Session)
	a.Equal(s.AuthUrl, "https://open.weixin.qq.com/connect/qrconnect")
	a.Equal(s.AccessToken, "1234567890")
}

func provider() *Provider {
	return New(os.Getenv("WECHAT_KEY"), os.Getenv("WECHAT_SECRET"), "/foo", WECHAT_LANG_CN)
}
