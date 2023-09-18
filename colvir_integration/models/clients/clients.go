package clients

import (
	"encoding/xml"

	"gitlab.apps.bcc.kz/digital-banking-platform/adapters/dbp-colvir-adapter-cap/colvir_integration/models/schema"
)

type LoadClientsBody struct {
	Text    string              `xml:",chardata"`
	Clients LoadClientsListElem `xml:"v1:loadClientsListElem"`
}

type LoadBalanceBody struct {
	Text    string              `xml:",chardata"`
	Clients LoadBalanceListElem `xml:"v1:loadBalanceListElem"`
}

// AccountType - enum with values: CURR, CCUR, TEDP, CARD, OTHERS
type AccountType string

// AccountExternalStatus - enum with values: A,B,C
type AccountExternalStatus string

// AccountIdentifierType - enum with values: clientid,accountid
type AccountIdentifierType string

// SaveClientElem ...
type SaveClientElem struct {
	XMLName                xml.Name `xml:"saveClientElem"`
	Client                 *Client  `xml:"client"`
	WithoutAdditionalCheck int      `xml:"withoutAdditionalCheck"`
	ActionForCheckFail     int      `xml:"actionForCheckFail"`
	OpenClientCard         int      `xml:"openClientCard"`
	*schema.AbstractRequest
}

// SaveClientResponseElem ...
type SaveClientResponseElem struct {
	XMLName              xml.Name                   `xml:"saveClientResponseElem"`
	ClientCode           string                     `xml:"clientCode"`
	NewClientCreated     bool                       `xml:"newClientCreated"`
	RolesOverwrited      bool                       `xml:"rolesOverwrited"`
	OpenClientCardResult *schema.DomainValue        `xml:"openClientCardResult"`
	Result               *schema.DocumentSaveResult `xml:"result"`
	LinkClientCardResult *schema.DomainValue        `xml:"linkClientCardResult"`
	*schema.AbstractResponse
}

// CheckClientElem ...
type CheckClientElem struct {
	XMLName xml.Name `xml:"checkClientElem"`
	Client  *Client  `xml:"client"`
	*schema.AbstractRequest
}

// CheckClientResponseElem ...
type CheckClientResponseElem struct {
	XMLName xml.Name                   `xml:"checkClientResponseElem"`
	Result  *schema.DocumentSaveResult `xml:"result"`
	*schema.AbstractResponse
}

// FindTaxPayerElem ...
type FindTaxPayerElem struct {
	XMLName xml.Name `xml:"findTaxPayerElem"`
	Iin     string   `xml:"iin"`
	*schema.AbstractRequest
}

// FindTaxPayerResponseElem ...
type FindTaxPayerResponseElem struct {
	XMLName          xml.Name `xml:"findTaxPayerResponseElem"`
	Title            string   `xml:"title"`
	TypeTaxpayer     int      `xml:"typeTaxpayer"`
	FlNotResident    int      `xml:"flNotResident"`
	TypePP           []int    `xml:"typePP"`
	FlInactive       int      `xml:"flInactive"`
	TaxAuthorityCode int      `xml:"taxAuthorityCode"`
	*schema.AbstractResponse
}

// LoadClientDetailsElem ...
type LoadClientDetailsElem struct {
	XMLName     xml.Name `xml:"loadClientDetailsElem"`
	Client      string   `xml:"client"`
	Department  string   `xml:"department"`
	AddressType string   `xml:"addressType"`
	*schema.AbstractRequest
}

// LoadClientDetailsResponseElem ...
type LoadClientDetailsResponseElem struct {
	XMLName xml.Name                 `xml:"loadClientDetailsResponseElem"`
	Result  *IndividualPersonDetails `xml:"result"`
	*schema.AbstractResponse
}

// LoadLegalClientDetailsElem ...
type LoadLegalClientDetailsElem struct {
	XMLName    xml.Name `xml:"loadLegalClientDetailsElem"`
	Client     string   `xml:"client"`
	Department string   `xml:"department"`
	*schema.AbstractRequest
}

// LoadLegalClientDetailsResponseElem ...
type LoadLegalClientDetailsResponseElem struct {
	XMLName xml.Name            `xml:"loadLegalClientDetailsResponseElem"`
	Result  *LegalPersonDetails `xml:"result"`
	*schema.AbstractResponse
}

// FindClientByRnnElem ...
type FindClientByRnnElem struct {
	XMLName  xml.Name `xml:"findClientByRnnElem"`
	Rnn      string   `xml:"rnn"`
	IinOrBin string   `xml:"iinOrBin"`
	*schema.AbstractRequest
}

// FindClientByRnnResponseElem ...
type FindClientByRnnResponseElem struct {
	XMLName xml.Name                `xml:"findClientByRnnResponseElem"`
	Result  *ClientEntityIdentifier `xml:"result"`
	*schema.AbstractResponse
}

// LoadDepartmentNameElem ...
type LoadDepartmentNameElem struct {
	XMLName    xml.Name `xml:"loadDepartmentNameElem"`
	Department string   `xml:"department"`
	*schema.AbstractRequest
}

// Result ...
type Result struct {
	XMLName        xml.Name `xml:"result"`
	DepartmentName string   `xml:"departmentName"`
}

// LoadDepartmentNameResponseElem ...
type LoadDepartmentNameResponseElem struct {
	XMLName xml.Name `xml:"loadDepartmentNameResponseElem"`
	Result  *Result  `xml:"result"`
	*schema.AbstractResponse
}

// LoadAuthClientDetailsElem ...
type LoadAuthClientDetailsElem struct {
	XMLName    xml.Name `xml:"loadAuthClientDetailsElem"`
	Client     string   `xml:"client"`
	Department string   `xml:"department"`
	PersonRole string   `xml:"personRole"`
	*schema.AbstractRequest
}

