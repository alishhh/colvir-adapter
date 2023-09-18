package schema

import (
	"encoding/xml"
)

//ISO код валюты
type CurrencyCode string

//Тип клиента
//Возможные значения: entrepreneur_joint, entrepreneur, individual, all, corporate, transit
type CustomerType string

// DocCreationWay ...
type DocCreationWay string

//Строка, которой запрещено быть пустой
type NotEmptyString string

//enum - urgent,normal
type ProcessingMethod string

//AlterCode - альтернативный код
type AlterCode string

//FieldMandatory - enum с возможными значениями: mandatory, optional, denied, undefined
type FieldMandatory string

//CurrencyRatesSource - enum с возможными значениями: MCRATE, MEXCRAT, E_RATE
type CurrencyRatesSource string

//RoundingMethodType - enum с возможными значениями: "DOWNWARDS","ARITHMETICALLY","UPWARDS"
type RoundingMethodType string

//AccountNumber - длина 5
type AccountNumber string

//AccountIban - длина 16-34
type AccountIban string

// AboutRequest ...
/*type AboutRequest struct {
	XMLName xml.Name `xml:"aboutRequest"`
	*AboutRequestType
}

// AboutResponse ...
type AboutResponse struct {
	XMLName xml.Name `xml:"aboutResponse"`
	*AboutResponseType
}*/

// EchoRequest ...
type EchoRequest struct {
	XMLName xml.Name `xml:"echoRequest"`
	*EchoRequestType
}

// EchoResponse ...
type EchoResponse struct {
	XMLName xml.Name `xml:"echoResponse"`
	*EchoResponseType
}

// DomainCodeRequestElem ...
type DomainCodeRequestElem struct {
	DomainCode string `xml:"domainCode"`
	*AbstractRequest
}

// DomainValuesLoadResultElem ...
type DomainValuesLoadResultElem struct {
	Value []*DomainValue `xml:"value"`
	*AbstractResponse
}

// DomainAnyCodeRequestElem ...
type DomainAnyCodeRequestElem struct {
	DomainName  string `xml:"domainName"`
	DomainField string `xml:"domainField"`
	*AbstractRequest
}

// DomainAnyValuesLoadResultElem ...
type DomainAnyValuesLoadResultElem struct {
	Value []*DomainValue `xml:"value"`
	*AbstractResponse
}

// LoadCurrencyRatesElem ...
type LoadCurrencyRatesElem struct {
	XMLName            xml.Name             `xml:"loadCurrencyRatesElem"`
	CurrencyRateSource *CurrencyRatesSource `xml:"currencyRateSource"`
	ActualDate         string               `xml:"actualDate"`
	RateQueryMCRATE    *RateQueryMCRATE     `xml:"rateQueryMCRATE"`
	RateQueryMEXCRAT   *ConditionAnd        `xml:"rateQueryMEXCRAT"`
	RateQueryERATE     *ConditionAnd        `xml:"rateQueryE_RATE"`
	*AbstractRequest
}

// LoadCurrencyRatesResponseElem ...
type LoadCurrencyRatesResponseElem struct {
	XMLName           xml.Name       `xml:"loadCurrencyRatesResponseElem"`
	CurrencyRatesList *CurrencyRates `xml:"currencyRatesList"`
	*AbstractListResponse
}

// LoadDocumentOperationsElem ...
type LoadDocumentOperationsElem struct {
	XMLName           xml.Name `xml:"loadDocumentOperationsElem"`
	ColvirReferenceId []string `xml:"colvirReferenceId"`
	*AbstractRequest
}

// LoadDocumentOperationsResponseElem ...
type LoadDocumentOperationsResponseElem struct {
	XMLName            xml.Name            `xml:"loadDocumentOperationsResponseElem"`
	DocumentOperations *DocumentOperations `xml:"documentOperations"`
	*AbstractResponse
}

// DomainCodeWithParametersRequestElem ...
type DomainCodeWithParametersRequestElem struct {
	DomainCode string `xml:"domainCode"`
	Parameters string `xml:"parameters"`
	*AbstractRequest
}

// DomainValuesWithParametersLoadResultElem ...
type DomainValuesWithParametersLoadResultElem struct {
	Value []*DomainValue `xml:"value"`
	*AbstractResponse
}

// LoadRoleFieldsRequestElem ...
type LoadRoleFieldsRequestElem struct {
	Role string `xml:"role"`
	*AbstractRequest
}

// LoadRoleFieldsLoadResultElem ...
type LoadRoleFieldsLoadResultElem struct {
	Value []*RoleField `xml:"value"`
	*AbstractResponse
}

// LoadAdditionalParametersElem ...
type LoadAdditionalParametersElem struct {
	XMLName  xml.Name      `xml:"loadAdditionalParametersElem"`
	Document *DocumentBase `xml:"document"`
	*AbstractRequest
}

// LoadAdditionalParametersResultElem ...
type LoadAdditionalParametersResultElem struct {
	XMLName        xml.Name      `xml:"loadAdditionalParametersResultElem"`
	ParametersList []*CustomForm `xml:"parametersList"`
	*AbstractResponse
}

// SaveAdditionalParametersElem ...
type SaveAdditionalParametersElem struct {
	XMLName  xml.Name      `xml:"saveAdditionalParametersElem"`
	Document *DocumentBase `xml:"document"`
	*AbstractRequest
}

// SaveAdditionalParametersResultElem ...
type SaveAdditionalParametersResultElem struct {
	XMLName xml.Name      `xml:"saveAdditionalParametersResultElem"`
	Result  []*CustomForm `xml:"result"`
	*AbstractResponse
}

