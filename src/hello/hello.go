package hello

import (
    "fmt"
    "github.com/gorilla/schema"
    "net/http"
    "timestamp"
)

func Timestamp(w http.ResponseWriter, r *http.Request) {
    type options struct {
        UTC bool        `schema:"utc"`
        Format string   `schema:"format"`
    }
    var params options
    decoder := schema.NewDecoder()
    err := decoder.Decode(&params, r.URL.Query())
    if err != nil {
        _, _ = w.Write([]byte(`Invalid request options!`))
        return
    }

    format := timestamp.AsFormat(params.Format)
    if format == timestamp.Unknown {
        format = timestamp.Verbose
    }
    timestamp.Options.UseUTC = params.UTC
    timestamp.Options.Representation = format
    _, _ = w.Write([]byte(fmt.Sprintf("Timestamp: %s", timestamp.Now())))
}