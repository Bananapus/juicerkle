# Juicerkle

Juicerkle generates claims and merkle proofs for [`nana-suckers`](https://github.com/bananapus/nana-suckers). For a full walkthrough, see [_Bridging in Juicebox v4_](https://filip.world/post/suckers/).

## Server Setup

To run Juicerkle, you'll need to have a working installation of [Go v1.22](https://go.dev/dl/) or greater.

To configure your server, copy the example config file and add any chains you'd like to support:

```bash
# copy the file
cp example.config.json config.json
# edit the file
vim config.json
```

Then you can build the binary with `go build` and run `./juicerkle`. It will start a server on `http://localhost:8080` by default. To run Juicerkle on a different port, set the `PORT` environment variable.

## Querying

To get a proof from Juicerkle, send a `POST` request to `/claims` with a specification in the following format:

| Field         | JS Type  | Description                                                                |
| ------------- | -------- | -------------------------------------------------------------------------- |
| `chainId`     | `int`    | The network ID for the sucker contract being claimed from.                 |
| `sucker`      | `string` | The address of the sucker being claimed from.                              |
| `token`       | `string` | The address of the `terminalToken` whose inbox tree is being claimed from. |
| `beneficiary` | `string` | The address of the beneficiary we're getting the available claims for.     |

For example, the following request:

```js
{
    "chainId": 10,
    "sucker": "0x5678…",
    "token": "0x000000000000000000000000000000000000EEEe",
    "beneficiary": "0x1234…" // jimmy.eth
}
```

This request is checking for available claims:

- On Optimism (`chainId` 10)
- For the sucker contract at `0x5678…`
- In its ETH inbox tree, represented by [`JBConstants.NATIVE_TOKEN`](https://github.com/Bananapus/nana-core/blob/main/src/libraries/JBConstants.sol)
- For the beneficiary `0x1234…` (jimmy.eth)

`juicerkle` would return an array of `BPClaims` which can be passed to `BPSucker.claim(…)`.

## Example Response

```js
[
  {
    Token: "0x000000000000000000000000000000000000eeee",
    Leaf: {
      Index: 0,
      Beneficiary: "0x1234…", // jimmy.eth
      ProjectTokenAmount: 1000000000000000000, // 1e18
      TerminalTokenAmount: 1000000000000000000, // 1e18
    },
    Proof: [
      [
        229, 206, 51, 48, 16, 242, 169, 29, 47, 33, 39, 105, 34, 55, 172, 232,
        217, 243, 168, 149, 38, 202, 133, 68, 191, 119, 165, 97, 59, 232, 212,
        14,
      ],
      [
        33, 40, 178, 36, 156, 7, 175, 252, 47, 196, 238, 239, 170, 52, 239, 153,
        66, 111, 173, 24, 113, 164, 25, 185, 54, 47, 170, 32, 232, 56, 97, 254,
      ],
      // More 32-byte chunks…
    ],
  },
  // More claims…
];
```
