#!/bin/bash
export CONFIG_DIR=${CONFIG_DIR:-${ICONBRIDGE_CONFIG_DIR}}
export CONTRACTS_DIR=${CONTRACTS_DIR:-${ICONBRIDGE_CONTRACTS_DIR}}
export SCRIPTS_DIR=${SCRIPTS_DIR:-${ICONBRIDGE_SCRIPTS_DIR}}

# GOLOOP Config
GOLOOPCHAIN=${GOLOOPCHAIN:-"goloop"}
export GOLOOP_RPC_URI=$ICON_ENDPOINT
export GOLOOP_RPC_ADMIN_URI=http://$GOLOOPCHAIN/admin/system
export GOLOOP_RPC_KEY_STORE=${ICON_KEY_STORE:-"$CONFIG_DIR/$GOLOOPCHAIN.keystore.json"}
export GOLOOP_RPC_KEY_SECRET=${ICON_SECRET:-"$CONFIG_DIR/$GOLOOPCHAIN.keysecret"}
export GOLOOP_RPC_STEP_LIMIT=${GOLOOP_RPC_STEP_LIMIT:-5000000000}
export GOLOOP_CHAINSCORE=cx000000000000000000000000000000000000000
export GOLOOP_RPC_NID=${GOLOOP_RPC_NID:-$(cat $CONFIG_DIR/nid.icon)}
export ICON_BMC_NET="$GOLOOP_RPC_NID.icon"

export ICON_NATIVE_COIN_SYM=("ICX")
export ICON_NATIVE_COIN_NAME=("btp-$ICON_BMC_NET-ICX")
export ICON_NATIVE_TOKEN_SYM=("sICX" "bnUSD")
export ICON_NATIVE_TOKEN_NAME=("btp-$ICON_BMC_NET-sICX" "btp-$ICON_BMC_NET-bnUSD")
export ICON_WRAPPED_COIN_SYM=("BNB" "BUSD" "USDT" "USDC" "BTCB" "ETH")
export ICON_WRAPPED_COIN_NAME=("btp-$BSC_BMC_NET-BNB" "btp-$BSC_BMC_NET-BUSD" "btp-$BSC_BMC_NET-USDT" "btp-$BSC_BMC_NET-USDC" "btp-$BSC_BMC_NET-BTCB" "btp-$BSC_BMC_NET-ETH")
export FEE_GATHERING_INTERVAL=1000

export ICON_NATIVE_COIN_FIXED_FEE=(4300000000000000000)
export ICON_NATIVE_COIN_FEE_NUMERATOR=(100)
export ICON_NATIVE_COIN_DECIMALS=(18)
export ICON_NATIVE_TOKEN_FIXED_FEE=(3900000000000000000 1500000000000000000)
export ICON_NATIVE_TOKEN_FEE_NUMERATOR=(100 100)
export ICON_NATIVE_TOKEN_DECIMALS=(18 18)
export ICON_WRAPPED_COIN_FIXED_FEE=(5000000000000000 1500000000000000000 1500000000000000000 1500000000000000000 62500000000000 750000000000000)
export ICON_WRAPPED_COIN_FEE_NUMERATOR=(100 100 100 100 100 100)
export ICON_WRAPPED_COIN_DECIMALS=(18 18 18 18 18 18)

export BSC_NID=${BSC_NID:-"97"}
export BSC_BMC_NET=${BSC_BMC_NET:-"0x61.bsc"}
export BSC_RPC_URI=${BSC_ENDPOINT:-"http://binancesmartchain:8545"}

export BSC_NATIVE_COIN_SYM=("BNB")
export BSC_NATIVE_COIN_NAME=("btp-$BSC_BMC_NET-BNB")
export BSC_NATIVE_TOKEN_SYM=("BUSD" "USDT" "USDC" "BTCB" "ETH")
export BSC_NATIVE_TOKEN_NAME=("btp-$BSC_BMC_NET-BUSD" "btp-$BSC_BMC_NET-USDT" "btp-$BSC_BMC_NET-USDC" "btp-$BSC_BMC_NET-BTCB" "btp-$BSC_BMC_NET-ETH")
export BSC_WRAPPED_COIN_SYM=("ICX" "sICX" "bnUSD")
export BSC_WRAPPED_COIN_NAME=("btp-$ICON_BMC_NET-ICX" "btp-$ICON_BMC_NET-sICX" "btp-$ICON_BMC_NET-bnUSD")

export BSC_NATIVE_COIN_FIXED_FEE=(5000000000000000)
export BSC_NATIVE_COIN_FEE_NUMERATOR=(100)
export BSC_NATIVE_COIN_DECIMALS=(18)
export BSC_NATIVE_TOKEN_FIXED_FEE=(1500000000000000000 1500000000000000000 1500000000000000000 62500000000000 750000000000000)
export BSC_NATIVE_TOKEN_FEE_NUMERATOR=(100 100 100 100 100)
export BSC_NATIVE_TOKEN_DECIMALS=(18 18 18 18 18)
export BSC_WRAPPED_COIN_FIXED_FEE=(4300000000000000000 3900000000000000000 1500000000000000000)
export BSC_WRAPPED_COIN_FEE_NUMERATOR=(100 100 100)
export BSC_WRAPPED_COIN_DECIMALS=(18 18 18)