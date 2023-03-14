// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2021 ICON Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity >=0.5.0 <=0.8.0;
pragma experimental ABIEncoderV2;

import "./Interfaces/IBSHProxy.sol";
import "./Interfaces/IBSHImpl.sol";
import "./Interfaces/IBMCPeriphery.sol";

import "./Libraries/Types.sol";
import "./Libraries/RLPEncodeStruct.sol";
import "./Libraries/RLPDecodeStruct.sol";
import "./Libraries/Strings.sol";
import "./Libraries/ParseAddress.sol";

import "@openzeppelin/contracts-upgradeable/utils/math/SafeMathUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract BSHImpl is IBSHImpl, Initializable {
    using SafeMathUpgradeable for uint256;
    using RLPEncodeStruct for Types.ServiceMessage;
    using RLPEncodeStruct for Types.TransferAssets;
    using RLPEncodeStruct for Types.Response;
    using RLPDecodeStruct for bytes;
    using ParseAddress for address;
    using ParseAddress for string;
    using Strings for string;
    using Strings for uint256;
    IBMCPeriphery private bmc;
    IBSHProxy private bshProxy;
    mapping(uint256 => Types.TransferAssets) public requests;
    uint256 private serialNo;
    string public serviceName;
    uint256 private constant RC_OK = 0;
    uint256 private constant RC_ERR = 1;
    uint256 private numOfPendingRequests;

    event HandleBTPMessageEvent(uint256 _sn, uint256 _code, string _msg);

    event TransferStart(
        address indexed _from,
        string _to,
        uint256 _sn,
        Types.Asset[] _assets
    );

    event TransferEnd(
        address indexed _from,
        uint256 _sn,
        uint256 _code,
        string _response
    );

    event TransferReceived(
        string indexed _from,
        address indexed _to,
        uint256 _sn,
        Types.Asset[] _assetDetails
    );

    modifier onlyBMC() {
        require(msg.sender == address(bmc), "Unauthorized");
        _;
    }

    modifier onlyBSHProxy() {
        require(msg.sender == address(bshProxy), "Unauthorized");
        _;
    }

    function initialize(
        address _bmc,
        address _bshProxy,
        string memory _serviceName
    ) public initializer {
        bmc = IBMCPeriphery(_bmc);
        //bmc.requestAddService(_serviceName, address(this));
        bshProxy = IBSHProxy(_bshProxy);
        serviceName = _serviceName;
        serialNo = 0;
    }

    /**
     @notice Check whether BSHPeriphery has any pending transferring requests
     @return true or false
    */
    function hasPendingRequest() external view override returns (bool) {
        return numOfPendingRequests != 0;
    }

    function handleFeeGathering(string memory _toFA, string memory _svc)
        external
        override
        onlyBMC
    {
        require(_svc.compareTo(serviceName) == true, "Invalid service");
        _toFA.splitBTPAddress();
        bshProxy.handleFeeTransfer(_toFA);
    }

    /**   @notice Notify that BSH contract has received unknown response
        The `_from` sender
        The `_sn` sequence number of service message
    */
    event UnknownResponse(string _from, uint256 _sn);

    function handleBTPMessage(
        string calldata _from,
        string calldata _svc,
        uint256 _sn,
        bytes calldata _msg
    ) external override onlyBMC {
        require(_svc.compareTo(serviceName) == true, "Invalid Service Name");
        Types.ServiceMessage memory _sm = _msg.decodeServiceMessage();
        if (_sm.serviceType == Types.ServiceType.REQUEST_TOKEN_TRANSFER) {
            Types.TransferAssets memory _ta = _sm.data.decodeTransferAsset();
            string memory _statusMsg;
            uint256 _status;
            try this.handleRequest(_ta) {
                _statusMsg = "Transfer Success";
                _status = RC_OK;
                emit TransferReceived(
                    _from,
                    _ta.to.parseAddress(),
                    _sn,
                    _ta.asset
                );
            } catch Error(string memory _err) {
                /**
                 * @dev Uncomment revert to debug errors
                 */
                //revert(_err);
                _statusMsg = _err;
                _status = RC_ERR;
            }

            sendBTPResponse(
                Types.ServiceType.RESPONSE_HANDLE_SERVICE,
                _from,
                _sn,
                _statusMsg,
                _status
            );
        } else if (
            _sm.serviceType == Types.ServiceType.RESPONSE_HANDLE_SERVICE
        ) {
            require(bytes(requests[_sn].from).length != 0, "Invalid SN");
            Types.Response memory response = _sm.data.decodeResponse();
            handleResponse(_sn, response.code, response.message);
        } else if (_sm.serviceType == Types.ServiceType.RESPONSE_UNKNOWN) {
            emit UnknownResponse(_from, _sn);
        } else {
            sendBTPResponse(
                Types.ServiceType.RESPONSE_UNKNOWN,
                _from,
                _sn,
                "UNKNOWN_TYPE",
                RC_ERR
            );
        }
    }

    function handleRequest(Types.TransferAssets memory transferAssets)
        external
    {
        require(msg.sender == address(this), "Unauthorized");
        //string memory _toNetwork;
        //string memory _toAddress;
        string memory _toAddress = transferAssets.to;
        //(_toNetwork, _toAddress) = transferAssets.to.splitBTPAddress();

        try this.checkParseAddress(_toAddress) {} catch {
            revert("Invalid Address");
        }
        Types.Asset[] memory _asset = transferAssets.asset;
        for (uint256 i = 0; i < _asset.length; i++) {
            // Check if the _toAddress is invalid
            uint256 _amount = _asset[i].value;
            string memory _tokenName = _asset[i].name;
            // Check if the token is registered already
            require(
                bshProxy.isTokenRegisterd(_tokenName) == true,
                "Unregistered Token"
            );
            try
                bshProxy.handleTransferRequest(
                    _toAddress.parseAddress(),
                    _tokenName,
                    _amount
                )
            {} catch Error(string memory _err) {
                revert(_err);
            }
        }
    }

    function sendBTPResponse(
        Types.ServiceType _serviceType,
        string memory _to,
        uint256 _sn,
        string memory _msg,
        uint256 _code
    ) private {
        bmc.sendMessage(
            _to,
            serviceName,
            _sn,
            Types
                .ServiceMessage(
                    _serviceType,
                    Types.Response(_code, _msg).encodeResponse()
                )
                .encodeServiceMessage()
        );
        emit HandleBTPMessageEvent(_sn, _code, _msg);
    }

    function checkParseAddress(string calldata _to) external pure {
        _to.parseAddress();
    }

    function handleBTPError(
        string calldata _src,
        string calldata _svc,
        uint256 _sn,
        uint256 _code,
        string calldata _msg
    ) external override onlyBMC {
        require(_svc.compareTo(serviceName) == true, "Invalid service");
        require(bytes(requests[_sn].from).length != 0, "Invalid SN");
        handleResponse(_sn, _code, _msg);
    }

    function handleResponse(
        uint256 _sn,
        uint256 _code,
        string memory _msg
    ) private {
        address _caller = requests[_sn].from.parseAddress();
        bshProxy.handleResponse(_caller, requests[_sn].asset, _code);
        delete requests[_sn];
        numOfPendingRequests--;
        emit TransferEnd(_caller, _sn, _code, _msg);
    }

    function sendServiceMessage(
        address _from,
        string memory _to,
        Types.Asset[] memory _assets
    ) external override onlyBSHProxy {
        // Send Service Message to BMC
        string memory _toNetwork;
        string memory _toAddress;
        (_toNetwork, _toAddress) = _to.splitBTPAddress();
        Types.TransferAssets memory _ta = Types.TransferAssets(
            _from.toString(),
            _toAddress,
            _assets
        );
        bytes memory serviceMessage = Types
            .ServiceMessage(
                Types.ServiceType.REQUEST_TOKEN_TRANSFER,
                _ta.encodeTransferAsset()
            )
            .encodeServiceMessage();
        serialNo++;
        bmc.sendMessage(_toNetwork, serviceName, serialNo, serviceMessage);

        /* requests[serialNo] = Types.TransferAssets(
            _from.toString(),
            _toAddress,
            new Types.Asset[](1)
        );*/

        requests[serialNo].from = _from.toString();
        requests[serialNo].to = _to;
        for (uint256 i = 0; i < _assets.length; i++) {
            requests[serialNo].asset.push(_assets[i]);
        }

        numOfPendingRequests++;
        emit TransferStart(_from, _to, serialNo, _assets);
    }
}
