const path  = require('path');

const solc = require('solc');
const fs = require('fs-extra') ; 

const buildPath = path.resolve(__dirname , 'build') ; 


fs.removeSync(buildPath) ; 

const contractPath = path.resolve(__dirname , 'contracts' , 'Decentify.sol' ) ; 

const source = fs.readFileSync(contractPath , 'UTF-8') ;

var input = {
    language: 'Solidity',
    sources: {
        'Decentify.sol' : {
            content: fs.readFileSync('contracts/Decentify.sol', 'utf8'),
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


const output = solc.compile(JSON.stringify(input)) ; 

fs.ensureDirSync(buildPath)  ; 

//output = JSON.parse(output) ; 
const out = JSON.parse(output) ; let BUILD_TARGET_PATH = './build/';
console.log(out.contracts) ; 
        for (let contractJson in out.contracts) {
        //      if (contractJson.startsWith("FunraiseFactory.") ||contractJson.startsWith("Funraiser.") ) {
        //         fs.writeFileSync(BUILD_TARGET_PATH + '/' + contractJson.replace(".sol", ".json"), JSON.stringify(output.contracts[contractJson][contractJson.replace(".sol", "")]), function (err) {
        //              console.log("\nJSON saved   -> OK\n    "+ BUILD_TARGET_PATH + contractJson.replace(".sol", ".json"));
        //              if (err) throw err;
        //          });
        //    }
         }

for(let contract in output){
      fs.outputJSONSync(
        path.resolve(buildPath , contract + '.json' ),
        output[contract]
      );
}

//console.log(JSON.parse(output[':FunraiseFactory']) );


// const solc = require("solc");
// const fs = require("fs");
// const path  = require('path');

// let BUILD_TARGET_PATH = './build/';

// let contractBuild = async () => {

//     // Creates target path in the case it doesn't exists
//     fs.mkdir(BUILD_TARGET_PATH, {recursive: true}, (err) => {
//         if (err) throw err;

//         // loading the source code from a solidity file
//         let input = {
//             language: 'Solidity',
//             sources: {
//                 //
//                 // CertificateManagement.sol
//                 'Decentify.sol': {
//                     content: fs.readFileSync('contracts/Decentify.sol', 'utf8'),
//                 },
                
//             },
//             settings: {
//                 outputSelection: {'*': {'*': ['*']}}
//             }
//         };

//         // compile the solidity code
//         let compiled = solc.compile(JSON.stringify(input));
//         console.log("Compiling    -> OK" );
//         let output = JSON.parse(compiled);

//         // for (let contractJson in output.contracts) {
//         //     if (contractJson.startsWith("Decentify.")) {
//         //         fs.writeFileSync(BUILD_TARGET_PATH + '/' + contractJson.replace(".sol", ".json"), JSON.stringify(output.contracts[contractJson][contractJson.replace(".sol", "")]), function (err) {
//         //             console.log("\nJSON saved   -> OK\n    "+ BUILD_TARGET_PATH + contractJson.replace(".sol", ".json"));
//         //             if (err) throw err;
//         //         });
//         //     }
//         // }

//         // Save bytecode
//         let bytecode = '0x' + output.contracts['Decentify.sol']['Decentify'].evm.bytecode.object;
//         console.log('\nByte Code    -> OK');

//         fs.writeFileSync(BUILD_TARGET_PATH + 'Decentify.bytecode', bytecode, function (err) {
//             console.log("\nCode saved   -> OK\n    "+ BUILD_TARGET_PATH + "Decentify.bytecode");
//             if (err) throw err;
//         });

//         // Logging of methods
//         console.log('\nIndentifiers -> OK\n');
//         let methodIdentifiers = output.contracts['Decentify.sol']['Decentify'].evm.methodIdentifiers;
//         for (let indentifier in methodIdentifiers) {
//             console.log("    " + indentifier);
//         }

//         // save public interface of contract
//         let abi = output.contracts['Decentify.sol']['Decentify'].abi;
//         fs.writeFileSync(BUILD_TARGET_PATH + 'Decentify.abi', JSON.stringify(abi), function (err) {
//             console.log("\nABI saved    -> OK\n    "+ BUILD_TARGET_PATH + "Decentify.abi");
//             console.log(" ");
//             if (err) throw err;
//         });
//     });
// };

// contractBuild();