// LoadDebugInfoElem ...
type LoadDebugInfoElem struct {
	XMLName   xml.Name `xml:"loadDebugInfoElem"`
	RequestId string   `xml:"requestId"`
	*AbstractRequest
}

// LoadDebugInfoResponseElem ...
type LoadDebugInfoResponseElem struct {
	XMLName       xml.Name                   `xml:"loadDebugInfoResponseElem"`
	DebugInfoList *LoadDebugInfoSearchResult `xml:"debugInfoList"`
	*AbstractResponse
}

// GetValueConvertedElem ...
type GetValueConvertedElem struct {
	XMLName      xml.Name                     `xml:"getValueConvertedElem"`
	Round        int                          `xml:"round"`
	CodeProfile  string                       `xml:"codeProfile"`
	CurrencyItem []*CurrencyConversionReqItem `xml:"currencyItem"`
	*AbstractRequest
}

// GetValueConvertedResponseElem ...
type GetValueConvertedResponseElem struct {
	XMLName xml.Name                     `xml:"getValueConvertedResponseElem"`
	Result  []*CurrencyConversionReqItem `xml:"result"`
	*AbstractResponse
}

// LoadBankHolidaysElem ...
type LoadBankHolidaysElem struct {
	XMLName      xml.Name `xml:"loadBankHolidaysElem"`
	StartDate    string   `xml:"startDate"`
	EndDate      string   `xml:"endDate"`
	CalendarCode string   `xml:"calendarCode"`
	*AbstractRequest
}

// LoadBankHolidaysResponseElem ...
type LoadBankHolidaysResponseElem struct {
	XMLName      xml.Name       `xml:"loadBankHolidaysResponseElem"`
	CalendarCode string         `xml:"calendarCode"`
	Holiday      []*BankHoliday `xml:"holiday"`
	*AbstractResponse
}

// LoadBanksElem ...
type LoadBanksElem struct {
	XMLName    xml.Name `xml:"loadBanksElem"`
	RegionCode string   `xml:"regionCode"`
	BicMask    []string `xml:"bicMask"`
	*AbstractRequest
}

// LoadBanksResponseElem ...
type LoadBanksResponseElem struct {
	XMLName xml.Name `xml:"loadBanksResponseElem"`
	Bank    []*Bank  `xml:"bank"`
	*AbstractResponse
}

// DomainHierarchyValuesRequestElem ...
type DomainHierarchyValuesRequestElem struct {
	DomainCode string `xml:"domainCode"`
	*AbstractRequest
}

// DomainHierarchyValuesWithParametersRequestElem ...
type DomainHierarchyValuesWithParametersRequestElem struct {
	DomainCode string `xml:"domainCode"`
	Parameters string `xml:"parameters"`
	*AbstractRequest
}

// DomainHierarchyValuesLoadResult ...
type DomainHierarchyValuesLoadResult struct {
	Value []*DomainHierarchyValue `xml:"value"`
	*AbstractResponse
}

// ConversionPurposeRequestElem ...
type ConversionPurposeRequestElem struct {
	Currency      string `xml:"currency"`
	OperationType string `xml:"operationType"`
	DebitAccount  string `xml:"debitAccount"`
	CreditAccount string `xml:"creditAccount"`
	*AbstractRequest
}

// ConversionPurposeLoadResult ...
type ConversionPurposeLoadResult struct {
	Items []*DomainValue `xml:"items"`
	*AbstractResponse
}

// GetParameterRequestElem ...
type GetParameterRequestElem struct {
	ParameterCode    string                        `xml:"parameterCode"`
	NOrd             int                           `xml:"nOrd"`
	MethodParamsID   *GetParameterMethodParamsID   //`xml:"methodParamsID"`
	MethodParamsCODE *GetParameterMethodParamsCODE //`xml:"methodParamsCODE"`
	*AbstractRequest
}

// GetParameterResponseElem ...
type GetParameterResponseElem struct {
	Result *ResultRawString `xml:"result"`
	*AbstractResponse
}

// LoadDecisionTableRequest ...
type LoadDecisionTableRequest struct {
	XMLName xml.Name `xml:"loadDecisionTableRequest"`
	Name    string   `xml:"name"`
	*AbstractRequest
}

// LoadDecisionTableResponse ...
type LoadDecisionTableResponse struct {
	XMLName       xml.Name       `xml:"loadDecisionTableResponse"`
	DecisionTable *DecisionTable `xml:"decisionTable"`
	*AbstractResponse
}

// LoadClientBankRelationsRequestElem ...
type LoadClientBankRelationsRequestElem struct {
	AffiliationFl bool `xml:"affiliationFl"`
	*AbstractRequest
}

// LoadClientBankRelationsResult ...
type LoadClientBankRelationsResult struct {
	Relations []*ClientBankRelation `xml:"relations"`
	*AbstractResponse
}

// LoadProductHierarchyRequest ...
type LoadProductHierarchyRequest struct {
	XMLName xml.Name `xml:"loadProductHierarchyRequest"`
	Code    string   `xml:"code"`
	*AbstractRequest
}

// LoadProductHierarchyResponse ...
type LoadProductHierarchyResponse struct {
	XMLName xml.Name                    `xml:"loadProductHierarchyResponse"`
	Value   []*HierarchyProductTemplate `xml:"value"`
	*AbstractResponse
}

