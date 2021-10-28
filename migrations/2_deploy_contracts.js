const Migrations = artifacts.require("FunraiseFactory");

module.exports = function (deployer) {
  deployer.deploy(Migrations);
};
