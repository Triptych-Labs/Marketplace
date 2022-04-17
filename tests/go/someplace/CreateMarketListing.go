// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package someplace

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreateMarketListing is the `createMarketListing` instruction.
type CreateMarketListing struct {
	Index *uint64
	Price *uint64

	// [0] = [WRITE] marketAuthority
	//
	// [1] = [WRITE] marketListing
	//
	// [2] = [WRITE] marketListingTokenAccount
	//
	// [3] = [WRITE, SIGNER] seller
	//
	// [4] = [WRITE] sellerNftTokenAccount
	//
	// [5] = [WRITE] sellerMarketTokenAccount
	//
	// [6] = [WRITE] nftMint
	//
	// [7] = [] systemProgram
	//
	// [8] = [] tokenProgram
	//
	// [9] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateMarketListingInstructionBuilder creates a new `CreateMarketListing` instruction builder.
func NewCreateMarketListingInstructionBuilder() *CreateMarketListing {
	nd := &CreateMarketListing{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

// SetIndex sets the "index" parameter.
func (inst *CreateMarketListing) SetIndex(index uint64) *CreateMarketListing {
	inst.Index = &index
	return inst
}

// SetPrice sets the "price" parameter.
func (inst *CreateMarketListing) SetPrice(price uint64) *CreateMarketListing {
	inst.Price = &price
	return inst
}

// SetMarketAuthorityAccount sets the "marketAuthority" account.
func (inst *CreateMarketListing) SetMarketAuthorityAccount(marketAuthority ag_solanago.PublicKey) *CreateMarketListing {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(marketAuthority).WRITE()
	return inst
}

// GetMarketAuthorityAccount gets the "marketAuthority" account.
func (inst *CreateMarketListing) GetMarketAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMarketListingAccount sets the "marketListing" account.
func (inst *CreateMarketListing) SetMarketListingAccount(marketListing ag_solanago.PublicKey) *CreateMarketListing {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(marketListing).WRITE()
	return inst
}

// GetMarketListingAccount gets the "marketListing" account.
func (inst *CreateMarketListing) GetMarketListingAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMarketListingTokenAccountAccount sets the "marketListingTokenAccount" account.
func (inst *CreateMarketListing) SetMarketListingTokenAccountAccount(marketListingTokenAccount ag_solanago.PublicKey) *CreateMarketListing {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(marketListingTokenAccount).WRITE()
	return inst
}

// GetMarketListingTokenAccountAccount gets the "marketListingTokenAccount" account.
func (inst *CreateMarketListing) GetMarketListingTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSellerAccount sets the "seller" account.
func (inst *CreateMarketListing) SetSellerAccount(seller ag_solanago.PublicKey) *CreateMarketListing {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(seller).WRITE().SIGNER()
	return inst
}

// GetSellerAccount gets the "seller" account.
func (inst *CreateMarketListing) GetSellerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSellerNftTokenAccountAccount sets the "sellerNftTokenAccount" account.
func (inst *CreateMarketListing) SetSellerNftTokenAccountAccount(sellerNftTokenAccount ag_solanago.PublicKey) *CreateMarketListing {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(sellerNftTokenAccount).WRITE()
	return inst
}

// GetSellerNftTokenAccountAccount gets the "sellerNftTokenAccount" account.
func (inst *CreateMarketListing) GetSellerNftTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSellerMarketTokenAccountAccount sets the "sellerMarketTokenAccount" account.
func (inst *CreateMarketListing) SetSellerMarketTokenAccountAccount(sellerMarketTokenAccount ag_solanago.PublicKey) *CreateMarketListing {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(sellerMarketTokenAccount).WRITE()
	return inst
}

// GetSellerMarketTokenAccountAccount gets the "sellerMarketTokenAccount" account.
func (inst *CreateMarketListing) GetSellerMarketTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetNftMintAccount sets the "nftMint" account.
func (inst *CreateMarketListing) SetNftMintAccount(nftMint ag_solanago.PublicKey) *CreateMarketListing {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(nftMint).WRITE()
	return inst
}

// GetNftMintAccount gets the "nftMint" account.
func (inst *CreateMarketListing) GetNftMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CreateMarketListing) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateMarketListing {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CreateMarketListing) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *CreateMarketListing) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *CreateMarketListing {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *CreateMarketListing) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetRentAccount sets the "rent" account.
func (inst *CreateMarketListing) SetRentAccount(rent ag_solanago.PublicKey) *CreateMarketListing {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *CreateMarketListing) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

func (inst CreateMarketListing) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateMarketListing,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateMarketListing) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateMarketListing) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Index == nil {
			return errors.New("Index parameter is not set")
		}
		if inst.Price == nil {
			return errors.New("Price parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.MarketAuthority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MarketListing is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.MarketListingTokenAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Seller is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SellerNftTokenAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SellerMarketTokenAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.NftMint is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *CreateMarketListing) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateMarketListing")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Index", *inst.Index))
						paramsBranch.Child(ag_format.Param("Price", *inst.Price))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   marketAuthority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("     marketListing", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("marketListingToken", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("            seller", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("    sellerNftToken", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta(" sellerMarketToken", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           nftMint", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("     systemProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("      tokenProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("              rent", inst.AccountMetaSlice.Get(9)))
					})
				})
		})
}

func (obj CreateMarketListing) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Index` param:
	err = encoder.Encode(obj.Index)
	if err != nil {
		return err
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateMarketListing) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Index`:
	err = decoder.Decode(&obj.Index)
	if err != nil {
		return err
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateMarketListingInstruction declares a new CreateMarketListing instruction with the provided parameters and accounts.
func NewCreateMarketListingInstruction(
	// Parameters:
	index uint64,
	price uint64,
	// Accounts:
	marketAuthority ag_solanago.PublicKey,
	marketListing ag_solanago.PublicKey,
	marketListingTokenAccount ag_solanago.PublicKey,
	seller ag_solanago.PublicKey,
	sellerNftTokenAccount ag_solanago.PublicKey,
	sellerMarketTokenAccount ag_solanago.PublicKey,
	nftMint ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *CreateMarketListing {
	return NewCreateMarketListingInstructionBuilder().
		SetIndex(index).
		SetPrice(price).
		SetMarketAuthorityAccount(marketAuthority).
		SetMarketListingAccount(marketListing).
		SetMarketListingTokenAccountAccount(marketListingTokenAccount).
		SetSellerAccount(seller).
		SetSellerNftTokenAccountAccount(sellerNftTokenAccount).
		SetSellerMarketTokenAccountAccount(sellerMarketTokenAccount).
		SetNftMintAccount(nftMint).
		SetSystemProgramAccount(systemProgram).
		SetTokenProgramAccount(tokenProgram).
		SetRentAccount(rent)
}
