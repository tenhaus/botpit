// Manages users and permissions

package auth

import "os"
import "fmt"
import "github.com/tenhaus/botpit/cloud"
import "golang.org/x/net/context"

// So I think we want to:
//
// Create a uuid for a user (probably pull it from a db)
// Create a service account with the uuid
// Create a topic with the uuid
// Somehow retrieve some token the user can use to authenticate
// Send the token back

func CreateUser(userHandle string, projectId string) (uuid string, error) {
  context, err := cloud.CloudContext(projectId)

  if err != nil {
    fmt.Println("Error creating context", err)
    return nil, err
  }

  return "this will be a uuid from somewhere", nil
}

func CreateServiceAccount(uuid string) {
}

func CreateTopic(context context.Context, uuid string) {
  pubsub.CreateTopic(context, uuid)
}