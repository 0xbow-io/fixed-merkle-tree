package fMerkleTree

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_BaseTree_Zeros(t *testing.T) {
	t.Run("should have correct zeros", func(t *testing.T) {
		zero, err := hex.DecodeString("2fe54c60d3acabf3343a35b6eba15db4821b340f76e741e2249685ed4899af6c")
		require.NoError(t, err)
		tree, err := NewMerkleTree(32, []Element{}, Element(zero), Poseidon2)
		require.NoError(t, err)
		require.Len(t, tree.BaseTree.zeros, 33)
		require.Equal(t, tree.BaseTree.zeros[1].Hex(), "1a332ca2cd2436bdc6796e6e4244ebf6f7e359868b7252e55342f766e4088082")
		require.Equal(t, tree.BaseTree.zeros[2].Hex(), "2fb19ac27499bdf9d7d3b387eff42b6d12bffbc6206e81d0ef0b0d6b24520ebd")
		require.Equal(t, tree.BaseTree.zeros[3].Hex(), "18d0d6e282d4eacbf18efc619a986db763b75095ed122fac7d4a49418daa42e1")
		require.Equal(t, tree.BaseTree.zeros[4].Hex(), "054dec40f76a0f5aaeff1a85a4a3721b92b4ad244362d30b0ef8ed7033de11d3")
		require.Equal(t, tree.BaseTree.zeros[5].Hex(), "1d24c91f8d40f1c2591edec19d392905cf5eb01eada48d71836177ef11aea5b2")
		require.Equal(t, tree.BaseTree.zeros[6].Hex(), "0fb63621cfc047eba2159faecfa55b120d7c81c0722633ef94e20e27675e378f")
		require.Equal(t, tree.BaseTree.zeros[7].Hex(), "277b08f214fe8c5504a79614cdec5abd7b6adc9133fe926398684c82fd798b44")
		require.Equal(t, tree.BaseTree.zeros[8].Hex(), "2633613437c1fd97f7c798e2ea30d52cfddee56d74f856a541320ae86ddaf2de")
		require.Equal(t, tree.BaseTree.zeros[9].Hex(), "00768963fa4b993fbfece3619bfaa3ca4afd7e3864f11b09a0849dbf4ad25807")
		require.Equal(t, tree.BaseTree.zeros[10].Hex(), "0e63ff9df484c1a21478bd27111763ef203177ec0a7ef3a3cd43ec909f587bb0")

		require.Equal(t, tree.BaseTree.zeros[11].Hex(), "0e6a4bfb0dd0ac8bf5517eaac48a95ba783dabe9f64494f9c892d3e8431eaab3")
		require.Equal(t, tree.BaseTree.zeros[12].Hex(), "0164a46b3ffff8baca00de7a130a63d105f1578076838502b99488505d5b3d35")
		require.Equal(t, tree.BaseTree.zeros[13].Hex(), "145a6f1521c02b250cc76eb35cd67c9b0b22473577de3778e4c51903836c8957")
		require.Equal(t, tree.BaseTree.zeros[14].Hex(), "29849fc5b55303a660bad33d986fd156d48516ec58a0f0a561a03b704a802254")
		require.Equal(t, tree.BaseTree.zeros[15].Hex(), "26639dd486b374e98ac6da34e8651b3fca58c51f1c2f857dd82045f27fc8dbe6")
		require.Equal(t, tree.BaseTree.zeros[16].Hex(), "2aa39214b887ee877e60afdb191390344c68177c30a0b8646649774174de5e33")
		require.Equal(t, tree.BaseTree.zeros[17].Hex(), "09b397d253e41a521d042ffe01f8c33ae37d4c7da21af68693aafb63d599d708")
		require.Equal(t, tree.BaseTree.zeros[18].Hex(), "02fbfd397ad901cea38553239aefec016fcb6a19899038503f04814cbb79a511")
		require.Equal(t, tree.BaseTree.zeros[19].Hex(), "266640a877ec97a91f6c95637f843eeac8718f53f311bac9cba7d958df646f9d")
		require.Equal(t, tree.BaseTree.zeros[20].Hex(), "29f9a0a07a22ab214d00aaa0190f54509e853f3119009baecb0035347606b0a9")

		require.Equal(t, tree.BaseTree.zeros[21].Hex(), "0a1fda67bffa0ab3a755f23fdcf922720820b6a96616a5ca34643cd0b935e3d6")
		require.Equal(t, tree.BaseTree.zeros[22].Hex(), "19507199eb76b5ec5abe538a01471d03efb6c6984739c77ec61ada2ba2afb389")
		require.Equal(t, tree.BaseTree.zeros[23].Hex(), "26bd93d26b751484942282e27acfb6d193537327a831df6927e19cdfc73c3e64")
		require.Equal(t, tree.BaseTree.zeros[24].Hex(), "2eb88a9c6b00a4bc6ea253268090fe1d255f6fe02d2eb745517723aae44d7386")
		require.Equal(t, tree.BaseTree.zeros[25].Hex(), "13e50d0bda78be97792df40273cbb16f0dc65c0697d81a82d07d0f6eee80a164")
		require.Equal(t, tree.BaseTree.zeros[26].Hex(), "2ea95776929000133246ff8d9fdcba179d0b262b9e910558309bac1c1ec03d7a")
		require.Equal(t, tree.BaseTree.zeros[27].Hex(), "1a640d6ef66e356c795396c0957b06a99891afe0c493f4d0bdfc0450764bae60")
		require.Equal(t, tree.BaseTree.zeros[28].Hex(), "2b17979f2c2048dd9e4ee5f482cced21435ea8cc54c32f80562e39a5016b0496")
		require.Equal(t, tree.BaseTree.zeros[29].Hex(), "29ba6a30de50542e261abfc7ee0c68911002d3acd4dd4c02ad59aa96805b20bb")
		require.Equal(t, tree.BaseTree.zeros[30].Hex(), "103fcf1c8a98ebe50285f6e669077a579308311fd44bb6895d5da7ba7fd3564e")

		require.Equal(t, tree.BaseTree.zeros[31].Hex(), "166bdd01780976e655f5278260c638dcf10fe7c136f37c9152cbcaabef901f4d")
		require.Equal(t, tree.BaseTree.zeros[32].Hex(), "2712c601a9b8b2abd396a619327095d3f1ea86a6c07d6df416a3973a1a4b3ce5")
	})
}
