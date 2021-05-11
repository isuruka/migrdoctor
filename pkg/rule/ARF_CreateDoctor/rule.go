package ARF_CreateDoctor
import (
    "migrdoctor/pkg/env"
    "migrdoctor/pkg/model"
)

type CFGContext struct {
    Doctor             *model.Doctor
    _xiDBNode_UserName interface{}
    _xiDBNode_Password interface{}
    vp2                *model.Doctor
    _returnValue       interface{}
    _errorValue        interface{}
}

func NewCFGContext(pDoctor *model.Doctor) *CFGContext {
    return &CFGContext{

        Doctor: pDoctor,
    }
}
func ARF_CreateDoctor(pDoctor *model.Doctor) interface{} {

    cfg := NewCFGContext(pDoctor)
    cfg.r0()
    return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

    cfg.xiModelPropertyNodeM0()

    if err :=
        cfg.xiDBNodeM10(); err != nil {
        return err
    }
    return nil
}

func (cfg *CFGContext) xiDBNodeM10() error {

    db := env.RDB
    db.Create(&cfg.vp2)
    cfg._returnValue = cfg.vp2
    return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
    cfg.vp2 = cfg.Doctor

    return nil
}
