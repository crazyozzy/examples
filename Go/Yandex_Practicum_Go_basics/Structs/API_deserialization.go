// Есть пример API-вызова в формате JSON:
// {
//     "header": {
//         "code": 0,
//         "message": ""
//     },
//     "data": [{
//         "type": "user",
//         "id": 100,
//         "attributes": {
//             "email": "bob@yandex.ru",
//             "article_ids": [10, 11, 12]
//         }
//     }]
// }
// К сожалению, ни Swagger-описания, ни статьи с API-ответом в любимом сервисе заметок — нет. Опишите данный объект в виде структуры на Go, в учебных целях отбросив «так делать нельзя» и «это не дело».
// На входе есть строка с сырыми данными, требуется написать функцию её десериализации:
// type Response struct {
//     // поля с тегами
// }

// func ReadResponse(rawResp string) (Response, error) {
//     // код десериализации
// }

package main

import (
	"encoding/json"
	"fmt"
)

const rawResp = `
{
    "header": {
        "code": 0,
        "message": ""
    },
    "data": [{
        "type": "user",
        "id": 100,
        "attributes": {
            "email": "bob@yandex.ru",
            "article_ids": [10, 11, 12]
        }
    }]
}
`

type (
	Response struct {
		Header ResponseHeader `json:"header"`
		Data ResponseData `json:"data,omitempty"`
	}

	ResponseHeader struct {
		Code int `json:"code"`
		Message string `json:"message,omitempty"`
	}

	ResponseData []ResponseDataElem

	ResponseDataElem struct {
		Type string `json:"type"`
		Id int `json:"id"`
		Attributes ResponseDataElemAttributes `json:"attributes"`
	}

	ResponseDataElemAttributes struct {
		Email string `json:"email"`
		Article_ids []int `json:"article_ids"`
	}
)


func ReadResponse(rawResp string) (Response, error) {
	resp := Response{}
    err := json.Unmarshal([]byte(rawResp), &resp)

	return resp, err
}

func main() {
	fmt.Println(ReadResponse(rawResp))
}


// Ответ:
// package main

// import (
//     "encoding/json"
//     "fmt"
// )

// const rawResp = `
// {
//     "header": {
//         "code": 0,
//         "message": ""
//     },
//     "data": [{
//         "type": "user",
//         "id": 100,
//         "attributes": {
//             "email": "bob@yandex.ru",
//             "article_ids": [10, 11, 12]
//         }
//     }]
// }
// `

// type (
//     Response struct {
//         Header ResponseHeader `json:"header"`
//         Data   ResponseData   `json:"data,omitempty"`
//     }

//     ResponseHeader struct {
//         Code    int    `json:"code"`
//         Message string `json:"message,omitempty"`
//     }

//     ResponseData []ResponseDataItem

//     ResponseDataItem struct {
//         Type       string                `json:"type"`
//         Id         int                   `json:"id"`
//         Attributes ResponseDataItemAttrs `json:"attributes"`
//     }

//     ResponseDataItemAttrs struct {
//         Email      string `json:"email"`
//         ArticleIds []int  `json:"article_ids"`
//     }
// )

// func ReadResponse(rawResp string) (Response, error) {
//     resp := Response{}
//     if err := json.Unmarshal([]byte(rawResp), &resp); err != nil {
//         return Response{}, fmt.Errorf("JSON unmarshal: %w", err)
//     }

//     return resp, nil
// }

// func main() {
//     resp, err := ReadResponse(rawResp)
//     if err != nil {
//         panic(err)
//     }
//     fmt.Printf("%+v\n", resp)
// }