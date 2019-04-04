package walletrpc

// TransferRequest is the request body of the Transfer client rpc call.
type TransferRequest struct {
	// Destinations - array of destinations to receive XMR:
	Destinations []Destination `json:"destinations"`
	// AccountIndex - return transfers for this account.
	AccountIndex uint64 `json:"account_index"`
	// SubaddrIndices - return transfers sent to these subaddresses
	SubaddrIndices []uint64 `json:"subaddr_indices,omitempty"`
	// Fee - unsigned int; Ignored, will be automatically calculated.
	Fee uint64 `json:"fee,omitempty"`
	// Mixin - unsigned int; Number of outpouts from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin,omitempty"`
	// RingSize - unsigned int; Number of outputs to mix in the transaction (this output + N decoys from the blockchain).
	RingSize uint64 `json:"ring_size,omitempty"`
	// unlock_time - unsigned int; Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`
	// payment_id - string; (Optional) Random 32-byte/64-character hex string to identify a transaction.
	PaymentID string `json:"payment_id,omitempty"`
	// get_tx_key - boolean; (Optional) Return the transaction key after sending.
	GetTxKey bool `json:"get_tx_key"`
	// priority - unsigned int; Set a priority for the transaction.
	// Accepted Values are: 0-3 for: default, unimportant, normal, elevated, priority.
	Priority Priority `json:"priority"`
	// do_not_relay - boolean; (Optional) If true, the newly created transaction will not be relayed to the monero network. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay,omitempty"`
	// get_tx_hex - boolean; Return the transaction as hex string after sending
	GetTxHex bool `json:"get_tx_hex,omitempty"`
}

// Destination to receive XMR
type Destination struct {
	// Amount - unsigned int; Amount to send to each destination, in atomic units.
	Amount uint64 `json:"amount"`
	// Address - string; Destination public address.
	Address string `json:"address"`
}

// TransferResponse is the successful output of a Client.Transfer()
type TransferResponse struct {
	// fee - Integer value of the fee charged for the txn.
	Fee uint64 `json:"fee"`
	// tx_hash - String for the publically searchable transaction hash
	TxHash string `json:"tx_hash"`
	// tx_key - String for the transaction key if get_tx_key is true, otherwise, blank string.
	TxKey string `json:"tx_key,omitempty"`
	// tx_blob - Transaction as hex string if get_tx_hex is true
	TxBlob string `json:"tx_blob,omitempty"`
}

// TransferSplitResponse is the successful output of a Client.TransferSplit()
type TransferSplitResponse struct {
	// fee_list - array of: integer. The amount of fees paid for every transaction.
	FeeList []uint64 `json:"fee_list"`
	// tx_hash_list - array of: string. The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`
	// tx_blob_list - array of: string. The tx as hex string for every transaction.
	TxBlobList []string `json:"tx_blob_list"`
	// amount_list - array of: integer. The amount transferred for every transaction..
	AmountList []uint64 `json:"amount_list"`
	// tx_key_list - array of: string. The transaction keys for every transaction.
	TxKeyList []string `json:"tx_key_list"`
}

// SweepAllRequest is the struct to send all unlocked balance to an address.
type SweepAllRequest struct {
	// address - string; Destination public address.
	Address string `json:"address"`
	// account_index - unsigned int; Sweep transactions from this account.
	AccountIndex uint64 `json:"account_index"`
	// priority - unsigned int; (Optional)
	Priority Priority `json:"priority,omitempty"`
	// mixin - unsigned int; Number of outpouts from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin"`
	// unlock_time - unsigned int; Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`
	// payment_id - string; (Optional) Random 32-byte/64-character hex string to identify a transaction.
	PaymentID string `json:"payment_id,omitempty"`
	// get_tx_keys - boolean; (Optional) Return the transaction keys after sending.
	GetTxKeys bool `json:"get_tx_keys,omitempty"`
	// below_amount - unsigned int; (Optional)
	BelowAmount uint64 `json:"below_amount"`
	// do_not_relay - boolean; (Optional)
	DoNotRelay bool `json:"do_not_relay,omitempty"`
	// get_tx_hex - boolean; (Optional) return the transactions as hex encoded string.
	GetTxHex bool `json:"get_tx_hex,omitempty"`
}

