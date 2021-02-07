package responses

const STRING = "string"
const JSON = "json"

type Response struct {
    Content interface{}
    Type    string
}

func Json(data interface{}) *Response {
    return &Response{
        Type:    JSON,
        Content: data,
    }
}

type Error struct {
    Code    string
    Message string
    //Link    string
}
