// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package someplace

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitMarket is the `initMarket` instruction.
type InitMarket struct {
	MarketUid *ag_solanago.PublicKey
	Name      *string

	// [0] = [WRITE] marketAuthority
	//
	// [1] = [WRITE] marketMint
	//
	// [2] = [WRITE, SIGNER] oracle
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitMarketInstructionBuilder creates a new `InitMarket` instruction builder.
func NewInitMarketInstructionBuilder() *InitMarket {
	nd := &InitMarket{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetMarketUid sets the "marketUid" parameter.
func (inst *InitMarket) SetMarketUid(marketUid ag_solanago.PublicKey) *InitMarket {
	inst.MarketUid = &marketUid
	return inst
}

// SetName sets the "name" parameter.
func (inst *InitMarket) SetName(name string) *InitMarket {
	inst.Name = &name
	return inst
}

// SetMarketAuthorityAccount sets the "marketAuthority" account.
func (inst *InitMarket) SetMarketAuthorityAccount(marketAuthority ag_solanago.PublicKey) *InitMarket {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(marketAuthority).WRITE()
	return inst
}

// GetMarketAuthorityAccount gets the "marketAuthority" account.
func (inst *InitMarket) GetMarketAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMarketMintAccount sets the "marketMint" account.
func (inst *InitMarket) SetMarketMintAccount(marketMint ag_solanago.PublicKey) *InitMarket {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(marketMint).WRITE()
	return inst
}

// GetMarketMintAccount gets the "marketMint" account.
func (inst *InitMarket) GetMarketMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOracleAccount sets the "oracle" account.
func (inst *InitMarket) SetOracleAccount(oracle ag_solanago.PublicKey) *InitMarket {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(oracle).WRITE().SIGNER()
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *InitMarket) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitMarket) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitMarket {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitMarket) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst InitMarket) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitMarket,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitMarket) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitMarket) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MarketUid == nil {
			return errors.New("MarketUid parameter is not set")
		}
		if inst.Name == nil {
			return errors.New("Name parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.MarketAuthority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MarketMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Oracle is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *InitMarket) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitMarket")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("MarketUid", *inst.MarketUid))
						paramsBranch.Child(ag_format.Param("     Name", *inst.Name))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("marketAuthority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("     marketMint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         oracle", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("  systemProgram", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj InitMarket) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MarketUid` param:
	err = encoder.Encode(obj.MarketUid)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitMarket) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MarketUid`:
	err = decoder.Decode(&obj.MarketUid)
	if err != nil {
		return err
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	return nil
}

// NewInitMarketInstruction declares a new InitMarket instruction with the provided parameters and accounts.
func NewInitMarketInstruction(
	// Parameters:
	marketUid ag_solanago.PublicKey,
	name string,
	// Accounts:
	marketAuthority ag_solanago.PublicKey,
	marketMint ag_solanago.PublicKey,
	oracle ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *InitMarket {
	return NewInitMarketInstructionBuilder().
		SetMarketUid(marketUid).
		SetName(name).
		SetMarketAuthorityAccount(marketAuthority).
		SetMarketMintAccount(marketMint).
		SetOracleAccount(oracle).
		SetSystemProgramAccount(systemProgram)
}
