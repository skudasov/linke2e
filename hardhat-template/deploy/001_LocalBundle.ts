import {DeployFunction} from "hardhat-deploy/types";
import {HardhatRuntimeEnvironment} from "hardhat/types";
import '@nomiclabs/hardhat-web3';

const deployLink = async (hre: HardhatRuntimeEnvironment) => {
  // @ts-ignore
  const {deploy} = hre.deployments;
  // @ts-ignore
  const [deployer] = await hre.ethers.getSigners()

  console.log("----------------------------------------------------")
  console.log('Deploying LinkToken');
  const link = await deploy('MockLink', {
    from: deployer.address,
    gasLimit: 4000000,
  });
  console.log("LinkToken deployed to: ", link.address)
  return link
}

const deployOracle = async (hre: HardhatRuntimeEnvironment, link: any) => {
  // @ts-ignore
  const {deploy} = hre.deployments;
  // @ts-ignore
  const [deployer] = await hre.ethers.getSigners()

  console.log("----------------------------------------------------")
  console.log('Deploying MockOracle');
  const oracle = await deploy('MockOracle', {
    from: deployer.address,
    gasLimit: 4000000,
    args: [link.address],
  });
  console.log("MockOracle deployed to: ", oracle.address)
  return oracle.address
}

const deployAPIConsumer = async (hre: HardhatRuntimeEnvironment, link: any) => {
  // @ts-ignore
  const {deploy} = hre.deployments;
  // @ts-ignore
  const [deployer] = await hre.ethers.getSigners()

  console.log("----------------------------------------------------")
  console.log('Deploying APIConsumer');
  const consumer = await deploy('APIConsumer', {
    from: deployer.address,
    gasLimit: 4000000,
    args: [link.address],
  });
  console.log("APIConsumer deployed to: ", consumer.address)
  return consumer.address
}

const fundLinkContract = async (hre: any, link: any, apiConsumerAddr: any) => {
  console.log(`Funding contract: ${apiConsumerAddr}`)
  //Fund with 10 LINK token
  const amount = '10000000000000000000'
  const accounts = await hre.ethers.getSigners()
  const signer = accounts[0]

  // Create connection to LINK token contract and initiate the transfer
  const linkTokenContract = new hre.ethers.Contract(link.address, link.abi, signer)
  const result = await linkTokenContract.transfer(apiConsumerAddr, amount)
  console.log(`Contract ${apiConsumerAddr} funded with 10 LINK. Transaction hash: ${result.hash}`)
}

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const link = await deployLink(hre)
  const oracleAddr = await deployOracle(hre, link)
  const APIConsumerAddr = await deployAPIConsumer(hre, link)
  await fundLinkContract(hre, link, APIConsumerAddr)
}

export default func
