// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// DeleteByQueryOption is a non-required DeleteByQuery option that gets applied to an HTTP request.
type DeleteByQueryOption func(r *transport.Request)

// WithDeleteByQueryType - a comma-separated list of document types to search; leave empty to perform the operation on all types.
func WithDeleteByQueryType(documentType []string) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQuerySource - true or false to return the _source field or not, or a list of fields to return.
func WithDeleteByQuerySource(source []string) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQuerySourceExclude - a list of fields to exclude from the returned _source field.
func WithDeleteByQuerySourceExclude(sourceExclude []string) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQuerySourceInclude - a list of fields to extract and return from the _source field.
func WithDeleteByQuerySourceInclude(sourceInclude []string) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (This includes "_all" string or when no indices have been specified).
func WithDeleteByQueryAllowNoIndices(allowNoIndices bool) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryAnalyzeWildcard - specify whether wildcard and prefix queries should be analyzed (default: false).
func WithDeleteByQueryAnalyzeWildcard(analyzeWildcard bool) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryAnalyzer - the analyzer to use for the query string.
func WithDeleteByQueryAnalyzer(analyzer string) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// DeleteByQueryConflicts - what to do when the delete-by-query hits version conflicts?
type DeleteByQueryConflicts int

const (
	// DeleteByQueryConflictsAbort can be used to set DeleteByQueryConflicts to "abort"
	DeleteByQueryConflictsAbort = iota
	// DeleteByQueryConflictsProceed can be used to set DeleteByQueryConflicts to "proceed"
	DeleteByQueryConflictsProceed = iota
)

// WithDeleteByQueryConflicts - what to do when the delete-by-query hits version conflicts?
func WithDeleteByQueryConflicts(conflicts DeleteByQueryConflicts) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// DeleteByQueryDefaultOperator - the default operator for query string query (AND or OR).
type DeleteByQueryDefaultOperator int

const (
	// DeleteByQueryDefaultOperatorAND can be used to set DeleteByQueryDefaultOperator to "AND"
	DeleteByQueryDefaultOperatorAND = iota
	// DeleteByQueryDefaultOperatorOR can be used to set DeleteByQueryDefaultOperator to "OR"
	DeleteByQueryDefaultOperatorOR = iota
)

// WithDeleteByQueryDefaultOperator - the default operator for query string query (AND or OR).
func WithDeleteByQueryDefaultOperator(defaultOperator DeleteByQueryDefaultOperator) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryDf - the field to use as default where no field prefix is given in the query string.
func WithDeleteByQueryDf(df string) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// DeleteByQueryExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
type DeleteByQueryExpandWildcards int

const (
	// DeleteByQueryExpandWildcardsOpen can be used to set DeleteByQueryExpandWildcards to "open"
	DeleteByQueryExpandWildcardsOpen = iota
	// DeleteByQueryExpandWildcardsClosed can be used to set DeleteByQueryExpandWildcards to "closed"
	DeleteByQueryExpandWildcardsClosed = iota
	// DeleteByQueryExpandWildcardsNone can be used to set DeleteByQueryExpandWildcards to "none"
	DeleteByQueryExpandWildcardsNone = iota
	// DeleteByQueryExpandWildcardsAll can be used to set DeleteByQueryExpandWildcards to "all"
	DeleteByQueryExpandWildcardsAll = iota
)

// WithDeleteByQueryExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
func WithDeleteByQueryExpandWildcards(expandWildcards DeleteByQueryExpandWildcards) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryFrom - starting offset (default: 0).
func WithDeleteByQueryFrom(from int) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func WithDeleteByQueryIgnoreUnavailable(ignoreUnavailable bool) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryLenient - specify whether format-based query failures (such as providing text to a numeric field) should be ignored.
func WithDeleteByQueryLenient(lenient bool) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryPreference - specify the node or shard the operation should be performed on (default: random).
func WithDeleteByQueryPreference(preference string) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryQ - query in the Lucene query string syntax.
func WithDeleteByQueryQ(q string) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryRefresh - should the effected indexes be refreshed?
func WithDeleteByQueryRefresh(refresh bool) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryRequestCache - specify if request cache should be used for this request or not, defaults to index level setting.
func WithDeleteByQueryRequestCache(requestCache bool) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryRequestsPerSecond - the throttle for this request in sub-requests per second. -1 means no throttle.
func WithDeleteByQueryRequestsPerSecond(requestsPerSecond int) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryRouting - a comma-separated list of specific routing values.
func WithDeleteByQueryRouting(routing []string) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryScroll - specify how long a consistent view of the index should be maintained for scrolled search.
func WithDeleteByQueryScroll(scroll time.Duration) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryScrollSize - size on the scroll request powering the update_by_query.
func WithDeleteByQueryScrollSize(scrollSize int) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQuerySearchTimeout - explicit timeout for each search request. Defaults to no timeout.
func WithDeleteByQuerySearchTimeout(searchTimeout time.Duration) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// DeleteByQuerySearchType - search operation type.
type DeleteByQuerySearchType int