// Details ...
type Details struct {
	XMLName             xml.Name `xml:"details"`
	ShowCurrencies      bool     `xml:"showCurrencies"`
	ShowPurposes        bool     `xml:"showPurposes"`
	ShowPeriods         bool     `xml:"showPeriods"`
	ShowParameters      bool     `xml:"showParameters"`
	ShowDossierDocTypes bool     `xml:"showDossierDocTypes"`
}

// LoadProductDetailsRequest ...
type LoadProductDetailsRequest struct {
	XMLName   xml.Name `xml:"loadProductDetailsRequest"`
	Code      string   `xml:"code"`
	Parameter []string `xml:"parameter"`
	Details   *Details `xml:"details"`
	*AbstractRequest
}

// LoadProductDetailsResponse ...
type LoadProductDetailsResponse struct {
	XMLName         xml.Name         `xml:"loadProductDetailsResponse"`
	ProductTemplate *ProductTemplate `xml:"productTemplate"`
	*AbstractResponse
}

// DomainCode is type of collateral
type DomainCode string

// DomainHierarchyCode is Ensuring types
type DomainHierarchyCode string

// IdentDocumentTypeExtended is Identificator
type IdentDocumentTypeExtended struct {
	Id int `xml:"id"`
	*DomainValue
}

// IdentDocumentTypeExtLoadResult is Value list
type IdentDocumentTypeExtLoadResult struct {
	Value []*IdentDocumentTypeExtended `xml:"value"`
	*AbstractResponse
}

// ProductPeriod is Time period (count of time unit)
type ProductPeriod struct {
	Id       int    `xml:"id"`
	Name     string `xml:"name"`
	EndDate  string `xml:"endDate"`
	TimeType string `xml:"timeType"`
	TimeUnit int    `xml:"timeUnit"`
}

// ProductPeriodLoadResult is Domain hierarchy value list
type ProductPeriodLoadResult struct {
	Value []*ProductPeriod `xml:"value"`
	*AbstractResponse
}

// Rate is Percent value list
type Rate struct {
	Name            string         `xml:"name"`
	ProductId       int            `xml:"productId"`
	CalculateRuleId int            `xml:"calculateRuleId"`
	PercentList     []*RatePercent `xml:"percentList"`
}

// RatePercent is Start date
type RatePercent struct {
	Value     string `xml:"value"`
	Condition string `xml:"condition"`
	FromDate  string `xml:"fromDate"`
}

// CurrencyValue is Currency
type CurrencyValue struct {
	AlterCode *AlterCode `xml:"alterCode"`
	*DomainValue
}

// PaymentType ...
type PaymentType struct {
	CutOffTime    string `xml:"cutOffTime"`
	CutOffTimeInt int    `xml:"cutOffTimeInt"`
	*DomainValue
}

// Bank ...
type Bank struct {
	Bic          string `xml:"bic"`
	Swift        string `xml:"swift"`
	Telex        string `xml:"telex"`
	Name         string `xml:"name"`
	NameAddress1 string `xml:"nameAddress1"`
	NameAddress2 string `xml:"nameAddress2"`
	NameAddress3 string `xml:"nameAddress3"`
	NameAddress4 string `xml:"nameAddress4"`
	RegionCode   string `xml:"regionCode"`
	Region       string `xml:"region"`
	EnglishName  string `xml:"englishName"`
	City         string `xml:"city"`
	IbanBankCode string `xml:"ibanBankCode"`
}

// BankHoliday ...
type BankHoliday struct {
	Day       string `xml:"day"`
	Month     string `xml:"month"`
	Year      string `xml:"year"`
	Holiday   string `xml:"holiday"`
	IsHoliday string `xml:"isHoliday"`
	WeekDay   string `xml:"weekDay"`
}

// AccountingTransactionSymbolCode ...
type AccountingTransactionSymbolCode string

// AccountStatusCode ...
type AccountStatusCode string

// BeneficiaryBankCode ...
type BeneficiaryBankCode string

// BeneficiaryCode ...
type BeneficiaryCode string

// BininnCode ...
type BininnCode string

// BusinessObjectCode ...
type BusinessObjectCode string

// CardStatusCode ...
type CardStatusCode string

// CardTypeCode ...
type CardTypeCode string

// ContractProviderCode ...
type ContractProviderCode string

// ContractTypeCode ...
type ContractTypeCode string

// DepartmentCode ...
type DepartmentCode string

// DocumentCode ...
type DocumentCode string

// DocumentStateCode ...
type DocumentStateCode string

// DocumentTypeCode ...
type DocumentTypeCode string

// ErrorCode ...
type ErrorCode string

// FiscalCodeOld ...
type FiscalCodeOld string

// KnpCode ...
type KnpCode string

// MassTypeCode ...
type MassTypeCode string

// MoneySenderCode ...
type MoneySenderCode string

// PayerCode ...
type PayerCode string

// RegionCode ...
type RegionCode string

// TaxBasicCode ...
type TaxBasicCode string

// TaxTypeCode ...
type TaxTypeCode string

// TimeOfDay ...
type TimeOfDay string

// AmountInCurrency ...
type AmountInCurrency struct {
	Amount   float64       `xml:"amount"`
	Currency *CurrencyCode `xml:"currency"`
}

// AccountReference ...
type AccountReference struct {
	BankCode      string `xml:"bankCode"`
	AccountNumber string `xml:"accountNumber"`
	Currency      string `xml:"currency"`
}

// UniversalReferenceItem ...
type UniversalReferenceItem struct {
	ReferenceCode  string `xml:"referenceCode"`
	Element        string `xml:"element"`
	UpperReference string `xml:"upperReference"`
	Action         string `xml:"action"`
	Content        string `xml:"content"`
}

