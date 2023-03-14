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

package foundation.icon.btp.bsh;

import foundation.icon.btp.bsh.types.Asset;
import foundation.icon.btp.bsh.types.BTPAddress;
import foundation.icon.btp.bsh.types.Balance;
import foundation.icon.btp.bsh.types.ErrorCodes;
import foundation.icon.btp.bsh.types.Token;
import foundation.icon.btp.bsh.types.TransferAsset;
import foundation.icon.btp.restrictions.RestrictionsScoreInterface;
import score.Address;
import score.ArrayDB;
import score.BranchDB;
import score.ByteArrayObjectWriter;
import score.Context;
import score.DictDB;
import score.ObjectReader;
import score.VarDB;
import score.annotation.EventLog;
import score.annotation.External;
import score.annotation.Optional;
import scorex.util.ArrayList;

import java.math.BigInteger;
import java.util.List;
import java.util.Map;

public class ServiceHandler {
    final static String RLPn = "RLPn";
    final static String _svc = "TokenBSH";
    private static final int MIN_OWNER_COUNT = 1;
    private static final int REQUEST_TOKEN_TRANSFER = 0;
    private static final int RESPONSE_HANDLE_SERVICE = 2;
    private static final int RESPONSE_UNKNOWN_ = 3;
    private static final int RC_OK = 0;
    private static final int RC_ERR = 1;
    private static final BigInteger FEE_DENOMINATOR = BigInteger.valueOf(10000);
    private final DictDB<BigInteger, byte[]> pendingDb = Context.newDictDB("pending", byte[].class);
    private final VarDB<BigInteger> serialNo = Context.newVarDB("serialNo", BigInteger.class);
    private final DictDB<String, String> tokenAddrDb = Context.newDictDB("token_addr", String.class);
    private final ArrayDB<String> tokenNameDb = Context.newArrayDB("token_name", String.class);
    private final DictDB<Address, Token> tokenDb = Context.newDictDB("tokens", Token.class);
    private final BranchDB<Address, DictDB<String, Balance>> balanceDB = Context.newBranchDB("balance", Balance.class);
    private final VarDB<Integer> numberOfOwners = Context.newVarDB("numberOfOwners", Integer.class);
    private final VarDB<Address> bmcDb = Context.newVarDB("bmc", Address.class);
    private final DictDB<String, BigInteger> feeCollector = Context.newDictDB("fee_collector", BigInteger.class);
    private final VarDB<String> fromAddr = Context.newVarDB("fromAddr", String.class);
    DictDB<Address, Boolean> ownersDb = Context.newDictDB("owners", Boolean.class);
    private final VarDB<Address> bsrDb = Context.newVarDB("bsr", Address.class);
    VarDB<Boolean> restriction = Context.newVarDB("restricton", Boolean.class);
    RestrictionsScoreInterface restrictionsInterface;
    private final VarDB<String> net = Context.newVarDB("net", String.class);

    public ServiceHandler(String _bmc) {
        // register the BMC link for this BSH
        bmcDb.set(Address.fromString(_bmc));
        serialNo.set(BigInteger.ZERO);
        ownersDb.set(Context.getOwner(), true);
        numberOfOwners.set(1);
        String _btpAddress = (String) Context.call(Address.fromString(_bmc), "getBtpAddress");
        BTPAddress btpAddress = BTPAddress.fromString(_btpAddress);
        net.set(btpAddress.getNet());
    }

    /**
     * @param address address of an owner for BSH to add
     */
    @External
    public void addOwner(Address address) {
        onlyOwner();
        ownersDb.set(address, true);
        numberOfOwners.set(numberOfOwners.get() + 1);
    }

    /**
     * @param address address of an owner for BSH to remove
     */
    @External
    public void removeOwner(Address address) {
        onlyOwner();
        checkMinOwners();
        ownersDb.set(address, false);
        numberOfOwners.set(numberOfOwners.get() - 1);
    }

