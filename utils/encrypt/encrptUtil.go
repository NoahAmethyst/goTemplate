package encrypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/google/uuid"
	"strings"
)

func GenerateUuid(format bool) string {
	uuidValue := uuid.NewString()
	if format {
		uuidValue = strings.Replace(uuidValue, "-", "", -1)
	}
	return uuidValue
}

func Md5WithSecretKey(content string, secretKey string) string {
	v := content + secretKey
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func Sha256HMAC(content []byte, secretKey string) string {

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write(content)
	expectedMAC := mac.Sum(nil)
	sha := hex.EncodeToString(expectedMAC)
	return sha

}

func GenerateRandomKeyPair() {
	apiSecret := make([]byte, 32)
	if _, err := rand.Read(apiSecret); err != nil {
		panic(err)
	}
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), apiSecret)
	apiKey := fmt.Sprintf("%x", privKey.PubKey().SerializeCompressed())
	apiSecretStr := fmt.Sprintf("%x", apiSecret)

	fmt.Printf("API_Key: %s\nAPI_SECRET: %s\n", apiKey, apiSecretStr)
}

func Hash256(s string) string {
	hashResult := sha256.Sum256([]byte(s))
	hashString := string(hashResult[:])
	return hashString
}
func Hash256x2(s string) string {
	return Hash256(Hash256(s))
}
