import web3 from "./Web3";

import FundraiseFactory from "../../../../build/contracts/FunraiseFactory.json";

const instance  = new web3.eth.Contract(
    JSON.parse(FundraiseFactory.abi),
    process.env.addressContract
)

export default instance ; 