// LoadAuthClientDetailsResponseElem ...
type LoadAuthClientDetailsResponseElem struct {
	XMLName xml.Name `xml:"loadAuthClientDetailsResponseElem"`
	Result  *Result  `xml:"result"`
	*schema.AbstractResponse
}

// LoadAdminClientDetailsElem ...
type LoadAdminClientDetailsElem struct {
	XMLName    xml.Name `xml:"loadAdminClientDetailsElem"`
	Client     string   `xml:"client"`
	Department string   `xml:"department"`
	PersonRole string   `xml:"personRole"`
	*schema.AbstractRequest
}

// LoadAdminClientDetailsResponseElem ...
type LoadAdminClientDetailsResponseElem struct {
	XMLName xml.Name `xml:"loadAdminClientDetailsResponseElem"`
	Result  *Result  `xml:"result"`
	*schema.AbstractResponse
}

// LoadClientsListElem ...
type LoadClientsListElem struct {
	XMLName                 xml.Name             `xml:"v1:loadClientsListElem,omitempty"`
	*schema.AbstractRequest                      //НЕ УБИРАТЬ С ЭТОЙ ПОЗИЦИИ, ПОРЯДОК ВАЖЕН
	ClientsListQuery        *schema.CustomQuery  `xml:"v1:clientsListQuery"`
	DocumentsQuery          ClientsListFilter    `xml:"v1:documentsQuery,omitempty"`
	ContactsQuery           ClientsListFilter    `xml:"v1:contactsQuery,omitempty"`
	ClientRole              string               `xml:"v1:clientRole,omitempty"`
	HideDetails             *int                 `xml:"v1:hideDetails,omitempty"`
	ClientsCode             string               `xml:"v1:clientsCode,omitempty"`
	TaxIdentificationNumber string               `xml:"v1:taxIdentificationNumber,omitempty"`
	DepartmentCode          string               `xml:"v1:departmentCode,omitempty"`
	ClientsRoleStatus       string               `xml:"v1:clientsRoleStatus,omitempty"`
	ClientsType             *schema.CustomerType `xml:"v1:clientsType,omitempty"`
}

// ClientsRoleStatus is Clients type
type ClientsRoleStatus string

// LoadClientsListResponseElem ...
type LoadClientsListResponseElem struct {
	XMLName    xml.Name                     `xml:"loadClientsListResponseElem"`
	ClientList *LoadClientsListSearchResult `xml:"clientList"`
	*schema.AbstractResponse
}

// ClientsListFilter ...
type ClientsListFilter struct {
	Type   *string             `xml:"v1:type,omitempty"`
	Filter *schema.CustomQuery `xml:"v1:filter,omitempty"`
}

// LoadClientAccountsListElem ...
type LoadClientAccountsListElem struct {
	XMLName                 xml.Name            `xml:"loadClientAccountsListElem"`
	ClientCode              string              `xml:"clientCode"`
	ClientAccount           string              `xml:"clientAccount"`
	ClientAccountsListQuery schema.ConditionAnd `xml:"clientAccountsListQuery"`
	*schema.AbstractRequest
}

// LoadClientAccountsListResponseElem ...
type LoadClientAccountsListResponseElem struct {
	XMLName      xml.Name                            `xml:"loadClientAccountsListResponseElem"`
	AccountsList *LoadClientAccountsListSearchResult `xml:"accountsList"`
	*schema.AbstractResponse
}

// LoadClientAccountsCountingElem ...
type LoadClientAccountsCountingElem struct {
	XMLName      xml.Name        `xml:"loadClientAccountsCountingElem"`
	AccountItems []AccountIdItem `xml:"accountItems"`
	Statuses     []string        `xml:"statuses"`
	*schema.AbstractRequest
}

// LoadClientAccountsCountingResponseElem ...
type LoadClientAccountsCountingResponseElem struct {
	XMLName              xml.Name                          `xml:"loadClientAccountsCountingResponseElem"`
	AccountsCointingList *LoadAccountsCountingSearchResult `xml:"accountsCointingList"`
	*schema.AbstractResponse
}

// LoadSignatoryListElem ...
type LoadSignatoryListElem struct {
	XMLName    xml.Name `xml:"loadSignatoryListElem"`
	ClientCode string   `xml:"clientCode"`
	*schema.AbstractRequest
}

// LoadSignatoryListResponseElem ...
type LoadSignatoryListResponseElem struct {
	XMLName       xml.Name                       `xml:"loadSignatoryListResponseElem"`
	SignatoryList *LoadSignatoryListSearchResult `xml:"signatoryList"`
	*schema.AbstractResponse
}

// LoadBalanceListElem ...
type LoadBalanceListElem struct {
	XMLName xml.Name `xml:"v1:loadBalanceListElem"`
	*schema.AbstractRequest
	AccountCode []string `xml:"v1:accountCode"`
}

// LoadBalanceListResponseElem ...
type LoadBalanceListResponseElem struct {
	XMLName            xml.Name                     `xml:"loadBalanceListResponseElem"`
	AccountBalanceList *LoadBalanceListSearchResult `xml:"accountBalanceList"`
	*schema.AbstractResponse
}

// LoadClientsListSearchResult ...
type LoadClientsListSearchResult struct {
	ClientListItem []*Client `xml:"clientListItem"`
}

// ClientRiskValuesList ...
type ClientRiskValuesList struct {
	XMLName         xml.Name           `xml:"clientRiskValuesList"`
	ClientRiskValue []*ClientRiskValue `xml:"clientRiskValue"`
}

