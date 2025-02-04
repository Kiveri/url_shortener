package usecase

import (
	"errors"
	"testing"

	"url/internal/domain/model"
	"url/internal/usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateUrlUseCase(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	type fields struct {
		urlRepo *mocks.UrlRepo
	}

	type args struct {
		req CreateUrlReq
	}

	tests := []struct {
		name    string
		args    args
		want    *model.URL
		wantErr bool
		before  func(f fields, args args)
	}{
		{
			name: "success",
			args: args{
				req: CreateUrlReq{
					LongURL: "https://golang.howtos.io/building-a-url-shortening-service-in-go/#prerequisites",
				},
			},
			want: &model.URL{
				ShortUrl: "166V4ZsOji",
				LongUrl:  "https://golang.howtos.io/building-a-url-shortening-service-in-go/#prerequisites",
			},
			before: func(f fields, args args) {
				url := model.NewURL("166V4ZsOji", args.req.LongURL)
				f.urlRepo.EXPECT().CreateUrl(url).Return(url, nil)
			},
		},
		{
			name: "error on create",
			args: args{
				req: CreateUrlReq{
					LongURL: "https://golang.howtos.io/building-a-url-shortening-service-in-go/#prerequisites",
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				url := model.NewURL("166V4ZsOji", args.req.LongURL)
				f.urlRepo.EXPECT().CreateUrl(url).Return(nil, errTest)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			f := fields{
				urlRepo: mocks.NewUrlRepo(t),
			}
			tt.before(f, tt.args)

			uc := NewUseCase(f.urlRepo)

			url, err := uc.CreateUrl(tt.args.req)

			if tt.wantErr {
				a.Error(err)

				return
			}
			a.NoError(err)
			a.Equal(tt.want, url)
		})
	}
}
