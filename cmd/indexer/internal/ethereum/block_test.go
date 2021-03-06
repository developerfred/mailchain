package ethereum_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/mock/gomock"
	"github.com/mailchain/mailchain/cmd/indexer/internal/actions"
	"github.com/mailchain/mailchain/cmd/indexer/internal/actions/actionstest"
	"github.com/mailchain/mailchain/cmd/indexer/internal/ethereum"
)

func TestBlock_Run(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	type fields struct {
		txProcessor actions.Transaction
	}
	type args struct {
		ctx      context.Context
		protocol string
		network  string
		blk      interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"success",
			fields{
				func() actions.Transaction {
					m := actionstest.NewMockTransaction(mockCtrl)
					m.EXPECT().Run(context.Background(), "ethereum", "mainnet", gomock.Any(), gomock.Any()).Return(nil)
					m.EXPECT().Run(context.Background(), "ethereum", "mainnet", gomock.Any(), gomock.Any()).Return(nil)
					return m
				}(),
			},
			args{
				context.Background(),
				"ethereum",
				"mainnet",
				types.NewBlock(&types.Header{},
					[]*types.Transaction{
						getTx(t, "0xcd3ccc846c566fbf76f38b1184ba02261a821c6942d18146b89b9d285aa29b9c"),
						getTx(t, "0xd2c574543459bf6704174fa869df4974220b71f67395eff6e20f1f5ec9f72d50"),
					}, nil, nil),
			},
			false,
		},
		{
			"err-run",
			fields{
				func() actions.Transaction {
					m := actionstest.NewMockTransaction(mockCtrl)
					m.EXPECT().Run(context.Background(), "ethereum", "mainnet", gomock.Any(), gomock.Any()).Return(nil)
					m.EXPECT().Run(context.Background(), "ethereum", "mainnet", gomock.Any(), gomock.Any()).Return(errors.New("error"))
					return m
				}(),
			},
			args{
				context.Background(),
				"ethereum",
				"mainnet",
				types.NewBlock(&types.Header{},
					[]*types.Transaction{
						getTx(t, "0xcd3ccc846c566fbf76f38b1184ba02261a821c6942d18146b89b9d285aa29b9c"),
						getTx(t, "0xd2c574543459bf6704174fa869df4974220b71f67395eff6e20f1f5ec9f72d50"),
					}, nil, nil),
			},
			true,
		},
		{
			"err-arg-blk",
			fields{
				func() actions.Transaction {
					m := actionstest.NewMockTransaction(mockCtrl)
					return m
				}(),
			},
			args{
				context.Background(),
				"ethereum",
				"mainnet",
				0,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := ethereum.NewBlockProcessor(tt.fields.txProcessor)
			if err := b.Run(tt.args.ctx, tt.args.protocol, tt.args.network, tt.args.blk); (err != nil) != tt.wantErr {
				t.Errorf("Block.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
