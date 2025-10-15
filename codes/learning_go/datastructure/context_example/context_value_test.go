package context_example

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

//
// Context values
//

// store user id in context
// use GUID to track service chain

type userKey int

// or:
// type userKey2 struct{}

type guidKey int

const _guidKey guidKey = 1

const (
	_ userKey = iota
	_userKey
)

const (
	UserIdCookieKey = "identity"
	GUIDHeaderKey   = "X-GUID"
)

func ContextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, _userKey, user)
}

func UserFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(_userKey).(string)
	return user, ok
}

// extract user id from cookie
func extractUserId(r *http.Request) (string, error) {
	userCookie, err := r.Cookie(UserIdCookieKey)
	if err != nil {
		return "", err
	}
	return userCookie.Value, nil
}

func contextWithGUID(ctx context.Context, guid string) context.Context {
	return context.WithValue(ctx, _guidKey, guid)
}

func guidFromContext(ctx context.Context) (string, bool) {
	guid, ok := ctx.Value(_guidKey).(string)
	return guid, ok
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user id
		userId, err := extractUserId(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
			return
		}

		ctx := r.Context()
		// guid
		if guid := r.Header.Get(GUIDHeaderKey); guid != "" {
			ctx = contextWithGUID(ctx, guid)
		} else {
			ctx = contextWithGUID(ctx, uuid.New().String())
		}

		ctx = ContextWithUser(ctx, userId)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

type Logger interface {
	Log(context.Context, string)
}

type RequestDecorator func(*http.Request) *http.Request

type StdoutLogger struct{}

type Logic interface {
	Process(ctx context.Context, data string) (string, error)
}

type LogicImpl struct {
	RequestDecorator RequestDecorator
	Logger           Logger
	Remote           string
}

func (l StdoutLogger) Log(ctx context.Context, message string) {
	if guid, ok := guidFromContext(ctx); ok {
		message = fmt.Sprintf("GUID: %s - %s", guid, message)
	}
	fmt.Println(message)
}

func Request(r *http.Request) *http.Request {
	ctx := r.Context()
	if guid, ok := guidFromContext(ctx); ok {
		r.Header.Add(GUIDHeaderKey, guid)
	}
	return r
}

func (l LogicImpl) Process(ctx context.Context, data string) (string, error) {
	l.Logger.Log(ctx, "start process with "+data)
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, l.Remote+"/get?query="+data, nil)
	if err != nil {
		l.Logger.Log(ctx, "error building remote request: "+err.Error())
		return "", err
	}

	r = l.RequestDecorator(r)
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		l.Logger.Log(ctx, "error making remote request: "+err.Error())
		return "", err
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		l.Logger.Log(ctx, "error parse response body: "+err.Error())
		return "", err
	}
	return string(respData), nil
}

type Controller struct {
	logic Logic
}

func (c Controller) Process(data string) (string, error) {
	ctx := context.Background()
	message, err := c.logic.Process(ctx, data)
	return message, err
}

func NewController() Controller {
	return Controller{
		logic: LogicImpl{
			RequestDecorator: Request,
			Logger:           StdoutLogger{},
			// see https://github.com/postmanlabs/httpbin
			Remote: "https://httpbin.org",
		},
	}
}

func TestContextValue(t *testing.T) {
	c := NewController()
	message, err := c.Process("alice")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(message)
}
