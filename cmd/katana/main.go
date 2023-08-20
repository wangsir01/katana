package katana

import (
	"github.com/projectdiscovery/gologger"
	"github.com/wangsir01/katana/internal/runner"
	"github.com/wangsir01/katana/pkg/types"
	"math"
	"os"
	"os/signal"
	"syscall"
)

var (
	//cfgFile string
	Options = &types.Options{}
)

//type Client struct {
//	Options types.Options
//}

func Crawler(url string) {
	//flagSet, err := readFlags()
	//if err != nil {
	//	gologger.Fatal().Msgf("Could not read flags: %s\n", err)
	//}
	Options.URLs = []string{url}
	if Options.HealthCheck {
		gologger.Fatal().Msgf("Options Failed to check\n")
		os.Exit(0)
	}

	katanaRunner, err := runner.New(Options)
	if err != nil || katanaRunner == nil {
		if Options.Version {
			return
		}
		gologger.Fatal().Msgf("could not create runner: %s\n", err)
	}
	defer katanaRunner.Close()

	// close handler
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		for range c {
			gologger.DefaultLogger.Info().Msg("- Ctrl+C pressed in Terminal")
			katanaRunner.Close()
			os.Exit(0)
		}
	}()

	if err := katanaRunner.ExecuteCrawling(); err != nil {
		gologger.Fatal().Msgf("could not execute crawling: %s", err)
	}
}

