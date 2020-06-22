package rtdata

import "testing"

func TestOperandStack_PushPopInt(t *testing.T) {
	type fields struct {
		size  int
		slots []Slot
	}
	type args struct {
		value int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "测试设置int",
			fields: fields{
				size:  0,
				slots: make([]Slot, 1),
			},
			args: args{
				value: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OperandStack{
				size:  tt.fields.size,
				slots: tt.fields.slots,
			}
			o.PushInt(tt.args.value)
			if tt.args.value != o.PopInt() {
				t.Errorf("value can't match excepted!!")
			}
		})
	}
}

func TestOperandStack_PushPopFloat(t *testing.T) {
	type fields struct {
		size  int
		slots []Slot
	}
	type args struct {
		value float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "测试设置float",
			fields: fields{
				size:  0,
				slots: make([]Slot, 1),
			},
			args: args{
				value: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OperandStack{
				size:  tt.fields.size,
				slots: tt.fields.slots,
			}
			o.PushFloat(tt.args.value)
			if tt.args.value != o.PopFloat() {
				t.Errorf("value can't match excepted!!")
			}
		})
	}
}

func TestOperandStack_PushPopLong(t *testing.T) {
	type fields struct {
		size  int
		slots []Slot
	}
	type args struct {
		value int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "测试设置long",
			fields: fields{
				size:  0,
				slots: make([]Slot, 2),
			},
			args: args{
				value: 2997924580,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OperandStack{
				size:  tt.fields.size,
				slots: tt.fields.slots,
			}
			o.PushLong(tt.args.value)
			if tt.args.value != o.PopLong() {
				t.Errorf("value can't match excepted!!")
			}
		})
	}
}

func TestOperandStack_PushPopDouble(t *testing.T) {
	type fields struct {
		size  int
		slots []Slot
	}
	type args struct {
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "测试设置double",
			fields: fields{
				size:  0,
				slots: make([]Slot, 2),
			},
			args: args{
				value: 2.71828182845,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OperandStack{
				size:  tt.fields.size,
				slots: tt.fields.slots,
			}
			o.PushDouble(tt.args.value)
			if tt.args.value != o.PopDouble() {
				t.Errorf("value can't match excepted!!")
			}
		})
	}
}

func TestOperandStack_PushPopRef(t *testing.T) {
	type fields struct {
		size  int
		slots []Slot
	}
	type args struct {
		value *Object
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "测试设置ref",
			fields: fields{
				size:  0,
				slots: make([]Slot, 1),
			},
			args: args{
				value: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OperandStack{
				size:  tt.fields.size,
				slots: tt.fields.slots,
			}
			o.PushRef(tt.args.value)
			if tt.args.value != o.PopRef() {
				t.Errorf("value can't match excepted!!")
			}
		})
	}
}