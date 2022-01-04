import { ethers } from "hardhat"

async function main() {
    const Store = await ethers.getContractFactory("Store")
    const store = await Store.deploy("0.1.0")

    await store.deployed()

    console.log("Store contract deployed to address: ", store.address)

    const Token = await ethers.getContractFactory("Token")
    const token = await Token.deploy("Token", "TKN", 31337)

    await token.deployed()

    console.log("Token contract deployed to address: ", token.address)
}

main()
    .then(() => process.exit(0))
    .catch(err => {
        console.error(err)
        process.exit(1)
    })