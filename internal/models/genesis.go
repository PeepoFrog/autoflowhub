package models

import "time"

type Genesis struct {
	AppHash         string          `json:"app_hash"`
	AppState        AppState        `json:"app_sate"`
	ChainId         string          `json:"chain_id"`
	ConsensusParams ConsensusParams `json:"consensus_params"`
	GenesisTime     time.Time       `json:"genesis_time"`
	InitialHeight   string          `json:"initial_height"`
}

type ConsensusParams struct {
	Block     BlockParams     `json:"block"`
	Evidence  EvidenceParams  `json:"evidence"`
	Validator ValidatorParams `json:"validator"`
	Version   string          `json:"version"`
}

type BlockParams struct {
	MaxBytes   uint64 `json:"max_bytes"`
	MaxGas     uint64 `json:"max_gas"`
	TimeIotaMs uint64 `json:"time_iota_ms"`
}

type EvidenceParams struct {
	MaxAgeNumBlocks uint64 `json:"max_age_num_blocks"`
	MaxAgeDuration  uint64 `json:"max_age_duration"`
	MaxBytes        uint64 `json:"max_bytes"`
}

type ValidatorParams struct {
	PubKeyTypes []string `json:"pub_key_types"`
}

type AppState struct {
	Auth           Auth           `json:"auth"`
	Bank           Bank           `json:"bank"`
	Basket         Basket         `json:"basket"`
	Collectives    Collectives    `json:"collectives"`
	Custody        Custody        `json:"custody"`
	CustomEvidence CustomEvidence `json:"customevidence"`
	CustomGov      CustomGov      `json:"customgov"`
	CustomStaking  CustomStaking  `json:"customstaking"`
	Distributor    Distributor    `json:"distributor"`
	Feeprocessing  FeeProcessing  `json:"feeprocessing"`
	Genutil        GenUtil        `json:"genutil"`
	Layer2         Layer2         `json:"layer2"`
	Multistatking  MultiStaking   `json:"multistaking"`
	Params         Params         `json:"params"`
	Recovery       Recovery       `json:"recovery"`
	Spending       Spending       `json:"spending"`
	Tokens         Tokens         `json:"tokens"`
	Ubi            UBI            `json:"ubi"`
	Upgrade        Upgrade        `json:"upgrade"`
}

type Auth struct {
	Params   AuthParams `json:"params"`
	Accounts []Account  `json:"accounts"`
}
type AuthParams struct {
	MaxMemoCharacters      uint64 `json:"max_memo_characters"`
	TxSigLimit             uint64 `json:"tx_sig_limit"`
	TxSizeCostPerByte      uint64 `json:"tx_size_cost_per_byte"`
	SigVerifyCostEd25519   uint64 `json:"sig_verify_cost_ed25519"`
	SigVerifyCostSecp256k1 uint64 `json:"sig_verify_cost_secp256k1"`
}

type Account struct {
	Type          string `json:"@type"`
	Address       string `json:"address"`
	PubKey        string `json:"pub_key"`
	AccountNumber uint64 `json:"account_number"`
	Sequence      uint64 `json:"sequence"`
}
type Bank struct {
	Params        BankParams    `json:"params"`
	Balances      []Balance     `json:"balances"`
	Supply        []interface{} `json:"supply"`         // Should be changed
	DenomMetadata []interface{} `json:"denom_metadata"` //Should be changed
}
type BankParams struct {
	SendEnabled        bool `json:"send_enabled"`
	DefaultSendEnabled bool `json:"default_send_enabled"`
}

type Balance struct {
	Address string `json:"address"`
	Coins   []Coin `json:"coins"`
}

type Coin struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type Basket struct {
	Baskets         []interface{} `json:"baskets"`
	HistoricalMints []interface{} `json:"historical_mints"`
	HistoricalBurns []interface{} `json:"historical_burns"`
	HistoricalSwaps []interface{} `json:"historical_swaps"`
}

type Collectives struct {
	Collectives  []interface{} `json:"collectives"`
	Contributers []interface{} `json:"contributers"`
}

type Custody struct {
}

type CustomEvidence struct {
	Evidence []interface{} `json:"evidence"`
}

type CustomGov struct {
	StartingProposalID          string                 `json:"starting_proposal_id"`
	NextRoleID                  string                 `json:"next_role_id"`
	Roles                       []CustomGovRole        `json:"roles"`
	RolePermissions             map[string]Permissions `json:"role_permissions"`
	NetworkActors               []NetworkActor         `json:"network_actors"`
	NetworkProperties           NetworkProperties      `json:"network_properties"`
	ExecutionFees               []ExecutionFee         `json:"execution_fees"`
	PoorNetworkMessages         PoorNetworkMessages    `json:"poor_network_messages"`
	Proposals                   []interface{}          `json:"proposals"`
	Votes                       []interface{}          `json:"votes"`
	DataRegistry                map[string]interface{} `json:"data_registry"`
	IdentityRecords             []IdentityRecord       `json:"identity_records"`
	LastIdentityRecordID        string                 `json:"last_identity_record_id"`
	IDRecordsVerifyRequests     []interface{}          `json:"id_records_verify_requests"`
	LastIDRecordVerifyRequestID string                 `json:"last_id_record_verify_request_id"`
	ProposalDurations           map[string]interface{} `json:"proposal_durations"`
}

type CustomGovRole struct {
	ID          int    `json:"id"`
	SID         string `json:"sid"`
	Description string `json:"description"`
}

type Permissions struct {
	Blacklist []int `json:"blacklist"`
	Whitelist []int `json:"whitelist"`
}

