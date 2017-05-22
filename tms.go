package main

import "C"

import (
	"encoding/json"
)

type PinEncryptionMethod string

const (
	_ PinEncryptionMethod = ""

	Dukpt         = "dukpt"
	MasterSession = "mk_sk"
)

type PaymentMethod string

const (
	_ PaymentMethod = ""

	CreditCard = "credit_card"
	DebitCard  = "debit_card"
)

type TmsUpdate struct {
	Version      string         `json:"version"`
	Acquirers    []*Acquirer    `json:"acquirers"`
	Applications []*Application `json:"applications"`
	PublicKeys   []*PublicKey   `json:"public_keys"`
	Terminal     *Terminal      `json:"terminal"`
	Merchant     *Merchant      `json:"merchant"`
	Checksum     string         `json:"checksum"`
}

type Merchant struct {
	ID       string        `json:"id"`
	Mcc      string        `json:"mcc"`
	Tcc      string        `json:"tcc"`
	Currency *CurrencyInfo `json:"currency"`
}

type CurrencyInfo struct {
	Code     int `json:"code"`
	Country  int `json:"country"`
	Exponent int `json:"exponent"`
}

type PublicKey struct {
	ID       string `json:"id"`
	RID      string `json:"rid"`
	Checksum []byte `json:"checksum"`
	Exponent []byte `json:"exponent"`
	Modulus  []byte `json:"modulus"`
}

type Application struct {
	ID                   string          `json:"id"`
	Name                 string          `json:"name"`
	AID                  string          `json:"aid"`
	AcquirerID           string          `json:"acquirer_id"`
	ActionCodes          *ActionCodeList `json:"action_codes"`
	AuthorizationEmvTags []int           `json:"authorization_emv_tags"`
	Brand                string          `json:"brand"`
	CapabilitySet        string          `json:"capability_set"`
	Ddol                 string          `json:"encoded_ddol"`
	Tdol                 string          `json:"encoded_tdol"`
	PaymentMethod        PaymentMethod   `json:"payment_method"`
	PublicKeyID          string          `json:"public_key_id"`
	Version              int             `json:"version"`
}

type ActionCodeList struct {
	Default string `json:"encoded_default"`
	Online  string `json:"encoded_online"`
	Denial  string `json:"encoded_denial"`
}

type Acquirer struct {
	ID                  string              `json:"id"`
	AbecsIndex          int                 `json:"abecs_index"`
	KeyIndex            int                 `json:"key_index"`
	PinEncryptionMethod PinEncryptionMethod `json:"pin_encryption_method"`
}

type Terminal struct {
	ID             string           `json:"id"`
	CapabilitySets []*CapabilitySet `json:"capability_sets"`
	RiskManagement *RiskManagement  `json:"risk_management"`
}

type CapabilitySet struct {
	Name                   string `json:"name"`
	Type                   int    `json:"encoded_type"`
	Capabilities           string `json:"encoded_capabilities"`
	AdditionalCapabilities string `json:"encoded_additional_capabilities"`
}

type RiskManagement struct {
	FloorLimit              int                       `json:"floor_limit"`
	RandomSelectionProfiles []*RandomSelectionProfile `json:"random_selection_profiles"`
}

type RandomSelectionProfile struct {
	Enabled                 bool   `json:"enabled"`
	ApplicationID           string `json:"application_id"`
	MaximumTargetPercentage int    `json:"maximum_target_percentage"`
	TargetPercentage        int    `json:"target_percentage"`
	ThresholdValue          int    `json:"threshold_value"`
}

type gambiarra struct {
	Status int `json:"status"`
}

//export gambiarraDaPorra
func gambiarraDaPorra() int {
	a := gambiarra{}

	if err := json.Unmarshal([]byte(`{"status":42}`), &a); err != nil {
		return 1000
	}

	return a.Status
}
