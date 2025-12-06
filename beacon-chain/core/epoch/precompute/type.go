package precompute

import "github.com/OffchainLabs/prysm/v6/consensus-types/primitives"

// Validator stores the pre computation of individual validator's attesting records these records
// consist of attestation votes, block inclusion record. Pre computing and storing such record
// is essential for process epoch optimizations.
// 设计目的：在纪元处理开始前批量预计算验证者状态，避免在纪元处理过程中重复计算，显著提升性能。
// Precomputes validator states before epoch processing to avoid redundant calculations during processing, significantly improving performance.
type Validator struct {
	// IsSlashed is true if the validator has been slashed.
	// IsSlashed 表示验证者是否被罚没
	IsSlashed bool
	// IsWithdrawableCurrentEpoch is true if the validator can withdraw current epoch.
	// IsWithdrawableCurrentEpoch 表示验证者当前纪元是否可以提取资金
	IsWithdrawableCurrentEpoch bool
	// IsActiveCurrentEpoch is true if the validator was active current epoch.
	// IsActiveCurrentEpoch 表示验证者当前纪元是否活跃
	IsActiveCurrentEpoch bool
	// IsActivePrevEpoch is true if the validator was active prev epoch.
	// IsActivePrevEpoch 表示验证者上个纪元是否活跃
	IsActivePrevEpoch bool
	// IsCurrentEpochAttester is true if the validator attested current epoch.
	// IsCurrentEpochAttester 表示验证者是否在当前纪元提交了证明
	IsCurrentEpochAttester bool
	// IsCurrentEpochTargetAttester is true if the validator attested current epoch target.
	// IsCurrentEpochTargetAttester 表示验证者是否在当前纪元正确投票给目标检查点
	IsCurrentEpochTargetAttester bool
	// IsPrevEpochAttester is true if the validator attested previous epoch.
	// IsPrevEpochAttester 表示验证者是否在上个纪元提交了证明
	IsPrevEpochAttester bool
	// IsPrevEpochSourceAttester is true if the validator attested to source previous epoch. [Only for Altair]
	// IsPrevEpochSourceAttester 表示验证者是否在上个纪元正确投票给源检查点 [仅Altair及以后]
	IsPrevEpochSourceAttester bool
	// IsPrevEpochTargetAttester is true if the validator attested previous epoch target.
	// IsPrevEpochTargetAttester 表示验证者是否在上个纪元正确投票给目标检查点
	IsPrevEpochTargetAttester bool
	// IsPrevEpochHeadAttester is true if the validator attested the previous epoch head.
	// IsPrevEpochHeadAttester 表示验证者是否在上个纪元正确投票给区块头
	IsPrevEpochHeadAttester bool
	//三种证明类型的作用：
	//源证明 (Source)：证明链的正确来源，确保最终性
	//目标证明 (Target)：证明正确的目标检查点
	//头证明 (Head)：证明正确的区块头，确保链的活性

	// 余额和包含信息字段 (Balance and Inclusion Information)
	// CurrentEpochEffectiveBalance is how much effective balance this validator has current epoch.
	// CurrentEpochEffectiveBalance 表示验证者当前纪元的有效余额
	//重要概念：
	//有效余额 ≠ 实际余额：有效余额是向下取整到1 ETH的倍数
	//最大32 ETH：无论实际质押多少，用于奖励计算的最大有效余额为32 ETH
	//迟滞更新：只有当实际余额变化超过阈值时才更新有效余额
	CurrentEpochEffectiveBalance uint64
	// InclusionSlot is the slot of when the attestation gets included in the chain.
	InclusionSlot primitives.Slot
	// InclusionDistance is the distance between the assigned slot and this validator's attestation was included in block.
	InclusionDistance primitives.Slot
	// ProposerIndex is the index of proposer at slot where this validator's attestation was included.
	ProposerIndex primitives.ValidatorIndex
	// BeforeEpochTransitionBalance is the validator balance prior to epoch transition.
	BeforeEpochTransitionBalance uint64
	// AfterEpochTransitionBalance is the validator balance after epoch transition.
	AfterEpochTransitionBalance uint64

	// InactivityScore of the validator. [New in Altair]
	InactivityScore uint64
}

// Balance stores the pre computation of the total participated balances for a given epoch
// Pre computing and storing such record is essential for process epoch optimizations.
type Balance struct {
	// ActiveCurrentEpoch is the total effective balance of all active validators during current epoch.
	// 当前纪元活跃总余额 (Total active balance in current epoch)
	ActiveCurrentEpoch uint64
	// ActivePrevEpoch is the total effective balance of all active validators during prev epoch.
	// 上个纪元源证明总余额 (Total source attestation balance in previous epoch)
	ActivePrevEpoch uint64
	// CurrentEpochAttested is the total effective balance of all validators who attested during current epoch.
	CurrentEpochAttested uint64
	// CurrentEpochTargetAttested is the total effective balance of all validators who attested
	// for epoch boundary block during current epoch.
	CurrentEpochTargetAttested uint64
	// PrevEpochAttested is the total effective balance of all validators who attested during prev epoch.
	PrevEpochAttested uint64
	// PrevEpochTargetAttested is the total effective balance of all validators who attested
	// for epoch boundary block during prev epoch.
	PrevEpochTargetAttested uint64
	// PrevEpochHeadAttested is the total effective balance of all validators who attested
	// correctly for head block during prev epoch.
	PrevEpochHeadAttested uint64
}
