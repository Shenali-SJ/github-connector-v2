package ARF_CreateRepoSubRule

import util "github-connector/pkg/util"

import (
	"github-connector/pkg/functions"
	"github-connector/pkg/model"
)

type CFGContext struct {
	Requestdata      *model.Requestdata
	gitToken         string
	ruleReturnString string
	_context         *util.CustomContext
	_ruleConfig      map[string]interface{}
	_returnValue     interface{}
	_errorValue      interface{}
}

func NewCFGContext(pRequestdata *model.Requestdata, pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		Requestdata: pRequestdata,
		_context:    pContext,
		_ruleConfig: pRuleConfig,
	}
}
func ARF_CreateRepoSubRule(pRequestdata *model.Requestdata, pContext *util.CustomContext, pRuleConfig map[string]interface{}) interface{} {

	cfg := NewCFGContext(pRequestdata, pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	cfg.xiHybridFunctionNodeM01()
	return nil
}

func (cfg *CFGContext) xiHybridFunctionNodeM01() error {
	var err error
	cfg.ruleReturnString, err = functions.CreateRepo(cfg.gitToken)
	cfg._returnValue = cfg.ruleReturnString
	return err
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
	cfg.gitToken = cfg.Requestdata.Token

	return nil
}
