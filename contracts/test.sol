pragma solidity >=0.7.0 <0.8.0;

contract TestContract {
    uint256 public totalSupply;

    function setValid(uint256 number) public {
        totalSupply += number;
    }

    function setInvalid(uint256 number) public {
        totalSupply += number;
        revert("whoops");
    }
    receive() external payable {
        if (totalSupply == 2) {
            revert("shit");
        }
        totalSupply = totalSupply + 1;
    }
}