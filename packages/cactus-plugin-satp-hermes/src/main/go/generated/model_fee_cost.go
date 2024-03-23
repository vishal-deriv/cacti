/*
SATP Gateway Client (Business Logic Orchestrator)

SATP is a protocol operating between two gateways that conducts the transfer of a digital asset from one gateway to another. The protocol establishes a secure channel between the endpoints and implements a 2-phase commit to ensure the properties of transfer atomicity, consistency, isolation and durability.  This API defines the gateway client facing API (business logic orchestrator, or BLO), which is named API-Type 1 in the SATP-Core specification.  **Additional Resources**: - [Proposed SATP Charter](https://datatracker.ietf.org/doc/charter-ietf-satp/) - [SATP Core draft](https://datatracker.ietf.org/doc/draft-ietf-satp-core) - [SATP Crash Recovery draft](https://datatracker.ietf.org/doc/draft-belchior-satp-gateway-recovery/) - [SATP Architecture draft](https://datatracker.ietf.org/doc/draft-ietf-satp-architecture/) - [SATP Use-Cases draft](https://datatracker.ietf.org/doc/draft-ramakrishna-sat-use-cases/) - [SATP Data sharing draft](https://datatracker.ietf.org/doc/draft-ramakrishna-satp-data-sharing) - [SATP View Addresses draft](https://datatracker.ietf.org/doc/draft-ramakrishna-satp-views-addresses)

API version: 0.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package generated

import (
	"encoding/json"
)

// checks if the FeeCost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeeCost{}

// FeeCost Details about a specific fee cost associated with the transaction.
type FeeCost struct {
	// Name of the fee cost.
	Name *string `json:"name,omitempty"`
	// The amount in string format including all decimals.
	Amount *string `json:"amount,omitempty"`
	// The amount in string format including all decimals.
	AmountUSD *string `json:"amountUSD,omitempty"`
	// The symbol of a token
	Token *string `json:"token,omitempty"`
	// Indicates if the fee is included in the transaction amount.
	Included *bool `json:"included,omitempty"`
}

// NewFeeCost instantiates a new FeeCost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeeCost() *FeeCost {
	this := FeeCost{}
	return &this
}

// NewFeeCostWithDefaults instantiates a new FeeCost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeCostWithDefaults() *FeeCost {
	this := FeeCost{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FeeCost) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeCost) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FeeCost) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FeeCost) SetName(v string) {
	o.Name = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *FeeCost) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeCost) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *FeeCost) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *FeeCost) SetAmount(v string) {
	o.Amount = &v
}

// GetAmountUSD returns the AmountUSD field value if set, zero value otherwise.
func (o *FeeCost) GetAmountUSD() string {
	if o == nil || IsNil(o.AmountUSD) {
		var ret string
		return ret
	}
	return *o.AmountUSD
}

// GetAmountUSDOk returns a tuple with the AmountUSD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeCost) GetAmountUSDOk() (*string, bool) {
	if o == nil || IsNil(o.AmountUSD) {
		return nil, false
	}
	return o.AmountUSD, true
}

// HasAmountUSD returns a boolean if a field has been set.
func (o *FeeCost) HasAmountUSD() bool {
	if o != nil && !IsNil(o.AmountUSD) {
		return true
	}

	return false
}

// SetAmountUSD gets a reference to the given string and assigns it to the AmountUSD field.
func (o *FeeCost) SetAmountUSD(v string) {
	o.AmountUSD = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *FeeCost) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeCost) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *FeeCost) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *FeeCost) SetToken(v string) {
	o.Token = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *FeeCost) GetIncluded() bool {
	if o == nil || IsNil(o.Included) {
		var ret bool
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeCost) GetIncludedOk() (*bool, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *FeeCost) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given bool and assigns it to the Included field.
func (o *FeeCost) SetIncluded(v bool) {
	o.Included = &v
}

func (o FeeCost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeeCost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.AmountUSD) {
		toSerialize["amountUSD"] = o.AmountUSD
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

type NullableFeeCost struct {
	value *FeeCost
	isSet bool
}

func (v NullableFeeCost) Get() *FeeCost {
	return v.value
}

func (v *NullableFeeCost) Set(val *FeeCost) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeCost) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeCost(val *FeeCost) *NullableFeeCost {
	return &NullableFeeCost{value: val, isSet: true}
}

func (v NullableFeeCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeeCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


