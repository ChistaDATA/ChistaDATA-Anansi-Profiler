package postgres

//func TestParseMessageOnlyLogV1(t *testing.T) {
//	type args struct {
//		extractedLog         stucts.ExtractedLog
//		dBPerfInfoRepository *stucts.DBPerfInfoRepository
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//		test    func(args) bool
//	}{
//		{
//			name: "first",
//			args: args{
//				extractedLog:         stucts.ExtractedLog{Message: "first"},
//				dBPerfInfoRepository: stucts.InitDBPerfInfoRepository(),
//			},
//			wantErr: false,
//			test: func(a args) bool {
//				return a.dBPerfInfoRepository.CurrentLog.Message == "first"
//			},
//		},
//		{
//			name: "append to first",
//			args: func() args {
//				dBPerfInfoRepository := stucts.InitDBPerfInfoRepository()
//				dBPerfInfoRepository.CurrentLog.Message = "first"
//				return args{
//					extractedLog:         stucts.ExtractedLog{Message: "second"},
//					dBPerfInfoRepository: dBPerfInfoRepository,
//				}
//			}(),
//			wantErr: false,
//			test: func(a args) bool {
//				return a.dBPerfInfoRepository.CurrentLog.Message == "first second"
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := ParseMessageOnlyLogV1(tt.args.extractedLog, tt.args.dBPerfInfoRepository); (err != nil) != tt.wantErr && tt.test(tt.args) {
//				t.Errorf("ParseMessageOnlyLogV1() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
