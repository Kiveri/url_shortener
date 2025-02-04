package usecase

import (
	"errors"
	"testing"

	"url/internal/domain/model"
	"url/internal/usecase/mocks"

	"github.com/stretchr/testify/assert"
)

func TestFindUrlUseCase(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	type fields struct {
		urlRepo *mocks.UrlRepo
	}

	type args struct {
		req FindUrlReq
	}

	url := &model.URL{
		ShortUrl: "166V4ZsOji",
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
				req: FindUrlReq{
					ShortUrl: url.ShortUrl,
				},
			},
			want: &model.URL{
				ShortUrl: url.ShortUrl,
			},
			before: func(f fields, args args) {
				f.urlRepo.EXPECT().FindUrl(args.req.ShortUrl).Return(url, nil)
			},
		},
		{
			name: "error on find",
			args: args{
				req: FindUrlReq{
					ShortUrl: "1234567890",
				},
			},
			wantErr: true,
			before: func(f fields, args args) {
				f.urlRepo.EXPECT().FindUrl(args.req.ShortUrl).Return(nil, errTest)
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

			u, err := uc.FindUrl(tt.args.req)
			if tt.wantErr {
				a.Error(err)

				return
			}
			a.NoError(err)
			a.Equal(tt.want, u)
		})
	}
}
