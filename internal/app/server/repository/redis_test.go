package repository

import (
	"sync"
	"testing"
	"time"
)

func TestInMemoryRedis_Del(t *testing.T) {
	type fields struct {
		store sync.Map
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Delete existed key",
			fields: fields{store: func() sync.Map {
				var store sync.Map
				store.Store("mykey", storeValue{
					value:  "myval",
					expiry: time.Time{},
				})
				return store
			}()},
			args: args{key: "mykey"},
			want: true,
		},
		{
			name: "Delete nonexistent key",
			args: args{key: "mykey"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemoryRedis{
				store: tt.fields.store,
			}
			if got := r.Del(tt.args.key); got != tt.want {
				t.Errorf("Del() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryRedis_Expire(t *testing.T) {
	type fields struct {
		store sync.Map
	}
	type args struct {
		key      string
		duration int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "expire key when it exists",
			fields: fields{store: func() sync.Map {
				var store sync.Map
				store.Store("mykey", storeValue{
					value:  "myval",
					expiry: time.Time{},
				})
				return store
			}()},
			args: args{key: "mykey", duration: 10},
			want: true,
		},
		{
			name:   "expire key when it doesn't exist",
			fields: fields{},
			args:   args{key: "mykey", duration: 10},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemoryRedis{
				store: tt.fields.store,
			}
			if got := r.Expire(tt.args.key, tt.args.duration); got != tt.want {
				t.Errorf("Expire() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryRedis_Get(t *testing.T) {
	type fields struct {
		store sync.Map
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   bool
		wantErr bool
	}{
		{
			name: "get by key when it exists",
			fields: fields{store: func() sync.Map {
				var store sync.Map
				store.Store("mykey", storeValue{
					value:  "myval",
					expiry: time.Time{},
				})
				return store
			}()},
			args:    args{key: "mykey"},
			want:    "myval",
			want1:   true,
			wantErr: false,
		},
		{
			name:    "get by key when it doesn't exist",
			fields:  fields{},
			args:    args{key: "mykey"},
			want:    "",
			want1:   false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemoryRedis{
				store: tt.fields.store,
			}
			got, got1, err := r.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInMemoryRedis_HGet(t *testing.T) {
	type fields struct {
		store sync.Map
	}
	type args struct {
		key   string
		field string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   bool
		wantErr bool
	}{
		{
			name:    "HGet when key doesn't exist",
			fields:  fields{},
			args:    args{key: "mykey", field: "some_field"},
			want:    "",
			want1:   false,
			wantErr: false,
		},
		{
			name: "HGet when key and field exists",
			fields: fields{store: func() sync.Map {
				var store sync.Map
				newMap := make(map[string]string)
				store.Store("mykey", storeValue{
					value: newMap,
				})

				newMap["some_field"] = "myval"
				return store
			}()},
			args:    args{key: "mykey", field: "some_field"},
			want:    "myval",
			want1:   true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemoryRedis{
				store: tt.fields.store,
			}
			got, got1, err := r.HGet(tt.args.key, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("HGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HGet() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("HGet() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInMemoryRedis_HSet(t *testing.T) {
	type fields struct {
		store sync.Map
	}
	type args struct {
		key   string
		field string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "HSet when a key holding the wrong kind of value",
			fields: fields{store: func() sync.Map {
				var store sync.Map
				store.Store("mykey", storeValue{
					value: "some_string",
				})

				return store
			}()},
			args:    args{key: "mykey", field: "some_field"},
			want:    false,
			wantErr: true,
		},
		{
			name: "HSet success",
			fields: fields{store: func() sync.Map {
				var store sync.Map
				newMap := make(map[string]string)
				store.Store("mykey", storeValue{
					value: newMap,
				})

				return store
			}()},
			args:    args{key: "mykey", field: "some_field", value: "value"},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemoryRedis{
				store: tt.fields.store,
			}
			got, err := r.HSet(tt.args.key, tt.args.field, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("HSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HSet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryRedis_Keys(t *testing.T) {
	type fields struct {
		store sync.Map
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Check keys command",
			fields: fields{func() sync.Map {
				var store sync.Map
				store.Store("firstname", storeValue{
					value: "some_string",
				})

				store.Store("lastname", storeValue{
					value: "some_string",
				})

				store.Store("age", storeValue{
					value: "35",
				})
				return store
			}()},
			args:    args{".*name.*"},
			want:    []string{"firstname", "lastname"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemoryRedis{
				store: tt.fields.store,
			}
			got, err := r.Keys(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("Keys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryRedis_LGet(t *testing.T) {
	type fields struct {
		store sync.Map
	}
	type args struct {
		key   string
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "LGet when a key doesn't exist",
			fields:  fields{},
			args:    args{key: "mykey", index: 0},
			want:    "",
			wantErr: true,
		},
		{
			name: "LGet when a key holding the wrong kind of value",
			fields: fields{store: func() sync.Map {
				var store sync.Map
				store.Store("mykey", storeValue{
					value: "some_string",
				})

				return store
			}()},
			args:    args{key: "mykey", index: 0},
			want:    "",
			wantErr: true,
		},
		{
			name: "LGet when index out of range",
			fields: fields{store: func() sync.Map {
				var store sync.Map
				newArr := make([]string, 0, 5)
				store.Store("mykey", storeValue{
					value: newArr,
				})

				return store
			}()},
			args:    args{key: "mykey", index: 100},
			want:    "",
			wantErr: true,
		},
		{
			name: "LGet success",
			fields: fields{store: func() sync.Map {
				var store sync.Map
				newArr := make([]string, 0, 5)
				newArr = append(newArr, "value")
				store.Store("mykey", storeValue{
					value: newArr,
				})

				return store
			}()},
			args:    args{key: "mykey", index: 0},
			want:    "value",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemoryRedis{
				store: tt.fields.store,
			}
			got, err := r.LGet(tt.args.key, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("LGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LGet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryRedis_LPush(t *testing.T) {
	type fields struct {
		store sync.Map
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "LPush when a key holding the wrong kind of value",
			fields: fields{store: func() sync.Map {
				var store sync.Map
				store.Store("mykey", storeValue{
					value: "some_string",
				})

				return store
			}()},
			args:    args{key: "mykey", value: "value"},
			want:    -1,
			wantErr: true,
		},
		{
			name: "LPush success",
			fields: fields{store: func() sync.Map {
				var store sync.Map
				newArr := make([]string, 0, 5)
				store.Store("mykey", storeValue{
					value: newArr,
				})

				return store
			}()},
			args:    args{key: "mykey", value: "value"},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemoryRedis{
				store: tt.fields.store,
			}
			got, err := r.LPush(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("LPush() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LPush() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryRedis_LSet(t *testing.T) {
	type fields struct {
		store sync.Map
	}
	type args struct {
		key   string
		index int
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "LSet when a key doesn't exist",
			fields:  fields{},
			args:    args{key: "mykey", index: 0},
			wantErr: true,
		},
		{
			name: "LSet when a key holding the wrong kind of value",
			fields: fields{store: func() sync.Map {
				var store sync.Map
				store.Store("mykey", storeValue{
					value: "some_string",
				})

				return store
			}()},
			args:    args{key: "mykey", index: 0},
			wantErr: true,
		},
		{
			name: "LSet when index out of range",
			fields: fields{store: func() sync.Map {
				var store sync.Map
				newArr := make([]string, 0, 5)
				store.Store("mykey", storeValue{
					value: newArr,
				})

				return store
			}()},
			args:    args{key: "mykey", index: 100},
			wantErr: true,
		},
		{
			name: "LSet success",
			fields: fields{store: func() sync.Map {
				var store sync.Map
				newArr := make([]string, 0, 5)
				newArr = append(newArr, "value")
				store.Store("mykey", storeValue{
					value: newArr,
				})

				return store
			}()},
			args:    args{key: "mykey", index: 0, value: "new_value"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemoryRedis{
				store: tt.fields.store,
			}
			if err := r.LSet(tt.args.key, tt.args.index, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("LSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
