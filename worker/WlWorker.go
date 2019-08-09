package worker

import (
	"fmt"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/repository"
)

type wlWorker struct {
}

func NewWlWorker() *wlWorker {
	return &wlWorker{}
}

//配送出库&配送出库二级明细表
func (w *wlWorker) Z3PsCkHpDt() {
	id := goToolCommon.Guid()
	log.Debug(fmt.Sprintf("Z3PsCkHpDt %s start", id))
	defer log.Debug(fmt.Sprintf("Z3PsCkHpDt %s complete", id))
	repWl := repository.NewRepWl()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3PsCkHpDt get rep online err: %s", err.Error())
		log.Error(errMsg)
		return
	}
	uCounter := 0
	udCounter := 0
	for {
		rList, err := repWl.GetZ3PsCkHpDtSy()
		if err != nil {
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3PsCkHpDt get sy error: return list can not be nil")
			log.Error(errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repWl.GetZ3PsCkHpDt(id)
			if err != nil {
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3PsCkHpDt get data error: return list can not be nil")
				log.Error(errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3PsCkHpDt(d)
					if err != nil {
						return
					}
				}
				uCounter = uCounter + 1
			}
			ddList, err := repWl.GetZ3PsCkHpFjDt(id)
			if ddList == nil {
				errMsg := fmt.Sprintf("Z3PsCkHpFjDt get data error: return list can not be nil")
				log.Error(errMsg)
				return
			}
			if len(ddList) > 0 {
				for _, d := range ddList {
					err := repOnLine.UpdateZ3PsCkHpFjDt(d)
					if err != nil {
						return
					}
				}
				udCounter = udCounter + 1
			}
			err = repWl.DelZ3PsCkHpDtSy(id)
			if err != nil {
				return
			}
			log.Debug(fmt.Sprintf("gs Z3PsCkHpDt[%s] Update", id))
		}
	}
	if uCounter > 0 || udCounter > 0 {
		log.Info(fmt.Sprintf("gs Z3PsCkHpDt Update %d,Z3PsCkHpFjDt Update %d", uCounter, udCounter))
	}
}

//工厂配送修正 z3psxzt 和z3psxzdt
func (w *wlWorker) Z3PsXzt() {
	log.Debug("Z3PsXzt")
}

//工厂配送冲红z3mdpscht 和z3mdpschdt
func (w *wlWorker) Z3MdPsCht() {
	log.Debug("Z3MdPsCht")
}

//工厂赊账销售出库z3shezhxsckt和z3shezhxshpmxt
func (w *wlWorker) Z3SheZhXsCkt() {
	log.Debug("Z3SheZhXsCkt")
}

//工厂完工入库shengchwgrkdt 和shengchwgrkt
func (w *wlWorker) ShengChWgRkt() {
	log.Debug("ShengChWgRkt")
}

//工厂即时台账表xttz
func (w *wlWorker) XtTz() {
	log.Debug("XtTz")
}

//工厂代门店订货单z3mddhdt
func (w *wlWorker) Z3MdDhDt() {
	log.Debug("Z3MdDhDt")
}

//工厂订货单修改z3mddhdt_md和z3mddht_md
func (w *wlWorker) Z3MdDhDtMd() {
	log.Debug("Z3MdDhDtMd")
}

//工厂订单提货单oodxv1ddshckt
func (w *wlWorker) OodXv1DdShCkt() {
	log.Debug("OodXv1DdShCkt")
}
