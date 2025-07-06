package go_starpago

const (
	PKPayInMethodWALLET = "PK_WALLET"
	PKPayInMethodDF     = "PK_DF"

	PKWalletPayTypeJAZZCASH  = "PK_WALLET_JAZZCASH"
	PKWalletPayTypeEASYPAISA = "PK_WALLET_EASYPAISA"
)

var PKPayInMethodMap = map[string]string{
	PKPayInMethodWALLET: "巴基斯坦代收", //代收只能是钱包, 无非是哪一种钱包,这个是下边map里二选一
	PKPayInMethodDF:     "巴基斯坦代付",
}

var PKWalletPayTypeMap = map[string]string{
	PKWalletPayTypeJAZZCASH:  "JazzCash（选择这种方式 extra.accountNo 必填）",
	PKWalletPayTypeEASYPAISA: "EasyPaisa",
}

// code -> 名称
// https://apifox.com/apidoc/shared/41f14b20-c5e5-44b8-b681-a05aae31f999/doc-6224970
var PKBankMap = map[string]string{
	"ALBARAKA":     "Al Baraka Islamic Bank",
	"ABL":          "Allied Bank Limited",
	"APNA":         "Apna Microfinance Bank",
	"AKBL":         "Askari Bank Limited",
	"BAHL":         "Bank Al Habib Limited",
	"BAFL":         "Bank Alfalah Limited",
	"BIPL":         "Bank Islami Pakistan Limited",
	"BOK":          "Bank of Khyber",
	"CITI":         "Citi Bank N.A.",
	"DIBPL":        "Dubai Islamic Bank Pakistan Limited",
	"FBL":          "Faysal Bank Limited",
	"FINCA":        "FINCA Microfinance Bank",
	"FWBL":         "First Women Bank Limited",
	"HBL":          "Habib Bank Limited",
	"HMB":          "Habib Metropolitan Bank Limited",
	"JSBL":         "JS Bank Limited",
	"MCB":          "MCB Bank Limited",
	"MCB_ISLAMIC":  "MCB Islamic Bank",
	"MEBL":         "Meezan Bank",
	"NBP":          "National Bank of Pakistan",
	"NRSP_MF_BANK": "NRSP Microfinance Bank",
	"NIB":          "NIB Bank Limited",
	"SBL":          "Samba Bank Limited",
	"SILK":         "Silk Bank Limited",
	"SINDH":        "Sindh Bank Limited",
	"SNBL":         "Soneri Bank Limited",
	"SCB":          "Standard Chartered Bank Ltd",
	"SMBL":         "Summit Bank Limited",
	"TMB":          "Telenor Microfinance Bank",
	"BOP":          "The Bank of Punjab",
	"U_BANK":       "U Microfinance Bank",
	"UBL":          "United Bank Limited",
	"JAZZCASH":     "JazzCash",
	"EASYPAISA":    "Easypaisa",
}

//-------------------菲律宾------------------------------

// 它是二合一了.   而巴基斯坦没有
const (
	PHPayInMethodGCASH = "PH_GCASH"
	PHPayInMethodMAYA  = "PH_MAYA"
	PHPayInMethodGRAB  = "PH_GRAB"
	PHPayInMethodQRIS  = "PH_QRIS"

	PHPayOutMethodWallet = "PH_DF_WALLET"
	PHPayOutMethodBANK   = "PH_DF_BANK"
)

var PHPayInMethodMap = map[string]string{
	PHPayInMethodGCASH: "菲律宾代收-gcash",
	PHPayInMethodMAYA:  "菲律宾代收-maya",
	PHPayInMethodGRAB:  "菲律宾代收-grab",
	PHPayInMethodQRIS:  "菲律宾代收-qrph",
}

var PHPayOutMethodMap = map[string]string{
	PHPayOutMethodWallet: "菲律宾-代付-钱包",
	PHPayOutMethodBANK:   "菲律宾-代付-银行卡",
}

