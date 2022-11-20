// Задание
// Определена API со следующим интерфейсом:
// type APIClient interface {
//     GetData(query string) (Response, error)
// }

// type Response struct {
//     Text       string
//     StatusCode int
// } 

// Реализуйте тип Mock с интерфейсом MockAPIClient:
// type MockAPIClient interface {
//     APIClient
//     SetResponse(resp Response, err error)
// }

// Ответ:
// Файл main.go:
// package main

// type APIClient interface {
//     GetData(query string) (Response, error)
// }

// type Response struct {
//     Text       string
//     StatusCode int
// }

// type MockAPIClient interface {
//     APIClient
//     SetResponse(resp Response, err error)
// }

// type Mock struct {
//     Resp Response
//     Err  error
// }

// func (m *Mock) GetData(query string) (Response, error) {
//     return m.Resp, m.Err
// }

// func (m *Mock) SetResponse(resp Response, err error) {
//     m.Resp = resp
//     m.Err = err
// }


// Файл main_test.go:
// package main

// import (
//     "errors"
//     "testing"
// )

// func TestMock(t *testing.T) {
//     tbl := []struct {
//         query string
//         resp  Response
//         err   error
//     }{
//         {"success", Response{"Success", 200}, nil},
//         {"error", Response{"", 500}, errors.New("something is wrong")},
//         {"empty", Response{"", 404}, nil},
//     }
//     for _, item := range tbl {
//         m := &Mock{}
//         m.SetResponse(item.resp, item.err)
//         resp, err := m.GetData(item.query)
//         if resp.Text != item.resp.Text {
//             t.Errorf(`%s: want %v got %v`, item.query, item.resp.Text, resp.Text)
//         }
//         if !errors.Is(err, item.err) {
//             t.Errorf(`%s: want %v got %v`, item.query, item.err, err)
//         }
//     }
// }