package logger

import "context"

type Writer struct{}

func (w *Writer) writeToExternal(ctx context.Context, data []byte) {}