// UniversalReferenceRequestElem ...
type UniversalReferenceRequestElem struct {
	Reference *UniversalReferenceItem `xml:"reference"`
	*AbstractRequest
}

// UniversalReferenceLoadResult ...
type UniversalReferenceLoadResult struct {
	Items []*UniversalReferenceItem `xml:"items"`
	*AbstractResponse
}

// LoadCalculationMethodListRequest ...
type LoadCalculationMethodListRequest struct {
	XMLName xml.Name      `xml:"loadCalculationMethodListRequest"`
	Code    string        `xml:"code"`
	Filter  *ConditionAnd `xml:"filter"`
	*AbstractRequest
}

// MethodList ...
type MethodList struct {
	XMLName xml.Name             `xml:"methodList"`
	Method  []*CalculationMethod `xml:"method"`
}

// LoadCalculationMethodListResponse ...
type LoadCalculationMethodListResponse struct {
	XMLName    xml.Name    `xml:"loadCalculationMethodListResponse"`
	MethodList *MethodList `xml:"methodList"`
	*AbstractListResponse
}

// LoadRoundingMethodListRequest ...
type LoadRoundingMethodListRequest struct {
	XMLName xml.Name      `xml:"loadRoundingMethodListRequest"`
	Filter  *ConditionAnd `xml:"filter"`
	*AbstractRequest
}

// LoadRoundingMethodListResponse ...
type LoadRoundingMethodListResponse struct {
	XMLName    xml.Name    `xml:"loadRoundingMethodListResponse"`
	MethodList *MethodList `xml:"methodList"`
	*AbstractListResponse
}

// LoadExchangeRateTypeListRequest ...
type LoadExchangeRateTypeListRequest struct {
	XMLName xml.Name      `xml:"loadExchangeRateTypeListRequest"`
	Filter  *ConditionAnd `xml:"filter"`
	*AbstractRequest
}

// ExchangeRateTypeList ...
type ExchangeRateTypeList struct {
	XMLName          xml.Name            `xml:"exchangeRateTypeList"`
	ExchangeRateType []*ExchangeRateType `xml:"exchangeRateType"`
}

// LoadExchangeRateTypeListResponse ...
type LoadExchangeRateTypeListResponse struct {
	XMLName              xml.Name              `xml:"loadExchangeRateTypeListResponse"`
	ExchangeRateTypeList *ExchangeRateTypeList `xml:"exchangeRateTypeList"`
	*AbstractListResponse
}

// DocumentBase ...
type DocumentBase struct {
	ColvirReferenceId string            `xml:"colvirReferenceId"`
	ExternalSystemId  string            `xml:"externalSystemId"`
	ProcessingMethod  *ProcessingMethod `xml:"processingMethod"`
	Status            *DomainValue      `xml:"status"`
	DocumentType      *DomainValue      `xml:"documentType"`
	Code              string            `xml:"code"`
	Amount            float64           `xml:"amount"`
	Currency          *CurrencyCode     `xml:"currency"`
	Rate              float64           `xml:"rate"`
	DocumentDate      string            `xml:"documentDate"`
	ValueDate         string            `xml:"valueDate"`
	ImportDate        string            `xml:"importDate"`
	AcceptedDate      string            `xml:"acceptedDate"`
	CorrectionDate    string            `xml:"correctionDate"`
	RefuseReason      string            `xml:"refuseReason"`
	CustomerType      *CustomerType     `xml:"customerType"`
	Description       string            `xml:"description"`
	CustomForms       []*CustomForm     `xml:"customForms"`
	DocCreationWay    string            `xml:"docCreationWay"`
}

// CustomForm ...
type CustomForm struct {
	Code       string          `xml:"code"`
	Status     *DomainValue    `xml:"status"`
	Parameters []*AddInfoValue `xml:"parameters"`
}

// AddInfoValue ...
type AddInfoValue struct {
	Value    string `xml:"value"`
	FromDate string `xml:"fromDate"`
	*DomainValue
}

// DomainValue ...
type DomainValue struct {
	Code        string `xml:"code"`
	Name        string `xml:"name"`
	Description string `xml:"description"`
}

// DomainValueFull ...
type DomainValueFull struct {
	LongName       string `xml:"longName"`
	Constvalue     string `xml:"constvalue"`
	Mode           string `xml:"mode"`
	ShortName      string `xml:"shortName"`
	IsArchived     bool   `xml:"isArchived"`
	AdditionalInfo string `xml:"additionalInfo"`
	*DomainValue
}

// ErrorValue ...
type ErrorValue struct {
	Severity string `xml:"severity"`
	*DomainValue
}

// DomainHierarchyValue ...
type DomainHierarchyValue struct {
	Id             string `xml:"id"`
	IdHi           string `xml:"idHi"`
	IsGroup        string `xml:"isGroup"`
	Level          string `xml:"level"`
	AdditionalInfo string `xml:"additionalInfo"`
	AlterCode      string `xml:"alterCode"`
	*DomainValue
}

// DocumentSaveResult ...
type DocumentSaveResult struct {
	ExecuteResult              *DomainValue  `xml:"executeResult"`
	ColvirReferenceId          string        `xml:"colvirReferenceId"`
	Status                     *DomainValue  `xml:"status"`
	AcceptedDate               string        `xml:"acceptedDate"`
	DocumentNumber             string        `xml:"documentNumber"`
	Date                       string        `xml:"date"`
	SaveCustomParametersStatus []*CustomForm `xml:"saveCustomParametersStatus"`
	*ValidationResponse
}

