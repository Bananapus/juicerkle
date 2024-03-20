# Juicerkle

Juicerkle is a merkle proof generation server intended for use with [`nana-suckers`](https://github.com/bananapus/nana-suckers).

## Server Setup

To run Juicerkle, you'll need to have a working installation of [Go v1.22](https://go.dev/dl/) or greater.

To configure your server, copy the example config file and add any chains you'd like to support:

```bash
# copy the file
cp example.config.json config.json
# edit the file
vim config.json
```

Then you can build the binary with `go build` and run `./juicerkle`. It will start a server on `localhost:8080` by default. To run Juicerkle on a different port, set the `PORT` environment variable.

## Querying

To get a proof from Juicerkle, send a `POST` request to `/proof` with a specification in the following format:

The JSON schema for the proof request is as follows:

| Field     | JS Type  | Description                                                                   |
| --------- | -------- | ----------------------------------------------------------------------------- |
| `chainId` | `int`    | The ID of the chain that the sucker contract is on.                           |
| `sucker`  | `string` | The sucker contract's address.                                                |
| `token`   | `string` | The address of the `terminalToken` being claimed from the contract.           |
| `index`   | `int`    | The index of the leaf to prove in the sucker contract's tree for the `token`. |

For example, the following json:

```json
{
  "chainId": 1,
  "sucker": "0xf19736127f87569423eC4F10FD8CC984bbcE0A17",
  "token": "0x000000000000000000000000000000000000EEEe",
  "index": 0
}
```

Would return the proof for the first leaf in the native token inbox tree (represented by `0x000000000000000000000000000000000000EEEe`) on the Ethereum mainnet sucker deployed to `0xf19736127f87569423eC4F10FD8CC984bbcE0A17`.