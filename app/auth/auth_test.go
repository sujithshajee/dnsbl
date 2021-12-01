package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	scrypt "github.com/elithrar/simple-scrypt"
	_ "github.com/mattn/go-sqlite3"
	gm "github.com/onsi/gomega"
	"github.com/sujithshajee/dnsbl/app/ent"
	"github.com/sujithshajee/dnsbl/app/ent/enttest"
)

func Test_GenerateUser(t *testing.T) {
	tcs := map[string]struct {
		client   *ent.Client
		setup    func(context.Context, *ent.Client)
		hasError bool
	}{
		"already exists": {
			client: enttest.Open(t, "sqlite3", "file:exists?mode=memory&cache=shared&_fk=1"),
			setup: func(ctx context.Context, cl *ent.Client) {
				cl.User.Create().SetUsername("username").SetPassword([]byte("pw")).SaveX(ctx)
			},
			hasError: true,
		},
		"success": {
			client: enttest.Open(t, "sqlite3", "file:success?mode=memory&cache=shared&_fk=1"),
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			gm.RegisterTestingT(t)

			ctx := context.Background()
			defer tc.client.Close()
			tx, err := tc.client.Tx(ctx)
			gm.Expect(err).NotTo(gm.HaveOccurred())
			cl := tx.Client()

			if tc.setup != nil {
				tc.setup(ctx, cl)
			}

			err = GenerateUser(ctx, cl, "username", "pw")
			if tc.hasError {
				gm.Expect(err).To(gm.HaveOccurred())
			} else {
				gm.Expect(err).NotTo(gm.HaveOccurred())
			}
		})
	}
}

func Test_CheckPassword(t *testing.T) {
	tcs := map[string]struct {
		pass     string
		hash     func() []byte
		hasError bool
	}{
		"bad": {
			pass:     "pass",
			hash:     func() []byte { return []byte("not hashed") },
			hasError: true,
		},
		"ok": {
			pass: "pass",
			hash: func() []byte {
				hash, _ := scrypt.GenerateFromPassword([]byte("pass"), scrypt.DefaultParams)
				return hash
			},
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			gm.RegisterTestingT(t)

			err := CheckPassword(tc.hash(), []byte(tc.pass))
			if tc.hasError {
				gm.Expect(err).To(gm.HaveOccurred())
			} else {
				gm.Expect(err).NotTo(gm.HaveOccurred())
			}
		})
	}
}

func Test_BasicAuth(t *testing.T) {
	authedReq := func(c *http.Client, url string) (*http.Response, error) {
		req, _ := http.NewRequest(http.MethodGet, url+"/path", nil)
		req.SetBasicAuth("username", "pw")
		return c.Do(req)
	}

	tcs := map[string]struct {
		client   *ent.Client
		setup    func(context.Context, *ent.Client)
		req      func(*http.Client, string) (*http.Response, error)
		hasError bool
	}{
		"ok": {
			client: enttest.Open(t, "sqlite3", "file:ok?mode=memory&cache=shared&_fk=1"),
			setup: func(ctx context.Context, cl *ent.Client) {
				_ = GenerateUser(ctx, cl, "username", "pw")
			},
			req: authedReq,
		},
		"bad auth": {
			client:   enttest.Open(t, "sqlite3", "file:bad_auth?mode=memory&cache=shared&_fk=1"),
			req:      authedReq,
			hasError: true,
		},
		"no auth": {
			req: func(c *http.Client, url string) (*http.Response, error) {
				return c.Get(url + "/path")
			},
			hasError: true,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			gm.RegisterTestingT(t)
			ctx := context.Background()

			if tc.setup != nil {
				tc.setup(ctx, tc.client)
			}

			s := httptest.NewServer(
				BasicAuth(tc.client, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
					_, _ = rw.Write([]byte("ok"))
				})),
			)
			defer s.Close()

			cl := s.Client()

			if tc.req != nil {
				r, _ := tc.req(cl, s.URL)
				if tc.hasError {
					gm.Expect(r.StatusCode).To(gm.Equal(http.StatusUnauthorized))
				} else {
					gm.Expect(r.StatusCode).To(gm.Equal(http.StatusOK))
				}
			} else {
				t.Fatal("request undefined")
			}
		})
	}
}