// DocumentDeleteResult ...
type DocumentDeleteResult struct {
	DeleteResult      *DomainValue `xml:"deleteResult"`
	ColvirReferenceId string       `xml:"colvirReferenceId"`
	Status            *DomainValue `xml:"status"`
	*ValidationResponse
}

// DomainDecimal ...
type DomainDecimal struct {
	Value float64 `xml:"value"`
	*DomainValue
}

// CurrencyRateProfile ...
type CurrencyRateProfile struct {
	RateTypeSell string `xml:"rateTypeSell"`
	RateTypeBuy  string `xml:"rateTypeBuy"`
}

// RateQueryMCRATE ...
type RateQueryMCRATE struct {
	Round        int                           `xml:"round"`
	CodeProfile  string                        `xml:"codeProfile"`
	CurrencyItem []*CurrencyConversionBaseItem `xml:"currencyItem"`
}

// CurrencyConversionBaseItem ...
type CurrencyConversionBaseItem struct {
	CurrencyFrom *CurrencyCode `xml:"currencyFrom"`
	CurrencyTo   *CurrencyCode `xml:"currencyTo"`
}

// CurrencySeniority ...
type CurrencySeniority struct {
	Multiplier int `xml:"multiplier"`
	*DomainValue
}

// CurrencyRate ...
type CurrencyRate struct {
	CurrencyFrom       *CurrencyCode      `xml:"currencyFrom"`
	CurrencyTo         *CurrencyCode      `xml:"currencyTo"`
	Group              *DomainValue       `xml:"group"`
	PurchaseRate       float64            `xml:"purchaseRate"`
	PurchaseAmountFrom float64            `xml:"purchaseAmountFrom"`
	PurchaseAmountTo   float64            `xml:"purchaseAmountTo"`
	SaleRate           float64            `xml:"saleRate"`
	SaleAmountFrom     float64            `xml:"saleAmountFrom"`
	SaleAmountTo       float64            `xml:"saleAmountTo"`
	PurchaseMargin     float64            `xml:"purchaseMargin"`
	SaleMargin         float64            `xml:"saleMargin"`
	NumberOfUnits      float64            `xml:"numberOfUnits"`
	DateFrom           string             `xml:"dateFrom"`
	DateTo             string             `xml:"dateTo"`
	Description        string             `xml:"description"`
	Seniority          *CurrencySeniority `xml:"seniority"`
	RateType           *DomainValue       `xml:"rateType"`
	IsUsed             bool               `xml:"isUsed"`
	IsActual           bool               `xml:"isActual"`
}

// CurrencyRates ...
type CurrencyRates struct {
	CurrencyRates []*CurrencyRate `xml:"currencyRates"`
	*BaseSearchResult
}

// DocumentOperation ...
type DocumentOperation struct {
	ColvirReferenceId string                  `xml:"colvirReferenceId"`
	DocumentNumber    string                  `xml:"documentNumber"`
	OperationCount    int                     `xml:"operationCount"`
	Result            string                  `xml:"result"`
	Operations        *DocumentOperationsList `xml:"operations"`
}

// DocumentOperationsList ...
type DocumentOperationsList struct {
	Operation []*Operation `xml:"operation"`
}

// DocumentOperations ...
type DocumentOperations struct {
	Document []*DocumentOperation `xml:"document"`
}

// Operation ...
type Operation struct {
	DocumentNumber     string          `xml:"documentNumber"`
	OperationNumber    string          `xml:"operationNumber"`
	ProcessId          string          `xml:"processId"`
	Scenario           *DomainValue    `xml:"scenario"`
	DocumentState      *DomainValue    `xml:"documentState"`
	ExecuteDate        string          `xml:"executeDate"`
	OperDate           string          `xml:"operDate"`
	ScenarioOperation  *DomainValue    `xml:"scenarioOperation"`
	User               *DomainValue    `xml:"user"`
	IsOperationWasUndo bool            `xml:"isOperationWasUndo"`
	IsCancelsOperation bool            `xml:"isCancelsOperation"`
	IsOperationWait    bool            `xml:"isOperationWait"`
	HasCashOrd         int             `xml:"hasCashOrd"`
	HasPayOrd          int             `xml:"hasPayOrd"`
	CurrencyCode       *CurrencyCode   `xml:"currencyCode"`
	Description        string          `xml:"description"`
	RequiredSanction   bool            `xml:"requiredSanction"`
	SanctionName       string          `xml:"sanctionName"`
	SubOperation       []*SubOperation `xml:"subOperation"`
}

// RoleFieldsCollections ...
type RoleFieldsCollections struct {
	RequiredRolesFields    []*RoleField `xml:"requiredRolesFields"`
	NotEditableRolesFields []*RoleField `xml:"notEditableRolesFields"`
}

// RoleField ...
type RoleField struct {
	Type          *DomainValue    `xml:"type"`
	Value         string          `xml:"value"`
	Domain        string          `xml:"domain"`
	DbmsType      string          `xml:"dbmsType"`
	Length        int             `xml:"length"`
	CustomerType  *CustomerType   `xml:"customerType"`
	IsRequired    bool            `xml:"isRequired"`
	IsDisabled    bool            `xml:"isDisabled"`
	IsNotEditable bool            `xml:"isNotEditable"`
	Mandatory     *FieldMandatory `xml:"mandatory"`
	IsExtended    bool            `xml:"isExtended"`
	NewValue      string          `xml:"newValue"`
	IsSended      bool            `xml:"isSended"`
	*DomainValue
}

