const path  = require('path');

const solc = require('solc');
const fs = require('fs-extra') ; 

const buildPath = path.resolve(__dirname , 'build') ; 


fs.removeSync(buildPath) ; 

const contractPath = path.resolve(__dirname , 'contracts' , 'Decentify.sol' ) ; 

const source = fs.readFileSync(contractPath , 'utf8') ;

var input = {
    language: 'Solidity',
    sources: {
        'Decentify.sol' : {
            content: source
        }
    },
    settings: {
        outputSelection: {
            '*': {
                '*': [ '*' ]
            }
        }
    }
}; 


const output = JSON.parse(solc.compile(JSON.stringify(input))) ; 

fs.ensureDirSync(buildPath)  ; 

for(let it in output.contracts){
    
}

console.log(output);