package httpserver

import (
	// stdlib
	"encoding/json"
	"fmt"
	"net/http"

	// other
	"github.com/labstack/echo"
)

// StrictJSONBinder implements Binder interface for Echo. It will parse
// JSON in strict mode throwing errors on schema mismatches.
type StrictJSONBinder struct{}

// Bind parses JSON input.
func (sjb *StrictJSONBinder) Bind(i interface{}, c echo.Context) error {
	req := c.Request()
	if req.ContentLength == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Request body can't be empty")
	}

	// Decode it.
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(i); err != nil {
		if ute, ok := err.(*json.UnmarshalTypeError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset))
		} else if se, ok := err.(*json.SyntaxError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error()))
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	return nil
}