// SubOperation ...
type SubOperation struct {
	ColvirReferenceId string       `xml:"colvirReferenceId"`
	OperationNumber   string       `xml:"operationNumber"`
	Name              string       `xml:"name"`
	Date              string       `xml:"date"`
	User              *DomainValue `xml:"user"`
	Description       string       `xml:"description"`
}

// DebugInfo ...
type DebugInfo struct {
	RequestId       string       `xml:"requestId"`
	Procedure       string       `xml:"procedure"`
	ProcedureParams string       `xml:"procedureParams"`
	ProcedureResult string       `xml:"procedureResult"`
	Stack           string       `xml:"stack"`
	Error           *DomainValue `xml:"error"`
	OperationDate   string       `xml:"operationDate"`
	OperationDay    string       `xml:"operationDay"`
	Debug           string       `xml:"debug"`
}

// LoadDebugInfoSearchResult ...
type LoadDebugInfoSearchResult struct {
	Value []*DebugInfo `xml:"value"`
}

// CurrencyConversionReqItem ...
type CurrencyConversionReqItem struct {
	Value float64 `xml:"value"`
	*CurrencyConversionBaseItem
}

// GetParameterMethodParamsID ...
type GetParameterMethodParamsID struct {
	XMLName      xml.Name `xml:"getParameterMethodParamsID"`
	DepartmentId int      `xml:"departmentId"`
	UserId       int      `xml:"userId"`
}

// GetParameterMethodParamsCODE ...
type GetParameterMethodParamsCODE struct {
	XMLName        xml.Name `xml:"getParameterMethodParamsCODE"`
	DepartmentCode string   `xml:"departmentCode"`
	UserCode       string   `xml:"userCode"`
}

// DecisionTable ...
type DecisionTable struct {
	Decision []*Decision `xml:"decision"`
}

// Case ...
type Case struct {
	XMLName xml.Name `xml:"case"`
	Code    string   `xml:"code"`
	Value   string   `xml:"value"`
}

// Decision ...
type Decision struct {
	RowNumberAttr int     `xml:"rowNumber,attr"`
	Case          []*Case `xml:"case"`
}

// ClientBankRelation ...
type ClientBankRelation struct {
	AffiliationFl bool `xml:"affiliationFl"`
	*DomainValue
}

// HierarchyProductTemplate ...
type HierarchyProductTemplate struct {
	ParentCode string `xml:"parentCode"`
	ParentName string `xml:"parentName"`
	Level      int    `xml:"level"`
	*ProductTemplate
}

// Docs ...
type Docs struct {
	XMLName xml.Name       `xml:"docs"`
	Doc     []*DomainValue `xml:"doc"`
}

// DocSet ...
type DocSet struct {
	XMLName xml.Name `xml:"docSet"`
	Docs    *Docs    `xml:"docs"`
	*DomainValue
}

// DossierDocSets ...
type DossierDocSets struct {
	XMLName xml.Name  `xml:"dossierDocSets"`
	DocSet  []*DocSet `xml:"docSet"`
}

// ProductTemplate ...
type ProductTemplate struct {
	IsGroup         bool                     `xml:"isGroup"`
	Concept         *DomainValue             `xml:"concept"`
	BusinessProcess *DomainValue             `xml:"businessProcess"`
	SubjectDomain   *DomainValue             `xml:"subjectDomain"`
	DossierId       int                      `xml:"dossierId"`
	RefName         string                   `xml:"refName"`
	Note            string                   `xml:"note"`
	IsArchival      bool                     `xml:"isArchival"`
	IsArrested      bool                     `xml:"isArrested"`
	Modificated     *Modification            `xml:"modificated"`
	Currency        []*ProductCurrency       `xml:"currency"`
	LoanPurpose     []*ProductPurposeOfLoan  `xml:"loanPurpose"`
	Period          []*ProductTemplatePeriod `xml:"period"`
	Parameter       []*ProductParameter      `xml:"parameter"`
	DossierDocSets  *DossierDocSets          `xml:"dossierDocSets"`
	*DomainValue
}

// ProductTemplatePeriod ...
type ProductTemplatePeriod struct {
	Days int `xml:"days"`
	*DomainValue
}

// ProductCurrency ...
type ProductCurrency struct {
	Currency                        *DomainValue `xml:"currency"`
	Department                      *DomainValue `xml:"department"`
	AmountCurrency                  *DomainValue `xml:"amountCurrency"`
	MinAmount                       string       `xml:"minAmount"`
	MaxAmount                       string       `xml:"maxAmount"`
	MinAmountFormula                *DomainValue `xml:"minAmountFormula"`
	MaxAmountFormula                *DomainValue `xml:"maxAmountFormula"`
	MinAddDeposit                   string       `xml:"minAddDeposit"`
	MaxAddDepsoit                   string       `xml:"maxAddDepsoit"`
	AddDepositSettings              *DomainValue `xml:"addDepositSettings"`
	MinBalance                      string       `xml:"minBalance"`
	MinBalanceAdjustment            *DomainValue `xml:"minBalanceAdjustment"`
	MinBalanceFormula               *DomainValue `xml:"minBalanceFormula"`
	MaxPartPayment                  string       `xml:"maxPartPayment"`
	MaxPartPaymentAdjustment        *DomainValue `xml:"maxPartPaymentAdjustment"`
	MaxPartPaymentFormula           *DomainValue `xml:"maxPartPaymentFormula"`
	MaxPercentPartPayment           float64      `xml:"maxPercentPartPayment"`
	MaxPercentPartPaymentAdjustment *DomainValue `xml:"maxPercentPartPaymentAdjustment"`
	IsForAllDep                     bool         `xml:"isForAllDep"`
}

