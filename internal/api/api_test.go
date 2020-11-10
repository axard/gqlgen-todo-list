package api

import (
	"context"
	"testing"
)

func TestUserIDFromContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Получение ID пользователя из пустой строки должно завершаться с ошибкой",
			args: args{
				context.WithValue(context.Background(), UserID, ""),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Получение ID пользователя из строки с буквами должно завершаться с ошибкой",
			args: args{
				context.WithValue(context.Background(), UserID, "five"),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Получение ID пользователя из типа отличного от строки должно завершаться с ошибкой",
			args: args{
				context.WithValue(context.Background(), UserID, 5),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Получение ID пользователя из строки с цифрами завершается без ошибок и возвращает переданное число",
			args: args{
				context.WithValue(context.Background(), UserID, "5"),
			},
			want:    5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UserIDFromContext(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserFromContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
