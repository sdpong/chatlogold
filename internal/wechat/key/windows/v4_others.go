//go:build !windows

package windows

import (
	"context"

	"github.com/sdpong/chatlogold/internal/wechat/model"
)

func (e *V4Extractor) Extract(ctx context.Context, proc *model.Process) (string, string, error) {
	return "", "", nil
}