// Client ...
type Client struct {
	ColvirReferenceId       string                    `xml:"colvirReferenceId"`
	Code                    string                    `xml:"code"`
	Department              *schema.DomainValue       `xml:"department"`
	Type                    schema.CustomerType       `xml:"type"`
	Name                    string                    `xml:"name"`
	LongName                string                    `xml:"longName"`
	FirstName               string                    `xml:"firstName"`
	MiddleName              string                    `xml:"middleName"`
	LastName                string                    `xml:"lastName"`
	PreviousName            string                    `xml:"previousName"`
	BirthDate               string                    `xml:"birthDate"`
	TaxIdentificationNumber *schema.DomainValue       `xml:"taxIdentificationNumber"`
	RolesGroup              string                    `xml:"rolesGroup"`
	RoleMask                string                    `xml:"roleMask"`
	State                   *schema.DomainValue       `xml:"state"`
	Status                  *schema.DomainValue       `xml:"status"`
	TypeExt                 *schema.DomainValue       `xml:"typeExt"`
	Kod                     *schema.DomainValue       `xml:"kod"`
	IsResident              bool                      `xml:"isResident"`
	Sex                     string                    `xml:"sex"`
	FromDate                string                    `xml:"fromDate"`
	Dreg                    string                    `xml:"dreg"`
	AutCode                 *schema.DomainValue       `xml:"autCode"`
	SectorId                *schema.DomainValue       `xml:"sectorId"`
	OkpoCode                string                    `xml:"okpoCode"`
	Country                 *schema.DomainValue       `xml:"country"`
	CitizenCountry          *schema.DomainValue       `xml:"citizenCountry"`
	IdentDocuments          []*IdentificationDocument `xml:"identDocuments"`
	Addresses               []*Address                `xml:"addresses"`
	DetailedAddresses       []*DetailedAddress        `xml:"detailedAddresses"`
	Emails                  []*Email                  `xml:"emails"`
	Phones                  []*Phone                  `xml:"phones"`
	ContactData             []*ContactData            `xml:"contactData"`
	AddInfoList             []*schema.AddInfoValue    `xml:"addInfoList"`
	ClientRiskValuesList    *ClientRiskValuesList     `xml:"clientRiskValuesList"`
	AffiliationFl           bool                      `xml:"affiliationFl"`
	ShiId                   *schema.DomainValue       `xml:"shiId"`
	Tariff                  *schema.DomainValue       `xml:"tariff"`
	CodeWord                string                    `xml:"codeWord"`
	CodeWordDate            string                    `xml:"codeWordDate"`
	CodeWordQuestion        string                    `xml:"codeWordQuestion"`
	ClientClassifier        []*ClientClassifier       `xml:"clientClassifier"`
	LatinFirstName          string                    `xml:"latinFirstName"`
	LatinMiddleName         string                    `xml:"latinMiddleName"`
	LatinLastName           string                    `xml:"latinLastName"`
	DeceaseDate             string                    `xml:"deceaseDate"`
	LinkedCards             []*LinkedCard             `xml:"linkedCards"`
	ClientBankRelationLinks []*ClientBankRelationLink `xml:"clientBankRelationLinks"`
}

// LinkedCard ...
type LinkedCard struct {
	LinkedCode string              `xml:"linkedCode"`
	Type       schema.CustomerType `xml:"type"`
	Name       string              `xml:"name"`
	LongName   string              `xml:"longName"`
}

// ClientsLinkedCard ...
type ClientsLinkedCard struct {
	ColvirReferenceId string `xml:"colvirReferenceId"`
	ClientCode        string `xml:"clientCode"`
	*LinkedCard
}

// ContactData ...
type ContactData struct {
	Nord      int                 `xml:"nord"`
	Type      *schema.DomainValue `xml:"type"`
	Kind      *schema.DomainValue `xml:"kind"`
	Value     string              `xml:"value"`
	IsDefault bool                `xml:"isDefault"`
}

// Address ...
type Address struct {
	Nord          int                 `xml:"nord"`
	Country       string              `xml:"country"`
	Type          *schema.DomainValue `xml:"type"`
	Kind          *schema.DomainValue `xml:"kind"`
	Okato         string              `xml:"okato"`
	AddressString string              `xml:"addressString"`
}

// Phone ...
type Phone struct {
	*ContactData
}

// Email ...
type Email struct {
	*ContactData
}

// ClientsContactData ...
type ClientsContactData struct {
	ColvirReferenceId string `xml:"colvirReferenceId"`
	ClientCode        string `xml:"clientCode"`
	IsPhone           bool   `xml:"isPhone"`
	IsEmail           bool   `xml:"isEmail"`
	IsWww             bool   `xml:"isWww"`
	*ContactData
}

// ClientsAddress ...
type ClientsAddress struct {
	ColvirReferenceId string `xml:"colvirReferenceId"`
	ClientCode        string `xml:"clientCode"`
	*Address
}

// IdentificationDocument ...
type IdentificationDocument struct {
	Nord           int                 `xml:"nord"`
	Type           *schema.DomainValue `xml:"type"`
	TypeExtCode    string              `xml:"typeExtCode"`
	Serie          string              `xml:"serie"`
	Number         string              `xml:"number"`
	IssueDate      string              `xml:"issueDate"`
	StartDate      string              `xml:"startDate"`
	ExpirationDate string              `xml:"expirationDate"`
	Issuer         string              `xml:"issuer"`
	IsDefault      bool                `xml:"isDefault"`
	IsArchival     bool                `xml:"isArchival"`
}

// ClientSaveResult ...
type ClientSaveResult struct {
	Result   string `xml:"result"`
	ClientId string `xml:"clientId"`
	*schema.ValidationResponse
}

