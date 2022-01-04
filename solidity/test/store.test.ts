import { ethers } from "hardhat"
import { expect } from "chai"

const toWei = (value: number) => {
    return ethers.utils.parseEther(value.toString())
}

const fromWei = (value: any) => {
    return ethers.utils.formatEther(
        typeof value === "string" ? value : value.toString()
    )
}

const formatString = (value: string) => {
    return ethers.utils.formatBytes32String(value)
}

describe("Store", () => {
    let owner
    let store: any
    let tx: any

    beforeEach(async () => {
        [owner] = await ethers.getSigners()

        const Store = await ethers.getContractFactory("Store")
        store = await Store.deploy("test")

        await store.deployed()
    })

    it("should emit ItemSet event after setting a new item", async () => {
        tx = store.setItem(formatString("foo"), formatString("bar"))

        await expect(tx).to.emit(store, "ItemSet").withArgs(formatString("foo"), formatString("bar"))
    })
})