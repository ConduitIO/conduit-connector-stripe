# Conduit Connector Stripe

### General
The Stripe connector is one of [Conduit](https://github.com/ConduitIO/conduit) plugins. It provides a source Stripe connector.

### Prerequisites
- [Go](https://go.dev/) 1.18
- (optional) [golangci-lint](https://github.com/golangci/golangci-lint) 1.45.2
- (optional) [mock](https://github.com/golang/mock) 1.6.0

### Configuration
The config passed to `Configure` can contain the following fields:

| name            | description                                                                                                 | required | example                      |
|-----------------|-------------------------------------------------------------------------------------------------------------|----------|------------------------------|
| `secretKey`     | Stripe [secret key](https://dashboard.stripe.com/apikeys).                                                  | yes      | "sk_51Kr0QrJit566F2YtZAwMlh" |
| `resourceName`  | The name of Stripe resource. A list of supported resources can be found [here](models/resources/README.md). | yes      | "plan"                       |

### How to build it
Run `make build`.

### Testing
Run `make test`.

### Stripe Source
The `Configure` method parses the configuration and validates them.

The `Open` method parses the current position, initializes an [HTTP client](https://github.com/hashicorp/go-retryablehttp), and initializes Snapshot (only if in the position IteratorType equals Snapshot) and CDC iterators.

The `Read` method calls the method `Next` of the current iterator and returns the next record.

The `Teardown` method calls the method `Close` of the http client, which calls `CloseIdleConnections` method of the [net/http](https://pkg.go.dev/net/http) package.

Stripe source connector supports Change data capture (CDC) process. 

### Snapshot
The `Snapshot` iterator runs first and receives data from the list of records of a defined Stripe resource (e.g. [list of a plan resource](https://stripe.com/docs/api/plans/list)).

**Note**: The `Snapshot` iterator creates a copy of the data, which is sorted by date of creation in descending order.

### Change Data Capture
The `CDC` iterator runs after the `Snapshot` and receives data from [Stripe events](https://api.stripe.com/v1/events) of the defined resource since the connector was set up.

### Position
Position is a JSON object with the following fields:

| name            | type    | description                                                                                                                                                         | 
|-----------------|---------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `IteratorType`  | string  | type of iterator (`snapshot`, `cdc`)                                                                                                                                |
| `CreatedAt`     | int64   | unix time from which the system should receive events of the resource in the CDC iterator (the parameter is set with the present time when the Position is created) |
| `Cursor`        | string  | resource or event identifier for receiving shifted data in the following requests                                                                                   |
| `Index`         | int     | current index of the returning record from the batch of previously received resources                                                                               |
Example:
```json
{
	"iterator_type":"cdc",
	"created_at":1652279623,
	"cursor":"evt_1KtXkmJit567F2YtZzGSIrsh",
	"index": 1
}
```

### Notes
- Data from Stripe is sorted by date of creation in descending order, with no manual sort option.
- If the user changes the resource name when the connector was created and was already in use, the connector will start working with the new resource from the beginning.