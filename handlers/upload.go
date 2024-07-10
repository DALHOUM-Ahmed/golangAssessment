package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"go-assessment/handlers/ethereum"

	shell "github.com/ipfs/go-ipfs-api"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	sh := shell.NewShell("ipfs:5001")

	tempFile, err := os.CreateTemp("", "ipfs-*")
	if err != nil {
		http.Error(w, "Failed to create temp file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Failed to copy file content", http.StatusInternalServerError)
		return
	}

	tempFile.Seek(0, 0)

	cid, err := sh.Add(tempFile)
	if err != nil {
		http.Error(w, "Failed to add file to IPFS", http.StatusInternalServerError)
		return
	}

	filePath := r.FormValue("filePath")
	err = ethereum.TransactWithContract("save", filePath, cid)
	if err != nil {
		http.Error(w, "Failed to save CID to blockchain", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully uploaded file to IPFS and saved CID to blockchain: %s", cid)
}

func GetCIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	filePath := r.URL.Query().Get("filePath")
	if filePath == "" {
		http.Error(w, "filePath is required", http.StatusBadRequest)
		return
	}

	cid, err := ethereum.CallContractFunction("get", filePath)
	if err != nil {
		http.Error(w, "Failed to get CID from blockchain", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "CID for file %s: %s", filePath, cid)
}
