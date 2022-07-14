package uploader_test

import (
	"image"
	"testing"

	"github.com/adammy/memepen-services/pkg/uploader"
	"github.com/stretchr/testify/assert"
)

func TestNewInMemoryRepository(t *testing.T) {
	u := uploader.NewLocalUploader()

	assert.NotNil(t, u)
	assert.Implements(t, (*uploader.Uploader)(nil), u)
}

func TestLocalUploader_UploadPNG(t *testing.T) {
	tests := map[string]struct {
		path         string
		img          image.Image
		expectedPath string
		error        bool
	}{
		"valid": {
			path: "../../../assets/memes/my-id",
			img: image.NewRGBA(image.Rectangle{
				Min: image.Point{X: 0, Y: 0},
				Max: image.Point{X: 1, Y: 1},
			}),
		},
		"invalid path": {
			path:  "non-existent-folder/memes/my-id",
			img:   image.NewRGBA(image.Rectangle{}),
			error: true,
		},
		"invalid image": {
			path:  "../../../assets/memes/my-id",
			img:   image.NewRGBA(image.Rectangle{}),
			error: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			u := uploader.NewLocalUploader()
			err := u.UploadPNG(tc.path, tc.img)

			if !tc.error {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
