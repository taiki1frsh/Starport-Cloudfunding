// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgFund } from "./types/cloudfunding/tx";
import { MsgCreateProject } from "./types/cloudfunding/tx";
import { MsgStopProject } from "./types/cloudfunding/tx";


const types = [
  ["/taikifuru.cloudfunding.cloudfunding.MsgFund", MsgFund],
  ["/taikifuru.cloudfunding.cloudfunding.MsgCreateProject", MsgCreateProject],
  ["/taikifuru.cloudfunding.cloudfunding.MsgStopProject", MsgStopProject],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgFund: (data: MsgFund): EncodeObject => ({ typeUrl: "/taikifuru.cloudfunding.cloudfunding.MsgFund", value: MsgFund.fromPartial( data ) }),
    msgCreateProject: (data: MsgCreateProject): EncodeObject => ({ typeUrl: "/taikifuru.cloudfunding.cloudfunding.MsgCreateProject", value: MsgCreateProject.fromPartial( data ) }),
    msgStopProject: (data: MsgStopProject): EncodeObject => ({ typeUrl: "/taikifuru.cloudfunding.cloudfunding.MsgStopProject", value: MsgStopProject.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
