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
	DB_TRIGGER      = "go_trigger"
	DB_TRIGGER_TYPE = "go_trigger_type"
)

const (
	NOTICE_HTTP = iota
	NOTICE_NSQ
)
