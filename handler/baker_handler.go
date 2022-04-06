package handler

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/Sakamoto0525/gRPC-Tutorial/gen/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func init() {
	// パンケーキの仕上がりに影響するseedを初期化
	rand.Seed(time.Now().UnixNano())
}

// BakerHandlerはパンケーキを焼く
type BakerHandler struct {
	report *report
}

type report struct {
	sync.Mutex // 複数人が同時に焼いても大丈夫なように
	data       map[api.Pancake_Menu]int
}

// NewBakerHandlerはBakerHandlerを初期化
func NewBakerHandler() *BakerHandler {
	return &BakerHandler{
		report: &report{
			data: make(map[api.Pancake_Menu]int),
		},
	}
}

// Bakeは指定されたメニューのパンケーキを焼いて、焼けたパンをResponseとして返す
func (h *BakerHandler) Bake(
	ctx context.Context,
	req *api.BakeRequest,
) (*api.BakeResponse, error) {
	// validation
	if req.Menu == api.Pancake_UNKNOWN ||
		req.Menu > api.Pancake_SPICY_CURRY {
		return nil, status.Errorf(codes.InvalidArgument, "パンケーキを選んでください。")
	}

	// パンケーキを焼いて、数を記録する
	now := time.Now()
	h.report.Lock()
	h.report.data[req.Menu] = h.report.data[req.Menu] + 1
	h.report.Unlock()

	// レスポンスを作って返す
	return &api.BakeResponse{
		Pancake: &api.Pancake{
			Menu:           req.Menu,
			ChefName:       "Sakamoto",
			TechnicalScore: rand.Float32(),
			CreateTime: &timestamppb.Timestamp{
				Seconds: now.Unix(),
				Nanos:   int32(now.Nanosecond()),
			},
		},
	}, nil
}

// Reportは焼けたパンの数を報告します
func (h *BakerHandler) Report(
	ctx context.Context,
	req *api.ReportRequest,
) (*api.ReportResponse, error) {

	counts := make([]*api.Report_BakeCount, len(h.report.data))

	// レポート作成
	h.report.Lock()
	for k, v := range h.report.data {
		counts = append(counts, &api.Report_BakeCount{
			Menu:  k,
			Const: int32(v),
		})
	}
	h.report.Unlock()

	// responseを返す
	return &api.ReportResponse{
		Report: &api.Report{
			BakeCounts: counts,
		},
	}, nil
}
