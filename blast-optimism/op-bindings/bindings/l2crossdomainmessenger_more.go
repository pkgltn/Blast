// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1002,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1003,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_1_0_1600\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1004,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_51_0_20\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_52_0_1568\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1006,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_101_0_1\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1007,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_102_0_1568\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1008,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_151_0_32\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_152_0_1568\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1010,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1011,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1012,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1013,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":1014,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":1015,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1016,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"discountedValues\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_mapping(t_bytes32,t_uint256)\"},{\"astId\":1017,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"208\",\"type\":\"t_array(t_uint256)41_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)41_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[41]\",\"numberOfBytes\":\"1312\",\"base\":\"t_uint256\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L2CrossDomainMessengerDeployedBin = "0x60806040526004361061015f5760003560e01c806383a74074116100c0578063a711986911610074578063b28ade2511610059578063b28ade25146103f8578063d764ad0b14610418578063ecc704281461042b57600080fd5b8063a711986914610395578063b1b1b209146103c857600080fd5b80639fa0bd8b116100a55780639fa0bd8b146102e65780639fce812c14610321578063a4e7f8bd1461035557600080fd5b806383a74074146102cf5780638cbeeef2146101fe57600080fd5b80634c1d6a69116101175780635644cfdf116100fc5780635644cfdf1461026a5780636e296e45146102805780638129fc1c146102ba57600080fd5b80634c1d6a69146101fe57806354fd4d501461021457600080fd5b80632828d7e8116101485780632828d7e8146101ac5780633dbb202b146101c15780633f827a5a146101d657600080fd5b8063028f85f7146101645780630c56849814610197575b600080fd5b34801561017057600080fd5b50610179601081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b3480156101a357600080fd5b50610179603f81565b3480156101b857600080fd5b50610179604081565b6101d46101cf3660046117e2565b610482565b005b3480156101e257600080fd5b506101eb600181565b60405161ffff909116815260200161018e565b34801561020a57600080fd5b50610179619c4081565b34801561022057600080fd5b5061025d6040518060400160405280600581526020017f312e372e3000000000000000000000000000000000000000000000000000000081525081565b60405161018e91906118b2565b34801561027657600080fd5b5061017961138881565b34801561028c57600080fd5b506102956106e6565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161018e565b3480156102c657600080fd5b506101d46107d2565b3480156102db57600080fd5b5061017962030d4081565b3480156102f257600080fd5b506103136103013660046118cc565b60cf6020526000908152604090205481565b60405190815260200161018e565b34801561032d57600080fd5b506102957f000000000000000000000000000000000000000000000000000000000000000081565b34801561036157600080fd5b506103856103703660046118cc565b60ce6020526000908152604090205460ff1681565b604051901515815260200161018e565b3480156103a157600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610295565b3480156103d457600080fd5b506103856103e33660046118cc565b60cb6020526000908152604090205460ff1681565b34801561040457600080fd5b506101796104133660046118e5565b610a21565b6101d4610426366004611939565b610a8f565b34801561043757600080fd5b5061031360cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b6105bb7f00000000000000000000000000000000000000000000000000000000000000006104b1858585610a21565b347fd764ad0b0000000000000000000000000000000000000000000000000000000061051d60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c6040516024016105399796959493929190611a04565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526113ba565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a33858561064060cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b86604051610652959493929190611a63565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2153016107b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084015b60405180910390fd5b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b6000546003907501000000000000000000000000000000000000000000900460ff16158015610820575060005460ff8083167401000000000000000000000000000000000000000090920416105b6108ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016107ac565b600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff60ff84167401000000000000000000000000000000000000000002167fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff909116177501000000000000000000000000000000000000000000179055610934611448565b6040517f4c802f3800000000000000000000000000000000000000000000000000000000815273430000000000000000000000000000000000000290634c802f389061098e90309060019060009061dead90600401611ae0565b600060405180830381600087803b1580156109a857600080fd5b505af11580156109bc573d6000803e3d6000fd5b5050600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff169055505060405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b6000611388619c4080603f610a3d604063ffffffff8816611b6c565b610a479190611b9c565b610a52601088611b6c565b610a5f9062030d40611bea565b610a699190611bea565b610a739190611bea565b610a7d9190611bea565b610a879190611bea565b949350505050565b60f087901c60028110610b4a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206f722031206d657373616765732061726520737570706f7274656460648201527f20617420746869732074696d6500000000000000000000000000000000000000608482015260a4016107ac565b8061ffff16600003610c3f576000610b9b878986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508f9250611521915050565b600081815260cb602052604090205490915060ff1615610c3d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c6179656400000000000000000060648201526084016107ac565b505b6000610c85898989898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061154092505050565b9050600073ffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef330181167f000000000000000000000000000000000000000000000000000000000000000090911603610d3657863411158015610d045750861580610d045750600034115b610d1057610d10611c16565b600082815260ce602052604090205460ff1615610d2f57610d2f611c16565b5034610e99565b3415610dea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a4016107ac565b600082815260ce602052604090205460ff16610e88576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c617965640000000000000000000000000000000060648201526084016107ac565b50600081815260cf60205260409020545b610ea288611563565b15610f55576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a4016107ac565b600082815260cb602052604090205460ff1615610ff4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c617965640000000000000000000060648201526084016107ac565b61101586611006611388619c40611bea565b67ffffffffffffffff166115ec565b158061103b575060cc5473ffffffffffffffffffffffffffffffffffffffff1661dead14155b1561116657600082815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a2600082815260cf602052604090208190557fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff320161115e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d6573736167650000000000000000000000000000000000000060648201526084016107ac565b5050506113b1565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8b1617905560006111f789619c405a6111ba9190611c45565b8489898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061160a92505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead1790559050801561128e57600083815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555184917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a26113ac565b600083815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555184917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a2600083815260cf602052604090208290557fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016113ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d6573736167650000000000000000000000000000000000000060648201526084016107ac565b505050505b50505050505050565b6040517fc2b3e5ac0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000169063c2b3e5ac90849061141090889088908790600401611c5c565b6000604051808303818588803b15801561142957600080fd5b505af115801561143d573d6000803e3d6000fd5b505050505050505050565b6000547501000000000000000000000000000000000000000000900460ff166114f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016107ac565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055565b600061152f85858585611624565b805190602001209050949350505050565b60006115508787878787876116bd565b8051906020012090509695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff82163014806115b2575073ffffffffffffffffffffffffffffffffffffffff8216734200000000000000000000000000000000000016145b806115e6575073ffffffffffffffffffffffffffffffffffffffff8216734300000000000000000000000000000000000002145b92915050565b600080603f83619c4001026040850201603f5a021015949350505050565b600080600080845160208601878a8af19695505050505050565b60608484848460405160240161163d9493929190611ca4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b60608686868686866040516024016116da96959493929190611cee565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461178057600080fd5b919050565b60008083601f84011261179757600080fd5b50813567ffffffffffffffff8111156117af57600080fd5b6020830191508360208285010111156117c757600080fd5b9250929050565b803563ffffffff8116811461178057600080fd5b600080600080606085870312156117f857600080fd5b6118018561175c565b9350602085013567ffffffffffffffff81111561181d57600080fd5b61182987828801611785565b909450925061183c9050604086016117ce565b905092959194509250565b6000815180845260005b8181101561186d57602081850181015186830182015201611851565b8181111561187f576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006118c56020830184611847565b9392505050565b6000602082840312156118de57600080fd5b5035919050565b6000806000604084860312156118fa57600080fd5b833567ffffffffffffffff81111561191157600080fd5b61191d86828701611785565b90945092506119309050602085016117ce565b90509250925092565b600080600080600080600060c0888a03121561195457600080fd5b873596506119646020890161175c565b95506119726040890161175c565b9450606088013593506080880135925060a088013567ffffffffffffffff81111561199c57600080fd5b6119a88a828b01611785565b989b979a50959850939692959293505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a0830152611a5660c0830184866119bb565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff86168152608060208201526000611a936080830186886119bb565b905083604083015263ffffffff831660608301529695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff8581168252608082019060038610611b0f57611b0f611ab1565b85602084015260028510611b2557611b25611ab1565b84604084015280841660608401525095945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681851681830481118215151615611b9357611b93611b3d565b02949350505050565b600067ffffffffffffffff80841680611bde577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600067ffffffffffffffff808316818516808303821115611c0d57611c0d611b3d565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600082821015611c5757611c57611b3d565b500390565b73ffffffffffffffffffffffffffffffffffffffff8416815267ffffffffffffffff83166020820152606060408201526000611c9b6060830184611847565b95945050505050565b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525060806040830152611cdd6080830185611847565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a0830152611d3960c0830184611847565b9897505050505050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2CrossDomainMessengerStorageLayoutJSON), L2CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2CrossDomainMessenger"] = L2CrossDomainMessengerStorageLayout
	deployedBytecodes["L2CrossDomainMessenger"] = L2CrossDomainMessengerDeployedBin
}
