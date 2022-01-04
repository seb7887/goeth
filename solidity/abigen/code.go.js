const fs = require("fs")

const [, , package, compiled, contract] = process.argv

const output = JSON.parse(fs.readFileSync(compiled))

function makeCodeFile(contract, path = `contracts/${contract}.sol`) {
    const Contract = output.contracts[`${path}:${contract}`]
    const binRuntime = Contract["bin-runtime"]

    return `// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
package ${package}

const ${contract}DeployedCode = "0x${binRuntime}"`
}

console.log(makeCodeFile(contract))