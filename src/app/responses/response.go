package responses

import "net/http"

const TEXT = "text"
const JSON = "json"

type Response struct {
    HttpStatusCode int
    Type           string
    Content        interface{}
}

func Json(data interface{}) *Response {
    switch data.(type) {
    case AuthError:
        return &Response{
            HttpStatusCode: http.StatusUnauthorized,
            Type:           JSON,
            Content:        data,
        }

    case ValidationError:
        return &Response{
            HttpStatusCode: http.StatusUnprocessableEntity,
            Type:           JSON,
            Content:        data,
        }

    default:
        return &Response{
            HttpStatusCode: http.StatusOK,
            Type:           JSON,
            Content:        data,
        }
    }
}

type ValidationError struct {
    Code    string            `json:"code"`
    Message string            `json:"message"`
    Errors  map[string]string `json:"errors"`
}

type AuthError struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

type Error struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}
