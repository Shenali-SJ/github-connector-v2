package ARF_Commit

import util "github-connector/pkg/util"

import (
	"github-connector/pkg/functions"
	"github-connector/pkg/model"
)

type CFGContext struct {
	Requestdata      *model.Requestdata
	currentCommit    []string
	blobSHA          string
	newTreeSHA       string
	updateHeadReturn string
	gitToken         string
	commitURL        string
	owner            string
	repo             string
	branch           string
	newCommitSHA     string
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
func ARF_Commit(pRequestdata *model.Requestdata, pContext *util.CustomContext, pRuleConfig map[string]interface{}) interface{} {

	cfg := NewCFGContext(pRequestdata, pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	cfg.xiModelPropertyNodeM02()

	cfg.xiHybridFunctionNodeM21()

	cfg.xiHybridFunctionNodeM13()

	cfg.xiHybridFunctionNodeM34()

	cfg.xiHybridFunctionNodeM45()

	cfg.xiHybridFunctionNodeM56()

	cfg.xiHybridFunctionNodeM67()
	return nil
}

func (cfg *CFGContext) xiHybridFunctionNodeM67() error {
	var err error
	cfg.updateHeadReturn, err = functions.UpdateHead(cfg.owner, cfg.repo, cfg.branch, cfg.gitToken, cfg.newCommitSHA)
	cfg._returnValue = cfg.updateHeadReturn
	return err
}

func (cfg *CFGContext) xiHybridFunctionNodeM56() error {
	var err error
	cfg.newCommitSHA, err = functions.CreateCommit(cfg.owner, cfg.repo, cfg.gitToken, cfg.currentCommit[0], cfg.newTreeSHA)
	cfg._returnValue = cfg.newCommitSHA
	return err
}

func (cfg *CFGContext) xiHybridFunctionNodeM45() error {
	var err error
	cfg.newTreeSHA, err = functions.CreateTree(cfg.owner, cfg.repo, cfg.gitToken, cfg.currentCommit[1], cfg.blobSHA)
	cfg._returnValue = cfg.newTreeSHA
	return err
}

func (cfg *CFGContext) xiHybridFunctionNodeM34() error {
	var err error
	cfg.blobSHA, err = functions.CreateBlob(cfg.owner, cfg.repo, cfg.gitToken)
	cfg._returnValue = cfg.blobSHA
	return err
}

func (cfg *CFGContext) xiHybridFunctionNodeM13() error {
	var err error
	cfg.currentCommit, err = functions.GetCurrentCommit(cfg.commitURL, cfg.gitToken)
	cfg._returnValue = cfg.currentCommit
	return err
}

func (cfg *CFGContext) xiHybridFunctionNodeM21() error {
	var err error
	cfg.commitURL, err = functions.GetHeadRef(cfg.owner, cfg.repo, cfg.branch, cfg.gitToken)
	cfg._returnValue = cfg.commitURL
	return err
}

func (cfg *CFGContext) xiModelPropertyNodeM02() error {

	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
	cfg.owner = cfg.Requestdata.Gitrepoowner
	cfg.repo = cfg.Requestdata.Repotocommit
	cfg.branch = cfg.Requestdata.Branchtocommit
	cfg.gitToken = cfg.Requestdata.Token

	return nil
}
