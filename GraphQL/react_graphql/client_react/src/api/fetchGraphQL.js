
async function fetchGraphQL(gqlQuery, params) {

  const endpoint = 'http://localhost:8080/graphql'

  const response = await fetch(endpoint, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
    },
    body: JSON.stringify({
      query: gqlQuery,
      variables: params
    })
  });
  return await response.json();
}

export default fetchGraphQL;