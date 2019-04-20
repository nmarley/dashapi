#! /bin/sh

# Examples of using cURL to access the API. Replace URL here
# w/the actual URL of service deployment for actual usage.

# TODO:
# example of posting invalid data
#
# curl --header "Content-Type: application/json" --data '{"addr": "7", "msg": "hi", "sig": "yo"}' http://localhost:7001/proposal

# TODO:
# example of posting valid data
#
# curl --header "Content-Type: application/json" --data '{"hash":"7854ab3a9e48479372e749d220c3eab525d1ca12668b4eb782cb57e75555775a","collateralHash":"4d6df7d6d118df4a0228e24e6c53682028155fff9ab2a4822f517d74e5d567fa","createdAt":"2019-04-10T14:28:51-03:00","countYes":190,"countNo":0,"countAbstain":1,"startAt":"2019-04-10T14:28:51-03:00","endAt":"2019-06-10T14:28:51-03:00","name":"dev-coffee-supply","url":"https://ycvDJkTKDUS7pGY5DyknhhjGX2FvqJbmJt.com/","address":"ycvDJkTKDUS7pGY5DyknhhjGX2FvqJbmJt","amount":3.1415}' http://localhost:7001/proposal


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
