package client

import (
	"fmt"
	"strconv"

	"github.com/go-zoox/fetch"
)

// Resource is a REST API resource.
type Resource[T any] struct {
	client    *Client
	namespace string
}

// List returns a resource item list.
func (r *Resource[T]) List(page int, pageSize int, query ...map[string]string) ([]T, error) {
	q := map[string]string{}
	if len(query) > 0 && query[0] != nil {
		for k, v := range query[0] {
			q[k] = v
		}
	}
	q["page"] = strconv.Itoa(page)
	q["pageSize"] = strconv.Itoa(pageSize)

	path := r.buildPath("List")
	response, err := fetch.Get(path, &fetch.Config{
		Query:   r.buildQuery(q),
		Headers: r.buildHeaders(),
	})
	if err != nil {
		return nil, err
	}

	if response.Status != 200 {
		return nil, r.buildError(response)
	}

	var result []T
	if err = response.UnmarshalJSON(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// Create creates a resource item.
func (r *Resource[T]) Create(mutation *T) error {
	path := r.buildPath("Create")
	response, err := fetch.Post(path, &fetch.Config{
		Headers: r.buildHeaders(),
		Body:    mutation,
	})
	if err != nil {
		return err
	}

	if response.Status != 201 {
		return r.buildError(response)
	}

	return nil
}

// Retrieve gets a resource item.
func (r *Resource[T]) Retrieve(id string) (*T, error) {
	path := r.buildPath("Retrieve")
	response, err := fetch.Get(path, &fetch.Config{
		Headers: r.buildHeaders(),
		Params:  map[string]string{"id": id},
	})
	var result T
	if err != nil {
		return &result, err
	}

	if response.Status != 200 {
		return &result, r.buildError(response)
	}

	if err = response.UnmarshalJSON(&result); err != nil {
		return &result, err
	}

	return &result, nil
}

// Update updates a resource item by id.
func (r *Resource[T]) Update(id string, mutation *T) error {
	path := r.buildPath("Update")
	response, err := fetch.Put(path, &fetch.Config{
		Headers: r.buildHeaders(),
		Params:  map[string]string{"id": id},
		Body:    mutation,
	})
	if err != nil {
		return err
	}

	if response.Status != 200 {
		return r.buildError(response)
	}

	return nil
}

// Delete deletes a resource item by id.
func (r *Resource[T]) Delete(id string) error {
	path := r.buildPath("Delete")
	response, err := fetch.Delete(path, &fetch.Config{
		Headers: r.buildHeaders(),
		Query:   r.buildQuery(map[string]string{"id": id}),
	})
	if err != nil {
		return err
	}

	if response.Status != 204 {
		return r.buildError(response)
	}

	return nil
}

func (r *Resource[T]) buildPath(typ string) string {
	switch typ {
	case "List":
		return fmt.Sprintf("%s/%s", r.client.Endpoint, r.namespace)
	case "Create":
		return fmt.Sprintf("%s/%s", r.client.Endpoint, r.namespace)
	case "Retrieve":
		return fmt.Sprintf("%s/%s/:id", r.client.Endpoint, r.namespace)
	case "Update":
		return fmt.Sprintf("%s/%s/:id", r.client.Endpoint, r.namespace)
	case "Delete":
		return fmt.Sprintf("%s/%s/:id", r.client.Endpoint, r.namespace)
	}

	return ""
}

func (r *Resource[T]) buildQuery(current map[string]string) map[string]string {
	if r.client.Config.Query == nil {
		return current
	}

	// merge
	merged := map[string]string{}
	for k, v := range r.client.Config.Query {
		merged[k] = v
	}

	for k, v := range current {
		merged[k] = v
	}

	return merged
}

func (r *Resource[T]) buildHeaders() map[string]string {
	// merge
	merged := map[string]string{
		"content-type": "application/json",
	}

	if r.client.Config.Headers == nil {
		return merged
	}

	for k, v := range r.client.Config.Headers {
		merged[k] = v
	}

	return merged
}

func (r *Resource[T]) buildError(response *fetch.Response) error {
	code := response.Get("code").Int()
	message := response.Get("message").String()
	if code == 0 {
		code = int64(response.Status)
	}
	if message == "" {
		message = response.String()
	}

	return fmt.Errorf("[%d] %s", code, message)
}