    private void onlyOwner() {
        assert (Context.getCaller() != null);
        if (!(Context.getOwner().equals(Context.getCaller()) || isAnOwner(Context.getCaller()))) {
            Context.revert(ErrorCodes.BSH_NO_PERMISSION, "No Permission");
        }
    }

    private void onlyBMC() {
        assert (Context.getCaller() != null);
        if (!Context.getCaller().equals(bmcDb.get())) {
            Context.revert(ErrorCodes.BSH_NO_PERMISSION, "No Permission");
        }
    }

    private void checkMinOwners() {
        Context.require(numberOfOwners.get() > MIN_OWNER_COUNT);
    }

    private boolean isAnOwner(Address caller) {
        if (ownersDb.get(caller) != null) {
            return ownersDb.get(caller);
        }
        return false;
    }

    @External(readonly = true)
    public boolean isOwner(Address _addr) {
        return isAnOwner(_addr);
    }

    /**
     * To change the tokenFee setting( Fixed fee and fee percentage)
     *
     * @param _name         name of the token
     * @param _feeNumerator fee numerator to calculate the fee percentage, Set Zero to retain existing value
     * @param _fixedFee     to update the fixed fee
     */
    @External
    public void setFeeRatio(String _name, BigInteger _feeNumerator, BigInteger _fixedFee) {
        onlyOwner();
        Context.require(_feeNumerator.compareTo(BigInteger.ONE) >= 0 &&
                        _feeNumerator.compareTo(FEE_DENOMINATOR) < 0,
                "The feeNumerator should be less than FEE_DENOMINATOR and feeNumerator should be greater than 1");
        String _tokenAddr = tokenAddrDb.get(_name);
        if (tokenAddrDb.get(_name) == null) {
            Context.revert(ErrorCodes.BSH_TOKEN_NOT_REGISTERED, "Token Not Registered");
        }
        Address tokenAddr = Address.fromString(_tokenAddr);
        Token _token = tokenDb.get(tokenAddr);
        _token.setFeeNumerator(_feeNumerator);
        _token.setFixedFee(_fixedFee);
        tokenDb.set(tokenAddr, _token);
    }

    /**
     * @param name    name of the token
     * @param address Address of the token contract
     */
    @External
    public void register(Address address, String name, String symbol, BigInteger decimals, BigInteger feeNumerator, BigInteger fixedFee) {
        onlyOwner();
        if (tokenAddrDb.get(name) != null) {
            Context.revert(ErrorCodes.BSH_TOKEN_EXISTS, "Token with same name exists already.");
        }
        if (feeNumerator.compareTo(BigInteger.ONE) < 0 && feeNumerator.compareTo(FEE_DENOMINATOR) >= 0) {
            Context.revert(ErrorCodes.BSH_INVALID_AMOUNT, "The feeNumerator should be less than FEE_DENOMINATOR and feeNumerator should be greater than 1");
        }
        tokenAddrDb.set(name, address.toString());
        tokenNameDb.add(name);
        tokenDb.set(address, new Token(name, symbol, decimals, feeNumerator, fixedFee));
    }

    @External(readonly = true)
    public String[] tokenNames() {
        String[] tokenNames = new String[tokenNameDb.size()];
        for (int i = 0; i < tokenNameDb.size(); i++) {
            tokenNames[i] = tokenNameDb.get(i);
        }
        return tokenNames;
    }

    @External
    public void withdraw(String _tokenName, BigInteger _value) {
        if (_value.compareTo(BigInteger.ZERO) <= 0) {
            Context.revert(ErrorCodes.BSH_INVALID_AMOUNT, "_value must be positive");
        }
        String tokenAddr = this.tokenAddrDb.getOrDefault(_tokenName, null);
        if (tokenAddr == null) {
            Context.revert(ErrorCodes.BSH_TOKEN_NOT_REGISTERED, "Token not registered");
        }
        Address owner = Context.getCaller();
        Balance balance = getBalance(owner, _tokenName);
        if (balance.getRefundable().compareTo(_value) < 0) {
            Context.revert(ErrorCodes.BSH_OVERDRAWN, "imbalance");
        }
        setBalance(owner, _tokenName, BigInteger.ZERO, BigInteger.ZERO, _value.negate());
        Context.call(Address.fromString(tokenAddr), "transfer", owner, _value, "transfer to Receiver".getBytes());
    }

