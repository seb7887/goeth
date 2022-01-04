import { ethers } from "hardhat"
import { expect } from "chai"

describe("Token", () => {
    let token: any

    before(async () => {
        const Token = await ethers.getContractFactory("Token")
        token = await Token.deploy("Token", "TKN", 31337)
        await token.deployed()
    })

    it("sets name and symbol when created", async () => {
        expect(await token.name()).to.equal("Token")
        expect(await token.symbol()).to.equal("TKN")
    })
})