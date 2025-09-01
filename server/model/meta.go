package model

import (
	"context"
	"encoding/hex"
	"errors"
	"log"

	"github.com/yuudi/ero-runner/server/utils"
	"gorm.io/gorm"
)

type meta struct {
	Key   string `gorm:"primaryKey"`
	Value string
}

type metaDataType struct {
	SecretKey []byte
}

var metaData *metaDataType

func initMeta() {
	ctx := context.Background()
	metaData = &metaDataType{}
	secretKeyHex, err := gorm.G[meta](db).Where("key = ?", "secret_key").First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newKey, err := utils.GenerateCryptoRandomBytes(32)
			if err != nil {
				log.Fatal(err)
			}
			metaData.SecretKey = newKey
			newKeyHex := hex.EncodeToString(newKey)
			if err := gorm.G[meta](db).Create(ctx, &meta{Key: "secret_key", Value: newKeyHex}); err != nil {
				log.Fatal(err)
			}
			return
		}
		log.Fatal(err)
	}
	metaData.SecretKey, err = hex.DecodeString(secretKeyHex.Value)
	if err != nil {
		log.Fatal(err)
	}
}

func GetMeta() *metaDataType {
	once.Do(initMeta)
	return metaData
}