    /**
     * Works as a deposit function from IRC2 token
     *
     * @param from  Address of the token transfer request sender
     * @param value amount of tokens to transfer
     * @param data  serialized data
     */

    @External
    public void tokenFallback(Address from, BigInteger value, @Optional byte[] data) {
        if (value.compareTo(BigInteger.ZERO) == 0) {
            Context.revert("Invalid Amount");
        }
        Token token = tokenDb.get(Context.getCaller());
        if (token == null) {
            Context.revert("Caller doesnt have a token Registered");
        }

        String _tokenName = token.getName();
        if (_tokenName == null || _tokenName.equals("")) {
            Context.revert("Token not registered");
        }
        setBalance(from, _tokenName, value, BigInteger.ZERO, BigInteger.ZERO);
    }

    /**
     * @param tokenName Name of the token to transfer
     * @param to        String Address of the receiver
     * @param value     Amount to transfer
     */
    @External
    public void transfer(String tokenName, BigInteger value, String to) {
        String tokenAddr = this.tokenAddrDb.getOrDefault(tokenName, null);
        if (tokenAddr == null) {
            Context.revert(ErrorCodes.BSH_TOKEN_NOT_REGISTERED, "Token not registered");
        }
        if (value.compareTo(BigInteger.ZERO) == 0) {
            Context.revert(ErrorCodes.BSH_INVALID_AMOUNT, "Invalid amount specified");
        }
        Address sender = Context.getCaller();
        Balance balance = getBalance(sender, tokenName);
        if (balance.getUsable().compareTo(value) == -1) {
            Context.revert(ErrorCodes.BSH_OVERDRAWN, "Overdrawn");
        }
        BTPAddress _to = BTPAddress.fromString(to);
        Token _tk = tokenDb.get(Address.fromString(tokenAddr));
        checkTransferRestrictions(tokenName, sender.toString(), _to.getContract(), value);
        setBalance(sender, tokenName, value.negate(), value, BigInteger.ZERO);
        BigInteger fee = value.multiply(_tk.getFeeNumerator()).divide(FEE_DENOMINATOR).add(_tk.getFixedFee());
        value = value.subtract(fee);
        BigInteger sn = generateSerialNumber();
        List<Asset> assets = new ArrayList<Asset>();
        assets.add(new Asset(tokenName, value, fee));
        byte[] msg = createMessage(REQUEST_TOKEN_TRANSFER, sender.toString(), _to.getContract(), assets);
        putPending(sn, msg);
        Context.call(bmcDb.get(), "sendMessage", _to.getNet(), _svc, serialNo.get(), msg);
        TransferStart(sender, to, sn, encodeToBytes(assets));
    }