// Region ...
type Region struct {
	XMLName          xml.Name `xml:"region"`
	IsCityRegionAttr bool     `xml:"isCityRegion,attr,omitempty"`
	schema.DomainDecimal
}

// DetailedAddress ...
type DetailedAddress struct {
	Nord          int                   `xml:"nord"`
	Type          *schema.DomainValue   `xml:"type"`
	Country       string                `xml:"country"`
	Region        *Region               `xml:"region"`
	District      *schema.DomainDecimal `xml:"district"`
	City          *schema.DomainDecimal `xml:"city"`
	CityPart      *schema.DomainDecimal `xml:"cityPart"`
	CityZone      *schema.DomainDecimal `xml:"cityZone"`
	Street        *schema.DomainDecimal `xml:"street"`
	HouseNumber   *schema.DomainDecimal `xml:"houseNumber"`
	Flat          *schema.DomainDecimal `xml:"flat"`
	BodyNumber    string                `xml:"bodyNumber"`
	BuildNumber   string                `xml:"buildNumber"`
	Zip           string                `xml:"zip"`
	Kind          *schema.DomainValue   `xml:"kind"`
	Okato         string                `xml:"okato"`
	AddressString string                `xml:"addressString"`
}

// ClientEntityIdentifier ...
type ClientEntityIdentifier struct {
	Client       string `xml:"client"`
	Department   string `xml:"department"`
	CustomerType string `xml:"customerType"`
}

// TaxPayerPersonDetails ...
type TaxPayerPersonDetails struct {
	Rnn string `xml:"rnn"`
	Iin string `xml:"iin"`
}

// BasePersonDetails ...
type BasePersonDetails struct {
	IsResident               int    `xml:"isResident"`
	Residence                string `xml:"residence"`
	IsArchived               int    `xml:"isArchived"`
	IsAccountOpeningDisabled int    `xml:"isAccountOpeningDisabled"`
	State                    string `xml:"state"`
	*TaxPayerPersonDetails
}

// RoledPersonDetails ...
type RoledPersonDetails struct {
	Role         string `xml:"role"`
	Prefix       string `xml:"prefix"`
	FirstName    string `xml:"firstName"`
	MiddleName   string `xml:"middleName"`
	LastName     string `xml:"lastName"`
	CustomerType string `xml:"customerType"`
	*TaxPayerPersonDetails
}

// IndividualPersonDetails ...
type IndividualPersonDetails struct {
	Prefix     string `xml:"prefix"`
	FirstName  string `xml:"firstName"`
	MiddleName string `xml:"middleName"`
	LastName   string `xml:"lastName"`
	Address    string `xml:"address"`
	*BasePersonDetails
}

// LegalPersonDetails ...
type LegalPersonDetails struct {
	Name      string `xml:"name"`
	ShortName string `xml:"shortName"`
	Property  string `xml:"property"`
	Branch    string `xml:"branch"`
	*BasePersonDetails
}

// AuthorizedPersonDetails ...
type AuthorizedPersonDetails struct {
	Education                 string                     `xml:"education"`
	FirstSignature            int                        `xml:"firstSignature"`
	SecondSignature           int                        `xml:"secondSignature"`
	AuthorizedClientsDocument *AuthorizedClientsDocument `xml:"authorizedClientsDocument"`
	*RoledPersonDetails
}

// AuthorizedClientsDocument ...
type AuthorizedClientsDocument struct {
	FromDate   string              `xml:"fromDate"`
	ToDate     string              `xml:"toDate"`
	State      *schema.DomainValue `xml:"state"`
	DocType    *schema.DomainValue `xml:"docType"`
	DocNum     string              `xml:"docNum"`
	NotaryName string              `xml:"notaryName"`
	IssueOrg   string              `xml:"issueOrg"`
}

// AdministrationPersonDetails ...
type AdministrationPersonDetails struct {
	AffiliatedCode        string `xml:"affiliatedCode"`
	AffiliatedDescription string `xml:"affiliatedDescription"`
	Note                  string `xml:"note"`
	StartDate             string `xml:"startDate"`
	FinishDate            string `xml:"finishDate"`
	*RoledPersonDetails
}

// LoadClientAccountsListSearchResult ...
type LoadClientAccountsListSearchResult struct {
	ClientAccountsListItem []*ClientAccount `xml:"clientAccountsListItem"`
}

// ClientAccount ...
type ClientAccount struct {
	ClientCode     string    `xml:"clientCode"`
	ClientAccounts []Account `xml:"clientAccounts"`
}

// AccountLockCurrent ...
type AccountLockCurrent struct {
	AccountCode      schema.AccountNumber `xml:"accountCode"`
	Value            *schema.DomainValue  `xml:"value"`
	AccountLocksList AccountLocksList     `xml:"accountLocksList"`
}

// AccountLocksList ...
type AccountLocksList struct {
	AccountLock []AccountLock `xml:"accountLock"`
}

// AccountLock ...
type AccountLock struct {
	AccountCode     schema.AccountNumber `xml:"accountCode"`
	LockId          string               `xml:"lockId"`
	LockType        *schema.DomainValue  `xml:"lockType"`
	FromDate        string               `xml:"fromDate"`
	ToDate          string               `xml:"toDate"`
	Description     string               `xml:"description"`
	DecTableCash    string               `xml:"decTableCash"`
	ByTblCash       *schema.DomainValue  `xml:"byTblCash"`
	DecTableNonCash string               `xml:"decTableNonCash"`
	ByTblNonCash    *schema.DomainValue  `xml:"byTblNonCash"`
	LockOrdNum      string               `xml:"lockOrdNum"`
	LockOrdDate     string               `xml:"lockOrdDate"`
	LockOrgCode     string               `xml:"lockOrgCode"`
	LockOrgName     string               `xml:"lockOrgName"`
	Refer           string               `xml:"refer"`
}

