# Task by greg-agacinski 

https://gist.github.com/greg-agacinski/9e70563c3fb88514fe99fe7dc787f921#file-backend-md

## Run Locally

You can use this command to run project using docker:

```bash
  docker run -p 8080:8080 -d -e API_KEY= You can get api key from random.org

```

## Description of challenge
Create a simple REST service in Go supporting the following GET operation:
```
/random/mean?requests={r}&length={l}
```
which performs `{r}` concurrent requests to [random.org](https://random.org) API asking for `{l}` number of random integers.

For each of `{r}` requests, the service must calculate standard deviation of the drawn integers and additionally standard deviation of sum of all sets.
Results must be presented in JSON format.

## Example
`GET /random/mean?requests=2&length=5`

Response:
```json
[
  {
     "stddev": 1,
     "data": [1, 2, 3, 4, 5]
  },
  {
     "stddev": 1,
     "data": [1, 2, 3, 4, 5]
  },
  { // stddev of sum
     "stddev": 1,
     "data": [1, 1, 2, 2, 3, 3, 4, 4, 5, 5]
  }
]
```

## Requirements
1. Proper error handling when communicating with external service (timeouts, invalid statuses).
2. Usage of contexts.
3. Solution should be delivered as a git repository.
4. Provide a `.Dockerfile` that builds the image with the application.
5. Application should run flawlessly after running `docker build ...` & `docker run ...` commands.


## 🔗 Links
[![portfolio](https://img.shields.io/badge/my_portfolio-000?style=for-the-badge&logo=ko-fi&logoColor=white)](https://github.com/fbrv01/)
[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/filip-bucholc/)

