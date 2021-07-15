package logger

import (
	"os"
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"

	"github.com/free5gc/logger_conf"
	"github.com/free5gc/logger_util"
)

var (
	log         *logrus.Logger
	JsonLog     *logrus.Logger
	SliceLog    *logrus.Entry
	AppLog      *logrus.Entry
	InitLog     *logrus.Entry
	CfgLog      *logrus.Entry
	GsmLog      *logrus.Entry
	PfcpLog     *logrus.Entry
	PduSessLog  *logrus.Entry
	CtxLog      *logrus.Entry
	ConsumerLog *logrus.Entry
	GinLog      *logrus.Entry
)

const SMF_SLICE_LOG = "/smf-slice-info.log"

func init() {
	log = logrus.New()
	log.SetReportCaller(false)

	log.Formatter = &formatter.Formatter{
		TimestampFormat: time.RFC3339,
		TrimMessages:    true,
		NoFieldsSpace:   true,
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
	}

	free5gcLogHook, err := logger_util.NewFileHook(logger_conf.Free5gcLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	if err == nil {
		log.Hooks.Add(free5gcLogHook)
	}

	selfLogHook, err := logger_util.NewFileHook(logger_conf.NfLogDir+"smf.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0o666)
	if err == nil {
		log.Hooks.Add(selfLogHook)
	}

	AppLog = log.WithFields(logrus.Fields{"component": "SMF", "category": "App"})
	InitLog = log.WithFields(logrus.Fields{"component": "SMF", "category": "Init"})
	CfgLog = log.WithFields(logrus.Fields{"component": "SMF", "category": "CFG"})
	PfcpLog = log.WithFields(logrus.Fields{"component": "SMF", "category": "PFCP"})
	PduSessLog = log.WithFields(logrus.Fields{"component": "SMF", "category": "PduSess"})
	GsmLog = log.WithFields(logrus.Fields{"component": "SMF", "category": "GSM"})
	CtxLog = log.WithFields(logrus.Fields{"component": "SMF", "category": "CTX"})
	ConsumerLog = log.WithFields(logrus.Fields{"component": "SMF", "category": "Consumer"})
	GinLog = log.WithFields(logrus.Fields{"component": "SMF", "category": "GIN"})

	JsonLog = logrus.New()
	JsonLog.SetReportCaller(false)
	JsonLog.SetFormatter(&logrus.JSONFormatter{})
	SliceLog = JsonLog.WithFields(logrus.Fields{"component": "SMF"})
	homeDirectory, err := os.UserHomeDir()
	if err == nil {
		file, err := os.OpenFile(homeDirectory+SMF_SLICE_LOG, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			SliceLog.Warnf("Error opening smf slice log")
		} else {
			JsonLog.SetOutput(file)
		}
	}

}

func SetLogLevel(level logrus.Level) {
	log.SetLevel(level)
}

func SetReportCaller(set bool) {
	log.SetReportCaller(set)
}