// ProductPurposeOfLoan ...
type ProductPurposeOfLoan struct {
	Purpose    *DomainValue   `xml:"purpose"`
	Collateral []*DomainValue `xml:"collateral"`
}

// HierarchyProductParameter ...
type HierarchyProductParameter struct {
	ParentCode string `xml:"parentCode"`
	ParentName string `xml:"parentName"`
	Level      int    `xml:"level"`
	*ProductParameter
}

// ProductParameter ...
type ProductParameter struct {
	Domain    *DomainValue `xml:"domain"`
	Value     string       `xml:"value"`
	ValueName string       `xml:"valueName"`
	*DomainValue
}

// Modification ...
type Modification struct {
	Date string       `xml:"date"`
	User *DomainValue `xml:"user"`
}

// ParametersValues ...
type ParametersValues struct {
	Values []*AddInfoValue `xml:"values"`
}

// FilterValues ...
type FilterValues struct {
	Code  string `xml:"code"`
	Value string `xml:"value"`
}

// AltValue ...
type AltValue struct {
	CodeAttr *NotEmptyString `xml:"code,attr"`
	Value    string          `xml:",chardata"`
}

// AltKey ...
type AltKey struct {
	CodeAttr *NotEmptyString `xml:"code,attr"`
	Value    []*AltValue     `xml:"value"`
}

// BusinessObject ...
type BusinessObject struct {
	Id   *NotEmptyString `xml:"id"`
	Code *NotEmptyString `xml:"code"`
	Name *NotEmptyString `xml:"name"`
}

// AvailableBusinessOperation ...
type AvailableBusinessOperation struct {
	Business        *BusinessObject `xml:"business"`
	Nord            int             `xml:"nord"`
	Code            *NotEmptyString `xml:"code"`
	Name            *NotEmptyString `xml:"name"`
	LongName        *NotEmptyString `xml:"LongName"`
	CancelOperation bool            `xml:"cancelOperation"`
	Enable          bool            `xml:"enable"`
}

// DocumentsAvailableOperation ...
type DocumentsAvailableOperation struct {
	DocumentReferenceId *NotEmptyString             `xml:"documentReferenceId"`
	BusinessOperation   *AvailableBusinessOperation `xml:"businessOperation"`
}

// SubjectArea ...
type SubjectArea struct {
	ReferenceId string `xml:"referenceId"`
	Code        string `xml:"code"`
	Name        string `xml:"name"`
}

// SubjectAreaConcept ...
type SubjectAreaConcept struct {
	ReferenceId string       `xml:"referenceId"`
	Code        string       `xml:"code"`
	Name        string       `xml:"name"`
	SubjectArea *SubjectArea `xml:"subjectArea"`
	Domain      *DomainValue `xml:"domain"`
	DataType    string       `xml:"dataType"`
}

// CalculationMethodParameter ...
type CalculationMethodParameter struct {
	Code     string       `xml:"code"`
	Nord     int          `xml:"nord"`
	Name     string       `xml:"name"`
	Domain   *DomainValue `xml:"domain"`
	EditMask string       `xml:"editMask"`
}

// CalculationMethod ...
type CalculationMethod struct {
	IdAttr                 int                           `xml:"id,attr,omitempty"`
	Code                   string                        `xml:"code"`
	Name                   string                        `xml:"name"`
	ViewMask               string                        `xml:"viewMask"`
	IsMultiplyByQuantity   bool                          `xml:"isMultiplyByQuantity"`
	SubjectAreaConceptType []*DomainValue                `xml:"subjectAreaConceptType"`
	Parameter              []*CalculationMethodParameter `xml:"parameter"`
}

// RoundingMethod ...
type RoundingMethod struct {
	IdAttr             int                 `xml:"id,attr,omitempty"`
	Code               string              `xml:"code"`
	Name               string              `xml:"name"`
	DecimalPlaces      int                 `xml:"decimalPlaces"`
	RoundingMethodType *RoundingMethodType `xml:"roundingMethodType"`
	Example            float64             `xml:"example"`
}

// ExchangeRateType ...
type ExchangeRateType struct {
	IdAttr             int    `xml:"id,attr,omitempty"`
	Code               string `xml:"code"`
	Name               string `xml:"name"`
	IsDefault          bool   `xml:"isDefault"`
	IsForwardRate      bool   `xml:"isForwardRate"`
	ForwardDaysCount   int    `xml:"forwardDaysCount"`
	IsForClientBank    bool   `xml:"isForClientBank"`
	IsForDepartment    bool   `xml:"isForDepartment"`
	IsForClientDealing bool   `xml:"isForClientDealing"`
	IsArchival         bool   `xml:"isArchival"`
	Note               string `xml:"note"`
}

// GetMoneyInWordsRequestElem ...
type GetMoneyInWordsRequestElem struct {
	XMLName  xml.Name `xml:"getMoneyInWordsRequestElem"`
	Amount   float64  `xml:"amount"`
	Currency string   `xml:"currency"`
	ShowZero bool     `xml:"showZero"`
	Language string   `xml:"language"`
	*AbstractRequest
}

// GetMoneyInWordsResponseElem ...
type GetMoneyInWordsResponseElem struct {
	XMLName       xml.Name     `xml:"getMoneyInWordsResponseElem"`
	AmountInWords string       `xml:"amountInWords"`
	Result        *DomainValue `xml:"result"`
	*AbstractResponse
}

