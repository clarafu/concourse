package emitter

import (
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/metric"
	"github.com/pkg/errors"

	influxclient "github.com/influxdata/influxdb1-client/v2"
)

type InfluxDBEmitter struct {
	Client        influxclient.Client
	Database      string
	BatchSize     int
	BatchDuration time.Duration
}

type InfluxDBConfig struct {
	Enabled bool   `yaml:"enabled,omitempty"`
	URL     string `yaml:"url,omitempty"`

	Database string `yaml:"database,omitempty"`

	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`

	InsecureSkipVerify bool `yaml:"insecure_skip_verify,omitempty"`

	BatchSize     uint32        `yaml:"batch_size,omitempty"`
	BatchDuration time.Duration `yaml:"batch_duration,omitempty"`
}

var (
	batch         []metric.Event
	lastBatchTime time.Time
)

func init() {
	batch = make([]metric.Event, 0)
	lastBatchTime = time.Now()
}

func (config *InfluxDBConfig) Description() string { return "InfluxDB" }
func (config *InfluxDBConfig) Validate() error {
	if config.URL == "" {
		return errors.New("url is missing")
	}

	return nil
}

func (config *InfluxDBConfig) NewEmitter() (metric.Emitter, error) {
	client, err := influxclient.NewHTTPClient(influxclient.HTTPConfig{
		Addr:               config.URL,
		Username:           config.Username,
		Password:           config.Password,
		InsecureSkipVerify: config.InsecureSkipVerify,
		Timeout:            time.Minute,
	})
	if err != nil {
		return &InfluxDBEmitter{}, err
	}

	return &InfluxDBEmitter{
		Client:        client,
		Database:      config.Database,
		BatchSize:     int(config.BatchSize),
		BatchDuration: config.BatchDuration,
	}, nil
}

func emitBatch(emitter *InfluxDBEmitter, logger lager.Logger, events []metric.Event) {

	logger.Debug("influxdb-emit-batch", lager.Data{
		"size": len(events),
	})
	bp, err := influxclient.NewBatchPoints(influxclient.BatchPointsConfig{
		Database: emitter.Database,
	})
	if err != nil {
		logger.Error("failed-to-construct-batch-points", err)
		return
	}

	for _, event := range events {
		tags := map[string]string{
			"host": event.Host,
		}

		for k, v := range event.Attributes {
			tags[k] = v
		}

		point, err := influxclient.NewPoint(
			event.Name,
			tags,
			map[string]interface{}{
				"value": event.Value,
			},
			event.Time,
		)
		if err != nil {
			logger.Error("failed-to-construct-point", err)
			continue
		}

		bp.AddPoint(point)
	}

	err = emitter.Client.Write(bp)
	if err != nil {
		logger.Error("failed-to-send-points",
			errors.Wrap(metric.ErrFailedToEmit, err.Error()))
		return
	}
}

func (emitter *InfluxDBEmitter) Emit(logger lager.Logger, event metric.Event) {
	batch = append(batch, event)
	duration := time.Since(lastBatchTime)
	if len(batch) >= emitter.BatchSize || duration >= emitter.BatchDuration {
		logger.Debug("influxdb-pre-emit-batch", lager.Data{
			"influxdb-batch-size":     emitter.BatchSize,
			"current-batch-size":      len(batch),
			"influxdb-batch-duration": emitter.BatchDuration,
			"current-duration":        duration,
		})
		emitter.SubmitBatch(logger)
	}
}

func (emitter *InfluxDBEmitter) SubmitBatch(logger lager.Logger) {
	batchToSubmit := make([]metric.Event, len(batch))
	copy(batchToSubmit, batch)
	batch = make([]metric.Event, 0)
	lastBatchTime = time.Now()
	go emitBatch(emitter, logger, batchToSubmit)
}
