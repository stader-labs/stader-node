package node

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	eCryto "github.com/ethereum/go-ethereum/crypto"
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
)

var (
	ExecutionClient = "Nethermind/v1.25.2+78c7bf5f/linux-x64/dotnet8.0.1"
	ConsensusClient = "Lighthouse/v4.6.0-rc.0-2e8e160/x86_64-linux"
)

func TestVerifySignature(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(eCryto.S256(), rand.Reader)
	if err != nil {
		t.Error(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Error(err)
	}

	pubkeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	req, err := makesNodeDiversityRequest(&stader_backend.NodeDiversity{
		ExecutionClient: ExecutionClient,
		ConsensusClient: ConsensusClient,
		NodeAddress:     crypto.PubkeyToAddress(*publicKeyECDSA).String(),
		NodePublicKey:   hex.EncodeToString(pubkeyBytes),
	}, privateKey)
	if err != nil {
		t.Error(err)
	}

	msgStr := req.Message
	verified := verifySignature(t, msgStr, req.Signature)

	if !verified {
		t.Error("verified should success")
	}
}

func TestVerifySignatureFailed(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(eCryto.S256(), rand.Reader)
	if err != nil {
		t.Error(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Error(err)
	}

	publickeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	privateKeyFake, err := ecdsa.GenerateKey(eCryto.S256(), rand.Reader)
	if err != nil {
		t.Error(err)
	}

	msg := stader_backend.NodeDiversity{
		ExecutionClient: ExecutionClient,
		ConsensusClient: ConsensusClient,
		NodeAddress:     crypto.PubkeyToAddress(*publicKeyECDSA).String(),
		NodePublicKey:   hex.EncodeToString(publickeyBytes),
	}
	req, err := makesNodeDiversityRequest(&msg, privateKeyFake)
	if err != nil {
		t.Error(err)
	}

	verified := verifySignature(t, req.Message, req.Signature)

	if verified {
		t.Error("verified should failed")
	}
}

func verifySignature(t *testing.T, msg *stader_backend.NodeDiversity, signEncoded string) bool {
	sign, err := hex.DecodeString(signEncoded)
	if err != nil {
		t.Error(err)
	}

	if msg.ConsensusClient != ConsensusClient {
		t.Error(err)
	}

	if msg.ExecutionClient != ExecutionClient {
		t.Error(err)
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}

	// 3. Calculate hash
	h := sha256.New()
	h.Write(msgBytes)
	msgHashed := h.Sum(nil)

	decodePubkey, err := hex.DecodeString(msg.NodePublicKey)
	if err != nil {
		t.Error(err)
	}

	// 4. Verify
	return eCryto.VerifySignature(decodePubkey, msgHashed, sign)
}
