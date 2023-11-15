package fMerkleTree

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewMerkleTree(t *testing.T) {
	t.Run("should have correct zero root", func(t *testing.T) {
		tree, err := NewMerkleTree(10, []Element{}, Element{0}, SHA256Hash)
		require.NoError(t, err)
		expected, _ := big.NewInt(0).SetString("cacc702f220ebafc52113f75e75f846a6f8866763a0e6a52ff2bf316424d99ff", 16)
		require.Equal(t, tree.Root().Hex(), hex.EncodeToString(expected.Bytes()))
	})

	t.Run("should have correct 1 element root", func(t *testing.T) {
		tree, err := NewMerkleTree(10, []Element{{1}}, Element{0}, SHA256Hash)
		require.NoError(t, err)
		expected, _ := big.NewInt(0).SetString("485dc9401c1957df7a9c9ef55819c0528a97bb9f1470c665aa38806c2b189364", 16)
		require.Equal(t, tree.Root().Hex(), hex.EncodeToString(expected.Bytes()))
	})

	t.Run("should have correct even elements root", func(t *testing.T) {
		tree, err := NewMerkleTree(10, []Element{{1}, {2}}, Element{0}, SHA256Hash)
		require.NoError(t, err)
		expected, _ := big.NewInt(0).SetString("33bf760045f2c4785da2d90dc442aa1eee200a9baf261b97f336f85ea8b183d2", 16)
		require.Equal(t, tree.Root().Hex(), hex.EncodeToString(expected.Bytes()))
	})

	t.Run("should have correct odd elements root", func(t *testing.T) {
		tree, err := NewMerkleTree(10, []Element{{1}, {2}, {3}}, Element{0}, SHA256Hash)
		require.NoError(t, err)
		expected, _ := big.NewInt(0).SetString("77f3ad4a2b33c519bbcf29ebb545c7a56675e49170819c1250cb59a56b6398e2", 16)
		require.Equal(t, tree.Root().Hex(), hex.EncodeToString(expected.Bytes()))
	})

}

func Test_MerkleTree_Insert(t *testing.T) {
	t.Run("should insert into empty tree", func(t *testing.T) {
		tree, err := NewMerkleTree(10, []Element{}, Element{0}, SHA256Hash)
		require.NoError(t, err)
		require.NoError(t, tree.Insert(Element{42}))
		expected, _ := big.NewInt(0).SetString("e87516e99655b9a121b1a8fb110e956b50056745e1f6f4e991bcb22c38ee091d", 16)
		require.Equal(t, tree.Root().Hex(), hex.EncodeToString(expected.Bytes()))
	})
	t.Run("should insert into odd tree", func(t *testing.T) {
		tree, err := NewMerkleTree(10, []Element{{1}}, Element{0}, SHA256Hash)
		require.NoError(t, err)
		require.NoError(t, tree.Insert(Element{42}))
		expected, _ := big.NewInt(0).SetString("681effa823c22bc59e9979681018d8926d9d19333d4c716ad2f17f2e950222bb", 16)
		require.Equal(t, tree.Root().Hex(), hex.EncodeToString(expected.Bytes()))
	})

	t.Run("should insert into even tree", func(t *testing.T) {
		tree, err := NewMerkleTree(10, []Element{{1}, {2}}, Element{0}, SHA256Hash)
		require.NoError(t, err)
		require.NoError(t, tree.Insert(Element{42}))
		expected, _ := big.NewInt(0).SetString("51345155cf6429903e07949039abdc0cf3bc37ebeafc45af9a029f5dedeb08cc", 16)
		require.Equal(t, tree.Root().Hex(), hex.EncodeToString(expected.Bytes()))
	})
}

func Test_MerkleTree_BulkInsert(t *testing.T) {
	t.Run("should update first element", func(t *testing.T) {
		tree, err := NewMerkleTree(10, []Element{{1}, {2}, {3}}, Element{0}, SHA256Hash)
		require.NoError(t, err)
		require.NoError(t, tree.BulkInsert([]Element{{4}, {5}, {6}}))
		expected, _ := big.NewInt(0).SetString("68f30f28e64ca789f80b0108d36f38deae22764b485ba57cbb1b486bb1711ff7", 16)
		require.Equal(t, tree.Root().Hex(), hex.EncodeToString(expected.Bytes()))
	})
}

