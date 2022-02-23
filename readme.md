# cloudfunding
**cloudfunding** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://starport.com).

This chain is built just to clouodfunding module that works like  very simple cloudfunding-like system.

Messageは、
１．CreateProject
２．Fund
３．StopProject（本当は自動でやりたい）

```
starport s message create-project target deadline description
starport s list project target collected deadline description state creator --no-message
starport s message fund id:uint amt
starport s message stop-project id:uint
```

データ構造

Listで、IndexはUintのIdをKeyとする。
list project target deadline creator

message project {
  uint64 id = 1;
  string target = 2;
  string collected = 3
  string deadline = 4;　（Blockheightで計算）
  string description = 5;
  string state = 6;
  string creator = 7;
}

EscrowAccountはModuleAccountにする。
つまり、最初はFundしたトークンはModuleAccountに送る。
その前に、CollectedとTargetを比べ、もしそのときのFundでCollectedがTargetを超えるようであればStateをSuccessにし、閉じる（？）。
（最終的にはDeadlineまで続けさせて、自動で閉じるようにしたい）