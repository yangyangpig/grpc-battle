package spliter

import (
	"fmt"
	"grpc-battle/pkg/dbdriver/mysql/algo"
	"grpc-battle/pkg/dbdriver/mysql/algo/simplehash"
	"testing"
)

func TestNewSpliter(t *testing.T) {
	spliter := NewSpliter(WithSetAlgorithm(func() algo.Algo {
		return simplehash.NewSimpleHash()
	}))

	databaseName := spliter.DatabaseName("ugc", 400)
	tableName := spliter.TableName("friend_ship", 200)

	fmt.Println(databaseName, "\n", tableName)
}
