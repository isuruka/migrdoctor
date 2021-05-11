package TRF_Main

import (
	"migrdoctor/pkg/env"
	"os"
)

type CFGContext struct {
	_returnValue interface{}
	_errorValue  interface{}
}

func NewCFGContext() *CFGContext {
	return &CFGContext{}
}
func TRF_Main() interface{} {

	cfg := NewCFGContext()
	cfg.r0()
	return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	cfg.xiConstantNodeM01()

	cfg.xiENVLNodeM12()

	cfg.xiCPNodeM23()
	return nil
}

func (cfg *CFGContext) xiCPNodeM23() error {

	env.RDBHOST = env.DEFAULTHOST
	env.RDBPORT = env.DEFAULTPORT
	env.RDBDBNAME = env.DEFAULTDB
	env.RDBUSER = env.DEFAULTUSER
	env.RDBPASSWORD = env.DEFAULTPWD
	return nil
}

func (cfg *CFGContext) xiENVLNodeM12() error {

	if val, exist := os.LookupEnv(env.DEFAULT_DIALET); exist {

		env.DEFAULTDIALET = val
	}
	if val, exist := os.LookupEnv(env.DEFAULT_DB); exist {

		env.DEFAULTDB = val
	}
	if val, exist := os.LookupEnv(env.DEFAULT_HOST); exist {

		env.DEFAULTHOST = val
	}
	if val, exist := os.LookupEnv(env.DEFAULT_PORT); exist {

		env.DEFAULTPORT = val
	}
	if val, exist := os.LookupEnv(env.DEFAULT_USER); exist {

		env.DEFAULTUSER = val
	}
	if val, exist := os.LookupEnv(env.DEFAULT_PWD); exist {

		env.DEFAULTPWD = val
	}
	if val, exist := os.LookupEnv(env.DEFAULT_URL); exist {

		env.DEFAULTURL = val
	}
	if val, exist := os.LookupEnv(env.REST_PORT); exist {

		env.REST_PORT = val
	}
	if val, exist := os.LookupEnv(env.GRPC_PORT); exist {

		env.GRPC_PORT = val
	}
	return nil
}

func (cfg *CFGContext) xiConstantNodeM01() error {
	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {

	return nil
}
