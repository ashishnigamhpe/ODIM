// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: account.proto

package account

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Account service

type AccountService interface {
	Create(ctx context.Context, in *CreateAccountRequest, opts ...client.CallOption) (*AccountResponse, error)
	GetAllAccounts(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...client.CallOption) (*AccountResponse, error)
	GetAccountServices(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error)
	Update(ctx context.Context, in *UpdateAccountRequest, opts ...client.CallOption) (*AccountResponse, error)
	Delete(ctx context.Context, in *DeleteAccountRequest, opts ...client.CallOption) (*AccountResponse, error)
}

type accountService struct {
	c    client.Client
	name string
}

func NewAccountService(name string, c client.Client) AccountService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "account"
	}
	return &accountService{
		c:    c,
		name: name,
	}
}

func (c *accountService) Create(ctx context.Context, in *CreateAccountRequest, opts ...client.CallOption) (*AccountResponse, error) {
	req := c.c.NewRequest(c.name, "Account.Create", in)
	out := new(AccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) GetAllAccounts(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error) {
	req := c.c.NewRequest(c.name, "Account.GetAllAccounts", in)
	out := new(AccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...client.CallOption) (*AccountResponse, error) {
	req := c.c.NewRequest(c.name, "Account.GetAccount", in)
	out := new(AccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) GetAccountServices(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error) {
	req := c.c.NewRequest(c.name, "Account.GetAccountServices", in)
	out := new(AccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Update(ctx context.Context, in *UpdateAccountRequest, opts ...client.CallOption) (*AccountResponse, error) {
	req := c.c.NewRequest(c.name, "Account.Update", in)
	out := new(AccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Delete(ctx context.Context, in *DeleteAccountRequest, opts ...client.CallOption) (*AccountResponse, error) {
	req := c.c.NewRequest(c.name, "Account.Delete", in)
	out := new(AccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountHandler interface {
	Create(context.Context, *CreateAccountRequest, *AccountResponse) error
	GetAllAccounts(context.Context, *AccountRequest, *AccountResponse) error
	GetAccount(context.Context, *GetAccountRequest, *AccountResponse) error
	GetAccountServices(context.Context, *AccountRequest, *AccountResponse) error
	Update(context.Context, *UpdateAccountRequest, *AccountResponse) error
	Delete(context.Context, *DeleteAccountRequest, *AccountResponse) error
}

func RegisterAccountHandler(s server.Server, hdlr AccountHandler, opts ...server.HandlerOption) error {
	type account interface {
		Create(ctx context.Context, in *CreateAccountRequest, out *AccountResponse) error
		GetAllAccounts(ctx context.Context, in *AccountRequest, out *AccountResponse) error
		GetAccount(ctx context.Context, in *GetAccountRequest, out *AccountResponse) error
		GetAccountServices(ctx context.Context, in *AccountRequest, out *AccountResponse) error
		Update(ctx context.Context, in *UpdateAccountRequest, out *AccountResponse) error
		Delete(ctx context.Context, in *DeleteAccountRequest, out *AccountResponse) error
	}
	type Account struct {
		account
	}
	h := &accountHandler{hdlr}
	return s.Handle(s.NewHandler(&Account{h}, opts...))
}

type accountHandler struct {
	AccountHandler
}

func (h *accountHandler) Create(ctx context.Context, in *CreateAccountRequest, out *AccountResponse) error {
	return h.AccountHandler.Create(ctx, in, out)
}

func (h *accountHandler) GetAllAccounts(ctx context.Context, in *AccountRequest, out *AccountResponse) error {
	return h.AccountHandler.GetAllAccounts(ctx, in, out)
}

func (h *accountHandler) GetAccount(ctx context.Context, in *GetAccountRequest, out *AccountResponse) error {
	return h.AccountHandler.GetAccount(ctx, in, out)
}

func (h *accountHandler) GetAccountServices(ctx context.Context, in *AccountRequest, out *AccountResponse) error {
	return h.AccountHandler.GetAccountServices(ctx, in, out)
}

func (h *accountHandler) Update(ctx context.Context, in *UpdateAccountRequest, out *AccountResponse) error {
	return h.AccountHandler.Update(ctx, in, out)
}

func (h *accountHandler) Delete(ctx context.Context, in *DeleteAccountRequest, out *AccountResponse) error {
	return h.AccountHandler.Delete(ctx, in, out)
}
