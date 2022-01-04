import { ethers } from "hardhat"

async function main() {
    const Store = await ethers.getContractFactory("Store")
    const store = await Store.deploy("0.1.0")

    await store.deployed()

    console.log("Store contract deployed to address: ", store.address)
}

main()
    .then(() => process.exit(0))
    .catch(err => {
        console.error(err)
        process.exit(1)
    })