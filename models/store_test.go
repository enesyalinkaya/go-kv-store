package models

import (
	"testing"

	memory "github.com/enesyalinkaya/go-kv-store/pkg/memoryDB"
)

var db = memory.NewMemoryClient("tmp/test.txt")

func loadSampleData() {
	db.Set("key1", "value1")
	db.Set("key2", "value2")
	db.Set("key3", "value3")
}

func Test_storeModel_Set(t *testing.T) {
	loadSampleData()
	type fields struct {
		db *memory.MemoryClient
	}
	type args struct {
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
				db: db,
			},
			args: args{
				key:   "setkey1",
				value: "setvalue1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &storeModel{
				db: tt.fields.db,
			}
			m.Set(tt.args.key, tt.args.value)
			value := m.Get(tt.args.key)
			if value != tt.args.value {
				t.Errorf("got %s, want %s", value, tt.args.value)
			}
		})
	}
}

func Test_storeModel_Get(t *testing.T) {
	loadSampleData()
	type fields struct {
		db *memory.MemoryClient
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "get value",
			fields: fields{
				db: db,
			},
			args: args{
				key: "key11",
			},
			want: "",
		},
		{
			name: "get value",
			fields: fields{
				db: db,
			},
			args: args{
				key: "key1",
			},
			want: "value1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &storeModel{
				db: tt.fields.db,
			}
			if got := m.Get(tt.args.key); got != tt.want {
				t.Errorf("storeModel.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storeModel_Flush(t *testing.T) {
	loadSampleData()
	type fields struct {
		db *memory.MemoryClient
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "get empty value",
			fields: fields{
				db: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &storeModel{
				db: tt.fields.db,
			}
			m.Flush()
			if got := m.Get("test1"); got != "" {
				t.Errorf("storeModel.Flush() = %v, want %v", got, "")
			}
		})
	}
}
