package rtdata

import "testing"

func TestLocalVars_SetGetInt(t *testing.T) {
	type args struct {
		index uint
		value int32
	}
	tests := []struct {
		name string
		vars LocalVars
		args args
	}{
		{
			name: "测试设置int",
			vars: NewLocalVars(1),
			args: args{
				index: 0,
				value: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vars.SetInt(tt.args.index, tt.args.value)
			if tt.args.value != tt.vars.GetInt(tt.args.index) {
				t.Errorf("value can't match excepted!!")
			}
		})
	}
}

func TestLocalVars_SetGetFloat(t *testing.T) {
	type args struct {
		index uint
		value float32
	}
	tests := []struct {
		name string
		vars LocalVars
		args args
	}{
		{
			name: "测试设置float",
			vars: NewLocalVars(1),
			args: args{
				index: 0,
				value: 100.01,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vars.SetFloat(tt.args.index, tt.args.value)
			if tt.args.value != tt.vars.GetFloat(tt.args.index) {
				t.Errorf("value can't match excepted!!")
			}
		})
	}
}

func TestLocalVars_SetGetLong(t *testing.T) {
	type args struct {
		index uint
		value int64
	}
	tests := []struct {
		name string
		vars LocalVars
		args args
	}{
		{
			name: "测试设置long",
			vars: NewLocalVars(2),
			args: args{
				index: 0,
				value: 2997924580,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vars.SetLong(tt.args.index, tt.args.value)
			if tt.args.value != tt.vars.GetLong(tt.args.index) {
				t.Errorf("value can't match excepted!!")
			}
		})
	}
}

func TestLocalVars_SetGetDouble(t *testing.T) {
	type args struct {
		index uint
		value float64
	}
	tests := []struct {
		name string
		vars LocalVars
		args args
	}{
		{
			name: "测试设置doublue",
			vars: NewLocalVars(2),
			args: args{
				index: 0,
				value: 2.71828182845,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vars.SetDouble(tt.args.index, tt.args.value)
			if tt.args.value != tt.vars.GetDouble(tt.args.index) {
				t.Errorf("value can't match excepted!!")
			}
		})
	}
}

func TestLocalVars_SetRef(t *testing.T) {
	type args struct {
		index uint
		ref   *Object
	}
	tests := []struct {
		name string
		vars LocalVars
		args args
	}{
		{
			name: "测试设置ref",
			vars: NewLocalVars(1),
			args: args{
				index: 0,
				ref: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vars.SetRef(tt.args.index, tt.args.ref)
			if tt.args.ref != tt.vars.GetRef(tt.args.index) {
				t.Errorf("value can't match excepted!!")
			}
		})
	}
}