// SweepAllResponse is a typical response of a SweepAllRequest
type SweepAllResponse struct {
	// tx_hash_list - array of: string. The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`
	// tx_blob_list - array of: string. The tx as hex string for every transaction.
	TxBlobList []string `json:"tx_blob_list"`
	// tx_key_list - array of: string. The transaction keys for every transaction.
	TxKeyList []string `json:"tx_key_list"`
}

// SweepSingleRequest is the struct to send single unlocked balance to an address.
type SweepSingleRequest struct {
	// address - string; Destination public address.
	Address string `json:"address"`
	// account_index - unsigned int; Sweep transactions from this account.
	AccountIndex uint64 `json:"account_index"`
	// priority - unsigned int; (Optional)
	Priority Priority `json:"priority,omitempty"`
	// mixin - unsigned int; Number of outpouts from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin,omitempty"`
	// RingSize - unsigned int; Number of outputs to mix in the transaction (this output + N decoys from the blockchain).
	RingSize uint64 `json:"ring_size,omitempty"`
	// unlock_time - unsigned int; Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`
	// payment_id - string; (Optional) Random 32-byte/64-character hex string to identify a transaction.
	PaymentID string `json:"payment_id,omitempty"`
	// get_tx_keys - boolean; (Optional) Return the transaction keys after sending.
	GetTxKeys bool `json:"get_tx_keys,omitempty"`
	// get_tx_metadata - boolean; (Optional) return the transaction metadata as a string. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata,omitempty"`
	//  string; Key image of specific output to sweep.
	KeyImage string `json:"key_image"`
	// below_amount - unsigned int; (Optional)
	BelowAmount uint64 `json:"below_amount"`
	// do_not_relay - boolean; (Optional)
	DoNotRelay bool `json:"do_not_relay,omitempty"`
	// get_tx_hex - boolean; (Optional) return the transactions as hex encoded string.
	GetTxHex bool `json:"get_tx_hex,omitempty"`
}

// SweepSingleResponse is a typical response of a SweepSingleRequest
type SweepSingleResponse struct {
	// tx_hash_list - string. The tx hashes of every transaction.
	TxHash string `json:"tx_hash"`
	// tx_key_list - string. The transaction keys for every transaction.
	TxKey string `json:"tx_key"`
	// amount - integer. The amount transferred for every transaction.
	Amount uint64 `json:"amount"`
	// fee - integer. The amount of fees paid for every transaction.
	Fee uint64 `json:"fee"`
	// tx_blob - string. The tx as hex string for every transaction.
	TxBlob string `json:"tx_blob"`
	// tx_metadata - string. List of transaction metadata needed to relay the transactions later.
	TxMetadata string `json:"tx_metadata"`
	// multisig_txset - string. The set of signing keys used in a multisig transaction (empty for non-multisig).
	Multisig_Txset string `json:"multisig_txset"`
	// unsigned_txset - string. Set of unsigned tx for cold-signing purposes.
	Unsigned_Txset string `json:"unsigned_txset"`
}

// Payment ...
type Payment struct {
	PaymentID   string `json:"payment_id"`
	TxHash      string `json:"tx_hash"`
	Amount      uint64 `json:"amount"`
	BlockHeight uint64 `json:"block_height"`
	UnlockTime  uint64 `json:"unlock_time"`
}

// GetTransfersRequest = GetTransfers body
type GetTransfersRequest struct {
	In             bool   `json:"in"`
	Out            bool   `json:"out"`
	Pending        bool   `json:"pending"`
	Failed         bool   `json:"failed"`
	Pool           bool   `json:"pool"`
	FilterByHeight bool   `json:"filter_by_height"`
	MinHeight      uint64 `json:"min_height"`
	MaxHeight      uint64 `json:"max_height"`
}

