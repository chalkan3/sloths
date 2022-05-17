package jump

import (
	"context"
	"encoding/json"

	nt "github.com/go-kit/kit/transport/nats"
	"github.com/nats-io/nats.go"
)

func SubscriberTransport(svc Service, nc *nats.Conn) error {
	errs := make(chan error)

	jumpSubscriberHandler := nt.NewSubscriber(
		JumpEndpoint(svc),
		decodeCountRequest,
		nt.EncodeJSONResponse,
	)

	_, err := nc.Subscribe("sloth", jumpSubscriberHandler.ServeMsg(nc))
	if err != nil {
		return err
	}
	nc.Flush()

	<-errs

	return nil
}

func decodeCountRequest(_ context.Context, msg *nats.Msg) (interface{}, error) {
	var request JumpRequest

	if err := json.Unmarshal(msg.Data, &request); err != nil {
		return nil, err
	}

	return request, nil
}
