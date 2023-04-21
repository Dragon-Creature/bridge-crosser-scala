package rest

import (
	"bytes"
	"encoding/json"
	"git.ssns.se/git/frozendragon/bridge-crosser-scala/internal/model"
	"git.ssns.se/git/frozendragon/bridge-crosser-scala/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostCalculateCrossingError(t *testing.T) {
	s := &service.InterfacesMock{
		CalculateCrossingFunc: func(request model.CrossingRequest) model.CrossingResponse {
			return model.CrossingResponse{}
		},
	}
	re := Rest{
		Service: s,
	}

	callEndpoint := func(body string) (model.Error, *httptest.ResponseRecorder) {
		request, err := http.NewRequest(http.MethodPost, "/api/calculate_crossing", bytes.NewReader([]byte(body)))
		require.NoError(t, err)

		resp := httptest.NewRecorder()
		re.PostCalculateCrossing(resp, request)

		data, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var actual model.Error
		err = json.Unmarshal(data, &actual)
		require.NoError(t, err)

		return actual, resp
	}

	t.Run("No body", func(t *testing.T) {
		actual, resp := callEndpoint("")

		expected := model.Error{
			HttpCode: http.StatusBadRequest,
			Message:  "unable to parse body",
		}

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Equal(t, expected, actual)
	})

	t.Run("empty body", func(t *testing.T) {
		actual, resp := callEndpoint("{}")

		expected := model.Error{
			HttpCode: http.StatusUnprocessableEntity,
			Message:  "bridges: non zero value required",
		}

		assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
		assert.Equal(t, expected, actual)
	})
}
