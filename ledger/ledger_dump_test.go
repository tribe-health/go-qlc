package ledger

import (
	"fmt"
	"testing"
	"time"

	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/mock"
)

func TestLedger_dump(t *testing.T) {
	teardownTestCase, l := setupTestCase(t)
	defer teardownTestCase(t)
	acc := new(types.AccountMeta)
	acc.Address = mock.Address()
	token := mock.TokenMeta(acc.Address)
	acc.Tokens = append(acc.Tokens, token)

	blk1 := mock.StateBlockWithoutWork()
	blk1.Address = acc.Address
	blk1.Token = acc.Tokens[0].Type

	blk2 := mock.StateBlockWithoutWork()
	blk2.Address = acc.Address
	blk2.Token = acc.Tokens[0].Type
	blk2.Previous = blk1.GetHash()

	blk3 := mock.StateBlockWithoutWork()
	blk3.Address = acc.Address
	blk3.Token = acc.Tokens[0].Type
	blk3.Previous = blk2.GetHash()

	blk4 := mock.StateBlockWithoutWork()
	blk4.Address = acc.Address
	blk4.Token = acc.Tokens[0].Type
	blk4.Previous = blk3.GetHash() // off chain

	blk5 := mock.StateBlockWithoutWork()
	blk5.Address = acc.Address
	blk5.Token = acc.Tokens[0].Type
	blk5.Previous = blk4.GetHash() // off chain

	acc.Tokens[0].Header = blk3.GetHash()
	fmt.Println(acc)

	if err := l.AddStateBlock(blk1); err != nil {
		t.Fatal(err)
	}
	if err := l.AddStateBlock(blk2); err != nil {
		t.Fatal(err)
	}
	if err := l.AddStateBlock(blk3); err != nil {
		t.Fatal(err)
	}
	if err := l.AddStateBlock(blk4); err != nil {
		t.Fatal(err)
	}
	if err := l.AddAccountMeta(acc, l.cache.GetCache()); err != nil {
		t.Fatal(err)
	}

	blkCache1 := mock.StateBlockWithoutWork()
	accCache1 := mock.AccountMeta(blkCache1.Address)
	blkCache1.Token = accCache1.Tokens[0].Type
	blkCache2 := mock.StateBlockWithoutWork()
	accCache2 := mock.AccountMeta(blkCache2.Address)
	blkCache2.Token = accCache2.Tokens[0].Type
	if err := l.AddBlockCache(blkCache1); err != nil {
		t.Fatal(err)
	}
	if err := l.AddBlockCache(blkCache2); err != nil {
		t.Fatal(err)
	}
	if err := l.AddAccountMetaCache(accCache1); err != nil {
		t.Fatal(err)
	}
	if err := l.AddUncheckedBlock(mock.Hash(), mock.StateBlockWithoutWork(), types.UncheckedKindLink, types.UnSynchronized); err != nil {
		fmt.Println(err)
		return
	}
	if err := l.Flush(); err != nil {
		t.Fatal(err)
	}
	if _, err := l.Dump(0); err != nil {
		t.Fatal(err)
	}

	start := time.Now()
	for {
		if r, err := l.Dump(1); err != nil {
			t.Fatal(err)
		} else {
			if r == "done" {
				return
			}
			time.Sleep(200 * time.Millisecond)
			if time.Since(start) > 10*time.Second {
				t.Fatal("timeout")
			}
		}

	}
}
