// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package ethconfig

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool/blobpool"
	"github.com/ethereum/go-ethereum/core/txpool/legacypool"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/miner"
)

// MarshalTOML marshals as TOML.
func (c Config) MarshalTOML() (interface{}, error) {
	type Config struct {
		Genesis                                   *core.Genesis `toml:",omitempty"`
		NetworkId                                 uint64
		SyncMode                                  SyncMode
		EthDiscoveryURLs                          []string
		SnapDiscoveryURLs                         []string
		NoPruning                                 bool
		NoPrefetch                                bool
		TxLookupLimit                             uint64                 `toml:",omitempty"`
		TransactionHistory                        uint64                 `toml:",omitempty"`
		StateHistory                              uint64                 `toml:",omitempty"`
		StateScheme                               string                 `toml:",omitempty"`
		RequiredBlocks                            map[uint64]common.Hash `toml:"-"`
		SkipBcVersionCheck                        bool                   `toml:"-"`
		DatabaseHandles                           int                    `toml:"-"`
		DatabaseCache                             int
		DatabaseFreezer                           string
		TrieCleanCache                            int
		TrieDirtyCache                            int
		TrieTimeout                               time.Duration
		SnapshotCache                             int
		Preimages                                 bool
		FilterLogCacheSize                        int
		Miner                                     miner.Config
		TxPool                                    legacypool.Config
		BlobPool                                  blobpool.Config
		GPO                                       gasprice.Config
		EnablePreimageRecording                   bool
		VMTrace                                   string
		VMTraceJsonConfig                         string
		RPCGasCap                                 uint64
		RPCEVMTimeout                             time.Duration
		RPCTxFeeCap                               float64
		OverrideCancun                            *uint64 `toml:",omitempty"`
		OverrideVerkle                            *uint64 `toml:",omitempty"`
		OverrideOptimismCanyon                    *uint64 `toml:",omitempty"`
		OverrideOptimismEcotone                   *uint64 `toml:",omitempty"`
		OverrideOptimismFjord                     *uint64 `toml:",omitempty"`
		OverrideOptimismGranite                   *uint64 `toml:",omitempty"`
		OverrideOptimismHolocene                  *uint64 `toml:",omitempty"`
		OverrideOptimismInterop                   *uint64 `toml:",omitempty"`
		OverrideOptimismIsthmus                   *uint64 `toml:",omitempty"`
		ApplySuperchainUpgrades                   bool    `toml:",omitempty"`
		RollupSequencerHTTP                       string
		RollupSequencerTxConditionalEnabled       bool
		RollupSequencerTxConditionalCostRateLimit int
		RollupHistoricalRPC                       string
		RollupHistoricalRPCTimeout                time.Duration
		RollupDisableTxPoolGossip                 bool
		RollupDisableTxPoolAdmission              bool
		RollupHaltOnIncompatibleProtocolVersion   string
		InteropMessageRPC                         string `toml:",omitempty"`
		InteropMempoolFiltering                   bool   `toml:",omitempty"`
	}
	var enc Config
	enc.Genesis = c.Genesis
	enc.NetworkId = c.NetworkId
	enc.SyncMode = c.SyncMode
	enc.EthDiscoveryURLs = c.EthDiscoveryURLs
	enc.SnapDiscoveryURLs = c.SnapDiscoveryURLs
	enc.NoPruning = c.NoPruning
	enc.NoPrefetch = c.NoPrefetch
	enc.TxLookupLimit = c.TxLookupLimit
	enc.TransactionHistory = c.TransactionHistory
	enc.StateHistory = c.StateHistory
	enc.StateScheme = c.StateScheme
	enc.RequiredBlocks = c.RequiredBlocks
	enc.SkipBcVersionCheck = c.SkipBcVersionCheck
	enc.DatabaseHandles = c.DatabaseHandles
	enc.DatabaseCache = c.DatabaseCache
	enc.DatabaseFreezer = c.DatabaseFreezer
	enc.TrieCleanCache = c.TrieCleanCache
	enc.TrieDirtyCache = c.TrieDirtyCache
	enc.TrieTimeout = c.TrieTimeout
	enc.SnapshotCache = c.SnapshotCache
	enc.Preimages = c.Preimages
	enc.FilterLogCacheSize = c.FilterLogCacheSize
	enc.Miner = c.Miner
	enc.TxPool = c.TxPool
	enc.BlobPool = c.BlobPool
	enc.GPO = c.GPO
	enc.EnablePreimageRecording = c.EnablePreimageRecording
	enc.VMTrace = c.VMTrace
	enc.VMTraceJsonConfig = c.VMTraceJsonConfig
	enc.RPCGasCap = c.RPCGasCap
	enc.RPCEVMTimeout = c.RPCEVMTimeout
	enc.RPCTxFeeCap = c.RPCTxFeeCap
	enc.OverrideCancun = c.OverrideCancun
	enc.OverrideVerkle = c.OverrideVerkle
	enc.OverrideOptimismCanyon = c.OverrideOptimismCanyon
	enc.OverrideOptimismEcotone = c.OverrideOptimismEcotone
	enc.OverrideOptimismFjord = c.OverrideOptimismFjord
	enc.OverrideOptimismGranite = c.OverrideOptimismGranite
	enc.OverrideOptimismHolocene = c.OverrideOptimismHolocene
	enc.OverrideOptimismInterop = c.OverrideOptimismInterop
	enc.OverrideOptimismIsthmus = c.OverrideOptimismIsthmus
	enc.ApplySuperchainUpgrades = c.ApplySuperchainUpgrades
	enc.RollupSequencerHTTP = c.RollupSequencerHTTP
	enc.RollupSequencerTxConditionalEnabled = c.RollupSequencerTxConditionalEnabled
	enc.RollupSequencerTxConditionalCostRateLimit = c.RollupSequencerTxConditionalCostRateLimit
	enc.RollupHistoricalRPC = c.RollupHistoricalRPC
	enc.RollupHistoricalRPCTimeout = c.RollupHistoricalRPCTimeout
	enc.RollupDisableTxPoolGossip = c.RollupDisableTxPoolGossip
	enc.RollupDisableTxPoolAdmission = c.RollupDisableTxPoolAdmission
	enc.RollupHaltOnIncompatibleProtocolVersion = c.RollupHaltOnIncompatibleProtocolVersion
	enc.InteropMessageRPC = c.InteropMessageRPC
	enc.InteropMempoolFiltering = c.InteropMempoolFiltering
	return &enc, nil
}