func Test_MerkleTree_Update(t *testing.T) {
	t.Run("should update first element", func(t *testing.T) {
		tree, err := NewMerkleTree(10, []Element{{1}, {2}, {3}, {4}, {5}}, Element{0}, SHA256Hash)
		require.NoError(t, err)
		require.NoError(t, tree.Update(0, Element{42}))
		expected, _ := big.NewInt(0).SetString("0fd63059b24afc353a157c12b75fe9ea7b46285fbabe0a331cb55f8317cbbd7a", 16)
		require.Equal(t, tree.Root().Hex(), hex.EncodeToString(expected.Bytes()))
	})

	t.Run("should update last element", func(t *testing.T) {
		tree, err := NewMerkleTree(10, []Element{{1}, {2}, {3}, {4}, {5}}, Element{0}, SHA256Hash)
		require.NoError(t, err)
		require.NoError(t, tree.Update(4, Element{42}))
		expected, _ := big.NewInt(0).SetString("d03b37f0c0054a9d984becc72f43fb74044e3b356477b27a561ea656b2cb8be6", 16)
		require.Equal(t, tree.Root().Hex(), hex.EncodeToString(expected.Bytes()))
	})

	t.Run("should update odd element", func(t *testing.T) {
		tree, err := NewMerkleTree(10, []Element{{1}, {2}, {3}, {4}, {5}}, Element{0}, SHA256Hash)
		require.NoError(t, err)
		require.NoError(t, tree.Update(1, Element{42}))
		expected, _ := big.NewInt(0).SetString("c303f51740299127c9be13c93675aeb193e813c1d8c93603780e87ccd39626e7", 16)
		require.Equal(t, tree.Root().Hex(), hex.EncodeToString(expected.Bytes()))
	})

	t.Run("should update even element", func(t *testing.T) {
		tree, err := NewMerkleTree(10, []Element{{1}, {2}, {3}, {4}, {5}}, Element{0}, SHA256Hash)
		require.NoError(t, err)
		require.NoError(t, tree.Update(2, Element{42}))
		expected, _ := big.NewInt(0).SetString("c3478a3fde5daaf2a99e8930e115b437e46f217f228ecaad33f967af70c0e683", 16)
		require.Equal(t, tree.Root().Hex(), hex.EncodeToString(expected.Bytes()))
	})
}

func TestMerkleTree_Serialization(t *testing.T) {
	tree, err := NewMerkleTree(10, []Element{{1}, {2}, {3}}, Element{0}, SHA256Hash)
	require.NoError(t, err)
	require.NoError(t, tree.BulkInsert([]Element{{4}, {5}, {6}}))
	expected, _ := big.NewInt(0).SetString("68f30f28e64ca789f80b0108d36f38deae22764b485ba57cbb1b486bb1711ff7", 16)
	require.Equal(t, tree.Root().Hex(), hex.EncodeToString(expected.Bytes()))
	data, err := tree.Serialize()
	require.NoError(t, err)

	tree2, err := DeserializeMerkleTree(data, SHA256Hash)
	require.NoError(t, err)
	require.Equal(t, tree.Root(), tree2.Root())
}

