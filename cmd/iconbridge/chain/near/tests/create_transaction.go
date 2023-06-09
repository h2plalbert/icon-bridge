package tests

// import (
// 	"crypto/ed25519"

// 	"github.com/btcsuite/btcutil/base58"
// 	"github.com/icon-project/icon-bridge/cmd/iconbridge/chain/base"
// 	"github.com/icon-project/icon-bridge/cmd/iconbridge/chain/near/tests/mock"
// 	"github.com/icon-project/icon-bridge/cmd/iconbridge/chain/near/types"
// 	"github.com/icon-project/icon-bridge/common/wallet"
// )

// type CreateTransactionTest struct {
// 	description string
// 	testData    []TestData
// }

// func (t CreateTransactionTest) Description() string {
// 	return t.description
// }

// func (t CreateTransactionTest) TestDatas() []TestData {
// 	return t.testData
// }

// type Input struct {
// 	Wallet                base.Wallet
// 	TransactionParameters types.TransactionParam
// 	Hash                  string
// }

// func Wallet(key string) base.Wallet {
// 	privateKey := base58.Decode(key)
// 	newPrivateKey := ed25519.PrivateKey(privateKey)
// 	nearWallet, _ := wallet.NewNearwalletFromPrivateKey(&newPrivateKey)

// 	return base.Wallet(nearWallet)
// }

// func init() {

// 	var testData = []TestData{
// 		{
// 			Description: "CreateTransaction Success",
// 			Input: Input{

