package helper

import (
	"time"

	badger "github.com/dgraph-io/badger/v4"

	"gin-starter/internal/config"
)

func SetCacheValue(key, val string, ttl time.Duration) error {
	return config.Config.CacheClient.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte(key), []byte(val)).WithTTL(ttl)
		return txn.SetEntry(entry)
	})
}

func GetCacheValue(key string) (string, error) {
	var result string

	err := config.Config.CacheClient.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		var valCopy []byte
		if err := item.Value(func(val []byte) error {
			valCopy = append([]byte{}, val...)

			result = string(valCopy)
			return nil
		}); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	return result, nil
}
