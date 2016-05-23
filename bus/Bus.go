// Handle communication between clients and the pit
package bus

import "google.golang.org/cloud/pubsub"
import "golang.org/x/net/context"
import "golang.org/x/oauth2/google"
import "google.golang.org/cloud"
import "google.golang.org/cloud/storage"
import "os"
import "fmt"

func OpenPit(projectId string, routingTopic string) string {
  context, err := cloudContext(projectId)
  if err != nil {
    fmt.Println("Error creating context", err)
    os.Exit(1)
  }

  pubsub.CreateTopic(context, routingTopic)
  return projectId
}

func cloudContext(projectId string) (context.Context, error) {
  ctx := context.Background()
	httpClient, err := google.DefaultClient(ctx, storage.ScopeFullControl, pubsub.ScopePubSub)
	if err != nil {
		return nil, err
	}
	return cloud.WithContext(ctx, projectId, httpClient), nil
}
