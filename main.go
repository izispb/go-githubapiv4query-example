package main

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// query MyQuery {
// 	search(type: REPOSITORY, first: 5, query: "stars:>0 sort:stars-desc") {
// 	  nodes {
// 		... on Repository {
// 		  stargazers {
// 			totalCount
// 		  }
// 		  nameWithOwner
// 		}
// 	  }
// 	}
//   }

var query struct {
	Search struct {
		// RepositoryCount int
		Nodes []struct {
			Repository struct {
				NameWithOwner string
				Stargazers    struct {
					TotalCount int
				}
			} `graphql:"... on Repository"`
		}
	} `graphql:"search (type: REPOSITORY, query: \"stars:>0 sort:stars-desc\", first:5)"`
}

// func print(v interface{}) {
// 	w := json.NewEncoder(os.Stdout)
// 	w.SetIndent("", "\t")
// 	err := w.Encode(v)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "86751fc6974b2866029bea464d6e360465d181aa"},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	err := client.Query(context.Background(), &query, nil)

	if err != nil {
	}

	for i, n := range query.Search.Nodes {

		fmt.Printf("%v: repo/ovner/stars: %v/%v\n", i+1, n.Repository.NameWithOwner, n.Repository.Stargazers.TotalCount)
	}
	// fmt.Println("query:", query.Search.Nodes)
	// print(query)
}
