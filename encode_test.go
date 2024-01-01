package fastconv

//
//func Test_encodeValue(t *testing.T) {
//	type args struct {
//		value   interface{}
//		buffer  *Buffer
//		encoder map[reflect.Type]string
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		{
//			name: "bool",
//			args: args{
//				value: true,
//				buffer: &Buffer{buf: []Value{
//					{
//						Type:  Bool,
//						Value: [8]byte{1},
//					},
//				}},
//				encoder: map[reflect.Type]string{},
//			},
//		},
//		{
//			name: "*bool",
//			args: args{
//				value: internal.Ptr(false),
//				buffer: &Buffer{buf: []Value{
//					{
//						Type:  Bool,
//						Value: [8]byte{0},
//					},
//				}},
//				encoder: map[reflect.Type]string{
//					internal.TypeFor[*bool](): "ptrEncoder.encode-fm",
//				},
//			},
//		},
//		{
//			name: "int",
//			args: args{
//				value: int(3),
//				buffer: &Buffer{buf: []Value{
//					{
//						Type:  Int,
//						Value: [8]byte{3},
//					},
//				}},
//				encoder: map[reflect.Type]string{},
//			},
//		},
//		{
//			name: "*int64",
//			args: args{
//				value: internal.Ptr(int64(3)),
//				buffer: &Buffer{buf: []Value{
//					{
//						Type:  Int,
//						Value: [8]byte{3},
//					},
//				}},
//				encoder: map[reflect.Type]string{
//					internal.TypeFor[*int64](): "ptrEncoder.encode-fm",
//				},
//			},
//		},
//		{
//			name: "uint",
//			args: args{
//				value: uint(3),
//				buffer: &Buffer{buf: []Value{
//					{
//						Type:  Uint,
//						Value: [8]byte{3},
//					},
//				}},
//				encoder: map[reflect.Type]string{},
//			},
//		},
//		{
//			name: "*uint64",
//			args: args{
//				value: internal.Ptr(uint64(3)),
//				buffer: &Buffer{buf: []Value{
//					{
//						Type:  Uint,
//						Value: [8]byte{3},
//					},
//				}},
//				encoder: map[reflect.Type]string{
//					internal.TypeFor[*uint64](): "ptrEncoder.encode-fm",
//				},
//			},
//		},
//		{
//			name: "float32",
//			args: args{
//				value: float32(3),
//				buffer: &Buffer{buf: []Value{
//					{
//						Type:  Float,
//						Value: [8]byte{0, 0, 0, 0, 0, 0, 8, 64},
//					},
//				}},
//				encoder: map[reflect.Type]string{},
//			},
//		},
//		{
//			name: "*float64",
//			args: args{
//				value: internal.Ptr(float64(3)),
//				buffer: &Buffer{buf: []Value{
//					{
//						Type:  Float,
//						Value: [8]byte{0, 0, 0, 0, 0, 0, 8, 64},
//					},
//				}},
//				encoder: map[reflect.Type]string{
//					internal.TypeFor[*float64](): "ptrEncoder.encode-fm",
//				},
//			},
//		},
//		{
//			name: "string",
//			args: args{
//				value: "3",
//				buffer: &Buffer{buf: []Value{
//					{
//						Type:  String,
//						Value: [8]byte{0, 0, 0, 0, 0, 0, 8, 64},
//					},
//				}},
//				encoder: map[reflect.Type]string{},
//			},
//		},
//		{
//			name: "*string",
//			args: args{
//				value: internal.Ptr("3"),
//				buffer: &Buffer{buf: []Value{
//					{
//						Type:  String,
//						Value: [8]byte{0, 0, 0, 0, 0, 0, 8, 64},
//					},
//				}},
//				encoder: map[reflect.Type]string{
//					internal.TypeFor[*string](): "ptrEncoder.encode-fm",
//				},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			l := &Buffer{}
//			l.Append(1)
//			v := reflect.ValueOf(tt.args.value)
//			encodeValue(l, 0, &l.buf[0], v)
//			assert.Equal(t, l, tt.args.buffer)
//			encoder := make(map[reflect.Type]string)
//			encoderCache.Range(func(key, value any) bool {
//				fn := runtime.FuncForPC(reflect.ValueOf(value.(encoderFunc)).Pointer())
//				encoder[key.(reflect.Type)] = strings.TrimPrefix(fn.Name(), "github.com/lvan100/fastconv.")
//				encoderCache.Delete(key)
//				return true
//			})
//			assert.Equal(t, encoder, tt.args.encoder)
//		})
//	}
//}
