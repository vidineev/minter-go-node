package validators

func GetValidatorsCountForBlock(block int64) int {
	return 100 // only for current testnet

	//count := 16 + (block/518400)*4
	//
	//if count > 256 {
	//	return 256
	//}
	//
	//return int(count)
}

func GetCandidatesCountForBlock(block int64) int {
	return GetValidatorsCountForBlock(block) * 3
}
