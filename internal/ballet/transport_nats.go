package ballet

import (
	"context"
	"encoding/json"

	nt "github.com/go-kit/kit/transport/nats"
	"github.com/nats-io/nats.go"
)

func SubscriberTransport(svc Service, nc *nats.Conn) error {
	errs := make(chan error)

	danceBalletSubscriberHandler := nt.NewSubscriber(
		DanceBalletEndpoint(svc),
		decodeDanceBalletRequest,
		nt.EncodeJSONResponse,
	)

	_, err := nc.Subscribe("sloth.dance.ballet", danceBalletSubscriberHandler.ServeMsg(nc))
	if err != nil {
		return err
	}
	nc.Flush()

	<-errs

	return nil
}

func decodeDanceBalletRequest(_ context.Context, msg *nats.Msg) (interface{}, error) {
	var request DanceBalletRequest

	if err := json.Unmarshal(msg.Data, &request); err != nil {
		return nil, err
	}

	return request, nil
}
