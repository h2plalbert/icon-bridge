
const BSHImpl = artifacts.require("BSHImpl");
const BSHProxy = artifacts.require("BSHProxy");
const { deployProxy } = require('@openzeppelin/truffle-upgrades');

const TruffleConfig = require('../truffle-config');
module.exports = async function (deployer, network) {
  if (network !== "development" && network != "soliditycoverage") {
    await deployProxy(BSHProxy, [parseInt(process.env.BSH_TOKEN_FEE)], { deployer });
    await deployProxy(BSHImpl, [process.env.BMC_PERIPHERY_ADDRESS, BSHProxy.address, process.env.BSH_SERVICE], { deployer });
    const bshProxy = await BSHProxy.deployed();
    await bshProxy.updateBSHImplementation(BSHImpl.address);
  }
};
