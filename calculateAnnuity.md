curl -X POST \
  http://localhost:8080/generate-plan \
  -H 'content-type: application/json' \
  -d '{
"loanAmount": "5000",
"nominalRate": "5.0",
"duration": 24,
"startDate": "2018-01-01T00:00:01Z"
}'