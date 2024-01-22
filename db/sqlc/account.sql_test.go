package db

import (
	"context"
	"testing"

	"github.com/mayuka-c/bank-app/utils"
	"github.com/stretchr/testify/require"
)

func RandomOwner() string {
	return utils.RandomString(6)
}

func RandomMoney() int64 {
	return utils.RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "INR"}
	return currencies[utils.RandInt(len(currencies))]
}

func createRandomAccount(t *testing.T) Account {
	var got Account
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		arg CreateAccountParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Account
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				db: testQueries.db,
			},
			args: args{
				ctx: context.Background(),
				arg: CreateAccountParams{
					Owner:    RandomOwner(),
					Balance:  RandomMoney(),
					Currency: RandomCurrency(),
				},
			},
			want:    Account{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			q := &Queries{
				db: tt.fields.db,
			}
			got, err = q.CreateAccount(tt.args.ctx, tt.args.arg)
			require.NoError(t, err)
			require.NotEmpty(t, got)

			require.Equal(t, tt.args.arg.Owner, got.Owner)
			require.Equal(t, tt.args.arg.Balance, got.Balance)
			require.Equal(t, tt.args.arg.Currency, got.Currency)

			require.NotZero(t, got.ID)
			require.NotZero(t, got.CreatedAt)
		})
	}

	return got
}

func TestQueries_CreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestQueries_GetAccount(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Success",
			fields: fields{
				db: testQueries.db,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.GetAccount(tt.args.ctx, tt.args.id)
			require.NoError(t, err)

			require.NotEmpty(t, got.Balance)
			require.NotEmpty(t, got.Owner)
			require.NotEmpty(t, got.Currency)
			require.NotEmpty(t, got.CreatedAt)
		})
	}
}

func TestQueries_UpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		arg UpdateAccountParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Account
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				db: testQueries.db,
			},
			args: args{
				ctx: context.Background(),
				arg: UpdateAccountParams{
					ID:      account1.ID,
					Balance: account1.Balance + 400,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queries{
				db: tt.fields.db,
			}
			got, err := q.UpdateAccount(tt.args.ctx, tt.args.arg)

			require.NoError(t, err)

			require.Equal(t, account1.ID, got.ID)
			require.Equal(t, account1.Owner, got.Owner)
			require.Equal(t, account1.Balance+400, got.Balance)
			require.Equal(t, account1.Currency, got.Currency)
			require.Equal(t, account1.CreatedAt, got.CreatedAt)
		})
	}
}