// Account ...
type Account struct {
	ClientCode           string               `xml:"clientCode"`
	Number               schema.AccountNumber `xml:"number"`
	Iban                 schema.AccountIban   `xml:"iban"`
	Type                 AccountType          `xml:"type"`
	Currency             *schema.DomainValue  `xml:"currency"`
	Activfl              bool                 `xml:"activfl"`
	StatusExtCode        string               `xml:"statusExtCode"`
	Status               *schema.DomainValue  `xml:"status"`
	Title                string               `xml:"title"`
	OwnerName            string               `xml:"ownerName"`
	Balance              float64              `xml:"balance"`
	BalanceNatVal        float64              `xml:"balanceNatVal"`
	BlockedBalance       float64              `xml:"blockedBalance"`
	BlockedBalanceNatVal float64              `xml:"blockedBalanceNatVal"`
	AccountLockCurrent   AccountLockCurrent   `xml:"accountLockCurrent"`
	Branch               *schema.DomainValue  `xml:"branch"`
	Bank                 *schema.DomainValue  `xml:"bank"`
	Product              *schema.DomainValue  `xml:"product"`
	AccountPlanType      *schema.DomainValue  `xml:"accountPlanType"`
}

// AccountIdItem ...
type AccountIdItem struct {
	IdType  string `xml:"idType"`
	IdValue string `xml:"idValue"`
}

// AccountCounting ...
type AccountCounting struct {
	IdType  AccountIdentifierType `xml:"idType"`
	IdValue schema.AccountNumber  `xml:"idValue"`
	Status  AccountExternalStatus `xml:"status"`
	Count   int                   `xml:"count"`
}

// LoadAccountsCountingSearchResult ...
type LoadAccountsCountingSearchResult struct {
	AccountCounting []AccountCounting `xml:"accountCounting"`
}

// LoadSignatoryListSearchResult ...
type LoadSignatoryListSearchResult struct {
	SignatoryList []*Signatory `xml:"signatoryList"`
}

// Signatory ...
type Signatory struct {
	Code        string `xml:"code"`
	Title       string `xml:"title"`
	Name        string `xml:"name"`
	Description string `xml:"description"`
	Signfl      bool   `xml:"signfl"`
	Begdate     string `xml:"begdate"`
	Enddate     string `xml:"enddate"`
	Sgnname     string `xml:"sgn_name"`
	DocNum      string `xml:"docNum"`
}

// Balance ...
type Balance struct {
	Currency                  string  `xml:"currency"`
	IntradayMovementDebit     float64 `xml:"intradayMovementDebit"`
	IntradayMovementCredit    float64 `xml:"intradayMovementCredit"`
	PendingTransactionDebit   float64 `xml:"pendingTransactionDebit"`
	PendingTransactionCredit  float64 `xml:"pendingTransactionCredit"`
	ReservedTransactionDebit  float64 `xml:"reservedTransactionDebit"`
	ReservedTransactionCredit float64 `xml:"reservedTransactionCredit"`
	LastStatementBalance      float64 `xml:"lastStatementBalance"`
	CurrentBalance            float64 `xml:"currentBalance"`
	OverdraftLimit            float64 `xml:"overdraftLimit"`
	TechnicalOverdraft        float64 `xml:"technicalOverdraft"`
	AvailableFunds            float64 `xml:"availableFunds"`
	BlockedBalance            float64 `xml:"blockedBalance"`
	ExpiredDocumentsBalance   float64 `xml:"expiredDocumentsBalance"`
}

// AccountBalance ...
type AccountBalance struct {
	Iban                      schema.AccountNumber   `xml:"iban"`
	TypeCode                  string                 `xml:"typeCode"`
	Type                      string                 `xml:"type"`
	Currency                  string                 `xml:"currency"`
	Activfl                   bool                   `xml:"activfl"`
	Status                    *AccountExternalStatus `xml:"status"`
	Title                     string                 `xml:"title"`
	OwnerName                 string                 `xml:"ownerName"`
	TaxIdentificationNumber   string                 `xml:"TaxIdentificationNumber"`
	DateOpened                string                 `xml:"dateOpened"`
	DateLastPosting           string                 `xml:"dateLastPosting"`
	DateLastStatement         string                 `xml:"dateLastStatement"`
	BalanceAccountCurrency    *Balance               `xml:"balanceAccountCurrency"`
	BalanceNatCurrency        *Balance               `xml:"balanceNatCurrency"`
	TermDepositOpenDate       string                 `xml:"termDepositOpenDate"`
	TermDepositClosingDate    string                 `xml:"termDepositClosingDate"`
	TermDepositContractNumber string                 `xml:"termDepositContractNumber"`
	TermDepositInterestRate   float64                `xml:"termDepositInterestRate"`
	TermDepositCapitalization *schema.DomainValue    `xml:"termDepositCapitalization"`
	TermDepositMinimumBalance float64                `xml:"termDepositMinimumBalance"`
}

// LoadBalanceListSearchResult ...
type LoadBalanceListSearchResult struct {
	AccountBalance []AccountBalance `xml:"accountBalance"`
}

// RoleAdditionalAttrSearchResult ...
type RoleAdditionalAttrSearchResult struct {
	RoleField []*schema.RoleField `xml:"roleField"`
}