// 				Wallet: Wallet("22yx6AjQgG1jGuAmPuEwLnVKFnuq5LU23dbU3JBZodKxrJ8dmmqpDZKtRSfiU4F8UQmv1RiZSrjWhQMQC3ye7M1J"),
// 				TransactionParameters: types.TransactionParam{
// 					To:   "dev-20211206025826-24100687319598",
// 					From: "69c003c3b80ed12ea02f5c67c9e8167f0ce3b2e8020a0f43b1029c4d787b0d21",
// 					RelayMessage: types.RelayMessageParam{
// 						Previous: "btp://0x1.icon/0xc294b1A62E82d3f135A8F9b2f9cAEAA23fbD6Cf5",
// 						Message:  "-QNlwvgA-AD5A125A1r5A1cAALkBVPkBUbkBTvkBS4IgALkBRfkBQgCVAZQ5KOt2a33MGDNmJIyCIn-oz8dWgwdSZIMHUmSFAukO3QC49hAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAgQAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAABAAACAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAACAAAAAAIAAAAEIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAACAAAAAAAAAAAAAAAAQAAAAAAAAAQAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAPgA-ACgvEWAgjiMdu0pUAxvQcurR0OVNFEmP5ZzhNMkpIVwqPr5Afn5AfYAuQHy-QHvo-IQoMOetu_ZC-YK5vgAe7UsNeBouviaGeP9VQgl1JDY7dRguFP4UaBjTIu0WBCx1UEf8MBqzibUyAkFNsmaQ3gnewIx7LIl96CegRcYwno7-6NZanIbuiC4o_93OVbLE-UByERTjAHFv4CAgICAgICAgICAgICAgLkBc_kBcCC5AWz5AWmVAZwHK-9UWr65aj97ov_fi4YftV3q-FSWTWVzc2FnZShzdHIsaW50LGJ5dGVzKbg6YnRwOi8vMHg1MDEucHJhLzB4NUNDMzA3MjY4YTEzOTNBQjlBNzY0QTIwREFDRTg0OEFCODI3NWM0NgL4-7j5-Pe4PmJ0cDovLzB4NThlYjFjLmljb24vY3g5YzA3MmJlZjU0NWFiZWI5NmEzZjdiYTJmZmRmOGI4NjFmYjU1ZGVhuDpidHA6Ly8weDUwMS5wcmEvMHg1Q0MzMDcyNjhhMTM5M0FCOUE3NjRBMjBEQUNFODQ4QUI4Mjc1YzQ2im5hdGl2ZWNvaW4BuG34awC4aPhmqmh4NDUwMmFhZDc5ODZhZDVhODQ4OTU1MTVmYWY3NmU5MGI1YjQ3ODY1NKoweDE1OEEzOTFGMzUwMEMzMjg4QWIyODY1MzcyMmE2NDU5RTc3MjZCMDHPzoNJQ1iJAIlj3YwsXgAA-AA=",
// 					},
// 				},
// 			},
// 			Expected: struct {
// 				Success interface{}
// 				Fail    interface{}
// 			}{
// 				Success: "QAAAADY5YzAwM2MzYjgwZWQxMmVhMDJmNWM2N2M5ZTgxNjdmMGNlM2IyZTgwMjBhMGY0M2IxMDI5YzRkNzg3YjBkMjEAacADw7gO0S6gL1xnyegWfwzjsugCCg9DsQKcTXh7DSEdAAAAAAAAACEAAABkZXYtMjAyMTEyMDYwMjU4MjYtMjQxMDA2ODczMTk1OTimx8l6NhI3ulzQBDj+se0H2mCBiedMOex5YU9ftNRRYAEAAAACFAAAAGhhbmRsZV9yZWxheV9tZXNzYWdl3wQAAHsic291cmNlIjoiYnRwOi8vMHgxLmljb24vMHhjMjk0YjFBNjJFODJkM2YxMzVBOEY5YjJmOWNBRUFBMjNmYkQ2Q2Y1IiwibWVzc2FnZSI6Ii1RTmx3dmdBLUFENUExMjVBMXI1QTFjQUFMa0JWUGtCVWJrQlR2a0JTNElnQUxrQlJma0JRZ0NWQVpRNUtPdDJhMzNNR0RObUpJeUNJbi1vejhkV2d3ZFNaSU1IVW1TRkF1a08zUUM0OWhBQUFBQUFBQUFCQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBZ1FBQUFBQUFBQUFBQUFBRUFBQUFBQUFBQUFBQUFBQUFBQUFBQUJBQUFDQUFBQUFBQUFBQUFBQUFBQUNBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFnQUFBQUFBQUFBQUFBQUFBQUFBRUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUlBQUFBQUFDQUFBQUFBSUFBQUFFSUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQWdBQUFBQUFBQ0FBQUFBQUFBQUFBQUFBQUFRQUFBQUFBQUFBUUFBQUFBQUFBQkFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFJQUFBQUFBQUFBQUFQZ0EtQUNndkVXQWdqaU1kdTBwVUF4dlFjdXJSME9WTkZFbVA1WnpoTk1rcElWd3FQcjVBZm41QWZZQXVRSHktUUh2by1JUW9NT2V0dV9aQy1ZSzV2Z0FlN1VzTmVCb3V2aWFHZVA5VlFnbDFKRFk3ZFJndUZQNFVhQmpUSXUwV0JDeDFVRWY4TUJxemliVXlBa0ZOc21hUTNnbmV3SXg3TElsOTZDZWdSY1l3bm83LTZOWmFuSWJ1aUM0b185M09WYkxFLVVCeUVSVGpBSEZ2NENBZ0lDQWdJQ0FnSUNBZ0lDQWdMa0JjX2tCY0NDNUFXejVBV21WQVp3SEstOVVXcjY1YWo5N292X2ZpNFlmdFYzcS1GU1dUV1Z6YzJGblpTaHpkSElzYVc1MExHSjVkR1Z6S2JnNlluUndPaTh2TUhnMU1ERXVjSEpoTHpCNE5VTkRNekEzTWpZNFlURXpPVE5CUWpsQk56WTBRVEl3UkVGRFJUZzBPRUZDT0RJM05XTTBOZ0w0LTdqNS1QZTRQbUowY0Rvdkx6QjROVGhsWWpGakxtbGpiMjR2WTNnNVl6QTNNbUpsWmpVME5XRmlaV0k1Tm1FelpqZGlZVEptWm1SbU9HSTROakZtWWpVMVpHVmh1RHBpZEhBNkx5OHdlRFV3TVM1d2NtRXZNSGcxUTBNek1EY3lOamhoTVRNNU0wRkNPVUUzTmpSQk1qQkVRVU5GT0RRNFFVSTRNamMxWXpRMmltNWhkR2wyWldOdmFXNEJ1RzM0YXdDNGFQaG1xbWg0TkRVd01tRmhaRGM1T0RaaFpEVmhPRFE0T1RVMU1UVm1ZV1kzTm1VNU1HSTFZalEzT0RZMU5Lb3dlREUxT0VFek9URkdNelV3TUVNek1qZzRRV0l5T0RZMU16Y3lNbUUyTkRVNVJUYzNNalpDTURIUHpvTkpRMWlKQUlsajNZd3NYZ0FBLUFBPSJ9AMBuMdkQAQAAAAAAAAAAAAAAAAAAAAAAAM2waD9GzXcTd1QKbrrEgEskh0rDUCDDSVhA2jgP5I5db0Hqxs21GEZPWNWlrtm0kIcRzYWnooPEH8fcv0TTwgg=",
// 			},
// 			MockStorage: func() mock.Storage {

