# logging

Package can be used for logging in json format

Usage:
------
    $ go get github.com/oleglpts/logging

    import (
        "github.com/oleglpts/logging"
        "github.com/ztrue/tracerr"
    )

    var __ = logging.GetMessage
    var ___ = logging.GetExtendedMessage
    ...........................................
    logging.Init("notificator", logging.INFO)
    log.Print(__(logging.DEBUG, "Start cycle"))
    ...........................................
    if err != nil {
        log.Print(___(logging.FATAL, err.Error(), map[string]string{},
                  map[string]string{}, "1", "database error",
                  tracerr.Sprint(tracerr.Wrap(err))))
    }