// LoadRoleAdditionalAttrElem ...
type LoadRoleAdditionalAttrElem struct {
	XMLName           xml.Name            `xml:"loadRoleAdditionalAttrElem"`
	Type              schema.CustomerType `xml:"type"`
	RolesMask         string              `xml:"rolesMask"`
	ColvirReferenceId string              `xml:"colvirReferenceId"`
	Parameters        string              `xml:"parameters"`
	*schema.AbstractRequest
}

// LoadRoleAdditionalAttrResponseElem ...
type LoadRoleAdditionalAttrResponseElem struct {
	XMLName        xml.Name                        `xml:"loadRoleAdditionalAttrResponseElem"`
	RoleFieldsList *RoleAdditionalAttrSearchResult `xml:"roleFieldsList"`
	*schema.AbstractResponse
}

// AddressType ...
type AddressType struct {
	LongName  string `xml:"longName"`
	Mandatory bool   `xml:"mandatory"`
	Multi     bool   `xml:"multi"`
	schema.DomainValue
}

// ClientAddressTypeSearchResult ...
type ClientAddressTypeSearchResult struct {
	AddressType []*AddressType `xml:"addressType"`
}

// LoadClientAddressTypeListElem ...
type LoadClientAddressTypeListElem struct {
	XMLName xml.Name            `xml:"loadClientAddressTypeListElem"`
	Code    string              `xml:"code"`
	Type    schema.CustomerType `xml:"type"`
	*schema.AbstractRequest
}

// LoadClientAddressTypeListResponseElem ...
type LoadClientAddressTypeListResponseElem struct {
	XMLName               xml.Name                       `xml:"loadClientAddressTypeListResponseElem"`
	ClientAddressTypeList *ClientAddressTypeSearchResult `xml:"clientAddressTypeList"`
	*schema.AbstractResponse
}

// SaveClientCodeWordRequest ...
type SaveClientCodeWordRequest struct {
	XMLName        xml.Name `xml:"saveClientCodeWordRequest"`
	ClientCode     string   `xml:"clientCode"`
	Codeword       string   `xml:"codeword"`
	ExpirationDate string   `xml:"expirationDate"`
	*schema.AbstractRequest
}

// SaveClientCodeWordResponse ...
type SaveClientCodeWordResponse struct {
	XMLName       xml.Name            `xml:"saveClientCodeWordResponse"`
	ExecuteResult *schema.DomainValue `xml:"executeResult"`
	*schema.AbstractResponse
}

// ClientRiskValue ...
type ClientRiskValue struct {
	ColvirReferenceId string              `xml:"colvirReferenceId"`
	Nord              int                 `xml:"nord"`
	FromDate          string              `xml:"fromDate"`
	ToDate            string              `xml:"toDate"`
	ActualFl          bool                `xml:"actualFl"`
	RiskValue         *schema.DomainValue `xml:"riskValue"`
	Reason            string              `xml:"reason"`
	Note              string              `xml:"note"`
	UserName          string              `xml:"userName"`
	CorrectionDate    string              `xml:"correctionDate"`
	RiskLevel         float64             `xml:"riskLevel"`
}

// RoleClassifierDescription ...
type RoleClassifierDescription struct {
	Role         string              `xml:"role"`
	Mandatory    *schema.DomainValue `xml:"mandatory"`
	Classifier   *schema.DomainValue `xml:"classifier"`
	Multiple     bool                `xml:"multiple"`
	CustomerType schema.CustomerType `xml:"customerType"`
}

// ClassifierValue ...
type ClassifierValue struct {
	ToDate string `xml:"toDate"`
	Nord   int    `xml:"nord"`
	*schema.AddInfoValue
}

// ClientClassifier ...
type ClientClassifier struct {
	ColvirReferenceId string              `xml:"colvirReferenceId"`
	Classifier        *schema.DomainValue `xml:"classifier"`
	Value             *ClassifierValue    `xml:"value"`
}

// LoadRoleClassifiersRequest ...
type LoadRoleClassifiersRequest struct {
	XMLName      xml.Name            `xml:"loadRoleClassifiersRequest"`
	Role         string              `xml:"role"`
	CustomerType schema.CustomerType `xml:"customerType"`
	*schema.AbstractRequest
}

// LoadRoleClassifiersResponse ...
type LoadRoleClassifiersResponse struct {
	XMLName xml.Name                     `xml:"loadRoleClassifiersResponse"`
	Result  []*RoleClassifierDescription `xml:"result"`
	*schema.AbstractResponse
}

// LoadClientClassifiersRequest ...
type LoadClientClassifiersRequest struct {
	XMLName    xml.Name `xml:"loadClientClassifiersRequest"`
	ClientCode string   `xml:"clientCode"`
	*schema.AbstractRequest
}

// LoadClientClassifiersResponse ...
type LoadClientClassifiersResponse struct {
	XMLName       xml.Name            `xml:"loadClientClassifiersResponse"`
	ExecuteResult *schema.DomainValue `xml:"executeResult"`
	Result        []*ClientClassifier `xml:"result"`
	*schema.AbstractResponse
}

// SaveClientClassifierRequest ...
type SaveClientClassifierRequest struct {
	XMLName    xml.Name `xml:"saveClientClassifierRequest"`
	ClientCode string   `xml:"clientCode"`
	Nord       int      `xml:"nord"`
	Classifier string   `xml:"classifier"`
	FromDate   string   `xml:"fromDate"`
	ToDate     string   `xml:"toDate"`
	Multiply   bool     `xml:"multiply"`
	Reason     string   `xml:"reason"`
	Value      string   `xml:"value"`
	*schema.AbstractRequest
}

// SaveClientClassifierResponse ...
type SaveClientClassifierResponse struct {
	XMLName       xml.Name            `xml:"saveClientClassifierResponse"`
	ExecuteResult *schema.DomainValue `xml:"executeResult"`
	Nord          int                 `xml:"nord"`
	*schema.AbstractResponse
}

