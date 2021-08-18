package ARF_Sidecar

import util "github-connector/pkg/util"

import (
	"github-connector/pkg/dto"
	"github-connector/pkg/functions"
	"github-connector/pkg/model"
	"github-connector/pkg/rule/ARF_CreateRepoSubRule"
	"github-connector/pkg/rule/ARF_DeleteRepoSubRule"
	"github-connector/pkg/xiLogger"
)

type CFGContext struct {
	Request               *dto.Request
	operation             string
	request               string
	requestJson           *model.Requestdata
	createRepoReturnValue string
	deleteRepoReturnValue string
	_context              *util.CustomContext
	_ruleConfig           map[string]interface{}
	_returnValue          interface{}
	_errorValue           interface{}
}

func NewCFGContext(pRequest *dto.Request, pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		Request:     pRequest,
		_context:    pContext,
		_ruleConfig: pRuleConfig,
	}
}
func ARF_Sidecar(pRequest *dto.Request, pContext *util.CustomContext, pRuleConfig map[string]interface{}) interface{} {

	cfg := NewCFGContext(pRequest, pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	if err := cfg.xiHybridFunctionNodeM01(); err != nil {
		xiLogger.Log.Info("Could not convert request string to JSON")
	}

	cfg.xiModelPropertyNodeM12()

	cfg.xiConstantNodeM23()

	cfg.xiSwitchNodeM34()
	return nil
}

func (cfg *CFGContext) xiSwitchNodeM34() error {

	switch _input := cfg.operation; {

	case _input == createRepoOperation:
		cfg.branchM41()
	case _input == deleteRepoOperation:
		cfg.branchM42()

	}
	return nil
}

const createRepoOperation = "createRepo"
const deleteRepoOperation = "deleteRepo"

func (cfg *CFGContext) xiConstantNodeM23() error {
	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM12() error {
	cfg.operation = cfg.requestJson.Operation

	return nil
}

func (cfg *CFGContext) xiHybridFunctionNodeM01() error {
	var err error
	cfg.requestJson, err = functions.ConvertStringToJson(cfg.request)
	cfg._returnValue = cfg.requestJson
	return err
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
	cfg.request = cfg.Request.Data
	cfg.requestJson = &model.Requestdata{}

	return nil
}
func (cfg *CFGContext) branchM41() error {

	cfg.xiModelPropertyNodeM5()

	cfg.xiSubRuleNodeM56()
	return nil
}

func (cfg *CFGContext) xiSubRuleNodeM56() error {
	cfg.createRepoReturnValue = ARF_CreateRepoSubRule.ARF_CreateRepoSubRule(cfg.requestJson, cfg._context, cfg._ruleConfig).(string)
	cfg._returnValue = cfg.createRepoReturnValue
	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM5() error {

	return nil
}
func (cfg *CFGContext) branchM42() error {

	cfg.xiModelPropertyNodeM7()

	cfg.xiSubRuleNodeM78()
	return nil
}

func (cfg *CFGContext) xiSubRuleNodeM78() error {
	cfg.deleteRepoReturnValue = ARF_DeleteRepoSubRule.ARF_DeleteRepoSubRule(cfg.requestJson, cfg._context, cfg._ruleConfig).(string)
	cfg._returnValue = cfg.deleteRepoReturnValue
	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM7() error {

	return nil
}
