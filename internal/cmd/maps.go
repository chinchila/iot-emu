package cmd

var COMMANDS_FUNCTION_MAP = map[string]interface{}{
	EXIT_LITERAL:   Exit,
	ADD_LITERAL:    Add,
	INFO_LITERAL:   Info,
	STATUS_LITERAL: Info,
	REM_LITERAL:    Rem,
	REMOVE_LITERAL: Rem,

	START_LITERAL: Start,
	STOP_LITERAL:  Stop,

	TADD_LITERAL:      ThingAdd,
	THING_ADD_LITERAL: ThingAdd,
}

var COMMANDS_USAGE_MAP = map[string]string{
	EXIT_LITERAL:   EXIT_USAGE,
	ADD_LITERAL:    ADD_USAGE,
	INFO_LITERAL:   INFO_USAGE,
	STATUS_LITERAL: STATUS_USAGE,
	REM_LITERAL:    REM_USAGE,
	REMOVE_LITERAL: REM_USAGE,

	START_LITERAL: START_USAGE,
	STOP_LITERAL:  STOP_USAGE,

	TADD_LITERAL:      TADD_USAGE,
	THING_ADD_LITERAL: TADD_USAGE,
}