// GetDecisionRequestElem ...
type GetDecisionRequestElem struct {
	XMLName       xml.Name        `xml:"getDecisionRequestElem"`
	DecisionTable string          `xml:"decisionTable"`
	Params        []*AddInfoValue `xml:"params"`
	*AbstractRequest
}

// GetDecisionResponseElem ...
type GetDecisionResponseElem struct {
	XMLName  xml.Name     `xml:"getDecisionResponseElem"`
	Decision string       `xml:"decision"`
	Result   *DomainValue `xml:"result"`
	*AbstractResponse
}

// LoadClientDocType ...
type LoadClientDocType struct {
	XMLName xml.Name      `xml:"loadClientDocType"`
	AttrFl  bool          `xml:"attrFl"`
	Code    string        `xml:"code"`
	Name    string        `xml:"name"`
	Filter  *ConditionAnd `xml:"filter"`
	*AbstractRequest
}

// LoadRegDocTypeRequest ...
type LoadRegDocTypeRequest struct {
	XMLName xml.Name      `xml:"loadRegDocTypeRequest"`
	Type    *CustomerType `xml:"type"`
	*LoadClientDocType
}

// LoadRegDocTypeResponse ...
type LoadRegDocTypeResponse struct {
	XMLName xml.Name      `xml:"loadRegDocTypeResponse"`
	Result  []*RegDocType `xml:"result"`
	*AbstractResponse
}

// Attributes ...
type Attributes struct {
	XMLName  xml.Name `xml:"attributes"`
	CtSer    string   `xml:"ctSer"`
	CtNum    string   `xml:"ctNum"`
	CtDtFrom string   `xml:"ctDtFrom"`
	CtDtTo   string   `xml:"ctDtTo"`
	CtGni    string   `xml:"ctGni"`
	CtOrg    string   `xml:"ctOrg"`
	CtRegnum string   `xml:"ctRegnum"`
	CtRegdt  string   `xml:"ctRegdt"`
	CtLic    string   `xml:"ctLic"`
	CtLctwrk string   `xml:"ctLctwrk"`
}

// RegDocType ...
type RegDocType struct {
	CliType    string        `xml:"cliType"`
	Usefl      string        `xml:"usefl"`
	UseflName  string        `xml:"useflName"`
	Prim       string        `xml:"prim"`
	DeadocCode string        `xml:"deadocCode"`
	DeadocName string        `xml:"deadocName"`
	Attributes []*Attributes `xml:"attributes"`
	*DomainValue
}

// LoadFamilyRelationshipTypeRequest ...
type LoadFamilyRelationshipTypeRequest struct {
	XMLName xml.Name      `xml:"loadFamilyRelationshipTypeRequest"`
	Code    string        `xml:"code"`
	Name    string        `xml:"name"`
	Sex     string        `xml:"sex"`
	Filter  *ConditionAnd `xml:"filter"`
	*AbstractRequest
}

// LoadFamilyRelationshipTypeResponse ...
type LoadFamilyRelationshipTypeResponse struct {
	XMLName xml.Name        `xml:"loadFamilyRelationshipTypeResponse"`
	Result  []*Relationship `xml:"result"`
	*AbstractResponse
}

// Relationship ...
type Relationship struct {
	Code          string       `xml:"code"`
	Name          string       `xml:"name"`
	Sex           *DomainValue `xml:"sex"`
	CloseFamilyFl bool         `xml:"closeFamilyFl"`
}

// LoadIdentDocTypeRequest ...
type LoadIdentDocTypeRequest struct {
	XMLName    xml.Name `xml:"loadIdentDocTypeRequest"`
	AddCodesFl bool     `xml:"addCodesFl"`
	*LoadClientDocType
}

// LoadIdentDocTypeResponse ...
type LoadIdentDocTypeResponse struct {
	XMLName xml.Name              `xml:"loadIdentDocTypeResponse"`
	Result  []*ClientIdentDocType `xml:"result"`
	*AbstractResponse
}

// ClientIdentDocType ...
type ClientIdentDocType struct {
	TaxCode    string                          `xml:"taxCode"`
	Residfl    string                          `xml:"residfl"`
	Prim       string                          `xml:"prim"`
	Attributes []*ClientIdentDocTypeAttributes `xml:"attributes"`
	AddCodes   []*IdentDocTypeAddCodes         `xml:"addCodes"`
	*DomainValue
}

// ClientIdentDocTypeAttributes ...
type ClientIdentDocTypeAttributes struct {
	CtPasSer  string `xml:"ctPasSer"`
	CtPasFin  string `xml:"ctPasFin"`
	CtDtFrom  string `xml:"ctDtFrom"`
	CtDtTo    string `xml:"ctDtTo"`
	CtCodeDep string `xml:"ctCodeDep"`
	CtDtSen   string `xml:"ctDtSen"`
}

// IdentDocTypeAddCodes ...
type IdentDocTypeAddCodes struct {
	Passcls     string `xml:"passcls"`
	PassclsName string `xml:"passclsName"`
	Code        string `xml:"code"`
}

// LoadAnyDomainValueFullRequest ...
type LoadAnyDomainValueFullRequest struct {
	XMLName    xml.Name `xml:"loadAnyDomainValueFullRequest"`
	DomainName string   `xml:"domainName"`
	*AbstractRequest
}

// LoadAnyDomainValueFullResponse ...
type LoadAnyDomainValueFullResponse struct {
	XMLName xml.Name           `xml:"loadAnyDomainValueFullResponse"`
	Value   []*DomainValueFull `xml:"value"`
	*AbstractResponse
}
