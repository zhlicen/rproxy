package rproxy

import (
	"context"
	"log"
	"net/http"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/mwitkow/grpc-proxy/proxy"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc/metadata"
)

var grpcRequestTotal *prometheus.CounterVec

func init() {
	grpcRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_request_total",
			Help: "Number of grpc requests in total",
		},
		[]string{"function", "endpoint"},
	)
	prometheus.MustRegister(grpcRequestTotal)
}

const grpcWebContentType = "application/grpc-web"

// grpcWebServerWrap translates grpc-web to native gRPC
var grpcWebServerWrap *grpcweb.WrappedGrpcServer

func serveGRPCWeb(w http.ResponseWriter, r *http.Request) (handled bool) {
	accessControlHeaders := strings.ToLower(r.Header.Get("Access-Control-Request-Headers"))
	setupCORSHeaders := func(w http.ResponseWriter) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Referer, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web")
		w.Header().Set("grpc-status", "")
		w.Header().Set("grpc-message", "")
	}
	if r.Method == http.MethodOptions && strings.Contains(accessControlHeaders, "x-grpc-web") {
		setupCORSHeaders(w)
		return true
	}

	if grpcWebServerWrap != nil && grpcWebServerWrap.IsGrpcWebRequest(r) {
		setupCORSHeaders(w)
		grpcWebServerWrap.HandleGrpcWebRequest(w, r)
		return true
	}

	return false
}

func registerGRPCServer(listener cmux.CMux) {
	// grpcListener := listener.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	grpcListener := listener.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	grpcServer := grpc.NewServer(getGRPCProxyServerOptions()...)
	grpcWebServerWrap = grpcweb.WrapServer(grpcServer)

	go func() {
		err := grpcServer.Serve(grpcListener)
		if err != nil {
			log.Fatalln("grpcServer.Serve:", err)
		}
	}()

}

var proxyDialOptions []grpc.DialOption
var proxyDialOptionsInsecure []grpc.DialOption

func init() {
	proxyDialOptions = []grpc.DialOption{grpc.WithCodec(proxy.Codec())}
	proxyDialOptionsInsecure = []grpc.DialOption{grpc.WithCodec(proxy.Codec()), grpc.WithInsecure()}
}

func getGRPCProxyServerOptions() []grpc.ServerOption {
	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		if strings.HasPrefix(fullMethodName, "/com.example.internal.") {
			return nil, nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
		}
		md, ok := metadata.FromIncomingContext(ctx)
		outCtx, _ := context.WithCancel(ctx)
		outCtx = metadata.NewOutgoingContext(outCtx, md.Copy())
		if ok {
			var dst string
			insecure := false
			GRPC.rangeRule(func(servicePrefix string, rule *GRPCRule) bool {
				if strings.HasPrefix(fullMethodName, servicePrefix) {
					dst = rule.ForwardDst
					insecure = rule.Insecure
					return false
				}
				return true
			})

			if dst != "" {
				grpcRequestTotal.With(prometheus.Labels{"function": fullMethodName, "endpoint": dst}).Inc()
				log.Printf("GRPC call[%s] matched -> [%s]", fullMethodName, dst)
				dialOpts := proxyDialOptions
				if insecure {
					dialOpts = proxyDialOptionsInsecure
				}
				conn, err := grpc.DialContext(outCtx, dst, dialOpts...)
				return outCtx, conn, err
			}
			log.Printf("GRPC call[%s] no rules found", fullMethodName)
			
		}
		return nil, nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
	}

	return []grpc.ServerOption{
		grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(proxy.TransparentHandler(director)),
	}

}
