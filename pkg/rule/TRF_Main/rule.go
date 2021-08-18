package TRF_Main

import util "github-connector/pkg/util"

import (
	"github-connector/pkg/env"
	"os"
)

type CFGContext struct {
	_context     *util.CustomContext
	_ruleConfig  map[string]interface{}
	_returnValue interface{}
	_errorValue  interface{}
}

func NewCFGContext(pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		_context:    pContext,
		_ruleConfig: pRuleConfig,
	}
}
func TRF_Main(pContext *util.CustomContext, pRuleConfig map[string]interface{}) interface{} {

	cfg := NewCFGContext(pContext, pRuleConfig)
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

	return nil
}

func (cfg *CFGContext) xiENVLNodeM12() error {

	if val, exist := os.LookupEnv(env.REST_PORT); exist {

		env.RESTPORT = val
	}
	if val, exist := os.LookupEnv(env.GRPC_PORT); exist {

		env.GRPCPORT = val
	}
	return nil
}

func (cfg *CFGContext) xiConstantNodeM01() error {
	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {

	return nil
}
