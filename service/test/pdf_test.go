package service_test

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const url string = "http://localhost:5001/convert"
const multi int = 5

func TestConvertPdfToJpeg(t *testing.T) {

	t.Run("convert one pdf to jpeg", func(t *testing.T) {
		body := []byte(`{"filename": "test"}`)
		res, err := http.Post(url, "application/json", bytes.NewReader(body))
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})

	t.Run("convert multiple pdf to jpeg", func(t *testing.T) {
		var wg sync.WaitGroup
		for i := 0; i < multi; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				body := []byte(fmt.Sprintf(`{"filename": "test-%d"}`, i))
				res, err := http.Post(url, "application/json", bytes.NewReader(body))
				if err != nil {
					t.Error(err)
				}
				assert.Equal(t, http.StatusCreated, res.StatusCode)
			}()
		}
		wg.Wait()
	})

	t.Run("convert inexistent pdf to jpeg", func(t *testing.T) {
		body := []byte(`{"filename": "___.pdf"}`)
		res, err := http.Post(url, "application/json", bytes.NewReader(body))
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("convert multiple inexistent pdf to jpeg", func(t *testing.T) {
		var wg sync.WaitGroup
		for i := 0; i < multi; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				body := []byte(fmt.Sprintf(`{"filename": "%d"}`, i))
				res, err := http.Post(url, "application/json", bytes.NewReader(body))
				if err != nil {
					t.Error(err)
				}
				assert.Equal(t, http.StatusBadRequest, res.StatusCode)
			}()
		}
		wg.Wait()
	})
}
