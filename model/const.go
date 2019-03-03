package model

const (
	ENV_AGENT_MONGO_ENDPOINT = "MONGO_ENDPOINT"
	ENV_AGENT_MONGO_NAME     = "MONGO_USER_NAME"
	ENV_AGENT_MONGO_PASSWD   = "MONGO_USER_PASSWD"
	ENV_AGENT_MONGO_DBNAME   = "MONGO_DB_NAME"
)

const (
	MODULENAME = "GoTrigger"
)

const (
	DB_TRIGGER        = "go_trigger"
	DB_TRIGGER_TYPE   = "go_trigger_type"
	DB_TRIGGER_PLUGIN = "go_trigger_plugin"
	DB_PLUGIN_KIND    = "go_trigger_plugin_kind"
	DB_SEQUENCE       = "go_trigger_sequence"
	DB_TRIGGER_JNL    = "go_trigger_jnl"
)

const (
	NOTICE_HTTP = iota
	NOTICE_NSQ
)

const (
	STATUS_INVOKE_SUCC = iota
	STATUS_INVOKE_FAILED
)

