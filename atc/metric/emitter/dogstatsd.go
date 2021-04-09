package emitter

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"code.cloudfoundry.org/lager"
	"github.com/DataDog/datadog-go/statsd"
	"github.com/concourse/concourse/atc/metric"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

type DogstatsdEmitter struct {
	client *statsd.Client
}

type DogstatsDBConfig struct {
	Enabled bool   `yaml:"enabled,omitempty"`
	Host    string `yaml:"agent_host,omitempty"`
	Port    string `yaml:"agent_port,omitempty"`
	Prefix  string `yaml:"prefix,omitempty"`
}

func (config *DogstatsDBConfig) Description() string { return "Datadog" }

func (config *DogstatsDBConfig) Validate() error {
	var errs *multierror.Error

	if config.Host == "" {
		errs = multierror.Append(errs, errors.New("host is missing"))
	}

	if config.Port == "" {
		errs = multierror.Append(errs, errors.New("port is missing"))
	}

	return errs.ErrorOrNil()
}

func (config *DogstatsDBConfig) NewEmitter() (metric.Emitter, error) {

	client, err := statsd.New(fmt.Sprintf("%s:%s", config.Host, config.Port))
	if err != nil {
		log.Fatal(err)
		return &DogstatsdEmitter{}, err
	}

	if config.Prefix != "" {
		if strings.HasSuffix(config.Prefix, ".") {
			client.Namespace = config.Prefix
		} else {
			client.Namespace = fmt.Sprintf("%s.", config.Prefix)
		}
	}

	return &DogstatsdEmitter{
		client: client,
	}, nil
}

var specialChars = regexp.MustCompile("[^a-zA-Z0-9_]+")

func (emitter *DogstatsdEmitter) Emit(logger lager.Logger, event metric.Event) {
	name := specialChars.ReplaceAllString(strings.Replace(strings.ToLower(event.Name), " ", "_", -1), "")

	tags := []string{
		fmt.Sprintf("host:%s", event.Host),
	}

	for k, v := range event.Attributes {
		tags = append(tags, fmt.Sprintf("%s:%s", k, v))
	}

	err := emitter.client.Gauge(name, event.Value, tags, 1)
	if err != nil {
		logger.Error("failed-to-send-metric",
			errors.Wrap(metric.ErrFailedToEmit, err.Error()))
		return
	}
}