// UnmarshalTOML unmarshals from TOML.
func (c *Config) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Config struct {
		Genesis                                   *core.Genesis `toml:",omitempty"`
		NetworkId                                 *uint64
		SyncMode                                  *SyncMode
		EthDiscoveryURLs                          []string
		SnapDiscoveryURLs                         []string
		NoPruning                                 *bool
		NoPrefetch                                *bool
		TxLookupLimit                             *uint64                `toml:",omitempty"`
		TransactionHistory                        *uint64                `toml:",omitempty"`
		StateHistory                              *uint64                `toml:",omitempty"`
		StateScheme                               *string                `toml:",omitempty"`
		RequiredBlocks                            map[uint64]common.Hash `toml:"-"`
		SkipBcVersionCheck                        *bool                  `toml:"-"`
		DatabaseHandles                           *int                   `toml:"-"`
		DatabaseCache                             *int
		DatabaseFreezer                           *string
		TrieCleanCache                            *int
		TrieDirtyCache                            *int
		TrieTimeout                               *time.Duration
		SnapshotCache                             *int
		Preimages                                 *bool
		FilterLogCacheSize                        *int
		Miner                                     *miner.Config
		TxPool                                    *legacypool.Config
		BlobPool                                  *blobpool.Config
		GPO                                       *gasprice.Config
		EnablePreimageRecording                   *bool
		VMTrace                                   *string
		VMTraceJsonConfig                         *string
		RPCGasCap                                 *uint64
		RPCEVMTimeout                             *time.Duration
		RPCTxFeeCap                               *float64
		OverrideCancun                            *uint64 `toml:",omitempty"`
		OverrideVerkle                            *uint64 `toml:",omitempty"`
		OverrideOptimismCanyon                    *uint64 `toml:",omitempty"`
		OverrideOptimismEcotone                   *uint64 `toml:",omitempty"`
		OverrideOptimismFjord                     *uint64 `toml:",omitempty"`
		OverrideOptimismGranite                   *uint64 `toml:",omitempty"`
		OverrideOptimismHolocene                  *uint64 `toml:",omitempty"`
		OverrideOptimismInterop                   *uint64 `toml:",omitempty"`
		OverrideOptimismIsthmus                   *uint64 `toml:",omitempty"`
		ApplySuperchainUpgrades                   *bool   `toml:",omitempty"`
		RollupSequencerHTTP                       *string
		RollupSequencerTxConditionalEnabled       *bool
		RollupSequencerTxConditionalCostRateLimit *int
		RollupHistoricalRPC                       *string
		RollupHistoricalRPCTimeout                *time.Duration
		RollupDisableTxPoolGossip                 *bool
		RollupDisableTxPoolAdmission              *bool
		RollupHaltOnIncompatibleProtocolVersion   *string
		InteropMessageRPC                         *string `toml:",omitempty"`
		InteropMempoolFiltering                   *bool   `toml:",omitempty"`
	}
	var dec Config
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Genesis != nil {
		c.Genesis = dec.Genesis
	}
	if dec.NetworkId != nil {
		c.NetworkId = *dec.NetworkId
	}
	if dec.SyncMode != nil {
		c.SyncMode = *dec.SyncMode
	}
	if dec.EthDiscoveryURLs != nil {
		c.EthDiscoveryURLs = dec.EthDiscoveryURLs
	}
	if dec.SnapDiscoveryURLs != nil {
		c.SnapDiscoveryURLs = dec.SnapDiscoveryURLs
	}
	if dec.NoPruning != nil {
		c.NoPruning = *dec.NoPruning
	}
	if dec.NoPrefetch != nil {
		c.NoPrefetch = *dec.NoPrefetch
	}
	if dec.TxLookupLimit != nil {
		c.TxLookupLimit = *dec.TxLookupLimit
	}
	if dec.TransactionHistory != nil {
		c.TransactionHistory = *dec.TransactionHistory
	}
	if dec.StateHistory != nil {
		c.StateHistory = *dec.StateHistory
	}
	if dec.StateScheme != nil {
		c.StateScheme = *dec.StateScheme
	}
	if dec.RequiredBlocks != nil {
		c.RequiredBlocks = dec.RequiredBlocks
	}
	if dec.SkipBcVersionCheck != nil {
		c.SkipBcVersionCheck = *dec.SkipBcVersionCheck
	}
	if dec.DatabaseHandles != nil {
		c.DatabaseHandles = *dec.DatabaseHandles
	}
	if dec.DatabaseCache != nil {
		c.DatabaseCache = *dec.DatabaseCache
	}
	if dec.DatabaseFreezer != nil {
		c.DatabaseFreezer = *dec.DatabaseFreezer
	}
	if dec.TrieCleanCache != nil {
		c.TrieCleanCache = *dec.TrieCleanCache
	}
	if dec.TrieDirtyCache != nil {
		c.TrieDirtyCache = *dec.TrieDirtyCache
	}
	if dec.TrieTimeout != nil {
		c.TrieTimeout = *dec.TrieTimeout
	}
	if dec.SnapshotCache != nil {
		c.SnapshotCache = *dec.SnapshotCache
	}
	if dec.Preimages != nil {
		c.Preimages = *dec.Preimages
	}
	if dec.FilterLogCacheSize != nil {
		c.FilterLogCacheSize = *dec.FilterLogCacheSize
	}
	if dec.Miner != nil {
		c.Miner = *dec.Miner
	}
	if dec.TxPool != nil {
		c.TxPool = *dec.TxPool
	}
	if dec.BlobPool != nil {
		c.BlobPool = *dec.BlobPool
	}
	if dec.GPO != nil {
		c.GPO = *dec.GPO
	}
	if dec.EnablePreimageRecording != nil {
		c.EnablePreimageRecording = *dec.EnablePreimageRecording
	}
	if dec.VMTrace != nil {
		c.VMTrace = *dec.VMTrace
	}
	if dec.VMTraceJsonConfig != nil {
		c.VMTraceJsonConfig = *dec.VMTraceJsonConfig
	}
	if dec.RPCGasCap != nil {
		c.RPCGasCap = *dec.RPCGasCap
	}
	if dec.RPCEVMTimeout != nil {
		c.RPCEVMTimeout = *dec.RPCEVMTimeout
	}
	if dec.RPCTxFeeCap != nil {
		c.RPCTxFeeCap = *dec.RPCTxFeeCap
	}
	if dec.OverrideCancun != nil {
		c.OverrideCancun = dec.OverrideCancun
	}
	if dec.OverrideVerkle != nil {
		c.OverrideVerkle = dec.OverrideVerkle
	}
	if dec.OverrideOptimismCanyon != nil {
		c.OverrideOptimismCanyon = dec.OverrideOptimismCanyon
	}
	if dec.OverrideOptimismEcotone != nil {
		c.OverrideOptimismEcotone = dec.OverrideOptimismEcotone
	}
	if dec.OverrideOptimismFjord != nil {
		c.OverrideOptimismFjord = dec.OverrideOptimismFjord
	}
	if dec.OverrideOptimismGranite != nil {
		c.OverrideOptimismGranite = dec.OverrideOptimismGranite
	}
	if dec.OverrideOptimismHolocene != nil {
		c.OverrideOptimismHolocene = dec.OverrideOptimismHolocene
	}
	if dec.OverrideOptimismInterop != nil {
		c.OverrideOptimismInterop = dec.OverrideOptimismInterop
	}
	if dec.OverrideOptimismIsthmus != nil {
		c.OverrideOptimismIsthmus = dec.OverrideOptimismIsthmus
	}
	if dec.ApplySuperchainUpgrades != nil {
		c.ApplySuperchainUpgrades = *dec.ApplySuperchainUpgrades
	}
	if dec.RollupSequencerHTTP != nil {
		c.RollupSequencerHTTP = *dec.RollupSequencerHTTP
	}
	if dec.RollupSequencerTxConditionalEnabled != nil {
		c.RollupSequencerTxConditionalEnabled = *dec.RollupSequencerTxConditionalEnabled
	}
	if dec.RollupSequencerTxConditionalCostRateLimit != nil {
		c.RollupSequencerTxConditionalCostRateLimit = *dec.RollupSequencerTxConditionalCostRateLimit
	}
	if dec.RollupHistoricalRPC != nil {
		c.RollupHistoricalRPC = *dec.RollupHistoricalRPC
	}
	if dec.RollupHistoricalRPCTimeout != nil {
		c.RollupHistoricalRPCTimeout = *dec.RollupHistoricalRPCTimeout
	}
	if dec.RollupDisableTxPoolGossip != nil {
		c.RollupDisableTxPoolGossip = *dec.RollupDisableTxPoolGossip
	}
	if dec.RollupDisableTxPoolAdmission != nil {
		c.RollupDisableTxPoolAdmission = *dec.RollupDisableTxPoolAdmission
	}
	if dec.RollupHaltOnIncompatibleProtocolVersion != nil {
		c.RollupHaltOnIncompatibleProtocolVersion = *dec.RollupHaltOnIncompatibleProtocolVersion
	}
	if dec.InteropMessageRPC != nil {
		c.InteropMessageRPC = *dec.InteropMessageRPC
	}
	if dec.InteropMempoolFiltering != nil {
		c.InteropMempoolFiltering = *dec.InteropMempoolFiltering
	}
	return nil
}
