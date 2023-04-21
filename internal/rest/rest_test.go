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

func TestPostCalculateCrossing(t *testing.T) {
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

	t.Run("no hikers", func(t *testing.T) {
		actual, resp := callEndpoint(`{
   "bridges":[
      {
         "id":"24a433ac-6c34-40cc-951a-3c7c65947c8a",
         "hikers":[],
         "length_in_feet":100
      },
      {
         "id":"ceecaf78-64a4-2fbb-7258-822611e76604",
         "length_in_feet":250,
         "hikers":[
            
         ]
      },
      {
         "id":"9df1f5a2-b9f9-c91b-a115-5b961cf6567b",
         "length_in_feet":300,
         "hikers":[
            
         ]
      }
   ]
}`)

		expected := model.Error{
			HttpCode: http.StatusUnprocessableEntity,
			Message:  "you need at least one hiker",
		}

		assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)
		assert.Equal(t, expected, actual)
	})

	t.Run("one hiker", func(t *testing.T) {
		_, resp := callEndpoint(`{
   "bridges":[
      {
         "id":"24a433ac-6c34-40cc-951a-3c7c65947c8a",
         "hikers":[
            {
               "id":"44ab4f16-f7bb-716c-449c-67389615e735",
               "speed_feet_in_minutes":10
            }
         ],
         "length_in_feet":100
      },
      {
         "id":"ceecaf78-64a4-2fbb-7258-822611e76604",
         "length_in_feet":250,
         "hikers":[
            
         ]
      },
      {
         "id":"9df1f5a2-b9f9-c91b-a115-5b961cf6567b",
         "length_in_feet":300,
         "hikers":[
            
         ]
      }
   ]
}`)

		assert.Equal(t, http.StatusOK, resp.Code)
	})
}
