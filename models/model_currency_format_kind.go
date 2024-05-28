package models
// CurrencyFormatKind : Currency format kind
type CurrencyFormatKind string

// List of CurrencyFormatKind
const (
	CURRENCYFORMATKIND_NONE CurrencyFormatKind = "None"
	CURRENCYFORMATKIND_DECIMAL CurrencyFormatKind = "Decimal"
	CURRENCYFORMATKIND_CURRENCY CurrencyFormatKind = "Currency"
	CURRENCYFORMATKIND_PERCENT CurrencyFormatKind = "Percent"
	CURRENCYFORMATKIND_SCIENTIFIC CurrencyFormatKind = "Scientific"
	CURRENCYFORMATKIND_SPELL_OUT CurrencyFormatKind = "SpellOut"
	CURRENCYFORMATKIND_ORDINAL CurrencyFormatKind = "Ordinal"
	CURRENCYFORMATKIND_CURRENCY_ISO_CODE CurrencyFormatKind = "CurrencyISOCode"
	CURRENCYFORMATKIND_CURRENCY_PLURAL CurrencyFormatKind = "CurrencyPlural"
	CURRENCYFORMATKIND_CURRENCY_ACCOUNTING CurrencyFormatKind = "CurrencyAccounting"
)
