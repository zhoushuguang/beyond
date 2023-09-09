package response

import "context"

func ErrHandler(ctx context.Context, err error) (int, any) {
	return 0, nil
}
