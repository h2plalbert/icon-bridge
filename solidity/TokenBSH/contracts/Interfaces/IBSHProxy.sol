// SPDX-License-Identifier: MIT

pragma solidity >=0.5.0 <=0.8.0;
pragma experimental ABIEncoderV2;

import "../Libraries/Types.sol";

/**
   @title Interface of BSHCore contract
   @dev This contract is used to handle coin transferring service
   Note: The coin of following interface can be:
   Native Coin : The native coin of this chain
   Wrapped Native Coin : A tokenized ERC1155 version of another native coin like ICX
*/
interface IBSHProxy {
    /**
       @notice Adding another Onwer.
       @dev Caller must be an Onwer of BTP network
       @param _owner    Address of a new Onwer.
    */
    function addOwner(address _owner) external;

    /**
       @notice Removing an existing Owner.
       @dev Caller must be an Owner of BTP network
       @dev If only one Owner left, unable to remove the last Owner
       @param _owner    Address of an Owner to be removed.
    */
    function removeOwner(address _owner) external;

    /**
       @notice Checking whether one specific address has Owner role.
       @dev Caller can be ANY
       @param _owner    Address needs to verify.
    */
    function isOwner(address _owner) external view returns (bool);

    /**
       @notice Get a list of current Owners
       @dev Caller can be ANY
       @return      An array of addresses of current Owners
    */

    function getOwners() external view returns (address[] memory);

    /**
        @notice update BSH Periphery address.
        @dev Caller must be an Owner of this contract
        _bshImpl Must be different with the existing one.
        @param _bshImpl    BSHPeriphery contract address.
    */
    function updateBSHImplementation(address _bshImpl) external;

    /**
        @notice set fee ratio.
        @dev Caller must be an Owner of this contract
        The transfer fee is calculated by feeNumerator/FEE_DEMONINATOR. 
        The feeNumetator should be less than FEE_DEMONINATOR
        _feeNumerator is set to `10` in construction by default, which means the default fee ratio is 0.1%.
        @param _feeNumerator    the fee numerator
    */
    function setFeeRatio(
        string calldata _name,
        uint256 _feeNumerator,
        uint256 _fixedFee
    ) external;

    /**
        @notice Registers a wrapped coin and id number of a supporting coin.
        @dev Caller must be an Owner of this contract
        _name Must be different with the native coin name.
        @dev '_id' of a wrapped coin is generated by using keccak256
          '_id' = 0 is fixed to assign to native coin
        @param _name    Coin name. 
    */
    function register(
        address _addr,
        string calldata _name,
        string calldata _symbol,
        uint256 _decimals,
        uint256 _feeNumerator,
        uint256 _fixedFee
    ) external;

    /**
        @notice Return a usable/locked/refundable balance of an account based on coinName.
        @return _usableBalance the balance that users are holding.
        @return _lockedBalance when users transfer the coin, 
                it will be locked until getting the Service Message Response.
        @return _refundableBalance refundable balance is the balance that will be refunded to users.
    */
    function getBalanceOf(address _owner, string memory _coinName)
        external
        view
        returns (
            uint256 _usableBalance,
            uint256 _lockedBalance,
            uint256 _refundableBalance
        );

    /**
        @notice Return a list accumulated Fees.
        @dev only return the asset that has Asset's value greater than 0
        @return _accumulatedFees An array of Asset
    */
    function getAccumulatedFees()
        external
        view
        returns (Types.Asset[] memory _accumulatedFees);

    /**
       @notice Allow users to deposit an amount of wrapped native coin `_coinName` from the `msg.sender` address into the BSHCore contract.
       @dev Caller must set to approve that the wrapped tokens can be transferred out of the `msg.sender` account by BSHCore contract.
       It MUST revert if the balance of the holder for token `_coinName` is lower than the `_value` sent.
       @param _coinName    A given name of a wrapped coin 
       @param _value       An amount request to transfer.
       @param _to          Target BTP address.
    */
    function transfer(
        string calldata _coinName,
        uint256 _value,
        string calldata _to
    ) external;

    /**
        @notice Reclaim the token's refundable balance by an owner.
        @dev Caller must be an owner of coin
        The amount to claim must be smaller or equal than refundable balance
        @param _coinName   A given name of coin
        @param _value       An amount of re-claiming tokens
    */
    //function reclaim(string calldata _coinName, uint256 _value) external;

    /**
        @notice return coin for the failed transfer.
        @dev Caller must be itself
        @param _to    account
        @param _coinName    coin name    
        @param _value    the minted amount   
    */
    /*  function refund(
        address _to,
        string calldata _coinName,
        uint256 _value
    ) external;

*/
    /**
        @notice Handle a request of Fee Gathering
        @dev    Caller must be an BSHPeriphery contract
        @param  _fa    BTP Address of Fee Aggregator 
    */
    function handleFeeTransfer(string calldata _fa) external;

    /**
        @notice Handle a response of a requested service
        @dev Caller must be an BSHPeriphery contract
        @param _caller   An address of originator of a requested service
        @param _assets   Type Asset 
        @param _code     Response Code
    */
    function handleResponse(
        address _caller,
        Types.Asset[] memory _assets,
        uint256 _code
    ) external;

    function isTokenRegisterd(string calldata _tokenName)
        external
        view
        returns (bool _registered);

    function handleTransferRequest(
        address _toAddress,
        string calldata _tokenName,
        uint256 _amount
    ) external;
}