const (
	// DeleteByQuerySearchTypeQueryThenFetch can be used to set DeleteByQuerySearchType to "query_then_fetch"
	DeleteByQuerySearchTypeQueryThenFetch = iota
	// DeleteByQuerySearchTypeDfsQueryThenFetch can be used to set DeleteByQuerySearchType to "dfs_query_then_fetch"
	DeleteByQuerySearchTypeDfsQueryThenFetch = iota
)

// WithDeleteByQuerySearchType - search operation type.
func WithDeleteByQuerySearchType(searchType DeleteByQuerySearchType) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQuerySize - number of hits to return (default: 10).
func WithDeleteByQuerySize(size int) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQuerySlices - the number of slices this task should be divided into. Defaults to 1 meaning the task isn't sliced into subtasks.
func WithDeleteByQuerySlices(slices int) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQuerySort - a comma-separated list of <field>:<direction> pairs.
func WithDeleteByQuerySort(sort []string) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryStats - specific 'tag' of the request for logging and statistical purposes.
func WithDeleteByQueryStats(stats []string) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryTerminateAfter - the maximum number of documents to collect for each shard, upon reaching which the query execution will terminate early.
func WithDeleteByQueryTerminateAfter(terminateAfter int) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryTimeout - time each individual bulk request should wait for shards that are unavailable.
func WithDeleteByQueryTimeout(timeout time.Duration) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryVersion - specify whether to return document version as part of a hit.
func WithDeleteByQueryVersion(version bool) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryWaitForActiveShards - sets the number of shard copies that must be active before proceeding with the delete by query operation. Defaults to 1, meaning the primary shard only. Set to "all" for all shard copies, otherwise set to any non-negative value less than or equal to the total number of copies for the shard (number of replicas + 1).
func WithDeleteByQueryWaitForActiveShards(waitForActiveShards string) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryWaitForCompletion - should the request should block until the delete-by-query is complete.
func WithDeleteByQueryWaitForCompletion(waitForCompletion bool) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryErrorTrace - include the stack trace of returned errors.
func WithDeleteByQueryErrorTrace(errorTrace bool) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryFilterPath - a comma-separated list of filters used to reduce the respone.
func WithDeleteByQueryFilterPath(filterPath []string) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryHuman - return human readable values for statistics.
func WithDeleteByQueryHuman(human bool) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryIgnore - ignores the specified HTTP status codes.
func WithDeleteByQueryIgnore(ignore []int) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQueryPretty - pretty format the returned JSON response.
func WithDeleteByQueryPretty(pretty bool) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// WithDeleteByQuerySourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithDeleteByQuerySourceParam(sourceParam string) DeleteByQueryOption {
	return func(r *transport.Request) {
	}
}

// DeleteByQuery - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/docs-delete-by-query.html for more info.
//
// index: a comma-separated list of index names to search; use "_all" or empty string to perform the operation on all indices.
//
// body: the search definition using the Query DSL.
//
// options: optional parameters.
func (a *API) DeleteByQuery(index []string, body map[string]interface{}, options ...DeleteByQueryOption) (*DeleteByQueryResponse, error) {
	req := a.transport.NewRequest("POST")
	for _, option := range options {
		option(req)
	}
	resp, err := a.transport.Do(req)
	return &DeleteByQueryResponse{resp}, err
}

// DeleteByQueryResponse is the response for DeleteByQuery.
type DeleteByQueryResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *DeleteByQueryResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}