    /**
     * Handles BTP Messages from other blockchains, accepts messages only from BMC.
     * If it fails,
     * then BMC will generate a BTP Message that includes error information, then
     * delivered to the source.
     *
     * @param from Network Address of source network / blockchain
     * @param svc  Service Name
     * @param sn   Serial Number of the message
     * @param msg  Serialised byte of service message (from, to, tokenName, value)
     *             in order
     */
    @External
    public void handleBTPMessage(String from, String svc, BigInteger sn, byte[] msg) {
        onlyBMC();
        if (svc.compareTo(_svc) != 0) {
            Context.revert(ErrorCodes.BSH_INVALID_SERVICE, "Invalid Service");
        }
        ObjectReader reader = Context.newByteArrayObjectReader(RLPn, msg);
        reader.beginList();
        int actionType = reader.readInt();
        ObjectReader readerTa = null;
        if (actionType == REQUEST_TOKEN_TRANSFER || actionType == RESPONSE_HANDLE_SERVICE) {
            readerTa = Context.newByteArrayObjectReader(RLPn, reader.readNullable(byte[].class));
        }
        if (actionType == REQUEST_TOKEN_TRANSFER) {
            TransferAsset _ta = TransferAsset.readObject(readerTa);
            Address dataTo = null;
            try {
                dataTo = Address.fromString(_ta.getTo());
            } catch (Exception e) {
                Context.revert(ErrorCodes.BSH_INVALID_ADDRESS_FORMAT, "Invalid Address format");
            }

            int code = RC_OK;
            for (Asset _asset : _ta.getAssets()) {
                String tokenName = _asset.getName();
                BigInteger value = _asset.getValue();
                checkTransferRestrictions(tokenName, _ta.getFrom(), _ta.getTo(), value);
                String tokenAddr = this.tokenAddrDb.getOrDefault(tokenName, null);

                if (tokenAddr != null) {
                    Context.call(Address.fromString(tokenAddr), "transfer", dataTo, value,
                            "transfer to Receiver".getBytes());
                } else {
                    // code = RC_ERR_UNREGISTERED_TOKEN;
                    Context.revert(ErrorCodes.BSH_TOKEN_NOT_REGISTERED, "Unregistered Token");
                }
            }
            // send response message for `req_token_transfer`
            byte[] res = createMessage(RESPONSE_HANDLE_SERVICE, code, "Transfer Success");
            Context.call(bmcDb.get(), "sendMessage", from, svc, sn, res);
            TransferReceived(from, dataTo, sn, encodeToBytes(_ta.getAssets()));
        } else if (actionType == RESPONSE_HANDLE_SERVICE) {
            if (!hasPending(sn)) {
                Context.revert(ErrorCodes.BSH_INVALID_SERIALNO, "Invalid Serial Number");
            }
            readerTa.beginList();
            int code = readerTa.readInt();
            byte[] respMsg = readerTa.readByteArray();
            byte[] pmsg = getPending(sn);
            ObjectReader pmsgReader = Context.newByteArrayObjectReader(RLPn, pmsg);
            pmsgReader.beginList();
            pmsgReader.skip();
            ObjectReader readerTasset = Context.newByteArrayObjectReader(RLPn, pmsgReader.readByteArray());
            TransferAsset pendingMsg = TransferAsset.readObject(readerTasset);
            fromAddr.set(pendingMsg.getFrom());
            Address pmsgFrom = Address.fromString(pendingMsg.getFrom());
            for (int i = 0; i < pendingMsg.getAssets().size(); i++) {
                Asset _asset = pendingMsg.getAssets().get(i);
                String _tokenName = _asset.getName();
                BigInteger _value = _asset.getValue();
                BigInteger _fee = _asset.getFee();
                BigInteger _totalAmount = _value.add(_fee);
                if (code == RC_OK) {
                    setBalance(pmsgFrom, _tokenName, BigInteger.ZERO, _totalAmount.negate(), BigInteger.ZERO);
                    feeCollector.set(_tokenName, feeCollector.getOrDefault(_tokenName, BigInteger.ZERO).add(_fee));
                } else {
                    setBalance(pmsgFrom, _tokenName, BigInteger.ZERO, _totalAmount.negate(), _totalAmount);
                }
            }
            // delete pending message
            deletePending(sn);
            Address _owner = Address.fromString(fromAddr.get());
            TransferEnd(_owner, sn, BigInteger.valueOf(code), respMsg);
        } else if (actionType == RESPONSE_UNKNOWN_) {
            UnknownResponse(from, sn);
        } else {
            byte[] res = createMessage(RESPONSE_UNKNOWN_);
            Context.call(bmcDb.get(), "sendMessage", from, _svc, serialNo.get(), res);
        }
    }

