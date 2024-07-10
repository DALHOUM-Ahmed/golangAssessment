# Go-IPFS File Registry Project

This project implements a file registry on the Binance Smart Chain (BSC) testnet. It uploads files to IPFS, stores the CID (Content Identifier) in a smart contract, and provides an API for uploading files and retrieving CIDs.

## Prerequisites

- Docker and Docker Compose

## Project Structure

```plaintext
go-assessment/
├── .env
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── handlers/
│   ├── ethereum/
│   │   └── ethereum.go
│   └── upload.go
├── abi/
│   └── FileRegistryABI.json
└── main.go
```

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/DALHOUM-Ahmed/golangAssessment.git
   cd golangAssessment
   ```

2. Create a `.env` file in the project directory with the following content:

   ```env
   IPFS_HOST=localhost:5001
   SERVER_PORT=8081
   RPC_URL=https://bsc-testnet-rpc.publicnode.com
   PRIVATE_KEY=the_owner_wallet_private_key
   CONTRACT_ADDRESS=0x52DF3ADCf7Ae617B22A36AfEd62cfebD13D94f44
   ```

3. Ensure the `abi/FileRegistryABI.json` file contains your smart contract's ABI.

## Running the Project

To run the IPFS node and the Go project together, use Docker Compose:

```bash
docker-compose up --build
```

To stop the services, either use `Ctrl+C` in the terminal or run:

```bash
docker-compose down
```

## Uploading a File

To upload a file, use the following `curl` command from the project directory:

```bash
curl -X POST -F "file=@weDidIt.png" -F "filePath=weDidIt.png" http://localhost:8081/v1/files
```

Example response:

```
Successfully uploaded file to IPFS and saved CID to blockchain: QmT8drg8K3D9Hz8zVmeNjfXXjG3mQnRG4wP8jkVgbVZDmt
```

## Visualizing the Uploaded Image

To visualize the uploaded image, use the following link in your browser:

```
http://localhost:8080/ipfs/QmT8drg8K3D9Hz8zVmeNjfXXjG3mQnRG4wP8jkVgbVZDmt
```

## Retrieving the CID from the Smart Contract

To request the CID from the register smart contract, use:

```bash
curl -X GET "http://localhost:8081/v1/files/get?filePath=weDidIt.png"
```

Example response:

```
CID for file weDidIt.png: QmT8drg8K3D9Hz8zVmeNjfXXjG3mQnRG4wP8jkVgbVZDmt
```

## Smart Contract Information

The register smart contract address with the code verified can be found here:

[https://testnet.bscscan.com/address/0x52DF3ADCf7Ae617B22A36AfEd62cfebD13D94f44](https://testnet.bscscan.com/address/0x52DF3ADCf7Ae617B22A36AfEd62cfebD13D94f44)

Only the owner can execute the save function. The owner is the address associated with the private key specified in the `.env` file.

## Notes

- Ensure your Go environment is set up correctly and dependencies are installed.
- Ensure Docker and Docker Compose are installed and running on your machine.

This project demonstrates integrating IPFS with a blockchain-based file registry using Go, Docker, and smart contracts on the BSC testnet.