func Test_MerkleTree_Path(t *testing.T) {
	t.Run("should work for even index", func(t *testing.T) {
		tree, err := NewMerkleTree(10, []Element{{1}, {2}, {3}, {4}, {5}}, Element{0}, SHA256Hash)
		require.NoError(t, err)
		path, err := tree.Path(2)
		require.NoError(t, err)
		require.ElementsMatch(t, path.PathIndices, []int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0})
		expected := []string{
			"04",
			"6b51d431df5d7f141cbececcf79edf3dd861c3b4069f0b11661a3eefacbba918",
			"3bcf81da18c3a06070dd115fd97e801192b100b798a397004604356a6da2995a",
			"461aa5c7bcac617fa44a126ed3a812b00f963c7fd7344113da0bd452024d81fb",
			"96405940c97198beddd8ec086d669c69e055f8e7288b962258e4cb61b6619dc2",
			"ee38900ad320e5e188c66c3ce41cea0553a1ec098876f790cee0b83df0531233",
			"b75f31fc6b41e45d67ac0d50ed3cae866053735154eb901fae40883ef7c56fb6",
			"fb4e521163400391274a5b2755591aaa5af4ddf32e099aa36aca5ff41f8e31bb",
			"92d2516263ebb61c30247f710ed300114c67ea889dd2db6e568da994b81aca61",
			"f374957d53b61981da8b5bb9f3cbb068bf2194d31e56adc38491f6b5c3017535",
		}
		for i := range path.PathElements {
			require.Equal(t, path.PathElements[i].Hex(), expected[i])
		}
	})

	t.Run("should work for odd index", func(t *testing.T) {
		tree, err := NewMerkleTree(10, []Element{{1}, {2}, {3}, {4}, {5}}, Element{0}, SHA256Hash)
		require.NoError(t, err)
		path, err := tree.Path(3)
		require.NoError(t, err)
		require.ElementsMatch(t, path.PathIndices, []int{1, 1, 0, 0, 0, 0, 0, 0, 0, 0})
		expected := []string{
			"03",
			"6b51d431df5d7f141cbececcf79edf3dd861c3b4069f0b11661a3eefacbba918",
			"3bcf81da18c3a06070dd115fd97e801192b100b798a397004604356a6da2995a",
			"461aa5c7bcac617fa44a126ed3a812b00f963c7fd7344113da0bd452024d81fb",
			"96405940c97198beddd8ec086d669c69e055f8e7288b962258e4cb61b6619dc2",
			"ee38900ad320e5e188c66c3ce41cea0553a1ec098876f790cee0b83df0531233",
			"b75f31fc6b41e45d67ac0d50ed3cae866053735154eb901fae40883ef7c56fb6",
			"fb4e521163400391274a5b2755591aaa5af4ddf32e099aa36aca5ff41f8e31bb",
			"92d2516263ebb61c30247f710ed300114c67ea889dd2db6e568da994b81aca61",
			"f374957d53b61981da8b5bb9f3cbb068bf2194d31e56adc38491f6b5c3017535",
		}
		for i := range path.PathElements {
			require.Equal(t, path.PathElements[i].Hex(), expected[i])
		}
	})

	t.Run("should fail on incorrect index", func(t *testing.T) {
		tree, err := NewMerkleTree(10, []Element{{1}, {2}, {3}, {4}, {5}}, Element{0}, SHA256Hash)
		require.NoError(t, err)
		_, err = tree.Path(-1)
		require.Error(t, err)

		_, err = tree.Path(5)
		require.Error(t, err)

	})
}

//68f30f28e64ca789f80b0108d36f38deae22764b485ba57cbb1b486bb1711ff7
// import { MerkleTree, PartialMerkleTree, TreeEdge } from '../src'
// import { assert, should } from 'chai'
// import { createHash } from 'crypto'
// import { it } from 'mocha'

// const sha256Hash = (left, right) => createHash('sha256').update(`${left}${right}`).digest('hex')
// const ZERO_ELEMENT = '21663839004416932945382355908790599225266501822907911457504978515578255421292'

// describe('MerkleTree', () => {

//   describe('#constructor', () => {

//   describe('#proof', () => {
//     it('should return proof for leaf', () => {
//       const tree = new MerkleTree(10, [1, 2, 3, 4, 5])
//       assert.deepEqual(tree.proof(4), tree.path(3))
//     })
//   })

//   describe('#getTreeEdge', () => {
//     it('should return correct treeEdge', () => {
//       const expectedEdge: TreeEdge = {
//         edgePath: {
//           pathElements: [
//             5,
//             '1390935134112885103361924701261056180224',
//             '1952916572242076545231119328171167580160',
//             '938972308169430750202858820582946897920',
//           ],
//           pathIndices: [0, 0, 1, 0],
//           pathPositions: [5, 0, 0, 0],
//           pathRoot: '3283298202329284319899364273680487022592',
//         },
//         edgeElement: 4,
//         edgeIndex: 4,
//         edgeElementsCount: 6,
//       }
//       const tree = new MerkleTree(4, [0, 1, 2, 3, 4, 5])
//       assert.deepEqual(tree.getTreeEdge(4), expectedEdge)
//     })
//     it('should fail if element not found', () => {
//       const tree = new MerkleTree(4, [0, 1, 2, 3, 4, 5])
//       const call = () => tree.getTreeEdge(6)
//       should().throw(call, 'Element not found')
//     })
//   })
//   describe('#getTreeSlices', () => {
//     let fullTree: MerkleTree
//     before(async () => {
//       const elements = Array.from({ length: 2 ** 10 }, (_, i) => i)
//       fullTree = new MerkleTree(10, elements)
//       return Promise.resolve()
//     })
//     it('should return correct slices count', () => {
//       const count = 4
//       const slicesCount = fullTree.getTreeSlices(4).length
//       should().equal(count, slicesCount)
//     }).timeout(10000)

