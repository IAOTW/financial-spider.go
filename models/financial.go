package models

import (
	"financial-spider.go/utils/db"
	"financial-spider.go/utils/tools"
	"strings"
)

// Financial 财务报表
type Financial struct {
	Code       string `db:"WHERE"` // 股票代码
	year       string // 年份
	ReportDate string `db:"WHERE"` // 财报季期
	reportType string // 季期类型（Q1~Q4，分别代表：一季报、半年报、三季报、年报；O，代表：其他）

	Dividend interface{} // 年度分红金额

	Ocf                  interface{} // 营业活动现金流量
	Cfi                  interface{} // 投资活动现金流量
	Cff                  interface{} // 筹资活动现金流量
	AssignDividendPorfit interface{} // 分配股利、利润或偿付利息支付的现金
	AcquisitionAssets    interface{} // 购建固定资产、无形资产和其他长期资产支付的现金
	InventoryLiquidating interface{} // 存货减少额

	Np       interface{} // 净利润
	Oi       interface{} // 营业收入
	Coe      interface{} // 营业成本
	CoeTotal interface{} // 营业总成本（含各种费用，销售费用、管理费用等）
	Eps      interface{} // 每股盈余|基本每股收益

	MonetaryFund          interface{} // 货币资金
	TradeFinassetNotfvtpl interface{} // 交易性金融资产
	TradeFinasset         interface{} // 交易性金融资产（历史遗留）
	DeriveFinasset        interface{} // 衍生金融资产

	FixedAsset interface{} // 固定资产
	Cip        interface{} // 在建工程

	CaTotal         interface{} // 流动资产总额
	NcaTotal        interface{} // 非流动资产总额
	ClTotal         interface{} // 流动负债总额
	NclTotal        interface{} // 非流动负债产总额
	Inventory       interface{} // 存货
	AccountsRece    interface{} // 应收账款
	AccountsPayable interface{} // 应付账款
}

// NewFinancial 新建财务报表对象
func NewFinancial(code string, reportDate string) *Financial {
	ymd := strings.Split(reportDate, "-")
	financial := Financial{Code: code, year: ymd[0], ReportDate: reportDate}
	switch ymd[1] {
	case "03":
		financial.reportType = "Q1"
	case "06":
		financial.reportType = "H1"
	case "09":
		financial.reportType = "Q3"
	case "12":
		financial.reportType = "FY"
	default:
		financial.reportType = "O"
	}
	return &financial
}

// InitData 初始化数据库数据
func (financial *Financial) InitData() {
	sql := "SELECT COUNT(*) AS cnt FROM financial WHERE code = ? AND year = ? AND report_date = ?"
	args := []interface{}{financial.Code, financial.year, financial.ReportDate}
	result := db.ExecSQL(sql, args...)

	if result[0]["cnt"].(int64) == 0 {
		sql = "INSERT INTO financial(code, year, report_date, report_type) VALUES(?, ?, ?, ?)"
		args = append(args, financial.reportType)
		db.ExecSQL(sql, args...)
	}
}

// UpdateData 更新数据库数据
func (financial *Financial) UpdateData() {
	sql, args := tools.MakeUpdateSqlAndArgs(financial)
	if sql == "" {
		return
	}
	db.ExecSQL(sql, args...)
}