//func readFlags() (*goflags.FlagSet, error) {
//	flagSet := goflags.NewFlagSet()
//	flagSet.SetDescription(`Katana is a fast crawler focused on execution in automation
//pipelines offering both headless and non-headless crawling.`)
//
//	flagSet.CreateGroup("input", "Input",
//		flagSet.StringSliceVarP(&Options.URLs, "list", "u", nil, "target url / list to crawl", goflags.FileCommaSeparatedStringSliceOptions),
//	)
//
//	flagSet.CreateGroup("config", "Configuration",
//		flagSet.StringSliceVarP(&Options.Resolvers, "resolvers", "r", nil, "list of custom resolver (file or comma separated)", goflags.FileCommaSeparatedStringSliceOptions),
//		flagSet.IntVarP(&Options.MaxDepth, "depth", "d", 3, "maximum depth to crawl"),
//		flagSet.BoolVarP(&Options.ScrapeJSResponses, "js-crawl", "jc", false, "enable endpoint parsing / crawling in javascript file"),
//		flagSet.BoolVarP(&Options.ScrapeJSLuiceResponses, "jsluice", "jsl", false, "enable jsluice parsing in javascript file (memory intensive)"),
//		flagSet.DurationVarP(&Options.CrawlDuration, "crawl-duration", "ct", 0, "maximum duration to crawl the target for (s, m, h, d) (default s)"),
//		flagSet.StringVarP(&Options.KnownFiles, "known-files", "kf", "", "enable crawling of known files (all,robotstxt,sitemapxml)"),
//		flagSet.IntVarP(&Options.BodyReadSize, "max-response-size", "mrs", math.MaxInt, "maximum response size to read"),
//		flagSet.IntVar(&Options.Timeout, "timeout", 10, "time to wait for request in seconds"),
//		flagSet.BoolVarP(&Options.AutomaticFormFill, "automatic-form-fill", "aff", false, "enable automatic form filling (experimental)"),
//		flagSet.BoolVarP(&Options.FormExtraction, "form-extraction", "fx", false, "extract form, input, textarea & select elements in jsonl output"),
//		flagSet.IntVar(&Options.Retries, "retry", 1, "number of times to retry the request"),
//		flagSet.StringVar(&Options.Proxy, "proxy", "", "http/socks5 proxy to use"),
//		flagSet.StringSliceVarP(&Options.CustomHeaders, "headers", "H", nil, "custom header/cookie to include in all http request in header:value format (file)", goflags.FileStringSliceOptions),
//		//flagSet.StringVar(&cfgFile, "config", "", "path to the katana configuration file"),
//		flagSet.StringVarP(&Options.FormConfig, "form-config", "fc", "", "path to custom form configuration file"),
//		flagSet.StringVarP(&Options.FieldConfig, "field-config", "flc", "", "path to custom field configuration file"),
//		flagSet.StringVarP(&Options.Strategy, "strategy", "s", "depth-first", "Visit strategy (depth-first, breadth-first)"),
//		flagSet.BoolVarP(&Options.IgnoreQueryParams, "ignore-query-params", "iqp", false, "Ignore crawling same path with different query-param values"),
//		flagSet.BoolVarP(&Options.TlsImpersonate, "tls-impersonate", "tlsi", false, "enable experimental client hello (ja3) tls randomization"),
//	)
//
//	flagSet.CreateGroup("debug", "Debug",
//		flagSet.BoolVarP(&Options.HealthCheck, "hc", "health-check", false, "run diagnostic check up"),
//		flagSet.StringVarP(&Options.ErrorLogFile, "error-log", "elog", "", "file to write sent requests error log"),
//	)
//
//	flagSet.CreateGroup("headless", "Headless",
//		flagSet.BoolVarP(&Options.Headless, "headless", "hl", false, "enable headless hybrid crawling (experimental)"),
//		flagSet.BoolVarP(&Options.UseInstalledChrome, "system-chrome", "sc", false, "use local installed chrome browser instead of katana installed"),
//		flagSet.BoolVarP(&Options.ShowBrowser, "show-browser", "sb", false, "show the browser on the screen with headless mode"),
//		flagSet.StringSliceVarP(&Options.HeadlessOptionalArguments, "headless-Options", "ho", nil, "start headless chrome with additional Options", goflags.FileCommaSeparatedStringSliceOptions),
//		flagSet.BoolVarP(&Options.HeadlessNoSandbox, "no-sandbox", "nos", false, "start headless chrome in --no-sandbox mode"),
//		flagSet.StringVarP(&Options.ChromeDataDir, "chrome-data-dir", "cdd", "", "path to store chrome browser data"),
//		flagSet.StringVarP(&Options.SystemChromePath, "system-chrome-path", "scp", "", "use specified chrome browser for headless crawling"),
//		flagSet.BoolVarP(&Options.HeadlessNoIncognito, "no-incognito", "noi", false, "start headless chrome without incognito mode"),
//		flagSet.StringVarP(&Options.ChromeWSUrl, "chrome-ws-url", "cwu", "", "use chrome browser instance launched elsewhere with the debugger listening at this URL"),
//		flagSet.BoolVarP(&Options.XhrExtraction, "xhr-extraction", "xhr", false, "extract xhr request url,method in jsonl output"),
//	)
//
//	flagSet.CreateGroup("scope", "Scope",
//		flagSet.StringSliceVarP(&Options.Scope, "crawl-scope", "cs", nil, "in scope url regex to be followed by crawler", goflags.FileCommaSeparatedStringSliceOptions),
//		flagSet.StringSliceVarP(&Options.OutOfScope, "crawl-out-scope", "cos", nil, "out of scope url regex to be excluded by crawler", goflags.FileCommaSeparatedStringSliceOptions),
//		flagSet.StringVarP(&Options.FieldScope, "field-scope", "fs", "rdn", "pre-defined scope field (dn,rdn,fqdn)"),
//		flagSet.BoolVarP(&Options.NoScope, "no-scope", "ns", false, "disables host based default scope"),
//		flagSet.BoolVarP(&Options.DisplayOutScope, "display-out-scope", "do", false, "display external endpoint from scoped crawling"),
//	)
//
//	//availableFields := strings.Join(output.FieldNames, ",")
//	flagSet.CreateGroup("filter", "Filter",
//		flagSet.StringSliceVarP(&Options.OutputMatchRegex, "match-regex", "mr", nil, "regex or list of regex to match on output url (cli, file)", goflags.FileStringSliceOptions),
//		flagSet.StringSliceVarP(&Options.OutputFilterRegex, "filter-regex", "fr", nil, "regex or list of regex to filter on output url (cli, file)", goflags.FileStringSliceOptions),
//		flagSet.StringVarP(&Options.Fields, "field", "f", "", fmt.Sprintf("field to display in output (%s)", availableFields)),
//		flagSet.StringVarP(&Options.StoreFields, "store-field", "sf", "", fmt.Sprintf("field to store in per-host output (%s)", availableFields)),
//		flagSet.StringSliceVarP(&Options.ExtensionsMatch, "extension-match", "em", nil, "match output for given extension (eg, -em php,html,js)", goflags.CommaSeparatedStringSliceOptions),
//		flagSet.StringSliceVarP(&Options.ExtensionFilter, "extension-filter", "ef", nil, "filter output for given extension (eg, -ef png,css)", goflags.CommaSeparatedStringSliceOptions),
//		flagSet.StringVarP(&Options.OutputMatchCondition, "match-condition", "mdc", "", "match response with dsl based condition"),
//		flagSet.StringVarP(&Options.OutputFilterCondition, "filter-condition", "fdc", "", "filter response with dsl based condition"),
//	)
//
//	flagSet.CreateGroup("ratelimit", "Rate-Limit",
//		flagSet.IntVarP(&Options.Concurrency, "concurrency", "c", 10, "number of concurrent fetchers to use"),
//		flagSet.IntVarP(&Options.Parallelism, "parallelism", "p", 10, "number of concurrent inputs to process"),
//		flagSet.IntVarP(&Options.Delay, "delay", "rd", 0, "request delay between each request in seconds"),
//		flagSet.IntVarP(&Options.RateLimit, "rate-limit", "rl", 150, "maximum requests to send per second"),
//		flagSet.IntVarP(&Options.RateLimitMinute, "rate-limit-minute", "rlm", 0, "maximum number of requests to send per minute"),
//	)
//
//	flagSet.CreateGroup("update", "Update",
//		flagSet.CallbackVarP(runner.GetUpdateCallback(), "update", "up", "update katana to latest version"),
//		flagSet.BoolVarP(&Options.DisableUpdateCheck, "disable-update-check", "duc", false, "disable automatic katana update check"),
//	)
//
//	flagSet.CreateGroup("output", "Output",
//		flagSet.StringVarP(&Options.OutputFile, "output", "o", "", "file to write output to"),
//		flagSet.BoolVarP(&Options.StoreResponse, "store-response", "sr", false, "store http requests/responses"),
//		flagSet.StringVarP(&Options.StoreResponseDir, "store-response-dir", "srd", "", "store http requests/responses to custom directory"),
//		flagSet.BoolVarP(&Options.OmitRaw, "omit-raw", "or", false, "omit raw requests/responses from jsonl output"),
//		flagSet.BoolVarP(&Options.OmitBody, "omit-body", "ob", false, "omit response body from jsonl output"),
//		flagSet.BoolVarP(&Options.JSON, "jsonl", "j", false, "write output in jsonl format"),
//		flagSet.BoolVarP(&Options.NoColors, "no-color", "nc", false, "disable output content coloring (ANSI escape codes)"),
//		flagSet.BoolVar(&Options.Silent, "silent", false, "display output only"),
//		flagSet.BoolVarP(&Options.Verbose, "verbose", "v", false, "display verbose output"),
//		flagSet.BoolVar(&Options.Debug, "debug", false, "display debug output"),
//		flagSet.BoolVar(&Options.Version, "version", false, "display project version"),
//	)
//
//	if err := flagSet.Parse(); err != nil {
//		return nil, errorutil.NewWithErr(err).Msgf("could not parse flags")
//	}
//
//	//if cfgFile != "" {
//	//	if err := flagSet.MergeConfigFile(cfgFile); err != nil {
//	//		return nil, errorutil.NewWithErr(err).Msgf("could not read config file")
//	//	}
//	//}
//	return flagSet, nil
//}

