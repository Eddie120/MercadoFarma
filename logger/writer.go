package logger

import "context"

type Writer struct{}

func (w *Writer) sendToSQS(ctx context.Context, data []byte) {

}