type NetworkActor struct {
	Address     string      `json:"address"`
	Roles       []string    `json:"roles"`
	Status      string      `json:"status"`
	Votes       []string    `json:"votes"`
	Permissions Permissions `json:"permissions"`
	Skin        string      `json:"skin"`
}

type NetworkProperties map[string]interface{}

type ExecutionFee struct {
	TransactionType   string `json:"transaction_type"`
	ExecutionFee      string `json:"execution_fee"`
	FailureFee        string `json:"failure_fee"`
	Timeout           string `json:"timeout"`
	DefaultParameters string `json:"default_parameters"`
}

type PoorNetworkMessages struct {
	Messages []string `json:"messages"`
}

type IdentityRecord struct {
	ID        string   `json:"id"`
	Address   string   `json:"address"`
	Key       string   `json:"key"`
	Value     string   `json:"value"`
	Date      string   `json:"date"`
	Verifiers []string `json:"verifiers"`
}
type CustomStaking struct {
	Validators []Validator `json:"validators"`
}

type Validator struct {
	ValKey string `json:"val_key"`
	PubKey PubKey `json:"pub_key"`
	Status string `json:"status"`
	Rank   string `json:"rank"`
	Streak string `json:"streak"`
}

type PubKey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}

type Distributor struct {
	FeesTreasury      []interface{} `json:"fees_treasury"`
	FeesCollected     []interface{} `json:"fees_collected"`
	SnapPeriod        string        `json:"snap_period"`
	ValidatorVotes    []interface{} `json:"validator_votes"`
	PreviousProposer  string        `json:"previous_proposer"`
	YearStartSnapshot YearSnapshot  `json:"year_start_snapshot"`
}

type YearSnapshot struct {
	SnapshotTime   string `json:"snapshot_time"`
	SnapshotAmount string `json:"snapshot_amount"`
}

type FeeProcessing struct {
}

type GenUtil struct {
	GenTxs []interface{} `json:"gen_txs"`
}

type Layer2 struct {
	Dapps  []interface{} `json:"dapps"`
	Bridge Bridge        `json:"bridge"`
}

type Bridge struct {
	Helper   interface{}   `json:"helper"`
	Accounts []interface{} `json:"accounts"`
	Tokens   []interface{} `json:"tokens"`
	Xams     []interface{} `json:"xams"`
}

type Params struct {
}

type Recovery struct {
	RecoveryRecords []interface{} `json:"recovery_records"`
	RecoveryTokens  []interface{} `json:"recovery_tokens"`
	Rewards         []interface{} `json:"rewards"`
	Rotations       []interface{} `json:"rotations"`
}

type Spending struct {
	Pools  []Pool        `json:"pools"`
	Claims []interface{} `json:"claims"`
}

type Pool struct {
	Name                    string        `json:"name"`
	ClaimStart              string        `json:"claim_start"`
	ClaimEnd                string        `json:"claim_end"`
	ClaimExpiry             string        `json:"claim_expiry"`
	Rates                   []Rate        `json:"rates"`
	VoteQuorum              string        `json:"vote_quorum"`
	VotePeriod              string        `json:"vote_period"`
	VoteEnactment           string        `json:"vote_enactment"`
	Owners                  Owners        `json:"owners"`
	Beneficiaries           Beneficiaries `json:"beneficiaries"`
	Balances                []interface{} `json:"balances"`
	DynamicRate             bool          `json:"dynamic_rate"`
	DynamicRatePeriod       string        `json:"dynamic_rate_period"`
	LastDynamicRateCalcTime string        `json:"last_dynamic_rate_calc_time"`
}

type Rate struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type Owners struct {
	OwnerRoles    []string      `json:"owner_roles"`
	OwnerAccounts []interface{} `json:"owner_accounts"`
}

type Beneficiaries struct {
	Roles    []Role        `json:"roles"`
	Accounts []interface{} `json:"accounts"`
}

type Role struct {
	Role   string `json:"role"`
	Weight string `json:"weight"`
}
type MultiStaking struct {
}

type Tokens struct {
	Aliases          []Alias         `json:"aliases"`
	Rates            []TokenRate     `json:"rates"`
	TokenBlackWhites TokenBlackWhite `json:"tokenBlackWhites"`
}

type Alias struct {
	Symbol      string   `json:"symbol"`
	Name        string   `json:"name"`
	Icon        string   `json:"icon"`
	Decimals    int      `json:"decimals"`
	Denoms      []string `json:"denoms"`
	Invalidated bool     `json:"invalidated"`
}

type TokenRate struct {
	Denom       string `json:"denom"`
	FeeRate     string `json:"fee_rate"`
	FeePayments bool   `json:"fee_payments"`
	StakeCap    string `json:"stake_cap"`
	StakeMin    string `json:"stake_min"`
	StakeToken  bool   `json:"stake_token"`
	Invalidated bool   `json:"invalidated"`
}

type TokenBlackWhite struct {
	Whitelisted []string `json:"whitelisted"`
	Blacklisted []string `json:"blacklisted"`
}

type UBI struct {
	UBIRecords []UBIRecord `json:"ubi_records"`
}

type UBIRecord struct {
	Name              string `json:"name"`
	DistributionStart string `json:"distribution_start"`
	DistributionEnd   string `json:"distribution_end"`
	DistributionLast  string `json:"distribution_last"`
	Amount            string `json:"amount"`
	Period            string `json:"period"`
	Pool              string `json:"pool"`
	Dynamic           bool   `json:"dynamic"`
}

type Upgrade struct {
	Version     string      `json:"version"`
	CurrentPlan interface{} `json:"current_plan"`
	NextPlan    interface{} `json:"next_plan"`
}
