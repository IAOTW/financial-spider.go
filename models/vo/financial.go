package vo

type FinancialData struct {
	ReportDate string `json:"REPORT_DATE"` // 报告期：yyyy-MM-dd HH:mm:ss

	Ocf                  interface{} `json:"NETCASH_OPERATE"`        // 经营活动产生的现金流量净额
	Cfi                  interface{} `json:"NETCASH_INVEST"`         // 投资活动产生的现金流量净额
	Cff                  interface{} `json:"NETCASH_FINANCE"`        // 筹资活动产生的现金流量净额
	AssignDividendPorfit interface{} `json:"ASSIGN_DIVIDEND_PORFIT"` // 分配股利、利润或偿付利息支付的现金
	AcquisitionAssets    interface{} `json:"CONSTRUCT_LONG_ASSET"`   // 购建固定资产、无形资产和其他长期资产支付的现金
	InventoryLiquidating interface{} `json:"INVENTORY_REDUCE"`       // 存货减少额

	Np       interface{} `json:"NETPROFIT"`            // 净利润
	Oi       interface{} `json:"TOTAL_OPERATE_INCOME"` // 营业收入
	Coe      interface{} `json:"OPERATE_COST"`         // 营业成本
	CoeTotal interface{} `json:"TOTAL_OPERATE_COST"`   // 营业总成本（含各种费用，销售费用、管理费用等）
	Eps      interface{} `json:"BASIC_EPS"`            // 每股盈余|基本每股收益

	MonetaryFund          interface{} `json:"MONETARYFUNDS"`           // 货币资金
	TradeFinassetNotfvtpl interface{} `json:"TRADE_FINASSET_NOTFVTPL"` // 交易性金融资产
	TradeFinasset         interface{} `json:"TRADE_FINASSET"`          // 交易性金融资产（历史遗留）
	DeriveFinasset        interface{} `json:"DERIVE_FINASSET"`         // 衍生金融资产

	FixedAsset interface{} `json:"FIXED_ASSET"` // 固定资产
	Cip        interface{} `json:"CIP"`         // 在建工程

	CaTotal         interface{} `json:"TOTAL_CURRENT_ASSETS"`    // 流动资产总额
	NcaTotal        interface{} `json:"TOTAL_NONCURRENT_ASSETS"` // 非流动资产总额
	ClTotal         interface{} `json:"TOTAL_CURRENT_LIAB"`      // 流动负债总额
	NclTotal        interface{} `json:"TOTAL_NONCURRENT_LIAB"`   // 非流动负债产总额
	Inventory       interface{} `json:"INVENTORY"`               // 存货
	AccountsRece    interface{} `json:"ACCOUNTS_RECE"`           // 应收账款
	AccountsPayable interface{} `json:"ACCOUNTS_PAYABLE"`        // 应付账款
}

// FinancialResult 现金流量表
type FinancialResult struct {
	Count int             `json:"count"`
	Pages int             `json:"pages"`
	Data  []FinancialData `json:"data"` // 长度等于传入日期数，这里是一个
	// 错误信息
	Type   string `json:"$type"`  // 1
	Status int    `json:"status"` // -1
}

// -----

type DividendData struct {
	Year  string      `json:"STATISTICS_YEAR"`
	Money interface{} `json:"TOTAL_DIVIDEND"`
}

type DividendPage struct {
	Count int            `json:"count"`
	Pages int            `json:"pages"`
	Data  []DividendData `json:"data"` // 只有一个
}

// DividendResult 分红信息
type DividendResult struct {
	Code    int          `json:"code"`
	Success bool         `json:"success"`
	Message string       `json:"msg"`
	Result  DividendPage `json:"result"`
}