    /**
     * @param src  Source
     * @param svc  Service Name
     * @param sn   Serial Number
     * @param code Error Code
     * @param msg  Serialized Message (from, to, tokenName, value) in order
     */
    @External
    public void handleBTPError(String src, String svc, BigInteger sn, int code, String msg) {
        onlyBMC();
        if (svc.compareTo(_svc) != 0) {
            Context.revert("Invalid Service name");
        }
        if (!hasPending(sn)) {
            Context.revert(ErrorCodes.BSH_INVALID_SERIALNO, "Invalid Serial Number");
        }
        byte[] pmsg = getPending(sn);
        ObjectReader reader = Context.newByteArrayObjectReader(RLPn, pmsg);
        reader.beginList();
        int actionType = reader.readInt();
        // Rollback token transfer
        if (actionType == REQUEST_TOKEN_TRANSFER) {
            ObjectReader readerTasset = Context.newByteArrayObjectReader(RLPn, reader.readByteArray());
            TransferAsset _ta = TransferAsset.readObject(readerTasset);
            Address from = Address.fromString(_ta.getFrom());
            Address to = Address.fromString(_ta.getTo());
            for (Asset _asset : _ta.getAssets()) {
                String tokenName = _asset.getName();
                BigInteger value = _asset.getValue();
                BigInteger fee = _asset.getValue();
                BigInteger _totalAmount = value.add(fee);
                readerTasset.end();
                //refund the user after failure
                setBalance(from, tokenName, BigInteger.ZERO, _totalAmount.negate(), _totalAmount);
            }
        }
        reader.end();
        // delete pending message
        deletePending(sn);
    }

    /**
     * Returns the Accumulated fees for all the assets
     */

    @External(readonly = true)
    public List<Map<String, BigInteger>> getAccumulatedFees() {
        List<Map<String, BigInteger>> tokens = new ArrayList<>();
        for (int i = 0; i < tokenNameDb.size(); i++) {
            if (feeCollector.getOrDefault(tokenNameDb.get(i), BigInteger.ZERO).compareTo(BigInteger.ZERO) != 0) {
                tokens.add(Map.of(tokenNameDb.get(i), feeCollector.get(tokenNameDb.get(i))));
            }
        }
        return tokens;
    }

    /**
     * @param _fa  Fee Aggregation address
     * @param _svc Service Name
     */
    @External
    public void handleFeeGathering(String _fa, String _svc) {
        onlyBMC();
        BTPAddress _addr = BTPAddress.fromString(_fa);
        List<Asset> _assets = new ArrayList<Asset>();
        for (int i = 0; i < tokenNameDb.size(); i++) {
            if (feeCollector.getOrDefault(tokenNameDb.get(i), BigInteger.ZERO).compareTo(BigInteger.ZERO) != 0) {
                Asset _asset = new Asset(tokenNameDb.get(i), feeCollector.get(tokenNameDb.get(i)), BigInteger.ZERO);
                _assets.add(_asset);
            }
        }
        if (_assets.size() > 0) {
            if (_addr.getNet().equals(net.get())) {
                for (Asset _asset : _assets) {
                    Address fa = Address.fromString(_addr.getContract());
                    String tokenName = _asset.getName();
                    BigInteger value = _asset.getValue();
                    String tokenAddr = this.tokenAddrDb.getOrDefault(tokenName, null);
                    if (tokenAddr != null) {
                        Context.call(Address.fromString(tokenAddr), "transfer", fa, value,
                                "transfer to Receiver".getBytes());
                        // feeCollector.set(tokenName, BigInteger.ZERO);
                    }
                }
            } else {
                BigInteger sn = generateSerialNumber();
                byte[] msg = createMessage(REQUEST_TOKEN_TRANSFER, Context.getAddress().toString(), _addr.getContract(),
                        _assets);
                putPending(sn, msg);
                Context.call(bmcDb.get(), "sendMessage", _addr.getNet(), _svc, serialNo.get(), msg);
                for (Asset _asset : _assets) {
                    String tokenName = _asset.getName();
                    feeCollector.set(tokenName, BigInteger.ZERO);
                }
                TransferStart(Context.getAddress(), _fa, sn, encodeToBytes(_assets));
            }
        }
    }

    @External(readonly = true)
    public Balance getBalance(Address user, String tokenName) {
        Balance defaultBal = new Balance(BigInteger.ZERO, BigInteger.ZERO, BigInteger.ZERO);
        return balanceDB.at(user).getOrDefault(tokenName, defaultBal);
    }

