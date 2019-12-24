package web_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/fernandoocampo/paymentgw/pkg/adapter/web"
	"github.com/fernandoocampo/paymentgw/pkg/portin"
	"github.com/stretchr/testify/mock"
)

func TestSuccessfulTransaction(t *testing.T) {
	service := new(testifyMock)
	service.On("Process", mock.AnythingOfType("*context.cancelCtx"), mock.AnythingOfType("*portin.NewPayment")).Return(anExpectedSuccessResult(), nil)
	ctx := context.TODO()
	endpoints := web.Endpoints{
		PaymentProcessorEndpoint: web.MakePaymentProcessorEndpoint(service),
	}
	mux := web.NewHTTPServer(ctx, endpoints)
	srv := httptest.NewServer(mux)
	defer srv.Close()
	want := `{"data":{"confirmation_id":"109c4af5-049c-4c3d-bcbe-703bcb7293c7","transaction_id":"1","company_key":"123","amount":50000.34,"success":true,"created_at":"2015-09-18T00:00:00Z"}}`

	givenJSONRequest := `{"transaction_id":"1","company_key":"123", "amount":50000.34}`
	req, errreq := http.NewRequest("POST", srv.URL+"/payments", bytes.NewBuffer([]byte(givenJSONRequest)))

	if errreq != nil {
		t.Fatal(errreq)
	}

	resp, _ := http.DefaultClient.Do(req)
	service.AssertExpectations(t)
	body, _ := ioutil.ReadAll(resp.Body)

	if have := strings.TrimSpace(string(body)); want != have {
		t.Errorf("%s: want %q, have %q", givenJSONRequest, want, have)
	}
}

func TestUnsuccessfulTransaction(t *testing.T) {
	service := new(testifyMock)
	service.On("Process", mock.AnythingOfType("*context.cancelCtx"), mock.AnythingOfType("*portin.NewPayment")).Return(anExpectedUnsuccessResult(), nil)
	ctx := context.TODO()
	endpoints := web.Endpoints{
		PaymentProcessorEndpoint: web.MakePaymentProcessorEndpoint(service),
	}
	mux := web.NewHTTPServer(ctx, endpoints)
	srv := httptest.NewServer(mux)
	defer srv.Close()
	want := `{"data":{"transaction_id":"1","company_key":"123","amount":50000.34,"success":false},"errors":["something went wrong"]}`

	givenJSONRequest := `{"transaction_id":"1","company_key":"123", "amount":50000.34}`
	req, errreq := http.NewRequest("POST", srv.URL+"/payments", bytes.NewBuffer([]byte(givenJSONRequest)))

	if errreq != nil {
		t.Fatal(errreq)
	}

	resp, _ := http.DefaultClient.Do(req)
	service.AssertExpectations(t)
	body, _ := ioutil.ReadAll(resp.Body)

	if have := strings.TrimSpace(string(body)); want != have {
		t.Errorf("%s: want %q, have %q", givenJSONRequest, want, have)
	}
}

func anExpectedSuccessResult() *portin.PaymentResult {
	ts := time.Date(2015, 9, 18, 0, 0, 0, 0, time.UTC)
	return &portin.PaymentResult{
		Amount:           50000.34,
		CompanyKey:       "123",
		ID:               "109c4af5-049c-4c3d-bcbe-703bcb7293c7",
		Success:          true,
		TxID:             "1",
		ValidationErrors: nil,
		Created:          &ts,
	}
}

func anExpectedUnsuccessResult() *portin.PaymentResult {
	return &portin.PaymentResult{
		Amount:           50000.34,
		CompanyKey:       "123",
		Success:          false,
		TxID:             "1",
		ValidationErrors: []error{fmt.Errorf("something went wrong")},
	}
}

type testifyMock struct {
	mock.Mock // just for academic purposes
}

func (t *testifyMock) Process(ctx context.Context, newPayment *portin.NewPayment) (*portin.PaymentResult, error) {
	args := t.Called(ctx, newPayment)
	return args.Get(0).(*portin.PaymentResult), args.Error(1)
}
