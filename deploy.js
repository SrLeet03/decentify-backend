
function func() {
    var contract;
    const contractAddress = "0x9abd5ea4f8fd7194efa51a3435bfceb24acbee31";

    const contractMessage = async () => {
        contract = new web3.eth.Contract(abi, contractAddress);
        let message = await contract.methods.getMessage().call();
        return message;
    }

    // Modern dapp browsers...
    if (window.ethereum) {
        window.web3 = new Web3(ethereum);
        try {
            // Request account access if needed
            await ethereum.enable();
            var message = await contractMessage();
            var elm = document.getElementById("message");
            elm.innerHTML = message;
        } catch (error) {
            // User denied account access...
        }
    }
    // Non-dapp browsers...
    else {
        console.log('Non-Ethereum browser detected. You should consider trying MetaMask!');
    }
}