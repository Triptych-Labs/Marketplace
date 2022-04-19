// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package someplace

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type QuestAccount struct {
	Stage              uint8
	StartTime          int64
	EndTime            int64
	DepositTokenAmount ag_solanago.PublicKey
	Initializer        ag_solanago.PublicKey
}

var QuestAccountDiscriminator = [8]byte{150, 179, 23, 90, 199, 60, 121, 92}

func (obj QuestAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(QuestAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Stage` param:
	err = encoder.Encode(obj.Stage)
	if err != nil {
		return err
	}
	// Serialize `StartTime` param:
	err = encoder.Encode(obj.StartTime)
	if err != nil {
		return err
	}
	// Serialize `EndTime` param:
	err = encoder.Encode(obj.EndTime)
	if err != nil {
		return err
	}
	// Serialize `DepositTokenAmount` param:
	err = encoder.Encode(obj.DepositTokenAmount)
	if err != nil {
		return err
	}
	// Serialize `Initializer` param:
	err = encoder.Encode(obj.Initializer)
	if err != nil {
		return err
	}
	return nil
}

func (obj *QuestAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(QuestAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[150 179 23 90 199 60 121 92]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Stage`:
	err = decoder.Decode(&obj.Stage)
	if err != nil {
		return err
	}
	// Deserialize `StartTime`:
	err = decoder.Decode(&obj.StartTime)
	if err != nil {
		return err
	}
	// Deserialize `EndTime`:
	err = decoder.Decode(&obj.EndTime)
	if err != nil {
		return err
	}
	// Deserialize `DepositTokenAmount`:
	err = decoder.Decode(&obj.DepositTokenAmount)
	if err != nil {
		return err
	}
	// Deserialize `Initializer`:
	err = decoder.Decode(&obj.Initializer)
	if err != nil {
		return err
	}
	return nil
}

type Batch struct {
	Name   string
	Oracle ag_solanago.PublicKey
	Data   CandyMachineData
}

var BatchDiscriminator = [8]byte{156, 194, 70, 44, 22, 88, 137, 44}

func (obj Batch) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(BatchDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Oracle` param:
	err = encoder.Encode(obj.Oracle)
	if err != nil {
		return err
	}
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Batch) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(BatchDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[156 194 70 44 22 88 137 44]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Oracle`:
	err = decoder.Decode(&obj.Oracle)
	if err != nil {
		return err
	}
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	return nil
}

type BatchReceipt struct {
	Id           uint64
	Name         string
	BatchAccount ag_solanago.PublicKey
	Oracle       ag_solanago.PublicKey
}

var BatchReceiptDiscriminator = [8]byte{121, 10, 11, 19, 248, 174, 122, 64}

func (obj BatchReceipt) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(BatchReceiptDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Id` param:
	err = encoder.Encode(obj.Id)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `BatchAccount` param:
	err = encoder.Encode(obj.BatchAccount)
	if err != nil {
		return err
	}
	// Serialize `Oracle` param:
	err = encoder.Encode(obj.Oracle)
	if err != nil {
		return err
	}
	return nil
}

func (obj *BatchReceipt) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(BatchReceiptDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[121 10 11 19 248 174 122 64]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Id`:
	err = decoder.Decode(&obj.Id)
	if err != nil {
		return err
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `BatchAccount`:
	err = decoder.Decode(&obj.BatchAccount)
	if err != nil {
		return err
	}
	// Deserialize `Oracle`:
	err = decoder.Decode(&obj.Oracle)
	if err != nil {
		return err
	}
	return nil
}

type Batches struct {
	Counter uint64
	Oracle  ag_solanago.PublicKey
}

var BatchesDiscriminator = [8]byte{177, 228, 145, 44, 38, 169, 111, 109}

func (obj Batches) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(BatchesDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Counter` param:
	err = encoder.Encode(obj.Counter)
	if err != nil {
		return err
	}
	// Serialize `Oracle` param:
	err = encoder.Encode(obj.Oracle)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Batches) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(BatchesDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[177 228 145 44 38 169 111 109]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Counter`:
	err = decoder.Decode(&obj.Counter)
	if err != nil {
		return err
	}
	// Deserialize `Oracle`:
	err = decoder.Decode(&obj.Oracle)
	if err != nil {
		return err
	}
	return nil
}

type TreasuryWhitelist struct {
	WhitelistId         uint64
	CandyMachineId      ag_solanago.PublicKey
	CandyMachineCreator ag_solanago.PublicKey
	TreasuryAuthority   ag_solanago.PublicKey
	Oracle              ag_solanago.PublicKey
}

var TreasuryWhitelistDiscriminator = [8]byte{193, 226, 184, 5, 35, 65, 189, 119}

func (obj TreasuryWhitelist) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(TreasuryWhitelistDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `WhitelistId` param:
	err = encoder.Encode(obj.WhitelistId)
	if err != nil {
		return err
	}
	// Serialize `CandyMachineId` param:
	err = encoder.Encode(obj.CandyMachineId)
	if err != nil {
		return err
	}
	// Serialize `CandyMachineCreator` param:
	err = encoder.Encode(obj.CandyMachineCreator)
	if err != nil {
		return err
	}
	// Serialize `TreasuryAuthority` param:
	err = encoder.Encode(obj.TreasuryAuthority)
	if err != nil {
		return err
	}
	// Serialize `Oracle` param:
	err = encoder.Encode(obj.Oracle)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TreasuryWhitelist) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(TreasuryWhitelistDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[193 226 184 5 35 65 189 119]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `WhitelistId`:
	err = decoder.Decode(&obj.WhitelistId)
	if err != nil {
		return err
	}
	// Deserialize `CandyMachineId`:
	err = decoder.Decode(&obj.CandyMachineId)
	if err != nil {
		return err
	}
	// Deserialize `CandyMachineCreator`:
	err = decoder.Decode(&obj.CandyMachineCreator)
	if err != nil {
		return err
	}
	// Deserialize `TreasuryAuthority`:
	err = decoder.Decode(&obj.TreasuryAuthority)
	if err != nil {
		return err
	}
	// Deserialize `Oracle`:
	err = decoder.Decode(&obj.Oracle)
	if err != nil {
		return err
	}
	return nil
}

type TreasuryAuthority struct {
	Whitelists           uint64
	Oracle               ag_solanago.PublicKey
	TreasuryDecimals     uint8
	TreasuryTokenAccount ag_solanago.PublicKey
	TreasuryMint         ag_solanago.PublicKey
	Adornment            string
}

var TreasuryAuthorityDiscriminator = [8]byte{56, 251, 232, 94, 128, 197, 59, 150}

func (obj TreasuryAuthority) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(TreasuryAuthorityDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Whitelists` param:
	err = encoder.Encode(obj.Whitelists)
	if err != nil {
		return err
	}
	// Serialize `Oracle` param:
	err = encoder.Encode(obj.Oracle)
	if err != nil {
		return err
	}
	// Serialize `TreasuryDecimals` param:
	err = encoder.Encode(obj.TreasuryDecimals)
	if err != nil {
		return err
	}
	// Serialize `TreasuryTokenAccount` param:
	err = encoder.Encode(obj.TreasuryTokenAccount)
	if err != nil {
		return err
	}
	// Serialize `TreasuryMint` param:
	err = encoder.Encode(obj.TreasuryMint)
	if err != nil {
		return err
	}
	// Serialize `Adornment` param:
	err = encoder.Encode(obj.Adornment)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TreasuryAuthority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(TreasuryAuthorityDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[56 251 232 94 128 197 59 150]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Whitelists`:
	err = decoder.Decode(&obj.Whitelists)
	if err != nil {
		return err
	}
	// Deserialize `Oracle`:
	err = decoder.Decode(&obj.Oracle)
	if err != nil {
		return err
	}
	// Deserialize `TreasuryDecimals`:
	err = decoder.Decode(&obj.TreasuryDecimals)
	if err != nil {
		return err
	}
	// Deserialize `TreasuryTokenAccount`:
	err = decoder.Decode(&obj.TreasuryTokenAccount)
	if err != nil {
		return err
	}
	// Deserialize `TreasuryMint`:
	err = decoder.Decode(&obj.TreasuryMint)
	if err != nil {
		return err
	}
	// Deserialize `Adornment`:
	err = decoder.Decode(&obj.Adornment)
	if err != nil {
		return err
	}
	return nil
}

type Listing struct {
	TreasuryAuthority ag_solanago.PublicKey
	Batch             ag_solanago.PublicKey
	Oracle            ag_solanago.PublicKey
	ConfigIndex       uint64
	Price             uint64
	LifecycleStart    uint64
	LifecycleEnd      uint64
}

var ListingDiscriminator = [8]byte{218, 32, 50, 73, 43, 134, 26, 58}

func (obj Listing) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ListingDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `TreasuryAuthority` param:
	err = encoder.Encode(obj.TreasuryAuthority)
	if err != nil {
		return err
	}
	// Serialize `Batch` param:
	err = encoder.Encode(obj.Batch)
	if err != nil {
		return err
	}
	// Serialize `Oracle` param:
	err = encoder.Encode(obj.Oracle)
	if err != nil {
		return err
	}
	// Serialize `ConfigIndex` param:
	err = encoder.Encode(obj.ConfigIndex)
	if err != nil {
		return err
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `LifecycleStart` param:
	err = encoder.Encode(obj.LifecycleStart)
	if err != nil {
		return err
	}
	// Serialize `LifecycleEnd` param:
	err = encoder.Encode(obj.LifecycleEnd)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Listing) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ListingDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[218 32 50 73 43 134 26 58]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `TreasuryAuthority`:
	err = decoder.Decode(&obj.TreasuryAuthority)
	if err != nil {
		return err
	}
	// Deserialize `Batch`:
	err = decoder.Decode(&obj.Batch)
	if err != nil {
		return err
	}
	// Deserialize `Oracle`:
	err = decoder.Decode(&obj.Oracle)
	if err != nil {
		return err
	}
	// Deserialize `ConfigIndex`:
	err = decoder.Decode(&obj.ConfigIndex)
	if err != nil {
		return err
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `LifecycleStart`:
	err = decoder.Decode(&obj.LifecycleStart)
	if err != nil {
		return err
	}
	// Deserialize `LifecycleEnd`:
	err = decoder.Decode(&obj.LifecycleEnd)
	if err != nil {
		return err
	}
	return nil
}

type Market struct {
	MarketDecimals uint8
	Listings       uint64
	Name           string
	MarketMint     ag_solanago.PublicKey
	MarketUid      ag_solanago.PublicKey
	Oracle         ag_solanago.PublicKey
}

var MarketDiscriminator = [8]byte{219, 190, 213, 55, 0, 227, 198, 154}

func (obj Market) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(MarketDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `MarketDecimals` param:
	err = encoder.Encode(obj.MarketDecimals)
	if err != nil {
		return err
	}
	// Serialize `Listings` param:
	err = encoder.Encode(obj.Listings)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `MarketMint` param:
	err = encoder.Encode(obj.MarketMint)
	if err != nil {
		return err
	}
	// Serialize `MarketUid` param:
	err = encoder.Encode(obj.MarketUid)
	if err != nil {
		return err
	}
	// Serialize `Oracle` param:
	err = encoder.Encode(obj.Oracle)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Market) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(MarketDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[219 190 213 55 0 227 198 154]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `MarketDecimals`:
	err = decoder.Decode(&obj.MarketDecimals)
	if err != nil {
		return err
	}
	// Deserialize `Listings`:
	err = decoder.Decode(&obj.Listings)
	if err != nil {
		return err
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `MarketMint`:
	err = decoder.Decode(&obj.MarketMint)
	if err != nil {
		return err
	}
	// Deserialize `MarketUid`:
	err = decoder.Decode(&obj.MarketUid)
	if err != nil {
		return err
	}
	// Deserialize `Oracle`:
	err = decoder.Decode(&obj.Oracle)
	if err != nil {
		return err
	}
	return nil
}

type MarketListing struct {
	MarketAuthority          ag_solanago.PublicKey
	NftMint                  ag_solanago.PublicKey
	SellerMarketTokenAccount ag_solanago.PublicKey
	Index                    uint64
	Price                    uint64
	ListedAt                 uint64
	Fulfilled                uint64
}

var MarketListingDiscriminator = [8]byte{175, 123, 31, 97, 53, 211, 229, 16}

func (obj MarketListing) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(MarketListingDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `MarketAuthority` param:
	err = encoder.Encode(obj.MarketAuthority)
	if err != nil {
		return err
	}
	// Serialize `NftMint` param:
	err = encoder.Encode(obj.NftMint)
	if err != nil {
		return err
	}
	// Serialize `SellerMarketTokenAccount` param:
	err = encoder.Encode(obj.SellerMarketTokenAccount)
	if err != nil {
		return err
	}
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
	// Serialize `ListedAt` param:
	err = encoder.Encode(obj.ListedAt)
	if err != nil {
		return err
	}
	// Serialize `Fulfilled` param:
	err = encoder.Encode(obj.Fulfilled)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MarketListing) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(MarketListingDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[175 123 31 97 53 211 229 16]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `MarketAuthority`:
	err = decoder.Decode(&obj.MarketAuthority)
	if err != nil {
		return err
	}
	// Deserialize `NftMint`:
	err = decoder.Decode(&obj.NftMint)
	if err != nil {
		return err
	}
	// Deserialize `SellerMarketTokenAccount`:
	err = decoder.Decode(&obj.SellerMarketTokenAccount)
	if err != nil {
		return err
	}
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
	// Deserialize `ListedAt`:
	err = decoder.Decode(&obj.ListedAt)
	if err != nil {
		return err
	}
	// Deserialize `Fulfilled`:
	err = decoder.Decode(&obj.Fulfilled)
	if err != nil {
		return err
	}
	return nil
}