// 				noncemap := mock.LoadNonceFromFile([]string{"69c003c3b80ed12ea02f5c67c9e8167f0ce3b2e8020a0f43b1029c4d787b0d21"})

// 				return mock.Storage{
// 					NonceMap: noncemap,
// 				}
// 			}(),
// 		},
// 		{
// 			Description: "CreateTransaction Fail",
// 			Input: Input{
// 				Wallet: base.Wallet(nil),
// 				TransactionParameters: types.TransactionParam{
// 					To:   "dev-20211206025826-24100687319598",
// 					From: "69c003c3b80ed12ea02f5c67c9e8167f0ce3b2e8020a0f43b1029c4d787b0d21",
// 					RelayMessage: types.RelayMessageParam{
// 						Previous: "btp://0x1.icon/0xc294b1A62E82d3f135A8F9b2f9cAEAA23fbD6Cf5",
// 						Message:  "-QNlwvgA-AD5A125A1r5A1cAALkBVPkBUbkBTvkBS4IgALkBRfkBQgCVAZQ5KOt2a33MGDNmJIyCIn-oz8dWgwdSZIMHUmSFAukO3QC49hAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAgQAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAABAAACAAAAAAAAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAACAAAAAAIAAAAEIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAACAAAAAAAAAAAAAAAAQAAAAAAAAAQAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAPgA-ACgvEWAgjiMdu0pUAxvQcurR0OVNFEmP5ZzhNMkpIVwqPr5Afn5AfYAuQHy-QHvo-IQoMOetu_ZC-YK5vgAe7UsNeBouviaGeP9VQgl1JDY7dRguFP4UaBjTIu0WBCx1UEf8MBqzibUyAkFNsmaQ3gnewIx7LIl96CegRcYwno7-6NZanIbuiC4o_93OVbLE-UByERTjAHFv4CAgICAgICAgICAgICAgLkBc_kBcCC5AWz5AWmVAZwHK-9UWr65aj97ov_fi4YftV3q-FSWTWVzc2FnZShzdHIsaW50LGJ5dGVzKbg6YnRwOi8vMHg1MDEucHJhLzB4NUNDMzA3MjY4YTEzOTNBQjlBNzY0QTIwREFDRTg0OEFCODI3NWM0NgL4-7j5-Pe4PmJ0cDovLzB4NThlYjFjLmljb24vY3g5YzA3MmJlZjU0NWFiZWI5NmEzZjdiYTJmZmRmOGI4NjFmYjU1ZGVhuDpidHA6Ly8weDUwMS5wcmEvMHg1Q0MzMDcyNjhhMTM5M0FCOUE3NjRBMjBEQUNFODQ4QUI4Mjc1YzQ2im5hdGl2ZWNvaW4BuG34awC4aPhmqmh4NDUwMmFhZDc5ODZhZDVhODQ4OTU1MTVmYWY3NmU5MGI1YjQ3ODY1NKoweDE1OEEzOTFGMzUwMEMzMjg4QWIyODY1MzcyMmE2NDU5RTc3MjZCMDHPzoNJQ1iJAIlj3YwsXgAA-AA=",
// 					},
// 				},
// 			},
// 			Expected: struct {
// 				Success interface{}
// 				Fail    interface{}
// 			}{
// 				Success: nil,
// 			},
// 		},
// 	}

// 	RegisterTest("CreateTransaction", CreateTransactionTest{
// 		description: "Test CreateTransaction",
// 		testData:    testData,
// 	})

// }
