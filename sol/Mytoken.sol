// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;


import "./IERC20.sol";


contract MyToken is IERC20 {
    mapping (address => uint256) private _balances;
    mapping (address => mapping (address => uint256)) private _allowed;
    uint256 private _totalSupply;

    address private _owner;

    string public symbol;

    constructor(string memory _symbol) {
        symbol = _symbol;
        _owner = msg.sender;
    }

    function totalSupply() override public view returns (uint256) {
        return _totalSupply;
    }

    function balanceOf(address owner) override public view returns (uint256) {
        return _balances[owner];
    }

    function approve(address spender, uint256 value) override public returns (bool) {
        require(spender != address(0));

        _allowed[msg.sender][spender] = value;

        emit Approval(msg.sender, spender, value);
        return true;
    }

    function allowance(address owner, address spender) override public view returns (uint256) {
        return _allowed[owner][spender];
    }

    function transferFrom(address from, address to, uint256 value) override public returns (bool) {
        require(value <= _balances[from]);
        require(value <= _allowed[from][msg.sender]);

        _balances[from] = _balances[from] - value;
        _balances[to] = _balances[to] + value;
        _allowed[from][msg.sender] = _allowed[from][msg.sender] - value;
        return true;
    }

    function transfer(address to, uint256 value) override public returns (bool) {
        require(value <= _balances[msg.sender]);
        require(to != address(0));

        _balances[msg.sender] = _balances[msg.sender] - value;
        _balances[to] = _balances[to] + value;
        emit Transfer(msg.sender, to, value);
        return true;
    }

    function mint(address to, uint256 value) public {
        require(msg.sender == _owner);
        _totalSupply = _totalSupply + value;
        _balances[to] = _balances[to] + value;
        emit Transfer(msg.sender, to, value);
    }
}