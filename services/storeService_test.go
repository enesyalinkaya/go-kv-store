package services

import (
	"context"
	"testing"

	"github.com/enesyalinkaya/go-kv-store/models"
	"github.com/enesyalinkaya/go-kv-store/pkg/memoryDB"
)

var memoryDBClient = memoryDB.NewMemoryClient("tmp", "test.txt")
var storeModel = models.NewStoreModel(memoryDBClient)
var ctx = context.Background()

func loadSampleData() {
	storeModel.Set("key1", "value1")
	storeModel.Set("key3", "value2")
	storeModel.Set("key2", "value3")
}
func Test_storeService_Get(t *testing.T) {
	loadSampleData()
	type fields struct {
		StoreModel models.StoreModel
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "Get value",
			fields: fields{StoreModel: storeModel},
			args:   args{ctx: ctx, key: "key1"},
			want:   "value1",
		},
		{
			name:   "Get empty value",
			fields: fields{StoreModel: storeModel},
			args:   args{ctx: ctx, key: "key4"},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &storeService{
				StoreModel: tt.fields.StoreModel,
			}
			if got := s.Get(tt.args.ctx, tt.args.key); got != tt.want {
				t.Errorf("storeService.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storeService_Set(t *testing.T) {
	loadSampleData()
	type fields struct {
		StoreModel models.StoreModel
	}
	type args struct {
		ctx   context.Context
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "set value",
			fields: fields{
				StoreModel: storeModel,
			},
			args: args{
				ctx:   ctx,
				key:   "setTestKey1",
				value: "setTestValue1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &storeService{
				StoreModel: tt.fields.StoreModel,
			}
			value := s.Set(tt.args.ctx, tt.args.key, tt.args.value)
			if value != tt.args.value {
				t.Errorf("got %s, want %s", value, tt.args.value)
			}
		})
	}
}

func Test_storeService_Flush(t *testing.T) {
	type fields struct {
		StoreModel models.StoreModel
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "get empty value",
			fields: fields{
				StoreModel: storeModel,
			},
			args: args{
				ctx: ctx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &storeService{
				StoreModel: tt.fields.StoreModel,
			}
			s.Flush(tt.args.ctx)
			if got := s.Get(tt.args.ctx, "test1"); got != "" {
				t.Errorf("storeModel.Flush() = %v, want %v", got, "")
			}
		})
	}
}