func init() {
	Options = &types.Options{
		// Configuration
		Resolvers:         nil,           // custom resolvers
		MaxDepth:          1,             // maximum depth to crawl
		ScrapeJSResponses: true,          // enable endpoint parsing / crawling in javascript file
		CrawlDuration:     0,             // maximum duration to crawl the target for
		KnownFiles:        "all",         // enable crawling of known files
		BodyReadSize:      math.MaxInt,   // maximum response size to read
		Timeout:           10,            // time to wait for request in seconds
		AutomaticFormFill: true,          // enable automatic form filling
		Retries:           1,             // number of times to retry the request
		Proxy:             "",            // http/socks5 proxy to use
		CustomHeaders:     nil,           // custom headers or cookies
		FormConfig:        "",            // path to custom form configuration file
		FieldConfig:       "",            // path to custom field configuration file
		Strategy:          "depth-first", // Visit strategy (depth-first, breadth-first)
		IgnoreQueryParams: false,         // Ignore crawling same path with different query-param values
		// Debug
		HealthCheck:  false, // run diagnostic check up
		ErrorLogFile: "",    // file to write sent requests error log
		// Headless conf
		Headless:                  true,  // enable headless hybrid crawling
		UseInstalledChrome:        false, // use local installed Chrome browser instead of katana installed
		ShowBrowser:               false, // show the browser on the screen with headless mode
		HeadlessOptionalArguments: nil,   // start headless chrome with additional Options
		HeadlessNoSandbox:         true,  // start headless chrome in --no-sandbox mode
		ChromeDataDir:             "",    // path to store chrome browser data
		SystemChromePath:          "",    // use specified Chrome browser for headless crawling
		HeadlessNoIncognito:       false, // start headless chrome without incognito mode
		// Scope conf
		Scope:           []string{}, // in scope url regex to be followed by crawler
		OutOfScope:      []string{}, // out of scope url regex to be excluded by crawler
		FieldScope:      "rdn",      // pre-defined scope field
		NoScope:         false,      // disables host based default scope
		DisplayOutScope: false,      // display external endpoint from scoped crawling
		// Filter conf
		OutputMatchRegex:  []string{},      // regex or list of regex to match on output url
		OutputFilterRegex: []string{},      // regex or list of regex to filter on output url
		Fields:            "",              // field to display in output
		StoreFields:       "",              // field to store in per-host output
		ExtensionsMatch:   []string{},      // match output for given extension
		ExtensionFilter:   []string{"css"}, // filter output for given extension
		// Rate-Limit conf
		Concurrency:     10,  // number of concurrent fetchers to use
		Parallelism:     10,  // number of concurrent inputs to process
		Delay:           0,   // request delay between each request in seconds
		RateLimit:       150, // maximum requests to send per second
		RateLimitMinute: 0,   // maximum number of requests to send per minute
		// Update conf
		DisableUpdateCheck: true, // disable automatic katana update check
		// Output conf
		OutputFile:       "",    // file to write output to
		StoreResponse:    false, // store http requests/responses
		StoreResponseDir: "",    // store http requests/responses to custom directory
		OmitRaw:          false, // omit raw requests/responses from jsonl output
		OmitBody:         false, // omit response body from jsonl output
		JSON:             false, // write output in jsonl format
		NoColors:         false, // disable output content coloring
		Silent:           true,  // display output only
		Verbose:          false, // display verbose output
		Debug:            false, // display debug output
		Version:          false,

		//OnResult: func(result output.Result) { // Callback function to execute for result
		//	select {
		//	case OutputDataCh <- result:
		//		// 如果能够执行这一行，说明通道没有阻塞
		//	default:
		//		// 如果执行到这一行，说明通道被阻塞了
		//		fmt.Println("通道阻塞")
		//	}
		//},
	}
}

//func init() {
//	// show detailed stacktrace in debug mode
//	if os.Getenv("DEBUG") == "true" {
//		errorutil.ShowStackTrace = true
//	}
//}