//     it('should be able to create partial tree from last slice', () => {
//       const [, , , lastSlice] = fullTree.getTreeSlices()
//       const partialTree = new PartialMerkleTree(10, lastSlice.edge, lastSlice.elements)
//       assert.deepEqual(fullTree.root, partialTree.root)
//     }).timeout(10000)

//     it('should be able to build full tree from slices', () => {
//       const slices = fullTree.getTreeSlices()
//       const lastSlice = slices.pop()
//       const partialTree = new PartialMerkleTree(10, lastSlice.edge, lastSlice.elements)
//       slices.reverse().forEach(({ edge, elements }) => {
//         partialTree.shiftEdge(edge, elements)
//       })
//       assert.deepEqual(fullTree.layers, partialTree.layers)
//     }).timeout(10000)

//     it('should return same path', () => {
//       const slices = fullTree.getTreeSlices()
//       const lastSlice = slices.pop()
//       const partialTree = new PartialMerkleTree(10, lastSlice.edge, lastSlice.elements)
//       slices.reverse().forEach(({ edge, elements }) => {
//         partialTree.shiftEdge(edge, elements)
//       })
//       assert.deepEqual(fullTree.path(100), partialTree.path(100))
//     }).timeout(10000)

//     it('should throw if invalid number of elements', () => {
//       const [firstSlice] = fullTree.getTreeSlices()
//       const call = () => new PartialMerkleTree(10, firstSlice.edge, firstSlice.elements)
//       should().throw(call, 'Invalid number of elements')
//     }).timeout(10000)
//   })
//   describe('#getters', () => {
//     const elements = [1, 2, 3, 4, 5]
//     const layers = [
//       [1, 2, 3, 4, 5],
//       [
//         '4027992409016347597424110157229339967488',
//         '923221781152860005594997320673730232320',
//         '752191049236692618445397735417537626112',

//       ],
//       [
//         '81822854828781486047086122479545722339328',
//         '3591172241203040147397382471352592629760',

//       ],
//       ['2729943778107054496417267081388406865920'],
//       ['4562739390655416913642128116127918718976'],
//     ]

//     it('should return same elements in array', () => {
//       const tree = new MerkleTree(10, elements)
//       assert.deepEqual(tree.elements, elements)
//     })
//     it('should return copy of elements array', () => {
//       const tree = new MerkleTree(10, elements)
//       const elements1 = tree.elements
//       tree.insert(6)
//       const elements2 = tree.elements
//       should().not.equal(elements1, elements2)
//     })

//     it('should return same layers in array', () => {
//       const tree = new MerkleTree(4, elements)
//       assert.deepEqual(tree.layers, layers)
//     })
//     it('should return copy of elements array', () => {
//       const tree = new MerkleTree(4, elements)
//       const layers1 = tree.layers
//       tree.insert(6)
//       const layers2 = tree.layers
//       should().not.equal(layers1, layers2)
//     })
//     it('should return correct zeros array', () => {
//       const zeros = [
//         0,
//         '1390935134112885103361924701261056180224',
//         '3223901263414086620636498663535535980544',
//         '938972308169430750202858820582946897920',
//         '3743880566844110745576746962917825445888',
//       ]
//       const tree = new MerkleTree(4, [])
//       assert.deepEqual(tree.zeros, zeros, 'Not equal')
//     })
//     it('should return copy of zeros array', () => {
//       const tree = new MerkleTree(4, [])
//       const zeros1 = tree.zeros
//       tree.insert(6)
//       const zeros2 = tree.zeros
//       should().not.equal(zeros1, zeros2)
//     })
//   })

//   describe('#serialize', () => {
//     it('should work', () => {
//       const src = new MerkleTree(10, [1, 2, 3, 4, 5, 6, 7, 8, 9])
//       const data = src.serialize()
//       const dst = MerkleTree.deserialize(data)
//       should().equal(src.root, dst.root)

//       src.insert(10)
//       dst.insert(10)

//       should().equal(src.root, dst.root)
//     })
//   })
//   describe('#toString', () => {
//     it('should return correct stringified representation', () => {
//       const src = new MerkleTree(10, [1, 2, 3, 4, 5, 6, 7, 8, 9])
//       const str = src.toString()
//       const dst = MerkleTree.deserialize(JSON.parse(str))
//       should().equal(src.root, dst.root)

//       src.insert(10)
//       dst.insert(10)

//       should().equal(src.root, dst.root)

//     })
//   })
// })