var PHBankMap = map[string]string{
	"GCASH":         "gcash",
	"MAYA":          "maya",
	"OMNIPAY":       "omnipay",
	"GRABPAY":       "grabpay",
	"COINS":         "coins.ph",
	"UBPH":          "Union Bank of the Philippines",
	"BDO":           "BDO Unibank,Inc.",
	"MBTC":          "Metropolitan Bank and Trust Company",
	"BPIF":          "Bank of the Philippine Islands / BPI Family",
	"LBPH":          "Land Bank of the Philippines",
	"SBC":           "Security Bank Corporation",
	"CBC":           "China Banking Corporation",
	"DBPH":          "Development Bank of the Philippines",
	"RCC":           "Rizal Commercial Banking Corporation",
	"PNB":           "Philippine National Bank",
	"APH":           "Alipay Philippines",
	"ABTB":          "AllBank (A Thrift Bank), Inc.",
	"AUBC":          "Asia United Bank Corporation",
	"BFS":           "Bananapay Fintech Services",
	"BMAB":          "Bangko Mabuhay (A Rural Bank), Inc.",
	"BOCHINA":       "Bank of China",
	"BOCOMMERCE":    "Bank of Commerce",
	"BBI":           "Bayanihan Bank Inc.",
	"BDB":           "BDO Network Bank",
	"BPIDB":         "BPI Direct BanKo",
	"CBRB":          "Camalig Bank, Inc. (A Rural Bank)",
	"CANTILAN_BANK": "Cantilan Bank Inc.",
	"CARD_BANK":     "CARD BANK Inc.",
	"CSMB":          "CARD SME Bank,Inc.A Thrift Bank",
	"CLRB":          "Cebuana Lhuillier Rural Bank,Inc",
	"CBPH":          "China Bank Savings,Inc.",
	"CIMB":          "CIMB Philippines,Inc.",
	"CIS":           "CIS Bayad Center,Inc.",
	"CRBRB":         "Community Rural Bank of Romblon,Inc.",
	"CTBC":          "CTBC Bank (Philippines) Corporation",
	"DCDPH":         "Dumaguete City Development Bank",
	"DGB":           "Dungganon Bank (A Microfinance Rural Bank),Inc.",
	"EWC":           "East West Banking Corporation",
	"EPGEMI":        "Easy Pay Global EMI Corp",
	"ESB":           "Equicom Savings Bank,Inc.",
	"GTB":           "GoTyme Bank Corporation",
	"HKBSPH":        "The Hong Kong and Shanghai Banking Corporation Limited, Philippine Branch",
	"INFOSERVE":     "INFOSERVE INCORPORATED",
	"ING":           "ING Bank N.V.",
	"IREMIT":        "I-Remit Inc.",
	"ISLAB":         "ISLA Bank (A Thrift Bank),Inc.",
	"LSB":           "Legazpi Savings Bank",
	"LFS":           "Lulu Financial Services",
	"LDB":           "Luzon Development Bank",
	"MBSM":          "Malayan Bank Savings and Mortgage Bank,Inc.",
	"MAYABANK":      "MAYA BANK,INC",
	"MAYBANK_PH":    "Maybank Philippines,Inc.",
	"MCB":           "Mindanao Consolidated Cooperative Bank",
	"OWNB":          "OWN BANK THE RURAL BANK OF CAVITE CITY INC.",
	"PASB":          "Pacific Ace Savings Bank,Inc.",
	"PRBRB":         "Partner Rural Bank (Cotabato),Inc.",
	"PBC":           "Philippine Bank of Communications",
	"PBB":           "Philippine Business Bank",
	"PDAE":          "Philippine Digital Asset Exchange",
	"PSB":           "Philippine Savings Bank",
	"PVB":           "Philippine Veterans Bank",
	"PHTB":          "PhilTrust Bank",
	"PPSPEPP":       "PPS-PEPP Financial Services Corporation",
	"PSBC":          "Producers Savings Bank Corporation",
	"QCBDB":         "Queen City Development Bank,Inc.",
	"QCRB":          "Quezon Capital Rural Bank,Inc.",
	"RANG_AY_BANK":  "RANG-AY BANK A Rural Bank Inc",
	"RBGP":          "Robinsons Bank Corporation",
	"RBRB":          "Rural Bank of Guinobatan,Inc.",
	"SBP":           "Seabank Philippines,Inc.",
	"SBC2":          "Security Bank Corporation 2",
	"ShopeePay":     "ShopeePay Philippines,Inc.",
	"SPEEDYPAY":     "SpeedyPay Inc.",
	"SCB":           "Standard Chartered Bank",
	"STARP":         "Starpay Corporation",
	"SBAA":          "Sterling Bank of Asia, Inc.,A Savings Bank",
	"SSB":           "Sun Savings Bank,Inc.",
	"TCI":           "Tayocash Inc.",
	"TTWI":          "TokTok Wallet Inc.",
	"TDB":           "Tonik Digital Bank",
	"TJT":           "Topjuan Tech Corporation",
	"TP":            "Traxion Pay Inc.",
	"UCPBSB":        "UCPB Savings Bank,Inc.",
	"UBSB":          "Union Digital Bank",
	"UCBPH":         "United Coconut Planters Bank",
	"UNOB":          "UnoBank Inc.",
	"USSCMONEY":     "USSC Money Services, Inc",
	"WDB":           "Wealth Development Bank Corporation",
	"WPI":           "Wise Pilipinas, Inc.",
	"ZT":            "Zybi Tech Inc.",
}
