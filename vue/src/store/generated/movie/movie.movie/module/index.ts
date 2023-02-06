// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateMovie } from "./types/movie/tx";
import { MsgUpdateMovie } from "./types/movie/tx";
import { MsgDeleteReview } from "./types/movie/tx";
import { MsgUpdateReview } from "./types/movie/tx";
import { MsgDeleteMovie } from "./types/movie/tx";
import { MsgCreateReview } from "./types/movie/tx";


const types = [
  ["/movie.movie.MsgCreateMovie", MsgCreateMovie],
  ["/movie.movie.MsgUpdateMovie", MsgUpdateMovie],
  ["/movie.movie.MsgDeleteReview", MsgDeleteReview],
  ["/movie.movie.MsgUpdateReview", MsgUpdateReview],
  ["/movie.movie.MsgDeleteMovie", MsgDeleteMovie],
  ["/movie.movie.MsgCreateReview", MsgCreateReview],
  
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
    msgCreateMovie: (data: MsgCreateMovie): EncodeObject => ({ typeUrl: "/movie.movie.MsgCreateMovie", value: MsgCreateMovie.fromPartial( data ) }),
    msgUpdateMovie: (data: MsgUpdateMovie): EncodeObject => ({ typeUrl: "/movie.movie.MsgUpdateMovie", value: MsgUpdateMovie.fromPartial( data ) }),
    msgDeleteReview: (data: MsgDeleteReview): EncodeObject => ({ typeUrl: "/movie.movie.MsgDeleteReview", value: MsgDeleteReview.fromPartial( data ) }),
    msgUpdateReview: (data: MsgUpdateReview): EncodeObject => ({ typeUrl: "/movie.movie.MsgUpdateReview", value: MsgUpdateReview.fromPartial( data ) }),
    msgDeleteMovie: (data: MsgDeleteMovie): EncodeObject => ({ typeUrl: "/movie.movie.MsgDeleteMovie", value: MsgDeleteMovie.fromPartial( data ) }),
    msgCreateReview: (data: MsgCreateReview): EncodeObject => ({ typeUrl: "/movie.movie.MsgCreateReview", value: MsgCreateReview.fromPartial( data ) }),
    
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