// DeleteClientClassifierRequest ...
type DeleteClientClassifierRequest struct {
	XMLName    xml.Name `xml:"deleteClientClassifierRequest"`
	ClientCode string   `xml:"clientCode"`
	Nord       int      `xml:"nord"`
	Reason     string   `xml:"reason"`
	*schema.AbstractRequest
}

// DeleteClientClassifierResponse ...
type DeleteClientClassifierResponse struct {
	XMLName       xml.Name            `xml:"deleteClientClassifierResponse"`
	ExecuteResult *schema.DomainValue `xml:"executeResult"`
	*schema.AbstractResponse
}

// LoadClientsFamilyListRequest ...
type LoadClientsFamilyListRequest struct {
	XMLName           xml.Name            `xml:"loadClientsFamilyListRequest"`
	ColvirReferenceId string              `xml:"colvirReferenceId"`
	FilterClient      schema.ConditionAnd `xml:"filterClient"`
	FilterFamily      schema.ConditionAnd `xml:"filterFamily"`
	*schema.AbstractRequest
}

// ClientsFamilyList ...
type ClientsFamilyList struct {
	XMLName       xml.Name         `xml:"clientsFamilyList"`
	ClientsFamily []*ClientsFamily `xml:"clientsFamily"`
}

// LoadClientsFamilyListResponse ...
type LoadClientsFamilyListResponse struct {
	XMLName           xml.Name            `xml:"loadClientsFamilyListResponse"`
	ExecuteResult     *schema.DomainValue `xml:"executeResult"`
	ClientsFamilyList *ClientsFamilyList  `xml:"clientsFamilyList"`
	*schema.AbstractResponse
}

// SaveClientsFamilyMemberRequest ...
type SaveClientsFamilyMemberRequest struct {
	XMLName      xml.Name      `xml:"saveClientsFamilyMemberRequest"`
	FamilyMember *FamilyMember `xml:"familyMember"`
	Reason       string        `xml:"reason"`
	*schema.AbstractRequest
}

// SaveClientsFamilyMemberResponse ...
type SaveClientsFamilyMemberResponse struct {
	XMLName       xml.Name            `xml:"saveClientsFamilyMemberResponse"`
	ClientCode    string              `xml:"clientCode"`
	Nord          int                 `xml:"nord"`
	ExecuteResult *schema.DomainValue `xml:"executeResult"`
	*schema.AbstractResponse
}

// DeleteClientsFamilyMemberRequest ...
type DeleteClientsFamilyMemberRequest struct {
	XMLName    xml.Name `xml:"deleteClientsFamilyMemberRequest"`
	ClientCode string   `xml:"clientCode"`
	Nord       int      `xml:"nord"`
	Reason     string   `xml:"reason"`
	*schema.AbstractRequest
}

// DeleteClientsFamilyMemberResponse ...
type DeleteClientsFamilyMemberResponse struct {
	XMLName       xml.Name            `xml:"deleteClientsFamilyMemberResponse"`
	ExecuteResult *schema.DomainValue `xml:"executeResult"`
	*schema.AbstractResponse
}

// ClientsFamily ...
type ClientsFamily struct {
	Code              string          `xml:"code"`
	ColvirReferenceId string          `xml:"colvirReferenceId"`
	FamilyMember      []*FamilyMember `xml:"familyMember"`
}

// FamilyMember ...
type FamilyMember struct {
	Nord              int                 `xml:"nord"`
	ColvirReferenceId string              `xml:"colvirReferenceId"`
	Code              string              `xml:"code"`
	Relationship      *Relationship       `xml:"relationship"`
	FirstName         string              `xml:"firstName"`
	MiddleName        string              `xml:"middleName"`
	LastName          string              `xml:"lastName"`
	Sex               *schema.DomainValue `xml:"sex"`
	BirthDate         string              `xml:"birthDate"`
	FromDate          string              `xml:"fromDate"`
	ToDate            string              `xml:"toDate"`
	ActivityKind      *schema.DomainValue `xml:"activityKind"`
	Education         *schema.DomainValue `xml:"education"`
	AverageIncome     float64             `xml:"averageIncome"`
	ClientCodeRef     string              `xml:"clientCodeRef"`
	ClientIdRef       string              `xml:"clientIdRef"`
	DependencyFl      bool                `xml:"dependencyFl"`
	DependencyStart   string              `xml:"dependencyStart"`
	DependencyEnd     string              `xml:"dependencyEnd"`
	BankWorkerFl      bool                `xml:"bankWorkerFl"`
	ActiveFl          bool                `xml:"activeFl"`
	Organization      string              `xml:"organization"`
	Note              string              `xml:"note"`
}

// Relationship ...
type Relationship struct {
	Code          string              `xml:"code"`
	LongName      string              `xml:"longName"`
	Sex           *schema.DomainValue `xml:"sex"`
	CloseFamilyFl bool                `xml:"closeFamilyFl"`
}

// LinkClientsCardsRequest ...
type LinkClientsCardsRequest struct {
	XMLName    xml.Name `xml:"linkClientsCardsRequest"`
	ClientCode string   `xml:"clientCode"`
	LinkedCode string   `xml:"linkedCode"`
	*schema.AbstractRequest
}

// LinkClientsCardsResponse ...
type LinkClientsCardsResponse struct {
	XMLName       xml.Name            `xml:"linkClientsCardsResponse"`
	ExecuteResult *schema.DomainValue `xml:"executeResult"`
	*schema.AbstractResponse
}

