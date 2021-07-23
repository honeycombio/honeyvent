package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	libhoney "github.com/honeycombio/libhoney-go"
	flag "github.com/jessevdk/go-flags"
)

var BuildID string

type Options struct {
	APIHost string `hidden:"true" long:"api_host" description:"APIHost for the Honeycomb API" default:"https://api.honeycomb.io/"`

	WriteKey  string   `short:"k" long:"writekey" description:"Team write key" required:"true"`
	Dataset   string   `short:"d" long:"dataset" description:"Name of the dataset" required:"true"`
	Timestamp string   `short:"t" long:"timestamp" description:"Set a specific timestamp (RFC3339) for this event (override the default of now)"`
	Name      []string `short:"n" long:"name" description:"Metric name"`
	Val       []string `short:"v" long:"value" description:"Metric value"`
	Verbose   bool     `short:"V" long:"verbose" description:"Show output"`
}

// parseAPIHost parses the provided APIHost argument and sets sensible defaults if they are not provided.
func parseAPIHost(host string) (*url.URL, error) {
	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	if len(u.Scheme) == 0 {
		u.Scheme = "https"
	}
	return u, nil
}

func main() {
	var opts Options
	flagParser := flag.NewParser(&opts, flag.Default)
	if extraArgs, err := flagParser.Parse(); err != nil || len(extraArgs) != 0 {
		if flagErr, ok := err.(*flag.Error); ok {
			if flagErr.Type == flag.ErrHelp {
				// asking for help isn't a failed run.
				os.Exit(0)
			}
		}
		errAndExit("command line parsing error - call with --help for usage")
	}

	if len(opts.Name) != len(opts.Val) {
		errAndExit("Must have a value for each metric name - call with --help for usage")
	}

	u, err := parseAPIHost(opts.APIHost)
	if err != nil {
		errAndExit(fmt.Sprintf("Unable to parse API host: %s", err.Error()))
	}

	c := libhoney.Config{
		APIKey:  opts.WriteKey,
		Dataset: opts.Dataset,
		APIHost: u.String(),
	}
	libhoney.Init(c)
	defer libhoney.Close()
	libhoney.UserAgentAddition = fmt.Sprintf("honeyvent/%s", BuildID)

	ev := libhoney.NewEvent()

	if opts.Timestamp != "" {
		t1, err := time.Parse(time.RFC3339, opts.Timestamp)
		if err != nil {
			errAndExit(fmt.Sprintf("Unable to parse timestamp: %v", err))
		}
		ev.Timestamp = t1
	}
	for i, name := range opts.Name {
		if val, err := strconv.Atoi(opts.Val[i]); err == nil {
			ev.AddField(name, val)
		} else if val := strings.ToLower(opts.Val[i]); val == "nan" || val == "inf" || val == "-inf" {
			// special exception because we don't want these special floats to parse as float
			// since then the field will get dropped on ingestion
			ev.AddField(name, val)
		} else if val, err := strconv.ParseFloat(opts.Val[i], 64); err == nil {
			ev.AddField(name, val)
		} else {
			// add it as a string
			ev.AddField(name, opts.Val[i])
		}
	}
	if opts.Verbose {
		fmt.Println("sending event", ev)
	}
	ev.Send()
	rs := libhoney.TxResponses()
	rsp := <-rs
	if opts.Verbose {
		fmt.Printf("sent event %+v\n", map[string]interface{}{
			"status_code": rsp.StatusCode,
			"body":        strings.TrimSpace(string(rsp.Body)),
			"duration":    rsp.Duration,
			"error":       rsp.Err,
		})
	}

	if rsp.Err != nil {
		errAndExit(rsp.Err.Error())
	}
}

func errAndExit(reason string) {
	fmt.Printf("Error: %s\n", reason)
	os.Exit(1)
}
