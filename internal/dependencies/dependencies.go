package dependencies

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/siyoga/jwt-auth-boilerplate/internal/adapter"
	"github.com/siyoga/jwt-auth-boilerplate/internal/config"
	"github.com/siyoga/jwt-auth-boilerplate/internal/database"
	"github.com/siyoga/jwt-auth-boilerplate/internal/handler"
	"github.com/siyoga/jwt-auth-boilerplate/internal/log"
	"github.com/siyoga/jwt-auth-boilerplate/internal/repository"
	"github.com/siyoga/jwt-auth-boilerplate/internal/server"
	"github.com/siyoga/jwt-auth-boilerplate/internal/service"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Dependencies interface {
		Close()
		WaitForInterrupt()

		HttpServer() server.Server
	}

	dependencies struct {
		cfg               config.Config
		log               log.Logger
		shutdownCallbacks []func()
		shutdownCh        chan os.Signal

		// http stuff
		server      server.Server
		reqHandler  handler.RequestHandler
		middlewares handler.Middleware

		authHandler handler.Handler

		// services
		authService service.AuthService

		// client
		psqlClient *database.PostgresClient

		// repos
		jwtRepo  repository.JwtRepository
		txRepo   repository.TxRepository
		userRepo repository.UserRepository

		// adapters
		timeAdapter   adapter.TimeAdapter
		randomAdapter adapter.RandomAdapter
	}
)

func NewDependencies(cfgPath string) (Dependencies, error) {
	cfg, err := config.NewConfig(cfgPath)
	if err != nil && err.Error() == "Config File \"config\" Not Found in \"[]\"" {
		cfg, err = config.NewConfig("./configs/local")
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.LevelKey = "lvl"
	encoderCfg.TimeKey = "t"
	z := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			zap.NewAtomicLevel(),
		),
		zap.AddCaller(),
	)

	return &dependencies{
		cfg:        *cfg,
		log:        log.NewLogger(z),
		shutdownCh: make(chan os.Signal),
	}, nil
}

func (d *dependencies) Close() {
	for i := len(d.shutdownCallbacks) - 1; i >= 0; i-- {
		d.shutdownCallbacks[i]()
	}
	d.log.Zap().Sync()
}

func (d *dependencies) HttpServer() server.Server {
	if d.server == nil {
		var err error
		msg := "initialize app server"
		if d.server, err = server.NewHttpServer(
			d.log,
			d.cfg.Base,
			d.Middlewares(),
			d.AuthHandler(),
		); err != nil {
			d.log.Zap().Panic(msg, zap.Error(err))
		}

		d.shutdownCallbacks = append(d.shutdownCallbacks, func() {
			msg := "shutting down app server"
			if err := d.server.Stop(); err != nil {
				d.log.Zap().Warn(msg, zap.Error(err))
				return
			}
			d.log.Zap().Info(msg)
		})
	}

	return d.server
}

func (d *dependencies) WaitForInterrupt() {
	signal.Notify(d.shutdownCh, syscall.SIGINT, syscall.SIGTERM)
	d.log.Zap().Info("Wait for receive interrupt signal")
	<-d.shutdownCh
	d.log.Zap().Info("Receive interrupt signal")
}