    @External
    public void setBalance(Address user, String tokenName, BigInteger usable, BigInteger locked,
                           BigInteger refundable) {
        Balance balanceBefore = getBalance(user, tokenName);
        Balance newBalance = new Balance(balanceBefore.getUsable().add(usable), balanceBefore.getLocked().add(locked),
                balanceBefore.getRefundable().add(refundable));
        balanceDB.at(user).set(tokenName, newBalance);
    }

    @External
    public void addRestrictor(Address _address) {
        onlyOwner();
        bsrDb.set(_address);
        restriction.set(true);
        restrictionsInterface = new RestrictionsScoreInterface(_address);
    }

    @External
    public void disableRestrictions() {
        onlyOwner();
        restriction.set(false);
    }

    private void checkTransferRestrictions(String _tokenName, String _from, String _to, BigInteger _value) {
        if (restriction.get() != null && restriction.get() && bsrDb.get() != null) {
            // restictonsInterface.validateRestriction(_tokenName, _from, _to, _value);
            Context.call(bsrDb.get(), "validateRestriction", _tokenName, _from, _to, _value);
        }
    }

    private byte[] createMessage(int type, Object... args) {
        ByteArrayObjectWriter writer = Context.newByteArrayObjectWriter(RLPn);
        if (type == REQUEST_TOKEN_TRANSFER) {
            writer.beginList(2);
            writer.write(REQUEST_TOKEN_TRANSFER);// ServiceType

            ByteArrayObjectWriter writerTa = Context.newByteArrayObjectWriter(RLPn);
            TransferAsset _ta = new TransferAsset((String) args[0], (String) args[1], (List<Asset>) args[2]);
            TransferAsset.writeObject(writerTa, _ta);

            writer.write(writerTa.toByteArray());// data
            writer.end();
        } else if (type == RESPONSE_HANDLE_SERVICE) {
            writer.beginList(2);
            writer.write(RESPONSE_HANDLE_SERVICE);// ServiceType

            ByteArrayObjectWriter writerResponse = Context.newByteArrayObjectWriter(RLPn);
            writerResponse.beginList(2);
            writerResponse.write((int) args[0]);// Code
            writerResponse.write((String) args[1]);// Message
            writerResponse.end();

            writer.write(writerResponse.toByteArray()); // data
            writer.end();
        } else if (type == RESPONSE_UNKNOWN_) {
            writer.beginList(1);
            writer.write(RESPONSE_UNKNOWN_);// ServiceType
            writer.end();
        }
        return writer.toByteArray();
    }

    private void putPending(BigInteger sn, byte[] msg) {
        pendingDb.set(sn, msg);
    }

    private boolean hasPending(BigInteger sn) {
        return pendingDb.get(sn) != null;
    }

    private void deletePending(BigInteger sn) {
        pendingDb.set(sn, null);
    }

    private byte[] getPending(BigInteger sn) {
        return pendingDb.get(sn);
    }

    private BigInteger generateSerialNumber() {
        BigInteger newSnNo = serialNo.getOrDefault(BigInteger.ZERO).add(BigInteger.ONE);
        serialNo.set(newSnNo);
        return newSnNo;
    }

    static byte[] encodeToBytes(List<Asset> assets) {
        ByteArrayObjectWriter writer = Context.newByteArrayObjectWriter("RLPn");
        writer.beginList(assets.size());
        for (Asset v : assets) {
            Asset.writeObject(writer, v);
        }
        writer.end();
        return writer.toByteArray();
    }

    @EventLog(indexed = 1)
    public void TransferStart(Address _from, String _to, BigInteger _sn, byte[] _assetDetails) {
    }

    @EventLog(indexed = 1)
    protected void TransferEnd(Address _from, BigInteger _sn, BigInteger _code, byte[] _msg) {
    }

    @EventLog(indexed = 1)
    protected void TransferReceived(String _from, Address _to, BigInteger _sn, byte[] _assetDetails) {
    }

    @EventLog(indexed = 1)
    public void UnknownResponse(String _from, BigInteger _sn) {
    }
}