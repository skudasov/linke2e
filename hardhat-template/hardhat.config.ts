// import fs from 'fs'
import 'dotenv/config'
import 'hardhat-deploy'
import 'hardhat-deploy-ethers'
import tasks from './tasks'
import 'hardhat-abi-exporter'
import 'hardhat-tracer'


// Load tasks
for (const tsk of tasks) { tsk() }

let mnemonic = process.env.MNEMONIC
if (!mnemonic) {
  // FOR DEV ONLY, SET IT IN .env ethconfig if you want to keep it private
  // (IT IS IMPORTANT TO HAVE A NON RANDOM MNEMONIC SO THAT SCRIPTS CAN ACT ON THE SAME ACCOUNTS)
  mnemonic = 'test test test test test test test test test test test junk'
}

const config = {
  abiExporter: {
    path: './abi_export',
    clear: true,
    flat: true,
    spacing: 0
  },
  solidity: {
    compilers: [
      {
        version: "0.7.0"
      },
      {
        version: "0.6.6",
      },
    ]
  },
  defaultNetwork: 'localhost',
  networks: {
    hardhat: {
      loggingEnabled: true,
      mining: {
        // auto: false,
        // interval: 5000
      },
    },
    localhost: {
      url: "http://0.0.0.0:8545",
      accounts: { mnemonic }
    },
    rinkeby: {
      url: `https://rinkeby.infura.io/v3/${process.env.INFURA_TOKEN}`,
      accounts: { mnemonic }
    },
    mainnet: {
      url: `https://mainnet.infura.io/v3/${process.env.INFURA_TOKEN}`,
      accounts: { mnemonic }
    },
    ropsten: {
      url: `https://ropsten.infura.io/v3/${process.env.INFURA_TOKEN}`,
      accounts: { mnemonic }
    },
    goerli: {
      url: `https://goerli.infura.io/v3/${process.env.INFURA_TOKEN}`,
      accounts: { mnemonic }
    },
    xdai: {
      url: 'https://dai.poa.network',
      chainId: 100,
      gas: 500000,
      gasPrice: 1000000000,
      accounts: { mnemonic }
    }
  }
}

export default config

