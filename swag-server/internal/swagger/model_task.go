/*
 * Sample REST server
 *
 * TODO
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type Task struct {

	Id int32 `json:"id,omitempty"`

	Text string `json:"text,omitempty"`

	Tags []string `json:"tags,omitempty"`

	Due time.Time `json:"due,omitempty"`
}
