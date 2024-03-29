// Code generated by protoc-gen-micro
// source: example.proto
package pb

import (
	context "context"
	v3 "github.com/unistack-org/micro-client-http/v3"
	api "github.com/unistack-org/micro/v3/api"
	client "github.com/unistack-org/micro/v3/client"
	codec "github.com/unistack-org/micro/v3/codec"
	server "github.com/unistack-org/micro/v3/server"
	http "net/http"
)

type userClient struct {
	c    client.Client
	name string
}

func NewUserClient(name string, c client.Client) UserClient {
	return &userClient{c: c, name: name}
}

func (c *userClient) FindAll(ctx context.Context, req *FindAllRequest, opts ...client.CallOption) (*FindAllResponse, error) {
	errmap := make(map[string]interface{}, 2)
	errmap["500"] = &Error{}
	errmap["200"] = &FindAllResponse{}
	opts = append(opts,
		v3.ErrorMap(errmap),
	)
	opts = append(opts,
		v3.Method(http.MethodGet),
		v3.Path("/api/v1/users"),
	)
	rsp := &FindAllResponse{}
	err := c.c.Call(ctx, c.c.NewRequest(c.name, "User.FindAll", req), rsp, opts...)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *userClient) FindByID(ctx context.Context, req *FindByIDRequest, opts ...client.CallOption) (*FindByIDResponse, error) {
	errmap := make(map[string]interface{}, 3)
	errmap["200"] = &FindByIDResponse{}
	errmap["404"] = &Error{}
	errmap["500"] = &Error{}
	opts = append(opts,
		v3.ErrorMap(errmap),
	)
	opts = append(opts,
		v3.Method(http.MethodGet),
		v3.Path("/api/v1/users/{user_id}"),
	)
	rsp := &FindByIDResponse{}
	err := c.c.Call(ctx, c.c.NewRequest(c.name, "User.FindByID", req), rsp, opts...)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *userClient) DownloadUserFile(ctx context.Context, req *DownloadRequest, opts ...client.CallOption) (*codec.Frame, error) {
	errmap := make(map[string]interface{}, 3)
	errmap["404"] = &Error{}
	errmap["500"] = &Error{}
	opts = append(opts,
		v3.ErrorMap(errmap),
	)
	opts = append(opts,
		v3.Method(http.MethodGet),
		v3.Path("/api/v1/users/{user_id}/file"),
	)
	rsp := &codec.Frame{}
	err := c.c.Call(ctx, c.c.NewRequest(c.name, "User.DownloadUserFile", req), rsp, opts...)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

type userServer struct {
	UserServer
}

func (h *userServer) FindAll(ctx context.Context, req *FindAllRequest, rsp *FindAllResponse) error {
	return h.UserServer.FindAll(ctx, req, rsp)
}

func (h *userServer) FindByID(ctx context.Context, req *FindByIDRequest, rsp *FindByIDResponse) error {
	return h.UserServer.FindByID(ctx, req, rsp)
}

func (h *userServer) DownloadUserFile(ctx context.Context, req *DownloadRequest, rsp *codec.Frame) error {
	return h.UserServer.DownloadUserFile(ctx, req, rsp)
}

func RegisterUserServer(s server.Server, sh UserServer, opts ...server.HandlerOption) error {
	type user interface {
		FindAll(ctx context.Context, req *FindAllRequest, rsp *FindAllResponse) error
		FindByID(ctx context.Context, req *FindByIDRequest, rsp *FindByIDResponse) error
		DownloadUserFile(ctx context.Context, req *DownloadRequest, rsp *codec.Frame) error
	}
	type User struct {
		user
	}
	h := &userServer{sh}
	var nopts []server.HandlerOption
	for _, endpoint := range NewUserEndpoints() {
		nopts = append(nopts, api.WithEndpoint(endpoint))
	}
	return s.Handle(s.NewHandler(&User{h}, append(nopts, opts...)...))
}
