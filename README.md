# Conduit Connector Stripe

### General
haris: nitpick: not a builtin plugin
The Stripe connector is one of [Conduit](https://github.com/ConduitIO/conduit) builtin plugins. It provides a source Stripe connector.

### Prerequisites
- [Go](https://go.dev/) 1.18
- (optional) [golangci-lint](https://github.com/golangci/golangci-lint) 1.45.2
- (optional) [mock](https://github.com/golang/mock) 1.6.0

### Configuration
The config passed to `Configure` can contain the following fields:

| name        | description                                                                                                      | required | example                      |
|-------------|------------------------------------------------------------------------------------------------------------------|----------|------------------------------|
| `key`       | Stripe [secret key](https://dashboard.stripe.com/apikeys).                                                       | yes      | "sk_51Kr0QrJit566F2YtZAwMlh" |
| `resource`  | The name of Stripe resource. A list of supported resources can be found [here](models/resources/README.md). | yes      | "plan"                       |

[comment]: <> (maha: we prefer having all the configs as camelCase.. so this could be maxRetry, maxRetries or retryMax)
| `retry_max` | The maximum number of requests to Stripe in case of failure. By default is 3. The maximum is 10.                 | no       | "5"                          |                                                                                           | yes      | "id"                                            |
| `limit`     | Count of records in one butch. By default is 50. The maximum is 100.                                             | no       | "70"                         |
haris: typo above butch -> batch. also, what are the reasons for setting max values for retries and the batch?
haris: nitpick, maybe rename limit to batch size. limit (at least to me) sounds more like a global limit.

### How to build it
Run `make build`.

### Testing
Run `make test`.
haris: is it possible to have integration tests?

### Stripe Source
The `Configure` method parses the configuration and validates them.
The `Open` method parses the current position, initializes an 
[HTTP client](https://github.com/hashicorp/go-retryablehttp), and initializes Snapshot (only if in the position IteratorType equals Snapshot) and CDC iterators.
The `Read` method calls the method `Next` of the current iterator and returns the next record.

[comment]: <> (maha: still in development?)
The `Ack` method checks if the record with the position was recorded (under development).

**Note:** Data from Stripe is sorted by creation date in descending order, with no manual sort option.

### Snapshot
The system retrieves data from the list of objects of a defined Stripe resource (e.g. resource [plan](https://stripe.com/docs/api/plans/list)).

The system makes the first request to get a list of resource objects with the `limit` parameter from the configuration and appends the data to the resulting slice.

The data in the resulting slice is ready to be returned by the `Read` method record by record.

Each time the `Read` method is called, the system sets the current record ID to the `Cursor` to the position.

After all the data from the resulting slice has been read, the system makes the next request to Stripe with the additional `starting_after` parameter which is `Cursor`.

The system stops making requests when the field `has_next` equals false. Then, the system sets the `IteratorType` with CDC value and clears the `Cursor` value.

haris: descending order by which field? CreatedAt I guess?
**Note:** The Snapshot process creates a copy of the data in descending order.

### Change Data Capture
The system retrieves data from [Stripe events](https://api.stripe.com/v1/events) of the defined resource.

Because of this, receiving data is divided by requests to Stripe with different "offset" parameters:
- starting_after
- ending_before

Each time the `Read` method is called, the system sets the iteration index of resulting slice to the position and updates `Cursor` value when the last slice element is read.

#### starting_after
The system makes requests in the loop to get all the data since the first start of the pipeline `CreatedAt`.

After each request, the system adds the received data to the final slice and makes the next request with the starting_after parameter, which is the event ID of the last element from the response.

The system stops making requests when the field `has_next` equals false.

Then the system reverses the whole resulting slice and sets the Cursor ID of the "freshest" event.

The data in the resulting slice is ready to be returned by the `Read` method record by record.

#### ending_before
After all the data from the resulting slice has been read, the system takes the next batch of data from Stripe.

The system makes requests in the loop to get all data starting from the position of the freshest event element (Cursor) using the ending_before parameter.

After each request, the system sets a new value for the next request before_ending (the ID of the first event element from the response), reverses the response, and appends the data to the resulting slice.

The system stops making requests when the field `has_next` equals false.

The data in the resulting slice is ready to be returned by the `Read` method record by record.

All the following data are taken by the [ending_before](#ending_before) script.

### Position
Position has the following:

[comment]: <> (maha: maybe show an example of a position and how it's split into these)
| name            | type    | description                            |
|-----------------|---------|----------------------------------------|
| `IteratorType`  | string  | `s` - Snapshot (by default), `c` - CDC |
| `Cursor`        | string  | `id` of the resource or event          |
| `CreatedAt`     | int64   | UTC timestamp                          |
| `Index`         | int     | iteration position for cached data     |
