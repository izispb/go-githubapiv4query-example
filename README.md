# go-githubapiv4query-example
GitHub API v4 on Go (oauth2 and githubv4) - sort repo by stars example 

query MyQuery {
	search(type: REPOSITORY, first: 5, query: "stars:>0 sort:stars-desc") {
	  nodes {
		... on Repository {
		  stargazers {
			totalCount
		  }
		  nameWithOwner
		}
	  }
	}
  }
  
Output:
1: repo/ovner/stars: freeCodeCamp/freeCodeCamp/307316
2: repo/ovner/stars: 996icu/996.ICU/248270
3: repo/ovner/stars: vuejs/vue/153590
4: repo/ovner/stars: facebook/react/140578
5: repo/ovner/stars: tensorflow/tensorflow/138649
