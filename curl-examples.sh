#! /bin/sh

# Examples of using cURL to access the API. Replace URL here
# w/the actual URL of service deployment for actual usage.

# TODO:
# example of posting invalid data
#
# curl --header "Content-Type: application/json" --data '{"addr": "7", "msg": "hi", "sig": "yo"}' http://localhost:7001/vote

# TODO:
# example of posting valid data
#
# curl --header "Content-Type: application/json" --data '{"end_epoch":1560187731, "name":"developer-salary", "payment_address": "yVQCPZ2kW6FyPguUriKRRLuBd1WqGbSgPR", "payment_amount":2, "start_epoch":1554917331, "type":1, "url":"https://yVQCPZ2kW6FyPguUriKRRLuBd1WqGbSgPR.com/"}' http://localhost:7001/proposal


# TODO:
# example listing valid proposals
#
# "valid" votes are those which are not superceded by any newer vote for the
# same MNO collateral address
#
# note: JWT_TOKEN must be set to a valid, signed token
#
# curl --silent --header "Authorization: Bearer $JWT_TOKEN" http://localhost:7001/validVotes

# TODO:
# example listing all proposals
#
# note: JWT_TOKEN must be set to a valid, signed token
#
# curl --silent --header "Authorization: Bearer $JWT_TOKEN" http://localhost:7001/allProposals

# TODO:
# * triggers
