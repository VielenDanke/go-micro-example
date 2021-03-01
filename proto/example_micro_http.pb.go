// Code generated by protoc-gen-micro
// source: example.proto
package pb

import (
	context "context"
	_ "github.com/unistack-org/micro-client-http/v3"
	api "github.com/unistack-org/micro/v3/api"
	client "github.com/unistack-org/micro/v3/client"
	server "github.com/unistack-org/micro/v3/server"
)

type postClient struct {
	c    client.Client
	name string
}

func NewPostClient(name string, c client.Client) PostClient {
	return &postClient{c: c, name: name}
}

func (c *postClient) FindByID(ctx context.Context, req *FindByIDRequest, opts ...client.CallOption) (*FindByIDResponse, error) {
	rsp := &FindByIDResponse{}
	err := c.c.Call(ctx, c.c.NewRequest(c.name, "Post.FindByID", req), rsp, opts...)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *postClient) GetPostFileByID(ctx context.Context, req *GetPostFileRequest, opts ...client.CallOption) (Post_GetPostFileByIDClient, error) {
	stream, err := c.c.Stream(ctx, c.c.NewRequest(c.name, "Post.GetPostFileByID", &GetPostFileRequest{}), opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(req); err != nil {
		return nil, err
	}
	return &postClientGetPostFileByID{stream}, nil
}

type postClientGetPostFileByID struct {
	stream client.Stream
}

func (s *postClientGetPostFileByID) Close() error {
	return s.stream.Close()
}

func (s *postClientGetPostFileByID) Context() context.Context {
	return s.stream.Context()
}

func (s *postClientGetPostFileByID) SendMsg(msg interface{}) error {
	return s.stream.Send(msg)
}

func (s *postClientGetPostFileByID) RecvMsg(msg interface{}) error {
	return s.stream.Recv(msg)
}

func (s *postClientGetPostFileByID) Recv() (*GetPostFileResponse, error) {
	msg := &GetPostFileResponse{}
	if err := s.stream.Recv(msg); err != nil {
		return nil, err
	}
	return msg, nil
}

type postServer struct {
	PostServer
}

func (h *postServer) FindByID(ctx context.Context, req *FindByIDRequest, rsp *FindByIDResponse) error {
	return h.PostServer.FindByID(ctx, req, rsp)
}

func (h *postServer) GetPostFileByID(ctx context.Context, stream server.Stream) error {
	msg := &GetPostFileRequest{}
	if err := stream.Recv(msg); err != nil {
		return err
	}
	return h.PostServer.GetPostFileByID(ctx, msg, &postGetPostFileByIDStream{stream})
}

type postGetPostFileByIDStream struct {
	stream server.Stream
}

func (s *postGetPostFileByIDStream) Close() error {
	return s.stream.Close()
}

func (s *postGetPostFileByIDStream) Context() context.Context {
	return s.stream.Context()
}

func (s *postGetPostFileByIDStream) SendMsg(msg interface{}) error {
	return s.stream.Send(msg)
}

func (s *postGetPostFileByIDStream) RecvMsg(msg interface{}) error {
	return s.stream.Recv(msg)
}

func (s *postGetPostFileByIDStream) Send(msg *GetPostFileResponse) error {
	return s.stream.Send(msg)
}

func RegisterPostServer(s server.Server, sh PostServer, opts ...server.HandlerOption) error {
	type post interface {
		FindByID(ctx context.Context, req *FindByIDRequest, rsp *FindByIDResponse) error
		GetPostFileByID(ctx context.Context, stream server.Stream) error
	}
	type Post struct {
		post
	}
	h := &postServer{sh}
	for _, endpoint := range NewPostEndpoints() {
		opts = append(opts, api.WithEndpoint(endpoint))
	}
	return s.Handle(s.NewHandler(&Post{h}, opts...))
}