// GetTransfersResponse = GetTransfers output
type GetTransfersResponse struct {
	In      []Transfer `json:"in"`
	Out     []Transfer `json:"out"`
	Pending []Transfer `json:"pending"`
	Failed  []Transfer `json:"failed"`
	Pool    []Transfer `json:"pool"`
}

// Transfer is the transfer data of
type Transfer struct {
	TxID         string        `json:"txid"`
	PaymentID    string        `json:"payment_id"`
	Height       uint64        `json:"height"`
	Timestamp    uint64        `json:"timestamp"`
	Amount       uint64        `json:"amount"`
	Fee          uint64        `json:"fee"`
	Note         string        `json:"note"`
	Destinations []Destination `json:"destinations,omitempty"` // TODO: check if deprecated
	Type         string        `json:"type"`
}

// IncTransfer is returned by IncomingTransfers
type IncTransfer struct {
	Amount uint64 `json:"amount"`
	// Mostly internal use, can be ignored by most users.
	GlobalIndex uint64 `json:"global_index"`
	//key_image - string; Key image for the incoming transfer's unspent output (empty unless verbose is true).
	KeyImage string `json:"key_image"`
	// spent - boolean; Indicates if this transfer has been spent.
	Spent bool `json:"spent"`
	// subaddr_index - unsigned int; Subaddress index for incoming transfer.
	SubAddressIndex struct {
		Major int64 `json:"major"`
		Minor int64 `json:"minor"`
	} `json:"subaddr_index"`
	// Several incoming transfers may share the same hash
	// if they were in the same transaction.
	TxHash string `json:"tx_hash"`
	// tx_size - unsigned int; Size of transaction in bytes.
	TxSize uint64 `json:"tx_size"`
}

// URIDef is the skeleton of the MakeURI()
type URIDef struct {
	// address - wallet address string
	Address string `json:"address"`
	// amount (optional) - the integer amount to receive, in atomic units
	Amount uint64 `json:"amount,omitempty"`
	// payment_id (optional) - 16 or 64 character hexadecimal payment id string
	PaymentID string `json:"payment_id,omitempty"`
	// recipient_name (optional) - string name of the payment recipient
	RecipientName string `json:"recipient_name,omitempty"`
	// tx_description (optional) - string describing the reason for the tx
	TxDescription string `json:"tx_description,omitempty"`
}

// SignedKeyImage - The key image is an alternate public key computed on a second base point,
// Hp(P), instead of G. It is required in traceable ring signature construction to ensure
// multiple signatures with the same real key are linked.
type SignedKeyImage struct {
	KeyImage  string `json:"key_image"`
	Signature string `json:"signature"`
}

// ImportKeyImageResponse is the result of ImportKeyImages()
type ImportKeyImageResponse struct {
	Height  uint64 `json:"height"`
	Spent   uint64 `json:"spent"`
	Unspent uint64 `json:"unspent"`
}

// AddressBookEntry Address
type AddressBookEntry struct {
	Address     string `json:"address"`
	Description string `json:"description,omitempty"`
	Index       uint64 `json:"index,omitempty"`
	PaymentID   string `json:"payment_id,omitempty"`
}

// GetAccountsResponse Accounts
type GetAccountsResponse struct {
	SubaddressAccounts   []SubaddressAccount `json:"subaddress_accounts"`
	TotalBalance         uint64              `json:"total_balance"`
	TotalUnlockedBalance uint64              `json:"total_unlocked_balance"`
}

// SubaddressAccount Subaddress accounts
type SubaddressAccount struct {
	AccountIndex    uint64 `json:"account_index"`
	Balance         uint64 `json:"balance"`
	BaseAddress     string `json:"base_address"`
	Label           string `json:"label"`
	Tag             string `json:"tag"`
	UnlockedBalance uint64 `json:"unlocked_balance"`
}

// CreateAccountResponse New account
type CreateAccountResponse struct {
	AccountIndex uint64 `json:"account_index"`
	Address      string `json:"address"`
}