// UndoLinkClientsCardsRequest ...
type UndoLinkClientsCardsRequest struct {
	XMLName    xml.Name `xml:"undoLinkClientsCardsRequest"`
	ClientCode string   `xml:"clientCode"`
	*schema.AbstractRequest
}

// UndoLinkClientsCardsResponse ...
type UndoLinkClientsCardsResponse struct {
	XMLName       xml.Name            `xml:"undoLinkClientsCardsResponse"`
	ExecuteResult *schema.DomainValue `xml:"executeResult"`
	*schema.AbstractResponse
}

// ClientBankRelationLink ...
type ClientBankRelationLink struct {
	ColvirReferenceId string              `xml:"colvirReferenceId"`
	Client            *schema.DomainValue `xml:"client"`
	Nord              int                 `xml:"nord"`
	FromDate          string              `xml:"fromDate"`
	ToDate            string              `xml:"toDate"`
	UpdateType        bool                `xml:"updateType"`
	*schema.ClientBankRelation
}

// LoadClientBankRelationLinkRequest ...
type LoadClientBankRelationLinkRequest struct {
	XMLName    xml.Name `xml:"loadClientBankRelationLinkRequest"`
	ClientCode string   `xml:"clientCode"`
	*schema.AbstractRequest
}

// LoadClientBankRelationLinkResponse ...
type LoadClientBankRelationLinkResponse struct {
	XMLName       xml.Name                  `xml:"loadClientBankRelationLinkResponse"`
	ExecuteResult *schema.DomainValue       `xml:"executeResult"`
	Result        []*ClientBankRelationLink `xml:"result"`
	*schema.AbstractResponse
}

// SaveClientBankRelationLinkRequest ...
type SaveClientBankRelationLinkRequest struct {
	XMLName             xml.Name `xml:"saveClientBankRelationLinkRequest"`
	ClientCode          string   `xml:"clientCode"`
	Nord                int      `xml:"nord"`
	RelationType        string   `xml:"relationType"`
	RelationDescription string   `xml:"relationDescription"`
	FromDate            string   `xml:"fromDate"`
	ToDate              string   `xml:"toDate"`
	UpdateType          bool     `xml:"updateType"`
	*schema.AbstractRequest
}

// SaveClientBankRelationLinkResponse ...
type SaveClientBankRelationLinkResponse struct {
	XMLName       xml.Name            `xml:"saveClientBankRelationLinkResponse"`
	ExecuteResult *schema.DomainValue `xml:"executeResult"`
	Nord          int                 `xml:"nord"`
	*schema.AbstractResponse
}

// DeleteClientBankRelationLinkRequest ...
type DeleteClientBankRelationLinkRequest struct {
	XMLName       xml.Name `xml:"deleteClientBankRelationLinkRequest"`
	ClientCode    string   `xml:"clientCode"`
	Nord          int      `xml:"nord"`
	DeleteType    bool     `xml:"deleteType"`
	DeleteHistory bool     `xml:"deleteHistory"`
	*schema.AbstractRequest
}

// DeleteClientBankRelationLinkResponse ...
type DeleteClientBankRelationLinkResponse struct {
	XMLName       xml.Name            `xml:"deleteClientBankRelationLinkResponse"`
	ExecuteResult *schema.DomainValue `xml:"executeResult"`
	*schema.AbstractResponse
}

// ChangeClientResidenceRequest ...
type ChangeClientResidenceRequest struct {
	XMLName    xml.Name `xml:"changeClientResidenceRequest"`
	ClientCode string   `xml:"clientCode"`
	IsResident bool     `xml:"isResident"`
	Residence  string   `xml:"residence"`
	*schema.AbstractRequest
}

// ChangeClientResidenceResponse ...
type ChangeClientResidenceResponse struct {
	XMLName       xml.Name            `xml:"changeClientResidenceResponse"`
	ExecuteResult *schema.DomainValue `xml:"executeResult"`
	*schema.AbstractResponse
}

// UndoChangeClientResidenceRequest ...
type UndoChangeClientResidenceRequest struct {
	XMLName    xml.Name `xml:"undoChangeClientResidenceRequest"`
	ClientCode string   `xml:"clientCode"`
	*schema.AbstractRequest
}

// UndoChangeClientResidenceResponse ...
type UndoChangeClientResidenceResponse struct {
	XMLName       xml.Name            `xml:"undoChangeClientResidenceResponse"`
	ExecuteResult *schema.DomainValue `xml:"executeResult"`
	*schema.AbstractResponse
}

// LoadClientsListCustomElem ...
type LoadClientsListCustomElem struct {
	XMLName      xml.Name               `xml:"loadClientsListCustomElem"`
	FilterCode   string                 `xml:"filterCode"`
	FilterParams []*schema.FilterValues `xml:"filterParams"`
	*schema.AbstractRequest
}

// LoadClientsListCustomResponseElem ...
type LoadClientsListCustomResponseElem struct {
	XMLName    xml.Name                     `xml:"loadClientsListCustomResponseElem"`
	Result     *schema.DomainValue          `xml:"result"`
	ClientList *LoadClientsListSearchResult `xml:"clientList"`
	*schema.AbstractResponse
}

// LoadClientDetailsFullElem ...
type LoadClientDetailsFullElem struct {
	XMLName    xml.Name `xml:"loadClientDetailsFullElem"`
	ClientCode string   `xml:"clientCode"`
	*schema.AbstractRequest
}

// LoadClientDetailsFullResponseElem ...
type LoadClientDetailsFullResponseElem struct {
	XMLName xml.Name  `xml:"loadClientDetailsFullResponseElem"`
	Client  []*Client `xml:"client"`
	*schema.AbstractResponse
}
